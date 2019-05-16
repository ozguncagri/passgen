package helpers

// SecureStringCombiner combines given string's based on values of every rune
// and rune's positions on string and creates unique number for word and sums
// all generated integers into one seed number for random functions
func SecureStringCombiner(values ...string) (seed int64) {
	// Loop through all passed sting arguments
	for _, value := range values {
		// loop through all runes on string
		for i, v := range value {
			// Add (character position + 2) * integer representation of rune to the seed
			seed += (int64(i) + 2) * int64(rune(v))
		}
	}

	return
}
