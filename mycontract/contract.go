// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mycontract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CoinPredictionStakingPendingReward is an auto generated low-level Go binding around an user-defined struct.
type CoinPredictionStakingPendingReward struct {
	StakeStatus   bool
	Unstakable    bool
	PendingReward *big.Int
}

// CoinPredictionStakingStake is an auto generated low-level Go binding around an user-defined struct.
type CoinPredictionStakingStake struct {
	Amount         *big.Int
	StartTime      *big.Int
	ClaimedRewards *big.Int
	Active         bool
	StakeApr       *big.Int
}

// MycontractMetaData contains all meta data concerning the Mycontract contract.
var MycontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialMinDurationYears\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialMaxDurationYears\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialApr\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newApr\",\"type\":\"uint256\"}],\"name\":\"AprUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"name\":\"LockToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxYears\",\"type\":\"uint256\"}],\"name\":\"MaxDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minYears\",\"type\":\"uint256\"}],\"name\":\"MinDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAprBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setApr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxDurationYears\",\"type\":\"uint256\"}],\"name\":\"setMaxDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinDurationYears\",\"type\":\"uint256\"}],\"name\":\"setMinDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"toggleLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnUnused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"withdrawUnusedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserPendingReward\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"stakeStatus\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"unstakable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingReward\",\"type\":\"uint256\"}],\"internalType\":\"structCoinPredictionStaking.PendingReward[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserStakes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakeApr\",\"type\":\"uint256\"}],\"internalType\":\"structCoinPredictionStaking.Stake[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserStakesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStakeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MycontractABI is the input ABI used to generate the binding from.
// Deprecated: Use MycontractMetaData.ABI instead.
var MycontractABI = MycontractMetaData.ABI

// Mycontract is an auto generated Go binding around an Ethereum contract.
type Mycontract struct {
	MycontractCaller     // Read-only binding to the contract
	MycontractTransactor // Write-only binding to the contract
	MycontractFilterer   // Log filterer for contract events
}

// MycontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MycontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MycontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MycontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MycontractSession struct {
	Contract     *Mycontract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MycontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MycontractCallerSession struct {
	Contract *MycontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MycontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MycontractTransactorSession struct {
	Contract     *MycontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MycontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MycontractRaw struct {
	Contract *Mycontract // Generic contract binding to access the raw methods on
}

// MycontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MycontractCallerRaw struct {
	Contract *MycontractCaller // Generic read-only contract binding to access the raw methods on
}

// MycontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MycontractTransactorRaw struct {
	Contract *MycontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMycontract creates a new instance of Mycontract, bound to a specific deployed contract.
func NewMycontract(address common.Address, backend bind.ContractBackend) (*Mycontract, error) {
	contract, err := bindMycontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mycontract{MycontractCaller: MycontractCaller{contract: contract}, MycontractTransactor: MycontractTransactor{contract: contract}, MycontractFilterer: MycontractFilterer{contract: contract}}, nil
}

// NewMycontractCaller creates a new read-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractCaller(address common.Address, caller bind.ContractCaller) (*MycontractCaller, error) {
	contract, err := bindMycontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractCaller{contract: contract}, nil
}

// NewMycontractTransactor creates a new write-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractTransactor(address common.Address, transactor bind.ContractTransactor) (*MycontractTransactor, error) {
	contract, err := bindMycontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractTransactor{contract: contract}, nil
}

// NewMycontractFilterer creates a new log filterer instance of Mycontract, bound to a specific deployed contract.
func NewMycontractFilterer(address common.Address, filterer bind.ContractFilterer) (*MycontractFilterer, error) {
	contract, err := bindMycontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MycontractFilterer{contract: contract}, nil
}

// bindMycontract binds a generic wrapper to an already deployed contract.
func bindMycontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MycontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.MycontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transact(opts, method, params...)
}

