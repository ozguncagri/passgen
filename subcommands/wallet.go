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
	"passgen/wallet"
	"strings"
	"unicode/utf8"

	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
)

// Wallet subcommand variable
var Wallet = &cobra.Command{
	Use:   "wallet",
	Short: "Manage wallet items",
	Long:  "Manage wallet items on your wallet config file",
	Run:   walletRunner,
}

var walletPath string

func init() {
	// Get current user's information
	user, userErr := user.Current()
	if userErr != nil {
		log.Fatalln(userErr)
	}

	walletFilePath := user.HomeDir + "/.passgenWallet"

	Wallet.Flags().StringVarP(&walletPath, "wallet-path", "", walletFilePath, "Set different wallet file path")
}

// wallet adds password generation item to the wallet
func walletRunner(cmd *cobra.Command, args []string) {
	//initialize empty wallet
	memoryWallet := wallet.PassgenWallet{
		Version: "1.0",
		Wallet:  make(map[string]wallet.Item),
	}

	// Read encrypted wallet as a byte array
	encryptedWallet, fileReadErr := ioutil.ReadFile(walletPath)
	if fileReadErr != nil {
		walletCreationApproval := false
		prompt := &survey.Confirm{
			Message: "Seems like you don't have any wallet. Would you like to create one?",
		}
		survey.AskOne(prompt, &walletCreationApproval, nil)

		if !walletCreationApproval {
			log.Fatalln("Wallet operations are aborted!")
		}
	}

	passwordSurveyMessage := "What is your wallet's password :"
	passwordSurveyHelp := "Write your wallet's password for unlocking it"
	if fileReadErr != nil {
		passwordSurveyMessage = "Enter password for your wallet :"
		passwordSurveyHelp = "Write a password for loging your new wallet"
	}

	walletPassword := ""
	prompt := &survey.Password{
		Message: passwordSurveyMessage,
		Help:    passwordSurveyHelp,
	}

	surveyErr := survey.AskOne(prompt, &walletPassword, func(val interface{}) error {
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

	// if there is no read error then decrypt and unmarshal wallet
	if fileReadErr == nil {
		decryptedWallet, decryptionErr := helpers.AES256Decrypt(walletPassword, encryptedWallet)
		if decryptionErr != nil {
			log.Fatalln(decryptionErr)
		}

		byteReader := bytes.NewReader(decryptedWallet)

		decoder := gob.NewDecoder(byteReader)
		err := decoder.Decode(&memoryWallet)
		if err != nil {
			log.Fatalln(err)
		}
	}

walletLoop:
	for {
		walletItemKey := ""
		allKeys := []string{"Generate", "Add", "List", "Update", "Remove", "Save & Lock", "Lock Without Saving"}

		prompt := &survey.Select{
			Message: "What do you want to do on your wallet :",
			Help:    "Select what to do on your wallet",
			Options: allKeys,
		}
		err := survey.AskOne(prompt, &walletItemKey, func(val interface{}) error {
			if utf8.RuneCountInString(val.(string)) == 0 {
				return errors.New("this is not valid selection")
			}
			return nil
		})
		if err != nil {
			log.Fatalln(err)
		}

		switch walletItemKey {
		case "Generate":
			wallet.Use(&memoryWallet)
		case "Add":
			wallet.Add(&memoryWallet)
		case "List":
			wallet.List(&memoryWallet)
		case "Update":
			wallet.Update(&memoryWallet)
		case "Remove":
			wallet.Remove(&memoryWallet)
		case "Save & Lock":
			err := wallet.Save(walletPassword, &memoryWallet, walletPath)
			if err != nil {
				log.Println(err)
				continue
			}
			break walletLoop
		case "Lock Without Saving":
			break walletLoop
		}
	}
}
