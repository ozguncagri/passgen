package generators

import "fmt"

// GeneratePassword function used for generating password with required information
func GeneratePassword(pool, appName, masterPassword string, length int) string {
	var seedNumber int64

	// combine all required information for password generation seed
	combination := fmt.Sprintf(
		"%s%s%s%s",
		sha3StringGenerator(appName),
		sha3StringGenerator(masterPassword),
		sha3StringGenerator(string(length)),
		sha3StringGenerator(pool),
	)

	// sum all integers correspoding to rune
	for _, v := range combination {
		seedNumber += int64(rune(v))
	}

	return string(generateRandomRuneArray(length, pool, seedNumber))
}
