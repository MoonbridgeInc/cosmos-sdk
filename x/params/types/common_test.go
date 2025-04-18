package types_test

import (
	"errors"
	"fmt"
	"time"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	keyUnbondingTime          = []byte("UnbondingTime")
	keyMaxValidators          = []byte("MaxValidators")
	keyBondDenom              = []byte("BondDenom")
	keyMaxRedelegationEntries = []byte("MaxRedelegationEntries")

	key  = storetypes.NewKVStoreKey("storekey")
	tkey = storetypes.NewTransientStoreKey("transientstorekey")
)

type params struct {
	UnbondingTime time.Duration `json:"unbonding_time" yaml:"unbonding_time"`
	MaxValidators uint16        `json:"max_validators" yaml:"max_validators"`
	BondDenom     string        `json:"bond_denom" yaml:"bond_denom"`
}

type paramsV2 struct {
	UnbondingTime          time.Duration `json:"unbonding_time" yaml:"unbonding_time"`
	MaxValidators          uint16        `json:"max_validators" yaml:"max_validators"`
	BondDenom              string        `json:"bond_denom" yaml:"bond_denom"`
	MaxRedelegationEntries uint32        `json:"max_redelegation_entries" yaml:"max_redelegation_entries"`
}

func validateUnbondingTime(i any) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < (24 * time.Hour) {
		return fmt.Errorf("unbonding time must be at least one day")
	}

	return nil
}

func validateMaxValidators(i any) error {
	_, ok := i.(uint16)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateBondDenom(i any) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return errors.New("denom cannot be empty")
	}

	return nil
}

func validateMaxRedelegationEntries(i any) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func (p *params) ParamSetPairs() types.ParamSetPairs {
	return types.ParamSetPairs{
		{keyUnbondingTime, &p.UnbondingTime, validateUnbondingTime},
		{keyMaxValidators, &p.MaxValidators, validateMaxValidators},
		{keyBondDenom, &p.BondDenom, validateBondDenom},
	}
}

func (p *paramsV2) ParamSetPairs() types.ParamSetPairs {
	return types.ParamSetPairs{
		{keyUnbondingTime, &p.UnbondingTime, validateUnbondingTime},
		{keyMaxValidators, &p.MaxValidators, validateMaxValidators},
		{keyBondDenom, &p.BondDenom, validateBondDenom},
		{keyMaxRedelegationEntries, &p.MaxRedelegationEntries, validateMaxRedelegationEntries},
	}
}

func paramKeyTable() types.KeyTable {
	return types.NewKeyTable().RegisterParamSet(&params{})
}
