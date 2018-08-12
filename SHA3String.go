package main

import (
	"crypto/sha512"
	"encoding/hex"
)

func sha3String(input string) (hashString string) {
	hash := sha512.New()
	hash.Write([]byte(input))
	hashString = hex.EncodeToString(hash.Sum(nil))
	return
}
