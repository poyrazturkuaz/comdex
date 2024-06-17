package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateGauge{}, "comdex/rewards/MsgCreateGauge", nil)
	cdc.RegisterConcrete(&ActivateExternalRewardsLockers{}, "comdex/rewards/activateExternalRewardsLockers", nil)
	cdc.RegisterConcrete(&ActivateExternalRewardsVault{}, "comdex/rewards/activateExternalRewardsVault", nil)
	cdc.RegisterConcrete(&ActivateExternalRewardsLend{}, "comdex/rewards/activateExternalRewardsLend", nil)
	cdc.RegisterConcrete(&ActivateExternalRewardsStableMint{}, "comdex/rewards/activateExternalRewardsStableMint", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateGauge{},
		&ActivateExternalRewardsLockers{},
		&ActivateExternalRewardsVault{},
		&ActivateExternalRewardsLend{},
		&ActivateExternalRewardsStableMint{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
