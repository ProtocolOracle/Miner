// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// PowERC20MetaData contains all meta data concerning the PowERC20 contract.
var PowERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tick\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_amt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"etherunes_protocol_Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"ethrunes_protocol_Inscribe\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"mine\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"inscription\",\"type\":\"string\"}],\"name\":\"setInscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_inscription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minedNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tick\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PowERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use PowERC20MetaData.ABI instead.
var PowERC20ABI = PowERC20MetaData.ABI

// PowERC20 is an auto generated Go binding around an Ethereum contract.
type PowERC20 struct {
	PowERC20Caller     // Read-only binding to the contract
	PowERC20Transactor // Write-only binding to the contract
	PowERC20Filterer   // Log filterer for contract events
}

// PowERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type PowERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PowERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PowERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PowERC20Session struct {
	Contract     *PowERC20         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PowERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PowERC20CallerSession struct {
	Contract *PowERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PowERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PowERC20TransactorSession struct {
	Contract     *PowERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PowERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type PowERC20Raw struct {
	Contract *PowERC20 // Generic contract binding to access the raw methods on
}

// PowERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PowERC20CallerRaw struct {
	Contract *PowERC20Caller // Generic read-only contract binding to access the raw methods on
}

// PowERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PowERC20TransactorRaw struct {
	Contract *PowERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPowERC20 creates a new instance of PowERC20, bound to a specific deployed contract.
func NewPowERC20(address common.Address, backend bind.ContractBackend) (*PowERC20, error) {
	contract, err := bindPowERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PowERC20{PowERC20Caller: PowERC20Caller{contract: contract}, PowERC20Transactor: PowERC20Transactor{contract: contract}, PowERC20Filterer: PowERC20Filterer{contract: contract}}, nil
}

// NewPowERC20Caller creates a new read-only instance of PowERC20, bound to a specific deployed contract.
func NewPowERC20Caller(address common.Address, caller bind.ContractCaller) (*PowERC20Caller, error) {
	contract, err := bindPowERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PowERC20Caller{contract: contract}, nil
}

// NewPowERC20Transactor creates a new write-only instance of PowERC20, bound to a specific deployed contract.
func NewPowERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*PowERC20Transactor, error) {
	contract, err := bindPowERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PowERC20Transactor{contract: contract}, nil
}

// NewPowERC20Filterer creates a new log filterer instance of PowERC20, bound to a specific deployed contract.
func NewPowERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*PowERC20Filterer, error) {
	contract, err := bindPowERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PowERC20Filterer{contract: contract}, nil
}

// bindPowERC20 binds a generic wrapper to an already deployed contract.
func bindPowERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PowERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PowERC20 *PowERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PowERC20.Contract.PowERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PowERC20 *PowERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowERC20.Contract.PowERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PowERC20 *PowERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PowERC20.Contract.PowERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PowERC20 *PowERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PowERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PowERC20 *PowERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PowERC20 *PowERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PowERC20.Contract.contract.Transact(opts, method, params...)
}

// FEE is a free data retrieval call binding the contract method 0xc57981b5.
//
// Solidity: function FEE() view returns(uint256)
func (_PowERC20 *PowERC20Caller) FEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEE is a free data retrieval call binding the contract method 0xc57981b5.
//
// Solidity: function FEE() view returns(uint256)
func (_PowERC20 *PowERC20Session) FEE() (*big.Int, error) {
	return _PowERC20.Contract.FEE(&_PowERC20.CallOpts)
}

// FEE is a free data retrieval call binding the contract method 0xc57981b5.
//
// Solidity: function FEE() view returns(uint256)
func (_PowERC20 *PowERC20CallerSession) FEE() (*big.Int, error) {
	return _PowERC20.Contract.FEE(&_PowERC20.CallOpts)
}

