package lk_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hyperboloide/lk"
)

// This example function creates a new license.
func ExampleLicense() {
	// create a new Private key:
	privateKey, err := lk.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	// create a license document:
	doc := struct {
		Email string
		End   time.Time
	}{
		"test@example.com",
		time.Now().Add(time.Hour * 24 * 365),
	}

	// marshall the document to bytes:
	docBytes, err := json.Marshal(doc)
	if err != nil {
		log.Fatal(err)
	}

	// generate your license with the private key:
	license, err := lk.NewLicense(privateKey, docBytes)
	if err != nil {
		log.Fatal(err)

	}

	// encode the new license to b64 and print it:
	str64, err := license.ToB64String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("License b64 encoded:\n%s\n", str64)

	// get the public key
	publicKey := privateKey.GetPublicKey()

	// validate the license:
	if ok, err := license.Verify(publicKey); err != nil {
		log.Fatal(err)

	} else if ok {
		fmt.Println("Valid license")

	} else {
		log.Fatal("Invalid license")

	}
}
