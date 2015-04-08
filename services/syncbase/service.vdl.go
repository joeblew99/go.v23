// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: service.vdl

package syncbase

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/rpc"
	"v.io/v23/vdl"
	"v.io/v23/verror"

	// VDL user imports
	"v.io/v23/security/access"
	"v.io/v23/services/permissions"
)

// Schema is a Database schema. Currently it's just a set of Table names, and
// Item types are not enforced anywhere.
// TODO(sadovsky): Iterate on schema representation and enforcement.
type Schema map[string]struct{}

func (Schema) __VDLReflect(struct {
	Name string "v.io/syncbase/v23/services/syncbase.Schema"
}) {
}

func init() {
	vdl.Register((*Schema)(nil))
}

var (
	ErrInvalidName = verror.Register("v.io/syncbase/v23/services/syncbase.InvalidName", verror.NoRetry, "{1:}{2:} invalid name: {3}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidName.ID), "{1:}{2:} invalid name: {3}")
}

// NewErrInvalidName returns an error with the ErrInvalidName ID.
func NewErrInvalidName(ctx *context.T, name string) error {
	return verror.New(ErrInvalidName, ctx, name)
}

// ServiceClientMethods is the client interface
// containing Service methods.
//
// Service represents a Vanadium syncbase service.
// Service.Glob operates over Universe names.
type ServiceClientMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/services/permissions"
	//
	//   type MyObject interface {
	//     permissions.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(acl access.Permissions, version string) error         {Red}
	//    GetPermissions() (acl access.Permissions, version string, err error) {Blue}
	//  }
	permissions.ObjectClientMethods
}

// ServiceClientStub adds universal methods to ServiceClientMethods.
type ServiceClientStub interface {
	ServiceClientMethods
	rpc.UniversalServiceMethods
}

// ServiceClient returns a client stub for Service.
func ServiceClient(name string) ServiceClientStub {
	return implServiceClientStub{name, permissions.ObjectClient(name)}
}

type implServiceClientStub struct {
	name string

	permissions.ObjectClientStub
}

// ServiceServerMethods is the interface a server writer
// implements for Service.
//
// Service represents a Vanadium syncbase service.
// Service.Glob operates over Universe names.
type ServiceServerMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/services/permissions"
	//
	//   type MyObject interface {
	//     permissions.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(acl access.Permissions, version string) error         {Red}
	//    GetPermissions() (acl access.Permissions, version string, err error) {Blue}
	//  }
	permissions.ObjectServerMethods
}

// ServiceServerStubMethods is the server interface containing
// Service methods, as expected by rpc.Server.
// There is no difference between this interface and ServiceServerMethods
// since there are no streaming methods.
type ServiceServerStubMethods ServiceServerMethods

