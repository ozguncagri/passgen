package helpers

// ProperCharacterCounter splits given string into array of rune and counts it
// for finding proper length for different types of characters including emojis
func ProperCharacterCounter(characters string) int {
	return len([]rune(characters))
}
