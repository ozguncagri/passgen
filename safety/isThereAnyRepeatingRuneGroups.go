package safety

import (
	"unicode/utf8"
)

// isThereAnyRepeatingRuneGroups checks password for any repeating character groups
func isThereAnyRepeatingRuneGroups(password string) bool {
	pwd := []rune(password)                       // define rune array from current password
	cache := make(map[string]int)                 // define character group cache for for detect repetition
	pwdLength := utf8.RuneCountInString(password) // define length of password

	// Loop through character selection range between 2 and length of password
	for charSelection := 2; charSelection <= pwdLength; charSelection++ {
		// Pick character groups looping through pwd with the range of char
		// selection until reach at the length of password - selection range to
		// avoid index out of range error
		for selectionIndex := 0; selectionIndex <= pwdLength-charSelection; selectionIndex++ {
			// Add selected character groups as a key of cache and increase it's value
			// it will increase same index if it's already exists
			cache[string(pwd[selectionIndex:selectionIndex+charSelection])]++
		}
	}

	// Loop through cache for checking it's values for any repeating character groups
	for _, value := range cache {
		if value > 1 {
			return true
		}
	}

	return false
}
