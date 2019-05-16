package generators

import "math/rand"

// generateRandomRuneArray function generates array of runes with selected length, selected scope and random seed
func generateRandomRuneArray(length int, scope string, seed int64) (randomRunes []rune) {
	pickingPool := GenerateKeyboardWritableRunePool(scope)

	// Loop trough character length
	for i := 0; i < length; i++ {
		// Add every loop's index to the seed for avoid getting same number from random's seed
		rand.Seed(seed + int64(i))

		// Get new random number on each pass
		randomNumber := int64(rand.Intn(len(pickingPool) - 1))

		// Append randomly picked runes to return array
		randomRunes = append(randomRunes, pickingPool[randomNumber])
	}

	return
}
