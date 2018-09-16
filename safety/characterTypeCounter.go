package safety

import (
	"strings"

	"passgen/generators"
)

// characterTypeCounter counts keyboard writable upper-case, lower-case, digit, symbol and non-standard characters
func characterTypeCounter(password string) (upperCount, lowerCount, numberCount, symbolCount, nonStandardCount int) {
	for _, v := range []rune(password) {
		if strings.ContainsAny(string(v), string(generators.GenerateKeyboardWritableRunePool("U"))) {
			upperCount++
			continue
		}

		if strings.ContainsAny(string(v), string(generators.GenerateKeyboardWritableRunePool("L"))) {
			lowerCount++
			continue
		}

		if strings.ContainsAny(string(v), string(generators.GenerateKeyboardWritableRunePool("N"))) {
			numberCount++
			continue
		}

		if strings.ContainsAny(string(v), string(generators.GenerateKeyboardWritableRunePool("S"))) {
			symbolCount++
			continue
		}

		if !strings.ContainsAny(string(v), string(generators.GenerateKeyboardWritableRunePool("ULNS"))) {
			nonStandardCount++
			continue
		}
	}
	return
}
