package interactors

import (
	"errors"
	"os"
	"strings"

	"passgen/generators"
	"passgen/helpers"

	"github.com/AlecAivazis/survey"
)

// AskForPassword is asks for master password
func AskForPassword() string {
	password := ""
	prompt := &survey.Password{
		Message: "What is your master password :",
		Help:    "Password should be at least 8 characters long and must be contain upper, lower case letters, digits and symbols",
	}

	err := survey.AskOne(prompt, &password, func(val interface{}) error {
		if helpers.ProperCharacterCounter(val.(string)) < 8 {
			return errors.New("value is too short. Min length is 8")
		}

		if !strings.ContainsAny(val.(string), string(generators.GenerateKeyboardWritableRunePool("U"))) {
			return errors.New("insecure password. Your password must contain at least one upper-case character")
		}

		if !strings.ContainsAny(val.(string), string(generators.GenerateKeyboardWritableRunePool("L"))) {
			return errors.New("insecure password. Your password must contain at least one lower-case character")
		}

		if !strings.ContainsAny(val.(string), string(generators.GenerateKeyboardWritableRunePool("N"))) {
			return errors.New("insecure password. Your password must contain at least one digit")
		}

		if !strings.ContainsAny(val.(string), string(generators.GenerateKeyboardWritableRunePool("S"))) {
			return errors.New("insecure password. Your password must contain at least one symbol")
		}

		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	return strings.TrimSpace(password)
}
