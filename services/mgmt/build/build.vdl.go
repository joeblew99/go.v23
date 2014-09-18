// This file was auto-generated by the veyron vdl tool.
// Source: build.vdl

// Package build supports building and describing Veyron binaries.
//
// TODO(jsimsa): Switch Architecture, Format, and OperatingSystem type
// to enum when supported.
package build

import (
	"veyron.io/veyron/veyron2/services/mgmt/binary"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_io "io"
	_gen_veyron2 "veyron.io/veyron/veyron2"
	_gen_context "veyron.io/veyron/veyron2/context"
	_gen_ipc "veyron.io/veyron/veyron2/ipc"
	_gen_naming "veyron.io/veyron/veyron2/naming"
	_gen_vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	_gen_wiretype "veyron.io/veyron/veyron2/wiretype"
)

// Architecture specifies the hardware architecture of a host.
type Architecture string

// Format specifies the file format of a host.
type Format string

// OperatingSystem specifies the operating system of a host.
type OperatingSystem string

// File records the name and contents of a file.
type File struct {
	Name     string
	Contents []byte
}

const X86 = Architecture("386")

const AMD64 = Architecture("amd64")

const ARM = Architecture("arm")

const UnsupportedArchitecture = Architecture("unsupported")

const ELF = Format("ELF")

const MACH = Format("MACH")

const PE = Format("PE")

const UnsupportedFormat = Format("unsupported")

const Darwin = OperatingSystem("darwin")

const Linux = OperatingSystem("linux")

const Windows = OperatingSystem("windows")

const UnsupportedOS = OperatingSystem("unsupported")

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// Builder describes an interface for building binaries from source.
// Builder is the interface the client binds and uses.
// Builder_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Builder_ExcludingUniversal interface {
	// Build streams sources to the build server, which then attempts to
	// build the sources and streams back the compiled binaries.
	Build(ctx _gen_context.T, Arch Architecture, OS OperatingSystem, opts ..._gen_ipc.CallOpt) (reply BuilderBuildCall, err error)
	// Describe generates a description for a binary identified by
	// the given Object name.
	Describe(ctx _gen_context.T, Name string, opts ..._gen_ipc.CallOpt) (reply binary.Description, err error)
}
type Builder interface {
	_gen_ipc.UniversalServiceMethods
	Builder_ExcludingUniversal
}

// BuilderService is the interface the server implements.
type BuilderService interface {

	// Build streams sources to the build server, which then attempts to
	// build the sources and streams back the compiled binaries.
	Build(context _gen_ipc.ServerContext, Arch Architecture, OS OperatingSystem, stream BuilderServiceBuildStream) (reply []byte, err error)
	// Describe generates a description for a binary identified by
	// the given Object name.
	Describe(context _gen_ipc.ServerContext, Name string) (reply binary.Description, err error)
}

// BuilderBuildCall is the interface for call object of the method
// Build in the service interface Builder.
type BuilderBuildCall interface {
	// RecvStream returns the recv portion of the stream
	RecvStream() interface {
		// Advance stages an element so the client can retrieve it
		// with Value.  Advance returns true iff there is an
		// element to retrieve.  The client must call Advance before
		// calling Value. Advance may block if an element is not
		// immediately available.
		Advance() bool

		// Value returns the element that was staged by Advance.
		// Value may panic if Advance returned false or was not
		// called at all.  Value does not block.
		Value() File

		// Err returns a non-nil error iff the stream encountered
		// any errors.  Err does not block.
		Err() error
	}

	// SendStream returns the send portion of the stream
	SendStream() interface {
		// Send places the item onto the output stream, blocking if there is no
		// buffer space available.  Calls to Send after having called Close
		// or Cancel will fail.  Any blocked Send calls will be unblocked upon
		// calling Cancel.
		Send(item File) error

		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.  This is
		// an optional call - it's used by streaming clients that need the
		// server to receive the io.EOF terminator before the client calls
		// Finish (for example, if the client needs to continue receiving items
		// from the server after having finished sending).
		// Calls to Close after having called Cancel will fail.
		// Like Send, Close blocks when there's no buffer space available.
		Close() error
	}

	// Finish performs the equivalent of SendStream().Close, then blocks until the server
	// is done, and returns the positional return values for call.
	// If Cancel has been called, Finish will return immediately; the output of
	// Finish could either be an error signalling cancelation, or the correct
	// positional return values from the server depending on the timing of the
	// call.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.
	// Finish should be called at most once.
	Finish() (reply []byte, err error)

	// Cancel cancels the RPC, notifying the server to stop processing.  It
	// is safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implBuilderBuildStreamSender struct {
	clientCall _gen_ipc.Call
}

