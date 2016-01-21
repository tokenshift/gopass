# gopass

Command line tool to generate, salt and hash passwords.

## Password Generation

Without any arguments, outputs several randomly generated passwords (using a
cryptographically secure random number generator):

```
$ gopass
qS&W7*9nfxn9j9MFaaQSDz9a
6KWQWpQfV2j78TNMWHm3bykc
gubastonakubrotradipolydra
correct horse battery staple
```

The first two passwords are completely random strings of 24 characters, with and
without special characters. The third is a "pronounceable" password of *at least*
24 characters, based on "A Random Word Generator for Pronounceable Passwords",
described [here](http://www.nist.gov/customcf/get_pdf.cfm?pub_id=901456) (PDF).
Last is the "XKCD algorithm" described [here](https://xkcd.com/936/).

This behavior can be customized with the following options:

* `--length`, `-l`
  Set the desired password length. Defaults to 24.

## Hashing

Any non-option parameters are treated as passwords to be hashed:

```
$ gopass mypassword
   Salt: c5c48c32ecc90397b537f941
    MD5: 2698c3f52c8b7b704ebee4faf792d3e9
SHA-256: d691e8c6d5126bd0ab324d8c692762882bd96a8655956ce51dbbfa82ce916927
  SHA-3: a280478416311e521baceed8f2be19608ee20e61aca6a8df7587f59b0bdc098905f13e4faec4b27bb269ac3386d2de80ab9b536fbadb2f2dd5314bf0a8fdffad
```

This can be customized with the following options:

* `--salt {type}`, `-s {type}`
** none
   Don't generate a salt string.
** append
   Generate a salt string and append it to the input password.
** prepend
   Generate a salt string and prepend it to the input password.
  Defaults to `append`.
* `--md5`, `--sha256`, `--sha3`
  Generate hashes using the listed algorithms. Any/all of these can be listed in
  a single call; they will be output in the provided order, and the `Type: `
  prefix omitted.