package usdtapi

type Tx struct {
	// The hex-encoded hash of the transaction
	Txid string
	// The transaction fee in bitcoins
	Fee string
	// The Bitcoin address of the sender
	Sendingaddress string
	// A Bitcoin address used as reference (if any)
	Referenceaddress string
	// Whether the transaction involes an address in the wallet
	Ismine bool
	// The transaction version
	Version uint32
	// The transaction type as number
	Type_int uint32
	// The transaction type as string
	Type string
	// The token property id
	Propertyid uint32
	// Can be divisible
	Divisible bool
	// Amount of the transaction
	Amount string
	// Whether the transaction is valid
	Valid bool
	// The hash of the block that contains the transaction
	Blockhash string
	// The timestamp of the block that contains the transaction
	Blocktime uint32
	// The position of the transaction in the block
	Positioninblock uint64
	// The height of the block that contains the transaction
	Block uint64
	// The number of transaction confirmations
	Confirmations uint64
}

type balanceResult struct {
	Balance, Reserved string
}
