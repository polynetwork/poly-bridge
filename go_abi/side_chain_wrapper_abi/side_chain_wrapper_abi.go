// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package side_chain_wrapper_abi

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
	MethodBurn = "burn"

	MethodLockProxy = "lockProxy"

	MethodToChainId = "toChainId"

	EventWrapperBurn = "WrapperBurn"
)

// SideChainWrapperTestABI is the input ABI used to generate the binding from.
const SideChainWrapperTestABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"WrapperBurn\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"toChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SideChainWrapperTestFuncSigs maps the 4-byte function signature to its string representation.
var SideChainWrapperTestFuncSigs = map[string]string{
	"b390c0ab": "burn(uint256,uint256)",
	"9d4dc021": "lockProxy()",
	"07bb7655": "toChainId()",
}

// SideChainWrapperTestBin is the compiled bytecode used for deploying new contracts.
var SideChainWrapperTestBin = "0x608060405234801561001057600080fd5b5060008054600160a01b600160e01b03196001600160a01b0319909116737d79d936da7833c7fe056eb450064f34a327dca81716600160a01b1790556102c08061005b6000396000f3fe6080604052600436106100345760003560e01c806307bb7655146100395780639d4dc0211461006b578063b390c0ab1461009c575b600080fd5b34801561004557600080fd5b5061004e6100d3565b6040805167ffffffffffffffff9092168252519081900360200190f35b34801561007757600080fd5b506100806100ea565b604080516001600160a01b039092168252519081900360200190f35b6100bf600480360360408110156100b257600080fd5b50803590602001356100f9565b604080519115158252519081900360200190f35b600054600160a01b900467ffffffffffffffff1681565b6000546001600160a01b031681565b600082820134811461013c5760405162461bcd60e51b815260040180806020018281038252603181526020018061025b6031913960400191505060405180910390fd5b6000805460408051631a6e5f3b60e01b815267ffffffffffffffff600160a01b84041660048201526024810188905290516001600160a01b0390921692631a6e5f3b926044808401936020939083900390910190829087803b1580156101a157600080fd5b505af11580156101b5573d6000803e3d6000fd5b505050506040513d60208110156101cb57600080fd5b5051610216576040805162461bcd60e51b81526020600482015260156024820152741b1bd8dad41c9bde1e4b989d5c9b8819985a5b1959605a1b604482015290519081900360640190fd5b6040805185815260208101859052815133927fb5c04ee83f14c08631fbf58fda41b37d2d2262481016213cbfac96596fd0b533928290030190a2506001939250505056fe74782e76616c75652073686f756c6420626520657175616c20746f2073756d206f6620616d6f756e7420616e6420666565a265627a7a723158200d5df34d2a1e5aba74ec523ec02f9713f7f46cc7a98adf6c394e07622ed835c264736f6c63430005110032"

