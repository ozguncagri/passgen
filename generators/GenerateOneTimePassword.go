package generators

import (
	"crypto/rand"
	"log"
	"math/big"
	"passgen/helpers"
	"strings"
)

// GenerateOneTimePassword function used for generating password with required information
func GenerateOneTimePassword(pool string, length int) string {
	generatedPassword := strings.Builder{}
	randomPool := GenerateKeyboardWritableRunePool(pool)

	for helpers.ProperCharacterCounter(generatedPassword.String()) < length {
		number, err := rand.Int(rand.Reader, big.NewInt(int64(len(randomPool)-1)))
		if err != nil {
			log.Fatalf("Error while generating secure number : %v", err)
		}
		generatedPassword.WriteRune(randomPool[int(number.Int64())])
	}
	return generatedPassword.String()
}
