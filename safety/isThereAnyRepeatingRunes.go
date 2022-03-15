package safety

// isThereAnyRepeatingRunes checks password for any repeating character
func isThereAnyRepeatingRunes(password string) bool {
	runeCountStorage := make(map[rune]int) // define map with key of rune and it's count as int

	// Loop trough rune array to count of some runes
	for _, v := range password {
		runeCountStorage[v]++
	}

	// Loop through map to find any repeating runes
	for _, value := range runeCountStorage {
		if value > 1 {
			return true
		}
	}

	return false
}