// DeploySideChainWrapperTest deploys a new Ethereum contract, binding an instance of SideChainWrapperTest to it.
func DeploySideChainWrapperTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SideChainWrapperTest, error) {
	parsed, err := abi.JSON(strings.NewReader(SideChainWrapperTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SideChainWrapperTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SideChainWrapperTest{SideChainWrapperTestCaller: SideChainWrapperTestCaller{contract: contract}, SideChainWrapperTestTransactor: SideChainWrapperTestTransactor{contract: contract}, SideChainWrapperTestFilterer: SideChainWrapperTestFilterer{contract: contract}}, nil
}

// SideChainWrapperTest is an auto generated Go binding around an Ethereum contract.
type SideChainWrapperTest struct {
	SideChainWrapperTestCaller     // Read-only binding to the contract
	SideChainWrapperTestTransactor // Write-only binding to the contract
	SideChainWrapperTestFilterer   // Log filterer for contract events
}

// SideChainWrapperTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type SideChainWrapperTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SideChainWrapperTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SideChainWrapperTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SideChainWrapperTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SideChainWrapperTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SideChainWrapperTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SideChainWrapperTestSession struct {
	Contract     *SideChainWrapperTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SideChainWrapperTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SideChainWrapperTestCallerSession struct {
	Contract *SideChainWrapperTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SideChainWrapperTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SideChainWrapperTestTransactorSession struct {
	Contract     *SideChainWrapperTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SideChainWrapperTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type SideChainWrapperTestRaw struct {
	Contract *SideChainWrapperTest // Generic contract binding to access the raw methods on
}

// SideChainWrapperTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SideChainWrapperTestCallerRaw struct {
	Contract *SideChainWrapperTestCaller // Generic read-only contract binding to access the raw methods on
}

// SideChainWrapperTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SideChainWrapperTestTransactorRaw struct {
	Contract *SideChainWrapperTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSideChainWrapperTest creates a new instance of SideChainWrapperTest, bound to a specific deployed contract.
func NewSideChainWrapperTest(address common.Address, backend bind.ContractBackend) (*SideChainWrapperTest, error) {
	contract, err := bindSideChainWrapperTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SideChainWrapperTest{SideChainWrapperTestCaller: SideChainWrapperTestCaller{contract: contract}, SideChainWrapperTestTransactor: SideChainWrapperTestTransactor{contract: contract}, SideChainWrapperTestFilterer: SideChainWrapperTestFilterer{contract: contract}}, nil
}

// NewSideChainWrapperTestCaller creates a new read-only instance of SideChainWrapperTest, bound to a specific deployed contract.
func NewSideChainWrapperTestCaller(address common.Address, caller bind.ContractCaller) (*SideChainWrapperTestCaller, error) {
	contract, err := bindSideChainWrapperTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SideChainWrapperTestCaller{contract: contract}, nil
}

// NewSideChainWrapperTestTransactor creates a new write-only instance of SideChainWrapperTest, bound to a specific deployed contract.
func NewSideChainWrapperTestTransactor(address common.Address, transactor bind.ContractTransactor) (*SideChainWrapperTestTransactor, error) {
	contract, err := bindSideChainWrapperTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SideChainWrapperTestTransactor{contract: contract}, nil
}

// NewSideChainWrapperTestFilterer creates a new log filterer instance of SideChainWrapperTest, bound to a specific deployed contract.
func NewSideChainWrapperTestFilterer(address common.Address, filterer bind.ContractFilterer) (*SideChainWrapperTestFilterer, error) {
	contract, err := bindSideChainWrapperTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SideChainWrapperTestFilterer{contract: contract}, nil
}

// bindSideChainWrapperTest binds a generic wrapper to an already deployed contract.
func bindSideChainWrapperTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SideChainWrapperTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SideChainWrapperTest *SideChainWrapperTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SideChainWrapperTest.Contract.SideChainWrapperTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SideChainWrapperTest *SideChainWrapperTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SideChainWrapperTest.Contract.SideChainWrapperTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SideChainWrapperTest *SideChainWrapperTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SideChainWrapperTest.Contract.SideChainWrapperTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SideChainWrapperTest *SideChainWrapperTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SideChainWrapperTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SideChainWrapperTest *SideChainWrapperTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SideChainWrapperTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SideChainWrapperTest *SideChainWrapperTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SideChainWrapperTest.Contract.contract.Transact(opts, method, params...)
}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_SideChainWrapperTest *SideChainWrapperTestCaller) LockProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SideChainWrapperTest.contract.Call(opts, &out, "lockProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_SideChainWrapperTest *SideChainWrapperTestSession) LockProxy() (common.Address, error) {
	return _SideChainWrapperTest.Contract.LockProxy(&_SideChainWrapperTest.CallOpts)
}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_SideChainWrapperTest *SideChainWrapperTestCallerSession) LockProxy() (common.Address, error) {
	return _SideChainWrapperTest.Contract.LockProxy(&_SideChainWrapperTest.CallOpts)
}