// YEAR is a free data retrieval call binding the contract method 0x83914540.
//
// Solidity: function YEAR() view returns(uint256)
func (_Mycontract *MycontractCaller) YEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YEAR is a free data retrieval call binding the contract method 0x83914540.
//
// Solidity: function YEAR() view returns(uint256)
func (_Mycontract *MycontractSession) YEAR() (*big.Int, error) {
	return _Mycontract.Contract.YEAR(&_Mycontract.CallOpts)
}

// YEAR is a free data retrieval call binding the contract method 0x83914540.
//
// Solidity: function YEAR() view returns(uint256)
func (_Mycontract *MycontractCallerSession) YEAR() (*big.Int, error) {
	return _Mycontract.Contract.YEAR(&_Mycontract.CallOpts)
}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_Mycontract *MycontractCaller) Apr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "apr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_Mycontract *MycontractSession) Apr() (*big.Int, error) {
	return _Mycontract.Contract.Apr(&_Mycontract.CallOpts)
}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_Mycontract *MycontractCallerSession) Apr() (*big.Int, error) {
	return _Mycontract.Contract.Apr(&_Mycontract.CallOpts)
}

// GetUserPendingReward is a free data retrieval call binding the contract method 0xc6710629.
//
// Solidity: function getUserPendingReward(address user) view returns((bool,bool,uint256)[])
func (_Mycontract *MycontractCaller) GetUserPendingReward(opts *bind.CallOpts, user common.Address) ([]CoinPredictionStakingPendingReward, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getUserPendingReward", user)

	if err != nil {
		return *new([]CoinPredictionStakingPendingReward), err
	}

	out0 := *abi.ConvertType(out[0], new([]CoinPredictionStakingPendingReward)).(*[]CoinPredictionStakingPendingReward)

	return out0, err

}

// GetUserPendingReward is a free data retrieval call binding the contract method 0xc6710629.
//
// Solidity: function getUserPendingReward(address user) view returns((bool,bool,uint256)[])
func (_Mycontract *MycontractSession) GetUserPendingReward(user common.Address) ([]CoinPredictionStakingPendingReward, error) {
	return _Mycontract.Contract.GetUserPendingReward(&_Mycontract.CallOpts, user)
}

// GetUserPendingReward is a free data retrieval call binding the contract method 0xc6710629.
//
// Solidity: function getUserPendingReward(address user) view returns((bool,bool,uint256)[])
func (_Mycontract *MycontractCallerSession) GetUserPendingReward(user common.Address) ([]CoinPredictionStakingPendingReward, error) {
	return _Mycontract.Contract.GetUserPendingReward(&_Mycontract.CallOpts, user)
}

// GetUserStakes is a free data retrieval call binding the contract method 0x842e2981.
//
// Solidity: function getUserStakes(address user) view returns((uint256,uint256,uint256,bool,uint256)[])
func (_Mycontract *MycontractCaller) GetUserStakes(opts *bind.CallOpts, user common.Address) ([]CoinPredictionStakingStake, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getUserStakes", user)

	if err != nil {
		return *new([]CoinPredictionStakingStake), err
	}

	out0 := *abi.ConvertType(out[0], new([]CoinPredictionStakingStake)).(*[]CoinPredictionStakingStake)

	return out0, err

}

// GetUserStakes is a free data retrieval call binding the contract method 0x842e2981.
//
// Solidity: function getUserStakes(address user) view returns((uint256,uint256,uint256,bool,uint256)[])
func (_Mycontract *MycontractSession) GetUserStakes(user common.Address) ([]CoinPredictionStakingStake, error) {
	return _Mycontract.Contract.GetUserStakes(&_Mycontract.CallOpts, user)
}

// GetUserStakes is a free data retrieval call binding the contract method 0x842e2981.
//
// Solidity: function getUserStakes(address user) view returns((uint256,uint256,uint256,bool,uint256)[])
func (_Mycontract *MycontractCallerSession) GetUserStakes(user common.Address) ([]CoinPredictionStakingStake, error) {
	return _Mycontract.Contract.GetUserStakes(&_Mycontract.CallOpts, user)
}

