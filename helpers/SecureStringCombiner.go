package helpers

// SecureStringCombiner combines given string's based on values of every rune
// and rune's positions on string and creates unique number for word and sums
// all generated integers into one seed number for random functions
func SecureStringCombiner(values ...string) (seed int64) {
	for _, value := range values {
		for i, v := range value {
			seed += (int64(i) + 2) * int64(rune(v))
		}
	}

	return
}