// MAXCOUNT is a free data retrieval call binding the contract method 0x77163c1d.
//
// Solidity: function MAX_COUNT() view returns(uint256)
func (_PowERC20 *PowERC20Caller) MAXCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "MAX_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCOUNT is a free data retrieval call binding the contract method 0x77163c1d.
//
// Solidity: function MAX_COUNT() view returns(uint256)
func (_PowERC20 *PowERC20Session) MAXCOUNT() (*big.Int, error) {
	return _PowERC20.Contract.MAXCOUNT(&_PowERC20.CallOpts)
}

// MAXCOUNT is a free data retrieval call binding the contract method 0x77163c1d.
//
// Solidity: function MAX_COUNT() view returns(uint256)
func (_PowERC20 *PowERC20CallerSession) MAXCOUNT() (*big.Int, error) {
	return _PowERC20.Contract.MAXCOUNT(&_PowERC20.CallOpts)
}

// Inscription is a free data retrieval call binding the contract method 0x0dd2e7df.
//
// Solidity: function _inscription() view returns(string)
func (_PowERC20 *PowERC20Caller) Inscription(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "_inscription")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Inscription is a free data retrieval call binding the contract method 0x0dd2e7df.
//
// Solidity: function _inscription() view returns(string)
func (_PowERC20 *PowERC20Session) Inscription() (string, error) {
	return _PowERC20.Contract.Inscription(&_PowERC20.CallOpts)
}

// Inscription is a free data retrieval call binding the contract method 0x0dd2e7df.
//
// Solidity: function _inscription() view returns(string)
func (_PowERC20 *PowERC20CallerSession) Inscription() (string, error) {
	return _PowERC20.Contract.Inscription(&_PowERC20.CallOpts)
}

// Amt is a free data retrieval call binding the contract method 0x16eab96b.
//
// Solidity: function amt() view returns(string)
func (_PowERC20 *PowERC20Caller) Amt(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "amt")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Amt is a free data retrieval call binding the contract method 0x16eab96b.
//
// Solidity: function amt() view returns(string)
func (_PowERC20 *PowERC20Session) Amt() (string, error) {
	return _PowERC20.Contract.Amt(&_PowERC20.CallOpts)
}

