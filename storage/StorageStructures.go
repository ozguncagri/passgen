package storage

// PassgenStorage is general storage structure
type PassgenStorage struct {
	Version string
	Storage map[string]Item
}

// Item is storage item for storing credentials
type Item struct {
	AppName  string
	UserName string
	Password string
}
