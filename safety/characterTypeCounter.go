package safety

import (
	"strings"

	"passgen/generators"
)

// characterTypeCounter counts keyboard writable upper-case, lower-case, digit, symbol and non-standard characters
func characterTypeCounter(password string) (upperCount, lowerCount, numberCount, symbolCount, nonStandardCount int) {
	// Loop through password runes and detect rune types
	for _, v := range []rune(password) {
		// Count standart upper case letters
		if strings.ContainsAny(string(v), string(generators.GenerateKeyboardWritableRunePool("U"))) {
			upperCount++
			continue
		}

		// Count standart lower case letters
		if strings.ContainsAny(string(v), string(generators.GenerateKeyboardWritableRunePool("L"))) {
			lowerCount++
			continue
		}

		// Count numbers
		if strings.ContainsAny(string(v), string(generators.GenerateKeyboardWritableRunePool("N"))) {
			numberCount++
			continue
		}

		// Count mostly used symbols
		if strings.ContainsAny(string(v), string(generators.GenerateKeyboardWritableRunePool("S"))) {
			symbolCount++
			continue
		}

		// Count out of scope characters
		if !strings.ContainsAny(string(v), string(generators.GenerateKeyboardWritableRunePool("ULNS"))) {
			nonStandardCount++
			continue
		}
	}
	return
}
