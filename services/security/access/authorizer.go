package access

import (
	"os"
	"reflect"

	"v.io/v23/security"
	"v.io/v23/vdl"
	"v.io/v23/verror"
)

const pkgPath = "v.io/v23/services/security/access"

var (
	errTagNeedsString      = verror.Register(pkgPath+".errTagNeedsString", verror.NoRetry, "{1:}{2:}tag type({3}) must be backed by a string not {4}{:_}")
	errNoMethodTags        = verror.Register(pkgPath+".errNoMethodTags", verror.NoRetry, "{1:}{2:}TaggedACLAuthorizer.Authorize called with an object ({3}, method {4}) that has no method tags; this is likely unintentional{:_}")
	errCantReadACLFromFile = verror.Register(pkgPath+".errCantReadACLFromFile", verror.NoRetry, "{1:}{2:}failed to read ACL from file{:_}")
)

// TaggedACLAuthorizer implements an authorization policy where access is
// granted if the remote end presents blessings included in the Access Control
// Lists (ACLs) associated with the set of relevant tags.
//
// The set of relevant tags is the subset of tags associated with the
// method (security.Call.MethodTags) that have the same type as tagType.
// Currently, tagType.Kind must be reflect.String, i.e., only tags that are
// named string types are supported.
//
// If multiple tags of tagType are associated with the method, then access is
// granted if the peer presents blessings that match the ACLs of each one of
// those tags. If no tags of tagType are associated with the method, then
// access is denied.
//
// If the TaggedACLMap provided is nil, then a nil authorizer is returned.
//
// Sample usage:
//
// (1) Attach tags to methods in the VDL (eg. myservice.vdl)
//   package myservice
//
//   type MyTag string
//   const (
//     ReadAccess  = MyTag("R")
//     WriteAccess = MyTag("W")
//   )
//
//   type MyService interface {
//     Get() ([]string, error)       {ReadAccess}
//     GetIndex(int) (string, error) {ReadAccess}
//
//     Set([]string) error           {WriteAccess}
//     SetIndex(int, string) error   {WriteAccess}
//
//     GetAndSet([]string) ([]string, error) {ReadAccess, WriteAccess}
//   }
//
// (2) Setup the ipc.Dispatcher to use the TaggedACLAuthorizer
//   import (
//     "reflect"
//     "v.io/x/ref/security/acl"
//
//     "v.io/v23/ipc"
//     "v.io/v23/security"
//   )
//
//   type dispatcher struct{}
//   func (d dispatcher) Lookup(suffix, method) (ipc.Invoker, security.Authorizer, error) {
//      acl := acl.TaggedACLMap{
//        "R": acl.ACL{In: []security.BlessingPattern{"alice/friends", "alice/family"} },
//        "W": acl.ACL{In: []security.BlessingPattern{"alice/family", "alice/colleagues" } },
//      }
//      typ := reflect.TypeOf(ReadAccess)  // equivalently, reflect.TypeOf(WriteAccess)
//      return newInvoker(), acl.TaggedACLAuthorizer(acl, typ), nil
//   }
//
// With the above dispatcher, the server will grant access to a peer with the blessing
// "alice/friend/bob" access only to the "Get" and "GetIndex" methods. A peer presenting
// the blessing "alice/colleague/carol" will get access only to the "Set" and "SetIndex"
// methods. A peer presenting "alice/family/mom" will get access to all methods, even
// GetAndSet - which requires that the blessing appear in the ACLs for both the
// ReadAccess and WriteAccess tags.
func TaggedACLAuthorizer(acls TaggedACLMap, tagType *vdl.Type) (security.Authorizer, error) {
	if tagType.Kind() != vdl.String {
		return nil, errTagType(tagType)
	}
	return &authorizer{acls, tagType}, nil
}

// TaggedACLAuthorizerFromFile applies the same authorization policy as
// TaggedACLAuthorizer, with the TaggedACLMap to be used sourced from a file named
// filename.
//
// Changes to the file are monitored and affect subsequent calls to Authorize.
// Currently, this is achieved by re-reading the file on every call to
// Authorize.
// TODO(ashankar,ataly): Use inotify or a similar mechanism to watch for
// changes.
func TaggedACLAuthorizerFromFile(filename string, tagType *vdl.Type) (security.Authorizer, error) {
	if tagType.Kind() != vdl.String {
		return nil, errTagType(tagType)
	}
	return &fileAuthorizer{filename, tagType}, nil
}

func errTagType(tt *vdl.Type) error {
	return verror.New(errTagNeedsString, nil, verror.New(errTagNeedsString, nil, tt, tt.Kind()))
}

type authorizer struct {
	acls    TaggedACLMap
	tagType *vdl.Type
}

func (a *authorizer) Authorize(call security.Call) error {
	// "Self-RPCs" are always authorized.
	if l, r := call.LocalBlessings().PublicKey(), call.RemoteBlessings().PublicKey(); l != nil && r != nil && reflect.DeepEqual(l, r) {
		return nil
	}

	blessings, invalid := call.RemoteBlessings().ForCall(call)
	grant := false
	if len(call.MethodTags()) == 0 {
		// The following error message leaks the fact that the server is likely
		// misconfigured, but that doesn't seem like a big deal.
		return verror.New(errNoMethodTags, call.Context(), call.Suffix(), call.Method())
	}
	for _, tag := range call.MethodTags() {
		if tag.Type() == a.tagType {
			if acl, exists := a.acls[tag.RawString()]; !exists || !acl.Includes(blessings...) {
				return NewErrACLMatch(nil, blessings, invalid)
			}
			grant = true
		}
	}
	if grant {
		return nil
	}
	return NewErrACLMatch(nil, blessings, invalid)
}

type fileAuthorizer struct {
	filename string
	tagType  *vdl.Type
}

func (a *fileAuthorizer) Authorize(call security.Call) error {
	acl, err := loadTaggedACLMapFromFile(a.filename)
	if err != nil {
		// TODO(ashankar): Information leak?
		return verror.New(errCantReadACLFromFile, call.Context(), err)
	}
	return (&authorizer{acl, a.tagType}).Authorize(call)
}

func loadTaggedACLMapFromFile(filename string) (TaggedACLMap, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadTaggedACLMap(file)
}
