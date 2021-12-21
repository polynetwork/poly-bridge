// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main_chain_lock_proxy_abi

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

var (
	MethodApprove = "approve"

	MethodLock = "lock"

	MethodAllowance = "allowance"

	MethodGetSideChainLockAmount = "getSideChainLockAmount"

	MethodName = "name"

	EventApproval = "Approval"

	EventCrossChainEvent = "CrossChainEvent"

	EventLockEvent = "LockEvent"

	EventUnlockEvent = "UnlockEvent"

	EventVerifyHeaderAndExecuteTxEvent = "VerifyHeaderAndExecuteTxEvent"
)

// IMainChainLockProxyABI is the input ABI used to generate the binding from.
const IMainChainLockProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"txId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyOrAssetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawdata\",\"type\":\"bytes\"}],\"name\":\"CrossChainEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fromChainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"crossChainTxHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fromChainTxHash\",\"type\":\"bytes\"}],\"name\":\"VerifyHeaderAndExecuteTxEvent\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getSideChainLockAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IMainChainLockProxyFuncSigs maps the 4-byte function signature to its string representation.
var IMainChainLockProxyFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"50d06e71": "getSideChainLockAmount(uint64)",
	"4bc68823": "lock(uint64,address,uint256)",
	"06fdde03": "name()",
}

// IMainChainLockProxy is an auto generated Go binding around an Ethereum contract.
type IMainChainLockProxy struct {
	IMainChainLockProxyCaller     // Read-only binding to the contract
	IMainChainLockProxyTransactor // Write-only binding to the contract
	IMainChainLockProxyFilterer   // Log filterer for contract events
}

// IMainChainLockProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMainChainLockProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMainChainLockProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMainChainLockProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMainChainLockProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMainChainLockProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMainChainLockProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMainChainLockProxySession struct {
	Contract     *IMainChainLockProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMainChainLockProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMainChainLockProxyCallerSession struct {
	Contract *IMainChainLockProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMainChainLockProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMainChainLockProxyTransactorSession struct {
	Contract     *IMainChainLockProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMainChainLockProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMainChainLockProxyRaw struct {
	Contract *IMainChainLockProxy // Generic contract binding to access the raw methods on
}

// IMainChainLockProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMainChainLockProxyCallerRaw struct {
	Contract *IMainChainLockProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IMainChainLockProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMainChainLockProxyTransactorRaw struct {
	Contract *IMainChainLockProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMainChainLockProxy creates a new instance of IMainChainLockProxy, bound to a specific deployed contract.
func NewIMainChainLockProxy(address common.Address, backend bind.ContractBackend) (*IMainChainLockProxy, error) {
	contract, err := bindIMainChainLockProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMainChainLockProxy{IMainChainLockProxyCaller: IMainChainLockProxyCaller{contract: contract}, IMainChainLockProxyTransactor: IMainChainLockProxyTransactor{contract: contract}, IMainChainLockProxyFilterer: IMainChainLockProxyFilterer{contract: contract}}, nil
}

// NewIMainChainLockProxyCaller creates a new read-only instance of IMainChainLockProxy, bound to a specific deployed contract.
func NewIMainChainLockProxyCaller(address common.Address, caller bind.ContractCaller) (*IMainChainLockProxyCaller, error) {
	contract, err := bindIMainChainLockProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMainChainLockProxyCaller{contract: contract}, nil
}

// NewIMainChainLockProxyTransactor creates a new write-only instance of IMainChainLockProxy, bound to a specific deployed contract.
func NewIMainChainLockProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IMainChainLockProxyTransactor, error) {
	contract, err := bindIMainChainLockProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMainChainLockProxyTransactor{contract: contract}, nil
}

// NewIMainChainLockProxyFilterer creates a new log filterer instance of IMainChainLockProxy, bound to a specific deployed contract.
func NewIMainChainLockProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IMainChainLockProxyFilterer, error) {
	contract, err := bindIMainChainLockProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMainChainLockProxyFilterer{contract: contract}, nil
}

// bindIMainChainLockProxy binds a generic wrapper to an already deployed contract.
func bindIMainChainLockProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMainChainLockProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMainChainLockProxy *IMainChainLockProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMainChainLockProxy.Contract.IMainChainLockProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMainChainLockProxy *IMainChainLockProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMainChainLockProxy.Contract.IMainChainLockProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMainChainLockProxy *IMainChainLockProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMainChainLockProxy.Contract.IMainChainLockProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMainChainLockProxy *IMainChainLockProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMainChainLockProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMainChainLockProxy *IMainChainLockProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMainChainLockProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMainChainLockProxy *IMainChainLockProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMainChainLockProxy.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMainChainLockProxy *IMainChainLockProxyCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IMainChainLockProxy.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMainChainLockProxy *IMainChainLockProxySession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IMainChainLockProxy.Contract.Allowance(&_IMainChainLockProxy.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMainChainLockProxy *IMainChainLockProxyCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IMainChainLockProxy.Contract.Allowance(&_IMainChainLockProxy.CallOpts, owner, spender)
}

// GetSideChainLockAmount is a free data retrieval call binding the contract method 0x50d06e71.
//
// Solidity: function getSideChainLockAmount(uint64 chainId) view returns(uint256)
func (_IMainChainLockProxy *IMainChainLockProxyCaller) GetSideChainLockAmount(opts *bind.CallOpts, chainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _IMainChainLockProxy.contract.Call(opts, &out, "getSideChainLockAmount", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSideChainLockAmount is a free data retrieval call binding the contract method 0x50d06e71.
//
// Solidity: function getSideChainLockAmount(uint64 chainId) view returns(uint256)
func (_IMainChainLockProxy *IMainChainLockProxySession) GetSideChainLockAmount(chainId uint64) (*big.Int, error) {
	return _IMainChainLockProxy.Contract.GetSideChainLockAmount(&_IMainChainLockProxy.CallOpts, chainId)
}

// GetSideChainLockAmount is a free data retrieval call binding the contract method 0x50d06e71.
//
// Solidity: function getSideChainLockAmount(uint64 chainId) view returns(uint256)
func (_IMainChainLockProxy *IMainChainLockProxyCallerSession) GetSideChainLockAmount(chainId uint64) (*big.Int, error) {
	return _IMainChainLockProxy.Contract.GetSideChainLockAmount(&_IMainChainLockProxy.CallOpts, chainId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IMainChainLockProxy *IMainChainLockProxyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IMainChainLockProxy.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IMainChainLockProxy *IMainChainLockProxySession) Name() (string, error) {
	return _IMainChainLockProxy.Contract.Name(&_IMainChainLockProxy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IMainChainLockProxy *IMainChainLockProxyCallerSession) Name() (string, error) {
	return _IMainChainLockProxy.Contract.Name(&_IMainChainLockProxy.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IMainChainLockProxy *IMainChainLockProxyTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMainChainLockProxy.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IMainChainLockProxy *IMainChainLockProxySession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMainChainLockProxy.Contract.Approve(&_IMainChainLockProxy.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IMainChainLockProxy *IMainChainLockProxyTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMainChainLockProxy.Contract.Approve(&_IMainChainLockProxy.TransactOpts, spender, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x4bc68823.
//
// Solidity: function lock(uint64 toChainId, address toAddress, uint256 amount) payable returns(bool)
func (_IMainChainLockProxy *IMainChainLockProxyTransactor) Lock(opts *bind.TransactOpts, toChainId uint64, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMainChainLockProxy.contract.Transact(opts, "lock", toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x4bc68823.
//
// Solidity: function lock(uint64 toChainId, address toAddress, uint256 amount) payable returns(bool)
func (_IMainChainLockProxy *IMainChainLockProxySession) Lock(toChainId uint64, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMainChainLockProxy.Contract.Lock(&_IMainChainLockProxy.TransactOpts, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x4bc68823.
//
// Solidity: function lock(uint64 toChainId, address toAddress, uint256 amount) payable returns(bool)
func (_IMainChainLockProxy *IMainChainLockProxyTransactorSession) Lock(toChainId uint64, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMainChainLockProxy.Contract.Lock(&_IMainChainLockProxy.TransactOpts, toChainId, toAddress, amount)
}

// IMainChainLockProxyApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IMainChainLockProxy contract.
type IMainChainLockProxyApprovalIterator struct {
	Event *IMainChainLockProxyApproval // Event containing the contract specifics and raw log

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
func (it *IMainChainLockProxyApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMainChainLockProxyApproval)
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
		it.Event = new(IMainChainLockProxyApproval)
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
func (it *IMainChainLockProxyApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMainChainLockProxyApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMainChainLockProxyApproval represents a Approval event raised by the IMainChainLockProxy contract.
type IMainChainLockProxyApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IMainChainLockProxyApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IMainChainLockProxy.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IMainChainLockProxyApprovalIterator{contract: _IMainChainLockProxy.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IMainChainLockProxyApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IMainChainLockProxy.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMainChainLockProxyApproval)
				if err := _IMainChainLockProxy.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) ParseApproval(log types.Log) (*IMainChainLockProxyApproval, error) {
	event := new(IMainChainLockProxyApproval)
	if err := _IMainChainLockProxy.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMainChainLockProxyCrossChainEventIterator is returned from FilterCrossChainEvent and is used to iterate over the raw logs and unpacked data for CrossChainEvent events raised by the IMainChainLockProxy contract.
type IMainChainLockProxyCrossChainEventIterator struct {
	Event *IMainChainLockProxyCrossChainEvent // Event containing the contract specifics and raw log

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
func (it *IMainChainLockProxyCrossChainEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMainChainLockProxyCrossChainEvent)
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
		it.Event = new(IMainChainLockProxyCrossChainEvent)
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
func (it *IMainChainLockProxyCrossChainEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMainChainLockProxyCrossChainEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMainChainLockProxyCrossChainEvent represents a CrossChainEvent event raised by the IMainChainLockProxy contract.
type IMainChainLockProxyCrossChainEvent struct {
	Sender               common.Address
	TxId                 []byte
	ProxyOrAssetContract common.Address
	ToChainId            uint64
	ToContract           []byte
	Rawdata              []byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCrossChainEvent is a free log retrieval operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) FilterCrossChainEvent(opts *bind.FilterOpts) (*IMainChainLockProxyCrossChainEventIterator, error) {

	logs, sub, err := _IMainChainLockProxy.contract.FilterLogs(opts, "CrossChainEvent")
	if err != nil {
		return nil, err
	}
	return &IMainChainLockProxyCrossChainEventIterator{contract: _IMainChainLockProxy.contract, event: "CrossChainEvent", logs: logs, sub: sub}, nil
}

// WatchCrossChainEvent is a free log subscription operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) WatchCrossChainEvent(opts *bind.WatchOpts, sink chan<- *IMainChainLockProxyCrossChainEvent) (event.Subscription, error) {

	logs, sub, err := _IMainChainLockProxy.contract.WatchLogs(opts, "CrossChainEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMainChainLockProxyCrossChainEvent)
				if err := _IMainChainLockProxy.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
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

// ParseCrossChainEvent is a log parse operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) ParseCrossChainEvent(log types.Log) (*IMainChainLockProxyCrossChainEvent, error) {
	event := new(IMainChainLockProxyCrossChainEvent)
	if err := _IMainChainLockProxy.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMainChainLockProxyLockEventIterator is returned from FilterLockEvent and is used to iterate over the raw logs and unpacked data for LockEvent events raised by the IMainChainLockProxy contract.
type IMainChainLockProxyLockEventIterator struct {
	Event *IMainChainLockProxyLockEvent // Event containing the contract specifics and raw log

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
func (it *IMainChainLockProxyLockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMainChainLockProxyLockEvent)
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
		it.Event = new(IMainChainLockProxyLockEvent)
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
func (it *IMainChainLockProxyLockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMainChainLockProxyLockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMainChainLockProxyLockEvent represents a LockEvent event raised by the IMainChainLockProxy contract.
type IMainChainLockProxyLockEvent struct {
	FromAssetHash common.Address
	FromAddress   common.Address
	ToChainId     uint64
	ToAssetHash   []byte
	ToAddress     []byte
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLockEvent is a free log retrieval operation binding the contract event 0x8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea5.
//
// Solidity: event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) FilterLockEvent(opts *bind.FilterOpts) (*IMainChainLockProxyLockEventIterator, error) {

	logs, sub, err := _IMainChainLockProxy.contract.FilterLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return &IMainChainLockProxyLockEventIterator{contract: _IMainChainLockProxy.contract, event: "LockEvent", logs: logs, sub: sub}, nil
}

// WatchLockEvent is a free log subscription operation binding the contract event 0x8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea5.
//
// Solidity: event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) WatchLockEvent(opts *bind.WatchOpts, sink chan<- *IMainChainLockProxyLockEvent) (event.Subscription, error) {

	logs, sub, err := _IMainChainLockProxy.contract.WatchLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMainChainLockProxyLockEvent)
				if err := _IMainChainLockProxy.contract.UnpackLog(event, "LockEvent", log); err != nil {
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

// ParseLockEvent is a log parse operation binding the contract event 0x8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea5.
//
// Solidity: event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) ParseLockEvent(log types.Log) (*IMainChainLockProxyLockEvent, error) {
	event := new(IMainChainLockProxyLockEvent)
	if err := _IMainChainLockProxy.contract.UnpackLog(event, "LockEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMainChainLockProxyUnlockEventIterator is returned from FilterUnlockEvent and is used to iterate over the raw logs and unpacked data for UnlockEvent events raised by the IMainChainLockProxy contract.
type IMainChainLockProxyUnlockEventIterator struct {
	Event *IMainChainLockProxyUnlockEvent // Event containing the contract specifics and raw log

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
func (it *IMainChainLockProxyUnlockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMainChainLockProxyUnlockEvent)
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
		it.Event = new(IMainChainLockProxyUnlockEvent)
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
func (it *IMainChainLockProxyUnlockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMainChainLockProxyUnlockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMainChainLockProxyUnlockEvent represents a UnlockEvent event raised by the IMainChainLockProxy contract.
type IMainChainLockProxyUnlockEvent struct {
	ToAssetHash common.Address
	ToAddress   common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnlockEvent is a free log retrieval operation binding the contract event 0xd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df.
//
// Solidity: event UnlockEvent(address toAssetHash, address toAddress, uint256 amount)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) FilterUnlockEvent(opts *bind.FilterOpts) (*IMainChainLockProxyUnlockEventIterator, error) {

	logs, sub, err := _IMainChainLockProxy.contract.FilterLogs(opts, "UnlockEvent")
	if err != nil {
		return nil, err
	}
	return &IMainChainLockProxyUnlockEventIterator{contract: _IMainChainLockProxy.contract, event: "UnlockEvent", logs: logs, sub: sub}, nil
}

// WatchUnlockEvent is a free log subscription operation binding the contract event 0xd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df.
//
// Solidity: event UnlockEvent(address toAssetHash, address toAddress, uint256 amount)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) WatchUnlockEvent(opts *bind.WatchOpts, sink chan<- *IMainChainLockProxyUnlockEvent) (event.Subscription, error) {

	logs, sub, err := _IMainChainLockProxy.contract.WatchLogs(opts, "UnlockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMainChainLockProxyUnlockEvent)
				if err := _IMainChainLockProxy.contract.UnpackLog(event, "UnlockEvent", log); err != nil {
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

// ParseUnlockEvent is a log parse operation binding the contract event 0xd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df.
//
// Solidity: event UnlockEvent(address toAssetHash, address toAddress, uint256 amount)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) ParseUnlockEvent(log types.Log) (*IMainChainLockProxyUnlockEvent, error) {
	event := new(IMainChainLockProxyUnlockEvent)
	if err := _IMainChainLockProxy.contract.UnpackLog(event, "UnlockEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMainChainLockProxyVerifyHeaderAndExecuteTxEventIterator is returned from FilterVerifyHeaderAndExecuteTxEvent and is used to iterate over the raw logs and unpacked data for VerifyHeaderAndExecuteTxEvent events raised by the IMainChainLockProxy contract.
type IMainChainLockProxyVerifyHeaderAndExecuteTxEventIterator struct {
	Event *IMainChainLockProxyVerifyHeaderAndExecuteTxEvent // Event containing the contract specifics and raw log

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
func (it *IMainChainLockProxyVerifyHeaderAndExecuteTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMainChainLockProxyVerifyHeaderAndExecuteTxEvent)
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
		it.Event = new(IMainChainLockProxyVerifyHeaderAndExecuteTxEvent)
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
func (it *IMainChainLockProxyVerifyHeaderAndExecuteTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMainChainLockProxyVerifyHeaderAndExecuteTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMainChainLockProxyVerifyHeaderAndExecuteTxEvent represents a VerifyHeaderAndExecuteTxEvent event raised by the IMainChainLockProxy contract.
type IMainChainLockProxyVerifyHeaderAndExecuteTxEvent struct {
	FromChainID      uint64
	ToContract       []byte
	CrossChainTxHash []byte
	FromChainTxHash  []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVerifyHeaderAndExecuteTxEvent is a free log retrieval operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) FilterVerifyHeaderAndExecuteTxEvent(opts *bind.FilterOpts) (*IMainChainLockProxyVerifyHeaderAndExecuteTxEventIterator, error) {

	logs, sub, err := _IMainChainLockProxy.contract.FilterLogs(opts, "VerifyHeaderAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return &IMainChainLockProxyVerifyHeaderAndExecuteTxEventIterator{contract: _IMainChainLockProxy.contract, event: "VerifyHeaderAndExecuteTxEvent", logs: logs, sub: sub}, nil
}

// WatchVerifyHeaderAndExecuteTxEvent is a free log subscription operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) WatchVerifyHeaderAndExecuteTxEvent(opts *bind.WatchOpts, sink chan<- *IMainChainLockProxyVerifyHeaderAndExecuteTxEvent) (event.Subscription, error) {

	logs, sub, err := _IMainChainLockProxy.contract.WatchLogs(opts, "VerifyHeaderAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMainChainLockProxyVerifyHeaderAndExecuteTxEvent)
				if err := _IMainChainLockProxy.contract.UnpackLog(event, "VerifyHeaderAndExecuteTxEvent", log); err != nil {
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

// ParseVerifyHeaderAndExecuteTxEvent is a log parse operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_IMainChainLockProxy *IMainChainLockProxyFilterer) ParseVerifyHeaderAndExecuteTxEvent(log types.Log) (*IMainChainLockProxyVerifyHeaderAndExecuteTxEvent, error) {
	event := new(IMainChainLockProxyVerifyHeaderAndExecuteTxEvent)
	if err := _IMainChainLockProxy.contract.UnpackLog(event, "VerifyHeaderAndExecuteTxEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

