package types

import "fmt"

func GetPairPoolAssetLpCoinDenom(denom string) string {
	return fmt.Sprintf("s%s", denom)
}

func GetPairPoolShadowLpCoinDenom(assetDenom string, shadowDenom string) string {
	return fmt.Sprintf("s%s-%s", shadowDenom, assetDenom)
}
