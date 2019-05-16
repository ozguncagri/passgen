package generators

import "passgen/helpers"

// GeneratePassword function used for generating password with passed arguments
func GeneratePassword(pool, key, masterPassword string, length int) string {
	// Securely generate seed number from SHA512 hash of key, master password, length and character pool
	seedNumber := helpers.SecureStringCombiner(
		helpers.SHA512Calculator(key),
		helpers.SHA512Calculator(masterPassword),
		helpers.SHA512Calculator(string(length)),
		helpers.SHA512Calculator(pool),
	)

	return string(generateRandomRuneArray(length, pool, seedNumber))
}
