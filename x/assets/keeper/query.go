package keeper

import (
	"github.com/gopherine/rentchain/x/assets/types"
)

var _ types.QueryServer = Keeper{}
