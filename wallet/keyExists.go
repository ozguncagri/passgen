package wallet

import (
	"passgen/config"
)

// IsKeyExists checks global wallet for existence of given key and returns boolean value
func IsKeyExists(keyName string) bool {
	for _, v := range config.GlobalConfig.Wallet {
		if v.Key == keyName {
			return true
		}
	}

	return false
}
