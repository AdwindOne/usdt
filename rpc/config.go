package rpc

// ConnConfig RPC connect configure of the omni core wallet
type ConnConfig struct {
	// The format is <host|ip:[port]>
	Host string
	// Give a username for Authentication
	User string
	// Give a password for Authentication
	Pass string
}
