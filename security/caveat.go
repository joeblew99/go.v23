// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"

	"v.io/v23/context"
	"v.io/v23/uniqueid"
	"v.io/v23/vdl"
	"v.io/v23/verror"
	"v.io/v23/vom"
)

var (
	errCaveatRegisteredTwice          = verror.Register(pkgPath+".errCaveatRegisteredTwice", verror.NoRetry, "{1:}{2:}Caveat with UUID {3} registered twice. Once with ({4}, fn={5}) from {6}, once with ({7}, fn={8}) from {9}{:_}")
	errBadCaveatDescriptorType        = verror.Register(pkgPath+".errBadCaveatDescriptorType", verror.NoRetry, "{1:}{2:}invalid caveat descriptor: vdl.Type({3}) cannot be converted to a Go type{:_}")
	errBadCaveatDescriptorKind        = verror.Register(pkgPath+".errBadCaveatDescriptorKind", verror.NoRetry, "{1:}{2:}invalid caveat validator: must be {3}, not {4}{:_}")
	errBadCaveatOutputNum             = verror.Register(pkgPath+".errBadCaveatOutputNum", verror.NoRetry, "{1:}{2:}invalid caveat validator: expected {3} outputs, not {4}{:_}")
	errBadCaveatOutput                = verror.Register(pkgPath+".errBadCaveatOutput", verror.NoRetry, "{1:}{2:}invalid caveat validator: output must be {3}, not {4}{:_}")
	errBadCaveatInputs                = verror.Register(pkgPath+".errBadCaveatInputs", verror.NoRetry, "{1:}{2:}invalid caveat validator: expected {3} inputs, not {4}{:_}")
	errBadCaveat1stArg                = verror.Register(pkgPath+".errBadCaveat1stArg", verror.NoRetry, "{1:}{2:}invalid caveat validator: first argument must be {3}, not {4}{:_}")
	errBadCaveat2ndArg                = verror.Register(pkgPath+".errBadCaveat2ndArg", verror.NoRetry, "{1:}{2:}invalid caveat validator: second argument must be {3}, not {4}{:_}")
	errBadCaveatRestriction           = verror.Register(pkgPath+".errBadCaveatRestriction", verror.NoRetry, "{1:}{2:}could not validate embedded restriction({3}): {4}{:_}")
	errCantUnmarshalDischargeKey      = verror.Register(pkgPath+".errCantUnmarshalDischargeKey", verror.NoRetry, "{1:}{2:}invalid {3}: failed to unmarshal discharger's public key: {4}{:_}")
	errInapproriateDischargeSignature = verror.Register(pkgPath+".errInapproriateDischargeSignature", verror.NoRetry, "{1:}{2:}signature on discharge for caveat {3} was not intended for discharges(purpose={4}){:_}")
	errBadDischargeSignature          = verror.Register(pkgPath+".errBadDischargeSignature", verror.NoRetry, "{1:}{2:}signature verification on discharge for caveat {3} failed{:_}")
)

type registryEntry struct {
	desc        CaveatDescriptor
	validatorFn reflect.Value
	paramType   reflect.Type
	registerer  string
}

// Instance of unconstrained use caveat, to be used by UnconstrainedCaveat().
var unconstrainedUseCaveat Caveat

func init() {
	var err error
	unconstrainedUseCaveat, err = NewCaveat(ConstCaveat, true)
	if err != nil {
		panic(fmt.Sprintf("Error in NewCaveat: %v", err))
	}
}

// caveatRegistry is used to implement a singleton global registry that maps
// the unique id of a caveat to its validation function.
//
// It is safe to invoke methods on caveatRegistry concurrently.
type caveatRegistry struct {
	mu     sync.RWMutex
	byUUID map[uniqueid.Id]registryEntry
}

var registry = &caveatRegistry{byUUID: make(map[uniqueid.Id]registryEntry)}

