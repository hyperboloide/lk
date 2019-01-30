package lk_test

import (
	"crypto/rand"

	"github.com/dchest/uniuri"
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

		invalidBytes := make([]byte, 42)
		rand.Read(invalidBytes)
		k2, err := lk.PrivateKeyFromBytes(invalidBytes)
		Ω(err).To(HaveOccurred())
		Ω(k2).To(BeNil())
	})

	It("should test private key b64", func() {
		b, err := k.ToB64String()
		Ω(err).To(BeNil())
		k1, err := lk.PrivateKeyFromB64String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))

		invalidB64Str := uniuri.NewLen(42)
		k2, err := lk.PrivateKeyFromB64String(invalidB64Str)
		Ω(err).To(HaveOccurred())
		Ω(k2).To(BeNil())
	})

	It("should test private key b32", func() {
		b, err := k.ToB32String()
		Ω(err).To(BeNil())
		k1, err := lk.PrivateKeyFromB32String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))

		invalidB32Str := uniuri.NewLen(42)
		k2, err := lk.PrivateKeyFromB32String(invalidB32Str)
		Ω(err).To(HaveOccurred())
		Ω(k2).To(BeNil())
	})

	It("should test private key hex", func() {
		b, err := k.ToHexString()
		Ω(err).To(BeNil())
		k1, err := lk.PrivateKeyFromHexString(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k))

		invalidB32Str := uniuri.NewLen(42)
		k2, err := lk.PrivateKeyFromHexString(invalidB32Str)
		Ω(err).To(HaveOccurred())
		Ω(k2).To(BeNil())
	})

	It("should test pubic key bytes", func() {
		b := k.GetPublicKey().ToBytes()
		k1, err := lk.PublicKeyFromBytes(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k.GetPublicKey()))

		invalidBytes := make([]byte, 42)
		rand.Read(invalidBytes)
		k2, err := lk.PublicKeyFromBytes(invalidBytes)
		Ω(err).To(HaveOccurred())
		Ω(k2).To(BeNil())
	})

	It("should test pubic key b64", func() {
		b := k.GetPublicKey().ToB64String()
		k1, err := lk.PublicKeyFromB64String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k.GetPublicKey()))

		invalidB64Str := uniuri.NewLen(42)
		k2, err := lk.PublicKeyFromB64String(invalidB64Str)
		Ω(err).To(HaveOccurred())
		Ω(k2).To(BeNil())
	})

	It("should test pubic key b32", func() {
		b := k.GetPublicKey().ToB32String()
		k1, err := lk.PublicKeyFromB32String(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k.GetPublicKey()))

		invalidB32Str := uniuri.NewLen(42)
		k2, err := lk.PublicKeyFromB32String(invalidB32Str)
		Ω(err).To(HaveOccurred())
		Ω(k2).To(BeNil())
	})

	It("should test pubic key hex", func() {
		b := k.GetPublicKey().ToHexString()
		k1, err := lk.PublicKeyFromHexString(b)
		Ω(err).To(BeNil())
		Ω(k1).To(Equal(k.GetPublicKey()))

		invalidHexStr := uniuri.NewLen(42)
		k2, err := lk.PublicKeyFromHexString(invalidHexStr)
		Ω(err).To(HaveOccurred())
		Ω(k2).To(BeNil())
	})

})
