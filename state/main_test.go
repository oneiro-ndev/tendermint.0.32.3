package state_test

import (
	"os"
	"testing"

	"github.com/oneiro-ndev/tendermint.0.32.3/types"
)

func TestMain(m *testing.M) {
	types.RegisterMockEvidencesGlobal()
	os.Exit(m.Run())
}
