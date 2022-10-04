# lkgen
`lkgen` is a license generation and validation utility.

It can be used standalone to generate and validate license keys

Install it with the following command :

```sh
go install github.com/hyperboloide/lk/lkgen
```

## Complete flow example:

1. Generate a private key file: `lkgen gen --output=./private.key`
2. Generate a public key to distribute with your app: `lkgen pub ./private.key --output=./pub.key`
3. Create the licence (here we use json but can be anything...): `echo '{"email":"user@example.com", "until":"2023-10-04"}' > license.tmp`
4. sign the license: `lkgen sign --input=./license.tmp --output=./license.signed private.key`. The file `license.signed` is your redistributable license.
5. validate the license:
```sh
lkgen verify --input=./license.signed ./pub.key
{"email":"user@example.com", "until":"2023-10-04"}
echo $?
0
```

## Reference documentation

```
lkgen --help-long
usage: lkgen [<flags>] <command> [<args> ...]

A command-line utility to generate private keys and licenses.

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.


  gen [<flags>]
    Generates a base32 encoded private key.

    -o, --output=OUTPUT  Output file (if not defined then stdout).

  pub [<flags>] <key>
    Get the public key.

    -o, --output=OUTPUT  Output file (if not defined then stdout).

  sign [<flags>] <key>
    Creates a license.

    -i, --input=INPUT    Input data file (if not defined then stdin).
    -o, --output=OUTPUT  Output file (if not defined then stdout).

  verify [<flags>] <key>
    Verifies a license.

    -i, --input=INPUT  Input license file (if not defined then stdin).

```
