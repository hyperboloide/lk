package lk

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

// License represents a license with some data and a hash.
type License struct {
	Data []byte
	R    *big.Int
	S    *big.Int
}

// NewLicense create a new license and sign it.
func NewLicense(k *PrivateKey, data []byte) (*License, error) {
	l := &License{
		Data: data,
	}

	if h, err := l.hash(); err != nil {
		return nil, err
	} else if r, s, err := ecdsa.Sign(rand.Reader, k.toEcdsa(), h); err != nil {
		return nil, err
	} else {
		l.R = r
		l.S = s
	}
	return l, nil
}

func (l *License) hash() ([]byte, error) {
	h256 := sha256.New()

	if _, err := h256.Write(l.Data); err != nil {
		return nil, err
	}
	return h256.Sum(nil), nil
}

// Verify the License with the public key
func (l *License) Verify(k *PublicKey) (bool, error) {
	h, err := l.hash()
	if err != nil {
		return false, err
	}

	return ecdsa.Verify(k.toEcdsa(), h, l.R, l.S), nil
}

// ToBytes transforms the licence to a base64 []byte.
func (l *License) ToBytes() ([]byte, error) {
	return toBytes(l)
}

// ToB64String transforms the licence to a base64 []byte.
func (l *License) ToB64String() (string, error) {
	return toB64String(l)
}

// ToB32String transforms the license to a base32 []byte.
func (l *License) ToB32String() (string, error) {
	return toB32String(l)
}

// ToHexString transforms the license to a hexadecimal []byte.
func (l *License) ToHexString() (string, error) {
	return toHexString(l)
}

// LicenseFromBytes returns a License from a []byte.
func LicenseFromBytes(b []byte) (*License, error) {
	l := &License{}
	return l, fromBytes(l, b)
}

// LicenseFromB64String returns a License from a base64 encoded
// string.
func LicenseFromB64String(str string) (*License, error) {
	l := &License{}
	return l, fromB64String(l, str)
}

// LicenseFromB32String returns a License from a base64 encoded
// string.
func LicenseFromB32String(str string) (*License, error) {
	l := &License{}
	return l, fromB32String(l, str)
}

// LicenseFromHexString returns a License from a hexadecimal encoded
// string.
func LicenseFromHexString(str string) (*License, error) {
	l := &License{}
	return l, fromHexString(l, str)
}
