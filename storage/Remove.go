package storage

import (
	"errors"
	"log"
	"passgen/helpers"
	"unicode/utf8"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Remove function removes an item from storage
func Remove(memoryStorage *PassgenStorage) {
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
		Message: "Choose storage item for remove :",
		Help:    "Pick an item for removing it from your storage",
		Options: allKeys,
	}
	err := survey.AskOne(prompt, &storageItemKey, func(val interface{}) error {
		if utf8.RuneCountInString(val.(string)) == 0 {
			return errors.New("this is not valid selection")
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	previousLength := len(memoryStorage.Storage)

	delete(memoryStorage.Storage, storageItemKey)

	if previousLength > len(memoryStorage.Storage) {
		helpers.ResultPrintf("\nStorage item (%v) removed\n\n", storageItemKey)
	}
}
