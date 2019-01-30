package lk

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/big"
)

// PrivateKey is the master key to create the licenses. Keep it in a secure
// location.
type PrivateKey ecdsa.PrivateKey

type pkContainer struct {
	Pub []byte
	D   *big.Int
}

// Curve is the elliptic.Curve to use. Default is elliptic.P384.
var Curve = elliptic.P384

// NewPrivateKey generates a new private key. The default elliptic.Curve used
// is elliptic.P384().
func NewPrivateKey() (*PrivateKey, error) {
	tmp, err := ecdsa.GenerateKey(Curve(), rand.Reader)
	if err != nil {
		return nil, err
	}
	key := PrivateKey(*tmp)
	return &key, nil
}

func (k *PrivateKey) toEcdsa() *ecdsa.PrivateKey {
	r := ecdsa.PrivateKey(*k)
	return &r
}

// ToBytes transforms the private key to a  []byte.
func (k PrivateKey) ToBytes() ([]byte, error) {
	ek := k.toEcdsa()
	c := &pkContainer{
		elliptic.Marshal(
			ek.PublicKey.Curve,
			ek.PublicKey.X,
			ek.PublicKey.Y),
		ek.D,
	}

	return toBytes(c)
}

// ToB64String transforms the private key to a base64 string.
func (k PrivateKey) ToB64String() (string, error) {
	b, err := k.ToBytes()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// ToB32String transforms the private key to a base32 string.
func (k PrivateKey) ToB32String() (string, error) {
	b, err := k.ToBytes()
	if err != nil {
		return "", err
	}
	return base32.StdEncoding.EncodeToString(b), nil
}

// ToHexString transforms the private key to a hexadecimal string
func (k PrivateKey) ToHexString() (string, error) {
	b, err := k.ToBytes()
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// PrivateKeyFromBytes returns a private key from a []byte.
func PrivateKeyFromBytes(b []byte) (*PrivateKey, error) {
	c := &pkContainer{}
	if err := fromBytes(c, b); err != nil {
		return nil, err
	}
	pk, err := PublicKeyFromBytes(c.Pub)
	if err != nil {
		return nil, err
	}
	k := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey(*pk),
		D:         c.D,
	}
	res := PrivateKey(k)
	return &res, nil
}

// PrivateKeyFromB64String returns a private key from a base64 encoded
// string.
func PrivateKeyFromB64String(str string) (*PrivateKey, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromBytes(b)
}

// PrivateKeyFromB32String returns a private key from a base32 encoded
// string.
func PrivateKeyFromB32String(str string) (*PrivateKey, error) {
	b, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromBytes(b)
}

// PrivateKeyFromHexString returns a private key from a hexadecimal encoded
// string.
func PrivateKeyFromHexString(str string) (*PrivateKey, error) {
	b, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromBytes(b)
}

// GetPublicKey returns the PublicKey associated with the private key.
func (k PrivateKey) GetPublicKey() *PublicKey {
	pk := PublicKey(k.PublicKey)
	return &pk
}

// PublicKey is used to check the validity of the licenses. You can share it
// freely.
type PublicKey ecdsa.PublicKey

func (k *PublicKey) toEcdsa() *ecdsa.PublicKey {
	r := ecdsa.PublicKey(*k)
	return &r
}

// ToBytes transforms the public key to a []byte.
func (k PublicKey) ToBytes() []byte {
	// return toBytes(k)
	ek := k.toEcdsa()
	return elliptic.Marshal(ek.Curve, ek.X, ek.Y)
}

// ToB64String transforms the public key to a base64 string.
func (k PublicKey) ToB64String() string {
	return base64.StdEncoding.EncodeToString(
		k.ToBytes(),
	)
}

// ToB32String transforms the public key to a base32 string.
func (k PublicKey) ToB32String() string {
	return base32.StdEncoding.EncodeToString(
		k.ToBytes(),
	)
}

// ToHexString transforms the public key to a hexadecimal string.
func (k PublicKey) ToHexString() string {
	return hex.EncodeToString(
		k.ToBytes(),
	)
}

// PublicKeyFromBytes returns a public key from a []byte.
func PublicKeyFromBytes(b []byte) (*PublicKey, error) {
	x, y := elliptic.Unmarshal(Curve(), b)
	if x == nil {
		return nil, errors.New("invalid key")
	}

	k := ecdsa.PublicKey{
		Curve: Curve(),
		X:     x,
		Y:     y,
	}
	r := PublicKey(k)
	return &r, nil
}

// PublicKeyFromB64String returns a public key from a base64 encoded
// string.
func PublicKeyFromB64String(str string) (*PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return PublicKeyFromBytes(b)
}

// PublicKeyFromB32String returns a public key from a base32 encoded
// string.
func PublicKeyFromB32String(str string) (*PublicKey, error) {
	b, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return PublicKeyFromBytes(b)
}

// PublicKeyFromHexString returns a public key from a hexadecimal encoded
// string.
func PublicKeyFromHexString(str string) (*PublicKey, error) {
	b, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return PublicKeyFromBytes(b)
}
