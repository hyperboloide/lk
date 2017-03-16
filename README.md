# license-key

[![Build Status](https://travis-ci.org/hyperboloide/lk.svg?branch=master)](https://travis-ci.org/hyperboloide/lk)
[![GoDoc](https://godoc.org/github.com/hyperboloide/lk?status.svg)](https://godoc.org/github.com/hyperboloide/lk)
[![Go Report Card](https://goreportcard.com/badge/github.com/hyperboloide/lk)](https://goreportcard.com/report/github.com/hyperboloide/lk)

A simple licensing library in Golang, that generates license files
containing arbitrary data.

Note that this implementation is quite basic and that in no way it could
prevent someone to hack your software. The goal of this project is only
to provide a convenient way for software publishers to generate license keys
and distribute them without too much hassle for the user.

### How does it works?

1. Generate a private key (and keep it secure).
2. Transform the data you want to provide (end date, user email...) to a byte array (using json or gob for example).
3. The library takes the data and create a cryptographically signed hash that is appended to the data.
4. Convert the result to a Base 64 string and send it to the end user.
5. when the user starts your program verify the signature using a public key.

### Example:

Bellow is a sample function that generate a key pair, signs a message and
verify it.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hyperboloide/lk"
)

// This example function creates a new license.
func main() {
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

```

### lkgen

A command line helper `lkgen` is also provided to generate private keys and create licenses.
Install it with the following command :

```sh
go install github.com/hyperboloide/lk/lkgen
```

See the usage bellow on how to use it (or enter `lkgen --help-long`):

```
usage: lkgen [<flags>] <command> [<args> ...]

A command-line utility to generate private keys and licenses.

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.


  gen [<flags>]
    Generates a base32 encoded private key.

    -c, --curve=p384     Elliptic curve to use.
    -o, --output=OUTPUT  Output file (if not defined then stdout).

  pub [<flags>] <key>
    Get the public key.

    -o, --output=OUTPUT  Output file (if not defined then stdout).

  sign [<flags>] <key>
    Creates a license.

    -i, --input=INPUT    Input data file (if not defined then stdin).
    -o, --output=OUTPUT  Output file (if not defined then stdout).
```
