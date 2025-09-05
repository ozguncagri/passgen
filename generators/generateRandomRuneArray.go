package generators

import (
	"fmt"
	"hash/fnv"
	"math/rand"
)

// generateRandomRuneArray generates a deterministic rune array with a string seed
func generateRandomRuneArray(length int, scope string, seed string) ([]rune, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be greater than zero")
	}

	pickingPool, err := GenerateKeyboardWritableRunePool(scope)
	if err != nil {
		return nil, err
	}

	randomRunes := make([]rune, length)
	r := rand.New(rand.NewSource(stringToSeed(seed)))

	for i := 0; i < length; i++ {
		randomRunes[i] = pickingPool[r.Intn(len(pickingPool))]
	}

	return randomRunes, nil
}

// stringToSeed converts a string into an int64 deterministically
func stringToSeed(s string) int64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return int64(h.Sum64())
}
