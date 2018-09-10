package wallet

import (
	"fmt"
	"log"

	"github.com/ozguncagri/passgen/config"
	"github.com/ozguncagri/passgen/generators"
)

// Add adds password generation item to the wallet
func Add() {
	key := generators.AskKeyNameForWallet()
	pool := generators.AskForCharPool()
	length := generators.AskForPasswordLength()

	config.GlobalConfig.Wallet[key] = config.WalletItem{
		Pool:   pool,
		Length: length,
	}

	err := config.Save()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\nYour choices are saved to wallet.")
}