// ToChainId is a free data retrieval call binding the contract method 0x07bb7655.
//
// Solidity: function toChainId() view returns(uint64)
func (_SideChainWrapperTest *SideChainWrapperTestCaller) ToChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SideChainWrapperTest.contract.Call(opts, &out, "toChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ToChainId is a free data retrieval call binding the contract method 0x07bb7655.
//
// Solidity: function toChainId() view returns(uint64)
func (_SideChainWrapperTest *SideChainWrapperTestSession) ToChainId() (uint64, error) {
	return _SideChainWrapperTest.Contract.ToChainId(&_SideChainWrapperTest.CallOpts)
}

// ToChainId is a free data retrieval call binding the contract method 0x07bb7655.
//
// Solidity: function toChainId() view returns(uint64)
func (_SideChainWrapperTest *SideChainWrapperTestCallerSession) ToChainId() (uint64, error) {
	return _SideChainWrapperTest.Contract.ToChainId(&_SideChainWrapperTest.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 amount, uint256 fee) payable returns(bool)
func (_SideChainWrapperTest *SideChainWrapperTestTransactor) Burn(opts *bind.TransactOpts, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _SideChainWrapperTest.contract.Transact(opts, "burn", amount, fee)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 amount, uint256 fee) payable returns(bool)
func (_SideChainWrapperTest *SideChainWrapperTestSession) Burn(amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _SideChainWrapperTest.Contract.Burn(&_SideChainWrapperTest.TransactOpts, amount, fee)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 amount, uint256 fee) payable returns(bool)
func (_SideChainWrapperTest *SideChainWrapperTestTransactorSession) Burn(amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _SideChainWrapperTest.Contract.Burn(&_SideChainWrapperTest.TransactOpts, amount, fee)
}

// SideChainWrapperTestWrapperBurnIterator is returned from FilterWrapperBurn and is used to iterate over the raw logs and unpacked data for WrapperBurn events raised by the SideChainWrapperTest contract.
type SideChainWrapperTestWrapperBurnIterator struct {
	Event *SideChainWrapperTestWrapperBurn // Event containing the contract specifics and raw log

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
func (it *SideChainWrapperTestWrapperBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SideChainWrapperTestWrapperBurn)
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
		it.Event = new(SideChainWrapperTestWrapperBurn)
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
func (it *SideChainWrapperTestWrapperBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SideChainWrapperTestWrapperBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SideChainWrapperTestWrapperBurn represents a WrapperBurn event raised by the SideChainWrapperTest contract.
type SideChainWrapperTestWrapperBurn struct {
	Sender common.Address
	Amount *big.Int
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWrapperBurn is a free log retrieval operation binding the contract event 0xb5c04ee83f14c08631fbf58fda41b37d2d2262481016213cbfac96596fd0b533.
//
// Solidity: event WrapperBurn(address indexed sender, uint256 amount, uint256 fee)
func (_SideChainWrapperTest *SideChainWrapperTestFilterer) FilterWrapperBurn(opts *bind.FilterOpts, sender []common.Address) (*SideChainWrapperTestWrapperBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SideChainWrapperTest.contract.FilterLogs(opts, "WrapperBurn", senderRule)
	if err != nil {
		return nil, err
	}
	return &SideChainWrapperTestWrapperBurnIterator{contract: _SideChainWrapperTest.contract, event: "WrapperBurn", logs: logs, sub: sub}, nil
}

// WatchWrapperBurn is a free log subscription operation binding the contract event 0xb5c04ee83f14c08631fbf58fda41b37d2d2262481016213cbfac96596fd0b533.
//
// Solidity: event WrapperBurn(address indexed sender, uint256 amount, uint256 fee)
func (_SideChainWrapperTest *SideChainWrapperTestFilterer) WatchWrapperBurn(opts *bind.WatchOpts, sink chan<- *SideChainWrapperTestWrapperBurn, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SideChainWrapperTest.contract.WatchLogs(opts, "WrapperBurn", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SideChainWrapperTestWrapperBurn)
				if err := _SideChainWrapperTest.contract.UnpackLog(event, "WrapperBurn", log); err != nil {
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

// ParseWrapperBurn is a log parse operation binding the contract event 0xb5c04ee83f14c08631fbf58fda41b37d2d2262481016213cbfac96596fd0b533.
//
// Solidity: event WrapperBurn(address indexed sender, uint256 amount, uint256 fee)
func (_SideChainWrapperTest *SideChainWrapperTestFilterer) ParseWrapperBurn(log types.Log) (*SideChainWrapperTestWrapperBurn, error) {
	event := new(SideChainWrapperTestWrapperBurn)
	if err := _SideChainWrapperTest.contract.UnpackLog(event, "WrapperBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
