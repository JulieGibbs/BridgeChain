//go:build !testnet && !mocknet && !stagenet
// +build !testnet,!mocknet,!stagenet

package cmd

const (
	Bech32PrefixAccAddr            = "human"
	Bech32PrefixAccPub             = "humanpub"
	Bech32PrefixValAddr            = "humanvaloper"
	Bech32PrefixValPub             = "humanvaloperpub"
	Bech32PrefixConsAddr           = "humanvalcons"
	Bech32PrefixConsPub            = "humanvalconspub"
	DenomRegex                     = `[a-zA-Z][a-zA-Z0-9:\\/\\\-\\_\\.]{2,127}`
	DIVERSIchainCoinType    uint32 = 931
	DIVERSIchainCoinPurpose uint32 = 44
	DIVERSIchainHDPath      string = `m/44'/931'/0'/0/0`
)
