// This file was auto-generated by the veyron vdl tool.
// Source: stats.vdl

// Package stats defines an interface to access statistical information for
// troubleshooting and monitoring purposes.
package stats

import (
	// VDL system imports
	"v.io/core/veyron2"
	"v.io/core/veyron2/context"
	"v.io/core/veyron2/i18n"
	"v.io/core/veyron2/ipc"
	"v.io/core/veyron2/vdl"
	"v.io/core/veyron2/verror"

	// VDL user imports
	"v.io/core/veyron2/services/security/access"
	"v.io/core/veyron2/services/watch"
)

var (
	ErrNoValue = verror.Register("v.io/core/veyron2/services/mgmt/stats.NoValue", verror.NoRetry, "{1:}{2:} object has no value, suffix: {3}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoValue.ID), "{1:}{2:} object has no value, suffix: {3}")
}

// NewErrNoValue returns an error with the ErrNoValue ID.
func NewErrNoValue(ctx *context.T, suffix string) error {
	return verror.New(ErrNoValue, ctx, suffix)
}

// StatsClientMethods is the client interface
// containing Stats methods.
//
// The Stats interface is used to access stats for troubleshooting and
// monitoring purposes. The stats objects are discoverable via the Globbable
// interface and watchable via the GlobWatcher interface.
//
// The types of the object values are implementation specific, but should be
// primarily numeric in nature, e.g. counters, memory usage, latency metrics,
// etc.
type StatsClientMethods interface {
	// GlobWatcher allows a client to receive updates for changes to objects
	// that match a pattern.  See the package comments for details.
	watch.GlobWatcherClientMethods
	// Value returns the current value of an object, or an error. The type
	// of the value is implementation specific.
	// Some objects may not have a value, in which case, Value() returns
	// a NoValue error.
	Value(*context.T, ...ipc.CallOpt) (vdl.AnyRep, error)
}

// StatsClientStub adds universal methods to StatsClientMethods.
type StatsClientStub interface {
	StatsClientMethods
	ipc.UniversalServiceMethods
}

// StatsClient returns a client stub for Stats.
func StatsClient(name string, opts ...ipc.BindOpt) StatsClientStub {
	var client ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(ipc.Client); ok {
			client = clientOpt
		}
	}
	return implStatsClientStub{name, client, watch.GlobWatcherClient(name, client)}
}

type implStatsClientStub struct {
	name   string
	client ipc.Client

	watch.GlobWatcherClientStub
}

func (c implStatsClientStub) c(ctx *context.T) ipc.Client {
	if c.client != nil {
		return c.client
	}
	return veyron2.GetClient(ctx)
}

func (c implStatsClientStub) Value(ctx *context.T, opts ...ipc.CallOpt) (o0 vdl.AnyRep, err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Value", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// StatsServerMethods is the interface a server writer
// implements for Stats.
//
// The Stats interface is used to access stats for troubleshooting and
// monitoring purposes. The stats objects are discoverable via the Globbable
// interface and watchable via the GlobWatcher interface.
//
// The types of the object values are implementation specific, but should be
// primarily numeric in nature, e.g. counters, memory usage, latency metrics,
// etc.
type StatsServerMethods interface {
	// GlobWatcher allows a client to receive updates for changes to objects
	// that match a pattern.  See the package comments for details.
	watch.GlobWatcherServerMethods
	// Value returns the current value of an object, or an error. The type
	// of the value is implementation specific.
	// Some objects may not have a value, in which case, Value() returns
	// a NoValue error.
	Value(ipc.ServerContext) (vdl.AnyRep, error)
}

// StatsServerStubMethods is the server interface containing
// Stats methods, as expected by ipc.Server.
// The only difference between this interface and StatsServerMethods
// is the streaming methods.
type StatsServerStubMethods interface {
	// GlobWatcher allows a client to receive updates for changes to objects
	// that match a pattern.  See the package comments for details.
	watch.GlobWatcherServerStubMethods
	// Value returns the current value of an object, or an error. The type
	// of the value is implementation specific.
	// Some objects may not have a value, in which case, Value() returns
	// a NoValue error.
	Value(ipc.ServerContext) (vdl.AnyRep, error)
}

// StatsServerStub adds universal methods to StatsServerStubMethods.
type StatsServerStub interface {
	StatsServerStubMethods
	// Describe the Stats interfaces.
	Describe__() []ipc.InterfaceDesc
}

// StatsServer returns a server stub for Stats.
// It converts an implementation of StatsServerMethods into
// an object that may be used by ipc.Server.
func StatsServer(impl StatsServerMethods) StatsServerStub {
	stub := implStatsServerStub{
		impl: impl,
		GlobWatcherServerStub: watch.GlobWatcherServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implStatsServerStub struct {
	impl StatsServerMethods
	watch.GlobWatcherServerStub
	gs *ipc.GlobState
}

func (s implStatsServerStub) Value(ctx ipc.ServerContext) (vdl.AnyRep, error) {
	return s.impl.Value(ctx)
}

func (s implStatsServerStub) Globber() *ipc.GlobState {
	return s.gs
}

func (s implStatsServerStub) Describe__() []ipc.InterfaceDesc {
	return []ipc.InterfaceDesc{StatsDesc, watch.GlobWatcherDesc}
}

// StatsDesc describes the Stats interface.
var StatsDesc ipc.InterfaceDesc = descStats

// descStats hides the desc to keep godoc clean.
var descStats = ipc.InterfaceDesc{
	Name:    "Stats",
	PkgPath: "v.io/core/veyron2/services/mgmt/stats",
	Doc:     "// The Stats interface is used to access stats for troubleshooting and\n// monitoring purposes. The stats objects are discoverable via the Globbable\n// interface and watchable via the GlobWatcher interface.\n//\n// The types of the object values are implementation specific, but should be\n// primarily numeric in nature, e.g. counters, memory usage, latency metrics,\n// etc.",
	Embeds: []ipc.EmbedDesc{
		{"GlobWatcher", "v.io/core/veyron2/services/watch", "// GlobWatcher allows a client to receive updates for changes to objects\n// that match a pattern.  See the package comments for details."},
	},
	Methods: []ipc.MethodDesc{
		{
			Name: "Value",
			Doc:  "// Value returns the current value of an object, or an error. The type\n// of the value is implementation specific.\n// Some objects may not have a value, in which case, Value() returns\n// a NoValue error.",
			OutArgs: []ipc.ArgDesc{
				{"", ``}, // vdl.AnyRep
				{"", ``}, // error
			},
			Tags: []vdl.AnyRep{access.Tag("Debug")},
		},
	},
}
