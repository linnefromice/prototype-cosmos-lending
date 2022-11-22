package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRepay = "repay"

var _ sdk.Msg = &MsgRepay{}

func NewMsgRepay(creator string, poolId uint64, amount uint64, isShadow bool) *MsgRepay {
	return &MsgRepay{
		Creator:  creator,
		PoolId:   poolId,
		Amount:   amount,
		IsShadow: isShadow,
	}
}

func (msg *MsgRepay) Route() string {
	return RouterKey
}

func (msg *MsgRepay) Type() string {
	return TypeMsgRepay
}

func (msg *MsgRepay) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRepay) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRepay) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
