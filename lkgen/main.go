package main

import (
	"crypto/elliptic"
	"io/ioutil"
	"log"
	"os"

	"github.com/hyperboloide/lk"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	cp224 = "p224"
	cp256 = "p256"
	cp384 = "p384"
	cp521 = "p521"
)

var (
	app = kingpin.New("lkgen", "A command-line utility to generate private keys and licenses to use with github.com/hyperboloide/lk library.")

	// Gen a private key.
	gen      = app.Command("gen", "Generates a base32 encoded private key.")
	genCurve = gen.Flag("curve", "Elliptic curve to use.").
			Short('c').
			Default(cp384).
			Enum(cp224, cp256, cp384, cp521)
	genOut = gen.Flag("output", "Output file (if not defined then stdout).").Short('o').String()

	// Pub returns the public key.
	pub    = app.Command("pub", "Get the public key.")
	pubKey = pub.Arg("key", "Path to private key to use.").String()
	pubOut = pub.Flag("output", "Output file (if not defined then stdout).").Short('o').String()

	// Sign a new license
	sign    = app.Command("sign", "Creates a license.")
	signKey = sign.Arg("key", "Path to private key to use.").String()
	signIn  = sign.Flag("input", "Input data file (if not defined then stdin).").Short('i').String()
	signOut = sign.Flag("output", "Output file (if not defined then stdout).").Short('o').String()
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
	}
}

func publicKey() {
	b, err := ioutil.ReadFile(*pubKey)
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
		if err := ioutil.WriteFile(*pubOut, []byte(str), 0600); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := os.Stdout.WriteString(str); err != nil {
			log.Fatal(err)
		}
	}
}

func signLicense() {
	b, err := ioutil.ReadFile(*signKey)
	if err != nil {
		log.Fatal(err)
	}
	pk, err := lk.PrivateKeyFromB32String(string(b[:]))
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	if *signIn != "" {
		data, err = ioutil.ReadFile(*signIn)
	} else {
		data, err = ioutil.ReadAll(os.Stdin)
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
		if err := ioutil.WriteFile(*signOut, []byte(str), 0600); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := os.Stdout.WriteString(str); err != nil {
			log.Fatal(err)
		}
	}
}

func genKey() {
	switch *genCurve {
	case cp224:
		lk.Curve = elliptic.P224
	case cp256:
		lk.Curve = elliptic.P256
	case cp521:
		lk.Curve = elliptic.P521
	default:
		lk.Curve = elliptic.P384
	}

	key, err := lk.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	str, err := key.ToB32String()
	if err != nil {
		log.Fatal(err)
	}

	if *genOut != "" {
		if err := ioutil.WriteFile(*genOut, []byte(str), 0600); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := os.Stdout.WriteString(str); err != nil {
			log.Fatal(err)
		}
	}
}
