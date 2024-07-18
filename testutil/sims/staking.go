package sims

import (
	"time"

	gogoany "github.com/cosmos/gogoproto/types/any"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	StakingNotBondedPoolName = "not_bonded_tokens_pool"
	StakingBondedPoolName    = "bonded_tokens_pool"
)

type StakingGenesisState struct {
	// params defines all the parameters of related to deposit.
	Params StakingParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// last_total_power tracks the total amounts of bonded tokens recorded during
	// the previous end block.
	LastTotalPower math.Int `protobuf:"bytes,2,opt,name=last_total_power,json=lastTotalPower,proto3,customtype=cosmossdk.io/math.Int" json:"last_total_power"`
	// validators defines the validator set at genesis.
	Validators []StakingValidator `protobuf:"bytes,4,rep,name=validators,proto3" json:"validators"`
	// delegations defines the delegations active at genesis.
	Delegations []StakingDelegation `protobuf:"bytes,5,rep,name=delegations,proto3" json:"delegations"`
}

func (s *StakingGenesisState) ToProto() (*anypb.Any, error) {
	params, err := s.Params.ToProto()
	if err != nil {
		return nil, err
	}

	// Create a map to hold the protobuf fields
	fields := map[string]interface{}{
		"params":           params,
		"last_total_power": s.LastTotalPower,
		"validators":       s.Validators,
		"delegations":      s.Delegations,
	}

	// Convert the map to a protobuf Struct
	pbStruct, err := structpb.NewStruct(fields)
	if err != nil {
		return nil, err
	}

	// Marshal the Struct into an Any message
	anyMsg, err := anypb.New(pbStruct)
	if err != nil {
		return nil, err
	}

	return anyMsg, nil
}

func ProtoToStakingGenesisState(protoMsg *anypb.Any) (*StakingGenesisState, error) {
	var s structpb.Struct
	if err := protoMsg.UnmarshalTo(&s); err != nil {
		return nil, err
	}

	genesisStake := &StakingGenesisState{}
	// genesisStake.Params = s.Fields["params"].
	// genesisStake.LastTotalPower = s.Fields["last_total_power"].
	// genesisStake.Validators = s.Fields["validators"].
	// genesisStake.Delegations = s.Fields["delegations"].

	return genesisStake, nil
}

type StakingParams struct {
	// unbonding_time is the time duration of unbonding.
	UnbondingTime time.Duration `protobuf:"bytes,1,opt,name=unbonding_time,json=unbondingTime,proto3,stdduration" json:"unbonding_time"`
	// max_validators is the maximum number of validators.
	MaxValidators uint32 `protobuf:"varint,2,opt,name=max_validators,json=maxValidators,proto3" json:"max_validators,omitempty"`
	// max_entries is the max entries for either unbonding delegation or redelegation (per pair/trio).
	MaxEntries uint32 `protobuf:"varint,3,opt,name=max_entries,json=maxEntries,proto3" json:"max_entries,omitempty"`
	// historical_entries is the number of historical entries to persist.
	HistoricalEntries uint32 `protobuf:"varint,4,opt,name=historical_entries,json=historicalEntries,proto3" json:"historical_entries,omitempty"`
	// bond_denom defines the bondable coin denomination.
	BondDenom string `protobuf:"bytes,5,opt,name=bond_denom,json=bondDenom,proto3" json:"bond_denom,omitempty"`
	// min_commission_rate is the chain-wide minimum commission rate that a validator can charge their delegators
	MinCommissionRate math.LegacyDec `protobuf:"bytes,6,opt,name=min_commission_rate,json=minCommissionRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min_commission_rate" yaml:"min_commission_rate"`
}

func (s *StakingParams) ToProto() (*anypb.Any, error) {
	// Create a map to hold the protobuf fields
	fields := map[string]interface{}{
		"unbonding_time":      s.UnbondingTime.String(),
		"max_validators":      s.MaxValidators,
		"max_entries":         s.MaxEntries,
		"historical_entries":  s.HistoricalEntries,
		"bondDenom":           s.BondDenom,
		"min_commission_rate": s.MinCommissionRate.String(),
	}

	// Convert the map to a protobuf Struct
	pbStruct, err := structpb.NewStruct(fields)
	if err != nil {
		return nil, err
	}

	// Marshal the Struct into an Any message
	anyMsg, err := anypb.New(pbStruct)
	if err != nil {
		return nil, err
	}

	return anyMsg, nil
}

type StakingValidator struct {
	// operator_address defines the address of the validator's operator; bech encoded in JSON.
	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	// consensus_pubkey is the consensus public key of the validator, as a Protobuf Any.
	ConsensusPubkey *gogoany.Any `protobuf:"bytes,2,opt,name=consensus_pubkey,json=consensusPubkey,proto3" json:"consensus_pubkey,omitempty"`
	// jailed defined whether the validator has been jailed from bonded status or not.
	Jailed bool `protobuf:"varint,3,opt,name=jailed,proto3" json:"jailed,omitempty"`
	// status is the validator status (bonded/unbonding/unbonded).
	Status int32 `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	// tokens define the delegated tokens (incl. self-delegation).
	Tokens math.Int `protobuf:"bytes,5,opt,name=tokens,proto3,customtype=cosmossdk.io/math.Int" json:"tokens"`
	// delegator_shares defines total shares issued to a validator's delegators.
	DelegatorShares math.LegacyDec `protobuf:"bytes,6,opt,name=delegator_shares,json=delegatorShares,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"delegator_shares"`
	// min_self_delegation is the validator's self declared minimum self delegation.
	MinSelfDelegation math.Int `protobuf:"bytes,11,opt,name=min_self_delegation,json=minSelfDelegation,proto3,customtype=cosmossdk.io/math.Int" json:"min_self_delegation"`
}

