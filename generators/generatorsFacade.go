package generators

import (
	"errors"
	"os"
	"passgen/config"
	"passgen/helpers"
	"strconv"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"
)

// AskForKeyName is asks for the key name for the password you are going to generate
func AskForKeyName() string {
	name := ""
	prompt := &survey.Input{
		Message: "What is key name for password :",
		Help:    "Key name should be at least 3 characters long",
	}

	err := survey.AskOne(prompt, &name, func(val interface{}) error {
		if helpers.ProperCharacterCounter(val.(string)) < 3 {
			return errors.New("value is too short. Min length is 3")
		}
		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	return name
}

// AskKeyNameForWallet is asks for the key name for the password you are going to generate and checks it in global wallet
func AskKeyNameForWallet() string {
	name := ""
	prompt := &survey.Input{
		Message: "What is key name for password :",
		Help:    "Key name should be at least 3 characters long",
	}

	err := survey.AskOne(prompt, &name, func(val interface{}) error {
		if helpers.ProperCharacterCounter(val.(string)) < 3 {
			return errors.New("value is too short. Min length is 3")
		}

		if _, ok := config.GlobalConfig.Wallet[val.(string)]; ok {
			return errors.New("key is already exists in wallet")
		}

		return nil
	})
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
		Help:    "Password should be at least 8 characters long and must be contain upper, lower case letters, digits and symbols",
	}

	err := survey.AskOne(prompt, &password, func(val interface{}) error {
		if helpers.ProperCharacterCounter(val.(string)) < 8 {
			return errors.New("value is too short. Min length is 8")
		}

		if !strings.ContainsAny(val.(string), string(GenerateKeyboardWritableRunePool("U"))) {
			return errors.New("unsecure password. Your password must contain at least one upper-case character")
		}

		if !strings.ContainsAny(val.(string), string(GenerateKeyboardWritableRunePool("L"))) {
			return errors.New("unsecure password. Your password must contain at least one lower-case character")
		}

		if !strings.ContainsAny(val.(string), string(GenerateKeyboardWritableRunePool("N"))) {
			return errors.New("unsecure password. Your password must contain at least one digit")
		}

		if !strings.ContainsAny(val.(string), string(GenerateKeyboardWritableRunePool("S"))) {
			return errors.New("unsecure password. Your password must contain at least one symbol")
		}

		return nil
	})
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
