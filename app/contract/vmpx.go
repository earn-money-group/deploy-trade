// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// VmpxMetaData contains all meta data concerning the Vmpx contract.
var VmpxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycles_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlockNumber_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTHORS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BATCH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cycles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VmpxABI is the input ABI used to generate the binding from.
// Deprecated: Use VmpxMetaData.ABI instead.
var VmpxABI = VmpxMetaData.ABI

// Vmpx is an auto generated Go binding around an Ethereum contract.
type Vmpx struct {
	VmpxCaller     // Read-only binding to the contract
	VmpxTransactor // Write-only binding to the contract
	VmpxFilterer   // Log filterer for contract events
}

// VmpxCaller is an auto generated read-only Go binding around an Ethereum contract.
type VmpxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmpxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VmpxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmpxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VmpxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmpxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VmpxSession struct {
	Contract     *Vmpx             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmpxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VmpxCallerSession struct {
	Contract *VmpxCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VmpxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VmpxTransactorSession struct {
	Contract     *VmpxTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmpxRaw is an auto generated low-level Go binding around an Ethereum contract.
type VmpxRaw struct {
	Contract *Vmpx // Generic contract binding to access the raw methods on
}

// VmpxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VmpxCallerRaw struct {
	Contract *VmpxCaller // Generic read-only contract binding to access the raw methods on
}

// VmpxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VmpxTransactorRaw struct {
	Contract *VmpxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVmpx creates a new instance of Vmpx, bound to a specific deployed contract.
func NewVmpx(address common.Address, backend bind.ContractBackend) (*Vmpx, error) {
	contract, err := bindVmpx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vmpx{VmpxCaller: VmpxCaller{contract: contract}, VmpxTransactor: VmpxTransactor{contract: contract}, VmpxFilterer: VmpxFilterer{contract: contract}}, nil
}

// NewVmpxCaller creates a new read-only instance of Vmpx, bound to a specific deployed contract.
func NewVmpxCaller(address common.Address, caller bind.ContractCaller) (*VmpxCaller, error) {
	contract, err := bindVmpx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VmpxCaller{contract: contract}, nil
}

// NewVmpxTransactor creates a new write-only instance of Vmpx, bound to a specific deployed contract.
func NewVmpxTransactor(address common.Address, transactor bind.ContractTransactor) (*VmpxTransactor, error) {
	contract, err := bindVmpx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VmpxTransactor{contract: contract}, nil
}

// NewVmpxFilterer creates a new log filterer instance of Vmpx, bound to a specific deployed contract.
func NewVmpxFilterer(address common.Address, filterer bind.ContractFilterer) (*VmpxFilterer, error) {
	contract, err := bindVmpx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VmpxFilterer{contract: contract}, nil
}

// bindVmpx binds a generic wrapper to an already deployed contract.
func bindVmpx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VmpxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vmpx *VmpxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vmpx.Contract.VmpxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vmpx *VmpxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vmpx.Contract.VmpxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vmpx *VmpxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vmpx.Contract.VmpxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vmpx *VmpxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vmpx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vmpx *VmpxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vmpx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vmpx *VmpxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vmpx.Contract.contract.Transact(opts, method, params...)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Vmpx *VmpxCaller) AUTHORS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "AUTHORS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Vmpx *VmpxSession) AUTHORS() (string, error) {
	return _Vmpx.Contract.AUTHORS(&_Vmpx.CallOpts)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Vmpx *VmpxCallerSession) AUTHORS() (string, error) {
	return _Vmpx.Contract.AUTHORS(&_Vmpx.CallOpts)
}

// BATCH is a free data retrieval call binding the contract method 0x5207d321.
//
// Solidity: function BATCH() view returns(uint256)
func (_Vmpx *VmpxCaller) BATCH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "BATCH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BATCH is a free data retrieval call binding the contract method 0x5207d321.
//
// Solidity: function BATCH() view returns(uint256)
func (_Vmpx *VmpxSession) BATCH() (*big.Int, error) {
	return _Vmpx.Contract.BATCH(&_Vmpx.CallOpts)
}

