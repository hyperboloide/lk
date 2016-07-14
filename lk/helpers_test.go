package lk

import (
	"crypto/rand"

	// . "github.com/hyperboloide/license-key/lk"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers", func() {

	It("should encode and decode correcty with bytes", func() {

		type TestStruct struct {
			Data []byte
			A    int
			B    int
		}

		b := make([]byte, 20)
		_, err := rand.Read(b)
		Ω(err).To(BeNil())

		t1 := &TestStruct{b, 1, 2}

		r1, err := toBytes(t1)
		Ω(err).To(BeNil())

		t2 := &TestStruct{}
		Ω(fromBytes(t2, r1)).To(BeNil())

		Ω(t2).To(Equal(t1))

	})

	It("should encode and decode correcty with b64 strings", func() {

		type TestStruct struct {
			Data []byte
			A    int
			B    int
		}

		b := make([]byte, 20)
		_, err := rand.Read(b)
		Ω(err).To(BeNil())

		t1 := &TestStruct{b, 1, 2}

		r1, err := toB64String(t1)
		Ω(err).To(BeNil())

		t2 := &TestStruct{}
		Ω(fromB64String(t2, r1)).To(BeNil())

		Ω(t2).To(Equal(t1))

	})

})
