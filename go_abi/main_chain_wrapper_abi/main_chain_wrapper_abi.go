// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main_chain_wrapper_abi

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
	MethodLock = "lock"

	MethodLockProxy = "lockProxy"

	EventWrapperLock = "WrapperLock"
)

// MainChainWrapperTestABI is the input ABI used to generate the binding from.
const MainChainWrapperTestABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"WrapperLock\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MainChainWrapperTestFuncSigs maps the 4-byte function signature to its string representation.
var MainChainWrapperTestFuncSigs = map[string]string{
	"d8d5965b": "lock(uint64,address,uint256,uint256)",
	"9d4dc021": "lockProxy()",
}

// MainChainWrapperTestBin is the compiled bytecode used for deploying new contracts.
var MainChainWrapperTestBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916737d79d936da7833c7fe056eb450064f34a327dca81790556102a2806100466000396000f3fe6080604052600436106100295760003560e01c80639d4dc0211461002e578063d8d5965b1461005f575b600080fd5b34801561003a57600080fd5b506100436100b5565b604080516001600160a01b039092168252519081900360200190f35b6100a16004803603608081101561007557600080fd5b5067ffffffffffffffff813516906001600160a01b0360208201351690604081013590606001356100c4565b604080519115158252519081900360200190f35b6000546001600160a01b031681565b60008282013481146101075760405162461bcd60e51b815260040180806020018281038252602a815260200180610244602a913960400191505060405180910390fd5b6000805460408051634bc6882360e01b815267ffffffffffffffff8a1660048201526001600160a01b0389811660248301526044820189905291519190921692634bc6882392606480820193602093909283900390910190829087803b15801561017057600080fd5b505af1158015610184573d6000803e3d6000fd5b505050506040513d602081101561019a57600080fd5b50516101db576040805162461bcd60e51b815260206004820152600b60248201526a1b1bd8dac819985a5b195960aa1b604482015290519081900360640190fd5b604080516001600160a01b038716815267ffffffffffffffff8816602082015280820186905260608101859052905133917ff914c302b53730e780733fa61614b90ae3074af7b3d95ceef40a267e7a2f501b919081900360800190a25060019594505050505056fe6d73672e76616c75652073686f756c6420626520657175616c20746f20616d6f756e74202b2066656521a265627a7a72315820404982276692f5a3453ca85aeec2506637ea55526b05902f741c585cd74dd04b64736f6c63430005110032"

