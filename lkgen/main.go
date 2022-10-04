package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hyperboloide/lk"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("lkgen", "A command-line utility to generate private keys and licenses.")

	// Gen a private key.
	gen    = app.Command("gen", "Generates a base32 encoded private key.")
	genOut = gen.Flag("output", "Output file (if not defined then stdout).").Short('o').String()

	// Pub returns the public key.
	pub    = app.Command("pub", "Get the public key.")
	pubKey = pub.Arg("key", "Path to private key to use.").Required().String()
	pubOut = pub.Flag("output", "Output file (if not defined then stdout).").Short('o').String()

	// Sign a new license
	sign    = app.Command("sign", "Creates a license.")
	signKey = sign.Arg("key", "Path to private key to use.").Required().String()
	signIn  = sign.Flag("input", "Input data file (if not defined then stdin).").Short('i').String()
	signOut = sign.Flag("output", "Output file (if not defined then stdout).").Short('o').String()

	// Verfify a license
	verify       = app.Command("verify", "Verifies a license.")
	verifyPubKey = verify.Arg("key", "Path to the public key to use.").Required().String()
	verifyIn     = verify.Flag("input", "Input license file (if not defined then stdin).").Short('i').String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	//Generate a private key
	case gen.FullCommand():
		genKey()

	case pub.FullCommand():
		publicKey()

	// Sign a license
	case sign.FullCommand():
		signLicense()

	case verify.FullCommand():
		verifyLicense()
	}
}

func publicKey() {
	b, err := os.ReadFile(*pubKey)
	if err != nil {
		log.Fatal(err)
	}

	pk, err := lk.PrivateKeyFromB32String(string(b[:]))
	if err != nil {
		log.Fatal(err)
	}

	key := pk.GetPublicKey()
	str := key.ToB32String()

	if *pubOut != "" {
		if err := os.WriteFile(*pubOut, []byte(str), 0600); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := os.Stdout.WriteString(str); err != nil {
			log.Fatal(err)
		}
	}
}

func signLicense() {
	b, err := os.ReadFile(*signKey)
	if err != nil {
		log.Fatal(err)
	}
	pk, err := lk.PrivateKeyFromB32String(string(b[:]))
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	if *signIn != "" {
		data, err = os.ReadFile(*signIn)
	} else {
		data, err = io.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Fatal(err)
	}

	l, err := lk.NewLicense(pk, data)
	if err != nil {
		log.Fatal(err)
	}

	str, err := l.ToB32String()
	if err != nil {
		log.Fatal(err)
	}

	if *signOut != "" {
		if err := os.WriteFile(*signOut, []byte(str), 0600); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := os.Stdout.WriteString(str); err != nil {
			log.Fatal(err)
		}
	}
}

func genKey() {
	key, err := lk.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	str, err := key.ToB32String()
	if err != nil {
		log.Fatal(err)
	}

	if *genOut != "" {
		if err := os.WriteFile(*genOut, []byte(str), 0600); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := os.Stdout.WriteString(str); err != nil {
			log.Fatal(err)
		}
	}
}

func verifyLicense() {
	b, err := os.ReadFile(*verifyPubKey)
	if err != nil {
		log.Print(*verifyPubKey)
		log.Fatal(err)
	}

	publicKey, err := lk.PublicKeyFromB32String(string(b))
	if err != nil {
		log.Fatal(err)
	}

	if *verifyIn != "" {
		b, err = os.ReadFile(*verifyIn)
	} else {
		b, err = io.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Fatal(err)
	}

	license, err := lk.LicenseFromB32String(string(b))
	if err != nil {
		log.Fatal(err)
	}

	if ok, err := license.Verify(publicKey); err != nil {
		log.Fatal(err)
	} else if !ok {
		log.Fatal("Invalid license signature")
	}
	fmt.Print(string(license.Data))
}
