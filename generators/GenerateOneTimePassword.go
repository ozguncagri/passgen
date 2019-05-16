package generators

import (
	"crypto/rand"
	"log"
	"math/big"
	"strings"
	"unicode/utf8"
)

// GenerateOneTimePassword function used for generating password with required information
func GenerateOneTimePassword(pool string, length int) string {
	generatedPassword := strings.Builder{}
	randomPool := GenerateKeyboardWritableRunePool(pool)

	// Loop until desired password length
	for utf8.RuneCountInString(generatedPassword.String()) < length {
		// Get secure random number from crypto library between 0 and character pool length
		number, err := rand.Int(rand.Reader, big.NewInt(int64(len(randomPool)-1)))
		if err != nil {
			log.Fatalf("Error while generating secure number : %v", err)
		}

		// Add selected rune to string
		generatedPassword.WriteRune(randomPool[int(number.Int64())])
	}

	return generatedPassword.String()
}