func (r *caveatRegistry) register(d CaveatDescriptor, validator interface{}) error {
	_, file, line, _ := runtime.Caller(2) // one for r.register, one for RegisterCaveatValidator
	registerer := fmt.Sprintf("%s:%d", file, line)
	r.mu.Lock()
	defer r.mu.Unlock()
	if e, exists := r.byUUID[d.Id]; exists {
		return verror.New(errCaveatRegisteredTwice, nil, d.Id, e.desc.ParamType, e.validatorFn.Interface(), e.registerer, d.ParamType, validator, registerer)
	}
	fn := reflect.ValueOf(validator)
	param := vdl.TypeToReflect(d.ParamType)
	if param == nil {
		// If you hit this error, https://github.com/veyron/release-issues/issues/907
		// might be the problem.
		return verror.New(errBadCaveatDescriptorType, nil, d.ParamType)
	}
	var (
		rtErr = reflect.TypeOf((*error)(nil)).Elem()
		rtCtx = reflect.TypeOf((*context.T)(nil))
	)
	if got, want := fn.Kind(), reflect.Func; got != want {
		return verror.New(errBadCaveatDescriptorKind, nil, want, got)
	}
	if got, want := fn.Type().NumOut(), 1; got != want {
		return verror.New(errBadCaveatOutputNum, nil, want, got)
	}
	if got, want := fn.Type().Out(0), rtErr; got != want {
		return verror.New(errBadCaveatOutput, nil, want, got)
	}
	if got, want := fn.Type().NumIn(), 2; got != want {
		return verror.New(errBadCaveatInputs, nil, want, got)
	}
	if got, want := fn.Type().In(0), rtCtx; got != want {
		return verror.New(errBadCaveat1stArg, nil, want, got)
	}
	if got, want := fn.Type().In(1), param; got != want {
		return verror.New(errBadCaveat2ndArg, nil, want, got)
	}
	r.byUUID[d.Id] = registryEntry{d, fn, param, registerer}
	return nil
}

func (r *caveatRegistry) lookup(uid uniqueid.Id) (registryEntry, bool) {
	r.mu.RLock()
	entry, exists := r.byUUID[uid]
	r.mu.RUnlock()
	return entry, exists
}

func (r *caveatRegistry) validate(uid uniqueid.Id, ctx *context.T, paramvom []byte) error {
	entry, exists := r.lookup(uid)
	if !exists {
		return NewErrCaveatNotRegistered(ctx, uid)
	}
	param := reflect.New(entry.paramType).Interface()
	if err := vom.Decode(paramvom, param); err != nil {
		t, _ := vdl.TypeFromReflect(entry.paramType)
		return NewErrCaveatParamCoding(ctx, uid, t, err)
	}
	err := entry.validatorFn.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(param).Elem()})[0].Interface()
	if err == nil {
		return nil
	}
	return NewErrCaveatValidation(ctx, err.(error))
}

// RegisterCaveatValidator associates a CaveatDescriptor with the
// implementation of the validation function.
//
// The validation function must act as if the caveat was obtained from the
// remote end of the call. In particular, if the caveat is a third-party
// caveat then 'call.RemoteDischarges()' must be used to validate it.
//
// This function must be called at most once per c.ID, and will panic on duplicate
// registrations.
func RegisterCaveatValidator(c CaveatDescriptor, validator interface{}) {
	if err := registry.register(c, validator); err != nil {
		panic(err)
	}
}

// NewCaveat returns a Caveat that requires validation by the validation
// function correponding to c and uses the provided parameters.
func NewCaveat(c CaveatDescriptor, param interface{}) (Caveat, error) {
	got := vdl.TypeOf(param)
	// If the user inputs a vdl.Value, use the type of the vdl.Value instead.
	if vv, ok := param.(*vdl.Value); ok {
		got = vv.Type()
	}
	noAnyInParam := c.ParamType.Walk(vdl.WalkAll, func(t *vdl.Type) bool {
		return t.Kind() != vdl.Any
	})
	if !noAnyInParam {
		return Caveat{}, NewErrCaveatParamAny(nil, c.Id)
	}
	if want := c.ParamType; got != want {
		return Caveat{}, NewErrCaveatParamTypeMismatch(nil, c.Id, got, want)
	}
	bytes, err := vom.Encode(param)
	if err != nil {
		return Caveat{}, NewErrCaveatParamCoding(nil, c.Id, c.ParamType, err)
	}
	return Caveat{c.Id, bytes}, nil
}

// ExpiryCaveat returns a Caveat that validates iff the current time is before t.
func ExpiryCaveat(t time.Time) (Caveat, error) {
	c, err := NewCaveat(ExpiryCaveatX, t)
	if err != nil {
		return c, err
	}
	return c, nil
}

// MethodCaveat returns a Caveat that validates iff the method being invoked by
// the peer is listed in an argument to this function.
func MethodCaveat(method string, additionalMethods ...string) (Caveat, error) {
	c, err := NewCaveat(MethodCaveatX, append(additionalMethods, method))
	if err != nil {
		return c, err
	}
	return c, nil
}

// digest returns a hash of the contents of c.
func (c *Caveat) digest(hash Hash) []byte {
	return hash.sum(append(hash.sum(c.Id[:]), hash.sum(c.ParamVom)...))
}

// Validate tests if 'c' is satisfied under 'call', returning nil if it is or an
// error otherwise.
//
// It assumes that 'c' was found on a credential obtained from the remote end of
// the call. In particular, if 'c' is a third-party caveat then it uses the
// call.RemoteDischarges() to validate it.
func (c *Caveat) Validate(ctx *context.T) error {
	return registry.validate(c.Id, ctx, c.ParamVom)
}

