package generators

import "math/rand"

// generateRandomRuneArray function generates array of runes with selected length, selected scope and random seed
func generateRandomRuneArray(length int, scope string, seed int64) (randomRunes []rune) {
	pickingPool := GenerateKeyboardWritableRunePool(scope)

	for i := 0; i < length; i++ {
		rand.Seed(seed + int64(i))
		randomNumber := int64(rand.Intn(len(pickingPool) - 1))
		randomRunes = append(randomRunes, pickingPool[randomNumber])
	}

	return
}
