package config

import (
	"log"
)

// PassgenConfig is general configuration structure
type PassgenConfig struct {
	Version string       `json:"version"`
	Wallet  []WalletItem `json:"wallet"`
}

// WalletItem is configuration item for generating password
type WalletItem struct {
	Key    string `json:"key"`
	Pool   string `json:"pool"`
	Length int    `json:"length"`
}

// GlobalConfig is application wide configuration structure variable
var GlobalConfig PassgenConfig

// AddWalletItem adds new wallet item to global config
func (pC *PassgenConfig) AddWalletItem(key, pool string, length int) {
	pC.Wallet = append(pC.Wallet, WalletItem{
		Key:    key,
		Pool:   pool,
		Length: length,
	})
}

func init() {
	err := load(&GlobalConfig)
	if err != nil {
		log.Fatalln(err)
	}
}
