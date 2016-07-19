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
	gen   = app.Command("gen", "Generates a base32 encoded private key.")
	curve = gen.Flag("curve", "Elliptic curve to use.").
		Short('c').
		Default(cp384).
		Enum(cp224, cp256, cp384, cp521)
	gout = gen.Flag("output", "Output file (if not defined then stdout).").Short('o').String()

	// Sign a new license
	sign = app.Command("sign", "Creates a license.")
	key  = sign.Arg("key", "Path to private key to use.").String()
	sin  = sign.Flag("input", "Input data file (if not defined then stdin).").Short('i').String()
	sout = sign.Flag("output", "Output file (if not defined then stdout).").Short('o').String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	//Generate a private key
	case gen.FullCommand():
		genKey()
	// Sign a license
	case sign.FullCommand():
		signLicense()
	}
}

func signLicense() {
	b, err := ioutil.ReadFile(*key)
	if err != nil {
		log.Fatal(err)
	}
	pk, err := lk.PrivateKeyFromB32String(string(b[:]))
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	if *sin != "" {
		data, err = ioutil.ReadFile(*sin)
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

	if *sout != "" {
		if err := ioutil.WriteFile(*sout, []byte(str), 0600); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := os.Stdout.WriteString(str); err != nil {
			log.Fatal(err)
		}
	}
}

func genKey() {
	switch *curve {
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

	if *gout != "" {
		if err := ioutil.WriteFile(*gout, []byte(str), 0600); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := os.Stdout.WriteString(str); err != nil {
			log.Fatal(err)
		}
	}
}
