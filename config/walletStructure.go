package config

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
