package wallet

// PassgenWallet is general wallet structure
type PassgenWallet struct {
	Version string
	Wallet  map[string]Item
}

// Item is wallet item for generating password
type Item struct {
	Pool   string
	Length int
}
