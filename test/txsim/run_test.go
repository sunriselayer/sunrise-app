//go:build !race

// known race in testnode
// ref: https://github.com/celestiaorg/celestia-app/issues/1369
package txsim_test

import (
	"context"
	"errors"
	"testing"
	"time"

	blob "github.com/sunriselayer/sunrise/api/sunrise/blob/v1"
	"github.com/sunriselayer/sunrise/app/encoding"
	"github.com/sunriselayer/sunrise/test/txsim"
	testencoding "github.com/sunriselayer/sunrise/test/util/encoding"
	"github.com/sunriselayer/sunrise/test/util/testnode"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

func TestTxSimulator(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping TestTxSimulator in short mode.")
	}

	encCfg := encoding.MakeConfig(testencoding.ModuleEncodingRegisters...)
	testCases := []struct {
		name        string
		sequences   []txsim.Sequence
		expMessages map[string]int64
		useFeegrant bool
	}{
		{
			name:      "send sequence",
			sequences: []txsim.Sequence{txsim.NewSendSequence(2, 1000, 100)},
			// we expect at least 5 bank send messages within 30 seconds
			expMessages: map[string]int64{sdk.MsgTypeURL(&bank.MsgSend{}): 5},
		},
		{
			name:      "stake sequence",
			sequences: []txsim.Sequence{txsim.NewStakeSequence(1000)},
			expMessages: map[string]int64{
				sdk.MsgTypeURL(&staking.MsgDelegate{}):                     1,
				sdk.MsgTypeURL(&distribution.MsgWithdrawDelegatorReward{}): 5,
				// NOTE: this sequence also makes redelegations but because the
				// testnet has only one validator, this never happens
			},
		},
		{
			name: "blob sequence",
			sequences: []txsim.Sequence{
				txsim.NewBlobSequence(
					txsim.NewRange(100, 1000),
					txsim.NewRange(1, 3)),
			},
			expMessages: map[string]int64{sdk.MsgTypeURL(&blob.MsgPayForBlobs{}): 10},
		},
		{
			name: "multi blob sequence",
			sequences: txsim.NewBlobSequence(
				txsim.NewRange(1000, 1000),
				txsim.NewRange(3, 3),
			).Clone(4),
			expMessages: map[string]int64{sdk.MsgTypeURL(&blob.MsgPayForBlobs{}): 20},
		},
		{
			name: "multi mixed sequence",
			sequences: append(append(
				txsim.NewSendSequence(2, 1000, 100).Clone(3),
				txsim.NewStakeSequence(1000).Clone(3)...),
				txsim.NewBlobSequence(txsim.NewRange(1000, 1000), txsim.NewRange(1, 3)).Clone(3)...),
			expMessages: map[string]int64{
				sdk.MsgTypeURL(&bank.MsgSend{}):                            15,
				sdk.MsgTypeURL(&staking.MsgDelegate{}):                     2,
				sdk.MsgTypeURL(&distribution.MsgWithdrawDelegatorReward{}): 10,
				sdk.MsgTypeURL(&blob.MsgPayForBlobs{}):                     10,
			},
		},
		{
			name: "multi mixed sequence using feegrant",
			sequences: append(append(
				txsim.NewSendSequence(2, 1000, 100).Clone(3),
				txsim.NewStakeSequence(1000).Clone(3)...),
				txsim.NewBlobSequence(txsim.NewRange(1000, 1000), txsim.NewRange(1, 3)).Clone(3)...),
			expMessages: map[string]int64{
				sdk.MsgTypeURL(&bank.MsgSend{}):                            15,
				sdk.MsgTypeURL(&staking.MsgDelegate{}):                     2,
				sdk.MsgTypeURL(&distribution.MsgWithdrawDelegatorReward{}): 10,
				sdk.MsgTypeURL(&blob.MsgPayForBlobs{}):                     10,
			},
			useFeegrant: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			keyring, rpcAddr, grpcAddr := Setup(t)
			time.Sleep(600 * time.Millisecond)

			opts := txsim.DefaultOptions().
				SuppressLogs().
				WithPollTime(time.Millisecond * 100)
			if tc.useFeegrant {
				opts.UseFeeGrant()
			}

			err := txsim.Run(
				ctx,
				grpcAddr,
				keyring,
				encCfg,
				opts,
				tc.sequences...,
			)
			// Expect all sequences to run for at least 20 seconds without error
			require.True(t, errors.Is(err, context.DeadlineExceeded), err.Error())

			blocks, err := testnode.ReadBlockchain(context.Background(), rpcAddr)
			require.NoError(t, err)
			for _, block := range blocks {
				txHashes, err := testnode.DecodeBlockData(block.Data)
				require.NoError(t, err, block.Height)
				require.GreaterOrEqual(t, len(txHashes), 1)
			}
		})
	}
}

func Setup(t testing.TB) (keyring.Keyring, string, string) {
	t.Helper()

	cfg := testnode.DefaultConfig().WithTimeoutCommit(300 * time.Millisecond).WithFundedAccounts("txsim-master")

	cctx, rpcAddr, grpcAddr := testnode.NewNetwork(t, cfg)

	return cctx.Keyring, rpcAddr, grpcAddr
}
