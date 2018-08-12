package main

import (
	"fmt"
)

func main() {
	appName := askForString(
		"Enter key name for password: ",
		"Enter key name for password: ")

	password := askForPassword(
		"Enter Master Password: ",
		"Password read error")

	pwdLength := askForInt(
		"Enter length of Password : ",
		"Please enter proper length for password!")

	pool := askForCharPool()

	generatedPassword := generatePassword(pool, appName, password, pwdLength)

	fmt.Printf("\nYour password is : %v\n", generatedPassword)
}
