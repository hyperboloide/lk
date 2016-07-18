package lk

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	It("should test private key", func() {

		k, err := NewPrivateKey()
		Ω(err).To(BeNil())

		b, err := k.ToBytes()
		Ω(err).To(BeNil())
		k1, err := PrivateKeyFromBytes(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))

		s, err := k.ToB64String()
		Ω(err).To(BeNil())
		k2, err := PrivateKeyFromB64String(s)
		Ω(err).To(BeNil())
		Ω(k2).To(Equal(k))

	})

	It("should test public key", func() {
		privK, err := NewPrivateKey()
		Ω(err).To(BeNil())

		k := privK.GetPublicKey()
		Ω(k).ToNot(BeNil())

		b := k.ToBytes()
		k1, err := PublicKeyFromBytes(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))

		s := k.ToB64String()
		Ω(s).ToNot(Equal(""))
		k2, err := PublicKeyFromB64String(s)
		Ω(err).To(BeNil())
		Ω(k2).To(Equal(k))

	})

})
