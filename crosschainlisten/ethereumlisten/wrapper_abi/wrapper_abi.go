// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrapper_abi

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

// IPolyWrapperABI is the input ABI used to generate the binding from.
const IPolyWrapperABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"net\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"PolyWrapperLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"txHash\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"efee\",\"type\":\"uint256\"}],\"name\":\"PolyWrapperSpeedUp\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"txHash\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"speedUp\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IPolyWrapperFuncSigs maps the 4-byte function signature to its string representation.
var IPolyWrapperFuncSigs = map[string]string{
	"c415b95c": "feeCollector()",
	"60de1a9b": "lock(address,uint64,bytes,uint256,uint256,uint256)",
	"9d4dc021": "lockProxy()",
	"5c975abb": "paused()",
	"d3ed7c76": "speedUp(address,bytes,uint256)",
}

// IPolyWrapper is an auto generated Go binding around an Ethereum contract.
type IPolyWrapper struct {
	IPolyWrapperCaller     // Read-only binding to the contract
	IPolyWrapperTransactor // Write-only binding to the contract
	IPolyWrapperFilterer   // Log filterer for contract events
}

// IPolyWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPolyWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPolyWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPolyWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPolyWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPolyWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPolyWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPolyWrapperSession struct {
	Contract     *IPolyWrapper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPolyWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPolyWrapperCallerSession struct {
	Contract *IPolyWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IPolyWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPolyWrapperTransactorSession struct {
	Contract     *IPolyWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPolyWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPolyWrapperRaw struct {
	Contract *IPolyWrapper // Generic contract binding to access the raw methods on
}

// IPolyWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPolyWrapperCallerRaw struct {
	Contract *IPolyWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// IPolyWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPolyWrapperTransactorRaw struct {
	Contract *IPolyWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPolyWrapper creates a new instance of IPolyWrapper, bound to a specific deployed contract.
func NewIPolyWrapper(address common.Address, backend bind.ContractBackend) (*IPolyWrapper, error) {
	contract, err := bindIPolyWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPolyWrapper{IPolyWrapperCaller: IPolyWrapperCaller{contract: contract}, IPolyWrapperTransactor: IPolyWrapperTransactor{contract: contract}, IPolyWrapperFilterer: IPolyWrapperFilterer{contract: contract}}, nil
}

// NewIPolyWrapperCaller creates a new read-only instance of IPolyWrapper, bound to a specific deployed contract.
func NewIPolyWrapperCaller(address common.Address, caller bind.ContractCaller) (*IPolyWrapperCaller, error) {
	contract, err := bindIPolyWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPolyWrapperCaller{contract: contract}, nil
}

// NewIPolyWrapperTransactor creates a new write-only instance of IPolyWrapper, bound to a specific deployed contract.
func NewIPolyWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*IPolyWrapperTransactor, error) {
	contract, err := bindIPolyWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPolyWrapperTransactor{contract: contract}, nil
}

// NewIPolyWrapperFilterer creates a new log filterer instance of IPolyWrapper, bound to a specific deployed contract.
func NewIPolyWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*IPolyWrapperFilterer, error) {
	contract, err := bindIPolyWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPolyWrapperFilterer{contract: contract}, nil
}

// bindIPolyWrapper binds a generic wrapper to an already deployed contract.
func bindIPolyWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPolyWrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPolyWrapper *IPolyWrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPolyWrapper.Contract.IPolyWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPolyWrapper *IPolyWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.IPolyWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPolyWrapper *IPolyWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.IPolyWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPolyWrapper *IPolyWrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPolyWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPolyWrapper *IPolyWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPolyWrapper *IPolyWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.contract.Transact(opts, method, params...)
}

// FeeCollector is a paid mutator transaction binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() returns(address)
func (_IPolyWrapper *IPolyWrapperTransactor) FeeCollector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPolyWrapper.contract.Transact(opts, "feeCollector")
}

// FeeCollector is a paid mutator transaction binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() returns(address)
func (_IPolyWrapper *IPolyWrapperSession) FeeCollector() (*types.Transaction, error) {
	return _IPolyWrapper.Contract.FeeCollector(&_IPolyWrapper.TransactOpts)
}