func (c *implBuilderBuildStreamSender) Send(item File) error {
	return c.clientCall.Send(item)
}

func (c *implBuilderBuildStreamSender) Close() error {
	return c.clientCall.CloseSend()
}

type implBuilderBuildStreamIterator struct {
	clientCall _gen_ipc.Call
	val        File
	err        error
}

func (c *implBuilderBuildStreamIterator) Advance() bool {
	c.val = File{}
	c.err = c.clientCall.Recv(&c.val)
	return c.err == nil
}

func (c *implBuilderBuildStreamIterator) Value() File {
	return c.val
}

func (c *implBuilderBuildStreamIterator) Err() error {
	if c.err == _gen_io.EOF {
		return nil
	}
	return c.err
}

// Implementation of the BuilderBuildCall interface that is not exported.
type implBuilderBuildCall struct {
	clientCall  _gen_ipc.Call
	writeStream implBuilderBuildStreamSender
	readStream  implBuilderBuildStreamIterator
}

func (c *implBuilderBuildCall) SendStream() interface {
	Send(item File) error
	Close() error
} {
	return &c.writeStream
}

func (c *implBuilderBuildCall) RecvStream() interface {
	Advance() bool
	Value() File
	Err() error
} {
	return &c.readStream
}

func (c *implBuilderBuildCall) Finish() (reply []byte, err error) {
	if ierr := c.clientCall.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implBuilderBuildCall) Cancel() {
	c.clientCall.Cancel()
}

type implBuilderServiceBuildStreamSender struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implBuilderServiceBuildStreamSender) Send(item File) error {
	return s.serverCall.Send(item)
}

type implBuilderServiceBuildStreamIterator struct {
	serverCall _gen_ipc.ServerCall
	val        File
	err        error
}

func (s *implBuilderServiceBuildStreamIterator) Advance() bool {
	s.val = File{}
	s.err = s.serverCall.Recv(&s.val)
	return s.err == nil
}

func (s *implBuilderServiceBuildStreamIterator) Value() File {
	return s.val
}

func (s *implBuilderServiceBuildStreamIterator) Err() error {
	if s.err == _gen_io.EOF {
		return nil
	}
	return s.err
}

// BuilderServiceBuildStream is the interface for streaming responses of the method
// Build in the service interface Builder.
type BuilderServiceBuildStream interface {
	// SendStream returns the send portion of the stream.
	SendStream() interface {
		// Send places the item onto the output stream, blocking if there is no buffer
		// space available.  If the client has canceled, an error is returned.
		Send(item File) error
	}
	// RecvStream returns the recv portion of the stream
	RecvStream() interface {
		// Advance stages an element so the client can retrieve it
		// with Value.  Advance returns true iff there is an
		// element to retrieve.  The client must call Advance before
		// calling Value.  Advance may block if an element is not
		// immediately available.
		Advance() bool

		// Value returns the element that was staged by Advance.
		// Value may panic if Advance returned false or was not
		// called at all.  Value does not block.
		Value() File

		// Err returns a non-nil error iff the stream encountered
		// any errors.  Err does not block.
		Err() error
	}
}

// Implementation of the BuilderServiceBuildStream interface that is not exported.
type implBuilderServiceBuildStream struct {
	writer implBuilderServiceBuildStreamSender
	reader implBuilderServiceBuildStreamIterator
}

func (s *implBuilderServiceBuildStream) SendStream() interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.  If the client has canceled, an error is returned.
	Send(item File) error
} {
	return &s.writer
}

