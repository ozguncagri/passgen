package interactors

import (
	"errors"
	"os"
	"strconv"

	"gopkg.in/AlecAivazis/survey.v1"
)

// AskForPasswordLength is asks user for the length of the password that you are going to generate
func AskForPasswordLength() int {
	length := 0
	prompt := &survey.Input{
		Message: "What is the length of password :",
		Help:    "Password length should be between 4 and 256 characters",
	}

	err := survey.AskOne(prompt, &length, func(val interface{}) error {
		convertedInt, err := strconv.Atoi(val.(string))
		if err != nil {
			return errors.New("value is not integer")
		}

		if convertedInt < 4 {
			return errors.New("password should be at least 4 characters for supporting pins. (Suggested min. length 16)")
		}

		if convertedInt > 256 {
			return errors.New("it is already unbreakable with our current technology. (Max : 256)")
		}
		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	return length
}