// GetUserStakesCount is a free data retrieval call binding the contract method 0x98dc8dea.
//
// Solidity: function getUserStakesCount(address user) view returns(uint256)
func (_Mycontract *MycontractCaller) GetUserStakesCount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getUserStakesCount", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserStakesCount is a free data retrieval call binding the contract method 0x98dc8dea.
//
// Solidity: function getUserStakesCount(address user) view returns(uint256)
func (_Mycontract *MycontractSession) GetUserStakesCount(user common.Address) (*big.Int, error) {
	return _Mycontract.Contract.GetUserStakesCount(&_Mycontract.CallOpts, user)
}

// GetUserStakesCount is a free data retrieval call binding the contract method 0x98dc8dea.
//
// Solidity: function getUserStakesCount(address user) view returns(uint256)
func (_Mycontract *MycontractCallerSession) GetUserStakesCount(user common.Address) (*big.Int, error) {
	return _Mycontract.Contract.GetUserStakesCount(&_Mycontract.CallOpts, user)
}

// MaxStakeDuration is a free data retrieval call binding the contract method 0x76f70003.
//
// Solidity: function maxStakeDuration() view returns(uint256)
func (_Mycontract *MycontractCaller) MaxStakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "maxStakeDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakeDuration is a free data retrieval call binding the contract method 0x76f70003.
//
// Solidity: function maxStakeDuration() view returns(uint256)
func (_Mycontract *MycontractSession) MaxStakeDuration() (*big.Int, error) {
	return _Mycontract.Contract.MaxStakeDuration(&_Mycontract.CallOpts)
}

// MaxStakeDuration is a free data retrieval call binding the contract method 0x76f70003.
//
// Solidity: function maxStakeDuration() view returns(uint256)
func (_Mycontract *MycontractCallerSession) MaxStakeDuration() (*big.Int, error) {
	return _Mycontract.Contract.MaxStakeDuration(&_Mycontract.CallOpts)
}

// MinStakeDuration is a free data retrieval call binding the contract method 0x5fec5c64.
//
// Solidity: function minStakeDuration() view returns(uint256)
func (_Mycontract *MycontractCaller) MinStakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "minStakeDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeDuration is a free data retrieval call binding the contract method 0x5fec5c64.
//
// Solidity: function minStakeDuration() view returns(uint256)
func (_Mycontract *MycontractSession) MinStakeDuration() (*big.Int, error) {
	return _Mycontract.Contract.MinStakeDuration(&_Mycontract.CallOpts)
}

