package generators

import "fmt"

func GeneratePassword(pool, appName, masterPassword string, length int) string {
	var asciiNumberSum int64

	combination := sha3StringGenerator(fmt.Sprintf(
		"%s%s%s",
		sha3StringGenerator(appName),
		sha3StringGenerator(masterPassword),
		sha3StringGenerator(string(length)),
	))

	for _, v := range combination {
		asciiNumberSum += int64(rune(v))
	}

	return string(generateRandomRuneArray(length, pool, asciiNumberSum))
}
