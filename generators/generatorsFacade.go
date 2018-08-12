package generators

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"
)

func AskForKeyName() string {
	name := ""

	prompt := &survey.Input{
		Message: "Enter key name for password:",
	}

	err := survey.AskOne(prompt, &name, survey.MinLength(1))

	if err != nil {
		os.Exit(1)
	}

	return name
}

func AskForPassword() string {
	password := ""

	prompt := &survey.Password{
		Message: "Enter Master Password:",
	}

	err := survey.AskOne(prompt, &password, survey.MinLength(8))

	if err != nil {
		os.Exit(1)
	}

	return strings.TrimSpace(password)
}

func AskForPasswordLength() int {
	length := 0

	prompt := &survey.Input{
		Message: "Enter length of Password :",
	}

	err := survey.AskOne(prompt, &length, func(val interface{}) error {
		convertedInt, err := strconv.Atoi(val.(string))
		if err != nil {
			return errors.New("Entered value is not valid integer")
		}

		if convertedInt < 4 {
			return errors.New("Password should be at least 4 characters for supporting pins. (Suggested min. length 16)")
		}
		return nil
	})

	if err != nil {
		os.Exit(1)
	}

	return length
}

func AskForCharPool() string {
	charPools := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select Character Pool Items:",
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
