package keeper

import(
	
	// "github.com/comdex-official/comdex/types"
	liquidationtypes "github.com/comdex-official/comdex/x/liquidationsV2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

)



func (k Keeper) AuctionActivator(ctx sdk.Context,liquidationData liquidationtypes.LockedVault) error {

	auctionType:=liquidationData.AppId
	

	
	return nil
}