// MinStakeDuration is a free data retrieval call binding the contract method 0x5fec5c64.
//
// Solidity: function minStakeDuration() view returns(uint256)
func (_Mycontract *MycontractCallerSession) MinStakeDuration() (*big.Int, error) {
	return _Mycontract.Contract.MinStakeDuration(&_Mycontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractSession) Owner() (common.Address, error) {
	return _Mycontract.Contract.Owner(&_Mycontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractCallerSession) Owner() (common.Address, error) {
	return _Mycontract.Contract.Owner(&_Mycontract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Mycontract *MycontractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Mycontract *MycontractSession) Token() (common.Address, error) {
	return _Mycontract.Contract.Token(&_Mycontract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Mycontract *MycontractCallerSession) Token() (common.Address, error) {
	return _Mycontract.Contract.Token(&_Mycontract.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Mycontract *MycontractCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Mycontract *MycontractSession) TotalStaked() (*big.Int, error) {
	return _Mycontract.Contract.TotalStaked(&_Mycontract.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Mycontract *MycontractCallerSession) TotalStaked() (*big.Int, error) {
	return _Mycontract.Contract.TotalStaked(&_Mycontract.CallOpts)
}

// Unlocked is a free data retrieval call binding the contract method 0x6a5e2650.
//
// Solidity: function unlocked() view returns(bool)
func (_Mycontract *MycontractCaller) Unlocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "unlocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Unlocked is a free data retrieval call binding the contract method 0x6a5e2650.
//
// Solidity: function unlocked() view returns(bool)
func (_Mycontract *MycontractSession) Unlocked() (bool, error) {
	return _Mycontract.Contract.Unlocked(&_Mycontract.CallOpts)
}

// Unlocked is a free data retrieval call binding the contract method 0x6a5e2650.
//
// Solidity: function unlocked() view returns(bool)
func (_Mycontract *MycontractCallerSession) Unlocked() (bool, error) {
	return _Mycontract.Contract.Unlocked(&_Mycontract.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 index) returns()
func (_Mycontract *MycontractTransactor) ClaimReward(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "claimReward", index)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 index) returns()
func (_Mycontract *MycontractSession) ClaimReward(index *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.ClaimReward(&_Mycontract.TransactOpts, index)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 index) returns()
func (_Mycontract *MycontractTransactorSession) ClaimReward(index *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.ClaimReward(&_Mycontract.TransactOpts, index)
}

// SetApr is a paid mutator transaction binding the contract method 0x59cee29c.
//
// Solidity: function setApr(uint256 newAprBasisPoints) returns()
func (_Mycontract *MycontractTransactor) SetApr(opts *bind.TransactOpts, newAprBasisPoints *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setApr", newAprBasisPoints)
}

// SetApr is a paid mutator transaction binding the contract method 0x59cee29c.
//
// Solidity: function setApr(uint256 newAprBasisPoints) returns()
func (_Mycontract *MycontractSession) SetApr(newAprBasisPoints *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SetApr(&_Mycontract.TransactOpts, newAprBasisPoints)
}

// SetApr is a paid mutator transaction binding the contract method 0x59cee29c.
//
// Solidity: function setApr(uint256 newAprBasisPoints) returns()
func (_Mycontract *MycontractTransactorSession) SetApr(newAprBasisPoints *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SetApr(&_Mycontract.TransactOpts, newAprBasisPoints)
}

// SetMaxDuration is a paid mutator transaction binding the contract method 0xcf0f34c4.
//
// Solidity: function setMaxDuration(uint256 newMaxDurationYears) returns()
func (_Mycontract *MycontractTransactor) SetMaxDuration(opts *bind.TransactOpts, newMaxDurationYears *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setMaxDuration", newMaxDurationYears)
}

// SetMaxDuration is a paid mutator transaction binding the contract method 0xcf0f34c4.
//
// Solidity: function setMaxDuration(uint256 newMaxDurationYears) returns()
func (_Mycontract *MycontractSession) SetMaxDuration(newMaxDurationYears *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SetMaxDuration(&_Mycontract.TransactOpts, newMaxDurationYears)
}

// SetMaxDuration is a paid mutator transaction binding the contract method 0xcf0f34c4.
//
// Solidity: function setMaxDuration(uint256 newMaxDurationYears) returns()
func (_Mycontract *MycontractTransactorSession) SetMaxDuration(newMaxDurationYears *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SetMaxDuration(&_Mycontract.TransactOpts, newMaxDurationYears)
}

// SetMinDuration is a paid mutator transaction binding the contract method 0x1674bade.
//
// Solidity: function setMinDuration(uint256 newMinDurationYears) returns()
func (_Mycontract *MycontractTransactor) SetMinDuration(opts *bind.TransactOpts, newMinDurationYears *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setMinDuration", newMinDurationYears)
}

// SetMinDuration is a paid mutator transaction binding the contract method 0x1674bade.
//
// Solidity: function setMinDuration(uint256 newMinDurationYears) returns()
func (_Mycontract *MycontractSession) SetMinDuration(newMinDurationYears *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SetMinDuration(&_Mycontract.TransactOpts, newMinDurationYears)
}

// SetMinDuration is a paid mutator transaction binding the contract method 0x1674bade.
//
// Solidity: function setMinDuration(uint256 newMinDurationYears) returns()
func (_Mycontract *MycontractTransactorSession) SetMinDuration(newMinDurationYears *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SetMinDuration(&_Mycontract.TransactOpts, newMinDurationYears)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Mycontract *MycontractTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Mycontract *MycontractSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Stake(&_Mycontract.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Mycontract *MycontractTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Stake(&_Mycontract.TransactOpts, amount)
}

// ToggleLock is a paid mutator transaction binding the contract method 0xff9413d8.
//
// Solidity: function toggleLock() returns()
func (_Mycontract *MycontractTransactor) ToggleLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "toggleLock")
}

// ToggleLock is a paid mutator transaction binding the contract method 0xff9413d8.
//
// Solidity: function toggleLock() returns()
func (_Mycontract *MycontractSession) ToggleLock() (*types.Transaction, error) {
	return _Mycontract.Contract.ToggleLock(&_Mycontract.TransactOpts)
}

// ToggleLock is a paid mutator transaction binding the contract method 0xff9413d8.
//
// Solidity: function toggleLock() returns()
func (_Mycontract *MycontractTransactorSession) ToggleLock() (*types.Transaction, error) {
	return _Mycontract.Contract.ToggleLock(&_Mycontract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mycontract *MycontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mycontract *MycontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.TransferOwnership(&_Mycontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mycontract *MycontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.TransferOwnership(&_Mycontract.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 index) returns()
func (_Mycontract *MycontractTransactor) Unstake(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "unstake", index)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 index) returns()
func (_Mycontract *MycontractSession) Unstake(index *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Unstake(&_Mycontract.TransactOpts, index)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 index) returns()
func (_Mycontract *MycontractTransactorSession) Unstake(index *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Unstake(&_Mycontract.TransactOpts, index)
}

// WithdrawUnusedTokens is a paid mutator transaction binding the contract method 0x75cd81b7.
//
// Solidity: function withdrawUnusedTokens() returns()
func (_Mycontract *MycontractTransactor) WithdrawUnusedTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "withdrawUnusedTokens")
}

// WithdrawUnusedTokens is a paid mutator transaction binding the contract method 0x75cd81b7.
//
// Solidity: function withdrawUnusedTokens() returns()
func (_Mycontract *MycontractSession) WithdrawUnusedTokens() (*types.Transaction, error) {
	return _Mycontract.Contract.WithdrawUnusedTokens(&_Mycontract.TransactOpts)
}

// WithdrawUnusedTokens is a paid mutator transaction binding the contract method 0x75cd81b7.
//
// Solidity: function withdrawUnusedTokens() returns()
func (_Mycontract *MycontractTransactorSession) WithdrawUnusedTokens() (*types.Transaction, error) {
	return _Mycontract.Contract.WithdrawUnusedTokens(&_Mycontract.TransactOpts)
}

// MycontractAprUpdatedIterator is returned from FilterAprUpdated and is used to iterate over the raw logs and unpacked data for AprUpdated events raised by the Mycontract contract.
type MycontractAprUpdatedIterator struct {
	Event *MycontractAprUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycontractAprUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractAprUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycontractAprUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycontractAprUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractAprUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractAprUpdated represents a AprUpdated event raised by the Mycontract contract.
type MycontractAprUpdated struct {
	NewApr *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAprUpdated is a free log retrieval operation binding the contract event 0x16bc525116a7ec45cf36d84a97ab6b444a8c5264cc1a9468bfae786131561115.
//
// Solidity: event AprUpdated(uint256 newApr)
func (_Mycontract *MycontractFilterer) FilterAprUpdated(opts *bind.FilterOpts) (*MycontractAprUpdatedIterator, error) {

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "AprUpdated")
	if err != nil {
		return nil, err
	}
	return &MycontractAprUpdatedIterator{contract: _Mycontract.contract, event: "AprUpdated", logs: logs, sub: sub}, nil
}

// WatchAprUpdated is a free log subscription operation binding the contract event 0x16bc525116a7ec45cf36d84a97ab6b444a8c5264cc1a9468bfae786131561115.
//
// Solidity: event AprUpdated(uint256 newApr)
func (_Mycontract *MycontractFilterer) WatchAprUpdated(opts *bind.WatchOpts, sink chan<- *MycontractAprUpdated) (event.Subscription, error) {

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "AprUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractAprUpdated)
				if err := _Mycontract.contract.UnpackLog(event, "AprUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAprUpdated is a log parse operation binding the contract event 0x16bc525116a7ec45cf36d84a97ab6b444a8c5264cc1a9468bfae786131561115.
//
// Solidity: event AprUpdated(uint256 newApr)
func (_Mycontract *MycontractFilterer) ParseAprUpdated(log types.Log) (*MycontractAprUpdated, error) {
	event := new(MycontractAprUpdated)
	if err := _Mycontract.contract.UnpackLog(event, "AprUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractLockToggledIterator is returned from FilterLockToggled and is used to iterate over the raw logs and unpacked data for LockToggled events raised by the Mycontract contract.
type MycontractLockToggledIterator struct {
	Event *MycontractLockToggled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycontractLockToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractLockToggled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycontractLockToggled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycontractLockToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractLockToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractLockToggled represents a LockToggled event raised by the Mycontract contract.
type MycontractLockToggled struct {
	Unlocked bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLockToggled is a free log retrieval operation binding the contract event 0x9011bbd43ef09c72953d907f4f5fc45830479ca653446fc9380711238b889276.
//
// Solidity: event LockToggled(bool unlocked)
func (_Mycontract *MycontractFilterer) FilterLockToggled(opts *bind.FilterOpts) (*MycontractLockToggledIterator, error) {

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "LockToggled")
	if err != nil {
		return nil, err
	}
	return &MycontractLockToggledIterator{contract: _Mycontract.contract, event: "LockToggled", logs: logs, sub: sub}, nil
}

// WatchLockToggled is a free log subscription operation binding the contract event 0x9011bbd43ef09c72953d907f4f5fc45830479ca653446fc9380711238b889276.
//
// Solidity: event LockToggled(bool unlocked)
func (_Mycontract *MycontractFilterer) WatchLockToggled(opts *bind.WatchOpts, sink chan<- *MycontractLockToggled) (event.Subscription, error) {

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "LockToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractLockToggled)
				if err := _Mycontract.contract.UnpackLog(event, "LockToggled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockToggled is a log parse operation binding the contract event 0x9011bbd43ef09c72953d907f4f5fc45830479ca653446fc9380711238b889276.
//
// Solidity: event LockToggled(bool unlocked)
func (_Mycontract *MycontractFilterer) ParseLockToggled(log types.Log) (*MycontractLockToggled, error) {
	event := new(MycontractLockToggled)
	if err := _Mycontract.contract.UnpackLog(event, "LockToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractMaxDurationUpdatedIterator is returned from FilterMaxDurationUpdated and is used to iterate over the raw logs and unpacked data for MaxDurationUpdated events raised by the Mycontract contract.
type MycontractMaxDurationUpdatedIterator struct {
	Event *MycontractMaxDurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycontractMaxDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractMaxDurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycontractMaxDurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycontractMaxDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractMaxDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractMaxDurationUpdated represents a MaxDurationUpdated event raised by the Mycontract contract.
type MycontractMaxDurationUpdated struct {
	MaxYears *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMaxDurationUpdated is a free log retrieval operation binding the contract event 0x5051341ab3f6d56c8295227b7476589d73ae29d5dfaff08dfc01f67f28a8a3e3.
//
// Solidity: event MaxDurationUpdated(uint256 maxYears)
func (_Mycontract *MycontractFilterer) FilterMaxDurationUpdated(opts *bind.FilterOpts) (*MycontractMaxDurationUpdatedIterator, error) {

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "MaxDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &MycontractMaxDurationUpdatedIterator{contract: _Mycontract.contract, event: "MaxDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxDurationUpdated is a free log subscription operation binding the contract event 0x5051341ab3f6d56c8295227b7476589d73ae29d5dfaff08dfc01f67f28a8a3e3.
//
// Solidity: event MaxDurationUpdated(uint256 maxYears)
func (_Mycontract *MycontractFilterer) WatchMaxDurationUpdated(opts *bind.WatchOpts, sink chan<- *MycontractMaxDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "MaxDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractMaxDurationUpdated)
				if err := _Mycontract.contract.UnpackLog(event, "MaxDurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxDurationUpdated is a log parse operation binding the contract event 0x5051341ab3f6d56c8295227b7476589d73ae29d5dfaff08dfc01f67f28a8a3e3.
//
// Solidity: event MaxDurationUpdated(uint256 maxYears)
func (_Mycontract *MycontractFilterer) ParseMaxDurationUpdated(log types.Log) (*MycontractMaxDurationUpdated, error) {
	event := new(MycontractMaxDurationUpdated)
	if err := _Mycontract.contract.UnpackLog(event, "MaxDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractMinDurationUpdatedIterator is returned from FilterMinDurationUpdated and is used to iterate over the raw logs and unpacked data for MinDurationUpdated events raised by the Mycontract contract.
type MycontractMinDurationUpdatedIterator struct {
	Event *MycontractMinDurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycontractMinDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractMinDurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycontractMinDurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycontractMinDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractMinDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractMinDurationUpdated represents a MinDurationUpdated event raised by the Mycontract contract.
type MycontractMinDurationUpdated struct {
	MinYears *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinDurationUpdated is a free log retrieval operation binding the contract event 0x049a6934bf12da797a1e19e9f70b578af0fdab9d72c2d4b8d82984a54c9e6400.
//
// Solidity: event MinDurationUpdated(uint256 minYears)
func (_Mycontract *MycontractFilterer) FilterMinDurationUpdated(opts *bind.FilterOpts) (*MycontractMinDurationUpdatedIterator, error) {

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "MinDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &MycontractMinDurationUpdatedIterator{contract: _Mycontract.contract, event: "MinDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchMinDurationUpdated is a free log subscription operation binding the contract event 0x049a6934bf12da797a1e19e9f70b578af0fdab9d72c2d4b8d82984a54c9e6400.
//
// Solidity: event MinDurationUpdated(uint256 minYears)
func (_Mycontract *MycontractFilterer) WatchMinDurationUpdated(opts *bind.WatchOpts, sink chan<- *MycontractMinDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "MinDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractMinDurationUpdated)
				if err := _Mycontract.contract.UnpackLog(event, "MinDurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinDurationUpdated is a log parse operation binding the contract event 0x049a6934bf12da797a1e19e9f70b578af0fdab9d72c2d4b8d82984a54c9e6400.
//
// Solidity: event MinDurationUpdated(uint256 minYears)
func (_Mycontract *MycontractFilterer) ParseMinDurationUpdated(log types.Log) (*MycontractMinDurationUpdated, error) {
	event := new(MycontractMinDurationUpdated)
	if err := _Mycontract.contract.UnpackLog(event, "MinDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mycontract contract.
type MycontractOwnershipTransferredIterator struct {
	Event *MycontractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycontractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractOwnershipTransferred represents a OwnershipTransferred event raised by the Mycontract contract.
type MycontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mycontract *MycontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MycontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MycontractOwnershipTransferredIterator{contract: _Mycontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mycontract *MycontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MycontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractOwnershipTransferred)
				if err := _Mycontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mycontract *MycontractFilterer) ParseOwnershipTransferred(log types.Log) (*MycontractOwnershipTransferred, error) {
	event := new(MycontractOwnershipTransferred)
	if err := _Mycontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Mycontract contract.
type MycontractRewardClaimedIterator struct {
	Event *MycontractRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycontractRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractRewardClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycontractRewardClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycontractRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractRewardClaimed represents a RewardClaimed event raised by the Mycontract contract.
type MycontractRewardClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 amount)
func (_Mycontract *MycontractFilterer) FilterRewardClaimed(opts *bind.FilterOpts, user []common.Address) (*MycontractRewardClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "RewardClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &MycontractRewardClaimedIterator{contract: _Mycontract.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 amount)
func (_Mycontract *MycontractFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *MycontractRewardClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "RewardClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractRewardClaimed)
				if err := _Mycontract.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaimed is a log parse operation binding the contract event 0x106f923f993c2149d49b4255ff723acafa1f2d94393f561d3eda32ae348f7241.
//
// Solidity: event RewardClaimed(address indexed user, uint256 amount)
func (_Mycontract *MycontractFilterer) ParseRewardClaimed(log types.Log) (*MycontractRewardClaimed, error) {
	event := new(MycontractRewardClaimed)
	if err := _Mycontract.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Mycontract contract.
type MycontractStakedIterator struct {
	Event *MycontractStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycontractStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycontractStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycontractStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractStaked represents a Staked event raised by the Mycontract contract.
type MycontractStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Mycontract *MycontractFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*MycontractStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &MycontractStakedIterator{contract: _Mycontract.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Mycontract *MycontractFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *MycontractStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractStaked)
				if err := _Mycontract.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Mycontract *MycontractFilterer) ParseStaked(log types.Log) (*MycontractStaked, error) {
	event := new(MycontractStaked)
	if err := _Mycontract.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Mycontract contract.
type MycontractUnstakedIterator struct {
	Event *MycontractUnstaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycontractUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractUnstaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycontractUnstaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycontractUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractUnstaked represents a Unstaked event raised by the Mycontract contract.
type MycontractUnstaked struct {
	User    common.Address
	Amount  *big.Int
	Rewards *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 rewards)
func (_Mycontract *MycontractFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*MycontractUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &MycontractUnstakedIterator{contract: _Mycontract.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 rewards)
func (_Mycontract *MycontractFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *MycontractUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractUnstaked)
				if err := _Mycontract.contract.UnpackLog(event, "Unstaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstaked is a log parse operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 rewards)
func (_Mycontract *MycontractFilterer) ParseUnstaked(log types.Log) (*MycontractUnstaked, error) {
	event := new(MycontractUnstaked)
	if err := _Mycontract.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractWithdrawnUnusedIterator is returned from FilterWithdrawnUnused and is used to iterate over the raw logs and unpacked data for WithdrawnUnused events raised by the Mycontract contract.
type MycontractWithdrawnUnusedIterator struct {
	Event *MycontractWithdrawnUnused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycontractWithdrawnUnusedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractWithdrawnUnused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycontractWithdrawnUnused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycontractWithdrawnUnusedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractWithdrawnUnusedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractWithdrawnUnused represents a WithdrawnUnused event raised by the Mycontract contract.
type MycontractWithdrawnUnused struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnUnused is a free log retrieval operation binding the contract event 0xfb3b624a488261ddeb4fb912d629fefafb3b725dcdd408b0a2b0e719d64cc510.
//
// Solidity: event WithdrawnUnused(address indexed owner, uint256 amount)
func (_Mycontract *MycontractFilterer) FilterWithdrawnUnused(opts *bind.FilterOpts, owner []common.Address) (*MycontractWithdrawnUnusedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "WithdrawnUnused", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MycontractWithdrawnUnusedIterator{contract: _Mycontract.contract, event: "WithdrawnUnused", logs: logs, sub: sub}, nil
}

// WatchWithdrawnUnused is a free log subscription operation binding the contract event 0xfb3b624a488261ddeb4fb912d629fefafb3b725dcdd408b0a2b0e719d64cc510.
//
// Solidity: event WithdrawnUnused(address indexed owner, uint256 amount)
func (_Mycontract *MycontractFilterer) WatchWithdrawnUnused(opts *bind.WatchOpts, sink chan<- *MycontractWithdrawnUnused, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "WithdrawnUnused", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractWithdrawnUnused)
				if err := _Mycontract.contract.UnpackLog(event, "WithdrawnUnused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnUnused is a log parse operation binding the contract event 0xfb3b624a488261ddeb4fb912d629fefafb3b725dcdd408b0a2b0e719d64cc510.
//
// Solidity: event WithdrawnUnused(address indexed owner, uint256 amount)
func (_Mycontract *MycontractFilterer) ParseWithdrawnUnused(log types.Log) (*MycontractWithdrawnUnused, error) {
	event := new(MycontractWithdrawnUnused)
	if err := _Mycontract.contract.UnpackLog(event, "WithdrawnUnused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
