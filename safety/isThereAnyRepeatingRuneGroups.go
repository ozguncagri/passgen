package safety

import (
	"unicode/utf8"
)

// isThereAnyRepeatingRuneGroups checks password for any repeating character groups
func isThereAnyRepeatingRuneGroups(password string) bool {
	pwd := []rune(password)
	cache := make(map[string]int)
	pwdLength := utf8.RuneCountInString(password)
	isThere := false

	for charSelection := 2; charSelection <= pwdLength; charSelection++ {
		for selectionIndex := 0; selectionIndex <= pwdLength-charSelection; selectionIndex++ {
			cache[string(pwd[selectionIndex:selectionIndex+charSelection])]++
		}
	}

	for _, value := range cache {
		if value > 1 {
			isThere = true
		}
	}

	return isThere
}
