package lk_test

import (
	"github.com/hyperboloide/lk"
)

func (s *Suite) TestKeys() {
	k, err := lk.NewPrivateKey()
	s.Require().NoError(err)

	s.Run("should test private key bytes", func() {
		b, err := k.ToBytes()
		s.Require().NoError(err)
		k1, err := lk.PrivateKeyFromBytes(b)
		s.Require().NoError(err)
		s.Require().Equal(k1, k)

		invalidBytes := s.RandomBytes(42)
		k2, err := lk.PrivateKeyFromBytes(invalidBytes)
		s.Require().Error(err)
		s.Require().Nil(k2)
	})

	s.Run("should test private key bytes", func() {
		b, err := k.ToBytes()
		s.Require().NoError(err)
		k1, err := lk.PrivateKeyFromBytes(b)
		s.Require().NoError(err)
		s.Require().Equal(k1, k)

		invalidBytes := s.RandomBytes(42)
		k2, err := lk.PrivateKeyFromBytes(invalidBytes)
		s.Require().Error(err)
		s.Require().Nil(k2)
	})

	tc := []struct {
		name               string
		privateKeyToString func(k *lk.PrivateKey) (string, error)
		publicKeyToString  func(k *lk.PublicKey) string
		fromString         func(string) (*lk.PrivateKey, error)
		fromPubStr         func(string) (*lk.PublicKey, error)
		randomStr          func(int) string
	}{
		{
			name: "b64",
			privateKeyToString: func(k *lk.PrivateKey) (string, error) {
				return k.ToB64String()
			},
			publicKeyToString: func(k *lk.PublicKey) string {
				return k.ToB64String()
			},
			fromString: lk.PrivateKeyFromB64String,
			fromPubStr: lk.PublicKeyFromB64String,
			randomStr:  s.RandomB64String,
		},
		{
			name: "b32",
			privateKeyToString: func(k *lk.PrivateKey) (string, error) {
				return k.ToB32String()
			},
			publicKeyToString: func(k *lk.PublicKey) string {
				return k.ToB32String()
			},
			fromString: lk.PrivateKeyFromB32String,
			fromPubStr: lk.PublicKeyFromB32String,
			randomStr:  s.RandomB32String,
		},
		{
			name: "hex",
			privateKeyToString: func(k *lk.PrivateKey) (string, error) {
				return k.ToHexString()
			},
			publicKeyToString: func(k *lk.PublicKey) string {
				return k.ToHexString()
			},
			fromString: lk.PrivateKeyFromHexString,
			fromPubStr: lk.PublicKeyFromHexString,
			randomStr:  s.RandomHexString,
		},
	}

	for _, tc := range tc {
		s.Run("should test private key "+tc.name, func() {
			b, err := tc.privateKeyToString(k)
			s.Require().NoError(err)
			k1, err := tc.fromString(b)
			s.Require().NoError(err)
			s.Require().Equal(k1, k)

			invalidStr := tc.randomStr(42)
			k2, err := tc.fromString(invalidStr)
			s.Require().Error(err)
			s.Require().Nil(k2)
		})

		s.Run("should test pubic key "+tc.name, func() {
			b := tc.publicKeyToString(k.GetPublicKey())
			k1, err := tc.fromPubStr(b)
			s.Require().NoError(err)
			s.Require().Equal(k1, k.GetPublicKey())

			invalidStr := s.RandomHexString(42)
			k2, err := tc.fromPubStr(invalidStr)
			s.Require().Error(err)
			s.Require().Nil(k2)
		})
	}

}
