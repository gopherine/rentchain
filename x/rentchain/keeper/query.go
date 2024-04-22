package keeper

import (
	"github.com/gopherine/rentchain/x/rentchain/types"
)

var _ types.QueryServer = Keeper{}