func (s *StakingValidator) ToProto() (*anypb.Any, error) {
	// Create a map to hold the protobuf fields
	fields := map[string]interface{}{
		"operator_address":    s.OperatorAddress,
		"consensus_pubkey":    s.ConsensusPubkey,
		"jailed":              s.Jailed,
		"status":              s.Status,
		"tokens":              s.Tokens.String(),
		"delegator_shares":    s.DelegatorShares.String(),
		"min_self_delegation": s.MinSelfDelegation.String(),
	}

	// Convert the map to a protobuf Struct
	pbStruct, err := structpb.NewStruct(fields)
	if err != nil {
		return nil, err
	}

	// Marshal the Struct into an Any message
	anyMsg, err := anypb.New(pbStruct)
	if err != nil {
		return nil, err
	}

	return anyMsg, nil
}

type StakingDelegation struct {
	// delegator_address is the encoded address of the delegator.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	// validator_address is the encoded address of the validator.
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// shares define the delegation shares received.
	Shares math.LegacyDec `protobuf:"bytes,3,opt,name=shares,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"shares"`
}

func (s *StakingDelegation) ToProto() (*anypb.Any, error) {
	// Create a map to hold the protobuf fields
	fields := map[string]interface{}{
		"delegator_address": s.DelegatorAddress,
		"validator_address": s.ValidatorAddress,
		"shares":            s.Shares.String(),
	}

	// Convert the map to a protobuf Struct
	pbStruct, err := structpb.NewStruct(fields)
	if err != nil {
		return nil, err
	}

	// Marshal the Struct into an Any message
	anyMsg, err := anypb.New(pbStruct)
	if err != nil {
		return nil, err
	}

	return anyMsg, nil
}

type StakingMsgCreateValidator struct {
	Commission        StakingValidatorCommission `protobuf:"bytes,2,opt,name=commission,proto3" json:"commission"`
	MinSelfDelegation math.Int                   `protobuf:"bytes,3,opt,name=min_self_delegation,json=minSelfDelegation,proto3,customtype=cosmossdk.io/math.Int" json:"min_self_delegation"`
	// The validator address bytes and delegator address bytes refer to the same account while creating validator (defer
	// only in bech32 notation).
	ValidatorAddress string       `protobuf:"bytes,5,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Pubkey           *gogoany.Any `protobuf:"bytes,6,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Value            sdk.Coin     `protobuf:"bytes,7,opt,name=value,proto3" json:"value"`
}

func (s *StakingMsgCreateValidator) ToProto() (*anypb.Any, error) {
	comm, err := s.Commission.ToProto()
	if err != nil {
		return nil, err
	}

	// Create a map to hold the protobuf fields
	fields := map[string]interface{}{
		"commission":          comm,
		"min_self_delegation": s.MinSelfDelegation.String(),
		"validator_address":   s.ValidatorAddress,
		"pubkey":              s.Pubkey,
		"value":               s.Value.String(),
	}

	// Convert the map to a protobuf Struct
	pbStruct, err := structpb.NewStruct(fields)
	if err != nil {
		return nil, err
	}

	// Marshal the Struct into an Any message
	anyMsg, err := anypb.New(pbStruct)
	if err != nil {
		return nil, err
	}

	return anyMsg, nil
}

type StakingValidatorCommission struct {
	// rate is the commission rate charged to delegators, as a fraction.
	Rate math.LegacyDec `protobuf:"bytes,1,opt,name=rate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"rate"`
	// max_rate defines the maximum commission rate which validator can ever charge, as a fraction.
	MaxRate math.LegacyDec `protobuf:"bytes,2,opt,name=max_rate,json=maxRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"max_rate"`
	// max_change_rate defines the maximum daily increase of the validator commission, as a fraction.
	MaxChangeRate math.LegacyDec `protobuf:"bytes,3,opt,name=max_change_rate,json=maxChangeRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"max_change_rate"`
}

func (s *StakingValidatorCommission) ToProto() (*anypb.Any, error) {
	// Create a map to hold the protobuf fields
	fields := map[string]interface{}{
		"rate":            s.Rate.String(),
		"max_rate":        s.MaxRate.String(),
		"max_change_rate": s.MaxChangeRate.String(),
	}

	// Convert the map to a protobuf Struct
	pbStruct, err := structpb.NewStruct(fields)
	if err != nil {
		return nil, err
	}

	// Marshal the Struct into an Any message
	anyMsg, err := anypb.New(pbStruct)
	if err != nil {
		return nil, err
	}

	return anyMsg, nil
}
