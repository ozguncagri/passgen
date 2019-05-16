package storage

import (
	"errors"
	"os"
	"passgen/helpers"
	"unicode/utf8"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Add adds credentials item to the storage
func Add(memoryStorage *PassgenStorage) {
	// Ask For Storage Application Name
	appName := ""
	appNamePrompt := &survey.Input{
		Message: "What is your application name :",
		Help:    "Unique application name that you want to store",
	}

	err := survey.AskOne(appNamePrompt, &appName, func(val interface{}) error {
		if utf8.RuneCountInString(val.(string)) < 1 {
			return errors.New("value is required for storage")
		}

		if _, ok := memoryStorage.Storage[val.(string)]; ok {
			return errors.New("application is already exists in storage")
		}

		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	// Ask For Application's User Name
	userName := ""
	userNamePrompt := &survey.Input{
		Message: "What is your user name :",
	}

	err = survey.AskOne(userNamePrompt, &userName, survey.Required)
	if err != nil {
		os.Exit(1)
	}

	// Ask For Application's Password
	password := ""
	passwordPrompt := &survey.Input{
		Message: "What is your password :",
	}

	err = survey.AskOne(passwordPrompt, &password, survey.Required)
	if err != nil {
		os.Exit(1)
	}

	// Add credentials to storage
	memoryStorage.Storage[appName] = Item{
		AppName:  appName,
		UserName: userName,
		Password: password,
	}

	helpers.ResultPrintf("\nYour new credentials are added to storage\n\n")
}
