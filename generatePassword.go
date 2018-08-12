package main

import "fmt"

func generatePassword(pool, appName, masterPassword string, length int) string {
	var asciiNumberSum int64

	combination := sha3String(fmt.Sprintf(
		"%s%s%s",
		sha3String(appName),
		sha3String(masterPassword),
		sha3String(string(length)),
	))

	for _, v := range combination {
		asciiNumberSum += int64(rune(v))
	}

	return string(generateRandomRuneArray(length, pool, asciiNumberSum))
}
