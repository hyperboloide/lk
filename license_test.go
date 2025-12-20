package lk_test

import (
	"bytes"

	"github.com/hyperboloide/lk"
)

func (s *Suite) TestExamples() {
	s.Run("Example complete", Example_complete)
	s.Run("Example license generation", Example_licenseGeneration)
	s.Run("Example license verification", Example_licenseVerification)
}

func (s *Suite) TestLicense() {

	var privateKey *lk.PrivateKey // private key for license generation
	var wrongKey *lk.PrivateKey   // wrong key for verification
	var license *lk.License       // license to be tested
	var theData []byte            // data to be signed

	s.Run("Generate test data", func() {
		var err error

		privateKey, err = lk.NewPrivateKey()
		s.Require().NoError(err)
		s.Require().NotNil(privateKey)

		wrongKey, err = lk.NewPrivateKey()
		s.Require().NoError(err)
		s.Require().NotNil(wrongKey)

		theData = s.RandomBytes(100)

		license, err = lk.NewLicense(privateKey, theData)
		s.Require().NoError(err)
		s.Require().NotNil(license)

		ok, err := license.Verify(privateKey.GetPublicKey())
		s.Require().NoError(err)
		s.Require().True(ok)
	})

	s.Run("Should not validate with wrong key", func() {
		ok, err := license.Verify(wrongKey.GetPublicKey())
		s.Require().NoError(err)
		s.Require().False(ok)
	})

	s.Run("Test license with bytes", func() {
		b2, err := license.ToBytes()
		s.Require().NoError(err)

		l2, err := lk.LicenseFromBytes(b2)
		s.Require().NoError(err)

		ok, err := l2.Verify(privateKey.GetPublicKey())
		s.Require().NoError(err)
		s.Require().True(ok)

		s.Require().True(bytes.Equal(license.Data, l2.Data))
	})

	s.Run("Test license with b64", func() {
		b2, err := license.ToB64String()
		s.Require().NoError(err)

		l2, err := lk.LicenseFromB64String(b2)
		s.Require().NoError(err)

		ok, err := l2.Verify(privateKey.GetPublicKey())
		s.Require().NoError(err)
		s.Require().True(ok)

		s.Require().True(bytes.Equal(license.Data, l2.Data))
	})

	s.Run("should test a license with b32", func() {
		b2, err := license.ToB32String()
		s.Require().NoError(err)

		l2, err := lk.LicenseFromB32String(b2)
		s.Require().NoError(err)

		ok, err := l2.Verify(privateKey.GetPublicKey())
		s.Require().NoError(err)
		s.Require().True(ok)

		s.Require().True(bytes.Equal(license.Data, l2.Data))
	})

	s.Run("should test a license with hex", func() {
		b2, err := license.ToHexString()
		s.Require().NoError(err)

		l2, err := lk.LicenseFromHexString(b2)
		s.Require().NoError(err)

		ok, err := l2.Verify(privateKey.GetPublicKey())
		s.Require().NoError(err)
		s.Require().True(ok)

		s.Require().True(bytes.Equal(license.Data, l2.Data))
	})
}
