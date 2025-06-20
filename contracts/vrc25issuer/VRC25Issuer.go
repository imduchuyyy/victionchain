// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrc25issuer

import (
	"math/big"
	"strings"

	"github.com/tomochain/tomochain"
	"github.com/tomochain/tomochain/accounts/abi"
	"github.com/tomochain/tomochain/accounts/abi/bind"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/core/types"
	"github.com/tomochain/tomochain/event"
)

// Vrc25issuerABI is the input ABI used to generate the binding from.
const Vrc25issuerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"minCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenCapacity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"charge\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Apply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"supporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Charge\",\"type\":\"event\"}]"

// Vrc25issuerBin is the compiled bytecode used for deploying new contracts.
const Vrc25issuerBin = `608060405234801561001057600080fd5b50604051602080610c318339810180604052810190808051906020019092919050505080806000819055505050610be58061004c6000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633fa615b01461007d5780638f3a981c146100a85780639d63848a146100ff578063c6b32f341461016b578063d9caed12146101a1578063fc6bd76a1461020e575b600080fd5b34801561008957600080fd5b50610092610244565b6040518082815260200191505060405180910390f35b3480156100b457600080fd5b506100e9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061024d565b6040518082815260200191505060405180910390f35b34801561010b57600080fd5b50610114610296565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561015757808201518184015260208101905061013c565b505050509050019250505060405180910390f35b61019f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610324565b005b3480156101ad57600080fd5b5061020c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610436565b005b610242600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107a0565b005b60008054905090565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600180548060200260200160405190810160405280929190818152602001828054801561031a57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116102d0575b5050505050905090565b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561036157600080fd5b600054341015151561037257600080fd5b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054141515610429576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f546f6b656e20616c7265616479206170706c696564000000000000000000000081525060200191505060405180910390fd5b610432826108ec565b5050565b60008390503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16631d1438486040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156104b657600080fd5b505af11580156104ca573d6000803e3d6000fd5b505050506040513d60208110156104e057600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1614151561057c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f4f6e6c79206973737565722063616e207769746864726177000000000000000081525060200191505060405180910390fd5b81600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151515610659576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f496e73756666696369656e7420636170616369747920746f207769746864726181526020017f770000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b6106ab82600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610b7c90919063ffffffff16565b600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015610734573d6000803e3d6000fd5b508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb846040518082815260200191505060405180910390a350505050565b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156107dd57600080fd5b60005434101515156107ee57600080fd5b61084034600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610b9890919063ffffffff16565b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e06346040518082815260200191505060405180910390a35050565b600081600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561092b57600080fd5b600054341015151561093c57600080fd5b8291503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16631d1438486040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156109ba57600080fd5b505af11580156109ce573d6000803e3d6000fd5b505050506040513d60208110156109e457600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16141515610a1757600080fd5b60018390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610acf34600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610b9890919063ffffffff16565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d716346040518082815260200191505060405180910390a3505050565b6000828211151515610b8d57600080fd5b818303905092915050565b6000808284019050838110151515610baf57600080fd5b80915050929150505600a165627a7a72305820e62ef40e00efd0d62b427e03e509da8acbce302390bbc1a8de6c741b69e9417d0029`

// DeployVrc25issuer deploys a new Ethereum contract, binding an instance of Vrc25issuer to it.
func DeployVrc25issuer(auth *bind.TransactOpts, backend bind.ContractBackend, value *big.Int) (common.Address, *types.Transaction, *Vrc25issuer, error) {
	parsed, err := abi.JSON(strings.NewReader(Vrc25issuerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Vrc25issuerBin), backend, value)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vrc25issuer{Vrc25issuerCaller: Vrc25issuerCaller{contract: contract}, Vrc25issuerTransactor: Vrc25issuerTransactor{contract: contract}, Vrc25issuerFilterer: Vrc25issuerFilterer{contract: contract}}, nil
}

// Vrc25issuer is an auto generated Go binding around an Ethereum contract.
type Vrc25issuer struct {
	Vrc25issuerCaller     // Read-only binding to the contract
	Vrc25issuerTransactor // Write-only binding to the contract
	Vrc25issuerFilterer   // Log filterer for contract events
}

