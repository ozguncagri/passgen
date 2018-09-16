package wallet

import (
	"errors"
	"log"

	"passgen/helpers"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Remove is remotes one ore more items from wallet
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
		if helpers.ProperCharacterCounter(val.(string)) == 0 {
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
