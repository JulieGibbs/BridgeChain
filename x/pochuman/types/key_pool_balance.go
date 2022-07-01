package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PoolBalanceKeyPrefix is the prefix to retrieve all PoolBalance
	PoolBalanceKeyPrefix = "PoolBalance/value/"
)

// PoolBalanceKey returns the store key to retrieve a PoolBalance from the index fields
func PoolBalanceKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