// Amt is a free data retrieval call binding the contract method 0x16eab96b.
//
// Solidity: function amt() view returns(string)
func (_PowERC20 *PowERC20CallerSession) Amt() (string, error) {
	return _PowERC20.Contract.Amt(&_PowERC20.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_PowERC20 *PowERC20Caller) Balance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "balance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_PowERC20 *PowERC20Session) Balance(arg0 common.Address) (*big.Int, error) {
	return _PowERC20.Contract.Balance(&_PowERC20.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_PowERC20 *PowERC20CallerSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _PowERC20.Contract.Balance(&_PowERC20.CallOpts, arg0)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_PowERC20 *PowERC20Caller) Challenge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "challenge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_PowERC20 *PowERC20Session) Challenge() (*big.Int, error) {
	return _PowERC20.Contract.Challenge(&_PowERC20.CallOpts)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_PowERC20 *PowERC20CallerSession) Challenge() (*big.Int, error) {
	return _PowERC20.Contract.Challenge(&_PowERC20.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_PowERC20 *PowERC20Caller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_PowERC20 *PowERC20Session) Counter() (*big.Int, error) {
	return _PowERC20.Contract.Counter(&_PowERC20.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_PowERC20 *PowERC20CallerSession) Counter() (*big.Int, error) {
	return _PowERC20.Contract.Counter(&_PowERC20.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_PowERC20 *PowERC20Caller) Difficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_PowERC20 *PowERC20Session) Difficulty() (*big.Int, error) {
	return _PowERC20.Contract.Difficulty(&_PowERC20.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_PowERC20 *PowERC20CallerSession) Difficulty() (*big.Int, error) {
	return _PowERC20.Contract.Difficulty(&_PowERC20.CallOpts)
}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_PowERC20 *PowERC20Caller) MinedNonces(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "minedNonces", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_PowERC20 *PowERC20Session) MinedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _PowERC20.Contract.MinedNonces(&_PowERC20.CallOpts, arg0, arg1)
}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_PowERC20 *PowERC20CallerSession) MinedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _PowERC20.Contract.MinedNonces(&_PowERC20.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PowERC20 *PowERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PowERC20 *PowERC20Session) Owner() (common.Address, error) {
	return _PowERC20.Contract.Owner(&_PowERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PowERC20 *PowERC20CallerSession) Owner() (common.Address, error) {
	return _PowERC20.Contract.Owner(&_PowERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PowERC20 *PowERC20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PowERC20 *PowERC20Session) Paused() (bool, error) {
	return _PowERC20.Contract.Paused(&_PowERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PowERC20 *PowERC20CallerSession) Paused() (bool, error) {
	return _PowERC20.Contract.Paused(&_PowERC20.CallOpts)
}

// Tick is a free data retrieval call binding the contract method 0x3eaf5d9f.
//
// Solidity: function tick() view returns(string)
func (_PowERC20 *PowERC20Caller) Tick(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PowERC20.contract.Call(opts, &out, "tick")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Tick is a free data retrieval call binding the contract method 0x3eaf5d9f.
//
// Solidity: function tick() view returns(string)
func (_PowERC20 *PowERC20Session) Tick() (string, error) {
	return _PowERC20.Contract.Tick(&_PowERC20.CallOpts)
}

// Tick is a free data retrieval call binding the contract method 0x3eaf5d9f.
//
// Solidity: function tick() view returns(string)
func (_PowERC20 *PowERC20CallerSession) Tick() (string, error) {
	return _PowERC20.Contract.Tick(&_PowERC20.CallOpts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x206a810f.
//
// Solidity: function batchTransfer(address _to, uint256 _amt) returns()
func (_PowERC20 *PowERC20Transactor) BatchTransfer(opts *bind.TransactOpts, _to common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _PowERC20.contract.Transact(opts, "batchTransfer", _to, _amt)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x206a810f.
//
// Solidity: function batchTransfer(address _to, uint256 _amt) returns()
func (_PowERC20 *PowERC20Session) BatchTransfer(_to common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _PowERC20.Contract.BatchTransfer(&_PowERC20.TransactOpts, _to, _amt)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x206a810f.
//
// Solidity: function batchTransfer(address _to, uint256 _amt) returns()
func (_PowERC20 *PowERC20TransactorSession) BatchTransfer(_to common.Address, _amt *big.Int) (*types.Transaction, error) {
	return _PowERC20.Contract.BatchTransfer(&_PowERC20.TransactOpts, _to, _amt)
}

// Mine is a paid mutator transaction binding the contract method 0x4d474898.
//
// Solidity: function mine(uint256 nonce) payable returns()
func (_PowERC20 *PowERC20Transactor) Mine(opts *bind.TransactOpts, nonce *big.Int) (*types.Transaction, error) {
	return _PowERC20.contract.Transact(opts, "mine", nonce)
}

// Mine is a paid mutator transaction binding the contract method 0x4d474898.
//
// Solidity: function mine(uint256 nonce) payable returns()
func (_PowERC20 *PowERC20Session) Mine(nonce *big.Int) (*types.Transaction, error) {
	return _PowERC20.Contract.Mine(&_PowERC20.TransactOpts, nonce)
}

// Mine is a paid mutator transaction binding the contract method 0x4d474898.
//
// Solidity: function mine(uint256 nonce) payable returns()
func (_PowERC20 *PowERC20TransactorSession) Mine(nonce *big.Int) (*types.Transaction, error) {
	return _PowERC20.Contract.Mine(&_PowERC20.TransactOpts, nonce)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PowERC20 *PowERC20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowERC20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PowERC20 *PowERC20Session) Pause() (*types.Transaction, error) {
	return _PowERC20.Contract.Pause(&_PowERC20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PowERC20 *PowERC20TransactorSession) Pause() (*types.Transaction, error) {
	return _PowERC20.Contract.Pause(&_PowERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PowERC20 *PowERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PowERC20 *PowERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _PowERC20.Contract.RenounceOwnership(&_PowERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PowERC20 *PowERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PowERC20.Contract.RenounceOwnership(&_PowERC20.TransactOpts)
}

// SetInscription is a paid mutator transaction binding the contract method 0xd593eb79.
//
// Solidity: function setInscription(string inscription) returns()
func (_PowERC20 *PowERC20Transactor) SetInscription(opts *bind.TransactOpts, inscription string) (*types.Transaction, error) {
	return _PowERC20.contract.Transact(opts, "setInscription", inscription)
}

// SetInscription is a paid mutator transaction binding the contract method 0xd593eb79.
//
// Solidity: function setInscription(string inscription) returns()
func (_PowERC20 *PowERC20Session) SetInscription(inscription string) (*types.Transaction, error) {
	return _PowERC20.Contract.SetInscription(&_PowERC20.TransactOpts, inscription)
}

// SetInscription is a paid mutator transaction binding the contract method 0xd593eb79.
//
// Solidity: function setInscription(string inscription) returns()
func (_PowERC20 *PowERC20TransactorSession) SetInscription(inscription string) (*types.Transaction, error) {
	return _PowERC20.Contract.SetInscription(&_PowERC20.TransactOpts, inscription)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address _to) returns()
func (_PowERC20 *PowERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _PowERC20.contract.Transact(opts, "transfer", _to)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address _to) returns()
func (_PowERC20 *PowERC20Session) Transfer(_to common.Address) (*types.Transaction, error) {
	return _PowERC20.Contract.Transfer(&_PowERC20.TransactOpts, _to)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address _to) returns()
func (_PowERC20 *PowERC20TransactorSession) Transfer(_to common.Address) (*types.Transaction, error) {
	return _PowERC20.Contract.Transfer(&_PowERC20.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PowERC20 *PowERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PowERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PowERC20 *PowERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PowERC20.Contract.TransferOwnership(&_PowERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PowERC20 *PowERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PowERC20.Contract.TransferOwnership(&_PowERC20.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PowERC20 *PowERC20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowERC20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PowERC20 *PowERC20Session) Unpause() (*types.Transaction, error) {
	return _PowERC20.Contract.Unpause(&_PowERC20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PowERC20 *PowERC20TransactorSession) Unpause() (*types.Transaction, error) {
	return _PowERC20.Contract.Unpause(&_PowERC20.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PowERC20 *PowERC20Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowERC20.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PowERC20 *PowERC20Session) Withdraw() (*types.Transaction, error) {
	return _PowERC20.Contract.Withdraw(&_PowERC20.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PowERC20 *PowERC20TransactorSession) Withdraw() (*types.Transaction, error) {
	return _PowERC20.Contract.Withdraw(&_PowERC20.TransactOpts)
}

// PowERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PowERC20 contract.
type PowERC20OwnershipTransferredIterator struct {
	Event *PowERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PowERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowERC20OwnershipTransferred)
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
		it.Event = new(PowERC20OwnershipTransferred)
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
func (it *PowERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowERC20OwnershipTransferred represents a OwnershipTransferred event raised by the PowERC20 contract.
type PowERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PowERC20 *PowERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PowERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PowERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PowERC20OwnershipTransferredIterator{contract: _PowERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PowERC20 *PowERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PowERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PowERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowERC20OwnershipTransferred)
				if err := _PowERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PowERC20 *PowERC20Filterer) ParseOwnershipTransferred(log types.Log) (*PowERC20OwnershipTransferred, error) {
	event := new(PowERC20OwnershipTransferred)
	if err := _PowERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowERC20PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PowERC20 contract.
type PowERC20PausedIterator struct {
	Event *PowERC20Paused // Event containing the contract specifics and raw log

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
func (it *PowERC20PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowERC20Paused)
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
		it.Event = new(PowERC20Paused)
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
func (it *PowERC20PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowERC20PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowERC20Paused represents a Paused event raised by the PowERC20 contract.
type PowERC20Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PowERC20 *PowERC20Filterer) FilterPaused(opts *bind.FilterOpts) (*PowERC20PausedIterator, error) {

	logs, sub, err := _PowERC20.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PowERC20PausedIterator{contract: _PowERC20.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PowERC20 *PowERC20Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PowERC20Paused) (event.Subscription, error) {

	logs, sub, err := _PowERC20.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowERC20Paused)
				if err := _PowERC20.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PowERC20 *PowERC20Filterer) ParsePaused(log types.Log) (*PowERC20Paused, error) {
	event := new(PowERC20Paused)
	if err := _PowERC20.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowERC20UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PowERC20 contract.
type PowERC20UnpausedIterator struct {
	Event *PowERC20Unpaused // Event containing the contract specifics and raw log

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
func (it *PowERC20UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowERC20Unpaused)
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
		it.Event = new(PowERC20Unpaused)
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
func (it *PowERC20UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowERC20UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowERC20Unpaused represents a Unpaused event raised by the PowERC20 contract.
type PowERC20Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PowERC20 *PowERC20Filterer) FilterUnpaused(opts *bind.FilterOpts) (*PowERC20UnpausedIterator, error) {

	logs, sub, err := _PowERC20.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PowERC20UnpausedIterator{contract: _PowERC20.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PowERC20 *PowERC20Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PowERC20Unpaused) (event.Subscription, error) {

	logs, sub, err := _PowERC20.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowERC20Unpaused)
				if err := _PowERC20.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PowERC20 *PowERC20Filterer) ParseUnpaused(log types.Log) (*PowERC20Unpaused, error) {
	event := new(PowERC20Unpaused)
	if err := _PowERC20.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowERC20EtherunesProtocolTransferIterator is returned from FilterEtherunesProtocolTransfer and is used to iterate over the raw logs and unpacked data for EtherunesProtocolTransfer events raised by the PowERC20 contract.
type PowERC20EtherunesProtocolTransferIterator struct {
	Event *PowERC20EtherunesProtocolTransfer // Event containing the contract specifics and raw log

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
func (it *PowERC20EtherunesProtocolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowERC20EtherunesProtocolTransfer)
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
		it.Event = new(PowERC20EtherunesProtocolTransfer)
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
func (it *PowERC20EtherunesProtocolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowERC20EtherunesProtocolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowERC20EtherunesProtocolTransfer represents a EtherunesProtocolTransfer event raised by the PowERC20 contract.
type PowERC20EtherunesProtocolTransfer struct {
	From    common.Address
	To      common.Address
	Content string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEtherunesProtocolTransfer is a free log retrieval operation binding the contract event 0x972b1b90380d0be7d8756f3bae915e3329bfbcce0cd513fc79b72884ef5a91e8.
//
// Solidity: event etherunes_protocol_Transfer(address indexed from, address indexed to, string content)
func (_PowERC20 *PowERC20Filterer) FilterEtherunesProtocolTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PowERC20EtherunesProtocolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PowERC20.contract.FilterLogs(opts, "etherunes_protocol_Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PowERC20EtherunesProtocolTransferIterator{contract: _PowERC20.contract, event: "etherunes_protocol_Transfer", logs: logs, sub: sub}, nil
}

// WatchEtherunesProtocolTransfer is a free log subscription operation binding the contract event 0x972b1b90380d0be7d8756f3bae915e3329bfbcce0cd513fc79b72884ef5a91e8.
//
// Solidity: event etherunes_protocol_Transfer(address indexed from, address indexed to, string content)
func (_PowERC20 *PowERC20Filterer) WatchEtherunesProtocolTransfer(opts *bind.WatchOpts, sink chan<- *PowERC20EtherunesProtocolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PowERC20.contract.WatchLogs(opts, "etherunes_protocol_Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowERC20EtherunesProtocolTransfer)
				if err := _PowERC20.contract.UnpackLog(event, "etherunes_protocol_Transfer", log); err != nil {
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

// ParseEtherunesProtocolTransfer is a log parse operation binding the contract event 0x972b1b90380d0be7d8756f3bae915e3329bfbcce0cd513fc79b72884ef5a91e8.
//
// Solidity: event etherunes_protocol_Transfer(address indexed from, address indexed to, string content)
func (_PowERC20 *PowERC20Filterer) ParseEtherunesProtocolTransfer(log types.Log) (*PowERC20EtherunesProtocolTransfer, error) {
	event := new(PowERC20EtherunesProtocolTransfer)
	if err := _PowERC20.contract.UnpackLog(event, "etherunes_protocol_Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowERC20EthrunesProtocolInscribeIterator is returned from FilterEthrunesProtocolInscribe and is used to iterate over the raw logs and unpacked data for EthrunesProtocolInscribe events raised by the PowERC20 contract.
type PowERC20EthrunesProtocolInscribeIterator struct {
	Event *PowERC20EthrunesProtocolInscribe // Event containing the contract specifics and raw log

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
func (it *PowERC20EthrunesProtocolInscribeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowERC20EthrunesProtocolInscribe)
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
		it.Event = new(PowERC20EthrunesProtocolInscribe)
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
func (it *PowERC20EthrunesProtocolInscribeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowERC20EthrunesProtocolInscribeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowERC20EthrunesProtocolInscribe represents a EthrunesProtocolInscribe event raised by the PowERC20 contract.
type PowERC20EthrunesProtocolInscribe struct {
	To      common.Address
	Content string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEthrunesProtocolInscribe is a free log retrieval operation binding the contract event 0xa9b52a0a385a0f4661bf0036806b0a6054494184a759463e02f802fdbf0254c7.
//
// Solidity: event ethrunes_protocol_Inscribe(address indexed to, string content)
func (_PowERC20 *PowERC20Filterer) FilterEthrunesProtocolInscribe(opts *bind.FilterOpts, to []common.Address) (*PowERC20EthrunesProtocolInscribeIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PowERC20.contract.FilterLogs(opts, "ethrunes_protocol_Inscribe", toRule)
	if err != nil {
		return nil, err
	}
	return &PowERC20EthrunesProtocolInscribeIterator{contract: _PowERC20.contract, event: "ethrunes_protocol_Inscribe", logs: logs, sub: sub}, nil
}

// WatchEthrunesProtocolInscribe is a free log subscription operation binding the contract event 0xa9b52a0a385a0f4661bf0036806b0a6054494184a759463e02f802fdbf0254c7.
//
// Solidity: event ethrunes_protocol_Inscribe(address indexed to, string content)
func (_PowERC20 *PowERC20Filterer) WatchEthrunesProtocolInscribe(opts *bind.WatchOpts, sink chan<- *PowERC20EthrunesProtocolInscribe, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PowERC20.contract.WatchLogs(opts, "ethrunes_protocol_Inscribe", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowERC20EthrunesProtocolInscribe)
				if err := _PowERC20.contract.UnpackLog(event, "ethrunes_protocol_Inscribe", log); err != nil {
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

// ParseEthrunesProtocolInscribe is a log parse operation binding the contract event 0xa9b52a0a385a0f4661bf0036806b0a6054494184a759463e02f802fdbf0254c7.
//
// Solidity: event ethrunes_protocol_Inscribe(address indexed to, string content)
func (_PowERC20 *PowERC20Filterer) ParseEthrunesProtocolInscribe(log types.Log) (*PowERC20EthrunesProtocolInscribe, error) {
	event := new(PowERC20EthrunesProtocolInscribe)
	if err := _PowERC20.contract.UnpackLog(event, "ethrunes_protocol_Inscribe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