// FeeCollector is a paid mutator transaction binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() returns(address)
func (_IPolyWrapper *IPolyWrapperTransactorSession) FeeCollector() (*types.Transaction, error) {
	return _IPolyWrapper.Contract.FeeCollector(&_IPolyWrapper.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) returns()
func (_IPolyWrapper *IPolyWrapperTransactor) Lock(opts *bind.TransactOpts, fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.contract.Transact(opts, "lock", fromAsset, toChainId, toAddress, amount, fee, id)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) returns()
func (_IPolyWrapper *IPolyWrapperSession) Lock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.Lock(&_IPolyWrapper.TransactOpts, fromAsset, toChainId, toAddress, amount, fee, id)
}

// Lock is a paid mutator transaction binding the contract method 0x60de1a9b.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee, uint256 id) returns()
func (_IPolyWrapper *IPolyWrapperTransactorSession) Lock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.Lock(&_IPolyWrapper.TransactOpts, fromAsset, toChainId, toAddress, amount, fee, id)
}

// LockProxy is a paid mutator transaction binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() returns(address)
func (_IPolyWrapper *IPolyWrapperTransactor) LockProxy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPolyWrapper.contract.Transact(opts, "lockProxy")
}

// LockProxy is a paid mutator transaction binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() returns(address)
func (_IPolyWrapper *IPolyWrapperSession) LockProxy() (*types.Transaction, error) {
	return _IPolyWrapper.Contract.LockProxy(&_IPolyWrapper.TransactOpts)
}

