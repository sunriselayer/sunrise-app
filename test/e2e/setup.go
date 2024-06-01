package e2e

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/p2p/pex"
	"github.com/cometbft/cometbft/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/sunriselayer/sunrise/app"
	"github.com/sunriselayer/sunrise/app/defaultoverrides"
	"github.com/sunriselayer/sunrise/app/encoding"
	"github.com/sunriselayer/sunrise/pkg/appconsts"
)

type Account struct {
	PubKey        cryptotypes.PubKey
	InitialTokens int64
}

func MakeGenesis(nodes []*Node, accounts []*Account) (types.GenesisDoc, error) {
	encCdc := encoding.MakeConfig(app.ModuleEncodingRegisters...)
	appGenState := app.ModuleBasics().DefaultGenesis(encCdc.Codec)
	bankGenesis := bank.DefaultGenesisState()
	stakingGenesis := staking.DefaultGenesisState()
	slashingGenesis := slashing.DefaultGenesisState()
	genAccs := []auth.GenesisAccount{}
	stakingGenesis.Params.BondDenom = appconsts.BondDenom
	delegations := make([]staking.Delegation, 0, len(nodes))
	valInfo := make([]slashing.SigningInfo, 0, len(nodes))
	balances := make([]bank.Balance, 0, len(accounts))
	var (
		validators  staking.Validators
		totalBonded int64
	)

	// setup the validator information on the state machine
	for _, node := range nodes {
		if !node.IsValidator() || node.StartHeight != 0 {
			continue
		}

		addr := node.AccountKey.PubKey().Address()
		pk, err := cryptocodec.FromTmPubKeyInterface(node.SignerKey.PubKey())
		if err != nil {
			return types.GenesisDoc{}, fmt.Errorf("converting public key for node %s: %w", node.Name, err)
		}
		pkAny, err := codectypes.NewAnyWithValue(pk)
		if err != nil {
			return types.GenesisDoc{}, err
		}

		validators.Validators = append(validators.Validators, staking.Validator{
			OperatorAddress: sdk.ValAddress(addr).String(),
			ConsensusPubkey: pkAny,
			Description: staking.Description{
				Moniker: node.Name,
			},
			Status:          staking.Bonded,
			Tokens:          sdkmath.NewInt(node.SelfDelegation),
			DelegatorShares: sdkmath.LegacyOneDec(),
			// 5% commission
			Commission:        staking.NewCommission(sdkmath.LegacyNewDecWithPrec(5, 2), sdkmath.LegacyOneDec(), sdkmath.LegacyOneDec()),
			MinSelfDelegation: sdkmath.ZeroInt(),
		})
		totalBonded += node.SelfDelegation
		consensusAddr := pk.Address()
		delegations = append(delegations, staking.NewDelegation(sdk.AccAddress(addr).String(), sdk.ValAddress(addr).String(), sdkmath.LegacyOneDec()))
		valInfo = append(valInfo, slashing.SigningInfo{
			Address:              sdk.ConsAddress(consensusAddr).String(),
			ValidatorSigningInfo: slashing.NewValidatorSigningInfo(sdk.ConsAddress(consensusAddr), 1, 0, time.Unix(0, 0), false, 0),
		})
	}
	stakingGenesis.Delegations = delegations
	stakingGenesis.Validators = validators.Validators
	slashingGenesis.SigningInfos = valInfo

	for idx, account := range accounts {
		addr := account.PubKey.Address()
		acc := auth.NewBaseAccount(addr.Bytes(), account.PubKey, uint64(idx), 0)
		genAccs = append(genAccs, acc)
		if account.InitialTokens == 0 {
			return types.GenesisDoc{}, fmt.Errorf("account %s has no initial tokens", addr)
		}
		balances = append(balances, bank.Balance{
			Address: sdk.AccAddress(addr).String(),
			Coins: sdk.NewCoins(
				sdk.NewCoin(appconsts.BondDenom, sdkmath.NewInt(account.InitialTokens)),
			),
		})
	}
	// add bonded amount to bonded pool module account
	balances = append(balances, bank.Balance{
		Address: auth.NewModuleAddress(staking.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(appconsts.BondDenom, sdkmath.NewInt(totalBonded))},
	})
	bankGenesis.Balances = bank.SanitizeGenesisBalances(balances)
	authGenesis := auth.NewGenesisState(auth.DefaultParams(), genAccs)

	// update the original genesis state
	appGenState[bank.ModuleName] = encCdc.Codec.MustMarshalJSON(bankGenesis)
	appGenState[auth.ModuleName] = encCdc.Codec.MustMarshalJSON(authGenesis)
	appGenState[staking.ModuleName] = encCdc.Codec.MustMarshalJSON(stakingGenesis)
	appGenState[slashing.ModuleName] = encCdc.Codec.MustMarshalJSON(slashingGenesis)

	if err := app.ModuleBasics().ValidateGenesis(encCdc.Codec, encCdc.TxConfig, appGenState); err != nil {
		return types.GenesisDoc{}, fmt.Errorf("validating genesis: %w", err)
	}

	appState, err := json.MarshalIndent(appGenState, "", " ")
	if err != nil {
		return types.GenesisDoc{}, fmt.Errorf("marshaling app state: %w", err)
	}

	// Validator set and app hash are set in InitChain
	return types.GenesisDoc{
		ChainID:         "testnet",
		GenesisTime:     time.Now().UTC(),
		ConsensusParams: defaultoverrides.DefaultConsensusParams(),
		AppState:        appState,
		// AppHash is not provided but computed after InitChain
	}, nil
}

func MakeConfig(node *Node) (*config.Config, error) {
	cfg := config.DefaultConfig()
	cfg.Moniker = node.Name
	cfg.RPC.ListenAddress = "tcp://0.0.0.0:26657"
	cfg.P2P.ExternalAddress = fmt.Sprintf("tcp://%v", node.AddressP2P(false))
	cfg.P2P.PersistentPeers = strings.Join(node.InitialPeers, ",")
	cfg.Consensus.TimeoutPropose = time.Second
	cfg.Consensus.TimeoutCommit = time.Second
	cfg.Instrumentation.Prometheus = true
	return cfg, nil
}

func WriteAddressBook(peers []string, file string) error {
	book := pex.NewAddrBook(file, false)
	for _, peer := range peers {
		addr, err := p2p.NewNetAddressString(peer)
		if err != nil {
			return fmt.Errorf("parsing peer address %s: %w", peer, err)
		}
		err = book.AddAddress(addr, addr)
		if err != nil {
			return fmt.Errorf("adding peer address %s: %w", peer, err)
		}
	}
	book.Save()
	return nil
}

func MakeAppConfig(_ *Node) (*serverconfig.Config, error) {
	srvCfg := serverconfig.DefaultConfig()
	srvCfg.MinGasPrices = fmt.Sprintf("0.001%s", appconsts.BondDenom)
	return srvCfg, srvCfg.ValidateBasic()
}
