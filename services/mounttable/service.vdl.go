// This file was auto-generated by the veyron vdl tool.
// Source: service.vdl

// Package mounttable defines a set of mount points and how to traverse them.
package mounttable

import (
	"veyron2/security"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_io "io"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

// MountedServer represents a server mounted on a specific name.
type MountedServer struct {
	// Server is the OA that's mounted.
	Server string
	// TTL is the remaining time (in seconds) before the mount entry expires.
	TTL uint32
}

// MountEntry represents a given name mounted in the mounttable.
type MountEntry struct {
	// Name is the mounted name.
	Name string
	// Servers (if present) specifies the mounted names.
	Servers []MountedServer
}

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// Globbable is the interface the client binds and uses.
// Globbable_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Globbable_ExcludingUniversal interface {
	// Glob returns all matching entries at the given server.
	Glob(ctx _gen_context.T, pattern string, opts ..._gen_ipc.CallOpt) (reply GlobbableGlobStream, err error)
}
type Globbable interface {
	_gen_ipc.UniversalServiceMethods
	Globbable_ExcludingUniversal
}

// GlobbableService is the interface the server implements.
type GlobbableService interface {

	// Glob returns all matching entries at the given server.
	Glob(context _gen_ipc.ServerContext, pattern string, stream GlobbableServiceGlobStream) (err error)
}

// GlobbableGlobStream is the interface for streaming responses of the method
// Glob in the service interface Globbable.
type GlobbableGlobStream interface {

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
	Value() MountEntry

	// Err returns a non-nil error iff the stream encountered
	// any errors.  Err does not block.
	Err() error

	// Finish blocks until the server is done and returns the positional
	// return values for call.
	//
	// If Cancel has been called, Finish will return immediately; the output of
	// Finish could either be an error signalling cancelation, or the correct
	// positional return values from the server depending on the timing of the
	// call.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return a non-EOF error.
	// Finish should be called at most once.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.  It
	// is safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

// Implementation of the GlobbableGlobStream interface that is not exported.
type implGlobbableGlobStream struct {
	clientCall _gen_ipc.Call
	val        MountEntry
	err        error
}

func (c *implGlobbableGlobStream) Advance() bool {
	c.val = MountEntry{}
	c.err = c.clientCall.Recv(&c.val)
	return c.err == nil
}

func (c *implGlobbableGlobStream) Value() MountEntry {
	return c.val
}

func (c *implGlobbableGlobStream) Err() error {
	if c.err == _gen_io.EOF {
		return nil
	}
	return c.err
}

func (c *implGlobbableGlobStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implGlobbableGlobStream) Cancel() {
	c.clientCall.Cancel()
}

// GlobbableServiceGlobStream is the interface for streaming responses of the method
// Glob in the service interface Globbable.
type GlobbableServiceGlobStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.  If the client has canceled, an error is returned.
	Send(item MountEntry) error
}

// Implementation of the GlobbableServiceGlobStream interface that is not exported.
type implGlobbableServiceGlobStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implGlobbableServiceGlobStream) Send(item MountEntry) error {
	return s.serverCall.Send(item)
}

// BindGlobbable returns the client stub implementing the Globbable
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindGlobbable(name string, opts ..._gen_ipc.BindOpt) (Globbable, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubGlobbable{client: client, name: name}

	return stub, nil
}

// NewServerGlobbable creates a new server stub.
//
// It takes a regular server implementing the GlobbableService
// interface, and returns a new server stub.
func NewServerGlobbable(server GlobbableService) interface{} {
	return &ServerStubGlobbable{
		service: server,
	}
}

// clientStubGlobbable implements Globbable.
type clientStubGlobbable struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubGlobbable) Glob(ctx _gen_context.T, pattern string, opts ..._gen_ipc.CallOpt) (reply GlobbableGlobStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Glob", []interface{}{pattern}, opts...); err != nil {
		return
	}
	reply = &implGlobbableGlobStream{clientCall: call}
	return
}

func (__gen_c *clientStubGlobbable) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubGlobbable) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubGlobbable) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubGlobbable wraps a server that implements
// GlobbableService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubGlobbable struct {
	service GlobbableService
}

func (__gen_s *ServerStubGlobbable) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Glob":
		return []interface{}{security.Label(1)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubGlobbable) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Glob"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "pattern", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},

		OutStream: 68,
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Server"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "TTL"},
			},
			"veyron2/services/mounttable.MountedServer", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x43, Name: "Servers"},
			},
			"veyron2/services/mounttable.MountEntry", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubGlobbable) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubGlobbable) Glob(call _gen_ipc.ServerCall, pattern string) (err error) {
	stream := &implGlobbableServiceGlobStream{serverCall: call}
	err = __gen_s.service.Glob(call, pattern, stream)
	return
}

