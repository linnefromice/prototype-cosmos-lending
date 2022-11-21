package types

import "fmt"

func (m *PairPool) GetAssetLpCoinDenom() string {
	return fmt.Sprintf("s%s", m.AssetLiquidity.Denom)
}

func (m *PairPool) GetShadowLpCoinDenom() string {
	return fmt.Sprintf("s%s-%s", m.ShadowLiquidity.Denom, m.AssetLiquidity.Denom)
}