func (s *implBuilderServiceBuildStream) RecvStream() interface {
	// Advance stages an element so the client can retrieve it
	// with Value.  Advance returns true iff there is an
	// element to retrieve.  The client must call Advance before
	// calling Value.  The client must call Cancel if it does
	// not iterate through all elements (i.e. until Advance
	// returns false).  Advance may block if an element is not
	// immediately available.
	Advance() bool

	// Value returns the element that was staged by Advance.
	// Value may panic if Advance returned false or was not
	// called at all.  Value does not block.
	Value() File

	// Err returns a non-nil error iff the stream encountered
	// any errors.  Err does not block.
	Err() error
} {
	return &s.reader
}

// BindBuilder returns the client stub implementing the Builder
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindBuilder(name string, opts ..._gen_ipc.BindOpt) (Builder, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		// Do nothing.
	case 1:
		if clientOpt, ok := opts[0].(_gen_ipc.Client); opts[0] == nil || ok {
			client = clientOpt
		} else {
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubBuilder{defaultClient: client, name: name}

	return stub, nil
}

// NewServerBuilder creates a new server stub.
//
// It takes a regular server implementing the BuilderService
// interface, and returns a new server stub.
func NewServerBuilder(server BuilderService) interface{} {
	return &ServerStubBuilder{
		service: server,
	}
}

// clientStubBuilder implements Builder.
type clientStubBuilder struct {
	defaultClient _gen_ipc.Client
	name          string
}

func (__gen_c *clientStubBuilder) client(ctx _gen_context.T) _gen_ipc.Client {
	if __gen_c.defaultClient != nil {
		return __gen_c.defaultClient
	}
	return _gen_veyron2.RuntimeFromContext(ctx).Client()
}

func (__gen_c *clientStubBuilder) Build(ctx _gen_context.T, Arch Architecture, OS OperatingSystem, opts ..._gen_ipc.CallOpt) (reply BuilderBuildCall, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Build", []interface{}{Arch, OS}, opts...); err != nil {
		return
	}
	reply = &implBuilderBuildCall{clientCall: call, writeStream: implBuilderBuildStreamSender{clientCall: call}, readStream: implBuilderBuildStreamIterator{clientCall: call}}
	return
}

func (__gen_c *clientStubBuilder) Describe(ctx _gen_context.T, Name string, opts ..._gen_ipc.CallOpt) (reply binary.Description, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Describe", []interface{}{Name}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBuilder) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBuilder) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubBuilder) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubBuilder wraps a server that implements
// BuilderService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubBuilder struct {
	service BuilderService
}

func (__gen_s *ServerStubBuilder) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Build":
		return []interface{}{}, nil
	case "Describe":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubBuilder) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Build"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Arch", Type: 65},
			{Name: "OS", Type: 66},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 68},
			{Name: "", Type: 69},
		},
		InStream:  70,
		OutStream: 70,
	}
	result.Methods["Describe"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Name", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 72},
			{Name: "", Type: 69},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x3, Name: "veyron.io/veyron/veyron2/services/mgmt/build.Architecture", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x3, Name: "veyron.io/veyron/veyron2/services/mgmt/build.OperatingSystem", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x43, Name: "", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x44, Name: "Contents"},
			},
			"veyron.io/veyron/veyron2/services/mgmt/build.File", []string(nil)},
		_gen_wiretype.MapType{Key: 0x3, Elem: 0x2, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x47, Name: "Profiles"},
			},
			"veyron.io/veyron/veyron2/services/mgmt/binary.Description", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubBuilder) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubBuilder) Build(call _gen_ipc.ServerCall, Arch Architecture, OS OperatingSystem) (reply []byte, err error) {
	stream := &implBuilderServiceBuildStream{reader: implBuilderServiceBuildStreamIterator{serverCall: call}, writer: implBuilderServiceBuildStreamSender{serverCall: call}}
	reply, err = __gen_s.service.Build(call, Arch, OS, stream)
	return
}

func (__gen_s *ServerStubBuilder) Describe(call _gen_ipc.ServerCall, Name string) (reply binary.Description, err error) {
	reply, err = __gen_s.service.Describe(call, Name)
	return
}
