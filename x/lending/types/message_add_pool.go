package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddPool = "add_pool"

var _ sdk.Msg = &MsgAddPool{}

func NewMsgAddPool(creator string, amount sdk.Coin, active bool) *MsgAddPool {
	return &MsgAddPool{
		Creator: creator,
		Amount:  amount,
		Active:  active,
	}
}

func (msg *MsgAddPool) Route() string {
	return RouterKey
}

func (msg *MsgAddPool) Type() string {
	return TypeMsgAddPool
}

func (msg *MsgAddPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
