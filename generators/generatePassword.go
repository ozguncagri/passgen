package generators

import "passgen/helpers"

// GeneratePassword function used for generating password with required information
func GeneratePassword(pool, key, masterPassword string, length int) string {
	var seedNumber int64

	// combine all required information for password generation seed
	combination := helpers.SHA3StringGenerator(key)
	combination += helpers.SHA3StringGenerator(masterPassword)
	combination += helpers.SHA3StringGenerator(string(length))
	combination += helpers.SHA3StringGenerator(pool)

	// sum all integers correspoding to rune
	for _, v := range combination {
		seedNumber += int64(rune(v))
	}

	return string(generateRandomRuneArray(length, pool, seedNumber))
}
