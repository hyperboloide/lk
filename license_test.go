package lk_test

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hyperboloide/lk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func ExampleNewLicense() {
	// create a new Private key:
	privKey, err := lk.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	// create a license document:
	doc := struct {
		Email string
		End   time.Time
	}{"test@example.com", time.Now().Add(time.Hour * 24 * 365)}

	// marshall the document to bytes:
	ulBytes, err := json.Marshal(doc)
	if err != nil {
		log.Fatal(err)
	}

	// generate your license with the private key:
	license, err := lk.NewLicense(privKey, ulBytes)
	if err != nil {
		log.Fatal(err)

	}

	// encode the new license to b64 and print it:
	str64, err := license.ToB64String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("License b64 encoded:\n%s\n\n", str64)
}

var _ = Describe("License", func() {

	var privateKey *lk.PrivateKey
	var wrongKey *lk.PrivateKey
	var license *lk.License
	var b []byte

	BeforeEach(func() {
		var err error

		privateKey, err = lk.NewPrivateKey()
		Ω(err).To(BeNil())
		Ω(privateKey).ToNot(BeNil())

		wrongKey, err = lk.NewPrivateKey()
		Ω(err).To(BeNil())
		Ω(privateKey).ToNot(BeNil())

		b = make([]byte, 100)
		_, err = rand.Read(b)
		Ω(err).To(BeNil())

		license, err = lk.NewLicense(privateKey, b)
		Ω(err).To(BeNil())
		Ω(license).ToNot(BeNil())
	})

	It("should test a license with bytes", func() {
		b2, err := license.ToBytes()
		Ω(err).To(BeNil())
		l2, err := lk.LicenseFromBytes(b2)
		Ω(err).To(BeNil())
		ok, err := l2.Verify(privateKey.GetPublicKey())
		Ω(err).To(BeNil())
		Ω(ok).To(BeTrue())
		Ω(bytes.Equal(license.Data, l2.Data)).To(BeTrue())

	})

	It("should not validate with wrong key", func() {
		ok, err := license.Verify(wrongKey.GetPublicKey())
		Ω(err).To(BeNil())
		Ω(ok).To(BeFalse())
	})

	It("should test a license with b64", func() {
		b2, err := license.ToB64String()
		Ω(err).To(BeNil())
		l2, err := lk.LicenseFromB64String(b2)
		Ω(err).To(BeNil())
		ok, err := l2.Verify(privateKey.GetPublicKey())
		Ω(err).To(BeNil())
		Ω(ok).To(BeTrue())
		Ω(bytes.Equal(license.Data, l2.Data)).To(BeTrue())

	})

	It("should test a license with b32", func() {
		b2, err := license.ToB32String()
		Ω(err).To(BeNil())
		l2, err := lk.LicenseFromB32String(b2)
		Ω(err).To(BeNil())
		ok, err := l2.Verify(privateKey.GetPublicKey())
		Ω(err).To(BeNil())
		Ω(ok).To(BeTrue())
		Ω(bytes.Equal(license.Data, l2.Data)).To(BeTrue())

	})

})