// BATCH is a free data retrieval call binding the contract method 0x5207d321.
//
// Solidity: function BATCH() view returns(uint256)
func (_Vmpx *VmpxCallerSession) BATCH() (*big.Int, error) {
	return _Vmpx.Contract.BATCH(&_Vmpx.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vmpx *VmpxCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vmpx *VmpxSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Vmpx.Contract.Allowance(&_Vmpx.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vmpx *VmpxCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Vmpx.Contract.Allowance(&_Vmpx.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Vmpx *VmpxCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Vmpx *VmpxSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Vmpx.Contract.BalanceOf(&_Vmpx.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Vmpx *VmpxCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Vmpx.Contract.BalanceOf(&_Vmpx.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Vmpx *VmpxCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Vmpx *VmpxSession) Cap() (*big.Int, error) {
	return _Vmpx.Contract.Cap(&_Vmpx.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Vmpx *VmpxCallerSession) Cap() (*big.Int, error) {
	return _Vmpx.Contract.Cap(&_Vmpx.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Vmpx *VmpxCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Vmpx *VmpxSession) Counter() (*big.Int, error) {
	return _Vmpx.Contract.Counter(&_Vmpx.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Vmpx *VmpxCallerSession) Counter() (*big.Int, error) {
	return _Vmpx.Contract.Counter(&_Vmpx.CallOpts)
}

// Cycles is a free data retrieval call binding the contract method 0x6dbe5554.
//
// Solidity: function cycles() view returns(uint256)
func (_Vmpx *VmpxCaller) Cycles(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "cycles")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cycles is a free data retrieval call binding the contract method 0x6dbe5554.
//
// Solidity: function cycles() view returns(uint256)
func (_Vmpx *VmpxSession) Cycles() (*big.Int, error) {
	return _Vmpx.Contract.Cycles(&_Vmpx.CallOpts)
}

// Cycles is a free data retrieval call binding the contract method 0x6dbe5554.
//
// Solidity: function cycles() view returns(uint256)
func (_Vmpx *VmpxCallerSession) Cycles() (*big.Int, error) {
	return _Vmpx.Contract.Cycles(&_Vmpx.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vmpx *VmpxCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vmpx *VmpxSession) Decimals() (uint8, error) {
	return _Vmpx.Contract.Decimals(&_Vmpx.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vmpx *VmpxCallerSession) Decimals() (uint8, error) {
	return _Vmpx.Contract.Decimals(&_Vmpx.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vmpx *VmpxCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vmpx *VmpxSession) Name() (string, error) {
	return _Vmpx.Contract.Name(&_Vmpx.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vmpx *VmpxCallerSession) Name() (string, error) {
	return _Vmpx.Contract.Name(&_Vmpx.CallOpts)
}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_Vmpx *VmpxCaller) StartBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "startBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_Vmpx *VmpxSession) StartBlockNumber() (*big.Int, error) {
	return _Vmpx.Contract.StartBlockNumber(&_Vmpx.CallOpts)
}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_Vmpx *VmpxCallerSession) StartBlockNumber() (*big.Int, error) {
	return _Vmpx.Contract.StartBlockNumber(&_Vmpx.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vmpx *VmpxCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vmpx *VmpxSession) Symbol() (string, error) {
	return _Vmpx.Contract.Symbol(&_Vmpx.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vmpx *VmpxCallerSession) Symbol() (string, error) {
	return _Vmpx.Contract.Symbol(&_Vmpx.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vmpx *VmpxCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vmpx.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vmpx *VmpxSession) TotalSupply() (*big.Int, error) {
	return _Vmpx.Contract.TotalSupply(&_Vmpx.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vmpx *VmpxCallerSession) TotalSupply() (*big.Int, error) {
	return _Vmpx.Contract.TotalSupply(&_Vmpx.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vmpx *VmpxTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vmpx.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vmpx *VmpxSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.Approve(&_Vmpx.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vmpx *VmpxTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.Approve(&_Vmpx.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Vmpx *VmpxTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Vmpx.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Vmpx *VmpxSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.DecreaseAllowance(&_Vmpx.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Vmpx *VmpxTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.DecreaseAllowance(&_Vmpx.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Vmpx *VmpxTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Vmpx.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Vmpx *VmpxSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.IncreaseAllowance(&_Vmpx.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Vmpx *VmpxTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.IncreaseAllowance(&_Vmpx.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 power) returns()
func (_Vmpx *VmpxTransactor) Mint(opts *bind.TransactOpts, power *big.Int) (*types.Transaction, error) {
	return _Vmpx.contract.Transact(opts, "mint", power)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 power) returns()
func (_Vmpx *VmpxSession) Mint(power *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.Mint(&_Vmpx.TransactOpts, power)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 power) returns()
func (_Vmpx *VmpxTransactorSession) Mint(power *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.Mint(&_Vmpx.TransactOpts, power)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Vmpx *VmpxTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vmpx.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Vmpx *VmpxSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.Transfer(&_Vmpx.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Vmpx *VmpxTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.Transfer(&_Vmpx.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Vmpx *VmpxTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vmpx.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Vmpx *VmpxSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.TransferFrom(&_Vmpx.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Vmpx *VmpxTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vmpx.Contract.TransferFrom(&_Vmpx.TransactOpts, from, to, amount)
}

// VmpxApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vmpx contract.
type VmpxApprovalIterator struct {
	Event *VmpxApproval // Event containing the contract specifics and raw log

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
func (it *VmpxApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmpxApproval)
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
		it.Event = new(VmpxApproval)
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
func (it *VmpxApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmpxApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmpxApproval represents a Approval event raised by the Vmpx contract.
type VmpxApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Vmpx *VmpxFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VmpxApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vmpx.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VmpxApprovalIterator{contract: _Vmpx.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Vmpx *VmpxFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VmpxApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vmpx.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmpxApproval)
				if err := _Vmpx.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Vmpx *VmpxFilterer) ParseApproval(log types.Log) (*VmpxApproval, error) {
	event := new(VmpxApproval)
	if err := _Vmpx.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmpxTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vmpx contract.
type VmpxTransferIterator struct {
	Event *VmpxTransfer // Event containing the contract specifics and raw log

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
func (it *VmpxTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmpxTransfer)
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
		it.Event = new(VmpxTransfer)
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
func (it *VmpxTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmpxTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmpxTransfer represents a Transfer event raised by the Vmpx contract.
type VmpxTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Vmpx *VmpxFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VmpxTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vmpx.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VmpxTransferIterator{contract: _Vmpx.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Vmpx *VmpxFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VmpxTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vmpx.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmpxTransfer)
				if err := _Vmpx.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Vmpx *VmpxFilterer) ParseTransfer(log types.Log) (*VmpxTransfer, error) {
	event := new(VmpxTransfer)
	if err := _Vmpx.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
