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