// MountTable defines the interface to talk to a mounttable.
// MountTable is the interface the client binds and uses.
// MountTable_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type MountTable_ExcludingUniversal interface {
	Globbable_ExcludingUniversal
	// Mount Server (a global name) onto the receiver.
	// Subsequent mounts add to the servers mounted there.  The multiple
	// servers are considered equivalent and are meant solely for
	// availability, i.e., no load balancing is guaranteed.
	//
	// TTL is the number of seconds the mount is to last unless refreshed by
	// another mount of the same server.  A TTL of 0 represents an infinite
	// duration.  A server with an expired TTL should never appear in the
	// results nor affect the operation of any MountTable method, and should
	// act as if it was never present as far as the interface is concerned.
	Mount(ctx _gen_context.T, Server string, TTL uint32, opts ..._gen_ipc.CallOpt) (err error)
	// Unmount removes Server from the mount point.  If Server is empty, remove
	// all servers mounted there.
	// Returns a non-nil error iff Server remains mounted at the mount point.
	Unmount(ctx _gen_context.T, Server string, opts ..._gen_ipc.CallOpt) (err error)
	// ResolveStep takes the next step in resolving a name.  Returns the next
	// servers to query and the suffix at those servers.
	ResolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (Servers []MountedServer, Suffix string, err error)
}
type MountTable interface {
	_gen_ipc.UniversalServiceMethods
	MountTable_ExcludingUniversal
}

// MountTableService is the interface the server implements.
type MountTableService interface {
	GlobbableService
	// Mount Server (a global name) onto the receiver.
	// Subsequent mounts add to the servers mounted there.  The multiple
	// servers are considered equivalent and are meant solely for
	// availability, i.e., no load balancing is guaranteed.
	//
	// TTL is the number of seconds the mount is to last unless refreshed by
	// another mount of the same server.  A TTL of 0 represents an infinite
	// duration.  A server with an expired TTL should never appear in the
	// results nor affect the operation of any MountTable method, and should
	// act as if it was never present as far as the interface is concerned.
	Mount(context _gen_ipc.ServerContext, Server string, TTL uint32) (err error)
	// Unmount removes Server from the mount point.  If Server is empty, remove
	// all servers mounted there.
	// Returns a non-nil error iff Server remains mounted at the mount point.
	Unmount(context _gen_ipc.ServerContext, Server string) (err error)
	// ResolveStep takes the next step in resolving a name.  Returns the next
	// servers to query and the suffix at those servers.
	ResolveStep(context _gen_ipc.ServerContext) (Servers []MountedServer, Suffix string, err error)
}

// BindMountTable returns the client stub implementing the MountTable
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindMountTable(name string, opts ..._gen_ipc.BindOpt) (MountTable, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubMountTable{client: client, name: name}
	stub.Globbable_ExcludingUniversal, _ = BindGlobbable(name, client)

	return stub, nil
}

// NewServerMountTable creates a new server stub.
//
// It takes a regular server implementing the MountTableService
// interface, and returns a new server stub.
func NewServerMountTable(server MountTableService) interface{} {
	return &ServerStubMountTable{
		ServerStubGlobbable: *NewServerGlobbable(server).(*ServerStubGlobbable),
		service:             server,
	}
}

// clientStubMountTable implements MountTable.
type clientStubMountTable struct {
	Globbable_ExcludingUniversal

	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubMountTable) Mount(ctx _gen_context.T, Server string, TTL uint32, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Mount", []interface{}{Server, TTL}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMountTable) Unmount(ctx _gen_context.T, Server string, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Unmount", []interface{}{Server}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMountTable) ResolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (Servers []MountedServer, Suffix string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "ResolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&Servers, &Suffix, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMountTable) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMountTable) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMountTable) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubMountTable wraps a server that implements
// MountTableService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubMountTable struct {
	ServerStubGlobbable

	service MountTableService
}

func (__gen_s *ServerStubMountTable) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	if resp, err := __gen_s.ServerStubGlobbable.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	switch method {
	case "Mount":
		return []interface{}{security.Label(2)}, nil
	case "Unmount":
		return []interface{}{security.Label(2)}, nil
	case "ResolveStep":
		return []interface{}{security.Label(1)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubMountTable) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Mount"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Server", Type: 3},
			{Name: "TTL", Type: 52},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["ResolveStep"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "Servers", Type: 67},
			{Name: "Suffix", Type: 3},
			{Name: "Error", Type: 65},
		},
	}
	result.Methods["Unmount"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Server", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Server"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "TTL"},
			},
			"veyron2/services/mounttable.MountedServer", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}}
	var ss _gen_ipc.ServiceSignature
	var firstAdded int
	ss, _ = __gen_s.ServerStubGlobbable.Signature(call)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.InArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.OutArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= _gen_wiretype.TypeIDFirst {
			v.InStream += _gen_wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= _gen_wiretype.TypeIDFirst {
			v.OutStream += _gen_wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case _gen_wiretype.SliceType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.ArrayType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.MapType:
			if wt.Key >= _gen_wiretype.TypeIDFirst {
				wt.Key += _gen_wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.StructType:
			for i, fld := range wt.Fields {
				if fld.Type >= _gen_wiretype.TypeIDFirst {
					wt.Fields[i].Type += _gen_wiretype.TypeID(firstAdded)
				}
			}
			d = wt
			// NOTE: other types are missing, but we are upgrading anyways.
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}

	return result, nil
}

func (__gen_s *ServerStubMountTable) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubMountTable) Mount(call _gen_ipc.ServerCall, Server string, TTL uint32) (err error) {
	err = __gen_s.service.Mount(call, Server, TTL)
	return
}

func (__gen_s *ServerStubMountTable) Unmount(call _gen_ipc.ServerCall, Server string) (err error) {
	err = __gen_s.service.Unmount(call, Server)
	return
}

func (__gen_s *ServerStubMountTable) ResolveStep(call _gen_ipc.ServerCall) (Servers []MountedServer, Suffix string, err error) {
	Servers, Suffix, err = __gen_s.service.ResolveStep(call)
	return
}
