package generators

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"unicode/utf8"
)

// GenerateOneTimePassword generates a secure random password of desired length
func GenerateOneTimePassword(scope string, length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than zero")
	}

	pool, err := GenerateKeyboardWritableRunePool(scope)
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	builder.Grow(length) // optimize memory allocation
	count := 0

	for count < length {
		// Generate a secure random index in the pool
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(pool))))
		if err != nil {
			return "", fmt.Errorf("error generating secure random number: %w", err)
		}

		builder.WriteRune(pool[index.Int64()])
		count++
	}

	// Ensure correct rune count in case pool has multi-byte characters
	password := builder.String()
	if utf8.RuneCountInString(password) != length {
		return "", fmt.Errorf("generated password length mismatch")
	}

	return password, nil
}
