package lk

import (
	"bytes"
	"crypto/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("License", func() {

	It("should test a license", func() {

		k, err := NewPrivateKey()
		Ω(err).To(BeNil())

		b := make([]byte, 100)
		_, err = rand.Read(b)
		Ω(err).To(BeNil())

		l, err := NewLicense(k, b)
		Ω(err).To(BeNil())
		Ω(l).ToNot(BeNil())

		b2, err := l.ToBytes()
		Ω(err).To(BeNil())
		l2, err := LicenseFromBytes(b2)
		Ω(err).To(BeNil())
		ok, err := l2.Verify(k.GetPublicKey())
		Ω(err).To(BeNil())
		Ω(ok).To(BeTrue())
		Ω(bytes.Equal(l.Data, l2.Data)).To(BeTrue())

		str, err := l.ToB64String()
		Ω(err).To(BeNil())
		l3, err := LicenseFromB64String(str)
		Ω(err).To(BeNil())
		ok, err = l3.Verify(k.GetPublicKey())
		Ω(err).To(BeNil())
		Ω(ok).To(BeTrue())
		Ω(bytes.Equal(l.Data, l3.Data)).To(BeTrue())

		k2, err := NewPrivateKey()
		Ω(err).To(BeNil())
		ok, err = l3.Verify(k2.GetPublicKey())
		Ω(err).To(BeNil())
		Ω(ok).To(BeFalse())
	})

})
