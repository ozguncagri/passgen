package config

import (
	"log"
)

// PassgenConfig is general configuration structure
type PassgenConfig struct {
	Version string                 `json:"version"`
	Wallet  map[string]*WalletItem `json:"wallet"`
}

// WalletItem is configuration item for generating password
type WalletItem struct {
	Pool   string `json:"pool"`
	Length int    `json:"length"`
}

// GlobalConfig is application wide configuration structure variable
var GlobalConfig *PassgenConfig

// Version is application's version number
var Version = "1.0"

func init() {
	//initialize global config
	GlobalConfig = new(PassgenConfig)
	GlobalConfig.Wallet = make(map[string]*WalletItem)

	err := load(GlobalConfig)
	if err != nil {
		log.Fatalln(err)
	}

	if GlobalConfig.Version == "" {
		GlobalConfig.Version = Version
	}
}
