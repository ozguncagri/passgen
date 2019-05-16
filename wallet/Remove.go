package wallet

import (
	"errors"
	"log"
	"passgen/helpers"
	"unicode/utf8"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Remove is removes an item from wallet
func Remove(memoryWallet *PassgenWallet) {
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
		Message: "Choose wallet item for remove :",
		Help:    "Pick any added item for removing it from your wallet",
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

	previousLength := len(memoryWallet.Wallet)

	delete(memoryWallet.Wallet, walletItemKey)

	if previousLength > len(memoryWallet.Wallet) {
		helpers.ResultPrintf("\nKey (%v) removed\n\n", walletItemKey)
	}
}
