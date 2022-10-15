package cli

import (
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

// flags for incentives module tx commands.
const (
	FlagStartTime = "start-time"

	FlagPoolID                     = "pool-id"
	FlagAppID                      = "app-id"
	FlagIsMasterPool               = "is-master-pool"
	FlagChildPoolIds               = "child-pool-ids"
	FlagAddLendExternalRewardsFile = "add-lend-external-rewards"
)

// FlagSetCreateGauge returns flags for creating gauge.
func FlagSetCreateGauge() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	// Global Flags
	fs.String(FlagStartTime, "", "Timestamp to begin distribution")

	// Msg Specific Flags - Liquidity GaugeType Flags.
	fs.Uint64(FlagPoolID, 0, "Pool Id")
	fs.Uint64(FlagAppID, 0, "App Id")
	fs.Bool(FlagIsMasterPool, false, "If gauge is for master pool or not, default false")
	fs.String(FlagChildPoolIds, "", "List of child pool ids, default [] i.e all pools")

	return fs
}

func ParseUint64SliceFromString(s string, separator string) ([]uint64, error) {
	stringSlice := strings.Split(s, separator)
	parsedInts := make([]uint64, 0, len(stringSlice))
	for _, s := range strings.Split(s, separator) {
		s = strings.TrimSpace(s)

		parsed, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return []uint64{}, err
		}
		parsedInts = append(parsedInts, parsed)
	}
	return parsedInts, nil
}

func FlagAddExternalLendRewardsMapping() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagAddLendExternalRewardsFile, "", "add lend-external-rewards json file path")
	return fs
}

type createAddLendExternalRewardsInputs struct {
	AppID              string `json:"app_id"`
	CPoolID            string `json:"c_pool_id"`
	AssetID            string `json:"asset_id"`
	TotalRewards       string `json:"total_rewards"`
	MasterPoolId       string `json:"master_pool_id"`
	Duration           string `json:"duration"`
	MinLockupTime      string `json:"min_lockup_time"`
	CSwapAppID         string `json:"c_swap_app_id"`
	CSwapMinLockAmount string `json:"c_swap_min_lock_amount"`
	Title              string
	Description        string
	Deposit            string
}
