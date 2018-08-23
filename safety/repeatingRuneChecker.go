package safety

// isThereAnyRepeatingRunes checks password for any repeating character
func isThereAnyRepeatingRunes(password string) bool {
	runeGroup := []rune(password)
	runeCountStorage := make(map[rune]int)

	for _, v := range runeGroup {
		runeCountStorage[v]++
	}

	for _, value := range runeCountStorage {
		if value > 1 {
			return true
		}
	}

	return false
}
