package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// KeysignVoteDataKeyPrefix is the prefix to retrieve all KeysignVoteData
	KeysignVoteDataKeyPrefix = "KeysignVoteData/value/"
)

// KeysignVoteDataKey returns the store key to retrieve a KeysignVoteData from the index fields
func KeysignVoteDataKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
