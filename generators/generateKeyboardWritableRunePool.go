package generators

import "strings"

// GenerateKeyboardWritableRunePool is generates array of runes depending on scope
func GenerateKeyboardWritableRunePool(scope string) (randomPool []rune) {
	lowers := []rune("abcdefghijklmnopqrstuvwxyz")
	uppers := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")
	symbols := []rune("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

	// Include lower characters if scope has 'L'
	if strings.Contains(strings.ToUpper(scope), "L") {
		randomPool = append(randomPool, lowers...)
	}

	// Include upper characters if scope has 'U'
	if strings.Contains(strings.ToUpper(scope), "U") {
		randomPool = append(randomPool, uppers...)
	}

	// Include numerical if scope has 'N'
	if strings.Contains(strings.ToUpper(scope), "N") {
		randomPool = append(randomPool, numbers...)
	}

	// Include symbols if scope has 'S'
	if strings.Contains(strings.ToUpper(scope), "S") {
		randomPool = append(randomPool, symbols...)
	}

	return
}
