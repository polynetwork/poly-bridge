// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swapper_abi

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

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEthCrossChainManagerABI is the input ABI used to generate the binding from.
const IEthCrossChainManagerABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_toContract\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_method\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthCrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainManagerFuncSigs = map[string]string{
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
}

// IEthCrossChainManager is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainManager struct {
	IEthCrossChainManagerCaller     // Read-only binding to the contract
	IEthCrossChainManagerTransactor // Write-only binding to the contract
	IEthCrossChainManagerFilterer   // Log filterer for contract events
}

// IEthCrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainManagerSession struct {
	Contract     *IEthCrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainManagerCallerSession struct {
	Contract *IEthCrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IEthCrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainManagerTransactorSession struct {
	Contract     *IEthCrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainManagerRaw struct {
	Contract *IEthCrossChainManager // Generic contract binding to access the raw methods on
}

// IEthCrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCallerRaw struct {
	Contract *IEthCrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactorRaw struct {
	Contract *IEthCrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainManager creates a new instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManager(address common.Address, backend bind.ContractBackend) (*IEthCrossChainManager, error) {
	contract, err := bindIEthCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManager{IEthCrossChainManagerCaller: IEthCrossChainManagerCaller{contract: contract}, IEthCrossChainManagerTransactor: IEthCrossChainManagerTransactor{contract: contract}, IEthCrossChainManagerFilterer: IEthCrossChainManagerFilterer{contract: contract}}, nil
}

// NewIEthCrossChainManagerCaller creates a new read-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainManagerCaller, error) {
	contract, err := bindIEthCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerCaller{contract: contract}, nil
}

// NewIEthCrossChainManagerTransactor creates a new write-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainManagerTransactor, error) {
	contract, err := bindIEthCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerTransactor{contract: contract}, nil
}

// NewIEthCrossChainManagerFilterer creates a new log filterer instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainManagerFilterer, error) {
	contract, err := bindIEthCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerFilterer{contract: contract}, nil
}

// bindIEthCrossChainManager binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactor) CrossChain(opts *bind.TransactOpts, _toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.contract.Transact(opts, "crossChain", _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// IEthCrossChainManagerProxyABI is the input ABI used to generate the binding from.
const IEthCrossChainManagerProxyABI = "[{\"inputs\":[],\"name\":\"getEthCrossChainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEthCrossChainManagerProxyFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainManagerProxyFuncSigs = map[string]string{
	"87939a7f": "getEthCrossChainManager()",
}

// IEthCrossChainManagerProxy is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainManagerProxy struct {
	IEthCrossChainManagerProxyCaller     // Read-only binding to the contract
	IEthCrossChainManagerProxyTransactor // Write-only binding to the contract
	IEthCrossChainManagerProxyFilterer   // Log filterer for contract events
}

// IEthCrossChainManagerProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainManagerProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainManagerProxySession struct {
	Contract     *IEthCrossChainManagerProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainManagerProxyCallerSession struct {
	Contract *IEthCrossChainManagerProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IEthCrossChainManagerProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainManagerProxyTransactorSession struct {
	Contract     *IEthCrossChainManagerProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyRaw struct {
	Contract *IEthCrossChainManagerProxy // Generic contract binding to access the raw methods on
}

// IEthCrossChainManagerProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyCallerRaw struct {
	Contract *IEthCrossChainManagerProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainManagerProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyTransactorRaw struct {
	Contract *IEthCrossChainManagerProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainManagerProxy creates a new instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxy(address common.Address, backend bind.ContractBackend) (*IEthCrossChainManagerProxy, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxy{IEthCrossChainManagerProxyCaller: IEthCrossChainManagerProxyCaller{contract: contract}, IEthCrossChainManagerProxyTransactor: IEthCrossChainManagerProxyTransactor{contract: contract}, IEthCrossChainManagerProxyFilterer: IEthCrossChainManagerProxyFilterer{contract: contract}}, nil
}

// NewIEthCrossChainManagerProxyCaller creates a new read-only instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainManagerProxyCaller, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyCaller{contract: contract}, nil
}

// NewIEthCrossChainManagerProxyTransactor creates a new write-only instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainManagerProxyTransactor, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyTransactor{contract: contract}, nil
}

// NewIEthCrossChainManagerProxyFilterer creates a new log filterer instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainManagerProxyFilterer, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyFilterer{contract: contract}, nil
}

// bindIEthCrossChainManagerProxy binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainManagerProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainManagerProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManagerProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.contract.Transact(opts, method, params...)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCaller) GetEthCrossChainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEthCrossChainManagerProxy.contract.Call(opts, &out, "getEthCrossChainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxySession) GetEthCrossChainManager() (common.Address, error) {
	return _IEthCrossChainManagerProxy.Contract.GetEthCrossChainManager(&_IEthCrossChainManagerProxy.CallOpts)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCallerSession) GetEthCrossChainManager() (common.Address, error) {
	return _IEthCrossChainManagerProxy.Contract.GetEthCrossChainManager(&_IEthCrossChainManagerProxy.CallOpts)
}

// ISwapABI is the input ABI used to generate the binding from.
const ISwapABI = "[{\"inputs\":[],\"name\":\"N_COINS\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"bool\",\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"_dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_min_token_out_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISwapFuncSigs maps the 4-byte function signature to its string representation.
var ISwapFuncSigs = map[string]string{
	"29357750": "N_COINS()",
	"503a1fc5": "add_liquidity_one_coin(uint256,address,uint256)",
	"4903b0d1": "balances(uint256)",
	"3883e119": "calc_token_amount(uint256[3],bool)",
	"cc2b27d7": "calc_withdraw_one_coin(uint256,int128)",
	"c6610657": "coins(uint256)",
	"0ed2fc95": "exchange(address,address,uint256,uint256)",
	"ddca3f43": "fee()",
	"5e0d443f": "get_dy(int128,int128,uint256)",
	"bb7b8b80": "get_virtual_price()",
	"82c63066": "lp_token()",
	"53834304": "remove_liquidity_one_coin(uint256,address,uint256)",
}

// ISwap is an auto generated Go binding around an Ethereum contract.
type ISwap struct {
	ISwapCaller     // Read-only binding to the contract
	ISwapTransactor // Write-only binding to the contract
	ISwapFilterer   // Log filterer for contract events
}

// ISwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapSession struct {
	Contract     *ISwap            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapCallerSession struct {
	Contract *ISwapCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapTransactorSession struct {
	Contract     *ISwapTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapRaw struct {
	Contract *ISwap // Generic contract binding to access the raw methods on
}

// ISwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapCallerRaw struct {
	Contract *ISwapCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapTransactorRaw struct {
	Contract *ISwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwap creates a new instance of ISwap, bound to a specific deployed contract.
func NewISwap(address common.Address, backend bind.ContractBackend) (*ISwap, error) {
	contract, err := bindISwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwap{ISwapCaller: ISwapCaller{contract: contract}, ISwapTransactor: ISwapTransactor{contract: contract}, ISwapFilterer: ISwapFilterer{contract: contract}}, nil
}

// NewISwapCaller creates a new read-only instance of ISwap, bound to a specific deployed contract.
func NewISwapCaller(address common.Address, caller bind.ContractCaller) (*ISwapCaller, error) {
	contract, err := bindISwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapCaller{contract: contract}, nil
}

// NewISwapTransactor creates a new write-only instance of ISwap, bound to a specific deployed contract.
func NewISwapTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapTransactor, error) {
	contract, err := bindISwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapTransactor{contract: contract}, nil
}

// NewISwapFilterer creates a new log filterer instance of ISwap, bound to a specific deployed contract.
func NewISwapFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapFilterer, error) {
	contract, err := bindISwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapFilterer{contract: contract}, nil
}

// bindISwap binds a generic wrapper to an already deployed contract.
func bindISwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwap *ISwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwap.Contract.ISwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwap *ISwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwap.Contract.ISwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwap *ISwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwap.Contract.ISwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwap *ISwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwap *ISwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwap *ISwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwap.Contract.contract.Transact(opts, method, params...)
}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(int128)
func (_ISwap *ISwapCaller) NCOINS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "N_COINS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(int128)
func (_ISwap *ISwapSession) NCOINS() (*big.Int, error) {
	return _ISwap.Contract.NCOINS(&_ISwap.CallOpts)
}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(int128)
func (_ISwap *ISwapCallerSession) NCOINS() (*big.Int, error) {
	return _ISwap.Contract.NCOINS(&_ISwap.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 ) view returns(uint256)
func (_ISwap *ISwapCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 ) view returns(uint256)
func (_ISwap *ISwapSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _ISwap.Contract.Balances(&_ISwap.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 ) view returns(uint256)
func (_ISwap *ISwapCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _ISwap.Contract.Balances(&_ISwap.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool _is_deposit) view returns(uint256)
func (_ISwap *ISwapCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts [3]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool _is_deposit) view returns(uint256)
func (_ISwap *ISwapSession) CalcTokenAmount(_amounts [3]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _ISwap.Contract.CalcTokenAmount(&_ISwap.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool _is_deposit) view returns(uint256)
func (_ISwap *ISwapCallerSession) CalcTokenAmount(_amounts [3]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _ISwap.Contract.CalcTokenAmount(&_ISwap.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_ISwap *ISwapCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "calc_withdraw_one_coin", amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_ISwap *ISwapSession) CalcWithdrawOneCoin(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _ISwap.Contract.CalcWithdrawOneCoin(&_ISwap.CallOpts, amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_ISwap *ISwapCallerSession) CalcWithdrawOneCoin(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _ISwap.Contract.CalcWithdrawOneCoin(&_ISwap.CallOpts, amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_ISwap *ISwapCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_ISwap *ISwapSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _ISwap.Contract.Coins(&_ISwap.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 ) view returns(address)
func (_ISwap *ISwapCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _ISwap.Contract.Coins(&_ISwap.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ISwap *ISwapCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ISwap *ISwapSession) Fee() (*big.Int, error) {
	return _ISwap.Contract.Fee(&_ISwap.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ISwap *ISwapCallerSession) Fee() (*big.Int, error) {
	return _ISwap.Contract.Fee(&_ISwap.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 _dx) view returns(uint256)
func (_ISwap *ISwapCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, _dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "get_dy", i, j, _dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 _dx) view returns(uint256)
func (_ISwap *ISwapSession) GetDy(i *big.Int, j *big.Int, _dx *big.Int) (*big.Int, error) {
	return _ISwap.Contract.GetDy(&_ISwap.CallOpts, i, j, _dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 _dx) view returns(uint256)
func (_ISwap *ISwapCallerSession) GetDy(i *big.Int, j *big.Int, _dx *big.Int) (*big.Int, error) {
	return _ISwap.Contract.GetDy(&_ISwap.CallOpts, i, j, _dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ISwap *ISwapCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ISwap *ISwapSession) GetVirtualPrice() (*big.Int, error) {
	return _ISwap.Contract.GetVirtualPrice(&_ISwap.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ISwap *ISwapCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _ISwap.Contract.GetVirtualPrice(&_ISwap.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_ISwap *ISwapCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "lp_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_ISwap *ISwapSession) LpToken() (common.Address, error) {
	return _ISwap.Contract.LpToken(&_ISwap.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_ISwap *ISwapCallerSession) LpToken() (common.Address, error) {
	return _ISwap.Contract.LpToken(&_ISwap.CallOpts)
}

// AddLiquidityOneCoin is a paid mutator transaction binding the contract method 0x503a1fc5.
//
// Solidity: function add_liquidity_one_coin(uint256 _amount, address tokenIn, uint256 _min_mint_amount) returns(uint256)
func (_ISwap *ISwapTransactor) AddLiquidityOneCoin(opts *bind.TransactOpts, _amount *big.Int, tokenIn common.Address, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "add_liquidity_one_coin", _amount, tokenIn, _min_mint_amount)
}

// AddLiquidityOneCoin is a paid mutator transaction binding the contract method 0x503a1fc5.
//
// Solidity: function add_liquidity_one_coin(uint256 _amount, address tokenIn, uint256 _min_mint_amount) returns(uint256)
func (_ISwap *ISwapSession) AddLiquidityOneCoin(_amount *big.Int, tokenIn common.Address, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.AddLiquidityOneCoin(&_ISwap.TransactOpts, _amount, tokenIn, _min_mint_amount)
}

// AddLiquidityOneCoin is a paid mutator transaction binding the contract method 0x503a1fc5.
//
// Solidity: function add_liquidity_one_coin(uint256 _amount, address tokenIn, uint256 _min_mint_amount) returns(uint256)
func (_ISwap *ISwapTransactorSession) AddLiquidityOneCoin(_amount *big.Int, tokenIn common.Address, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.AddLiquidityOneCoin(&_ISwap.TransactOpts, _amount, tokenIn, _min_mint_amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x0ed2fc95.
//
// Solidity: function exchange(address tokenIn, address tokenOut, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_ISwap *ISwapTransactor) Exchange(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "exchange", tokenIn, tokenOut, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x0ed2fc95.
//
// Solidity: function exchange(address tokenIn, address tokenOut, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_ISwap *ISwapSession) Exchange(tokenIn common.Address, tokenOut common.Address, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.Exchange(&_ISwap.TransactOpts, tokenIn, tokenOut, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x0ed2fc95.
//
// Solidity: function exchange(address tokenIn, address tokenOut, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_ISwap *ISwapTransactorSession) Exchange(tokenIn common.Address, tokenOut common.Address, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.Exchange(&_ISwap.TransactOpts, tokenIn, tokenOut, _dx, _min_dy)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x53834304.
//
// Solidity: function remove_liquidity_one_coin(uint256 _amount, address tokenOut, uint256 _min_token_out_amount) returns(uint256)
func (_ISwap *ISwapTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _amount *big.Int, tokenOut common.Address, _min_token_out_amount *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "remove_liquidity_one_coin", _amount, tokenOut, _min_token_out_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x53834304.
//
// Solidity: function remove_liquidity_one_coin(uint256 _amount, address tokenOut, uint256 _min_token_out_amount) returns(uint256)
func (_ISwap *ISwapSession) RemoveLiquidityOneCoin(_amount *big.Int, tokenOut common.Address, _min_token_out_amount *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidityOneCoin(&_ISwap.TransactOpts, _amount, tokenOut, _min_token_out_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x53834304.
//
// Solidity: function remove_liquidity_one_coin(uint256 _amount, address tokenOut, uint256 _min_token_out_amount) returns(uint256)
func (_ISwap *ISwapTransactorSession) RemoveLiquidityOneCoin(_amount *big.Int, tokenOut common.Address, _min_token_out_amount *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidityOneCoin(&_ISwap.TransactOpts, _amount, tokenOut, _min_token_out_amount)
}

// IpABI is the input ABI used to generate the binding from.
const IpABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_min_token_out_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IpFuncSigs maps the 4-byte function signature to its string representation.
var IpFuncSigs = map[string]string{
	"503a1fc5": "add_liquidity_one_coin(uint256,address,uint256)",
	"0ed2fc95": "exchange(address,address,uint256,uint256)",
	"53834304": "remove_liquidity_one_coin(uint256,address,uint256)",
}

// Ip is an auto generated Go binding around an Ethereum contract.
type Ip struct {
	IpCaller     // Read-only binding to the contract
	IpTransactor // Write-only binding to the contract
	IpFilterer   // Log filterer for contract events
}

// IpCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpSession struct {
	Contract     *Ip               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpCallerSession struct {
	Contract *IpCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpTransactorSession struct {
	Contract     *IpTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpRaw struct {
	Contract *Ip // Generic contract binding to access the raw methods on
}

// IpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpCallerRaw struct {
	Contract *IpCaller // Generic read-only contract binding to access the raw methods on
}

// IpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpTransactorRaw struct {
	Contract *IpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIp creates a new instance of Ip, bound to a specific deployed contract.
func NewIp(address common.Address, backend bind.ContractBackend) (*Ip, error) {
	contract, err := bindIp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ip{IpCaller: IpCaller{contract: contract}, IpTransactor: IpTransactor{contract: contract}, IpFilterer: IpFilterer{contract: contract}}, nil
}

// NewIpCaller creates a new read-only instance of Ip, bound to a specific deployed contract.
func NewIpCaller(address common.Address, caller bind.ContractCaller) (*IpCaller, error) {
	contract, err := bindIp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpCaller{contract: contract}, nil
}

// NewIpTransactor creates a new write-only instance of Ip, bound to a specific deployed contract.
func NewIpTransactor(address common.Address, transactor bind.ContractTransactor) (*IpTransactor, error) {
	contract, err := bindIp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpTransactor{contract: contract}, nil
}

// NewIpFilterer creates a new log filterer instance of Ip, bound to a specific deployed contract.
func NewIpFilterer(address common.Address, filterer bind.ContractFilterer) (*IpFilterer, error) {
	contract, err := bindIp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpFilterer{contract: contract}, nil
}

// bindIp binds a generic wrapper to an already deployed contract.
func bindIp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ip *IpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ip.Contract.IpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ip *IpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ip.Contract.IpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ip *IpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ip.Contract.IpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ip *IpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ip.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ip *IpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ip.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ip *IpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ip.Contract.contract.Transact(opts, method, params...)
}

// AddLiquidityOneCoin is a paid mutator transaction binding the contract method 0x503a1fc5.
//
// Solidity: function add_liquidity_one_coin(uint256 _amount, address tokenIn, uint256 _min_mint_amount) returns(uint256)
func (_Ip *IpTransactor) AddLiquidityOneCoin(opts *bind.TransactOpts, _amount *big.Int, tokenIn common.Address, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Ip.contract.Transact(opts, "add_liquidity_one_coin", _amount, tokenIn, _min_mint_amount)
}

// AddLiquidityOneCoin is a paid mutator transaction binding the contract method 0x503a1fc5.
//
// Solidity: function add_liquidity_one_coin(uint256 _amount, address tokenIn, uint256 _min_mint_amount) returns(uint256)
func (_Ip *IpSession) AddLiquidityOneCoin(_amount *big.Int, tokenIn common.Address, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Ip.Contract.AddLiquidityOneCoin(&_Ip.TransactOpts, _amount, tokenIn, _min_mint_amount)
}

// AddLiquidityOneCoin is a paid mutator transaction binding the contract method 0x503a1fc5.
//
// Solidity: function add_liquidity_one_coin(uint256 _amount, address tokenIn, uint256 _min_mint_amount) returns(uint256)
func (_Ip *IpTransactorSession) AddLiquidityOneCoin(_amount *big.Int, tokenIn common.Address, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Ip.Contract.AddLiquidityOneCoin(&_Ip.TransactOpts, _amount, tokenIn, _min_mint_amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x0ed2fc95.
//
// Solidity: function exchange(address tokenIn, address tokenOut, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Ip *IpTransactor) Exchange(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Ip.contract.Transact(opts, "exchange", tokenIn, tokenOut, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x0ed2fc95.
//
// Solidity: function exchange(address tokenIn, address tokenOut, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Ip *IpSession) Exchange(tokenIn common.Address, tokenOut common.Address, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Ip.Contract.Exchange(&_Ip.TransactOpts, tokenIn, tokenOut, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x0ed2fc95.
//
// Solidity: function exchange(address tokenIn, address tokenOut, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Ip *IpTransactorSession) Exchange(tokenIn common.Address, tokenOut common.Address, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Ip.Contract.Exchange(&_Ip.TransactOpts, tokenIn, tokenOut, _dx, _min_dy)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x53834304.
//
// Solidity: function remove_liquidity_one_coin(uint256 _amount, address tokenOut, uint256 _min_token_out_amount) returns(uint256)
func (_Ip *IpTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _amount *big.Int, tokenOut common.Address, _min_token_out_amount *big.Int) (*types.Transaction, error) {
	return _Ip.contract.Transact(opts, "remove_liquidity_one_coin", _amount, tokenOut, _min_token_out_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x53834304.
//
// Solidity: function remove_liquidity_one_coin(uint256 _amount, address tokenOut, uint256 _min_token_out_amount) returns(uint256)
func (_Ip *IpSession) RemoveLiquidityOneCoin(_amount *big.Int, tokenOut common.Address, _min_token_out_amount *big.Int) (*types.Transaction, error) {
	return _Ip.Contract.RemoveLiquidityOneCoin(&_Ip.TransactOpts, _amount, tokenOut, _min_token_out_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x53834304.
//
// Solidity: function remove_liquidity_one_coin(uint256 _amount, address tokenOut, uint256 _min_token_out_amount) returns(uint256)
func (_Ip *IpTransactorSession) RemoveLiquidityOneCoin(_amount *big.Int, tokenOut common.Address, _min_token_out_amount *big.Int) (*types.Transaction, error) {
	return _Ip.Contract.RemoveLiquidityOneCoin(&_Ip.TransactOpts, _amount, tokenOut, _min_token_out_amount)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122040df1f01086c8c0a497e2837ccd47a8aa444eb03f91ec7576750123caa860b5e64736f6c634300060c0033"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122010a0072a5651c6ef4e3ce7e975e2347044e1b908a9ae7393b25b26ec52d4b2e764736f6c634300060c0033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SwapProxyABI is the input ABI used to generate the binding from.
const SwapProxyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inAssetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outLPAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"}],\"name\":\"AddLiquidityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inLPAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outAssetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"}],\"name\":\"RemoveLiquidityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"backChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"backAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"backAddress\",\"type\":\"bytes\"}],\"name\":\"RollBackEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inAssetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outAssetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"}],\"name\":\"SwapEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"argsBs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"fromContractAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"argsBs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"fromContractAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"}],\"name\":\"addUnderlying\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"assetHashMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"assetPoolMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"}],\"name\":\"bindAssetHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"bindPoolAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawAssetHash\",\"type\":\"bytes\"}],\"name\":\"bindPoolAssetAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"targetProxyHash\",\"type\":\"bytes\"}],\"name\":\"bindProxyHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"targetSwapperHash\",\"type\":\"bytes\"}],\"name\":\"bindSwapperHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"}],\"name\":\"getBalanceFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerProxyContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"poolAddressMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"proxyHashMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"argsBs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"fromContractAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"argsBs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"fromContractAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"}],\"name\":\"removeUnderlying\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethCCMProxyAddr\",\"type\":\"address\"}],\"name\":\"setManagerProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"argsBs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"fromContractAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"argsBs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"fromContractAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"}],\"name\":\"swapUnderlying\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"swapperHashMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"argsBs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"fromContractAddr\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"}],\"name\":\"unlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SwapProxyFuncSigs maps the 4-byte function signature to its string representation.
var SwapProxyFuncSigs = map[string]string{
	"3b2ae647": "add(bytes,bytes,uint64)",
	"72abb8a5": "addUnderlying(bytes,bytes,uint64)",
	"4f7d9808": "assetHashMap(address,uint64)",
	"85dbc866": "assetPoolMap(uint64,uint64,bytes)",
	"3348f63b": "bindAssetHash(address,uint64,bytes)",
	"9a1231c8": "bindPoolAddress(uint64,address)",
	"78901796": "bindPoolAssetAddress(uint64,uint64,address,bytes)",
	"379b98f6": "bindProxyHash(uint64,bytes)",
	"9ad24ba5": "bindSwapperHash(uint64,bytes)",
	"59c589a1": "getBalanceFor(address)",
	"8f32d59b": "isOwner()",
	"84a6d055": "lock(address,uint64,bytes,uint256)",
	"d798f881": "managerProxyContract()",
	"8da5cb5b": "owner()",
	"98669474": "poolAddressMap(uint64)",
	"9e5767aa": "proxyHashMap(uint64)",
	"f072f520": "remove(bytes,bytes,uint64)",
	"f03e2fad": "removeUnderlying(bytes,bytes,uint64)",
	"715018a6": "renounceOwnership()",
	"af9980f0": "setManagerProxy(address)",
	"72c345ec": "swap(bytes,bytes,uint64)",
	"ece088b3": "swapUnderlying(bytes,bytes,uint64)",
	"db3e29f1": "swapperHashMap(uint64)",
	"f2fde38b": "transferOwnership(address)",
	"06af4b9f": "unlock(bytes,bytes,uint64)",
	"fc56058e": "update(address[],address)",
}

// SwapProxyBin is the compiled bytecode used for deploying new contracts.
var SwapProxyBin = "0x60806040523480156200001157600080fd5b5060006200001e6200006e565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35062000072565b3390565b615b6f80620000826000396000f3fe6080604052600436106101815760003560e01c80638f32d59b116100d1578063d798f8811161008a578063f03e2fad11610064578063f03e2fad14610f14578063f072f52014611055578063f2fde38b14611196578063fc56058e146111c957610181565b8063d798f88114610d8b578063db3e29f114610da0578063ece088b314610dd357610181565b80638f32d59b14610bda5780639866947414610bef5780639a1231c814610c225780639ad24ba514610c645780639e5767aa14610d25578063af9980f014610d5857610181565b8063715018a61161013e5780637890179611610118578063789017961461094157806384a6d05514610a1b57806385dbc86614610adf5780638da5cb5b14610bc557610181565b8063715018a6146106a857806372abb8a5146106bf57806372c345ec1461080057610181565b806306af4b9f146101865780633348f63b146102db578063379b98f6146103aa5780633b2ae6471461046b5780634f7d9808146105ac57806359c589a114610663575b600080fd5b34801561019257600080fd5b506102c7600480360360608110156101a957600080fd5b810190602081018135600160201b8111156101c357600080fd5b8201836020820111156101d557600080fd5b803590602001918460018302840111600160201b831117156101f657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561024857600080fd5b82018360208201111561025a57600080fd5b803590602001918460018302840111600160201b8311171561027b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b031691506112829050565b604080519115158252519081900360200190f35b3480156102e757600080fd5b506102c7600480360360608110156102fe57600080fd5b6001600160a01b03823516916001600160401b0360208201351691810190606081016040820135600160201b81111561033657600080fd5b82018360208201111561034857600080fd5b803590602001918460018302840111600160201b8311171561036957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611561945050505050565b3480156103b657600080fd5b506102c7600480360360408110156103cd57600080fd5b6001600160401b038235169190810190604081016020820135600160201b8111156103f757600080fd5b82018360208201111561040957600080fd5b803590602001918460018302840111600160201b8311171561042a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506115ef945050505050565b34801561047757600080fd5b506102c76004803603606081101561048e57600080fd5b810190602081018135600160201b8111156104a857600080fd5b8201836020820111156104ba57600080fd5b803590602001918460018302840111600160201b831117156104db57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561052d57600080fd5b82018360208201111561053f57600080fd5b803590602001918460018302840111600160201b8311171561056057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b0316915061166b9050565b3480156105b857600080fd5b506105ee600480360360408110156105cf57600080fd5b5080356001600160a01b031690602001356001600160401b031661189b565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610628578181015183820152602001610610565b50505050905090810190601f1680156106555780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561066f57600080fd5b506106966004803603602081101561068657600080fd5b50356001600160a01b031661193f565b60408051918252519081900360200190f35b3480156106b457600080fd5b506106bd6119da565b005b3480156106cb57600080fd5b506102c7600480360360608110156106e257600080fd5b810190602081018135600160201b8111156106fc57600080fd5b82018360208201111561070e57600080fd5b803590602001918460018302840111600160201b8311171561072f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561078157600080fd5b82018360208201111561079357600080fd5b803590602001918460018302840111600160201b831117156107b457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b03169150611a6b9050565b34801561080c57600080fd5b506102c76004803603606081101561082357600080fd5b810190602081018135600160201b81111561083d57600080fd5b82018360208201111561084f57600080fd5b803590602001918460018302840111600160201b8311171561087057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156108c257600080fd5b8201836020820111156108d457600080fd5b803590602001918460018302840111600160201b831117156108f557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b031691506121c39050565b34801561094d57600080fd5b506102c76004803603608081101561096457600080fd5b6001600160401b0382358116926020810135909116916001600160a01b036040830135169190810190608081016060820135600160201b8111156109a757600080fd5b8201836020820111156109b957600080fd5b803590602001918460018302840111600160201b831117156109da57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506122f7945050505050565b6102c760048036036080811015610a3157600080fd5b6001600160a01b03823516916001600160401b0360208201351691810190606081016040820135600160201b811115610a6957600080fd5b820183602082011115610a7b57600080fd5b803590602001918460018302840111600160201b83111715610a9c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250612478915050565b348015610aeb57600080fd5b50610ba960048036036060811015610b0257600080fd5b6001600160401b038235811692602081013590911691810190606081016040820135600160201b811115610b3557600080fd5b820183602082011115610b4757600080fd5b803590602001918460018302840111600160201b83111715610b6857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612765945050505050565b604080516001600160a01b039092168252519081900360200190f35b348015610bd157600080fd5b50610ba96127a6565b348015610be657600080fd5b506102c76127b5565b348015610bfb57600080fd5b50610ba960048036036020811015610c1257600080fd5b50356001600160401b03166127d9565b348015610c2e57600080fd5b506102c760048036036040811015610c4557600080fd5b5080356001600160401b031690602001356001600160a01b03166127f4565b348015610c7057600080fd5b506102c760048036036040811015610c8757600080fd5b6001600160401b038235169190810190604081016020820135600160201b811115610cb157600080fd5b820183602082011115610cc357600080fd5b803590602001918460018302840111600160201b83111715610ce457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612879945050505050565b348015610d3157600080fd5b506105ee60048036036020811015610d4857600080fd5b50356001600160401b03166128eb565b348015610d6457600080fd5b506106bd60048036036020811015610d7b57600080fd5b50356001600160a01b0316612951565b348015610d9757600080fd5b50610ba96129ba565b348015610dac57600080fd5b506105ee60048036036020811015610dc357600080fd5b50356001600160401b03166129c9565b348015610ddf57600080fd5b506102c760048036036060811015610df657600080fd5b810190602081018135600160201b811115610e1057600080fd5b820183602082011115610e2257600080fd5b803590602001918460018302840111600160201b83111715610e4357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610e9557600080fd5b820183602082011115610ea757600080fd5b803590602001918460018302840111600160201b83111715610ec857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b03169150612a319050565b348015610f2057600080fd5b506102c760048036036060811015610f3757600080fd5b810190602081018135600160201b811115610f5157600080fd5b820183602082011115610f6357600080fd5b803590602001918460018302840111600160201b83111715610f8457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610fd657600080fd5b820183602082011115610fe857600080fd5b803590602001918460018302840111600160201b8311171561100957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b0316915061317e9050565b34801561106157600080fd5b506102c76004803603606081101561107857600080fd5b810190602081018135600160201b81111561109257600080fd5b8201836020820111156110a457600080fd5b803590602001918460018302840111600160201b831117156110c557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561111757600080fd5b82018360208201111561112957600080fd5b803590602001918460018302840111600160201b8311171561114a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b031691506138439050565b3480156111a257600080fd5b506106bd600480360360208110156111b957600080fd5b50356001600160a01b0316613977565b3480156111d557600080fd5b506106bd600480360360408110156111ec57600080fd5b810190602081018135600160201b81111561120657600080fd5b82018360208201111561121857600080fd5b803590602001918460208302840111600160201b8311171561123957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550505090356001600160a01b031691506139ca9050565b600154604080516387939a7f60e01b815290516000926001600160a01b03169182916387939a7f91600480820192602092909190829003018186803b1580156112ca57600080fd5b505afa1580156112de573d6000803e3d6000fd5b505050506040513d60208110156112f457600080fd5b50516001600160a01b0316611307613b0e565b6001600160a01b03161461134c5760405162461bcd60e51b815260040180806020018281038252602d815260200180615827602d913960400191505060405180910390fd5b6113546155d2565b61135d86613b12565b90508451600014156113a05760405162461bcd60e51b815260040180806020018281038252602b815260200180615854602b913960400191505060405180910390fd5b6001600160401b03841660009081526002602052604090206113c29086613b5e565b6113fd5760405162461bcd60e51b8152600401808060200182810382526022815260200180615ad26022913960400191505060405180910390fd5b805151611451576040805162461bcd60e51b815260206004820152601b60248201527f746f4173736574486173682063616e6e6f7420626520656d7074790000000000604482015290519081900360640190fd5b60006114608260000151613c12565b9050816020015151600014156114b9576040805162461bcd60e51b8152602060048201526019602482015278746f416464726573732063616e6e6f7420626520656d70747960381b604482015290519081900360640190fd5b60006114c88360200151613c12565b90506114d982828560400151613c5c565b6115145760405162461bcd60e51b815260040180806020018281038252603c8152602001806157c1603c913960400191505060405180910390fd5b60408084015181516001600160a01b038086168252841660208201528083019190915290516000805160206156de8339815191529181900360600190a1600194505050505b509392505050565b600061156b6127b5565b6115aa576040805162461bcd60e51b815260206004820181905260248201526000805160206158a1833981519152604482015290519081900360640190fd5b6001600160a01b03841660009081526006602090815260408083206001600160401b0387168452825290912083516115e4928501906155f3565b506001949350505050565b60006115f96127b5565b611638576040805162461bcd60e51b815260206004820181905260248201526000805160206158a1833981519152604482015290519081900360640190fd5b6001600160401b03831660009081526002602090815260409091208351611661928501906155f3565b5060019392505050565b600154604080516387939a7f60e01b815290516000926001600160a01b03169182916387939a7f91600480820192602092909190829003018186803b1580156116b357600080fd5b505afa1580156116c7573d6000803e3d6000fd5b505050506040513d60208110156116dd57600080fd5b50516001600160a01b03166116f0613b0e565b6001600160a01b0316146117355760405162461bcd60e51b815260040180806020018281038252602d815260200180615827602d913960400191505060405180910390fd5b306001600160a01b03166372abb8a58686866040518463ffffffff1660e01b8152600401808060200180602001846001600160401b03168152602001838103835286818151815260200191508051906020019080838360005b838110156117a657818101518382015260200161178e565b50505050905090810190601f1680156117d35780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156118065781810151838201526020016117ee565b50505050905090810190601f1680156118335780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561185557600080fd5b505af192505050801561187a57506040513d602081101561187557600080fd5b505160015b611891576118888584613cee565b60019150611559565b5060019150611559565b60066020908152600092835260408084208252918352918190208054825160026001831615610100026000190190921691909104601f8101859004850282018501909352828152929091908301828280156119375780601f1061190c57610100808354040283529160200191611937565b820191906000526020600020905b81548152906001019060200180831161191a57829003601f168201915b505050505081565b60006001600160a01b038216611957575030316119d5565b604080516370a0823160e01b8152306004820152905183916001600160a01b038316916370a0823191602480820192602092909190829003018186803b1580156119a057600080fd5b505afa1580156119b4573d6000803e3d6000fd5b505050506040513d60208110156119ca57600080fd5b505191506119d59050565b919050565b6119e26127b5565b611a21576040805162461bcd60e51b815260206004820181905260248201526000805160206158a1833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b600030611a76613b0e565b6001600160a01b031614611abb5760405162461bcd60e51b815260040180806020018281038252604b8152602001806159bf604b913960600191505060405180910390fd5b611ac3615671565b611acc85613e40565b9050835160001415611b0f5760405162461bcd60e51b815260040180806020018281038252602581526020018061592f6025913960400191505060405180910390fd5b6001600160401b0383166000908152600360205260409020611b319085613b5e565b611b6c5760405162461bcd60e51b815260040180806020018281038252602481526020018061579d6024913960400191505060405180910390fd5b6040808201516001600160401b03166000908152600460205220546001600160a01b031680611bd6576040805162461bcd60e51b81526020600482015260116024820152701c1bdbdb08191bc81b9bdd08195e1cda5d607a1b604482015290519081900360640190fd5b60006005600084604001516001600160401b03166001600160401b031681526020019081526020016000206000866001600160401b03166001600160401b0316815260200190815260200160002083608001516040518082805190602001908083835b60208310611c585780518252601f199092019160209182019101611c39565b51815160001960209485036101000a019081169019919091161790529201948552506040519384900301909220546001600160a01b03169250505080611ce5576040805162461bcd60e51b815260206004820152601b60248201527f696e4173736574486173682063616e6e6f7420626520656d7074790000000000604482015290519081900360640190fd5b6000826001600160a01b03166382c630666040518163ffffffff1660e01b815260040160206040518083038186803b158015611d2057600080fd5b505afa158015611d34573d6000803e3d6000fd5b505050506040513d6020811015611d4a57600080fd5b505160e085015151909150611da2576040805162461bcd60e51b8152602060048201526019602482015278746f416464726573732063616e6e6f7420626520656d70747960381b604482015290519081900360640190fd5b6001600160a01b03811660009081526006602090815260408083206060888101516001600160401b03168552908352928190208054825160026001831615610100026000190190921691909104601f810185900485028201850190935282815292909190830182828015611e575780601f10611e2c57610100808354040283529160200191611e57565b820191906000526020600020905b815481529060010190602001808311611e3a57829003601f168201915b50505050509050805160001415611eb1576040805162461bcd60e51b81526020600482015260196024820152780cadae0e8f240d2d8d8cacec2d840e8de82e6e6cae890c2e6d603b1b604482015290519081900360640190fd5b6000611ec7858588600001518960200151613f00565b9050611edd86606001518760e001518484613fc3565b611ee657600080fd5b8551604080516001600160a01b038716815230602082015280820192909252516000805160206156de8339815191529181900360600190a17fa184af1adb02eb56c0f9fbbed6a596b24a1f909dc75a1a3371ce1da92ee851a0866040015185886000015188858b60600151888d60e0015160405180896001600160401b03168152602001886001600160a01b03168152602001878152602001866001600160a01b03168152602001858152602001846001600160401b031681526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611fe0578181015183820152602001611fc8565b50505050905090810190601f16801561200d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015612040578181015183820152602001612028565b50505050905090810190601f16801561206d5780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390a160008051602061597d83398151915285308860600151858a60e001518660405180876001600160a01b03168152602001866001600160a01b03168152602001856001600160401b031681526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156121135781810151838201526020016120fb565b50505050905090810190601f1680156121405780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561217357818101518382015260200161215b565b50505050905090810190601f1680156121a05780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a15060019998505050505050505050565b600154604080516387939a7f60e01b815290516000926001600160a01b03169182916387939a7f91600480820192602092909190829003018186803b15801561220b57600080fd5b505afa15801561221f573d6000803e3d6000fd5b505050506040513d602081101561223557600080fd5b50516001600160a01b0316612248613b0e565b6001600160a01b03161461228d5760405162461bcd60e51b815260040180806020018281038252602d815260200180615827602d913960400191505060405180910390fd5b60405163ece088b360e01b81526001600160401b0384166044820152606060048201908152865160648301528651309263ece088b39289928992899291829160248101916084909101906020880190808383600083156117a657818101518382015260200161178e565b60006123016127b5565b612340576040805162461bcd60e51b815260206004820181905260248201526000805160206158a1833981519152604482015290519081900360640190fd5b6001600160a01b03831660009081526006602090815260408083206001600160401b038816845290915290206123769083613b5e565b6123c7576040805162461bcd60e51b815260206004820152601860248201527f696e76616c696420636861696e2d617373657420706169720000000000000000604482015290519081900360640190fd5b6001600160401b0380861660009081526005602090815260408083209388168352928152908290209151845186939286929182918401908083835b602083106124215780518252601f199092019160209182019101612402565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922080546001600160a01b0319166001600160a01b0394909416939093179092555060019695505050505050565b6000816124c5576040805162461bcd60e51b8152602060048201526016602482015275616d6f756e742063616e6e6f74206265207a65726f2160501b604482015290519081900360640190fd5b6124cf858361432f565b61250a5760405162461bcd60e51b815260040180806020018281038252603f8152602001806158f0603f913960400191505060405180910390fd5b6001600160a01b03851660009081526006602090815260408083206001600160401b038816845282529182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156125ba5780601f1061258f576101008083540402835291602001916125ba565b820191906000526020600020905b81548152906001019060200180831161259d57829003601f168201915b50505050509050805160001415612614576040805162461bcd60e51b81526020600482015260196024820152780cadae0e8f240d2d8d8cacec2d840e8de82e6e6cae890c2e6d603b1b604482015290519081900360640190fd5b61262085858386613fc3565b61262957600080fd5b60008051602061597d83398151915286612641613b0e565b8784888860405180876001600160a01b03168152602001866001600160a01b03168152602001856001600160401b031681526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156126b95781810151838201526020016126a1565b50505050905090810190601f1680156126e65780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015612719578181015183820152602001612701565b50505050905090810190601f1680156127465780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a150600195945050505050565b60056020908152600093845260408085208252928452919092208251808401830180519281529083019390920192909220919052546001600160a01b031681565b6000546001600160a01b031690565b600080546001600160a01b03166127ca613b0e565b6001600160a01b031614905090565b6004602052600090815260409020546001600160a01b031681565b60006127fe6127b5565b61283d576040805162461bcd60e51b815260206004820181905260248201526000805160206158a1833981519152604482015290519081900360640190fd5b506001600160401b038216600090815260046020526040902080546001600160a01b0383166001600160a01b0319909116179055600192915050565b60006128836127b5565b6128c2576040805162461bcd60e51b815260206004820181905260248201526000805160206158a1833981519152604482015290519081900360640190fd5b6001600160401b03831660009081526003602090815260409091208351611661928501906155f3565b600260208181526000928352604092839020805484516001821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156119375780601f1061190c57610100808354040283529160200191611937565b6129596127b5565b612998576040805162461bcd60e51b815260206004820181905260248201526000805160206158a1833981519152604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031681565b60036020908152600091825260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156119375780601f1061190c57610100808354040283529160200191611937565b600030612a3c613b0e565b6001600160a01b031614612a815760405162461bcd60e51b815260040180806020018281038252604b8152602001806159bf604b913960600191505060405180910390fd5b612a89615671565b612a9285613e40565b9050835160001415612ad55760405162461bcd60e51b815260040180806020018281038252602581526020018061592f6025913960400191505060405180910390fd5b6001600160401b0383166000908152600360205260409020612af79085613b5e565b612b325760405162461bcd60e51b815260040180806020018281038252602481526020018061579d6024913960400191505060405180910390fd5b6040808201516001600160401b03166000908152600460205220546001600160a01b031680612b9c576040805162461bcd60e51b81526020600482015260116024820152701c1bdbdb08191bc81b9bdd08195e1cda5d607a1b604482015290519081900360640190fd5b60006005600084604001516001600160401b03166001600160401b031681526020019081526020016000206000866001600160401b03166001600160401b0316815260200190815260200160002083608001516040518082805190602001908083835b60208310612c1e5780518252601f199092019160209182019101612bff565b51815160001960209485036101000a019081169019919091161790529201948552506040519384900301909220546001600160a01b03169250505080612cab576040805162461bcd60e51b815260206004820152601b60248201527f696e4173736574486173682063616e6e6f7420626520656d7074790000000000604482015290519081900360640190fd5b60006005600085604001516001600160401b03166001600160401b03168152602001908152602001600020600085606001516001600160401b03166001600160401b031681526020019081526020016000208460c001516040518082805190602001908083835b60208310612d315780518252601f199092019160209182019101612d12565b51815160001960209485036101000a019081169019919091161790529201948552506040519384900301909220546001600160a01b03169250505080612dba576040805162461bcd60e51b81526020600482015260196024820152781d185c99d95d08185cdcd95d08191bc81b9bdd08195e1cda5d603a1b604482015290519081900360640190fd5b60e084015151612e0d576040805162461bcd60e51b8152602060048201526019602482015278746f416464726573732063616e6e6f7420626520656d70747960381b604482015290519081900360640190fd5b60c084015151612e60576040805162461bcd60e51b81526020600482015260196024820152780cadae0e8f240d2d8d8cacec2d840e8de82e6e6cae890c2e6d603b1b604482015290519081900360640190fd5b6000612e7784848760000151858960200151614452565b9050612e9185606001518660e001518760c0015184613fc3565b612e9a57600080fd5b8451604080516001600160a01b038616815230602082015280820192909252516000805160206156de8339815191529181900360600190a17f8cad61375db78f5b40b47b2bced1c95123d2b8e29bf6cefdb314b83d20af9dbb856040015184876000015185858a606001518b60c001518c60e0015160405180896001600160401b03168152602001886001600160a01b03168152602001878152602001866001600160a01b03168152602001858152602001846001600160401b031681526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015612f98578181015183820152602001612f80565b50505050905090810190601f168015612fc55780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015612ff8578181015183820152602001612fe0565b50505050905090810190601f1680156130255780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390a160008051602061597d833981519152823087606001518860c001518960e001518660405180876001600160a01b03168152602001866001600160a01b03168152602001856001600160401b031681526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156130cf5781810151838201526020016130b7565b50505050905090810190601f1680156130fc5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561312f578181015183820152602001613117565b50505050905090810190601f16801561315c5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a150600198975050505050505050565b600030613189613b0e565b6001600160a01b0316146131ce5760405162461bcd60e51b815260040180806020018281038252604b8152602001806159bf604b913960600191505060405180910390fd5b6131d6615671565b6131df85613e40565b90508351600014156132225760405162461bcd60e51b815260040180806020018281038252602581526020018061592f6025913960400191505060405180910390fd5b6001600160401b03831660009081526003602052604090206132449085613b5e565b61327f5760405162461bcd60e51b815260040180806020018281038252602481526020018061579d6024913960400191505060405180910390fd5b6040808201516001600160401b03166000908152600460205220546001600160a01b0316806132e9576040805162461bcd60e51b81526020600482015260116024820152701c1bdbdb08191bc81b9bdd08195e1cda5d607a1b604482015290519081900360640190fd5b61338e60066000836001600160a01b03166382c630666040518163ffffffff1660e01b815260040160206040518083038186803b15801561332957600080fd5b505afa15801561333d573d6000803e3d6000fd5b505050506040513d602081101561335357600080fd5b50516001600160a01b03168152602081810192909252604090810160009081206001600160401b038916825290925290206080840151613b5e565b6133c95760405162461bcd60e51b815260040180806020018281038252602a8152602001806157fd602a913960400191505060405180910390fd5b60006005600084604001516001600160401b03166001600160401b031681526020019081526020016000206000866001600160401b03166001600160401b031681526020019081526020016000208360c001516040518082805190602001908083835b6020831061344b5780518252601f19909201916020918201910161342c565b51815160001960209485036101000a019081169019919091161790529201948552506040519384900301909220546001600160a01b031692505050806134d4576040805162461bcd60e51b81526020600482015260196024820152781d185c99d95d08185cdcd95d08191bc81b9bdd08195e1cda5d603a1b604482015290519081900360640190fd5b60e083015151613527576040805162461bcd60e51b8152602060048201526019602482015278746f416464726573732063616e6e6f7420626520656d70747960381b604482015290519081900360640190fd5b600061353d838560000151848760200151614517565b905061355784606001518560e001518660c0015184613fc3565b61356057600080fd5b8351604080516001600160a01b038616815230602082015280820192909252516000805160206156de8339815191529181900360600190a17febe708b5c4cf4393d89ea503656ecc48372f1a5deeb302d22b4e219fb64fe40d8460400151848660000151858589606001518a60c001518b60e0015160405180896001600160401b03168152602001886001600160a01b03168152602001878152602001866001600160a01b03168152602001858152602001846001600160401b031681526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561365e578181015183820152602001613646565b50505050905090810190601f16801561368b5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156136be5781810151838201526020016136a6565b50505050905090810190601f1680156136eb5780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390a160008051602061597d833981519152823086606001518760c001518860e001518660405180876001600160a01b03168152602001866001600160a01b03168152602001856001600160401b031681526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561379557818101518382015260200161377d565b50505050905090810190601f1680156137c25780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156137f55781810151838201526020016137dd565b50505050905090810190601f1680156138225780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a1506001979650505050505050565b600154604080516387939a7f60e01b815290516000926001600160a01b03169182916387939a7f91600480820192602092909190829003018186803b15801561388b57600080fd5b505afa15801561389f573d6000803e3d6000fd5b505050506040513d60208110156138b557600080fd5b50516001600160a01b03166138c8613b0e565b6001600160a01b03161461390d5760405162461bcd60e51b815260040180806020018281038252602d815260200180615827602d913960400191505060405180910390fd5b60405163f03e2fad60e01b81526001600160401b0384166044820152606060048201908152865160648301528651309263f03e2fad9289928992899291829160248101916084909101906020880190808383600083156117a657818101518382015260200161178e565b61397f6127b5565b6139be576040805162461bcd60e51b815260206004820181905260248201526000805160206158a1833981519152604482015290519081900360640190fd5b6139c781614637565b50565b6139d26127b5565b613a11576040805162461bcd60e51b815260206004820181905260248201526000805160206158a1833981519152604482015290519081900360640190fd5b60005b8251811015613b0957613ac6838281518110613a2c57fe5b602002602001015183858481518110613a4157fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015613a9557600080fd5b505afa158015613aa9573d6000803e3d6000fd5b505050506040513d6020811015613abf57600080fd5b50516146d7565b613b015760405162461bcd60e51b8152600401808060200182810382526025815260200180615af46025913960400191505060405180910390fd5b600101613a14565b505050565b3390565b613b1a6155d2565b613b226155d2565b6000613b2e84826146ee565b9083529050613b3d84826146ee565b60208401919091529050613b5184826147c6565b5060408301525092915050565b60008060019050835460026001808316156101000203821604845180821460018114613b8d5760009450613c06565b8215613c06576020831060018114613beb57600189600052602060002060208a018581015b600284828410011415613be2578151835414613bd15760009950600093505b600183019250602082019150613bb2565b50505050613c04565b61010080860402945060208801518514613c0457600095505b505b50929695505050505050565b60008151601414613c545760405162461bcd60e51b81526004018080602001828103825260238152602001806157216023913960400191505060405180910390fd5b506014015190565b60006001600160a01b038416613ca8576040516001600160a01b0384169083156108fc029084906000818181858888f19350505050158015613ca2573d6000803e3d6000fd5b50611661565b613cb38484846146d7565b6116615760405162461bcd60e51b81526004018080602001828103825260338152602001806157446033913960400191505060405180910390fd5b613cf6615671565b613cff83613e40565b9050613d19828260a0015183608001518460000151613fc3565b613d2257600080fd5b7f1b01a3b7239821e3f9b220a948c8a5036776bfe981b6c7a56056668eb56b502a8282608001518360a0015160405180846001600160401b031681526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015613d9e578181015183820152602001613d86565b50505050905090810190601f168015613dcb5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015613dfe578181015183820152602001613de6565b50505050905090810190601f168015613e2b5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1505050565b613e48615671565b613e50615671565b6000613e5c84826147c6565b9083529050613e6b84826147c6565b60208401919091529050613e7f84826148c3565b6001600160401b0390911660408401529050613e9b84826148c3565b6001600160401b0390911660608401529050613eb784826146ee565b60808401919091529050613ecb84826146ee565b60a08401919091529050613edf84826146ee565b60c08401919091529050613ef384826146ee565b5060e08301525092915050565b600084613f176001600160a01b0386168284614969565b613f2b6001600160a01b0386168786614969565b6000816001600160a01b031663503a1fc58688876040518463ffffffff1660e01b815260040180848152602001836001600160a01b031681526020018281526020019350505050602060405180830381600087803b158015613f8c57600080fd5b505af1158015613fa0573d6000803e3d6000fd5b505050506040513d6020811015613fb657600080fd5b5051979650505050505050565b6000613fcd6155d2565b60405180606001604052808581526020018681526020018481525090506060613ff582614a7c565b90506000600160009054906101000a90046001600160a01b03166001600160a01b03166387939a7f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561404757600080fd5b505afa15801561405b573d6000803e3d6000fd5b505050506040513d602081101561407157600080fd5b50516001600160401b038916600090815260026020818152604092839020805484516001821615610100026000190190911693909304601f8101839004830284018301909452838352939450849360609390918301828280156141155780601f106140ea57610100808354040283529160200191614115565b820191906000526020600020905b8154815290600101906020018083116140f857829003601f168201915b50505050509050805160001415614173576040805162461bcd60e51b815260206004820152601960248201527f656d70747920696c6c6567616c20746f50726f78794861736800000000000000604482015290519081900360640190fd5b816001600160a01b031663bd5cf6258b83876040518463ffffffff1660e01b815260040180846001600160401b03168152602001806020018060200180602001848103845286818151815260200191508051906020019080838360005b838110156141e85781810151838201526020016141d0565b50505050905090810190601f1680156142155780820380516001836020036101000a031916815260200191505b508481038352600681526020018065756e6c6f636b60d01b815250602001848103825285818151815260200191508051906020019080838360005b83811015614268578181015183820152602001614250565b50505050905090810190601f1680156142955780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b1580156142b857600080fd5b505af11580156142cc573d6000803e3d6000fd5b505050506040513d60208110156142e257600080fd5b505161431f5760405162461bcd60e51b815260040180806020018281038252602f8152602001806158c1602f913960400191505060405180910390fd5b5060019998505050505050505050565b60006001600160a01b0383166143be573461437b5760405162461bcd60e51b8152600401808060200182810382526021815260200180615b196021913960400191505060405180910390fd5b8134146143b95760405162461bcd60e51b81526004018080602001828103825260298152602001806159546029913960400191505060405180910390fd5b614449565b34156143fb5760405162461bcd60e51b815260040180806020018281038252602281526020018061587f6022913960400191505060405180910390fd5b61440e83614407613b0e565b3085614ba9565b6144495760405162461bcd60e51b81526004018080602001828103825260338152602001806157446033913960400191505060405180910390fd5b50600192915050565b6000856144696001600160a01b0387168284614969565b61447d6001600160a01b0387168887614969565b60408051630ed2fc9560e01b81526001600160a01b038881166004830152868116602483015260448201889052606482018690529151600092841691630ed2fc9591608480830192602092919082900301818787803b1580156144df57600080fd5b505af11580156144f3573d6000803e3d6000fd5b505050506040513d602081101561450957600080fd5b505198975050505050505050565b600080859050614598866000836001600160a01b03166382c630666040518163ffffffff1660e01b815260040160206040518083038186803b15801561455c57600080fd5b505afa158015614570573d6000803e3d6000fd5b505050506040513d602081101561458657600080fd5b50516001600160a01b03169190614969565b6145d68686836001600160a01b03166382c630666040518163ffffffff1660e01b815260040160206040518083038186803b15801561455c57600080fd5b6000816001600160a01b031663538343048787876040518463ffffffff1660e01b815260040180848152602001836001600160a01b031681526020018281526020019350505050602060405180830381600087803b158015613f8c57600080fd5b6001600160a01b03811661467c5760405162461bcd60e51b81526004018080602001828103825260268152602001806157776026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000836115e46001600160a01b0382168585614bcd565b60606000806146fd8585614c1f565b865190955090915081850111801590614717575080840184105b6147525760405162461bcd60e51b8152600401808060200182810382526024815260200180615a786024913960400191505060405180910390fd5b60608115801561476d576040519150602082016040526147b7565b6040519150601f8316801560200281840101848101888315602002848c0101015b818310156147a657805183526020928301920161478e565b5050848452601f01601f1916604052505b509250830190505b9250929050565b600080835183602001111580156147df57508260200183105b61481a5760405162461bcd60e51b81526004018080602001828103825260238152602001806156fe6023913960400191505060405180910390fd5b600060405160206000600182038760208a0101515b8383101561484f5780821a8386015360018301925060018203915061482f565b50505081016040525190506001600160ff1b038111156148b6576040805162461bcd60e51b815260206004820152601760248201527f56616c75652065786365656473207468652072616e6765000000000000000000604482015290519081900360640190fd5b9460209390930193505050565b600080835183600801111580156148dc57508260080183105b6149175760405162461bcd60e51b8152600401808060200182810382526022815260200180615a346022913960400191505060405180910390fd5b600060405160086000600182038760208a0101515b8383101561494c5780821a8386015360018301925060018203915061492c565b505050808201604052602003900351956008949094019450505050565b8015806149ef575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156149c157600080fd5b505afa1580156149d5573d6000803e3d6000fd5b505050506040513d60208110156149eb57600080fd5b5051155b614a2a5760405162461bcd60e51b8152600401808060200182810382526036815260200180615a9c6036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052613b09908490614e38565b606080614a8c8360000151614fed565b614a998460200151614fed565b614aa685604001516150b3565b6040516020018084805190602001908083835b60208310614ad85780518252601f199092019160209182019101614ab9565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310614b205780518252601f199092019160209182019101614b01565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310614b685780518252601f199092019160209182019101614b49565b6001836020036101000a0380198251168184511680821785525050505050509050019350505050604051602081830303815290604052905080915050919050565b600084614bc16001600160a01b038216868686615150565b50600195945050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052613b09908490614e38565b6000806000614c2e85856151aa565b94509050600060fd60f81b6001600160f81b031983161415614ccc57614c548686615228565b955061ffff16905060fd8110801590614c6f575061ffff8111155b614cc0576040805162461bcd60e51b815260206004820152601f60248201527f4e65787455696e7431362c2076616c7565206f7574736964652072616e676500604482015290519081900360640190fd5b92508391506147bf9050565b607f60f91b6001600160f81b031983161415614d5c57614cec86866152b1565b955063ffffffff16905061ffff81118015614d0b575063ffffffff8111155b614cc0576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b6001600160f81b03198083161415614ddd57614d7886866148c3565b95506001600160401b0316905063ffffffff8111614cc0576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b5060f881901c60fd8110614cc0576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b614e4182615357565b614e92576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310614ed05780518252601f199092019160209182019101614eb1565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614f32576040519150601f19603f3d011682016040523d82523d6000602084013e614f37565b606091505b509150915081614f8e576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115614fe757808060200190516020811015614faa57600080fd5b5051614fe75760405162461bcd60e51b815260040180806020018281038252602a815260200180615a0a602a913960400191505060405180910390fd5b50505050565b8051606090614ffb81615393565b836040516020018083805190602001908083835b6020831061502e5780518252601f19909201916020918201910161500f565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106150765780518252601f199092019160209182019101615057565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b60606001600160ff1b03821115615111576040805162461bcd60e51b815260206004820152601b60248201527f56616c756520657863656564732075696e743235352072616e67650000000000604482015290519081900360640190fd5b60405160208082526000601f5b828210156151405785811a82602086010153600191909101906000190161511e565b5050506040818101905292915050565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052614fe7908590614e38565b600080835183600101111580156151c357508260010183105b615214576040805162461bcd60e51b815260206004820181905260248201527f4e657874427974652c204f66667365742065786365656473206d6178696d756d604482015290519081900360640190fd5b505081810160200151600182019250929050565b6000808351836002011115801561524157508260020183105b61527c5760405162461bcd60e51b815260040180806020018281038252602281526020018061599d6022913960400191505060405180910390fd5b6000604051846020870101518060011a82538060001a6001830153506002818101604052601d19909101519694019450505050565b600080835183600401111580156152ca57508260040183105b6153055760405162461bcd60e51b8152600401808060200182810382526022815260200180615a566022913960400191505060405180910390fd5b600060405160046000600182038760208a0101515b8383101561533a5780821a8386015360018301925060018203915061531a565b505050808201604052602003900351956004949094019450505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470811580159061538b5750808214155b949350505050565b606060fd826001600160401b031610156153b7576153b0826154d9565b90506119d5565b61ffff826001600160401b031611615495576153d660fd60f81b6154f5565b6153df83615509565b6040516020018083805190602001908083835b602083106154115780518252601f1990920191602091820191016153f2565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106154595780518252601f19909201916020918201910161543a565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405290506119d5565b63ffffffff826001600160401b0316116154bf576154b6607f60f91b6154f5565b6153df8361554c565b6154d06001600160f81b03196154f5565b6153df8361558f565b604080516001815260f89290921b602083015260218201905290565b60606155038260f81c6154d9565b92915050565b6040516002808252606091906000601f5b8282101561553c5785811a82602086010153600191909101906000190161551a565b5050506022810160405292915050565b6040516004808252606091906000601f5b8282101561557f5785811a82602086010153600191909101906000190161555d565b5050506024810160405292915050565b6040516008808252606091906000601f5b828210156155c25785811a8260208601015360019190910190600019016155a0565b5050506028810160405292915050565b60405180606001604052806060815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061563457805160ff1916838001178555615661565b82800160010185558215615661579182015b82811115615661578251825591602001919060010190615646565b5061566d9291506156c8565b5090565b604051806101000160405280600081526020016000815260200160006001600160401b0316815260200160006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b5b8082111561566d57600081556001016156c956fed90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df4e65787455696e743235352c206f66667365742065786365656473206d6178696d756d6279746573206c656e67746820646f6573206e6f74206d6174636820616464726573737472616e7366657220657263323020617373657420746f206c6f636b5f70726f787920636f6e7472616374206661696c6564214f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737366726f6d207377617070657220636f6e74726163742061646472657373206572726f72217472616e736665722061737365742066726f6d206c6f636b5f70726f787920636f6e747261637420746f20746f41646472657373206661696c65642166726f6d20417373657420646f206e6f74206d6174636820706f6f6c20746f6b656e20616464726573736d736753656e646572206973206e6f742045746843726f7373436861696e4d616e61676572436f6e747261637466726f6d2070726f787920636f6e747261637420616464726573732063616e6e6f7420626520656d70747974686572652073686f756c64206265206e6f206574686572207472616e73666572214f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657245746843726f7373436861696e4d616e616765722063726f7373436861696e206578656375746564206572726f72217472616e736665722061737365742066726f6d2066726f6d4164647265737320746f206c6f636b5f70726f787920636f6e747261637420206661696c65642166726f6d20636f6e747261637420616464726573732063616e6e6f7420626520656d7074797472616e73666572726564206574686572206973206e6f7420657175616c20746f20616d6f756e74218636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea54e65787455696e7431362c206f66667365742065786365656473206d6178696d756d7468697320697320616e20696e7465726e616c5f66756e6374696f6e20696e2074686520666f726d206f662065787465726e616c5f66756e6374696f6e20666f72207472792f63617463685361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565644e65787455696e7436342c206f66667365742065786365656473206d6178696d756d4e65787455696e7433322c206f66667365742065786365656473206d6178696d756d4e65787456617242797465732c206f66667365742065786365656473206d6178696d756d5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e636546726f6d2050726f787920636f6e74726163742061646472657373206572726f72217472616e7366657220617373657420746f206e6577436f6e7472616374206661696c6564217472616e736665727265642065746865722063616e6e6f74206265207a65726f21a264697066735822122054c3f21b869b053796c323dcb3ea449dc183675f1dca4a8c9a7863039e45584e64736f6c634300060c0033"

// DeploySwapProxy deploys a new Ethereum contract, binding an instance of SwapProxy to it.
func DeploySwapProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwapProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwapProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwapProxy{SwapProxyCaller: SwapProxyCaller{contract: contract}, SwapProxyTransactor: SwapProxyTransactor{contract: contract}, SwapProxyFilterer: SwapProxyFilterer{contract: contract}}, nil
}

// SwapProxy is an auto generated Go binding around an Ethereum contract.
type SwapProxy struct {
	SwapProxyCaller     // Read-only binding to the contract
	SwapProxyTransactor // Write-only binding to the contract
	SwapProxyFilterer   // Log filterer for contract events
}

// SwapProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapProxySession struct {
	Contract     *SwapProxy        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapProxyCallerSession struct {
	Contract *SwapProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SwapProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapProxyTransactorSession struct {
	Contract     *SwapProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SwapProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapProxyRaw struct {
	Contract *SwapProxy // Generic contract binding to access the raw methods on
}

// SwapProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapProxyCallerRaw struct {
	Contract *SwapProxyCaller // Generic read-only contract binding to access the raw methods on
}

// SwapProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapProxyTransactorRaw struct {
	Contract *SwapProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapProxy creates a new instance of SwapProxy, bound to a specific deployed contract.
func NewSwapProxy(address common.Address, backend bind.ContractBackend) (*SwapProxy, error) {
	contract, err := bindSwapProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapProxy{SwapProxyCaller: SwapProxyCaller{contract: contract}, SwapProxyTransactor: SwapProxyTransactor{contract: contract}, SwapProxyFilterer: SwapProxyFilterer{contract: contract}}, nil
}

// NewSwapProxyCaller creates a new read-only instance of SwapProxy, bound to a specific deployed contract.
func NewSwapProxyCaller(address common.Address, caller bind.ContractCaller) (*SwapProxyCaller, error) {
	contract, err := bindSwapProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapProxyCaller{contract: contract}, nil
}

// NewSwapProxyTransactor creates a new write-only instance of SwapProxy, bound to a specific deployed contract.
func NewSwapProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapProxyTransactor, error) {
	contract, err := bindSwapProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapProxyTransactor{contract: contract}, nil
}

// NewSwapProxyFilterer creates a new log filterer instance of SwapProxy, bound to a specific deployed contract.
func NewSwapProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapProxyFilterer, error) {
	contract, err := bindSwapProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapProxyFilterer{contract: contract}, nil
}

// bindSwapProxy binds a generic wrapper to an already deployed contract.
func bindSwapProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapProxy *SwapProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapProxy.Contract.SwapProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapProxy *SwapProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapProxy.Contract.SwapProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapProxy *SwapProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapProxy.Contract.SwapProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapProxy *SwapProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapProxy *SwapProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapProxy *SwapProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapProxy.Contract.contract.Transact(opts, method, params...)
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_SwapProxy *SwapProxyCaller) AssetHashMap(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) ([]byte, error) {
	var out []interface{}
	err := _SwapProxy.contract.Call(opts, &out, "assetHashMap", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_SwapProxy *SwapProxySession) AssetHashMap(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _SwapProxy.Contract.AssetHashMap(&_SwapProxy.CallOpts, arg0, arg1)
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_SwapProxy *SwapProxyCallerSession) AssetHashMap(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _SwapProxy.Contract.AssetHashMap(&_SwapProxy.CallOpts, arg0, arg1)
}

// AssetPoolMap is a free data retrieval call binding the contract method 0x85dbc866.
//
// Solidity: function assetPoolMap(uint64 , uint64 , bytes ) view returns(address)
func (_SwapProxy *SwapProxyCaller) AssetPoolMap(opts *bind.CallOpts, arg0 uint64, arg1 uint64, arg2 []byte) (common.Address, error) {
	var out []interface{}
	err := _SwapProxy.contract.Call(opts, &out, "assetPoolMap", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetPoolMap is a free data retrieval call binding the contract method 0x85dbc866.
//
// Solidity: function assetPoolMap(uint64 , uint64 , bytes ) view returns(address)
func (_SwapProxy *SwapProxySession) AssetPoolMap(arg0 uint64, arg1 uint64, arg2 []byte) (common.Address, error) {
	return _SwapProxy.Contract.AssetPoolMap(&_SwapProxy.CallOpts, arg0, arg1, arg2)
}

// AssetPoolMap is a free data retrieval call binding the contract method 0x85dbc866.
//
// Solidity: function assetPoolMap(uint64 , uint64 , bytes ) view returns(address)
func (_SwapProxy *SwapProxyCallerSession) AssetPoolMap(arg0 uint64, arg1 uint64, arg2 []byte) (common.Address, error) {
	return _SwapProxy.Contract.AssetPoolMap(&_SwapProxy.CallOpts, arg0, arg1, arg2)
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address fromAssetHash) view returns(uint256)
func (_SwapProxy *SwapProxyCaller) GetBalanceFor(opts *bind.CallOpts, fromAssetHash common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapProxy.contract.Call(opts, &out, "getBalanceFor", fromAssetHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address fromAssetHash) view returns(uint256)
func (_SwapProxy *SwapProxySession) GetBalanceFor(fromAssetHash common.Address) (*big.Int, error) {
	return _SwapProxy.Contract.GetBalanceFor(&_SwapProxy.CallOpts, fromAssetHash)
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address fromAssetHash) view returns(uint256)
func (_SwapProxy *SwapProxyCallerSession) GetBalanceFor(fromAssetHash common.Address) (*big.Int, error) {
	return _SwapProxy.Contract.GetBalanceFor(&_SwapProxy.CallOpts, fromAssetHash)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_SwapProxy *SwapProxyCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SwapProxy.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_SwapProxy *SwapProxySession) IsOwner() (bool, error) {
	return _SwapProxy.Contract.IsOwner(&_SwapProxy.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_SwapProxy *SwapProxyCallerSession) IsOwner() (bool, error) {
	return _SwapProxy.Contract.IsOwner(&_SwapProxy.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_SwapProxy *SwapProxyCaller) ManagerProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapProxy.contract.Call(opts, &out, "managerProxyContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_SwapProxy *SwapProxySession) ManagerProxyContract() (common.Address, error) {
	return _SwapProxy.Contract.ManagerProxyContract(&_SwapProxy.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_SwapProxy *SwapProxyCallerSession) ManagerProxyContract() (common.Address, error) {
	return _SwapProxy.Contract.ManagerProxyContract(&_SwapProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapProxy *SwapProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapProxy *SwapProxySession) Owner() (common.Address, error) {
	return _SwapProxy.Contract.Owner(&_SwapProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapProxy *SwapProxyCallerSession) Owner() (common.Address, error) {
	return _SwapProxy.Contract.Owner(&_SwapProxy.CallOpts)
}

// PoolAddressMap is a free data retrieval call binding the contract method 0x98669474.
//
// Solidity: function poolAddressMap(uint64 ) view returns(address)
func (_SwapProxy *SwapProxyCaller) PoolAddressMap(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _SwapProxy.contract.Call(opts, &out, "poolAddressMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolAddressMap is a free data retrieval call binding the contract method 0x98669474.
//
// Solidity: function poolAddressMap(uint64 ) view returns(address)
func (_SwapProxy *SwapProxySession) PoolAddressMap(arg0 uint64) (common.Address, error) {
	return _SwapProxy.Contract.PoolAddressMap(&_SwapProxy.CallOpts, arg0)
}

// PoolAddressMap is a free data retrieval call binding the contract method 0x98669474.
//
// Solidity: function poolAddressMap(uint64 ) view returns(address)
func (_SwapProxy *SwapProxyCallerSession) PoolAddressMap(arg0 uint64) (common.Address, error) {
	return _SwapProxy.Contract.PoolAddressMap(&_SwapProxy.CallOpts, arg0)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_SwapProxy *SwapProxyCaller) ProxyHashMap(opts *bind.CallOpts, arg0 uint64) ([]byte, error) {
	var out []interface{}
	err := _SwapProxy.contract.Call(opts, &out, "proxyHashMap", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_SwapProxy *SwapProxySession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _SwapProxy.Contract.ProxyHashMap(&_SwapProxy.CallOpts, arg0)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_SwapProxy *SwapProxyCallerSession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _SwapProxy.Contract.ProxyHashMap(&_SwapProxy.CallOpts, arg0)
}

// SwapperHashMap is a free data retrieval call binding the contract method 0xdb3e29f1.
//
// Solidity: function swapperHashMap(uint64 ) view returns(bytes)
func (_SwapProxy *SwapProxyCaller) SwapperHashMap(opts *bind.CallOpts, arg0 uint64) ([]byte, error) {
	var out []interface{}
	err := _SwapProxy.contract.Call(opts, &out, "swapperHashMap", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SwapperHashMap is a free data retrieval call binding the contract method 0xdb3e29f1.
//
// Solidity: function swapperHashMap(uint64 ) view returns(bytes)
func (_SwapProxy *SwapProxySession) SwapperHashMap(arg0 uint64) ([]byte, error) {
	return _SwapProxy.Contract.SwapperHashMap(&_SwapProxy.CallOpts, arg0)
}

// SwapperHashMap is a free data retrieval call binding the contract method 0xdb3e29f1.
//
// Solidity: function swapperHashMap(uint64 ) view returns(bytes)
func (_SwapProxy *SwapProxyCallerSession) SwapperHashMap(arg0 uint64) ([]byte, error) {
	return _SwapProxy.Contract.SwapperHashMap(&_SwapProxy.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0x3b2ae647.
//
// Solidity: function add(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactor) Add(opts *bind.TransactOpts, argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "add", argsBs, fromContractAddr, fromChainId)
}

// Add is a paid mutator transaction binding the contract method 0x3b2ae647.
//
// Solidity: function add(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxySession) Add(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.Add(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// Add is a paid mutator transaction binding the contract method 0x3b2ae647.
//
// Solidity: function add(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) Add(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.Add(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// AddUnderlying is a paid mutator transaction binding the contract method 0x72abb8a5.
//
// Solidity: function addUnderlying(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactor) AddUnderlying(opts *bind.TransactOpts, argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "addUnderlying", argsBs, fromContractAddr, fromChainId)
}

// AddUnderlying is a paid mutator transaction binding the contract method 0x72abb8a5.
//
// Solidity: function addUnderlying(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxySession) AddUnderlying(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.AddUnderlying(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// AddUnderlying is a paid mutator transaction binding the contract method 0x72abb8a5.
//
// Solidity: function addUnderlying(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) AddUnderlying(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.AddUnderlying(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_SwapProxy *SwapProxyTransactor) BindAssetHash(opts *bind.TransactOpts, fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "bindAssetHash", fromAssetHash, toChainId, toAssetHash)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_SwapProxy *SwapProxySession) BindAssetHash(fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _SwapProxy.Contract.BindAssetHash(&_SwapProxy.TransactOpts, fromAssetHash, toChainId, toAssetHash)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) BindAssetHash(fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _SwapProxy.Contract.BindAssetHash(&_SwapProxy.TransactOpts, fromAssetHash, toChainId, toAssetHash)
}

// BindPoolAddress is a paid mutator transaction binding the contract method 0x9a1231c8.
//
// Solidity: function bindPoolAddress(uint64 poolId, address poolAddress) returns(bool)
func (_SwapProxy *SwapProxyTransactor) BindPoolAddress(opts *bind.TransactOpts, poolId uint64, poolAddress common.Address) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "bindPoolAddress", poolId, poolAddress)
}

// BindPoolAddress is a paid mutator transaction binding the contract method 0x9a1231c8.
//
// Solidity: function bindPoolAddress(uint64 poolId, address poolAddress) returns(bool)
func (_SwapProxy *SwapProxySession) BindPoolAddress(poolId uint64, poolAddress common.Address) (*types.Transaction, error) {
	return _SwapProxy.Contract.BindPoolAddress(&_SwapProxy.TransactOpts, poolId, poolAddress)
}

// BindPoolAddress is a paid mutator transaction binding the contract method 0x9a1231c8.
//
// Solidity: function bindPoolAddress(uint64 poolId, address poolAddress) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) BindPoolAddress(poolId uint64, poolAddress common.Address) (*types.Transaction, error) {
	return _SwapProxy.Contract.BindPoolAddress(&_SwapProxy.TransactOpts, poolId, poolAddress)
}

// BindPoolAssetAddress is a paid mutator transaction binding the contract method 0x78901796.
//
// Solidity: function bindPoolAssetAddress(uint64 poolId, uint64 chainId, address assetAddress, bytes rawAssetHash) returns(bool)
func (_SwapProxy *SwapProxyTransactor) BindPoolAssetAddress(opts *bind.TransactOpts, poolId uint64, chainId uint64, assetAddress common.Address, rawAssetHash []byte) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "bindPoolAssetAddress", poolId, chainId, assetAddress, rawAssetHash)
}

// BindPoolAssetAddress is a paid mutator transaction binding the contract method 0x78901796.
//
// Solidity: function bindPoolAssetAddress(uint64 poolId, uint64 chainId, address assetAddress, bytes rawAssetHash) returns(bool)
func (_SwapProxy *SwapProxySession) BindPoolAssetAddress(poolId uint64, chainId uint64, assetAddress common.Address, rawAssetHash []byte) (*types.Transaction, error) {
	return _SwapProxy.Contract.BindPoolAssetAddress(&_SwapProxy.TransactOpts, poolId, chainId, assetAddress, rawAssetHash)
}

// BindPoolAssetAddress is a paid mutator transaction binding the contract method 0x78901796.
//
// Solidity: function bindPoolAssetAddress(uint64 poolId, uint64 chainId, address assetAddress, bytes rawAssetHash) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) BindPoolAssetAddress(poolId uint64, chainId uint64, assetAddress common.Address, rawAssetHash []byte) (*types.Transaction, error) {
	return _SwapProxy.Contract.BindPoolAssetAddress(&_SwapProxy.TransactOpts, poolId, chainId, assetAddress, rawAssetHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_SwapProxy *SwapProxyTransactor) BindProxyHash(opts *bind.TransactOpts, toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "bindProxyHash", toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_SwapProxy *SwapProxySession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _SwapProxy.Contract.BindProxyHash(&_SwapProxy.TransactOpts, toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _SwapProxy.Contract.BindProxyHash(&_SwapProxy.TransactOpts, toChainId, targetProxyHash)
}

// BindSwapperHash is a paid mutator transaction binding the contract method 0x9ad24ba5.
//
// Solidity: function bindSwapperHash(uint64 toChainId, bytes targetSwapperHash) returns(bool)
func (_SwapProxy *SwapProxyTransactor) BindSwapperHash(opts *bind.TransactOpts, toChainId uint64, targetSwapperHash []byte) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "bindSwapperHash", toChainId, targetSwapperHash)
}

// BindSwapperHash is a paid mutator transaction binding the contract method 0x9ad24ba5.
//
// Solidity: function bindSwapperHash(uint64 toChainId, bytes targetSwapperHash) returns(bool)
func (_SwapProxy *SwapProxySession) BindSwapperHash(toChainId uint64, targetSwapperHash []byte) (*types.Transaction, error) {
	return _SwapProxy.Contract.BindSwapperHash(&_SwapProxy.TransactOpts, toChainId, targetSwapperHash)
}

// BindSwapperHash is a paid mutator transaction binding the contract method 0x9ad24ba5.
//
// Solidity: function bindSwapperHash(uint64 toChainId, bytes targetSwapperHash) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) BindSwapperHash(toChainId uint64, targetSwapperHash []byte) (*types.Transaction, error) {
	return _SwapProxy.Contract.BindSwapperHash(&_SwapProxy.TransactOpts, toChainId, targetSwapperHash)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_SwapProxy *SwapProxyTransactor) Lock(opts *bind.TransactOpts, fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "lock", fromAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_SwapProxy *SwapProxySession) Lock(fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _SwapProxy.Contract.Lock(&_SwapProxy.TransactOpts, fromAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) Lock(fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _SwapProxy.Contract.Lock(&_SwapProxy.TransactOpts, fromAssetHash, toChainId, toAddress, amount)
}

// Remove is a paid mutator transaction binding the contract method 0xf072f520.
//
// Solidity: function remove(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactor) Remove(opts *bind.TransactOpts, argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "remove", argsBs, fromContractAddr, fromChainId)
}

// Remove is a paid mutator transaction binding the contract method 0xf072f520.
//
// Solidity: function remove(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxySession) Remove(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.Remove(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// Remove is a paid mutator transaction binding the contract method 0xf072f520.
//
// Solidity: function remove(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) Remove(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.Remove(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// RemoveUnderlying is a paid mutator transaction binding the contract method 0xf03e2fad.
//
// Solidity: function removeUnderlying(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactor) RemoveUnderlying(opts *bind.TransactOpts, argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "removeUnderlying", argsBs, fromContractAddr, fromChainId)
}

// RemoveUnderlying is a paid mutator transaction binding the contract method 0xf03e2fad.
//
// Solidity: function removeUnderlying(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxySession) RemoveUnderlying(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.RemoveUnderlying(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// RemoveUnderlying is a paid mutator transaction binding the contract method 0xf03e2fad.
//
// Solidity: function removeUnderlying(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) RemoveUnderlying(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.RemoveUnderlying(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapProxy *SwapProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapProxy *SwapProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapProxy.Contract.RenounceOwnership(&_SwapProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapProxy *SwapProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapProxy.Contract.RenounceOwnership(&_SwapProxy.TransactOpts)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address ethCCMProxyAddr) returns()
func (_SwapProxy *SwapProxyTransactor) SetManagerProxy(opts *bind.TransactOpts, ethCCMProxyAddr common.Address) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "setManagerProxy", ethCCMProxyAddr)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address ethCCMProxyAddr) returns()
func (_SwapProxy *SwapProxySession) SetManagerProxy(ethCCMProxyAddr common.Address) (*types.Transaction, error) {
	return _SwapProxy.Contract.SetManagerProxy(&_SwapProxy.TransactOpts, ethCCMProxyAddr)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address ethCCMProxyAddr) returns()
func (_SwapProxy *SwapProxyTransactorSession) SetManagerProxy(ethCCMProxyAddr common.Address) (*types.Transaction, error) {
	return _SwapProxy.Contract.SetManagerProxy(&_SwapProxy.TransactOpts, ethCCMProxyAddr)
}

// Swap is a paid mutator transaction binding the contract method 0x72c345ec.
//
// Solidity: function swap(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactor) Swap(opts *bind.TransactOpts, argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "swap", argsBs, fromContractAddr, fromChainId)
}

// Swap is a paid mutator transaction binding the contract method 0x72c345ec.
//
// Solidity: function swap(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxySession) Swap(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.Swap(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// Swap is a paid mutator transaction binding the contract method 0x72c345ec.
//
// Solidity: function swap(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) Swap(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.Swap(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// SwapUnderlying is a paid mutator transaction binding the contract method 0xece088b3.
//
// Solidity: function swapUnderlying(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactor) SwapUnderlying(opts *bind.TransactOpts, argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "swapUnderlying", argsBs, fromContractAddr, fromChainId)
}

// SwapUnderlying is a paid mutator transaction binding the contract method 0xece088b3.
//
// Solidity: function swapUnderlying(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxySession) SwapUnderlying(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.SwapUnderlying(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// SwapUnderlying is a paid mutator transaction binding the contract method 0xece088b3.
//
// Solidity: function swapUnderlying(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) SwapUnderlying(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.SwapUnderlying(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapProxy *SwapProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapProxy *SwapProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapProxy.Contract.TransferOwnership(&_SwapProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapProxy *SwapProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapProxy.Contract.TransferOwnership(&_SwapProxy.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x06af4b9f.
//
// Solidity: function unlock(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactor) Unlock(opts *bind.TransactOpts, argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "unlock", argsBs, fromContractAddr, fromChainId)
}

// Unlock is a paid mutator transaction binding the contract method 0x06af4b9f.
//
// Solidity: function unlock(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxySession) Unlock(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.Unlock(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// Unlock is a paid mutator transaction binding the contract method 0x06af4b9f.
//
// Solidity: function unlock(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_SwapProxy *SwapProxyTransactorSession) Unlock(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _SwapProxy.Contract.Unlock(&_SwapProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// Update is a paid mutator transaction binding the contract method 0xfc56058e.
//
// Solidity: function update(address[] tokens, address newContract) returns()
func (_SwapProxy *SwapProxyTransactor) Update(opts *bind.TransactOpts, tokens []common.Address, newContract common.Address) (*types.Transaction, error) {
	return _SwapProxy.contract.Transact(opts, "update", tokens, newContract)
}

// Update is a paid mutator transaction binding the contract method 0xfc56058e.
//
// Solidity: function update(address[] tokens, address newContract) returns()
func (_SwapProxy *SwapProxySession) Update(tokens []common.Address, newContract common.Address) (*types.Transaction, error) {
	return _SwapProxy.Contract.Update(&_SwapProxy.TransactOpts, tokens, newContract)
}

// Update is a paid mutator transaction binding the contract method 0xfc56058e.
//
// Solidity: function update(address[] tokens, address newContract) returns()
func (_SwapProxy *SwapProxyTransactorSession) Update(tokens []common.Address, newContract common.Address) (*types.Transaction, error) {
	return _SwapProxy.Contract.Update(&_SwapProxy.TransactOpts, tokens, newContract)
}

// SwapProxyAddLiquidityEventIterator is returned from FilterAddLiquidityEvent and is used to iterate over the raw logs and unpacked data for AddLiquidityEvent events raised by the SwapProxy contract.
type SwapProxyAddLiquidityEventIterator struct {
	Event *SwapProxyAddLiquidityEvent // Event containing the contract specifics and raw log

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
func (it *SwapProxyAddLiquidityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapProxyAddLiquidityEvent)
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
		it.Event = new(SwapProxyAddLiquidityEvent)
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
func (it *SwapProxyAddLiquidityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapProxyAddLiquidityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapProxyAddLiquidityEvent represents a AddLiquidityEvent event raised by the SwapProxy contract.
type SwapProxyAddLiquidityEvent struct {
	ToPoolId         uint64
	InAssetAddress   common.Address
	InAmount         *big.Int
	PoolTokenAddress common.Address
	OutLPAmount      *big.Int
	ToChainId        uint64
	ToAssetHash      []byte
	ToAddress        []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidityEvent is a free log retrieval operation binding the contract event 0xa184af1adb02eb56c0f9fbbed6a596b24a1f909dc75a1a3371ce1da92ee851a0.
//
// Solidity: event AddLiquidityEvent(uint64 toPoolId, address inAssetAddress, uint256 inAmount, address poolTokenAddress, uint256 outLPAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress)
func (_SwapProxy *SwapProxyFilterer) FilterAddLiquidityEvent(opts *bind.FilterOpts) (*SwapProxyAddLiquidityEventIterator, error) {

	logs, sub, err := _SwapProxy.contract.FilterLogs(opts, "AddLiquidityEvent")
	if err != nil {
		return nil, err
	}
	return &SwapProxyAddLiquidityEventIterator{contract: _SwapProxy.contract, event: "AddLiquidityEvent", logs: logs, sub: sub}, nil
}

// WatchAddLiquidityEvent is a free log subscription operation binding the contract event 0xa184af1adb02eb56c0f9fbbed6a596b24a1f909dc75a1a3371ce1da92ee851a0.
//
// Solidity: event AddLiquidityEvent(uint64 toPoolId, address inAssetAddress, uint256 inAmount, address poolTokenAddress, uint256 outLPAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress)
func (_SwapProxy *SwapProxyFilterer) WatchAddLiquidityEvent(opts *bind.WatchOpts, sink chan<- *SwapProxyAddLiquidityEvent) (event.Subscription, error) {

	logs, sub, err := _SwapProxy.contract.WatchLogs(opts, "AddLiquidityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapProxyAddLiquidityEvent)
				if err := _SwapProxy.contract.UnpackLog(event, "AddLiquidityEvent", log); err != nil {
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

// ParseAddLiquidityEvent is a log parse operation binding the contract event 0xa184af1adb02eb56c0f9fbbed6a596b24a1f909dc75a1a3371ce1da92ee851a0.
//
// Solidity: event AddLiquidityEvent(uint64 toPoolId, address inAssetAddress, uint256 inAmount, address poolTokenAddress, uint256 outLPAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress)
func (_SwapProxy *SwapProxyFilterer) ParseAddLiquidityEvent(log types.Log) (*SwapProxyAddLiquidityEvent, error) {
	event := new(SwapProxyAddLiquidityEvent)
	if err := _SwapProxy.contract.UnpackLog(event, "AddLiquidityEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapProxyLockEventIterator is returned from FilterLockEvent and is used to iterate over the raw logs and unpacked data for LockEvent events raised by the SwapProxy contract.
type SwapProxyLockEventIterator struct {
	Event *SwapProxyLockEvent // Event containing the contract specifics and raw log

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
func (it *SwapProxyLockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapProxyLockEvent)
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
		it.Event = new(SwapProxyLockEvent)
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
func (it *SwapProxyLockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapProxyLockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapProxyLockEvent represents a LockEvent event raised by the SwapProxy contract.
type SwapProxyLockEvent struct {
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
func (_SwapProxy *SwapProxyFilterer) FilterLockEvent(opts *bind.FilterOpts) (*SwapProxyLockEventIterator, error) {

	logs, sub, err := _SwapProxy.contract.FilterLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return &SwapProxyLockEventIterator{contract: _SwapProxy.contract, event: "LockEvent", logs: logs, sub: sub}, nil
}

// WatchLockEvent is a free log subscription operation binding the contract event 0x8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea5.
//
// Solidity: event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount)
func (_SwapProxy *SwapProxyFilterer) WatchLockEvent(opts *bind.WatchOpts, sink chan<- *SwapProxyLockEvent) (event.Subscription, error) {

	logs, sub, err := _SwapProxy.contract.WatchLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapProxyLockEvent)
				if err := _SwapProxy.contract.UnpackLog(event, "LockEvent", log); err != nil {
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
func (_SwapProxy *SwapProxyFilterer) ParseLockEvent(log types.Log) (*SwapProxyLockEvent, error) {
	event := new(SwapProxyLockEvent)
	if err := _SwapProxy.contract.UnpackLog(event, "LockEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SwapProxy contract.
type SwapProxyOwnershipTransferredIterator struct {
	Event *SwapProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwapProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapProxyOwnershipTransferred)
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
		it.Event = new(SwapProxyOwnershipTransferred)
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
func (it *SwapProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapProxyOwnershipTransferred represents a OwnershipTransferred event raised by the SwapProxy contract.
type SwapProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapProxy *SwapProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapProxyOwnershipTransferredIterator{contract: _SwapProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapProxy *SwapProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapProxyOwnershipTransferred)
				if err := _SwapProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SwapProxy *SwapProxyFilterer) ParseOwnershipTransferred(log types.Log) (*SwapProxyOwnershipTransferred, error) {
	event := new(SwapProxyOwnershipTransferred)
	if err := _SwapProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapProxyRemoveLiquidityEventIterator is returned from FilterRemoveLiquidityEvent and is used to iterate over the raw logs and unpacked data for RemoveLiquidityEvent events raised by the SwapProxy contract.
type SwapProxyRemoveLiquidityEventIterator struct {
	Event *SwapProxyRemoveLiquidityEvent // Event containing the contract specifics and raw log

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
func (it *SwapProxyRemoveLiquidityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapProxyRemoveLiquidityEvent)
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
		it.Event = new(SwapProxyRemoveLiquidityEvent)
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
func (it *SwapProxyRemoveLiquidityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapProxyRemoveLiquidityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapProxyRemoveLiquidityEvent represents a RemoveLiquidityEvent event raised by the SwapProxy contract.
type SwapProxyRemoveLiquidityEvent struct {
	ToPoolId         uint64
	PoolTokenAddress common.Address
	InLPAmount       *big.Int
	OutAssetAddress  common.Address
	OutAmount        *big.Int
	ToChainId        uint64
	ToAssetHash      []byte
	ToAddress        []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityEvent is a free log retrieval operation binding the contract event 0xebe708b5c4cf4393d89ea503656ecc48372f1a5deeb302d22b4e219fb64fe40d.
//
// Solidity: event RemoveLiquidityEvent(uint64 toPoolId, address poolTokenAddress, uint256 inLPAmount, address outAssetAddress, uint256 outAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress)
func (_SwapProxy *SwapProxyFilterer) FilterRemoveLiquidityEvent(opts *bind.FilterOpts) (*SwapProxyRemoveLiquidityEventIterator, error) {

	logs, sub, err := _SwapProxy.contract.FilterLogs(opts, "RemoveLiquidityEvent")
	if err != nil {
		return nil, err
	}
	return &SwapProxyRemoveLiquidityEventIterator{contract: _SwapProxy.contract, event: "RemoveLiquidityEvent", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityEvent is a free log subscription operation binding the contract event 0xebe708b5c4cf4393d89ea503656ecc48372f1a5deeb302d22b4e219fb64fe40d.
//
// Solidity: event RemoveLiquidityEvent(uint64 toPoolId, address poolTokenAddress, uint256 inLPAmount, address outAssetAddress, uint256 outAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress)
func (_SwapProxy *SwapProxyFilterer) WatchRemoveLiquidityEvent(opts *bind.WatchOpts, sink chan<- *SwapProxyRemoveLiquidityEvent) (event.Subscription, error) {

	logs, sub, err := _SwapProxy.contract.WatchLogs(opts, "RemoveLiquidityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapProxyRemoveLiquidityEvent)
				if err := _SwapProxy.contract.UnpackLog(event, "RemoveLiquidityEvent", log); err != nil {
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

// ParseRemoveLiquidityEvent is a log parse operation binding the contract event 0xebe708b5c4cf4393d89ea503656ecc48372f1a5deeb302d22b4e219fb64fe40d.
//
// Solidity: event RemoveLiquidityEvent(uint64 toPoolId, address poolTokenAddress, uint256 inLPAmount, address outAssetAddress, uint256 outAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress)
func (_SwapProxy *SwapProxyFilterer) ParseRemoveLiquidityEvent(log types.Log) (*SwapProxyRemoveLiquidityEvent, error) {
	event := new(SwapProxyRemoveLiquidityEvent)
	if err := _SwapProxy.contract.UnpackLog(event, "RemoveLiquidityEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapProxyRollBackEventIterator is returned from FilterRollBackEvent and is used to iterate over the raw logs and unpacked data for RollBackEvent events raised by the SwapProxy contract.
type SwapProxyRollBackEventIterator struct {
	Event *SwapProxyRollBackEvent // Event containing the contract specifics and raw log

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
func (it *SwapProxyRollBackEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapProxyRollBackEvent)
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
		it.Event = new(SwapProxyRollBackEvent)
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
func (it *SwapProxyRollBackEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapProxyRollBackEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapProxyRollBackEvent represents a RollBackEvent event raised by the SwapProxy contract.
type SwapProxyRollBackEvent struct {
	BackChainId   uint64
	BackAssetHash []byte
	BackAddress   []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollBackEvent is a free log retrieval operation binding the contract event 0x1b01a3b7239821e3f9b220a948c8a5036776bfe981b6c7a56056668eb56b502a.
//
// Solidity: event RollBackEvent(uint64 backChainId, bytes backAssetHash, bytes backAddress)
func (_SwapProxy *SwapProxyFilterer) FilterRollBackEvent(opts *bind.FilterOpts) (*SwapProxyRollBackEventIterator, error) {

	logs, sub, err := _SwapProxy.contract.FilterLogs(opts, "RollBackEvent")
	if err != nil {
		return nil, err
	}
	return &SwapProxyRollBackEventIterator{contract: _SwapProxy.contract, event: "RollBackEvent", logs: logs, sub: sub}, nil
}

// WatchRollBackEvent is a free log subscription operation binding the contract event 0x1b01a3b7239821e3f9b220a948c8a5036776bfe981b6c7a56056668eb56b502a.
//
// Solidity: event RollBackEvent(uint64 backChainId, bytes backAssetHash, bytes backAddress)
func (_SwapProxy *SwapProxyFilterer) WatchRollBackEvent(opts *bind.WatchOpts, sink chan<- *SwapProxyRollBackEvent) (event.Subscription, error) {

	logs, sub, err := _SwapProxy.contract.WatchLogs(opts, "RollBackEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapProxyRollBackEvent)
				if err := _SwapProxy.contract.UnpackLog(event, "RollBackEvent", log); err != nil {
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

// ParseRollBackEvent is a log parse operation binding the contract event 0x1b01a3b7239821e3f9b220a948c8a5036776bfe981b6c7a56056668eb56b502a.
//
// Solidity: event RollBackEvent(uint64 backChainId, bytes backAssetHash, bytes backAddress)
func (_SwapProxy *SwapProxyFilterer) ParseRollBackEvent(log types.Log) (*SwapProxyRollBackEvent, error) {
	event := new(SwapProxyRollBackEvent)
	if err := _SwapProxy.contract.UnpackLog(event, "RollBackEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapProxySwapEventIterator is returned from FilterSwapEvent and is used to iterate over the raw logs and unpacked data for SwapEvent events raised by the SwapProxy contract.
type SwapProxySwapEventIterator struct {
	Event *SwapProxySwapEvent // Event containing the contract specifics and raw log

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
func (it *SwapProxySwapEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapProxySwapEvent)
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
		it.Event = new(SwapProxySwapEvent)
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
func (it *SwapProxySwapEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapProxySwapEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapProxySwapEvent represents a SwapEvent event raised by the SwapProxy contract.
type SwapProxySwapEvent struct {
	ToPoolId        uint64
	InAssetAddress  common.Address
	InAmount        *big.Int
	OutAssetAddress common.Address
	OutAmount       *big.Int
	ToChainId       uint64
	ToAssetHash     []byte
	ToAddress       []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSwapEvent is a free log retrieval operation binding the contract event 0x8cad61375db78f5b40b47b2bced1c95123d2b8e29bf6cefdb314b83d20af9dbb.
//
// Solidity: event SwapEvent(uint64 toPoolId, address inAssetAddress, uint256 inAmount, address outAssetAddress, uint256 outAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress)
func (_SwapProxy *SwapProxyFilterer) FilterSwapEvent(opts *bind.FilterOpts) (*SwapProxySwapEventIterator, error) {

	logs, sub, err := _SwapProxy.contract.FilterLogs(opts, "SwapEvent")
	if err != nil {
		return nil, err
	}
	return &SwapProxySwapEventIterator{contract: _SwapProxy.contract, event: "SwapEvent", logs: logs, sub: sub}, nil
}

// WatchSwapEvent is a free log subscription operation binding the contract event 0x8cad61375db78f5b40b47b2bced1c95123d2b8e29bf6cefdb314b83d20af9dbb.
//
// Solidity: event SwapEvent(uint64 toPoolId, address inAssetAddress, uint256 inAmount, address outAssetAddress, uint256 outAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress)
func (_SwapProxy *SwapProxyFilterer) WatchSwapEvent(opts *bind.WatchOpts, sink chan<- *SwapProxySwapEvent) (event.Subscription, error) {

	logs, sub, err := _SwapProxy.contract.WatchLogs(opts, "SwapEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapProxySwapEvent)
				if err := _SwapProxy.contract.UnpackLog(event, "SwapEvent", log); err != nil {
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

// ParseSwapEvent is a log parse operation binding the contract event 0x8cad61375db78f5b40b47b2bced1c95123d2b8e29bf6cefdb314b83d20af9dbb.
//
// Solidity: event SwapEvent(uint64 toPoolId, address inAssetAddress, uint256 inAmount, address outAssetAddress, uint256 outAmount, uint64 toChainId, bytes toAssetHash, bytes toAddress)
func (_SwapProxy *SwapProxyFilterer) ParseSwapEvent(log types.Log) (*SwapProxySwapEvent, error) {
	event := new(SwapProxySwapEvent)
	if err := _SwapProxy.contract.UnpackLog(event, "SwapEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapProxyUnlockEventIterator is returned from FilterUnlockEvent and is used to iterate over the raw logs and unpacked data for UnlockEvent events raised by the SwapProxy contract.
type SwapProxyUnlockEventIterator struct {
	Event *SwapProxyUnlockEvent // Event containing the contract specifics and raw log

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
func (it *SwapProxyUnlockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapProxyUnlockEvent)
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
		it.Event = new(SwapProxyUnlockEvent)
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
func (it *SwapProxyUnlockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapProxyUnlockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapProxyUnlockEvent represents a UnlockEvent event raised by the SwapProxy contract.
type SwapProxyUnlockEvent struct {
	ToAssetHash common.Address
	ToAddress   common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnlockEvent is a free log retrieval operation binding the contract event 0xd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df.
//
// Solidity: event UnlockEvent(address toAssetHash, address toAddress, uint256 amount)
func (_SwapProxy *SwapProxyFilterer) FilterUnlockEvent(opts *bind.FilterOpts) (*SwapProxyUnlockEventIterator, error) {

	logs, sub, err := _SwapProxy.contract.FilterLogs(opts, "UnlockEvent")
	if err != nil {
		return nil, err
	}
	return &SwapProxyUnlockEventIterator{contract: _SwapProxy.contract, event: "UnlockEvent", logs: logs, sub: sub}, nil
}

// WatchUnlockEvent is a free log subscription operation binding the contract event 0xd90288730b87c2b8e0c45bd82260fd22478aba30ae1c4d578b8daba9261604df.
//
// Solidity: event UnlockEvent(address toAssetHash, address toAddress, uint256 amount)
func (_SwapProxy *SwapProxyFilterer) WatchUnlockEvent(opts *bind.WatchOpts, sink chan<- *SwapProxyUnlockEvent) (event.Subscription, error) {

	logs, sub, err := _SwapProxy.contract.WatchLogs(opts, "UnlockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapProxyUnlockEvent)
				if err := _SwapProxy.contract.UnpackLog(event, "UnlockEvent", log); err != nil {
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
func (_SwapProxy *SwapProxyFilterer) ParseUnlockEvent(log types.Log) (*SwapProxyUnlockEvent, error) {
	event := new(SwapProxyUnlockEvent)
	if err := _SwapProxy.contract.UnpackLog(event, "UnlockEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
var UtilsBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122087e312e847fef1ea92504dbcad6887b1a3ff962404b37031a2a36fa8eb3b94e364736f6c634300060c0033"

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySinkABI is the input ABI used to generate the binding from.
const ZeroCopySinkABI = "[]"

// ZeroCopySinkBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySinkBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202968b32839632efd416e504a4d25fbb7e2c49c0139cd18bba97c90866a85002664736f6c634300060c0033"

// DeployZeroCopySink deploys a new Ethereum contract, binding an instance of ZeroCopySink to it.
func DeployZeroCopySink(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySink, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySinkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// ZeroCopySink is an auto generated Go binding around an Ethereum contract.
type ZeroCopySink struct {
	ZeroCopySinkCaller     // Read-only binding to the contract
	ZeroCopySinkTransactor // Write-only binding to the contract
	ZeroCopySinkFilterer   // Log filterer for contract events
}

// ZeroCopySinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySinkSession struct {
	Contract     *ZeroCopySink     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySinkCallerSession struct {
	Contract *ZeroCopySinkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZeroCopySinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySinkTransactorSession struct {
	Contract     *ZeroCopySinkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZeroCopySinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySinkRaw struct {
	Contract *ZeroCopySink // Generic contract binding to access the raw methods on
}

// ZeroCopySinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySinkCallerRaw struct {
	Contract *ZeroCopySinkCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactorRaw struct {
	Contract *ZeroCopySinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySink creates a new instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySink(address common.Address, backend bind.ContractBackend) (*ZeroCopySink, error) {
	contract, err := bindZeroCopySink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// NewZeroCopySinkCaller creates a new read-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySinkCaller, error) {
	contract, err := bindZeroCopySink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkCaller{contract: contract}, nil
}

// NewZeroCopySinkTransactor creates a new write-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySinkTransactor, error) {
	contract, err := bindZeroCopySink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkTransactor{contract: contract}, nil
}

// NewZeroCopySinkFilterer creates a new log filterer instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySinkFilterer, error) {
	contract, err := bindZeroCopySink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkFilterer{contract: contract}, nil
}

// bindZeroCopySink binds a generic wrapper to an already deployed contract.
func bindZeroCopySink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.ZeroCopySinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySourceABI is the input ABI used to generate the binding from.
const ZeroCopySourceABI = "[]"

// ZeroCopySourceBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySourceBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204f3af594df5985466c8a77613d232351f4cd6ad04b792cba10096d4b24adc80b64736f6c634300060c0033"

// DeployZeroCopySource deploys a new Ethereum contract, binding an instance of ZeroCopySource to it.
func DeployZeroCopySource(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySource, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySourceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// ZeroCopySource is an auto generated Go binding around an Ethereum contract.
type ZeroCopySource struct {
	ZeroCopySourceCaller     // Read-only binding to the contract
	ZeroCopySourceTransactor // Write-only binding to the contract
	ZeroCopySourceFilterer   // Log filterer for contract events
}

// ZeroCopySourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySourceSession struct {
	Contract     *ZeroCopySource   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySourceCallerSession struct {
	Contract *ZeroCopySourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ZeroCopySourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySourceTransactorSession struct {
	Contract     *ZeroCopySourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZeroCopySourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySourceRaw struct {
	Contract *ZeroCopySource // Generic contract binding to access the raw methods on
}

// ZeroCopySourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySourceCallerRaw struct {
	Contract *ZeroCopySourceCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactorRaw struct {
	Contract *ZeroCopySourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySource creates a new instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySource(address common.Address, backend bind.ContractBackend) (*ZeroCopySource, error) {
	contract, err := bindZeroCopySource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// NewZeroCopySourceCaller creates a new read-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySourceCaller, error) {
	contract, err := bindZeroCopySource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceCaller{contract: contract}, nil
}

// NewZeroCopySourceTransactor creates a new write-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySourceTransactor, error) {
	contract, err := bindZeroCopySource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceTransactor{contract: contract}, nil
}

// NewZeroCopySourceFilterer creates a new log filterer instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySourceFilterer, error) {
	contract, err := bindZeroCopySource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceFilterer{contract: contract}, nil
}

// bindZeroCopySource binds a generic wrapper to an already deployed contract.
func bindZeroCopySource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.ZeroCopySourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transact(opts, method, params...)
}