// Vrc25issuerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Vrc25issuerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Vrc25issuerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Vrc25issuerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Vrc25issuerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Vrc25issuerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Vrc25issuerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Vrc25issuerSession struct {
	Contract     *Vrc25issuer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Vrc25issuerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Vrc25issuerCallerSession struct {
	Contract *Vrc25issuerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Vrc25issuerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Vrc25issuerTransactorSession struct {
	Contract     *Vrc25issuerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Vrc25issuerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Vrc25issuerRaw struct {
	Contract *Vrc25issuer // Generic contract binding to access the raw methods on
}

// Vrc25issuerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Vrc25issuerCallerRaw struct {
	Contract *Vrc25issuerCaller // Generic read-only contract binding to access the raw methods on
}

// Vrc25issuerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Vrc25issuerTransactorRaw struct {
	Contract *Vrc25issuerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVrc25issuer creates a new instance of Vrc25issuer, bound to a specific deployed contract.
func NewVrc25issuer(address common.Address, backend bind.ContractBackend) (*Vrc25issuer, error) {
	contract, err := bindVrc25issuer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vrc25issuer{Vrc25issuerCaller: Vrc25issuerCaller{contract: contract}, Vrc25issuerTransactor: Vrc25issuerTransactor{contract: contract}, Vrc25issuerFilterer: Vrc25issuerFilterer{contract: contract}}, nil
}

// NewVrc25issuerCaller creates a new read-only instance of Vrc25issuer, bound to a specific deployed contract.
func NewVrc25issuerCaller(address common.Address, caller bind.ContractCaller) (*Vrc25issuerCaller, error) {
	contract, err := bindVrc25issuer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Vrc25issuerCaller{contract: contract}, nil
}

// NewVrc25issuerTransactor creates a new write-only instance of Vrc25issuer, bound to a specific deployed contract.
func NewVrc25issuerTransactor(address common.Address, transactor bind.ContractTransactor) (*Vrc25issuerTransactor, error) {
	contract, err := bindVrc25issuer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Vrc25issuerTransactor{contract: contract}, nil
}

// NewVrc25issuerFilterer creates a new log filterer instance of Vrc25issuer, bound to a specific deployed contract.
func NewVrc25issuerFilterer(address common.Address, filterer bind.ContractFilterer) (*Vrc25issuerFilterer, error) {
	contract, err := bindVrc25issuer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Vrc25issuerFilterer{contract: contract}, nil
}

// bindVrc25issuer binds a generic wrapper to an already deployed contract.
func bindVrc25issuer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Vrc25issuerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vrc25issuer *Vrc25issuerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vrc25issuer.Contract.Vrc25issuerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vrc25issuer *Vrc25issuerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrc25issuer.Contract.Vrc25issuerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vrc25issuer *Vrc25issuerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vrc25issuer.Contract.Vrc25issuerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vrc25issuer *Vrc25issuerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vrc25issuer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vrc25issuer *Vrc25issuerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrc25issuer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vrc25issuer *Vrc25issuerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vrc25issuer.Contract.contract.Transact(opts, method, params...)
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_Vrc25issuer *Vrc25issuerCaller) GetTokenCapacity(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vrc25issuer.contract.Call(opts, out, "getTokenCapacity", token)
	return *ret0, err
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_Vrc25issuer *Vrc25issuerSession) GetTokenCapacity(token common.Address) (*big.Int, error) {
	return _Vrc25issuer.Contract.GetTokenCapacity(&_Vrc25issuer.CallOpts, token)
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_Vrc25issuer *Vrc25issuerCallerSession) GetTokenCapacity(token common.Address) (*big.Int, error) {
	return _Vrc25issuer.Contract.GetTokenCapacity(&_Vrc25issuer.CallOpts, token)
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_Vrc25issuer *Vrc25issuerCaller) MinCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vrc25issuer.contract.Call(opts, out, "minCap")
	return *ret0, err
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_Vrc25issuer *Vrc25issuerSession) MinCap() (*big.Int, error) {
	return _Vrc25issuer.Contract.MinCap(&_Vrc25issuer.CallOpts)
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_Vrc25issuer *Vrc25issuerCallerSession) MinCap() (*big.Int, error) {
	return _Vrc25issuer.Contract.MinCap(&_Vrc25issuer.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_Vrc25issuer *Vrc25issuerCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Vrc25issuer.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_Vrc25issuer *Vrc25issuerSession) Tokens() ([]common.Address, error) {
	return _Vrc25issuer.Contract.Tokens(&_Vrc25issuer.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_Vrc25issuer *Vrc25issuerCallerSession) Tokens() ([]common.Address, error) {
	return _Vrc25issuer.Contract.Tokens(&_Vrc25issuer.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_Vrc25issuer *Vrc25issuerTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Vrc25issuer.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_Vrc25issuer *Vrc25issuerSession) Apply(token common.Address) (*types.Transaction, error) {
	return _Vrc25issuer.Contract.Apply(&_Vrc25issuer.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_Vrc25issuer *Vrc25issuerTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _Vrc25issuer.Contract.Apply(&_Vrc25issuer.TransactOpts, token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_Vrc25issuer *Vrc25issuerTransactor) Charge(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Vrc25issuer.contract.Transact(opts, "charge", token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_Vrc25issuer *Vrc25issuerSession) Charge(token common.Address) (*types.Transaction, error) {
	return _Vrc25issuer.Contract.Charge(&_Vrc25issuer.TransactOpts, token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_Vrc25issuer *Vrc25issuerTransactorSession) Charge(token common.Address) (*types.Transaction, error) {
	return _Vrc25issuer.Contract.Charge(&_Vrc25issuer.TransactOpts, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(token address, receiver address, amount uint256) returns()
func (_Vrc25issuer *Vrc25issuerTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vrc25issuer.contract.Transact(opts, "withdraw", token, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(token address, receiver address, amount uint256) returns()
func (_Vrc25issuer *Vrc25issuerSession) Withdraw(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vrc25issuer.Contract.Withdraw(&_Vrc25issuer.TransactOpts, token, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(token address, receiver address, amount uint256) returns()
func (_Vrc25issuer *Vrc25issuerTransactorSession) Withdraw(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vrc25issuer.Contract.Withdraw(&_Vrc25issuer.TransactOpts, token, receiver, amount)
}

// Vrc25issuerApplyIterator is returned from FilterApply and is used to iterate over the raw logs and unpacked data for Apply events raised by the Vrc25issuer contract.
type Vrc25issuerApplyIterator struct {
	Event *Vrc25issuerApply // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  tomochain.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vrc25issuerApplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vrc25issuerApply)
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
		it.Event = new(Vrc25issuerApply)
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
func (it *Vrc25issuerApplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vrc25issuerApplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vrc25issuerApply represents a Apply event raised by the Vrc25issuer contract.
type Vrc25issuerApply struct {
	Issuer common.Address
	Token  common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApply is a free log retrieval operation binding the contract event 0x2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d716.
//
// Solidity: event Apply(issuer indexed address, token indexed address, value uint256)
func (_Vrc25issuer *Vrc25issuerFilterer) FilterApply(opts *bind.FilterOpts, issuer []common.Address, token []common.Address) (*Vrc25issuerApplyIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vrc25issuer.contract.FilterLogs(opts, "Apply", issuerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &Vrc25issuerApplyIterator{contract: _Vrc25issuer.contract, event: "Apply", logs: logs, sub: sub}, nil
}

// WatchApply is a free log subscription operation binding the contract event 0x2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d716.
//
// Solidity: event Apply(issuer indexed address, token indexed address, value uint256)
func (_Vrc25issuer *Vrc25issuerFilterer) WatchApply(opts *bind.WatchOpts, sink chan<- *Vrc25issuerApply, issuer []common.Address, token []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vrc25issuer.contract.WatchLogs(opts, "Apply", issuerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vrc25issuerApply)
				if err := _Vrc25issuer.contract.UnpackLog(event, "Apply", log); err != nil {
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

// Vrc25issuerChargeIterator is returned from FilterCharge and is used to iterate over the raw logs and unpacked data for Charge events raised by the Vrc25issuer contract.
type Vrc25issuerChargeIterator struct {
	Event *Vrc25issuerCharge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  tomochain.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vrc25issuerChargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vrc25issuerCharge)
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
		it.Event = new(Vrc25issuerCharge)
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
func (it *Vrc25issuerChargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vrc25issuerChargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vrc25issuerCharge represents a Charge event raised by the Vrc25issuer contract.
type Vrc25issuerCharge struct {
	Supporter common.Address
	Token     common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCharge is a free log retrieval operation binding the contract event 0x5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e06.
//
// Solidity: event Charge(supporter indexed address, token indexed address, value uint256)
func (_Vrc25issuer *Vrc25issuerFilterer) FilterCharge(opts *bind.FilterOpts, supporter []common.Address, token []common.Address) (*Vrc25issuerChargeIterator, error) {

	var supporterRule []interface{}
	for _, supporterItem := range supporter {
		supporterRule = append(supporterRule, supporterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vrc25issuer.contract.FilterLogs(opts, "Charge", supporterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &Vrc25issuerChargeIterator{contract: _Vrc25issuer.contract, event: "Charge", logs: logs, sub: sub}, nil
}

// WatchCharge is a free log subscription operation binding the contract event 0x5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e06.
//
// Solidity: event Charge(supporter indexed address, token indexed address, value uint256)
func (_Vrc25issuer *Vrc25issuerFilterer) WatchCharge(opts *bind.WatchOpts, sink chan<- *Vrc25issuerCharge, supporter []common.Address, token []common.Address) (event.Subscription, error) {

	var supporterRule []interface{}
	for _, supporterItem := range supporter {
		supporterRule = append(supporterRule, supporterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vrc25issuer.contract.WatchLogs(opts, "Charge", supporterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vrc25issuerCharge)
				if err := _Vrc25issuer.contract.UnpackLog(event, "Charge", log); err != nil {
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

// Vrc25issuerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Vrc25issuer contract.
type Vrc25issuerWithdrawIterator struct {
	Event *Vrc25issuerWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  tomochain.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Vrc25issuerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vrc25issuerWithdraw)
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
		it.Event = new(Vrc25issuerWithdraw)
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
func (it *Vrc25issuerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vrc25issuerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vrc25issuerWithdraw represents a Withdraw event raised by the Vrc25issuer contract.
type Vrc25issuerWithdraw struct {
	Token    common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(token indexed address, receiver indexed address, value uint256)
func (_Vrc25issuer *Vrc25issuerFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*Vrc25issuerWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Vrc25issuer.contract.FilterLogs(opts, "Withdraw", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &Vrc25issuerWithdrawIterator{contract: _Vrc25issuer.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(token indexed address, receiver indexed address, value uint256)
func (_Vrc25issuer *Vrc25issuerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Vrc25issuerWithdraw, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Vrc25issuer.contract.WatchLogs(opts, "Withdraw", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vrc25issuerWithdraw)
				if err := _Vrc25issuer.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
