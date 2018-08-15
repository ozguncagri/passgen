package wallet

import (
	"passgen/config"
)

func IsKeyExists(keyName string) bool {
	for _, v := range config.GlobalConfig.Wallet {
		if v.Key == keyName {
			return true
		}
	}

	return false
}