// ThirdPartyDetails returns nil if c is not a third party caveat, or details about
// the third party otherwise.
func (c *Caveat) ThirdPartyDetails() ThirdPartyCaveat {
	if c.Id == PublicKeyThirdPartyCaveatX.Id {
		var param publicKeyThirdPartyCaveat
		if err := vom.Decode(c.ParamVom, &param); err != nil {
			// TODO(jsimsa): Decide what (if any) logging mechanism to use.
			// vlog.Errorf("Error decoding PublicKeyThirdPartyCaveat: %v", err)
		}
		return &param
	}
	return nil
}

func (c Caveat) String() string {
	var param interface{}
	if err := vom.Decode(c.ParamVom, &param); err == nil {
		return fmt.Sprintf("%v(%T=%v)", c.Id, param, param)
	}
	return fmt.Sprintf("%v(%d bytes of param)", c.Id, len(c.ParamVom))
}

// UnconstrainedUse returns a Caveat implementation that never fails to
// validate. This is useful only for providing unconstrained
// blessings/discharges to another principal.
func UnconstrainedUse() Caveat {
	return unconstrainedUseCaveat
}

// NewPublicKeyCaveat returns a third-party caveat, i.e., the returned
// Caveat will be valid only when a discharge signed by discharger
// is issued.
//
// Location specifies the expected address at which the third-party
// service is found (and which issues discharges).
//
// The discharger will validate all provided caveats (caveat,
// additionalCaveats) before issuing a discharge.
func NewPublicKeyCaveat(discharger PublicKey, location string, requirements ThirdPartyRequirements, caveat Caveat, additionalCaveats ...Caveat) (Caveat, error) {
	param := publicKeyThirdPartyCaveat{
		Caveats:                append(additionalCaveats, caveat),
		DischargerLocation:     location,
		DischargerRequirements: requirements,
	}
	var err error
	if param.DischargerKey, err = discharger.MarshalBinary(); err != nil {
		return Caveat{}, err
	}
	if _, err := rand.Read(param.Nonce[:]); err != nil {
		return Caveat{}, err
	}
	c, err := NewCaveat(PublicKeyThirdPartyCaveatX, param)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c *publicKeyThirdPartyCaveat) ID() string {
	key, err := c.discharger(nil)
	if err != nil {
		// TODO(jsimsa): Decide what (if any) logging mechanism to use.
		// vlog.Error(err)
		return ""
	}
	hash := key.hash()
	bytes := append(hash.sum(c.Nonce[:]), hash.sum(c.DischargerKey)...)
	for _, cav := range c.Caveats {
		bytes = append(bytes, cav.digest(hash)...)
	}
	return base64.StdEncoding.EncodeToString(hash.sum(bytes))
}

func (c *publicKeyThirdPartyCaveat) Location() string { return c.DischargerLocation }
func (c *publicKeyThirdPartyCaveat) Requirements() ThirdPartyRequirements {
	return c.DischargerRequirements
}

func (c *publicKeyThirdPartyCaveat) Dischargeable(ctx *context.T) error {
	// Validate the caveats embedded within this third-party caveat.
	for _, cav := range c.Caveats {
		if err := cav.Validate(ctx); err != nil {
			return verror.New(errBadCaveatRestriction, ctx, cav, err)
		}
	}
	return nil
}

func (c *publicKeyThirdPartyCaveat) discharger(cxt *context.T) (PublicKey, error) {
	key, err := UnmarshalPublicKey(c.DischargerKey)
	if err != nil {
		return nil, verror.New(errCantUnmarshalDischargeKey, cxt, fmt.Sprintf("%T", *c), err)
	}
	return key, nil
}

func (c publicKeyThirdPartyCaveat) String() string {
	return fmt.Sprintf("%v@%v [%+v]", c.ID(), c.Location(), c.Requirements())
}

func (d *publicKeyDischarge) digest(hash Hash) []byte {
	msg := hash.sum([]byte(d.ThirdPartyCaveatId))
	for _, cav := range d.Caveats {
		msg = append(msg, cav.digest(hash)...)
	}
	return hash.sum(msg)
}

func (d *publicKeyDischarge) verify(cxt *context.T, key PublicKey) error {
	if !bytes.Equal(d.Signature.Purpose, dischargePurpose) {
		return verror.New(errInapproriateDischargeSignature, cxt, d.ThirdPartyCaveatId, d.Signature.Purpose)
	}
	if !d.Signature.Verify(key, d.digest(key.hash())) {
		return verror.New(errBadDischargeSignature, cxt, d.ThirdPartyCaveatId)
	}
	return nil
}

func (d *publicKeyDischarge) sign(signer Signer) error {
	var err error
	d.Signature, err = signer.Sign(dischargePurpose, d.digest(signer.PublicKey().hash()))
	return err
}

func (d *publicKeyDischarge) String() string {
	return fmt.Sprint(*d)
}
