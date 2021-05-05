package keeper

import (
	"github.com/ylgr/my-star/x/mystar/types"
)

var _ types.QueryServer = Keeper{}
