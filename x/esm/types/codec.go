package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDepositESM{}, "comdex/esm/deposit-esm", nil)
	cdc.RegisterConcrete(&MsgExecuteESM{}, "comdex/esm/execute-esm", nil)
	cdc.RegisterConcrete(&MsgKillRequest{}, "comdex/esm/stop-all-actions", nil)
	cdc.RegisterConcrete(&MsgCollateralRedemptionRequest{}, "comdex/esm/redeem-collateral", nil)
	cdc.RegisterConcrete(&Params{}, "comdex/esm/Params", nil)
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "comdex/esm/MsgUpdateParams")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgDepositESM{},
		&MsgExecuteESM{},
		&MsgKillRequest{},
		&MsgCollateralRedemptionRequest{},
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterLegacyAminoCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	Amino.Seal()
}
