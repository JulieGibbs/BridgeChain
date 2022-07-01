package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TransactionDataKeyPrefix is the prefix to retrieve all TransactionData
	TransactionDataKeyPrefix = "TransactionData/value/"
)

// TransactionDataKey returns the store key to retrieve a TransactionData from the index fields
func TransactionDataKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
