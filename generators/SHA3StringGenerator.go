package generators

import (
	"crypto/sha512"
	"encoding/hex"
)

// sha3StringGenerator generates hash of input string and returns hash as string
func sha3StringGenerator(input string) (hashString string) {
	hash := sha512.New()
	hash.Write([]byte(input))
	hashString = hex.EncodeToString(hash.Sum(nil))
	return
}
