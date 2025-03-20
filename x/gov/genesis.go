package gov

import (
	"fmt"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/keeper"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

// InitGenesis - store genesis parameters
func InitGenesis(ctx sdk.Context, ak types.AccountKeeper, bk types.BankKeeper, k *keeper.Keeper, data *v1.GenesisState) {
	// check if the gov module account exists
	moduleAcc := k.GetGovernanceAccount(ctx)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	startingProposalID := data.StartingProposalId

	// Calculate total deposits
	totalDeposits := sdk.NewCoins()
	for _, deposit := range data.Deposits {
		totalDeposits = totalDeposits.Add(deposit.Amount...)
	}

	// Set deposits and votes
	for _, deposit := range data.Deposits {
		depositor, err := sdk.AccAddressFromBech32(deposit.Depositor)
		if err != nil {
			panic(err)
		}

		proposalID := deposit.ProposalId

		if err := k.Deposits.Set(ctx, collections.Join(proposalID, depositor), *deposit); err != nil {
			panic(err)
		}
	}

	for _, vote := range data.Votes {
		voter, err := sdk.AccAddressFromBech32(vote.Voter)
		if err != nil {
			panic(err)
		}

		proposalID := vote.ProposalId

		if err := k.Votes.Set(ctx, collections.Join(proposalID, voter), *vote); err != nil {
			panic(err)
		}
	}

	// If there are deposits, ensure the module account has the correct balance
	if !totalDeposits.IsZero() {
		balance := bk.GetAllBalances(ctx, moduleAcc.GetAddress())
		if !balance.Equal(totalDeposits) {
			// If the balance doesn't match, send coins from the mint module
			if err := bk.SendCoinsFromModuleToModule(ctx, "mint", types.ModuleName, totalDeposits); err != nil {
				panic(fmt.Sprintf("failed to send coins from mint module to gov module: %s", err))
			}
		}
	}

	// Set proposals
	for _, proposal := range data.Proposals {
		if err := k.Proposals.Set(ctx, proposal.Id, *proposal); err != nil {
			panic(err)
		}
	}

	if err := k.ProposalID.Set(ctx, startingProposalID); err != nil {
		panic(err)
	}

	// Always set the constitution from genesis state
	if err := k.Constitution.Set(ctx, data.Constitution); err != nil {
		panic(err)
	}

	if data.Params != nil {
		if err := k.Params.Set(ctx, *data.Params); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis - output genesis parameters
func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) (*v1.GenesisState, error) {
	startingProposalID, err := k.ProposalID.Peek(ctx)
	if err != nil {
		return nil, err
	}

	var proposals v1.Proposals
	err = k.Proposals.Walk(ctx, nil, func(_ uint64, value v1.Proposal) (stop bool, err error) {
		proposals = append(proposals, &value)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	constitution, err := k.Constitution.Get(ctx)
	if err != nil {
		if err.Error() == "collections: not found: key 'no_key' of type string" {
			constitution = ""
		} else {
			return nil, err
		}
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var proposalsDeposits v1.Deposits
	err = k.Deposits.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], value v1.Deposit) (stop bool, err error) {
		proposalsDeposits = append(proposalsDeposits, &value)
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	// export proposals votes
	var proposalsVotes v1.Votes
	err = k.Votes.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], value v1.Vote) (stop bool, err error) {
		proposalsVotes = append(proposalsVotes, &value)
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	return &v1.GenesisState{
		StartingProposalId: startingProposalID,
		Deposits:           proposalsDeposits,
		Votes:              proposalsVotes,
		Proposals:          proposals,
		Params:             &params,
		Constitution:       constitution,
	}, nil
}