// ServiceServerStub adds universal methods to ServiceServerStubMethods.
type ServiceServerStub interface {
	ServiceServerStubMethods
	// Describe the Service interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ServiceServer returns a server stub for Service.
// It converts an implementation of ServiceServerMethods into
// an object that may be used by rpc.Server.
func ServiceServer(impl ServiceServerMethods) ServiceServerStub {
	stub := implServiceServerStub{
		impl:             impl,
		ObjectServerStub: permissions.ObjectServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implServiceServerStub struct {
	impl ServiceServerMethods
	permissions.ObjectServerStub
	gs *rpc.GlobState
}

func (s implServiceServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implServiceServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ServiceDesc, permissions.ObjectDesc}
}

// ServiceDesc describes the Service interface.
var ServiceDesc rpc.InterfaceDesc = descService

// descService hides the desc to keep godoc clean.
var descService = rpc.InterfaceDesc{
	Name:    "Service",
	PkgPath: "v.io/syncbase/v23/services/syncbase",
	Doc:     "// Service represents a Vanadium syncbase service.\n// Service.Glob operates over Universe names.",
	Embeds: []rpc.EmbedDesc{
		{"Object", "v.io/v23/services/permissions", "// Object provides access control for Vanadium objects.\n//\n// Vanadium services implementing dynamic access control would typically embed\n// this interface and tag additional methods defined by the service with one of\n// Admin, Read, Write, Resolve etc. For example, the VDL definition of the\n// object would be:\n//\n//   package mypackage\n//\n//   import \"v.io/v23/security/access\"\n//   import \"v.io/v23/services/permissions\"\n//\n//   type MyObject interface {\n//     permissions.Object\n//     MyRead() (string, error) {access.Read}\n//     MyWrite(string) error    {access.Write}\n//   }\n//\n// If the set of pre-defined tags is insufficient, services may define their\n// own tag type and annotate all methods with this new type.\n//\n// Instead of embedding this Object interface, define SetPermissions and\n// GetPermissions in their own interface. Authorization policies will typically\n// respect annotations of a single type. For example, the VDL definition of an\n// object would be:\n//\n//  package mypackage\n//\n//  import \"v.io/v23/security/access\"\n//\n//  type MyTag string\n//\n//  const (\n//    Blue = MyTag(\"Blue\")\n//    Red  = MyTag(\"Red\")\n//  )\n//\n//  type MyObject interface {\n//    MyMethod() (string, error) {Blue}\n//\n//    // Allow clients to change access via the access.Object interface:\n//    SetPermissions(acl access.Permissions, version string) error         {Red}\n//    GetPermissions() (acl access.Permissions, version string, err error) {Blue}\n//  }"},
	},
}

// UniverseClientMethods is the client interface
// containing Universe methods.
//
// Universe represents a collection of Databases.
// Universe.Glob operates over Database names.
// We expect there to be one Universe per app, likely created by the node
// manager as part of the app installation procedure.
type UniverseClientMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/services/permissions"
	//
	//   type MyObject interface {
	//     permissions.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(acl access.Permissions, version string) error         {Red}
	//    GetPermissions() (acl access.Permissions, version string, err error) {Blue}
	//  }
	permissions.ObjectClientMethods
	// Create creates this Universe.
	// If acl is nil, Permissions is inherited (copied) from the Service.
	// Create requires the caller to have Write permission at the Service.
	Create(ctx *context.T, acl access.Permissions, opts ...rpc.CallOpt) error
	// Delete deletes this Universe.
	Delete(*context.T, ...rpc.CallOpt) error
}

// UniverseClientStub adds universal methods to UniverseClientMethods.
type UniverseClientStub interface {
	UniverseClientMethods
	rpc.UniversalServiceMethods
}

// UniverseClient returns a client stub for Universe.
func UniverseClient(name string) UniverseClientStub {
	return implUniverseClientStub{name, permissions.ObjectClient(name)}
}

type implUniverseClientStub struct {
	name string

	permissions.ObjectClientStub
}

func (c implUniverseClientStub) Create(ctx *context.T, i0 access.Permissions, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Create", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implUniverseClientStub) Delete(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Delete", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

// UniverseServerMethods is the interface a server writer
// implements for Universe.
//
// Universe represents a collection of Databases.
// Universe.Glob operates over Database names.
// We expect there to be one Universe per app, likely created by the node
// manager as part of the app installation procedure.
type UniverseServerMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/services/permissions"
	//
	//   type MyObject interface {
	//     permissions.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(acl access.Permissions, version string) error         {Red}
	//    GetPermissions() (acl access.Permissions, version string, err error) {Blue}
	//  }
	permissions.ObjectServerMethods
	// Create creates this Universe.
	// If acl is nil, Permissions is inherited (copied) from the Service.
	// Create requires the caller to have Write permission at the Service.
	Create(call rpc.ServerCall, acl access.Permissions) error
	// Delete deletes this Universe.
	Delete(rpc.ServerCall) error
}

// UniverseServerStubMethods is the server interface containing
// Universe methods, as expected by rpc.Server.
// There is no difference between this interface and UniverseServerMethods
// since there are no streaming methods.
type UniverseServerStubMethods UniverseServerMethods

// UniverseServerStub adds universal methods to UniverseServerStubMethods.
type UniverseServerStub interface {
	UniverseServerStubMethods
	// Describe the Universe interfaces.
	Describe__() []rpc.InterfaceDesc
}

// UniverseServer returns a server stub for Universe.
// It converts an implementation of UniverseServerMethods into
// an object that may be used by rpc.Server.
func UniverseServer(impl UniverseServerMethods) UniverseServerStub {
	stub := implUniverseServerStub{
		impl:             impl,
		ObjectServerStub: permissions.ObjectServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implUniverseServerStub struct {
	impl UniverseServerMethods
	permissions.ObjectServerStub
	gs *rpc.GlobState
}

func (s implUniverseServerStub) Create(call rpc.ServerCall, i0 access.Permissions) error {
	return s.impl.Create(call, i0)
}

func (s implUniverseServerStub) Delete(call rpc.ServerCall) error {
	return s.impl.Delete(call)
}

func (s implUniverseServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implUniverseServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{UniverseDesc, permissions.ObjectDesc}
}

// UniverseDesc describes the Universe interface.
var UniverseDesc rpc.InterfaceDesc = descUniverse

// descUniverse hides the desc to keep godoc clean.
var descUniverse = rpc.InterfaceDesc{
	Name:    "Universe",
	PkgPath: "v.io/syncbase/v23/services/syncbase",
	Doc:     "// Universe represents a collection of Databases.\n// Universe.Glob operates over Database names.\n// We expect there to be one Universe per app, likely created by the node\n// manager as part of the app installation procedure.",
	Embeds: []rpc.EmbedDesc{
		{"Object", "v.io/v23/services/permissions", "// Object provides access control for Vanadium objects.\n//\n// Vanadium services implementing dynamic access control would typically embed\n// this interface and tag additional methods defined by the service with one of\n// Admin, Read, Write, Resolve etc. For example, the VDL definition of the\n// object would be:\n//\n//   package mypackage\n//\n//   import \"v.io/v23/security/access\"\n//   import \"v.io/v23/services/permissions\"\n//\n//   type MyObject interface {\n//     permissions.Object\n//     MyRead() (string, error) {access.Read}\n//     MyWrite(string) error    {access.Write}\n//   }\n//\n// If the set of pre-defined tags is insufficient, services may define their\n// own tag type and annotate all methods with this new type.\n//\n// Instead of embedding this Object interface, define SetPermissions and\n// GetPermissions in their own interface. Authorization policies will typically\n// respect annotations of a single type. For example, the VDL definition of an\n// object would be:\n//\n//  package mypackage\n//\n//  import \"v.io/v23/security/access\"\n//\n//  type MyTag string\n//\n//  const (\n//    Blue = MyTag(\"Blue\")\n//    Red  = MyTag(\"Red\")\n//  )\n//\n//  type MyObject interface {\n//    MyMethod() (string, error) {Blue}\n//\n//    // Allow clients to change access via the access.Object interface:\n//    SetPermissions(acl access.Permissions, version string) error         {Red}\n//    GetPermissions() (acl access.Permissions, version string, err error) {Blue}\n//  }"},
	},
	Methods: []rpc.MethodDesc{
		{
			Name: "Create",
			Doc:  "// Create creates this Universe.\n// If acl is nil, Permissions is inherited (copied) from the Service.\n// Create requires the caller to have Write permission at the Service.",
			InArgs: []rpc.ArgDesc{
				{"acl", ``}, // access.Permissions
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Delete",
			Doc:  "// Delete deletes this Universe.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
	},
}

// DatabaseClientMethods is the client interface
// containing Database methods.
//
// Database represents a collection of Tables. Batch operations, queries, sync,
// watch, etc. all currently operate at the Database level. A Database's etag
// covers both its Permissions and its schema.
// Database.Glob operates over Table names.
//
// TODO(sadovsky): Add Watch method.
// TODO(sadovsky): Support batch operations.
// TODO(sadovsky): Iterate on the schema management API as we figure out how to
// deal with schema versioning and sync.
type DatabaseClientMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/services/permissions"
	//
	//   type MyObject interface {
	//     permissions.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(acl access.Permissions, version string) error         {Red}
	//    GetPermissions() (acl access.Permissions, version string, err error) {Blue}
	//  }
	permissions.ObjectClientMethods
	// Create creates this Database.
	// If acl is nil, Permissions is inherited (copied) from the Universe.
	// Create requires the caller to have Write permission at the Universe.
	Create(ctx *context.T, acl access.Permissions, opts ...rpc.CallOpt) error
	// Delete deletes this Database.
	Delete(*context.T, ...rpc.CallOpt) error
	// UpdateSchema updates the schema for this Database, creating and deleting
	// Tables under the hood as needed.
	UpdateSchema(ctx *context.T, schema Schema, etag string, opts ...rpc.CallOpt) error
	// GetSchema returns the schema for this Database.
	GetSchema(*context.T, ...rpc.CallOpt) (schema Schema, etag string, err error)
}

// DatabaseClientStub adds universal methods to DatabaseClientMethods.
type DatabaseClientStub interface {
	DatabaseClientMethods
	rpc.UniversalServiceMethods
}

// DatabaseClient returns a client stub for Database.
func DatabaseClient(name string) DatabaseClientStub {
	return implDatabaseClientStub{name, permissions.ObjectClient(name)}
}

type implDatabaseClientStub struct {
	name string

	permissions.ObjectClientStub
}

func (c implDatabaseClientStub) Create(ctx *context.T, i0 access.Permissions, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Create", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implDatabaseClientStub) Delete(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Delete", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implDatabaseClientStub) UpdateSchema(ctx *context.T, i0 Schema, i1 string, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "UpdateSchema", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implDatabaseClientStub) GetSchema(ctx *context.T, opts ...rpc.CallOpt) (o0 Schema, o1 string, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "GetSchema", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0, &o1)
	return
}

// DatabaseServerMethods is the interface a server writer
// implements for Database.
//
// Database represents a collection of Tables. Batch operations, queries, sync,
// watch, etc. all currently operate at the Database level. A Database's etag
// covers both its Permissions and its schema.
// Database.Glob operates over Table names.
//
// TODO(sadovsky): Add Watch method.
// TODO(sadovsky): Support batch operations.
// TODO(sadovsky): Iterate on the schema management API as we figure out how to
// deal with schema versioning and sync.
type DatabaseServerMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/services/permissions"
	//
	//   type MyObject interface {
	//     permissions.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(acl access.Permissions, version string) error         {Red}
	//    GetPermissions() (acl access.Permissions, version string, err error) {Blue}
	//  }
	permissions.ObjectServerMethods
	// Create creates this Database.
	// If acl is nil, Permissions is inherited (copied) from the Universe.
	// Create requires the caller to have Write permission at the Universe.
	Create(call rpc.ServerCall, acl access.Permissions) error
	// Delete deletes this Database.
	Delete(rpc.ServerCall) error
	// UpdateSchema updates the schema for this Database, creating and deleting
	// Tables under the hood as needed.
	UpdateSchema(call rpc.ServerCall, schema Schema, etag string) error
	// GetSchema returns the schema for this Database.
	GetSchema(rpc.ServerCall) (schema Schema, etag string, err error)
}

// DatabaseServerStubMethods is the server interface containing
// Database methods, as expected by rpc.Server.
// There is no difference between this interface and DatabaseServerMethods
// since there are no streaming methods.
type DatabaseServerStubMethods DatabaseServerMethods

// DatabaseServerStub adds universal methods to DatabaseServerStubMethods.
type DatabaseServerStub interface {
	DatabaseServerStubMethods
	// Describe the Database interfaces.
	Describe__() []rpc.InterfaceDesc
}

// DatabaseServer returns a server stub for Database.
// It converts an implementation of DatabaseServerMethods into
// an object that may be used by rpc.Server.
func DatabaseServer(impl DatabaseServerMethods) DatabaseServerStub {
	stub := implDatabaseServerStub{
		impl:             impl,
		ObjectServerStub: permissions.ObjectServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implDatabaseServerStub struct {
	impl DatabaseServerMethods
	permissions.ObjectServerStub
	gs *rpc.GlobState
}

func (s implDatabaseServerStub) Create(call rpc.ServerCall, i0 access.Permissions) error {
	return s.impl.Create(call, i0)
}

func (s implDatabaseServerStub) Delete(call rpc.ServerCall) error {
	return s.impl.Delete(call)
}

func (s implDatabaseServerStub) UpdateSchema(call rpc.ServerCall, i0 Schema, i1 string) error {
	return s.impl.UpdateSchema(call, i0, i1)
}

func (s implDatabaseServerStub) GetSchema(call rpc.ServerCall) (Schema, string, error) {
	return s.impl.GetSchema(call)
}

func (s implDatabaseServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implDatabaseServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{DatabaseDesc, permissions.ObjectDesc}
}

// DatabaseDesc describes the Database interface.
var DatabaseDesc rpc.InterfaceDesc = descDatabase

// descDatabase hides the desc to keep godoc clean.
var descDatabase = rpc.InterfaceDesc{
	Name:    "Database",
	PkgPath: "v.io/syncbase/v23/services/syncbase",
	Doc:     "// Database represents a collection of Tables. Batch operations, queries, sync,\n// watch, etc. all currently operate at the Database level. A Database's etag\n// covers both its Permissions and its schema.\n// Database.Glob operates over Table names.\n//\n// TODO(sadovsky): Add Watch method.\n// TODO(sadovsky): Support batch operations.\n// TODO(sadovsky): Iterate on the schema management API as we figure out how to\n// deal with schema versioning and sync.",
	Embeds: []rpc.EmbedDesc{
		{"Object", "v.io/v23/services/permissions", "// Object provides access control for Vanadium objects.\n//\n// Vanadium services implementing dynamic access control would typically embed\n// this interface and tag additional methods defined by the service with one of\n// Admin, Read, Write, Resolve etc. For example, the VDL definition of the\n// object would be:\n//\n//   package mypackage\n//\n//   import \"v.io/v23/security/access\"\n//   import \"v.io/v23/services/permissions\"\n//\n//   type MyObject interface {\n//     permissions.Object\n//     MyRead() (string, error) {access.Read}\n//     MyWrite(string) error    {access.Write}\n//   }\n//\n// If the set of pre-defined tags is insufficient, services may define their\n// own tag type and annotate all methods with this new type.\n//\n// Instead of embedding this Object interface, define SetPermissions and\n// GetPermissions in their own interface. Authorization policies will typically\n// respect annotations of a single type. For example, the VDL definition of an\n// object would be:\n//\n//  package mypackage\n//\n//  import \"v.io/v23/security/access\"\n//\n//  type MyTag string\n//\n//  const (\n//    Blue = MyTag(\"Blue\")\n//    Red  = MyTag(\"Red\")\n//  )\n//\n//  type MyObject interface {\n//    MyMethod() (string, error) {Blue}\n//\n//    // Allow clients to change access via the access.Object interface:\n//    SetPermissions(acl access.Permissions, version string) error         {Red}\n//    GetPermissions() (acl access.Permissions, version string, err error) {Blue}\n//  }"},
	},
	Methods: []rpc.MethodDesc{
		{
			Name: "Create",
			Doc:  "// Create creates this Database.\n// If acl is nil, Permissions is inherited (copied) from the Universe.\n// Create requires the caller to have Write permission at the Universe.",
			InArgs: []rpc.ArgDesc{
				{"acl", ``}, // access.Permissions
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Delete",
			Doc:  "// Delete deletes this Database.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "UpdateSchema",
			Doc:  "// UpdateSchema updates the schema for this Database, creating and deleting\n// Tables under the hood as needed.",
			InArgs: []rpc.ArgDesc{
				{"schema", ``}, // Schema
				{"etag", ``},   // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "GetSchema",
			Doc:  "// GetSchema returns the schema for this Database.",
			OutArgs: []rpc.ArgDesc{
				{"schema", ``}, // Schema
				{"etag", ``},   // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
	},
}

// TableClientMethods is the client interface
// containing Table methods.
//
// Table represents a collection of Items (rows). Table implements built-in
// methods such as Glob but currently does not provide any custom methods.
// Table.Glob operates over the primary keys of Items in the Table.
// All Permissions checks are performed against the Database Permissions.
type TableClientMethods interface {
}

// TableClientStub adds universal methods to TableClientMethods.
type TableClientStub interface {
	TableClientMethods
	rpc.UniversalServiceMethods
}

// TableClient returns a client stub for Table.
func TableClient(name string) TableClientStub {
	return implTableClientStub{name}
}

type implTableClientStub struct {
	name string
}

// TableServerMethods is the interface a server writer
// implements for Table.
//
// Table represents a collection of Items (rows). Table implements built-in
// methods such as Glob but currently does not provide any custom methods.
// Table.Glob operates over the primary keys of Items in the Table.
// All Permissions checks are performed against the Database Permissions.
type TableServerMethods interface {
}

// TableServerStubMethods is the server interface containing
// Table methods, as expected by rpc.Server.
// There is no difference between this interface and TableServerMethods
// since there are no streaming methods.
type TableServerStubMethods TableServerMethods

// TableServerStub adds universal methods to TableServerStubMethods.
type TableServerStub interface {
	TableServerStubMethods
	// Describe the Table interfaces.
	Describe__() []rpc.InterfaceDesc
}

// TableServer returns a server stub for Table.
// It converts an implementation of TableServerMethods into
// an object that may be used by rpc.Server.
func TableServer(impl TableServerMethods) TableServerStub {
	stub := implTableServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implTableServerStub struct {
	impl TableServerMethods
	gs   *rpc.GlobState
}

func (s implTableServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implTableServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{TableDesc}
}

// TableDesc describes the Table interface.
var TableDesc rpc.InterfaceDesc = descTable

// descTable hides the desc to keep godoc clean.
var descTable = rpc.InterfaceDesc{
	Name:    "Table",
	PkgPath: "v.io/syncbase/v23/services/syncbase",
	Doc:     "// Table represents a collection of Items (rows). Table implements built-in\n// methods such as Glob but currently does not provide any custom methods.\n// Table.Glob operates over the primary keys of Items in the Table.\n// All Permissions checks are performed against the Database Permissions.",
}

// ItemClientMethods is the client interface
// containing Item methods.
//
// Item represents a single row in a Table. The type of data stored in an Item
// is dictated by the Database schema. The relative name of this Item must be
// its encoded primary key.
// All Permissions checks are performed against the Database Permissions.
type ItemClientMethods interface {
	// Get returns the value for this Item.
	Get(*context.T, ...rpc.CallOpt) (*vdl.Value, error)
	// Put writes the given value for this Item. The value's primary key field
	// must match Item.Key().
	Put(ctx *context.T, value *vdl.Value, opts ...rpc.CallOpt) error
	// Delete deletes this Item.
	Delete(*context.T, ...rpc.CallOpt) error
}

// ItemClientStub adds universal methods to ItemClientMethods.
type ItemClientStub interface {
	ItemClientMethods
	rpc.UniversalServiceMethods
}

// ItemClient returns a client stub for Item.
func ItemClient(name string) ItemClientStub {
	return implItemClientStub{name}
}

type implItemClientStub struct {
	name string
}

func (c implItemClientStub) Get(ctx *context.T, opts ...rpc.CallOpt) (o0 *vdl.Value, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Get", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implItemClientStub) Put(ctx *context.T, i0 *vdl.Value, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Put", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implItemClientStub) Delete(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Delete", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

// ItemServerMethods is the interface a server writer
// implements for Item.
//
// Item represents a single row in a Table. The type of data stored in an Item
// is dictated by the Database schema. The relative name of this Item must be
// its encoded primary key.
// All Permissions checks are performed against the Database Permissions.
type ItemServerMethods interface {
	// Get returns the value for this Item.
	Get(rpc.ServerCall) (*vdl.Value, error)
	// Put writes the given value for this Item. The value's primary key field
	// must match Item.Key().
	Put(call rpc.ServerCall, value *vdl.Value) error
	// Delete deletes this Item.
	Delete(rpc.ServerCall) error
}

// ItemServerStubMethods is the server interface containing
// Item methods, as expected by rpc.Server.
// There is no difference between this interface and ItemServerMethods
// since there are no streaming methods.
type ItemServerStubMethods ItemServerMethods

// ItemServerStub adds universal methods to ItemServerStubMethods.
type ItemServerStub interface {
	ItemServerStubMethods
	// Describe the Item interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ItemServer returns a server stub for Item.
// It converts an implementation of ItemServerMethods into
// an object that may be used by rpc.Server.
func ItemServer(impl ItemServerMethods) ItemServerStub {
	stub := implItemServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implItemServerStub struct {
	impl ItemServerMethods
	gs   *rpc.GlobState
}

func (s implItemServerStub) Get(call rpc.ServerCall) (*vdl.Value, error) {
	return s.impl.Get(call)
}

func (s implItemServerStub) Put(call rpc.ServerCall, i0 *vdl.Value) error {
	return s.impl.Put(call, i0)
}

func (s implItemServerStub) Delete(call rpc.ServerCall) error {
	return s.impl.Delete(call)
}

func (s implItemServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implItemServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ItemDesc}
}

// ItemDesc describes the Item interface.
var ItemDesc rpc.InterfaceDesc = descItem

// descItem hides the desc to keep godoc clean.
var descItem = rpc.InterfaceDesc{
	Name:    "Item",
	PkgPath: "v.io/syncbase/v23/services/syncbase",
	Doc:     "// Item represents a single row in a Table. The type of data stored in an Item\n// is dictated by the Database schema. The relative name of this Item must be\n// its encoded primary key.\n// All Permissions checks are performed against the Database Permissions.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Get",
			Doc:  "// Get returns the value for this Item.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // *vdl.Value
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "Put",
			Doc:  "// Put writes the given value for this Item. The value's primary key field\n// must match Item.Key().",
			InArgs: []rpc.ArgDesc{
				{"value", ``}, // *vdl.Value
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Delete",
			Doc:  "// Delete deletes this Item.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
	},
}
