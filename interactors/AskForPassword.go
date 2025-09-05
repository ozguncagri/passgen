package interactors

import (
	"errors"
	"log"
	"os"
	"passgen/generators"
	"strings"
	"unicode/utf8"

	"gopkg.in/AlecAivazis/survey.v1"
)

// AskForPassword is asks for master password
func AskForPassword() string {
	password := ""
	prompt := &survey.Password{
		Message: "What is your master password :",
		Help:    "Password should be at least 8 characters long and must be contain upper, lower case letters, digits and symbols",
	}

	upperCasePool, err := generators.GenerateKeyboardWritableRunePool("U")
	if err != nil {
		log.Fatalln(err)
	}

	lowerCasePool, err := generators.GenerateKeyboardWritableRunePool("L")
	if err != nil {
		log.Fatalln(err)
	}

	numberPool, err := generators.GenerateKeyboardWritableRunePool("N")
	if err != nil {
		log.Fatalln(err)
	}

	symbolPool, err := generators.GenerateKeyboardWritableRunePool("S")
	if err != nil {
		log.Fatalln(err)
	}

	err = survey.AskOne(prompt, &password, func(val interface{}) error {
		if utf8.RuneCountInString(val.(string)) < 8 {
			return errors.New("value is too short. Min length is 8")
		}

		if !strings.ContainsAny(val.(string), string(upperCasePool)) {
			return errors.New("insecure password. Your password must contain at least one upper-case character")
		}

		if !strings.ContainsAny(val.(string), string(lowerCasePool)) {
			return errors.New("insecure password. Your password must contain at least one lower-case character")
		}

		if !strings.ContainsAny(val.(string), string(numberPool)) {
			return errors.New("insecure password. Your password must contain at least one digit")
		}

		if !strings.ContainsAny(val.(string), string(symbolPool)) {
			return errors.New("insecure password. Your password must contain at least one symbol")
		}

		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	return password
}
