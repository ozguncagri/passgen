package safety

import "github.com/ozguncagri/passgen/helpers"

// isThereAnyRepeatingRuneGroups checks password for any repeating character groups
func isThereAnyRepeatingRuneGroups(password string) bool {
	pwd := []rune(password)
	cache := make(map[string]int)
	pwdLength := helpers.ProperCharacterCounter(password)
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
