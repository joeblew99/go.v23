// This file was auto-generated by the veyron vdl tool.
// Source: wire.vdl

// TODO(ashankar,ataly): Remove this package before the 0.1 release.
//
// Package wire defines the wire types of veyron security primitives.
//
// It defines types for representinve ECDSA public keys, signatures, caveats
// and identity implementations. It also provides methods for encoding/decoding
// these primitives into the higher level, language specific types. For example,
// converting the wire representation of an ECDSA public key to a Go object
// of type crypto.ecdsa.PublicKey.
package wire

import (
	"veyron.io/veyron/veyron2/security"
)

// KeyCurve defines a namespace for elliptic curves.
type KeyCurve byte

// PublicKey represents an ECDSA PublicKey.
type PublicKey struct {
	// Curve identifies the curve of an ECDSA PublicKey.
	Curve KeyCurve
	// XY is the marshaled form of a point on the curve using the format specified
	// in section 4.3.6 of ANSI X9.62.
	XY []byte
}

// TODO(ataly, ashankar): Get rid of this Caveat type and use security.Caveat instead.
// Caveat represents a veyron2/security.Caveat.
type Caveat struct {
	// Service is a pattern identifying the services that the caveat encoded in Bytes
	// is bound to.
	Service security.BlessingPattern
	// Bytes is a serialized representation of the embedded caveat.
	Bytes []byte
}

// Certificate is a signed assertion binding a name to a public key under a certain set
// of caveats. The issuer of a Certificate is the principal that possesses the private key
// under which the Certificate was signed. The Certificate's signature is over the contents
// of the Certificate along with the Signature of the issuer.
type Certificate struct {
	// Name specified in the certificate, e.g., Alice, Bob. Name must not have the
	// character "/".
	Name string
	// PublicKey is the ECDSA public key associated with the Certificate.
	PublicKey PublicKey
	// Caveats under which the certificate is valid.
	Caveats []Caveat
	// Signature of the contents of the certificate.
	Signature security.Signature
}

// ChainPublicID represents the chain implementation of PublicIDs from veyron.io/veyron/veyron/runtimes/google/security.
// It consists of a chain of certificates such that each certificate is signed using the private key
// of the previous certificate (i.e., issuer). The certificate's signature is over its contents along
// with the signature of the issuer certificate (this is done to bind this certificate to the issuer
// chain). The first certificate of the chain is "self signed". The last certificate's public key is
// considered the PublicID's public key. The chain of certificates, if valid, effectively binds a chain
// of names to the PublicID's public key.
type ChainPublicID struct {
	// Certificates specifies the chain of certificates for the PublicID.
	Certificates []Certificate
}

// ChainPrivateID represents the chain implementation of PrivateIDs from veyron.io/veyron/veyron/runtimes/google/security.
type ChainPrivateID struct {
	// PublicID associated with the PrivateID.
	PublicID ChainPublicID
	// Secret represents the secret integer that together with an ECDSA public key makes up the
	// corresponding private key.
	Secret []byte
}

// KeyCurveP256 describes the NIST P256 curve.
const KeyCurveP256 = KeyCurve(0)
