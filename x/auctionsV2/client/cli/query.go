package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"strconv"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/comdex-official/comdex/x/auctionsV2/types"
)

func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group auctions queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		queryAuction(),
		queryAuctions(),
		queryBids(),
		queryBid(),
		queryAuctionParams(),
		queryUserLimitOrderBidsByAssetID(),
		queryLimitOrderBids(),
		queryLimitBidProtocolData(),
		queryAuctionFeesCollectionData(),
		queryLimitBidProtocolDataWithAddress(),
		queryBidsFilter(),
		queryAuctionsHistory(),
	)

	return cmd
}

func queryAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auction [auction id] [history]",
		Short: "Query auction",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			auctionID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			history, err := strconv.ParseBool(args[1])
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			res, err := queryClient.Auction(
				context.Background(),
				&types.QueryAuctionRequest{
					AuctionId: auctionID,
					History:   history,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func queryAuctions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auctions [type] [history]",
		Short: "Query all auctions",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			auctionType, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			history, err := strconv.ParseBool(args[1])
			if err != nil {
				return err
			}
			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			res, err := queryClient.Auctions(
				context.Background(),
				&types.QueryAuctionsRequest{
					AuctionType: auctionType,
					History:     history,
					Pagination:  pagination,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "auctions")
	return cmd
}

func queryBids() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bids [bidder] [bid-type] [history]",
		Short: "Query bids by bidder address",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bidder := args[0]
			bidType, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			history, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			res, err := queryClient.Bids(
				context.Background(),
				&types.QueryBidsRequest{
					Bidder:     bidder,
					BidType:    bidType,
					History:    history,
					Pagination: pagination,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "bids")

	return cmd
}

func queryBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bid [bid-id]",
		Short: "Query bid by bid id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bidID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			res, err := queryClient.Bid(
				context.Background(),
				&types.QueryBidRequest{
					BidId: bidID,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func queryAuctionParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auction-params",
		Short: "Query auction params",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(ctx)
			res, err := queryClient.AuctionParams(
				context.Background(),
				&types.QueryAuctionParamsRequest{},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "auction-params")

	return cmd
}

func queryUserLimitOrderBidsByAssetID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-limit-order-bids-by-asset-id [bidder] [collateral-token-id] [debt-token-id] ",
		Short: "Query limit order bids by bidder address and asset id",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bidder := args[0]
			collateralID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			debtID, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			res, err := queryClient.UserLimitBidsByAssetID(
				context.Background(),
				&types.QueryUserLimitBidsByAssetIDRequest{
					Bidder:            bidder,
					CollateralTokenId: collateralID,
					DebtTokenId:       debtID,
					Pagination:        pagination,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user-limit-order-bids-by-asset-id")

	return cmd
}

func queryLimitOrderBids() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "limit-order-bids [collateral-token-id] [debt-token-id] ",
		Short: "Query limit order bids by asset id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			collateralID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			debtID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			res, err := queryClient.LimitBids(
				context.Background(),
				&types.QueryLimitBidsRequest{
					CollateralTokenId: collateralID,
					DebtTokenId:       debtID,
					Pagination:        pagination,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "limit-order-bids")

	return cmd
}

func queryLimitBidProtocolData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "limit-bid-protocol-data",
		Short: "Query Limit Bid Protocol Data",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := queryClient.LimitBidProtocolData(
				context.Background(),
				&types.QueryLimitBidProtocolDataRequest{
					Pagination: pagination,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "limit-bid-protocol-data")

	return cmd
}

func queryAuctionFeesCollectionData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auction-fees-collection-data",
		Short: "Query Auction Fees Collection Data",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := queryClient.AuctionFeesCollectionData(
				context.Background(),
				&types.QueryAuctionFeesCollectionFromLimitBidTxRequest{
					Pagination: pagination,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "auction-fees-collection-data")

	return cmd
}

func queryLimitBidProtocolDataWithAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "limit-bid-protocol-data-with-user [bidder]",
		Short: "Query Limit Bid Protocol Data with bidder address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			bidder := args[0]
			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := queryClient.LimitBidProtocolDataWithUser(
				context.Background(),
				&types.QueryLimitBidProtocolDataWithUserRequest{
					Bidder:     bidder,
					Pagination: pagination,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "limit-bid-protocol-data-with-user")

	return cmd
}

func queryBidsFilter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bids-filter [bidder] [bid-type] [history]",
		Short: "Query bids by bidder address",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bidder := args[0]
			bidType, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			history, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			res, err := queryClient.Bids(
				context.Background(),
				&types.QueryBidsRequest{
					Bidder:     bidder,
					BidType:    bidType,
					History:    history,
					Pagination: pagination,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "bids")

	return cmd
}

func queryAuctionsHistory() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auctions-history [type]",
		Short: "Query all auctions history",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			auctionType, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(ctx)
			res, err := queryClient.AuctionsHistory(
				context.Background(),
				&types.QueryAuctionsHistoryRequest{
					AuctionType: auctionType,
					Pagination:  pagination,
				},
			)
			if err != nil {
				return err
			}
			return ctx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "auctions-history")
	return cmd
}
