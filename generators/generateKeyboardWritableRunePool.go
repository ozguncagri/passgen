package generators

import (
	"fmt"
	"strings"
)

// GenerateKeyboardWritableRunePool generates a slice of runes depending on the scope
func GenerateKeyboardWritableRunePool(scope string) ([]rune, error) {
	lowers := []rune("abcdefghijklmnopqrstuvwxyz")
	uppers := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")
	symbols := []rune("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

	var pool []rune

	scopeUpper := strings.ToUpper(scope)

	if strings.Contains(scopeUpper, "L") {
		pool = append(pool, lowers...)
	}
	if strings.Contains(scopeUpper, "U") {
		pool = append(pool, uppers...)
	}
	if strings.Contains(scopeUpper, "N") {
		pool = append(pool, numbers...)
	}
	if strings.Contains(scopeUpper, "S") {
		pool = append(pool, symbols...)
	}

	if len(pool) == 0 {
		return nil, fmt.Errorf("invalid scope: no characters available")
	}

	return pool, nil
}
