package types

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

const (
	// ModuleName defines the module name
	ModuleName = "pochuman"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_pochuman"

	// Ethereum Chain
	CHAIN_ETHEREUM = "Ethereum"

	// Human Cosmos Chain
	CHAIN_HUMAN = "Human"

	// Pay Status -- Available
	PAY_AVAILABLE = "Available"

	// Pay Status -- Confirmed
	PAY_CONFIRMED = "Confirmed"

	// Pay Status -- Paid
	PAY_PAID = "Paid"

	// Pay Status -- Failed to Pay
	PAY_FAILED = "FailedToPay"

	// Pay Status -- UnAvailable
	PAY_UNAVAILABLE = "UnAvailable"

	// Pay Status -- UnAvailable
	PAY_KEYSIGNED = "KeySigned"

	MAIN_VALIDATOR_MONIKER = "validator"

	// Human Chain
	Humanchain_Pool_Address = "human17zc58s96rxj79jtqqsnzt3wtx3tern6al2pagq"
)

const (
	VOTE_DIFF        = 60 * 5 // 5 min
	VALID_CNT_SIGNER = 3      // should be 2-3 if it's 4 nodes
)

var (
	// GlobalStoreKeyPrefix is a prefix for global primitive state variable
	GlobalStoreKeyPrefix = []byte{0x00}

	// Requested transaction count store key
	ReqTransactionCountStoreKey = append(GlobalStoreKeyPrefix, []byte("ReqTransactionCount")...)

	// Requested transaction count store key
	ObserveTxCountStoreKey = append(GlobalStoreKeyPrefix, []byte("ObserveTxCount")...)

	// Requested keysign count store key
	KeysignTxCountStoreKey = append(GlobalStoreKeyPrefix, []byte("KeysignTxCount")...)
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// Generate SHA256 encrypted string from plain text
// which will be used as signed key
func EncryptMsgSHA256(plain string) string {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return ""
	}

	// Get public key
	publicKey := privateKey.PublicKey

	// Generate RSA encrypted msg.
	encryptedMsg := RSA_OAEP_Encrypt(plain, publicKey)

	return encryptedMsg
}

// RAS encrypt
func RSA_OAEP_Encrypt(secretMessage string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), label)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(ciphertext)
}
