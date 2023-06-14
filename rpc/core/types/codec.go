package core_types

import (
	amino "github.com/tendermint/go-amino"
	"github.com/oneiro-ndev/tendermint.0.32.3/types"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}
