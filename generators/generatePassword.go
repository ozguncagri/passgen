package generators

import (
	"log"
	"strconv"
	"strings"
)

// GeneratePassword function used for generating password with passed arguments
func GeneratePassword(pool, key, masterPassword string, length int) string {
	var stringer strings.Builder
	stringer.WriteString(key)
	stringer.WriteString(masterPassword)
	stringer.WriteString(strconv.Itoa(length))
	stringer.WriteString(pool)

	runes, err := generateRandomRuneArray(length, pool, stringer.String())
	if err != nil {
		log.Fatalln(err)
	}

	output := string(runes)

	return output
}
