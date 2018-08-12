package main

import (
	"fmt"
)

func main() {
	appName := askForKeyName()

	password := askForPassword()

	pwdLength := askForPasswordLength()

	pool := askForCharPool()

	generatedPassword := generatePassword(pool, appName, password, pwdLength)

	fmt.Printf("\nYour password is : %v\n", generatedPassword)
}
