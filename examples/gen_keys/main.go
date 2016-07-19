package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hyperboloide/lk"
)

func main() {

	// Generates a Private key:
	privKey, err := lk.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	str64, err := privKey.ToB64String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Private Key b64 encoded:\n%s\n\n", str64)

	// Get the public key:
	pubKey := privKey.GetPublicKey()
	fmt.Printf("Public Key b64 encoded:\n%s\n\n", pubKey.ToB64String())

	// Generate a license:
	userLicense := struct {
		Email string
		End   time.Time
	}{"test@example.com", time.Now().Add(time.Hour * 24 * 365)}

	ulBytes, err := json.Marshal(userLicense)
	if err != nil {
		log.Fatal(err)
	}

	license, err := lk.NewLicense(privKey, ulBytes)
	if err != nil {
		log.Fatal(err)
	}
	str64, err = license.ToB64String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("License b64 encoded:\n%s\n\n", str64)

	// Verify the Lisence
	ok, err := license.Verify(pubKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Is the license valid?\n%t\n", ok)
}
