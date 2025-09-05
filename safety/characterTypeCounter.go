package safety

import (
	"log"
	"strings"

	"passgen/generators"
)

// characterTypeCounter counts keyboard writable upper-case, lower-case, digit, symbol and non-standard characters
func characterTypeCounter(password string) (upperCount, lowerCount, numberCount, symbolCount, nonStandardCount int) {
	upperCasePool, err := generators.GenerateKeyboardWritableRunePool("U")
	if err != nil {
		log.Fatalln(err)
	}

	lowerCasePool, err := generators.GenerateKeyboardWritableRunePool("L")
	if err != nil {
		log.Fatalln(err)
	}

	numberPool, err := generators.GenerateKeyboardWritableRunePool("N")
	if err != nil {
		log.Fatalln(err)
	}

	symbolPool, err := generators.GenerateKeyboardWritableRunePool("S")
	if err != nil {
		log.Fatalln(err)
	}

	allPool := string(upperCasePool) + string(lowerCasePool) + string(numberPool) + string(symbolPool)

	// Loop through password runes and detect rune types
	for _, v := range password {
		// Count standart upper case letters
		if strings.ContainsAny(string(v), string(upperCasePool)) {
			upperCount++
			continue
		}

		// Count standart lower case letters
		if strings.ContainsAny(string(v), string(lowerCasePool)) {
			lowerCount++
			continue
		}

		// Count numbers
		if strings.ContainsAny(string(v), string(numberPool)) {
			numberCount++
			continue
		}

		// Count mostly used symbols
		if strings.ContainsAny(string(v), string(symbolPool)) {
			symbolCount++
			continue
		}

		// Count out of scope characters
		if !strings.ContainsAny(string(v), string(allPool)) {
			nonStandardCount++
			continue
		}
	}
	return
}
