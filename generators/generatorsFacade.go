package generators

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"
)

// AskForKeyName is asks for the key name for the password you are going to generate
func AskForKeyName() string {
	name := ""

	prompt := &survey.Input{
		Message: "What is key name for password :",
	}

	err := survey.AskOne(prompt, &name, survey.MinLength(1))

	if err != nil {
		os.Exit(1)
	}

	return name
}

// AskForPassword is asks for master password
func AskForPassword() string {
	password := ""

	prompt := &survey.Password{
		Message: "What is your master password :",
	}

	err := survey.AskOne(prompt, &password, survey.MinLength(8))

	if err != nil {
		os.Exit(1)
	}

	return strings.TrimSpace(password)
}

// AskForPasswordLength is asks user for the length of the password that you are going to generate
func AskForPasswordLength() int {
	length := 0

	prompt := &survey.Input{
		Message: "What is the length of password :",
	}

	err := survey.AskOne(prompt, &length, func(val interface{}) error {
		convertedInt, err := strconv.Atoi(val.(string))
		if err != nil {
			return errors.New("value is not integer")
		}

		if convertedInt < 4 {
			return errors.New("password should be at least 4 characters for supporting pins. (Suggested min. length 16)")
		}
		return nil
	})

	if err != nil {
		os.Exit(1)
	}

	return length
}

// AskForCharPool is ask for which character group are you going to use for the password you are generating
func AskForCharPool() string {
	charPools := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select Character Pool Items :",
		Options: []string{"Upper", "Lower", "Number", "Symbols"},
	}
	err := survey.AskOne(prompt, &charPools, survey.Required)
	if err != nil {
		os.Exit(1)
	}

	for i, v := range charPools {
		charPools[i] = string(v[0])
	}

	return strings.ToUpper(strings.Join(charPools, ""))
}
