package ante

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestGetTxPriority(t *testing.T) {
	cases := []struct {
		name        string
		fee         sdk.Coins
		gas         int64
		expectedPri int64
	}{
		{
			name:        "1 TIA fee large gas",
			fee:         sdk.NewCoins(sdk.NewInt64Coin("usr", 1_000_000)),
			gas:         1000000,
			expectedPri: 1000000,
		},
		{
			name:        "1 usr fee small gas",
			fee:         sdk.NewCoins(sdk.NewInt64Coin("usr", 1)),
			gas:         1,
			expectedPri: 1000000,
		},
		{
			name:        "2 usr fee small gas",
			fee:         sdk.NewCoins(sdk.NewInt64Coin("usr", 2)),
			gas:         1,
			expectedPri: 2000000,
		},
		{
			name:        "1_000_000 TIA fee normal gas tx",
			fee:         sdk.NewCoins(sdk.NewInt64Coin("usr", 1_000_000_000_000)),
			gas:         75000,
			expectedPri: 13333333333333,
		},
		{
			name:        "0.001 usr gas price",
			fee:         sdk.NewCoins(sdk.NewInt64Coin("usr", 1_000)),
			gas:         1_000_000,
			expectedPri: 1000,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pri := getTxPriority(tc.fee, tc.gas)
			assert.Equal(t, tc.expectedPri, pri)
		})
	}
}