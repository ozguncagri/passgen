package wallet

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"passgen/helpers"

	"github.com/AlecAivazis/survey"
)

// Update selects and updates wallet item
func Update(memoryWallet *PassgenWallet) {
	if len(memoryWallet.Wallet) == 0 {
		helpers.NegativePrintf("\nThere is no item in your wallet\n\n")
		return
	}

	walletItemKey := ""
	var allKeys []string

	for keys := range memoryWallet.Wallet {
		allKeys = append(allKeys, keys)
	}

	prompt := &survey.Select{
		Message: "Choose wallet item for update :",
		Help:    "Pick any wallet item for change it's character pool and length settings",
		Options: allKeys,
	}
	err := survey.AskOne(prompt, &walletItemKey, func(val interface{}) error {
		if helpers.ProperCharacterCounter(val.(string)) == 0 {
			return errors.New("this is not valid selection")
		}
		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	//Select Char Pool
	charPools := []string{}
	charPoolsPrompt := &survey.MultiSelect{
		Message: fmt.Sprintf("Select new character pool items for (%v) :", walletItemKey),
		Help:    "Selection of which character groups will be use for password generation",
		Options: []string{"Upper", "Lower", "Number", "Symbols"},
	}

	err = survey.AskOne(charPoolsPrompt, &charPools, survey.Required)
	if err != nil {
		os.Exit(1)
	}

	for i, v := range charPools {
		charPools[i] = string(v[0])
	}

	selectedCharPool := strings.ToUpper(strings.Join(charPools, ""))

	//ask for length
	length := 0
	lengthPrompt := &survey.Input{
		Message: fmt.Sprintf("What is the new length of password for (%v) :", walletItemKey),
		Help:    "Password length should be between 4 and 256 characters",
	}

	err = survey.AskOne(lengthPrompt, &length, func(val interface{}) error {
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

	memoryWallet.Wallet[walletItemKey] = Item{
		Pool:   selectedCharPool,
		Length: length,
	}

	helpers.ResultPrintf("\nCharacter pool and length values of key (%v) updated\n\n")
}
