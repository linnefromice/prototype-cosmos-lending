package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/linnefromice/lending/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgRepay_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRepay
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRepay{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRepay{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
