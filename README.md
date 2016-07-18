# license-key

[![Build Status](https://travis-ci.org/hyperboloide/license-key.svg?branch=master)](https://travis-ci.org/hyperboloide/license-key)
[![GoDoc](https://godoc.org/github.com/hyperboloide/qmail?status.svg)](https://godoc.org/github.com/hyperboloide/license-key/lk)

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
