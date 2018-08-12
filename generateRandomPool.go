package main

import "strings"

func generateRandomPool(scope string) (randomPool []rune) {
	lowers := []rune("abcdefghijklmnopqrstuvwxyz")
	uppers := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")
	symbols := []rune("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

	//Include lower characters if scope has 'L'
	if strings.Contains(string(scope), "L") {
		randomPool = append(randomPool, lowers...)
	}

	//Include upper characters if scope has 'U'
	if strings.Contains(string(scope), "U") {
		randomPool = append(randomPool, uppers...)
	}

	//Include numerical if scope has 'N'
	if strings.Contains(string(scope), "N") {
		randomPool = append(randomPool, numbers...)
	}

	//Include symbols if scope has 'S'
	if strings.Contains(string(scope), "S") {
		randomPool = append(randomPool, symbols...)
	}

	return
}
