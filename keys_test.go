package lk_test

import (
	"github.com/hyperboloide/lk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	var k *lk.PrivateKey

	BeforeEach(func() {
		var err error
		k, err = lk.NewPrivateKey()
		Ω(err).To(BeNil())
	})

	It("should test private key bytes", func() {
		b, err := k.ToBytes()
		Ω(err).To(BeNil())
		k1, err := lk.PrivateKeyFromBytes(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))
	})

	It("should test private key b64", func() {
		b, err := k.ToB64String()
		Ω(err).To(BeNil())
		k1, err := lk.PrivateKeyFromB64String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))
	})

	It("should test private key b32", func() {
		b, err := k.ToB32String()
		Ω(err).To(BeNil())
		k1, err := lk.PrivateKeyFromB32String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))
	})

	It("should test pubic key bytes", func() {
		b := k.GetPublicKey().ToBytes()
		k1, err := lk.PublicKeyFromBytes(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k.GetPublicKey()))
	})

	It("should test pubic key b64", func() {
		b := k.GetPublicKey().ToB64String()
		k1, err := lk.PublicKeyFromB64String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k.GetPublicKey()))
	})

	It("should test pubic key b32", func() {
		b := k.GetPublicKey().ToB32String()
		k1, err := lk.PublicKeyFromB32String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k.GetPublicKey()))
	})

})
