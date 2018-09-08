package helpers

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512Calculator generates hash of input string and returns hash as string
func SHA512Calculator(input string) string {
	hash := sha512.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
