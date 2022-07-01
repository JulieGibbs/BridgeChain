package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FeeBalanceKeyPrefix is the prefix to retrieve all FeeBalance
	FeeBalanceKeyPrefix = "FeeBalance/value/"
)

// FeeBalanceKey returns the store key to retrieve a FeeBalance from the index fields
func FeeBalanceKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
