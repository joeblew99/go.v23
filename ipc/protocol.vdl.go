// This file was auto-generated by the veyron vdl tool.
// Source: protocol.vdl

package ipc

// Request describes the request header sent by the client to the server.  A
// non-zero request header is sent at the beginning of the RPC call, followed by
// the positional args.  Thereafter a zero request header is sent before each
// streaming arg, terminated by a non-zero request header with EndStreamArgs set
// to true.
type Request struct {
	// Suffix of the name used to identify the object hosting the service.
	Suffix string
	// Method to invoke on the service.
	Method string
	// NumPosArgs is the number of positional arguments, which follow this message
	// (and any blessings) on the request stream.
	NumPosArgs uint64
	// EndStreamArgs is true iff no more streaming arguments will be sent.  No
	// more data will be sent on the request stream.
	EndStreamArgs bool
	// Timeout is the duration after which the request should be cancelled.  This
	// is a hint to the server, to avoid wasted work.
	//
	// TODO(toddw): Change to time.Time when a built-in idl time package is added.
	Timeout int64
	// HasBlessing is true iff a blessing credential, bound to the identity of
	// the server (provided by the client) appears immediately after this request
	// message.
	// TODO(toddw,ashankar): Ideally, this would be the blessing itself, but
	// vom currently does not allow for data-type interfaces.
	HasBlessing bool
}

// Response describes the response header sent by the server to the client.  A
// zero response header is sent before each streaming arg.  Thereafter a
// non-zero response header is sent at the end of the RPC call, right before
// the positional results.
type Response struct {
	// Error in processing the RPC at the server. Implies EndStreamResults.
	Error error
	// EndStreamResults is true iff no more streaming results will be sent; the
	// remainder of the stream consists of NumPosResults positional results.
	EndStreamResults bool
	// NumPosResults is the number of positional results, which immediately follow
	// on the response stream.  After these results, no further data will be sent
	// on the response stream.
	NumPosResults uint64
}

const (
	// NoTimeout specifies that no timeout is desired.
	// NoTimeout is set to the maximum value for int64 (i.e. 2^63-1),
	// as opposed to 0 (which may be intended as as instant timeout),
	// or negative integers (which may indicate a bug).
	NoTimeout = int64(9223372036854775807)
)
