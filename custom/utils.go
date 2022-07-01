package custom

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ParseCoinNormalized parses and normalize a cli input for one coin type, returning errors if invalid or on an empty string
// as well.
// Expected format: "{amount}{denomination}"
func ParseCoinNormalized(coinStr string) (coin sdk.Coin, err error) {
	decCoin, err := sdk.ParseDecCoin(coinStr)
	if err != nil {
		return sdk.Coin{}, err
	}

	coin, _ = normalizeDecCoin(decCoin).TruncateDecimal()
	return coin, nil
}

// NormalizeCoins normalize and truncate a list of decimal coins
func ParseCoinsNormalized(coinStr string) (sdk.Coins, error) {
	coins, err := sdk.ParseDecCoins(coinStr)
	if err != nil {
		return sdk.Coins{}, err
	}
	result := make([]sdk.Coin, 0, len(coins))

	for _, coin := range coins {
		newCoin, _ := normalizeDecCoin(coin).TruncateDecimal()
		result = append(result, newCoin)
	}

	return result, nil
}

func normalizeDecCoin(coin sdk.DecCoin) sdk.DecCoin {
	if err := sdk.ValidateDenom(coin.Denom); err != nil {
		return sdk.DecCoin{}
	}

	srcUnit, ok := sdk.GetDenomUnit(coin.Denom)
	if !ok {
		return coin
	}

	if coin.Denom[0] != 'u' {
		denom := "u" + coin.Denom
		dstUnit, ok := sdk.GetDenomUnit(denom)
		if !ok {
			return coin
		}
		return sdk.NewDecCoinFromDec(denom, coin.Amount.Mul(srcUnit).Quo(dstUnit))
	}

	return coin
}
