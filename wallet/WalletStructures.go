package wallet

// PassgenWallet is general wallet structure
type PassgenWallet struct {
	Version string          `json:"version"`
	Wallet  map[string]Item `json:"wallet"`
}

// Item is wallet item for generating password
type Item struct {
	Pool   string `json:"pool"`
	Length int    `json:"length"`
}

// GlobalWallet is application wide configuration structure variable
var GlobalWallet PassgenWallet

// IsWalletUnlocked keeps control over wallet object with changing on app
var IsWalletUnlocked = false

// IsWalletExists is semaphore for existence of wallet file
var IsWalletExists = false

func init() {
	//initialize global config's wallet map
	GlobalWallet.Wallet = make(map[string]Item)
}
