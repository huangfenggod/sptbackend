// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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
)

// SptMetaData contains all meta data concerning the Spt contract.
var SptMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blackaddress\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blackaddress\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"tranferBNB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToLp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pancake\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceBNB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLpTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blackaddress\",\"type\":\"address\"}],\"name\":\"isBlackList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpdivid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpdividend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mingpool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SptABI is the input ABI used to generate the binding from.
// Deprecated: Use SptMetaData.ABI instead.
var SptABI = SptMetaData.ABI

// Spt is an auto generated Go binding around an Ethereum contract.
type Spt struct {
	SptCaller     // Read-only binding to the contract
	SptTransactor // Write-only binding to the contract
	SptFilterer   // Log filterer for contract events
}

// SptCaller is an auto generated read-only Go binding around an Ethereum contract.
type SptCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SptTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SptTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SptFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SptFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SptSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SptSession struct {
	Contract     *Spt              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SptCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SptCallerSession struct {
	Contract *SptCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SptTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SptTransactorSession struct {
	Contract     *SptTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SptRaw is an auto generated low-level Go binding around an Ethereum contract.
type SptRaw struct {
	Contract *Spt // Generic contract binding to access the raw methods on
}

// SptCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SptCallerRaw struct {
	Contract *SptCaller // Generic read-only contract binding to access the raw methods on
}

// SptTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SptTransactorRaw struct {
	Contract *SptTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpt creates a new instance of Spt, bound to a specific deployed contract.
func NewSpt(address common.Address, backend bind.ContractBackend) (*Spt, error) {
	contract, err := bindSpt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Spt{SptCaller: SptCaller{contract: contract}, SptTransactor: SptTransactor{contract: contract}, SptFilterer: SptFilterer{contract: contract}}, nil
}

// NewSptCaller creates a new read-only instance of Spt, bound to a specific deployed contract.
func NewSptCaller(address common.Address, caller bind.ContractCaller) (*SptCaller, error) {
	contract, err := bindSpt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SptCaller{contract: contract}, nil
}

// NewSptTransactor creates a new write-only instance of Spt, bound to a specific deployed contract.
func NewSptTransactor(address common.Address, transactor bind.ContractTransactor) (*SptTransactor, error) {
	contract, err := bindSpt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SptTransactor{contract: contract}, nil
}

// NewSptFilterer creates a new log filterer instance of Spt, bound to a specific deployed contract.
func NewSptFilterer(address common.Address, filterer bind.ContractFilterer) (*SptFilterer, error) {
	contract, err := bindSpt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SptFilterer{contract: contract}, nil
}

// bindSpt binds a generic wrapper to an already deployed contract.
func bindSpt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SptABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spt *SptRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Spt.Contract.SptCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spt *SptRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spt.Contract.SptTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spt *SptRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spt.Contract.SptTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spt *SptCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Spt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spt *SptTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spt *SptTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spt.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Spt *SptCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Spt *SptSession) Owner() (common.Address, error) {
	return _Spt.Contract.Owner(&_Spt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Spt *SptCallerSession) Owner() (common.Address, error) {
	return _Spt.Contract.Owner(&_Spt.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Spt *SptCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Spt *SptSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Spt.Contract.Allowance(&_Spt.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Spt *SptCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Spt.Contract.Allowance(&_Spt.CallOpts, owner, spender)
}

// BalanceBNB is a free data retrieval call binding the contract method 0x0b04fffe.
//
// Solidity: function balanceBNB() view returns(uint256)
func (_Spt *SptCaller) BalanceBNB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "balanceBNB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceBNB is a free data retrieval call binding the contract method 0x0b04fffe.
//
// Solidity: function balanceBNB() view returns(uint256)
func (_Spt *SptSession) BalanceBNB() (*big.Int, error) {
	return _Spt.Contract.BalanceBNB(&_Spt.CallOpts)
}

// BalanceBNB is a free data retrieval call binding the contract method 0x0b04fffe.
//
// Solidity: function balanceBNB() view returns(uint256)
func (_Spt *SptCallerSession) BalanceBNB() (*big.Int, error) {
	return _Spt.Contract.BalanceBNB(&_Spt.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Spt *SptCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Spt *SptSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Spt.Contract.BalanceOf(&_Spt.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Spt *SptCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Spt.Contract.BalanceOf(&_Spt.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Spt *SptCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Spt *SptSession) Decimals() (uint8, error) {
	return _Spt.Contract.Decimals(&_Spt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Spt *SptCallerSession) Decimals() (uint8, error) {
	return _Spt.Contract.Decimals(&_Spt.CallOpts)
}

// GetAccountLp is a free data retrieval call binding the contract method 0xcfe619e3.
//
// Solidity: function getAccountLp(address account) view returns(uint256)
func (_Spt *SptCaller) GetAccountLp(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "getAccountLp", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountLp is a free data retrieval call binding the contract method 0xcfe619e3.
//
// Solidity: function getAccountLp(address account) view returns(uint256)
func (_Spt *SptSession) GetAccountLp(account common.Address) (*big.Int, error) {
	return _Spt.Contract.GetAccountLp(&_Spt.CallOpts, account)
}

// GetAccountLp is a free data retrieval call binding the contract method 0xcfe619e3.
//
// Solidity: function getAccountLp(address account) view returns(uint256)
func (_Spt *SptCallerSession) GetAccountLp(account common.Address) (*big.Int, error) {
	return _Spt.Contract.GetAccountLp(&_Spt.CallOpts, account)
}

// GetLpTotalSupply is a free data retrieval call binding the contract method 0x2476a1a0.
//
// Solidity: function getLpTotalSupply() view returns(uint256)
func (_Spt *SptCaller) GetLpTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "getLpTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLpTotalSupply is a free data retrieval call binding the contract method 0x2476a1a0.
//
// Solidity: function getLpTotalSupply() view returns(uint256)
func (_Spt *SptSession) GetLpTotalSupply() (*big.Int, error) {
	return _Spt.Contract.GetLpTotalSupply(&_Spt.CallOpts)
}

// GetLpTotalSupply is a free data retrieval call binding the contract method 0x2476a1a0.
//
// Solidity: function getLpTotalSupply() view returns(uint256)
func (_Spt *SptCallerSession) GetLpTotalSupply() (*big.Int, error) {
	return _Spt.Contract.GetLpTotalSupply(&_Spt.CallOpts)
}

// IsBlackList is a free data retrieval call binding the contract method 0xb36d6919.
//
// Solidity: function isBlackList(address blackaddress) view returns(bool)
func (_Spt *SptCaller) IsBlackList(opts *bind.CallOpts, blackaddress common.Address) (bool, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "isBlackList", blackaddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackList is a free data retrieval call binding the contract method 0xb36d6919.
//
// Solidity: function isBlackList(address blackaddress) view returns(bool)
func (_Spt *SptSession) IsBlackList(blackaddress common.Address) (bool, error) {
	return _Spt.Contract.IsBlackList(&_Spt.CallOpts, blackaddress)
}

// IsBlackList is a free data retrieval call binding the contract method 0xb36d6919.
//
// Solidity: function isBlackList(address blackaddress) view returns(bool)
func (_Spt *SptCallerSession) IsBlackList(blackaddress common.Address) (bool, error) {
	return _Spt.Contract.IsBlackList(&_Spt.CallOpts, blackaddress)
}

// Lpdivid is a free data retrieval call binding the contract method 0xf83ae961.
//
// Solidity: function lpdivid() view returns(uint256)
func (_Spt *SptCaller) Lpdivid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "lpdivid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Lpdivid is a free data retrieval call binding the contract method 0xf83ae961.
//
// Solidity: function lpdivid() view returns(uint256)
func (_Spt *SptSession) Lpdivid() (*big.Int, error) {
	return _Spt.Contract.Lpdivid(&_Spt.CallOpts)
}

// Lpdivid is a free data retrieval call binding the contract method 0xf83ae961.
//
// Solidity: function lpdivid() view returns(uint256)
func (_Spt *SptCallerSession) Lpdivid() (*big.Int, error) {
	return _Spt.Contract.Lpdivid(&_Spt.CallOpts)
}

// Lpdividend is a free data retrieval call binding the contract method 0x52ecf3d4.
//
// Solidity: function lpdividend() view returns(address)
func (_Spt *SptCaller) Lpdividend(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "lpdividend")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lpdividend is a free data retrieval call binding the contract method 0x52ecf3d4.
//
// Solidity: function lpdividend() view returns(address)
func (_Spt *SptSession) Lpdividend() (common.Address, error) {
	return _Spt.Contract.Lpdividend(&_Spt.CallOpts)
}

// Lpdividend is a free data retrieval call binding the contract method 0x52ecf3d4.
//
// Solidity: function lpdividend() view returns(address)
func (_Spt *SptCallerSession) Lpdividend() (common.Address, error) {
	return _Spt.Contract.Lpdividend(&_Spt.CallOpts)
}

// Mingpool is a free data retrieval call binding the contract method 0xf9e048c2.
//
// Solidity: function mingpool() view returns(address)
func (_Spt *SptCaller) Mingpool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "mingpool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mingpool is a free data retrieval call binding the contract method 0xf9e048c2.
//
// Solidity: function mingpool() view returns(address)
func (_Spt *SptSession) Mingpool() (common.Address, error) {
	return _Spt.Contract.Mingpool(&_Spt.CallOpts)
}

// Mingpool is a free data retrieval call binding the contract method 0xf9e048c2.
//
// Solidity: function mingpool() view returns(address)
func (_Spt *SptCallerSession) Mingpool() (common.Address, error) {
	return _Spt.Contract.Mingpool(&_Spt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Spt *SptCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Spt *SptSession) Name() (string, error) {
	return _Spt.Contract.Name(&_Spt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Spt *SptCallerSession) Name() (string, error) {
	return _Spt.Contract.Name(&_Spt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Spt *SptCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Spt *SptSession) Symbol() (string, error) {
	return _Spt.Contract.Symbol(&_Spt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Spt *SptCallerSession) Symbol() (string, error) {
	return _Spt.Contract.Symbol(&_Spt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Spt *SptCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Spt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Spt *SptSession) TotalSupply() (*big.Int, error) {
	return _Spt.Contract.TotalSupply(&_Spt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Spt *SptCallerSession) TotalSupply() (*big.Int, error) {
	return _Spt.Contract.TotalSupply(&_Spt.CallOpts)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address blackaddress) returns(bool)
func (_Spt *SptTransactor) AddBlackList(opts *bind.TransactOpts, blackaddress common.Address) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "addBlackList", blackaddress)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address blackaddress) returns(bool)
func (_Spt *SptSession) AddBlackList(blackaddress common.Address) (*types.Transaction, error) {
	return _Spt.Contract.AddBlackList(&_Spt.TransactOpts, blackaddress)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address blackaddress) returns(bool)
func (_Spt *SptTransactorSession) AddBlackList(blackaddress common.Address) (*types.Transaction, error) {
	return _Spt.Contract.AddBlackList(&_Spt.TransactOpts, blackaddress)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(address account) returns(bool)
func (_Spt *SptTransactor) AddWhiteList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "addWhiteList", account)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(address account) returns(bool)
func (_Spt *SptSession) AddWhiteList(account common.Address) (*types.Transaction, error) {
	return _Spt.Contract.AddWhiteList(&_Spt.TransactOpts, account)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(address account) returns(bool)
func (_Spt *SptTransactorSession) AddWhiteList(account common.Address) (*types.Transaction, error) {
	return _Spt.Contract.AddWhiteList(&_Spt.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Spt *SptTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Spt *SptSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Spt.Contract.Approve(&_Spt.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Spt *SptTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Spt.Contract.Approve(&_Spt.TransactOpts, spender, amount)
}

// GetValue is a paid mutator transaction binding the contract method 0x20965255.
//
// Solidity: function getValue() payable returns(bool)
func (_Spt *SptTransactor) GetValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "getValue")
}

// GetValue is a paid mutator transaction binding the contract method 0x20965255.
//
// Solidity: function getValue() payable returns(bool)
func (_Spt *SptSession) GetValue() (*types.Transaction, error) {
	return _Spt.Contract.GetValue(&_Spt.TransactOpts)
}

// GetValue is a paid mutator transaction binding the contract method 0x20965255.
//
// Solidity: function getValue() payable returns(bool)
func (_Spt *SptTransactorSession) GetValue() (*types.Transaction, error) {
	return _Spt.Contract.GetValue(&_Spt.TransactOpts)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address blackaddress) returns(bool)
func (_Spt *SptTransactor) RemoveBlackList(opts *bind.TransactOpts, blackaddress common.Address) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "removeBlackList", blackaddress)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address blackaddress) returns(bool)
func (_Spt *SptSession) RemoveBlackList(blackaddress common.Address) (*types.Transaction, error) {
	return _Spt.Contract.RemoveBlackList(&_Spt.TransactOpts, blackaddress)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address blackaddress) returns(bool)
func (_Spt *SptTransactorSession) RemoveBlackList(blackaddress common.Address) (*types.Transaction, error) {
	return _Spt.Contract.RemoveBlackList(&_Spt.TransactOpts, blackaddress)
}

// RemoveWhiteList is a paid mutator transaction binding the contract method 0x2042e5c2.
//
// Solidity: function removeWhiteList(address account) returns(bool)
func (_Spt *SptTransactor) RemoveWhiteList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "removeWhiteList", account)
}

// RemoveWhiteList is a paid mutator transaction binding the contract method 0x2042e5c2.
//
// Solidity: function removeWhiteList(address account) returns(bool)
func (_Spt *SptSession) RemoveWhiteList(account common.Address) (*types.Transaction, error) {
	return _Spt.Contract.RemoveWhiteList(&_Spt.TransactOpts, account)
}

// RemoveWhiteList is a paid mutator transaction binding the contract method 0x2042e5c2.
//
// Solidity: function removeWhiteList(address account) returns(bool)
func (_Spt *SptTransactorSession) RemoveWhiteList(account common.Address) (*types.Transaction, error) {
	return _Spt.Contract.RemoveWhiteList(&_Spt.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Spt *SptTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Spt *SptSession) RenounceOwnership() (*types.Transaction, error) {
	return _Spt.Contract.RenounceOwnership(&_Spt.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Spt *SptTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Spt.Contract.RenounceOwnership(&_Spt.TransactOpts)
}

// TranferBNB is a paid mutator transaction binding the contract method 0x289e79cb.
//
// Solidity: function tranferBNB(uint256 amount) payable returns(bool)
func (_Spt *SptTransactor) TranferBNB(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "tranferBNB", amount)
}

// TranferBNB is a paid mutator transaction binding the contract method 0x289e79cb.
//
// Solidity: function tranferBNB(uint256 amount) payable returns(bool)
func (_Spt *SptSession) TranferBNB(amount *big.Int) (*types.Transaction, error) {
	return _Spt.Contract.TranferBNB(&_Spt.TransactOpts, amount)
}

// TranferBNB is a paid mutator transaction binding the contract method 0x289e79cb.
//
// Solidity: function tranferBNB(uint256 amount) payable returns(bool)
func (_Spt *SptTransactorSession) TranferBNB(amount *big.Int) (*types.Transaction, error) {
	return _Spt.Contract.TranferBNB(&_Spt.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Spt *SptTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Spt *SptSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Spt.Contract.Transfer(&_Spt.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Spt *SptTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Spt.Contract.Transfer(&_Spt.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address spender, address to, uint256 amount) returns(bool)
func (_Spt *SptTransactor) TransferFrom(opts *bind.TransactOpts, spender common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "transferFrom", spender, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address spender, address to, uint256 amount) returns(bool)
func (_Spt *SptSession) TransferFrom(spender common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Spt.Contract.TransferFrom(&_Spt.TransactOpts, spender, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address spender, address to, uint256 amount) returns(bool)
func (_Spt *SptTransactorSession) TransferFrom(spender common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Spt.Contract.TransferFrom(&_Spt.TransactOpts, spender, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Spt *SptTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Spt *SptSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Spt.Contract.TransferOwnership(&_Spt.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Spt *SptTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Spt.Contract.TransferOwnership(&_Spt.TransactOpts, newOwner)
}

// TransferToLp is a paid mutator transaction binding the contract method 0x34614468.
//
// Solidity: function transferToLp(uint256 amount) returns(uint256)
func (_Spt *SptTransactor) TransferToLp(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Spt.contract.Transact(opts, "transferToLp", amount)
}

// TransferToLp is a paid mutator transaction binding the contract method 0x34614468.
//
// Solidity: function transferToLp(uint256 amount) returns(uint256)
func (_Spt *SptSession) TransferToLp(amount *big.Int) (*types.Transaction, error) {
	return _Spt.Contract.TransferToLp(&_Spt.TransactOpts, amount)
}

// TransferToLp is a paid mutator transaction binding the contract method 0x34614468.
//
// Solidity: function transferToLp(uint256 amount) returns(uint256)
func (_Spt *SptTransactorSession) TransferToLp(amount *big.Int) (*types.Transaction, error) {
	return _Spt.Contract.TransferToLp(&_Spt.TransactOpts, amount)
}

// SptApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Spt contract.
type SptApprovalIterator struct {
	Event *SptApproval // Event containing the contract specifics and raw log

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
func (it *SptApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SptApproval)
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
		it.Event = new(SptApproval)
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
func (it *SptApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SptApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SptApproval represents a Approval event raised by the Spt contract.
type SptApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Spt *SptFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SptApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Spt.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SptApprovalIterator{contract: _Spt.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Spt *SptFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SptApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Spt.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SptApproval)
				if err := _Spt.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Spt *SptFilterer) ParseApproval(log types.Log) (*SptApproval, error) {
	event := new(SptApproval)
	if err := _Spt.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SptTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Spt contract.
type SptTransferIterator struct {
	Event *SptTransfer // Event containing the contract specifics and raw log

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
func (it *SptTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SptTransfer)
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
		it.Event = new(SptTransfer)
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
func (it *SptTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SptTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SptTransfer represents a Transfer event raised by the Spt contract.
type SptTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Spt *SptFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SptTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Spt.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SptTransferIterator{contract: _Spt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Spt *SptFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SptTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Spt.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SptTransfer)
				if err := _Spt.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Spt *SptFilterer) ParseTransfer(log types.Log) (*SptTransfer, error) {
	event := new(SptTransfer)
	if err := _Spt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
