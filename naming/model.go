package naming

import (
	"net"
	"time"

	"veyron.io/veyron/veyron2/context"

	verror "veyron.io/veyron/veyron2/verror2"
)

const (
	pkgPath = "veyron.io/veyron/veyron2/naming"
)

var (
	ErrNameExists              = verror.Register(pkgPath+".nameExists", verror.NoRetry, "{1} {2} Name exists {_}")
	ErrNoSuchName              = verror.Register(pkgPath+".nameDoesntExist", verror.NoRetry, "{1} {2} Name {3} doesn't exist {_}")
	ErrNoSuchNameRoot          = verror.Register(pkgPath+".rootNameDoesntExist", verror.NoRetry, "{1} {2} Namespace root name {3} doesn't exist {_}")
	ErrResolutionDepthExceeded = verror.Register(pkgPath+".resolutionDepthExceeded", verror.NoRetry, "{1} {2} Resolution depth exceeded {_}")
	ErrNoMountTable            = verror.Register(pkgPath+".noMounttable", verror.NoRetry, "{1} {2} No mounttable {_}")
)

// Endpoint represents unique identifiers for entities communicating over a
// network.  End users don't use endpoints - they deal solely with object names,
// with the MountTable providing translation of object names to endpoints.
type Endpoint interface {
	// Network returns "veyron" so that Endpoint can implement net.Addr.
	Network() string

	// String returns a string representation of the endpoint.
	//
	// The String method formats the endpoint as:
	//   @<version>@<version specific fields>@@
	// Where version is an unsigned integer.
	//
	// Version 1 is the current version for network address information:
	//   @1@<protocol>@<address>@<routingid>@@
	// Where protocol is the underlying network protocol (tcp, bluetooth etc.)
	// and address is the address specific to that protocol (host:port for
	// tcp, MAC address for bluetooth etc.)
	//
	// Version 2 is the current version for RPC:
	//   @2@<protocol>@<address>@<routingid>@<rpc version>@<rpc codec>@@
	//
	// Version 3 is the new version for RPC:
	//   @3@<protocol>@<address>@<routingid>@<rpc version>@<rpc codec>@m|s@@
	//
	// Along with Network, this method ensures that Endpoint implements net.Addr.
	String() string

	// VersionedString returns a string in the specified format. If the version
	// number is unsupported, the current 'default' version will be used.
	VersionedString(version int) string

	// RoutingID returns the RoutingID associated with this Endpoint.
	RoutingID() RoutingID

	// Addrs returns a net.Addr whose String method will return the
	// the underlying network address encoded in the endpoint rather than
	// the endpoint string itself.
	// For example, for TCP based endpoints it will return a net.Addr
	// whose network is "tcp" and string representation is <host>:<port>,
	// than the full Veyron endpoint as per the String method above.
	Addr() net.Addr

	// ServesMountTable returns true if this endpoint serves a mount table.
	ServesMountTable() bool
}

// MountedServer represents a server mounted under an object name.
//
// TODO(toddw): Consolidate with VDLMountedServer once vdl supports time.
type MountedServer struct {
	Server  string    // Server is an object address (OA): endpoint + suffix
	Expires time.Time // Absolute time after which the mount expires.
}

// MountEntry represents a name mounted in the mounttable.
//
// TODO(toddw): Consolidate with VDLMountEntry once vdl supports time.
type MountEntry struct {
	// Name is the mounted name.
	Name string
	// Servers (if present) specifies the mounted names (Link is empty).
	Servers []MountedServer
	// mt is true if servers refer to another mount table.
	mt bool
	// An error occurred fulfilling the request.
	Error error
}

// ServesMountTable returns true if the mount entry represents servers that are
// mount tables.
// TODO(p): When the endpoint actually has this fact encoded in, use that.
func (e *MountEntry) ServesMountTable() bool { return e.mt }

// SetServesMountTable sets whether or not this is a mount table.
func (e *MountEntry) SetServesMountTable(v bool) { e.mt = v }

// CacheCtl is a cache control for the resolution cache.
type CacheCtl interface {
	CacheCtl()
}

// DisbleCache disables the resolution cache when set to true and enables if false.
// As a side effect one can flush the cache by disabling and then reenabling it.
type DisableCache bool

func (DisableCache) CacheCtl() {}

// MountOpt is the interface for all Mount options.
type MountOpt interface {
	NSMountOpt()
}

// ReplaceMountOpt requests the mount to replace the previous mount.
type ReplaceMountOpt bool

func (ReplaceMountOpt) NSMountOpt() {}

// ServesMountTableOpt means the target is a mount table.
type ServesMountTableOpt bool

func (ServesMountTableOpt) NSMountOpt()  {}
func (ServesMountTableOpt) EndpointOpt() {}

// TODO(p): Perhaps add an ACL Opt.

// Namespace provides translation from object names to server object addresses.
// It represents the interface to a client side library for the MountTable
// service
type Namespace interface {
	// Mount the server object address under the object name, expiring after
	// the ttl. ttl of zero implies an implementation-specific high value
	// (essentially, forever).
	Mount(ctx context.T, name, server string, ttl time.Duration, opts ...MountOpt) error

	// Unmount the server object address from the object name, or if server
	// is empty, unmount all server OAs from the object name.
	Unmount(ctx context.T, name, server string) error

	// Resolve the object name into its mounted servers.
	Resolve(ctx context.T, name string) (names []string, err error)

	// Resolve the object name into its mounted servers.
	ResolveX(ctx context.T, name string) (entry *MountEntry, err error)

	// ResolveToMountTable resolves the object name into the mounttables
	// directly responsible for the name.
	ResolveToMountTable(ctx context.T, name string) (names []string, err error)

	// ResolveToMountTable resolves the object name into the mounttables
	// directly responsible for the name.
	ResolveToMountTableX(ctx context.T, name string) (entry *MountEntry, err error)

	// FlushCacheEntry flushes resolution information cached for the name.  If
	// anything was flushed it returns true.
	FlushCacheEntry(name string) bool

	// CacheCtl sets controls and returns the current control values.
	CacheCtl(ctls ...CacheCtl) []CacheCtl

	// TODO(caprita): consider adding a version of Unresolve to the
	// IDL-generated stub (in addition to UnresolveStep).

	// Unresolve returns the object name that resolves to the given name.
	// It can be the given name itself, though typically the service at the
	// given name will return the name of a mount table, which is then
	// followed up the namespace ancestry to obtain a name rooted at a
	// 'global' (i.e., widely accessible) mount table.
	Unresolve(ctx context.T, name string) (names []string, err error)

	// Glob returns all names matching pattern.  If recursive is true, it also
	// returns all names below the matching ones.
	Glob(ctx context.T, pattern string) (chan MountEntry, error)

	// SetRoots sets the roots that the local Namespace is
	// relative to. All relative names passed to the methods above
	// will be interpreted as relative to these roots. The roots
	// will be tried in the order that they are specified in the parameter
	// list for SetRoots. Calling SetRoots with no arguments will clear the
	// currently configured set of roots.
	SetRoots(roots ...string) error

	// Roots returns the currently configured roots. An empty slice is
	// returned if no roots are configured.
	Roots() []string
}

// When this prefix is present at the beginning of an object name suffix, the
// server may intercept the request and handle it internally. This is used to
// provide debugging, monitoring and other common functionality across all
// servers. Applications cannot use any name component that starts with this
// prefix.
const ReservedNamePrefix = "__"
