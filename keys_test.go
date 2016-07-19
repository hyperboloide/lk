package lk

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	It("should test private key bytes", func() {
		k, err := NewPrivateKey()
		Ω(err).To(BeNil())

		b, err := k.ToBytes()
		Ω(err).To(BeNil())
		k1, err := PrivateKeyFromBytes(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))
	})

	It("should test private key b64", func() {
		k, err := NewPrivateKey()
		Ω(err).To(BeNil())

		b, err := k.ToB64String()
		Ω(err).To(BeNil())
		k1, err := PrivateKeyFromB64String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))
	})

	It("should test private key b32", func() {
		k, err := NewPrivateKey()
		Ω(err).To(BeNil())

		b, err := k.ToB32String()
		Ω(err).To(BeNil())
		k1, err := PrivateKeyFromB32String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))
	})

	It("should test pubic key bytes", func() {
		privK, err := NewPrivateKey()
		Ω(err).To(BeNil())

		k := privK.GetPublicKey()
		Ω(k).ToNot(BeNil())

		b := k.ToBytes()
		k1, err := PublicKeyFromBytes(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))
	})

	It("should test pubic key b64", func() {
		privK, err := NewPrivateKey()
		Ω(err).To(BeNil())

		k := privK.GetPublicKey()
		Ω(k).ToNot(BeNil())

		b := k.ToB64String()
		k1, err := PublicKeyFromB64String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))
	})

	It("should test pubic key b32", func() {
		privK, err := NewPrivateKey()
		Ω(err).To(BeNil())

		k := privK.GetPublicKey()
		Ω(k).ToNot(BeNil())

		b := k.ToB32String()
		k1, err := PublicKeyFromB32String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))
	})

})
