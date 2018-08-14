package generators

// GeneratePassword function used for generating password with required information
func GeneratePassword(pool, key, masterPassword string, length int) string {
	var seedNumber int64

	// combine all required information for password generation seed
	combination := sha3StringGenerator(key)
	combination += sha3StringGenerator(masterPassword)
	combination += sha3StringGenerator(string(length))
	combination += sha3StringGenerator(pool)

	// sum all integers correspoding to rune
	for _, v := range combination {
		seedNumber += int64(rune(v))
	}

	return string(generateRandomRuneArray(length, pool, seedNumber))
}