// LockProxy is a paid mutator transaction binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() returns(address)
func (_IPolyWrapper *IPolyWrapperTransactorSession) LockProxy() (*types.Transaction, error) {
	return _IPolyWrapper.Contract.LockProxy(&_IPolyWrapper.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_IPolyWrapper *IPolyWrapperTransactor) Paused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPolyWrapper.contract.Transact(opts, "paused")
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_IPolyWrapper *IPolyWrapperSession) Paused() (*types.Transaction, error) {
	return _IPolyWrapper.Contract.Paused(&_IPolyWrapper.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns(bool)
func (_IPolyWrapper *IPolyWrapperTransactorSession) Paused() (*types.Transaction, error) {
	return _IPolyWrapper.Contract.Paused(&_IPolyWrapper.TransactOpts)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) returns()
func (_IPolyWrapper *IPolyWrapperTransactor) SpeedUp(opts *bind.TransactOpts, fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.contract.Transact(opts, "speedUp", fromAsset, txHash, fee)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) returns()
func (_IPolyWrapper *IPolyWrapperSession) SpeedUp(fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.SpeedUp(&_IPolyWrapper.TransactOpts, fromAsset, txHash, fee)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) returns()
func (_IPolyWrapper *IPolyWrapperTransactorSession) SpeedUp(fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.SpeedUp(&_IPolyWrapper.TransactOpts, fromAsset, txHash, fee)
}

// IPolyWrapperPolyWrapperLockIterator is returned from FilterPolyWrapperLock and is used to iterate over the raw logs and unpacked data for PolyWrapperLock events raised by the IPolyWrapper contract.
type IPolyWrapperPolyWrapperLockIterator struct {
	Event *IPolyWrapperPolyWrapperLock // Event containing the contract specifics and raw log

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
func (it *IPolyWrapperPolyWrapperLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPolyWrapperPolyWrapperLock)
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
		it.Event = new(IPolyWrapperPolyWrapperLock)
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
func (it *IPolyWrapperPolyWrapperLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPolyWrapperPolyWrapperLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPolyWrapperPolyWrapperLock represents a PolyWrapperLock event raised by the IPolyWrapper contract.
type IPolyWrapperPolyWrapperLock struct {
	FromAsset common.Address
	Sender    common.Address
	ToChainId uint64
	ToAddress []byte
	Net       *big.Int
	Fee       *big.Int
	Id        *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPolyWrapperLock is a free log retrieval operation binding the contract event 0x2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e.
//
// Solidity: event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint256 net, uint256 fee, uint256 id)
func (_IPolyWrapper *IPolyWrapperFilterer) FilterPolyWrapperLock(opts *bind.FilterOpts, fromAsset []common.Address, sender []common.Address) (*IPolyWrapperPolyWrapperLockIterator, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPolyWrapper.contract.FilterLogs(opts, "PolyWrapperLock", fromAssetRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IPolyWrapperPolyWrapperLockIterator{contract: _IPolyWrapper.contract, event: "PolyWrapperLock", logs: logs, sub: sub}, nil
}

// WatchPolyWrapperLock is a free log subscription operation binding the contract event 0x2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e.
//
// Solidity: event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint256 net, uint256 fee, uint256 id)
func (_IPolyWrapper *IPolyWrapperFilterer) WatchPolyWrapperLock(opts *bind.WatchOpts, sink chan<- *IPolyWrapperPolyWrapperLock, fromAsset []common.Address, sender []common.Address) (event.Subscription, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPolyWrapper.contract.WatchLogs(opts, "PolyWrapperLock", fromAssetRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPolyWrapperPolyWrapperLock)
				if err := _IPolyWrapper.contract.UnpackLog(event, "PolyWrapperLock", log); err != nil {
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

// ParsePolyWrapperLock is a log parse operation binding the contract event 0x2b0591052cc6602e870d3994f0a1b173fdac98c215cb3b0baf84eaca5a0aa81e.
//
// Solidity: event PolyWrapperLock(address indexed fromAsset, address indexed sender, uint64 toChainId, bytes toAddress, uint256 net, uint256 fee, uint256 id)
func (_IPolyWrapper *IPolyWrapperFilterer) ParsePolyWrapperLock(log types.Log) (*IPolyWrapperPolyWrapperLock, error) {
	event := new(IPolyWrapperPolyWrapperLock)
	if err := _IPolyWrapper.contract.UnpackLog(event, "PolyWrapperLock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IPolyWrapperPolyWrapperSpeedUpIterator is returned from FilterPolyWrapperSpeedUp and is used to iterate over the raw logs and unpacked data for PolyWrapperSpeedUp events raised by the IPolyWrapper contract.
type IPolyWrapperPolyWrapperSpeedUpIterator struct {
	Event *IPolyWrapperPolyWrapperSpeedUp // Event containing the contract specifics and raw log

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
func (it *IPolyWrapperPolyWrapperSpeedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPolyWrapperPolyWrapperSpeedUp)
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
		it.Event = new(IPolyWrapperPolyWrapperSpeedUp)
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
func (it *IPolyWrapperPolyWrapperSpeedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPolyWrapperPolyWrapperSpeedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPolyWrapperPolyWrapperSpeedUp represents a PolyWrapperSpeedUp event raised by the IPolyWrapper contract.
type IPolyWrapperPolyWrapperSpeedUp struct {
	FromAsset common.Address
	TxHash    common.Hash
	Sender    common.Address
	Efee      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPolyWrapperSpeedUp is a free log retrieval operation binding the contract event 0xf6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef6.
//
// Solidity: event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint256 efee)
func (_IPolyWrapper *IPolyWrapperFilterer) FilterPolyWrapperSpeedUp(opts *bind.FilterOpts, fromAsset []common.Address, txHash [][]byte, sender []common.Address) (*IPolyWrapperPolyWrapperSpeedUpIterator, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPolyWrapper.contract.FilterLogs(opts, "PolyWrapperSpeedUp", fromAssetRule, txHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IPolyWrapperPolyWrapperSpeedUpIterator{contract: _IPolyWrapper.contract, event: "PolyWrapperSpeedUp", logs: logs, sub: sub}, nil
}

// WatchPolyWrapperSpeedUp is a free log subscription operation binding the contract event 0xf6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef6.
//
// Solidity: event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint256 efee)
func (_IPolyWrapper *IPolyWrapperFilterer) WatchPolyWrapperSpeedUp(opts *bind.WatchOpts, sink chan<- *IPolyWrapperPolyWrapperSpeedUp, fromAsset []common.Address, txHash [][]byte, sender []common.Address) (event.Subscription, error) {

	var fromAssetRule []interface{}
	for _, fromAssetItem := range fromAsset {
		fromAssetRule = append(fromAssetRule, fromAssetItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IPolyWrapper.contract.WatchLogs(opts, "PolyWrapperSpeedUp", fromAssetRule, txHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPolyWrapperPolyWrapperSpeedUp)
				if err := _IPolyWrapper.contract.UnpackLog(event, "PolyWrapperSpeedUp", log); err != nil {
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

// ParsePolyWrapperSpeedUp is a log parse operation binding the contract event 0xf6579aef3e0d086d986c5d6972659f8a0d8602ef7945b054be1b88e088773ef6.
//
// Solidity: event PolyWrapperSpeedUp(address indexed fromAsset, bytes indexed txHash, address indexed sender, uint256 efee)
func (_IPolyWrapper *IPolyWrapperFilterer) ParsePolyWrapperSpeedUp(log types.Log) (*IPolyWrapperPolyWrapperSpeedUp, error) {
	event := new(IPolyWrapperPolyWrapperSpeedUp)
	if err := _IPolyWrapper.contract.UnpackLog(event, "PolyWrapperSpeedUp", log); err != nil {
		return nil, err
	}
	return event, nil
}
