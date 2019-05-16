package subcommands

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"passgen/generators"
	"passgen/helpers"
	"passgen/storage"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

// Storage subcommand variable
var Storage = &cobra.Command{
	Use:   "storage",
	Short: "Manage stored passwords",
	Long:  "Manage (non-regeneratable) stored passwords",
	Run:   storageRunner,
}

var storagePath string

func init() {
	// Get current user's information
	user, userErr := user.Current()
	if userErr != nil {
		log.Fatalln(userErr)
	}

	storageFilePath := user.HomeDir + "/.passgen/storage"

	Storage.Flags().StringVarP(&storagePath, "storage-path", "", storageFilePath, "Set different storage file path")
}

func storageRunner(cmd *cobra.Command, args []string) {
	//initialize empty storage
	memoryStorage := storage.PassgenStorage{
		Version: "1.0",
		Storage: make(map[string]storage.Item),
	}

	// Read encrypted storage as a byte array
	encryptedStorage, fileReadErr := ioutil.ReadFile(storagePath)
	if fileReadErr != nil {
		storageCreationApproval := false
		prompt := &survey.Confirm{
			Message: "Seems like you don't have any password storage. Would you like to create one?",
		}
		survey.AskOne(prompt, &storageCreationApproval, nil)

		if !storageCreationApproval {
			log.Fatalln("Storage operations are aborted!")
		}
	}

	// Update survey messages depending on storage's existence
	passwordSurveyMessage := "What is your storage's password :"
	passwordSurveyHelp := "Write your storage's password for unlocking it"
	if fileReadErr != nil {
		passwordSurveyMessage = "Enter password for your storage :"
		passwordSurveyHelp = "Write a password for locking your new storage"
	}

	storagePassword := ""
	prompt := &survey.Password{
		Message: passwordSurveyMessage,
		Help:    passwordSurveyHelp,
	}

	surveyErr := survey.AskOne(prompt, &storagePassword, func(val interface{}) error {
		if utf8.RuneCountInString(val.(string)) < 8 {
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
	if surveyErr != nil {
		os.Exit(1)
	}

	// If there is no read error then decrypt and unmarshal storage
	if fileReadErr == nil {
		decryptedStorage, decryptionErr := helpers.AES256Decrypt(storagePassword, encryptedStorage)
		if decryptionErr != nil {
			log.Fatalln(decryptionErr)
		}

		byteReader := bytes.NewReader(decryptedStorage)

		decoder := gob.NewDecoder(byteReader)
		err := decoder.Decode(&memoryStorage)
		if err != nil {
			log.Fatalln(err)
		}
	}

storageLoop:
	for {
		storageItemKey := ""
		allKeys := []string{"Add", "List", "Update", "Remove", "Save & Lock", "Lock Without Saving"}

		prompt := &survey.Select{
			Message: "What do you want to do on your storage :",
			Help:    "Select what to do on your storage",
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

		switch storageItemKey {
		case "Add":
			storage.Add(&memoryStorage)
		case "List":
			storage.List(&memoryStorage)
		case "Update":
			storage.Update(&memoryStorage)
		case "Remove":
			storage.Remove(&memoryStorage)
		case "Save & Lock":
			err := storage.Save(storagePassword, &memoryStorage, storagePath)
			if err != nil {
				log.Println(err)
				continue
			}
			break storageLoop
		case "Lock Without Saving":
			break storageLoop
		}
	}
}
