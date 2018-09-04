package generators

import "passgen/helpers"

// GeneratePassword function used for generating password with required information
func GeneratePassword(pool, key, masterPassword string, length int) string {
	seedNumber := helpers.SecureStringCombiner(
		helpers.SHA3Calculator(key),
		helpers.SHA3Calculator(masterPassword),
		helpers.SHA3Calculator(string(length)),
		helpers.SHA3Calculator(pool),
	)

	return string(generateRandomRuneArray(length, pool, seedNumber))
}
