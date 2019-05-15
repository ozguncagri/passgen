package storage

import (
	"errors"
	"os"
	"passgen/helpers"
	"unicode/utf8"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Update selects and updates storage item
func Update(memoryStorage *PassgenStorage) {
	if len(memoryStorage.Storage) == 0 {
		helpers.NegativePrintf("\nThere is no item in your storage\n\n")
		return
	}

	storageItemKey := ""
	var allKeys []string

	for keys := range memoryStorage.Storage {
		allKeys = append(allKeys, keys)
	}

	prompt := &survey.Select{
		Message: "Choose storage item for update :",
		Help:    "Pick any storage item for updating it's credentials",
		Options: allKeys,
	}
	err := survey.AskOne(prompt, &storageItemKey, func(val interface{}) error {
		if utf8.RuneCountInString(val.(string)) == 0 {
			return errors.New("this is not valid selection")
		}
		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	userName := ""
	userNamePrompt := &survey.Input{
		Message: "User Name :",
		Default: memoryStorage.Storage[storageItemKey].UserName,
	}
	err = survey.AskOne(userNamePrompt, &userName, survey.Required)
	if err != nil {
		os.Exit(1)
	}

	password := ""
	passwordPrompt := &survey.Input{
		Message: "Password :",
		Default: memoryStorage.Storage[storageItemKey].Password,
	}
	err = survey.AskOne(passwordPrompt, &password, survey.Required)
	if err != nil {
		os.Exit(1)
	}

	memoryStorage.Storage[storageItemKey] = Item{
		AppName:  storageItemKey,
		UserName: userName,
		Password: password,
	}

	helpers.ResultPrintf("\nStorage item (%v) updated\n\n", storageItemKey)
}
