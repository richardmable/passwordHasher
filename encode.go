package main

import (
	"crypto/sha512"
	b64 "encoding/base64"
)

// function to take a password in byte array
// hash with SHA512 and return as Base64
func hashPassword(pwd []byte) string {
	// generate new hash.Hash with SHA512
	hasher := sha512.New()
	hasher.Write(pwd)
	// Convert the byestring array to Base64
	// string w/url safe encoding
	return b64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