// DeployMainChainWrapperTest deploys a new Ethereum contract, binding an instance of MainChainWrapperTest to it.
func DeployMainChainWrapperTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MainChainWrapperTest, error) {
	parsed, err := abi.JSON(strings.NewReader(MainChainWrapperTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MainChainWrapperTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MainChainWrapperTest{MainChainWrapperTestCaller: MainChainWrapperTestCaller{contract: contract}, MainChainWrapperTestTransactor: MainChainWrapperTestTransactor{contract: contract}, MainChainWrapperTestFilterer: MainChainWrapperTestFilterer{contract: contract}}, nil
}

// MainChainWrapperTest is an auto generated Go binding around an Ethereum contract.
type MainChainWrapperTest struct {
	MainChainWrapperTestCaller     // Read-only binding to the contract
	MainChainWrapperTestTransactor // Write-only binding to the contract
	MainChainWrapperTestFilterer   // Log filterer for contract events
}

// MainChainWrapperTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainChainWrapperTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainChainWrapperTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainChainWrapperTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainChainWrapperTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainChainWrapperTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainChainWrapperTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainChainWrapperTestSession struct {
	Contract     *MainChainWrapperTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MainChainWrapperTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainChainWrapperTestCallerSession struct {
	Contract *MainChainWrapperTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MainChainWrapperTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainChainWrapperTestTransactorSession struct {
	Contract     *MainChainWrapperTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MainChainWrapperTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainChainWrapperTestRaw struct {
	Contract *MainChainWrapperTest // Generic contract binding to access the raw methods on
}

// MainChainWrapperTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainChainWrapperTestCallerRaw struct {
	Contract *MainChainWrapperTestCaller // Generic read-only contract binding to access the raw methods on
}

// MainChainWrapperTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainChainWrapperTestTransactorRaw struct {
	Contract *MainChainWrapperTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainChainWrapperTest creates a new instance of MainChainWrapperTest, bound to a specific deployed contract.
func NewMainChainWrapperTest(address common.Address, backend bind.ContractBackend) (*MainChainWrapperTest, error) {
	contract, err := bindMainChainWrapperTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainChainWrapperTest{MainChainWrapperTestCaller: MainChainWrapperTestCaller{contract: contract}, MainChainWrapperTestTransactor: MainChainWrapperTestTransactor{contract: contract}, MainChainWrapperTestFilterer: MainChainWrapperTestFilterer{contract: contract}}, nil
}

// NewMainChainWrapperTestCaller creates a new read-only instance of MainChainWrapperTest, bound to a specific deployed contract.
func NewMainChainWrapperTestCaller(address common.Address, caller bind.ContractCaller) (*MainChainWrapperTestCaller, error) {
	contract, err := bindMainChainWrapperTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainChainWrapperTestCaller{contract: contract}, nil
}

// NewMainChainWrapperTestTransactor creates a new write-only instance of MainChainWrapperTest, bound to a specific deployed contract.
func NewMainChainWrapperTestTransactor(address common.Address, transactor bind.ContractTransactor) (*MainChainWrapperTestTransactor, error) {
	contract, err := bindMainChainWrapperTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainChainWrapperTestTransactor{contract: contract}, nil
}

// NewMainChainWrapperTestFilterer creates a new log filterer instance of MainChainWrapperTest, bound to a specific deployed contract.
func NewMainChainWrapperTestFilterer(address common.Address, filterer bind.ContractFilterer) (*MainChainWrapperTestFilterer, error) {
	contract, err := bindMainChainWrapperTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainChainWrapperTestFilterer{contract: contract}, nil
}

// bindMainChainWrapperTest binds a generic wrapper to an already deployed contract.
func bindMainChainWrapperTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainChainWrapperTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainChainWrapperTest *MainChainWrapperTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainChainWrapperTest.Contract.MainChainWrapperTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainChainWrapperTest *MainChainWrapperTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainChainWrapperTest.Contract.MainChainWrapperTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainChainWrapperTest *MainChainWrapperTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainChainWrapperTest.Contract.MainChainWrapperTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainChainWrapperTest *MainChainWrapperTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainChainWrapperTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainChainWrapperTest *MainChainWrapperTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainChainWrapperTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainChainWrapperTest *MainChainWrapperTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainChainWrapperTest.Contract.contract.Transact(opts, method, params...)
}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_MainChainWrapperTest *MainChainWrapperTestCaller) LockProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MainChainWrapperTest.contract.Call(opts, &out, "lockProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_MainChainWrapperTest *MainChainWrapperTestSession) LockProxy() (common.Address, error) {
	return _MainChainWrapperTest.Contract.LockProxy(&_MainChainWrapperTest.CallOpts)
}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_MainChainWrapperTest *MainChainWrapperTestCallerSession) LockProxy() (common.Address, error) {
	return _MainChainWrapperTest.Contract.LockProxy(&_MainChainWrapperTest.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xd8d5965b.
//
// Solidity: function lock(uint64 toChainId, address toAddress, uint256 amount, uint256 fee) payable returns(bool)
func (_MainChainWrapperTest *MainChainWrapperTestTransactor) Lock(opts *bind.TransactOpts, toChainId uint64, toAddress common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _MainChainWrapperTest.contract.Transact(opts, "lock", toChainId, toAddress, amount, fee)
}

// Lock is a paid mutator transaction binding the contract method 0xd8d5965b.
//
// Solidity: function lock(uint64 toChainId, address toAddress, uint256 amount, uint256 fee) payable returns(bool)
func (_MainChainWrapperTest *MainChainWrapperTestSession) Lock(toChainId uint64, toAddress common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _MainChainWrapperTest.Contract.Lock(&_MainChainWrapperTest.TransactOpts, toChainId, toAddress, amount, fee)
}

// Lock is a paid mutator transaction binding the contract method 0xd8d5965b.
//
// Solidity: function lock(uint64 toChainId, address toAddress, uint256 amount, uint256 fee) payable returns(bool)
func (_MainChainWrapperTest *MainChainWrapperTestTransactorSession) Lock(toChainId uint64, toAddress common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _MainChainWrapperTest.Contract.Lock(&_MainChainWrapperTest.TransactOpts, toChainId, toAddress, amount, fee)
}

// MainChainWrapperTestWrapperLockIterator is returned from FilterWrapperLock and is used to iterate over the raw logs and unpacked data for WrapperLock events raised by the MainChainWrapperTest contract.
type MainChainWrapperTestWrapperLockIterator struct {
	Event *MainChainWrapperTestWrapperLock // Event containing the contract specifics and raw log

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
func (it *MainChainWrapperTestWrapperLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainChainWrapperTestWrapperLock)
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
		it.Event = new(MainChainWrapperTestWrapperLock)
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
func (it *MainChainWrapperTestWrapperLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainChainWrapperTestWrapperLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainChainWrapperTestWrapperLock represents a WrapperLock event raised by the MainChainWrapperTest contract.
type MainChainWrapperTestWrapperLock struct {
	Sender    common.Address
	ToAddress common.Address
	ToChainId uint64
	Amount    *big.Int
	Fee       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWrapperLock is a free log retrieval operation binding the contract event 0xf914c302b53730e780733fa61614b90ae3074af7b3d95ceef40a267e7a2f501b.
//
// Solidity: event WrapperLock(address indexed sender, address toAddress, uint64 toChainId, uint256 amount, uint256 fee)
func (_MainChainWrapperTest *MainChainWrapperTestFilterer) FilterWrapperLock(opts *bind.FilterOpts, sender []common.Address) (*MainChainWrapperTestWrapperLockIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MainChainWrapperTest.contract.FilterLogs(opts, "WrapperLock", senderRule)
	if err != nil {
		return nil, err
	}
	return &MainChainWrapperTestWrapperLockIterator{contract: _MainChainWrapperTest.contract, event: "WrapperLock", logs: logs, sub: sub}, nil
}

// WatchWrapperLock is a free log subscription operation binding the contract event 0xf914c302b53730e780733fa61614b90ae3074af7b3d95ceef40a267e7a2f501b.
//
// Solidity: event WrapperLock(address indexed sender, address toAddress, uint64 toChainId, uint256 amount, uint256 fee)
func (_MainChainWrapperTest *MainChainWrapperTestFilterer) WatchWrapperLock(opts *bind.WatchOpts, sink chan<- *MainChainWrapperTestWrapperLock, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MainChainWrapperTest.contract.WatchLogs(opts, "WrapperLock", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainChainWrapperTestWrapperLock)
				if err := _MainChainWrapperTest.contract.UnpackLog(event, "WrapperLock", log); err != nil {
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

// ParseWrapperLock is a log parse operation binding the contract event 0xf914c302b53730e780733fa61614b90ae3074af7b3d95ceef40a267e7a2f501b.
//
// Solidity: event WrapperLock(address indexed sender, address toAddress, uint64 toChainId, uint256 amount, uint256 fee)
func (_MainChainWrapperTest *MainChainWrapperTestFilterer) ParseWrapperLock(log types.Log) (*MainChainWrapperTestWrapperLock, error) {
	event := new(MainChainWrapperTestWrapperLock)
	if err := _MainChainWrapperTest.contract.UnpackLog(event, "WrapperLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
