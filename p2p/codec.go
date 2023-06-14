package p2p

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/oneiro-ndev/tendermint.0.32.3/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}
