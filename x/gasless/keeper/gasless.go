package keeper

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/comdex-official/comdex/x/gasless/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetAvailableTxPaths(ctx sdk.Context) []string {
	return k.interfaceRegistry.ListImplementations("cosmos.base.v1beta1.Msg")
}

func (k Keeper) GetAllContractInfos(ctx sdk.Context) (contractInfos []wasmtypes.ContractInfo) {
	contractInfos = []wasmtypes.ContractInfo{}
	k.wasmKeeper.IterateContractInfo(ctx, func(aa sdk.AccAddress, ci wasmtypes.ContractInfo) bool {
		contractInfos = append(contractInfos, ci)
		return false
	})
	return contractInfos
}

func (k Keeper) GetAllContractsByCode(ctx sdk.Context, codeID uint64) (contracts []string) {
	contracts = []string{}
	k.wasmKeeper.IterateContractsByCode(ctx, codeID, func(address sdk.AccAddress) bool {
		contracts = append(contracts, address.String())
		return false
	})
	return contracts
}

func (k Keeper) GetAllAvailableContracts(ctx sdk.Context) (availableContracts []types.AvailableContracts) {
	availableContracts = []types.AvailableContracts{}
	contractInfos := k.GetAllContractInfos(ctx)
	for _, ci := range contractInfos {
		contracts := k.GetAllContractsByCode(ctx, ci.CodeID)
		for _, c := range contracts {
			availableContracts = append(availableContracts, types.AvailableContracts{
				CodeID:          ci.CodeID,
				ContractAddress: c,
				ContractLable:   ci.Label,
			})
		}
	}
	return availableContracts
}
