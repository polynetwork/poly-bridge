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

// IWETHABI is the input ABI used to generate the binding from.
const IWETHABI = "[{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWETHFuncSigs maps the 4-byte function signature to its string representation.
var IWETHFuncSigs = map[string]string{
	"d0e30db0": "deposit()",
	"a9059cbb": "transfer(address,uint256)",
	"2e1a7d4d": "withdraw(uint256)",
}

// IWETH is an auto generated Go binding around an Ethereum contract.
type IWETH struct {
	IWETHCaller     // Read-only binding to the contract
	IWETHTransactor // Write-only binding to the contract
	IWETHFilterer   // Log filterer for contract events
}

// IWETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETHSession struct {
	Contract     *IWETH            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETHCallerSession struct {
	Contract *IWETHCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETHTransactorSession struct {
	Contract     *IWETHTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWETHRaw struct {
	Contract *IWETH // Generic contract binding to access the raw methods on
}

// IWETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETHCallerRaw struct {
	Contract *IWETHCaller // Generic read-only contract binding to access the raw methods on
}

// IWETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETHTransactorRaw struct {
	Contract *IWETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH creates a new instance of IWETH, bound to a specific deployed contract.
func NewIWETH(address common.Address, backend bind.ContractBackend) (*IWETH, error) {
	contract, err := bindIWETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH{IWETHCaller: IWETHCaller{contract: contract}, IWETHTransactor: IWETHTransactor{contract: contract}, IWETHFilterer: IWETHFilterer{contract: contract}}, nil
}

// NewIWETHCaller creates a new read-only instance of IWETH, bound to a specific deployed contract.
func NewIWETHCaller(address common.Address, caller bind.ContractCaller) (*IWETHCaller, error) {
	contract, err := bindIWETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHCaller{contract: contract}, nil
}

// NewIWETHTransactor creates a new write-only instance of IWETH, bound to a specific deployed contract.
func NewIWETHTransactor(address common.Address, transactor bind.ContractTransactor) (*IWETHTransactor, error) {
	contract, err := bindIWETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHTransactor{contract: contract}, nil
}

// NewIWETHFilterer creates a new log filterer instance of IWETH, bound to a specific deployed contract.
func NewIWETHFilterer(address common.Address, filterer bind.ContractFilterer) (*IWETHFilterer, error) {
	contract, err := bindIWETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETHFilterer{contract: contract}, nil
}

// bindIWETH binds a generic wrapper to an already deployed contract.
func bindIWETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH *IWETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH.Contract.IWETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH *IWETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.Contract.IWETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH *IWETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH.Contract.IWETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH *IWETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH *IWETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH *IWETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHSession) Deposit() (*types.Transaction, error) {
	return _IWETH.Contract.Deposit(&_IWETH.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHTransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH.Contract.Deposit(&_IWETH.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Transfer(&_IWETH.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Transfer(&_IWETH.TransactOpts, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Withdraw(&_IWETH.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Withdraw(&_IWETH.TransactOpts, arg0)
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

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"5c975abb": "paused()",
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[]"

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
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

// SwapperABI is the input ABI used to generate the binding from.
const SwapperABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_swapChianId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_lockProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_CCMP\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_swapProxyHash\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AddLiquidityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"SwapEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"assetInPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"asset1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"asset2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"asset3\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"}],\"name\":\"bind3Asset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fromAssetHash\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"}],\"name\":\"bindAssetAndPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"extractFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"}],\"name\":\"getBalanceFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerProxyContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"poolTokenMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"poolTokenAddress\",\"type\":\"address\"}],\"name\":\"registerPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"poolTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"asset1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"asset2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"asset3\",\"type\":\"bytes\"}],\"name\":\"registerPoolWith3Assets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lockProxy\",\"type\":\"address\"}],\"name\":\"setLockProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethCCMProxyAddr\",\"type\":\"address\"}],\"name\":\"setManagerProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_swapChianId\",\"type\":\"uint64\"}],\"name\":\"setSwapChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"swapProxyAddr\",\"type\":\"bytes\"}],\"name\":\"setSwapProxyHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"name\":\"setWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toPoolId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAssetHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapProxyHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SwapperFuncSigs maps the 4-byte function signature to its string representation.
var SwapperFuncSigs = map[string]string{
	"ad5c4648": "WETH()",
	"4672b083": "add_liquidity(address,uint64,uint64,bytes,uint256,uint256,uint256,uint256)",
	"84eae14e": "assetInPool(bytes,uint64)",
	"de77f69b": "bind3Asset(bytes,bytes,bytes,uint64)",
	"9462c985": "bindAssetAndPool(bytes,uint64)",
	"9a8a0592": "chainId()",
	"1745399d": "extractFee(address)",
	"c415b95c": "feeCollector()",
	"59c589a1": "getBalanceFor(address)",
	"8f32d59b": "isOwner()",
	"4ad43eda": "lockProxyAddress()",
	"d798f881": "managerProxyContract()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"af249861": "poolTokenMap(uint64)",
	"a3e28eb4": "registerPool(uint64,address)",
	"16675e4f": "registerPoolWith3Assets(uint64,address,bytes,bytes,bytes)",
	"773c9d9c": "remove_liquidity(address,uint64,uint64,bytes,bytes,uint256,uint256,uint256,uint256)",
	"715018a6": "renounceOwnership()",
	"a42dce80": "setFeeCollector(address)",
	"6f2b6ee6": "setLockProxy(address)",
	"af9980f0": "setManagerProxy(address)",
	"f916aa85": "setSwapChainId(uint64)",
	"1f237452": "setSwapProxyHash(bytes)",
	"5b769f3c": "setWETH(address)",
	"559ebe57": "swap(address,uint64,uint64,bytes,bytes,uint256,uint256,uint256,uint256)",
	"78b556fa": "swapChainId()",
	"5f56bc53": "swapProxyHash()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// SwapperBin is the compiled bytecode used for deploying new contracts.
var SwapperBin = "0x60806040523480156200001157600080fd5b5060405162004ac638038062004ac6833981810160405260e08110156200003757600080fd5b815160208301516040808501516060860151608087015160a088015160c089018051955197999698949793969295919483019291846401000000008211156200007f57600080fd5b9083019060208201858111156200009557600080fd5b8251640100000000811182820188101715620000b057600080fd5b82525081516020918201929091019080838360005b83811015620000df578181015183820152602001620000c5565b50505050905090810190601f1680156200010d5780820380516001836020036101000a031916815260200191505b506040525050506000620001266200026e60201b60201c565b600080546001600160a01b0319166001600160a01b03831690811782556040519293509160008051602062004aa6833981519152908290a3506000805461ffff60a01b1916600160a81b1790556001600160401b038616620001b8576040805162461bcd60e51b8152602060048201526006602482015265085b1959d85b60d21b604482015290519081900360640190fd5b620001c38762000272565b60068054600160401b600160801b031916680100000000000000006001600160401b03891602179055600380546001600160a01b03199081166001600160a01b03878116919091179092556004805482168684161790556001805490911691841691909117905580516200023f90600590602084019062000393565b5050600680546001600160401b0319166001600160401b039590951694909417909355506200042f9350505050565b3390565b6200027c620002dc565b620002ce576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b620002d98162000302565b50565b600080546001600160a01b0316620002f36200026e565b6001600160a01b031614905090565b6001600160a01b038116620003495760405162461bcd60e51b815260040180806020018281038252602681526020018062004a806026913960400191505060405180910390fd5b600080546040516001600160a01b038085169392169160008051602062004aa683398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003d657805160ff191683800117855562000406565b8280016001018555821562000406579182015b8281111562000406578251825591602001919060010190620003e9565b506200041492915062000418565b5090565b5b8082111562000414576000815560010162000419565b614641806200043f6000396000f3fe6080604052600436106101d85760003560e01c80638456cb5911610102578063ad5c464811610095578063d798f88111610064578063d798f88114610cfe578063de77f69b14610d13578063f2fde38b14610ed9578063f916aa8514610f0c576101d8565b8063ad5c464814610c6e578063af24986114610c83578063af9980f014610cb6578063c415b95c14610ce9576101d8565b80639462c985116100d15780639462c98514610b285780639a8a059214610be4578063a3e28eb414610bf9578063a42dce8014610c3b576101d8565b80638456cb5914610a2d57806384eae14e14610a425780638da5cb5b14610afe5780638f32d59b14610b13576101d8565b806359c589a11161017a5780636f2b6ee6116101495780636f2b6ee614610850578063715018a614610883578063773c9d9c1461089857806378b556fa146109fc576101d8565b806359c589a1146107395780635b769f3c1461077e5780635c975abb146107b15780635f56bc53146107c6576101d8565b80633f4ba83a116101b65780633f4ba83a1461049c5780634672b083146104b15780634ad43eda146105a4578063559ebe57146105d5576101d8565b806316675e4f146101dd5780631745399d146103b85780631f237452146103eb575b600080fd5b3480156101e957600080fd5b506103b6600480360360a081101561020057600080fd5b6001600160401b03823516916001600160a01b0360208201351691810190606081016040820135600160201b81111561023857600080fd5b82018360208201111561024a57600080fd5b803590602001918460018302840111600160201b8311171561026b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156102bd57600080fd5b8201836020820111156102cf57600080fd5b803590602001918460018302840111600160201b831117156102f057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561034257600080fd5b82018360208201111561035457600080fd5b803590602001918460018302840111600160201b8311171561037557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610f3f945050505050565b005b3480156103c457600080fd5b506103b6600480360360208110156103db57600080fd5b50356001600160a01b031661115a565b3480156103f757600080fd5b506103b66004803603602081101561040e57600080fd5b810190602081018135600160201b81111561042857600080fd5b82018360208201111561043a57600080fd5b803590602001918460018302840111600160201b8311171561045b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061127c945050505050565b3480156104a857600080fd5b506103b66112da565b61059060048036036101008110156104c857600080fd5b6001600160a01b03823516916001600160401b03602082013581169260408301359091169190810190608081016060820135600160201b81111561050b57600080fd5b82018360208201111561051d57600080fd5b803590602001918460018302840111600160201b8311171561053e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040810135906060013561132b565b604080519115158252519081900360200190f35b3480156105b057600080fd5b506105b9611bcf565b604080516001600160a01b039092168252519081900360200190f35b61059060048036036101208110156105ec57600080fd5b6001600160a01b03823516916001600160401b03602082013581169260408301359091169190810190608081016060820135600160201b81111561062f57600080fd5b82018360208201111561064157600080fd5b803590602001918460018302840111600160201b8311171561066257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156106b457600080fd5b8201836020820111156106c657600080fd5b803590602001918460018302840111600160201b831117156106e757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060208101359060408101359060600135611bde565b34801561074557600080fd5b5061076c6004803603602081101561075c57600080fd5b50356001600160a01b031661253b565b60408051918252519081900360200190f35b34801561078a57600080fd5b506103b6600480360360208110156107a157600080fd5b50356001600160a01b03166125d6565b3480156107bd57600080fd5b5061059061263f565b3480156107d257600080fd5b506107db61264f565b6040805160208082528351818301528351919283929083019185019080838360005b838110156108155781810151838201526020016107fd565b50505050905090810190601f1680156108425780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561085c57600080fd5b506103b66004803603602081101561087357600080fd5b50356001600160a01b03166126dd565b34801561088f57600080fd5b506103b6612791565b61059060048036036101208110156108af57600080fd5b6001600160a01b03823516916001600160401b03602082013581169260408301359091169190810190608081016060820135600160201b8111156108f257600080fd5b82018360208201111561090457600080fd5b803590602001918460018302840111600160201b8311171561092557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561097757600080fd5b82018360208201111561098957600080fd5b803590602001918460018302840111600160201b831117156109aa57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060208101359060408101359060600135612822565b348015610a0857600080fd5b50610a11612fc6565b604080516001600160401b039092168252519081900360200190f35b348015610a3957600080fd5b506103b6612fd5565b348015610a4e57600080fd5b5061059060048036036040811015610a6557600080fd5b810190602081018135600160201b811115610a7f57600080fd5b820183602082011115610a9157600080fd5b803590602001918460018302840111600160201b83111715610ab257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b031691506130249050565b348015610b0a57600080fd5b506105b9613054565b348015610b1f57600080fd5b50610590613063565b348015610b3457600080fd5b5061059060048036036040811015610b4b57600080fd5b810190602081018135600160201b811115610b6557600080fd5b820183602082011115610b7757600080fd5b803590602001918460018302840111600160201b83111715610b9857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b031691506130879050565b348015610bf057600080fd5b50610a11613162565b348015610c0557600080fd5b5061059060048036036040811015610c1c57600080fd5b5080356001600160401b031690602001356001600160a01b031661317d565b348015610c4757600080fd5b506103b660048036036020811015610c5e57600080fd5b50356001600160a01b0316613202565b348015610c7a57600080fd5b506105b96132b6565b348015610c8f57600080fd5b506105b960048036036020811015610ca657600080fd5b50356001600160401b03166132c5565b348015610cc257600080fd5b506103b660048036036020811015610cd957600080fd5b50356001600160a01b03166132e0565b348015610cf557600080fd5b506105b9613349565b348015610d0a57600080fd5b506105b9613358565b348015610d1f57600080fd5b506103b660048036036080811015610d3657600080fd5b810190602081018135600160201b811115610d5057600080fd5b820183602082011115610d6257600080fd5b803590602001918460018302840111600160201b83111715610d8357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610dd557600080fd5b820183602082011115610de757600080fd5b803590602001918460018302840111600160201b83111715610e0857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610e5a57600080fd5b820183602082011115610e6c57600080fd5b803590602001918460018302840111600160201b83111715610e8d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b031691506133679050565b348015610ee557600080fd5b506103b660048036036020811015610efc57600080fd5b50356001600160a01b0316613548565b348015610f1857600080fd5b506103b660048036036020811015610f2f57600080fd5b50356001600160401b0316613598565b610f47613063565b610f86576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b6001600160401b03851660009081526008602090815260409182902080546001600160a01b0319166001600160a01b03881617905590518451600192600792879290918291908401908083835b60208310610ff25780518252601f199092019160209182019101610fd3565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185206001600160401b038c16600090815290835220805460ff19169515159590951790945550508351600192600792869290918291908401908083835b602083106110785780518252601f199092019160209182019101611059565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185206001600160401b038c16600090815290835220805460ff19169515159590951790945550508251600192600792859290918291908401908083835b602083106110fe5780518252601f1990920191602091820191016110df565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160401b039a909a1660009081529990525050909520805460ff19169515159590951790945550505050565b6002546001600160a01b031633146111a9576040805162461bcd60e51b815260206004820152600d60248201526c10b332b2a1b7b63632b1ba37b960991b604482015290519081900360640190fd5b6001600160a01b0381166111e95760405133904780156108fc02916000818181858888f193505050501580156111e3573d6000803e3d6000fd5b50611279565b600254604080516370a0823160e01b81523060048201529051611279926001600160a01b0390811692908516916370a0823191602480820192602092909190829003018186803b15801561123c57600080fd5b505afa158015611250573d6000803e3d6000fd5b505050506040513d602081101561126657600080fd5b50516001600160a01b0384169190613602565b50565b611284613063565b6112c3576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b80516112d6906005906020840190614482565b5050565b6112e2613063565b611321576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b611329613659565b565b60008054600160a81b900460ff1661138a576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff60a81b191690819055600160a01b900460ff16156113e8576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6113f28986613701565b6113fd898685613772565b9450611409898661387a565b6001600160a01b0389161561141e578861142b565b6001546001600160a01b03165b6001600160401b0389166000908152600860205260409020549099506001600160a01b031661149b576040805162461bcd60e51b815260206004820152601760248201527619da5d995b881c1bdbdb08191bc81b9bdd08195e1cda5d604a1b604482015290519081900360640190fd5b60076114a68a613994565b6040518082805190602001908083835b602083106114d55780518252601f1990920191602091820191016114b6565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160401b038d1660009081529152929092205460ff1691506115709050576040805162461bcd60e51b815260206004820152601d60248201527f696e70757420746f6b656e206e6f7420696e20676976656e20706f6f6c000000604482015290519081900360640190fd5b85516115b5576040805162461bcd60e51b815260206004820152600f60248201526e656d70747920746f4164647265737360881b604482015290519081900360640190fd5b60148601516001600160a01b038116611606576040805162461bcd60e51b815260206004820152600e60248201526d7a65726f20746f4164647265737360901b604482015290519081900360640190fd5b61160e614500565b6040518061010001604052808881526020018781526020018b6001600160401b031681526020018a6001600160401b0316815260200161164d8d613994565b815260200161166261165d6139af565b613994565b81526020016116716000613994565b815260200189905290506060611686826139b3565b90506000600460009054906101000a90046001600160a01b03166001600160a01b03166387939a7f6040518163ffffffff1660e01b815260040160206040518083038186803b1580156116d857600080fd5b505afa1580156116ec573d6000803e3d6000fd5b505050506040513d602081101561170257600080fd5b505160065460405163bd5cf62560e01b81526001600160401b039091166004820181815260806024840190815260058054600260001960018316156101000201909116046084860181905295965086956001600160a01b0387169563bd5cf625959492938a939290916044810191606482019160a40190879080156117c85780601f1061179d576101008083540402835291602001916117c8565b820191906000526020600020905b8154815290600101906020018083116117ab57829003601f168201915b5050848103835260038152602001806218591960ea1b815250602001848103825285818151815260200191508051906020019080838360005b83811015611819578181015183820152602001611801565b50505050905090810190601f1680156118465780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561186957600080fd5b505af115801561187d573d6000803e3d6000fd5b505050506040513d602081101561189357600080fd5b50516118d05760405162461bcd60e51b815260040180806020018281038252602f8152602001806145b3602f913960400191505060405180910390fd5b505050507f8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea58a6118fe6139af565b6006546001600160401b03166119146000613994565b60058b60405180876001600160a01b03168152602001866001600160a01b03168152602001856001600160401b031681526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561198b578181015183820152602001611973565b50505050905090810190601f1680156119b85780820380516001836020036101000a031916815260200191505b50838103825285546002600019610100600184161502019091160480825260209091019086908015611a2b5780601f10611a0057610100808354040283529160200191611a2b565b820191906000526020600020905b815481529060010190602001808311611a0e57829003601f168201915b50509850505050505050505060405180910390a17f7b634860445c375b3604695e3d36b0ca94d7342cacaae46d96b8727e86522d328a611a696139af565b8b8b611a756000613994565b8c8c8b8b604051808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160401b03168152602001876001600160401b031681526020018060200180602001868152602001858152602001848152602001838103835288818151815260200191508051906020019080838360005b83811015611b08578181015183820152602001611af0565b50505050905090810190601f168015611b355780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b83811015611b68578181015183820152602001611b50565b50505050905090810190601f168015611b955780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390a160019150506000805460ff60a81b1916600160a81b17905598975050505050505050565b6003546001600160a01b031681565b60008054600160a81b900460ff16611c3d576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff60a81b191690819055600160a01b900460ff1615611c9b576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b611ca58a86613701565b611cb08a8685613772565b9450611cbc8a8661387a565b6001600160a01b038a1615611cd15789611cde565b6001546001600160a01b03165b6001600160401b038a16600090815260086020526040902054909a506001600160a01b0316611d4e576040805162461bcd60e51b815260206004820152601760248201527619da5d995b881c1bdbdb08191bc81b9bdd08195e1cda5d604a1b604482015290519081900360640190fd5b6007611d598b613994565b6040518082805190602001908083835b60208310611d885780518252601f199092019160209182019101611d69565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160401b038e1660009081529152929092205460ff169150611e239050576040805162461bcd60e51b815260206004820152601d60248201527f696e70757420746f6b656e206e6f7420696e20676976656e20706f6f6c000000604482015290519081900360640190fd5b6007876040518082805190602001908083835b60208310611e555780518252601f199092019160209182019101611e36565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160401b038e1660009081529152929092205460ff169150611ef09050576040805162461bcd60e51b815260206004820152601e60248201527f6f757470757420746f6b656e206e6f7420696e20676976656e20706f6f6c0000604482015290519081900360640190fd5b8551611f35576040805162461bcd60e51b815260206004820152600f60248201526e656d70747920746f4164647265737360881b604482015290519081900360640190fd5b60148601516001600160a01b038116611f86576040805162461bcd60e51b815260206004820152600e60248201526d7a65726f20746f4164647265737360901b604482015290519081900360640190fd5b611f8e614500565b6040518061010001604052808881526020018781526020018c6001600160401b031681526020018b6001600160401b03168152602001611fcd8e613994565b8152602001611fdd61165d6139af565b81526020018a81526020018981525090506060611ff9826139b3565b90506000600460009054906101000a90046001600160a01b03166001600160a01b03166387939a7f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561204b57600080fd5b505afa15801561205f573d6000803e3d6000fd5b505050506040513d602081101561207557600080fd5b505160065460405163bd5cf62560e01b81526001600160401b039091166004820181815260806024840190815260058054600260001960018316156101000201909116046084860181905295965086956001600160a01b0387169563bd5cf625959492938a939290916044810191606482019160a401908790801561213b5780601f106121105761010080835404028352916020019161213b565b820191906000526020600020905b81548152906001019060200180831161211e57829003601f168201915b505084810383526004815260200180630737761760e41b815250602001848103825285818151815260200191508051906020019080838360005b8381101561218d578181015183820152602001612175565b50505050905090810190601f1680156121ba5780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b1580156121dd57600080fd5b505af11580156121f1573d6000803e3d6000fd5b505050506040513d602081101561220757600080fd5b50516122445760405162461bcd60e51b815260040180806020018281038252602f8152602001806145b3602f913960400191505060405180910390fd5b505050507f8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea58b6122726139af565b6006546001600160401b03166122886000613994565b60058b60405180876001600160a01b03168152602001866001600160a01b03168152602001856001600160401b031681526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156122ff5781810151838201526020016122e7565b50505050905090810190601f16801561232c5780820380516001836020036101000a031916815260200191505b5083810382528554600260001961010060018416150201909116048082526020909101908690801561239f5780601f106123745761010080835404028352916020019161239f565b820191906000526020600020905b81548152906001019060200180831161238257829003601f168201915b50509850505050505050505060405180910390a17f9e37e0e96b266241aa70174d3c6d60151148a5b4181a57fb3d9475aa39ed06728b6123dd6139af565b8c8c8c8c8c8b8b604051808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160401b03168152602001876001600160401b031681526020018060200180602001868152602001858152602001848152602001838103835288818151815260200191508051906020019080838360005b8381101561247357818101518382015260200161245b565b50505050905090810190601f1680156124a05780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b838110156124d35781810151838201526020016124bb565b50505050905090810190601f1680156125005780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390a160019150506000805460ff60a81b1916600160a81b1790559998505050505050505050565b60006001600160a01b038216612553575030316125d1565b604080516370a0823160e01b8152306004820152905183916001600160a01b038316916370a0823191602480820192602092909190829003018186803b15801561259c57600080fd5b505afa1580156125b0573d6000803e3d6000fd5b505050506040513d60208110156125c657600080fd5b505191506125d19050565b919050565b6125de613063565b61261d576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b600054600160a01b900460ff1690565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156126d55780601f106126aa576101008083540402835291602001916126d5565b820191906000526020600020905b8154815290600101906020018083116126b857829003601f168201915b505050505081565b6126e5613063565b612724576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b6001600160a01b03811661276f576040805162461bcd60e51b815260206004820152600d60248201526c656d747079206164647265737360981b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b612799613063565b6127d8576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008054600160a81b900460ff16612881576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff60a81b191690819055600160a01b900460ff16156128df576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6128e98a86613701565b6128f48a8685613772565b94506129008a8661387a565b6001600160a01b038a16156129155789612922565b6001546001600160a01b03165b6001600160401b038a16600090815260086020526040902054909a506001600160a01b0316612992576040805162461bcd60e51b815260206004820152601760248201527619da5d995b881c1bdbdb08191bc81b9bdd08195e1cda5d604a1b604482015290519081900360640190fd5b6001600160401b0389166000908152600860205260409020546001600160a01b038b8116911614612a0a576040805162461bcd60e51b815260206004820181905260248201527f696e70757420746f6b656e206973206e6f7420706f6f6c204c5020746f6b656e604482015290519081900360640190fd5b6007876040518082805190602001908083835b60208310612a3c5780518252601f199092019160209182019101612a1d565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160401b038e1660009081529152929092205460ff169150612ad79050576040805162461bcd60e51b815260206004820152601e60248201527f6f757470757420746f6b656e206e6f7420696e20676976656e20706f6f6c0000604482015290519081900360640190fd5b8551612b1c576040805162461bcd60e51b815260206004820152600f60248201526e656d70747920746f4164647265737360881b604482015290519081900360640190fd5b60148601516001600160a01b038116612b6d576040805162461bcd60e51b815260206004820152600e60248201526d7a65726f20746f4164647265737360901b604482015290519081900360640190fd5b612b75614500565b6040518061010001604052808881526020018781526020018c6001600160401b031681526020018b6001600160401b03168152602001612bb48e613994565b8152602001612bc461165d6139af565b81526020018a81526020018981525090506060612be0826139b3565b90506000600460009054906101000a90046001600160a01b03166001600160a01b03166387939a7f6040518163ffffffff1660e01b815260040160206040518083038186803b158015612c3257600080fd5b505afa158015612c46573d6000803e3d6000fd5b505050506040513d6020811015612c5c57600080fd5b505160065460405163bd5cf62560e01b81526001600160401b039091166004820181815260806024840190815260058054600260001960018316156101000201909116046084860181905295965086956001600160a01b0387169563bd5cf625959492938a939290916044810191606482019160a4019087908015612d225780601f10612cf757610100808354040283529160200191612d22565b820191906000526020600020905b815481529060010190602001808311612d0557829003601f168201915b5050848103835260068152602001806572656d6f766560d01b815250602001848103825285818151815260200191508051906020019080838360005b83811015612d76578181015183820152602001612d5e565b50505050905090810190601f168015612da35780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b158015612dc657600080fd5b505af1158015612dda573d6000803e3d6000fd5b505050506040513d6020811015612df057600080fd5b5051612e2d5760405162461bcd60e51b815260040180806020018281038252602f8152602001806145b3602f913960400191505060405180910390fd5b505050507f8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea58b612e5b6139af565b6006546001600160401b0316612e716000613994565b60058b60405180876001600160a01b03168152602001866001600160a01b03168152602001856001600160401b031681526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015612ee8578181015183820152602001612ed0565b50505050905090810190601f168015612f155780820380516001836020036101000a031916815260200191505b50838103825285546002600019610100600184161502019091160480825260209091019086908015612f885780601f10612f5d57610100808354040283529160200191612f88565b820191906000526020600020905b815481529060010190602001808311612f6b57829003601f168201915b50509850505050505050505060405180910390a17f7ee445799431a22b707efdb3f751a430c4f01f12d902e952200041d81255a41e8b6123dd6139af565b6006546001600160401b031681565b612fdd613063565b61301c576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b611329613c8e565b81516020818401810180516007825292820194820194909420919093529091526000908152604090205460ff1681565b6000546001600160a01b031690565b600080546001600160a01b03166130786139af565b6001600160a01b031614905090565b6000613091613063565b6130d0576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b60016007846040518082805190602001908083835b602083106131045780518252601f1990920191602091820191016130e5565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160401b039790971660009081529690525050909220805460ff191692151592909217909155506001919050565b6006546801000000000000000090046001600160401b031681565b6000613187613063565b6131c6576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b506001600160401b038216600090815260086020526040902080546001600160a01b0383166001600160a01b0319909116179055600192915050565b61320a613063565b613249576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b6001600160a01b038116613294576040805162461bcd60e51b815260206004820152600d60248201526c656d747079206164647265737360981b604482015290519081900360640190fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031681565b6008602052600090815260409020546001600160a01b031681565b6132e8613063565b613327576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b600480546001600160a01b0319166001600160a01b0392909216919091179055565b6002546001600160a01b031681565b6004546001600160a01b031681565b61336f613063565b6133ae576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b60016007856040518082805190602001908083835b602083106133e25780518252601f1990920191602091820191016133c3565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185206001600160401b038816600090815290835220805460ff19169515159590951790945550508451600192600792879290918291908401908083835b602083106134685780518252601f199092019160209182019101613449565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185206001600160401b038816600090815290835220805460ff19169515159590951790945550508351600192600792869290918291908401908083835b602083106134ee5780518252601f1990920191602091820191016134cf565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160401b039690961660009081529590525050909120805460ff1916911515919091179055505050565b613550613063565b61358f576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b61127981613d1c565b6135a0613063565b6135df576040805162461bcd60e51b81526020600482018190526024820152600080516020614593833981519152604482015290519081900360640190fd5b6006805467ffffffffffffffff19166001600160401b0392909216919091179055565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052613654908490613dbc565b505050565b600054600160a01b900460ff166136ae576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6136e46139af565b604080516001600160a01b039092168252519081900360200190a1565b6001600160a01b03821661375d57803414613758576040805162461bcd60e51b815260206004820152601260248201527134b739bab33334b1b4b2b73a1032ba3432b960711b604482015290519081900360640190fd5b6112d6565b6112d66001600160a01b038316333084613f71565b60006001600160a01b038416613827578234146137cb576040805162461bcd60e51b815260206004820152601260248201527134b739bab33334b1b4b2b73a1032ba3432b960711b604482015290519081900360640190fd5b818311613816576040805162461bcd60e51b8152602060048201526014602482015273616d6f756e74206c657373207468616e2066656560601b604482015290519081900360640190fd5b6138208383613fcb565b9050613873565b813414613870576040805162461bcd60e51b815260206004820152601260248201527134b739bab33334b1b4b2b73a1032ba3432b960711b604482015290519081900360640190fd5b50815b9392505050565b6001600160a01b03821661397a57600160009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b1580156138d857600080fd5b505af11580156138ec573d6000803e3d6000fd5b50506001546003546040805163a9059cbb60e01b81526001600160a01b03928316600482015260248101889052905191909216945063a9059cbb9350604480830193506020928290030181600087803b15801561394857600080fd5b505af115801561395c573d6000803e3d6000fd5b505050506040513d602081101561397257600080fd5b506112d69050565b6003546112d6906001600160a01b03848116911683613602565b604080516014815260609290921b6020830152818101905290565b3390565b6060806139c3836000015161400d565b6139d0846020015161400d565b6139dd85604001516140aa565b6139ea86606001516140aa565b6139f787608001516140ed565b613a048860a001516140ed565b613a118960c001516140ed565b613a1e8a60e001516140ed565b6040516020018089805190602001908083835b60208310613a505780518252601f199092019160209182019101613a31565b51815160209384036101000a60001901801990921691161790528b5191909301928b0191508083835b60208310613a985780518252601f199092019160209182019101613a79565b51815160209384036101000a60001901801990921691161790528a5191909301928a0191508083835b60208310613ae05780518252601f199092019160209182019101613ac1565b51815160209384036101000a600019018019909216911617905289519190930192890191508083835b60208310613b285780518252601f199092019160209182019101613b09565b51815160209384036101000a600019018019909216911617905288519190930192880191508083835b60208310613b705780518252601f199092019160209182019101613b51565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b60208310613bb85780518252601f199092019160209182019101613b99565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310613c005780518252601f199092019160209182019101613be1565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310613c485780518252601f199092019160209182019101613c29565b6001836020036101000a03801982511681845116808217855250505050505090500198505050505050505050604051602081830303815290604052905080915050919050565b600054600160a01b900460ff1615613ce0576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586136e46139af565b6001600160a01b038116613d615760405162461bcd60e51b815260040180806020018281038252602681526020018061456d6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b613dc5826141b3565b613e16576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310613e545780518252601f199092019160209182019101613e35565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613eb6576040519150601f19603f3d011682016040523d82523d6000602084013e613ebb565b606091505b509150915081613f12576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115613f6b57808060200190516020811015613f2e57600080fd5b5051613f6b5760405162461bcd60e51b815260040180806020018281038252602a8152602001806145e2602a913960400191505060405180910390fd5b50505050565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052613f6b908590613dbc565b600061387383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506141ef565b60606001600160ff1b0382111561406b576040805162461bcd60e51b815260206004820152601b60248201527f56616c756520657863656564732075696e743235352072616e67650000000000604482015290519081900360640190fd5b60405160208082526000601f5b8282101561409a5785811a826020860101536001919091019060001901614078565b5050506040818101905292915050565b6040516008808252606091906000601f5b828210156140dd5785811a8260208601015360019190910190600019016140bb565b5050506028810160405292915050565b80516060906140fb81614286565b836040516020018083805190602001908083835b6020831061412e5780518252601f19909201916020918201910161410f565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106141765780518252601f199092019160209182019101614157565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081158015906141e75750808214155b949350505050565b6000818484111561427e5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561424357818101518382015260200161422b565b50505050905090810190601f1680156142705780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b606060fd826001600160401b031610156142aa576142a3826143cc565b90506125d1565b61ffff826001600160401b031611614388576142c960fd60f81b6143e8565b6142d2836143fc565b6040516020018083805190602001908083835b602083106143045780518252601f1990920191602091820191016142e5565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061434c5780518252601f19909201916020918201910161432d565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405290506125d1565b63ffffffff826001600160401b0316116143b2576143a9607f60f91b6143e8565b6142d28361443f565b6143c36001600160f81b03196143e8565b6142d2836140aa565b604080516001815260f89290921b602083015260218201905290565b60606143f68260f81c6143cc565b92915050565b6040516002808252606091906000601f5b8282101561442f5785811a82602086010153600191909101906000190161440d565b5050506022810160405292915050565b6040516004808252606091906000601f5b828210156144725785811a826020860101536001919091019060001901614450565b5050506024810160405292915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106144c357805160ff19168380011785556144f0565b828001600101855582156144f0579182015b828111156144f05782518255916020019190600101906144d5565b506144fc929150614557565b5090565b604051806101000160405280600081526020016000815260200160006001600160401b0316815260200160006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b5b808211156144fc576000815560010161455856fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657245746843726f7373436861696e4d616e616765722063726f7373436861696e206578656375746564206572726f72215361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a26469706673582212208d94ef13710664f80d28d5604959b0c9ea7f935e6e050e10d2223fdedc5c0c7764736f6c634300060c00334f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573738be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"

// DeploySwapper deploys a new Ethereum contract, binding an instance of Swapper to it.
func DeploySwapper(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _chainId uint64, _swapChianId uint64, _lockProxy common.Address, _CCMP common.Address, _weth common.Address, _swapProxyHash []byte) (common.Address, *types.Transaction, *Swapper, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwapperBin), backend, _owner, _chainId, _swapChianId, _lockProxy, _CCMP, _weth, _swapProxyHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swapper{SwapperCaller: SwapperCaller{contract: contract}, SwapperTransactor: SwapperTransactor{contract: contract}, SwapperFilterer: SwapperFilterer{contract: contract}}, nil
}

// Swapper is an auto generated Go binding around an Ethereum contract.
type Swapper struct {
	SwapperCaller     // Read-only binding to the contract
	SwapperTransactor // Write-only binding to the contract
	SwapperFilterer   // Log filterer for contract events
}

// SwapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapperSession struct {
	Contract     *Swapper          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapperCallerSession struct {
	Contract *SwapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SwapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapperTransactorSession struct {
	Contract     *SwapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SwapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapperRaw struct {
	Contract *Swapper // Generic contract binding to access the raw methods on
}

// SwapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapperCallerRaw struct {
	Contract *SwapperCaller // Generic read-only contract binding to access the raw methods on
}

// SwapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapperTransactorRaw struct {
	Contract *SwapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapper creates a new instance of Swapper, bound to a specific deployed contract.
func NewSwapper(address common.Address, backend bind.ContractBackend) (*Swapper, error) {
	contract, err := bindSwapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swapper{SwapperCaller: SwapperCaller{contract: contract}, SwapperTransactor: SwapperTransactor{contract: contract}, SwapperFilterer: SwapperFilterer{contract: contract}}, nil
}

// NewSwapperCaller creates a new read-only instance of Swapper, bound to a specific deployed contract.
func NewSwapperCaller(address common.Address, caller bind.ContractCaller) (*SwapperCaller, error) {
	contract, err := bindSwapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapperCaller{contract: contract}, nil
}

// NewSwapperTransactor creates a new write-only instance of Swapper, bound to a specific deployed contract.
func NewSwapperTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapperTransactor, error) {
	contract, err := bindSwapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapperTransactor{contract: contract}, nil
}

// NewSwapperFilterer creates a new log filterer instance of Swapper, bound to a specific deployed contract.
func NewSwapperFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapperFilterer, error) {
	contract, err := bindSwapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapperFilterer{contract: contract}, nil
}

// bindSwapper binds a generic wrapper to an already deployed contract.
func bindSwapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapper *SwapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapper.Contract.SwapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapper *SwapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapper *SwapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapper *SwapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapper *SwapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapper *SwapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapper.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Swapper *SwapperCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Swapper *SwapperSession) WETH() (common.Address, error) {
	return _Swapper.Contract.WETH(&_Swapper.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Swapper *SwapperCallerSession) WETH() (common.Address, error) {
	return _Swapper.Contract.WETH(&_Swapper.CallOpts)
}

// AssetInPool is a free data retrieval call binding the contract method 0x84eae14e.
//
// Solidity: function assetInPool(bytes , uint64 ) view returns(bool)
func (_Swapper *SwapperCaller) AssetInPool(opts *bind.CallOpts, arg0 []byte, arg1 uint64) (bool, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "assetInPool", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetInPool is a free data retrieval call binding the contract method 0x84eae14e.
//
// Solidity: function assetInPool(bytes , uint64 ) view returns(bool)
func (_Swapper *SwapperSession) AssetInPool(arg0 []byte, arg1 uint64) (bool, error) {
	return _Swapper.Contract.AssetInPool(&_Swapper.CallOpts, arg0, arg1)
}

// AssetInPool is a free data retrieval call binding the contract method 0x84eae14e.
//
// Solidity: function assetInPool(bytes , uint64 ) view returns(bool)
func (_Swapper *SwapperCallerSession) AssetInPool(arg0 []byte, arg1 uint64) (bool, error) {
	return _Swapper.Contract.AssetInPool(&_Swapper.CallOpts, arg0, arg1)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_Swapper *SwapperCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_Swapper *SwapperSession) ChainId() (uint64, error) {
	return _Swapper.Contract.ChainId(&_Swapper.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_Swapper *SwapperCallerSession) ChainId() (uint64, error) {
	return _Swapper.Contract.ChainId(&_Swapper.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Swapper *SwapperCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Swapper *SwapperSession) FeeCollector() (common.Address, error) {
	return _Swapper.Contract.FeeCollector(&_Swapper.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Swapper *SwapperCallerSession) FeeCollector() (common.Address, error) {
	return _Swapper.Contract.FeeCollector(&_Swapper.CallOpts)
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address fromAssetHash) view returns(uint256)
func (_Swapper *SwapperCaller) GetBalanceFor(opts *bind.CallOpts, fromAssetHash common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "getBalanceFor", fromAssetHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address fromAssetHash) view returns(uint256)
func (_Swapper *SwapperSession) GetBalanceFor(fromAssetHash common.Address) (*big.Int, error) {
	return _Swapper.Contract.GetBalanceFor(&_Swapper.CallOpts, fromAssetHash)
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address fromAssetHash) view returns(uint256)
func (_Swapper *SwapperCallerSession) GetBalanceFor(fromAssetHash common.Address) (*big.Int, error) {
	return _Swapper.Contract.GetBalanceFor(&_Swapper.CallOpts, fromAssetHash)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Swapper *SwapperCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Swapper *SwapperSession) IsOwner() (bool, error) {
	return _Swapper.Contract.IsOwner(&_Swapper.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Swapper *SwapperCallerSession) IsOwner() (bool, error) {
	return _Swapper.Contract.IsOwner(&_Swapper.CallOpts)
}

// LockProxyAddress is a free data retrieval call binding the contract method 0x4ad43eda.
//
// Solidity: function lockProxyAddress() view returns(address)
func (_Swapper *SwapperCaller) LockProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "lockProxyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockProxyAddress is a free data retrieval call binding the contract method 0x4ad43eda.
//
// Solidity: function lockProxyAddress() view returns(address)
func (_Swapper *SwapperSession) LockProxyAddress() (common.Address, error) {
	return _Swapper.Contract.LockProxyAddress(&_Swapper.CallOpts)
}

// LockProxyAddress is a free data retrieval call binding the contract method 0x4ad43eda.
//
// Solidity: function lockProxyAddress() view returns(address)
func (_Swapper *SwapperCallerSession) LockProxyAddress() (common.Address, error) {
	return _Swapper.Contract.LockProxyAddress(&_Swapper.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_Swapper *SwapperCaller) ManagerProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "managerProxyContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_Swapper *SwapperSession) ManagerProxyContract() (common.Address, error) {
	return _Swapper.Contract.ManagerProxyContract(&_Swapper.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_Swapper *SwapperCallerSession) ManagerProxyContract() (common.Address, error) {
	return _Swapper.Contract.ManagerProxyContract(&_Swapper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapper *SwapperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapper *SwapperSession) Owner() (common.Address, error) {
	return _Swapper.Contract.Owner(&_Swapper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapper *SwapperCallerSession) Owner() (common.Address, error) {
	return _Swapper.Contract.Owner(&_Swapper.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Swapper *SwapperCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Swapper *SwapperSession) Paused() (bool, error) {
	return _Swapper.Contract.Paused(&_Swapper.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Swapper *SwapperCallerSession) Paused() (bool, error) {
	return _Swapper.Contract.Paused(&_Swapper.CallOpts)
}

// PoolTokenMap is a free data retrieval call binding the contract method 0xaf249861.
//
// Solidity: function poolTokenMap(uint64 ) view returns(address)
func (_Swapper *SwapperCaller) PoolTokenMap(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "poolTokenMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolTokenMap is a free data retrieval call binding the contract method 0xaf249861.
//
// Solidity: function poolTokenMap(uint64 ) view returns(address)
func (_Swapper *SwapperSession) PoolTokenMap(arg0 uint64) (common.Address, error) {
	return _Swapper.Contract.PoolTokenMap(&_Swapper.CallOpts, arg0)
}

// PoolTokenMap is a free data retrieval call binding the contract method 0xaf249861.
//
// Solidity: function poolTokenMap(uint64 ) view returns(address)
func (_Swapper *SwapperCallerSession) PoolTokenMap(arg0 uint64) (common.Address, error) {
	return _Swapper.Contract.PoolTokenMap(&_Swapper.CallOpts, arg0)
}

// SwapChainId is a free data retrieval call binding the contract method 0x78b556fa.
//
// Solidity: function swapChainId() view returns(uint64)
func (_Swapper *SwapperCaller) SwapChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "swapChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SwapChainId is a free data retrieval call binding the contract method 0x78b556fa.
//
// Solidity: function swapChainId() view returns(uint64)
func (_Swapper *SwapperSession) SwapChainId() (uint64, error) {
	return _Swapper.Contract.SwapChainId(&_Swapper.CallOpts)
}

// SwapChainId is a free data retrieval call binding the contract method 0x78b556fa.
//
// Solidity: function swapChainId() view returns(uint64)
func (_Swapper *SwapperCallerSession) SwapChainId() (uint64, error) {
	return _Swapper.Contract.SwapChainId(&_Swapper.CallOpts)
}

// SwapProxyHash is a free data retrieval call binding the contract method 0x5f56bc53.
//
// Solidity: function swapProxyHash() view returns(bytes)
func (_Swapper *SwapperCaller) SwapProxyHash(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "swapProxyHash")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SwapProxyHash is a free data retrieval call binding the contract method 0x5f56bc53.
//
// Solidity: function swapProxyHash() view returns(bytes)
func (_Swapper *SwapperSession) SwapProxyHash() ([]byte, error) {
	return _Swapper.Contract.SwapProxyHash(&_Swapper.CallOpts)
}

// SwapProxyHash is a free data retrieval call binding the contract method 0x5f56bc53.
//
// Solidity: function swapProxyHash() view returns(bytes)
func (_Swapper *SwapperCallerSession) SwapProxyHash() ([]byte, error) {
	return _Swapper.Contract.SwapProxyHash(&_Swapper.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4672b083.
//
// Solidity: function add_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_Swapper *SwapperTransactor) AddLiquidity(opts *bind.TransactOpts, fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "add_liquidity", fromAssetHash, toPoolId, toChainId, toAddress, amount, minOutAmount, fee, id)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4672b083.
//
// Solidity: function add_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_Swapper *SwapperSession) AddLiquidity(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.AddLiquidity(&_Swapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAddress, amount, minOutAmount, fee, id)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4672b083.
//
// Solidity: function add_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_Swapper *SwapperTransactorSession) AddLiquidity(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.AddLiquidity(&_Swapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAddress, amount, minOutAmount, fee, id)
}

// Bind3Asset is a paid mutator transaction binding the contract method 0xde77f69b.
//
// Solidity: function bind3Asset(bytes asset1, bytes asset2, bytes asset3, uint64 poolId) returns()
func (_Swapper *SwapperTransactor) Bind3Asset(opts *bind.TransactOpts, asset1 []byte, asset2 []byte, asset3 []byte, poolId uint64) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "bind3Asset", asset1, asset2, asset3, poolId)
}

// Bind3Asset is a paid mutator transaction binding the contract method 0xde77f69b.
//
// Solidity: function bind3Asset(bytes asset1, bytes asset2, bytes asset3, uint64 poolId) returns()
func (_Swapper *SwapperSession) Bind3Asset(asset1 []byte, asset2 []byte, asset3 []byte, poolId uint64) (*types.Transaction, error) {
	return _Swapper.Contract.Bind3Asset(&_Swapper.TransactOpts, asset1, asset2, asset3, poolId)
}

// Bind3Asset is a paid mutator transaction binding the contract method 0xde77f69b.
//
// Solidity: function bind3Asset(bytes asset1, bytes asset2, bytes asset3, uint64 poolId) returns()
func (_Swapper *SwapperTransactorSession) Bind3Asset(asset1 []byte, asset2 []byte, asset3 []byte, poolId uint64) (*types.Transaction, error) {
	return _Swapper.Contract.Bind3Asset(&_Swapper.TransactOpts, asset1, asset2, asset3, poolId)
}

// BindAssetAndPool is a paid mutator transaction binding the contract method 0x9462c985.
//
// Solidity: function bindAssetAndPool(bytes fromAssetHash, uint64 poolId) returns(bool)
func (_Swapper *SwapperTransactor) BindAssetAndPool(opts *bind.TransactOpts, fromAssetHash []byte, poolId uint64) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "bindAssetAndPool", fromAssetHash, poolId)
}

// BindAssetAndPool is a paid mutator transaction binding the contract method 0x9462c985.
//
// Solidity: function bindAssetAndPool(bytes fromAssetHash, uint64 poolId) returns(bool)
func (_Swapper *SwapperSession) BindAssetAndPool(fromAssetHash []byte, poolId uint64) (*types.Transaction, error) {
	return _Swapper.Contract.BindAssetAndPool(&_Swapper.TransactOpts, fromAssetHash, poolId)
}

// BindAssetAndPool is a paid mutator transaction binding the contract method 0x9462c985.
//
// Solidity: function bindAssetAndPool(bytes fromAssetHash, uint64 poolId) returns(bool)
func (_Swapper *SwapperTransactorSession) BindAssetAndPool(fromAssetHash []byte, poolId uint64) (*types.Transaction, error) {
	return _Swapper.Contract.BindAssetAndPool(&_Swapper.TransactOpts, fromAssetHash, poolId)
}

// ExtractFee is a paid mutator transaction binding the contract method 0x1745399d.
//
// Solidity: function extractFee(address token) returns()
func (_Swapper *SwapperTransactor) ExtractFee(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "extractFee", token)
}

// ExtractFee is a paid mutator transaction binding the contract method 0x1745399d.
//
// Solidity: function extractFee(address token) returns()
func (_Swapper *SwapperSession) ExtractFee(token common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.ExtractFee(&_Swapper.TransactOpts, token)
}

// ExtractFee is a paid mutator transaction binding the contract method 0x1745399d.
//
// Solidity: function extractFee(address token) returns()
func (_Swapper *SwapperTransactorSession) ExtractFee(token common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.ExtractFee(&_Swapper.TransactOpts, token)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Swapper *SwapperTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Swapper *SwapperSession) Pause() (*types.Transaction, error) {
	return _Swapper.Contract.Pause(&_Swapper.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Swapper *SwapperTransactorSession) Pause() (*types.Transaction, error) {
	return _Swapper.Contract.Pause(&_Swapper.TransactOpts)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xa3e28eb4.
//
// Solidity: function registerPool(uint64 poolId, address poolTokenAddress) returns(bool)
func (_Swapper *SwapperTransactor) RegisterPool(opts *bind.TransactOpts, poolId uint64, poolTokenAddress common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "registerPool", poolId, poolTokenAddress)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xa3e28eb4.
//
// Solidity: function registerPool(uint64 poolId, address poolTokenAddress) returns(bool)
func (_Swapper *SwapperSession) RegisterPool(poolId uint64, poolTokenAddress common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.RegisterPool(&_Swapper.TransactOpts, poolId, poolTokenAddress)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xa3e28eb4.
//
// Solidity: function registerPool(uint64 poolId, address poolTokenAddress) returns(bool)
func (_Swapper *SwapperTransactorSession) RegisterPool(poolId uint64, poolTokenAddress common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.RegisterPool(&_Swapper.TransactOpts, poolId, poolTokenAddress)
}

// RegisterPoolWith3Assets is a paid mutator transaction binding the contract method 0x16675e4f.
//
// Solidity: function registerPoolWith3Assets(uint64 poolId, address poolTokenAddress, bytes asset1, bytes asset2, bytes asset3) returns()
func (_Swapper *SwapperTransactor) RegisterPoolWith3Assets(opts *bind.TransactOpts, poolId uint64, poolTokenAddress common.Address, asset1 []byte, asset2 []byte, asset3 []byte) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "registerPoolWith3Assets", poolId, poolTokenAddress, asset1, asset2, asset3)
}

// RegisterPoolWith3Assets is a paid mutator transaction binding the contract method 0x16675e4f.
//
// Solidity: function registerPoolWith3Assets(uint64 poolId, address poolTokenAddress, bytes asset1, bytes asset2, bytes asset3) returns()
func (_Swapper *SwapperSession) RegisterPoolWith3Assets(poolId uint64, poolTokenAddress common.Address, asset1 []byte, asset2 []byte, asset3 []byte) (*types.Transaction, error) {
	return _Swapper.Contract.RegisterPoolWith3Assets(&_Swapper.TransactOpts, poolId, poolTokenAddress, asset1, asset2, asset3)
}

// RegisterPoolWith3Assets is a paid mutator transaction binding the contract method 0x16675e4f.
//
// Solidity: function registerPoolWith3Assets(uint64 poolId, address poolTokenAddress, bytes asset1, bytes asset2, bytes asset3) returns()
func (_Swapper *SwapperTransactorSession) RegisterPoolWith3Assets(poolId uint64, poolTokenAddress common.Address, asset1 []byte, asset2 []byte, asset3 []byte) (*types.Transaction, error) {
	return _Swapper.Contract.RegisterPoolWith3Assets(&_Swapper.TransactOpts, poolId, poolTokenAddress, asset1, asset2, asset3)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x773c9d9c.
//
// Solidity: function remove_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_Swapper *SwapperTransactor) RemoveLiquidity(opts *bind.TransactOpts, fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "remove_liquidity", fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x773c9d9c.
//
// Solidity: function remove_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_Swapper *SwapperSession) RemoveLiquidity(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.RemoveLiquidity(&_Swapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x773c9d9c.
//
// Solidity: function remove_liquidity(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_Swapper *SwapperTransactorSession) RemoveLiquidity(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.RemoveLiquidity(&_Swapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapper *SwapperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapper *SwapperSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swapper.Contract.RenounceOwnership(&_Swapper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapper *SwapperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swapper.Contract.RenounceOwnership(&_Swapper.TransactOpts)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address collector) returns()
func (_Swapper *SwapperTransactor) SetFeeCollector(opts *bind.TransactOpts, collector common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "setFeeCollector", collector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address collector) returns()
func (_Swapper *SwapperSession) SetFeeCollector(collector common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.SetFeeCollector(&_Swapper.TransactOpts, collector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address collector) returns()
func (_Swapper *SwapperTransactorSession) SetFeeCollector(collector common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.SetFeeCollector(&_Swapper.TransactOpts, collector)
}

// SetLockProxy is a paid mutator transaction binding the contract method 0x6f2b6ee6.
//
// Solidity: function setLockProxy(address _lockProxy) returns()
func (_Swapper *SwapperTransactor) SetLockProxy(opts *bind.TransactOpts, _lockProxy common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "setLockProxy", _lockProxy)
}

// SetLockProxy is a paid mutator transaction binding the contract method 0x6f2b6ee6.
//
// Solidity: function setLockProxy(address _lockProxy) returns()
func (_Swapper *SwapperSession) SetLockProxy(_lockProxy common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.SetLockProxy(&_Swapper.TransactOpts, _lockProxy)
}

// SetLockProxy is a paid mutator transaction binding the contract method 0x6f2b6ee6.
//
// Solidity: function setLockProxy(address _lockProxy) returns()
func (_Swapper *SwapperTransactorSession) SetLockProxy(_lockProxy common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.SetLockProxy(&_Swapper.TransactOpts, _lockProxy)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address ethCCMProxyAddr) returns()
func (_Swapper *SwapperTransactor) SetManagerProxy(opts *bind.TransactOpts, ethCCMProxyAddr common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "setManagerProxy", ethCCMProxyAddr)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address ethCCMProxyAddr) returns()
func (_Swapper *SwapperSession) SetManagerProxy(ethCCMProxyAddr common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.SetManagerProxy(&_Swapper.TransactOpts, ethCCMProxyAddr)
}

// SetManagerProxy is a paid mutator transaction binding the contract method 0xaf9980f0.
//
// Solidity: function setManagerProxy(address ethCCMProxyAddr) returns()
func (_Swapper *SwapperTransactorSession) SetManagerProxy(ethCCMProxyAddr common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.SetManagerProxy(&_Swapper.TransactOpts, ethCCMProxyAddr)
}

// SetSwapChainId is a paid mutator transaction binding the contract method 0xf916aa85.
//
// Solidity: function setSwapChainId(uint64 _swapChianId) returns()
func (_Swapper *SwapperTransactor) SetSwapChainId(opts *bind.TransactOpts, _swapChianId uint64) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "setSwapChainId", _swapChianId)
}

// SetSwapChainId is a paid mutator transaction binding the contract method 0xf916aa85.
//
// Solidity: function setSwapChainId(uint64 _swapChianId) returns()
func (_Swapper *SwapperSession) SetSwapChainId(_swapChianId uint64) (*types.Transaction, error) {
	return _Swapper.Contract.SetSwapChainId(&_Swapper.TransactOpts, _swapChianId)
}

// SetSwapChainId is a paid mutator transaction binding the contract method 0xf916aa85.
//
// Solidity: function setSwapChainId(uint64 _swapChianId) returns()
func (_Swapper *SwapperTransactorSession) SetSwapChainId(_swapChianId uint64) (*types.Transaction, error) {
	return _Swapper.Contract.SetSwapChainId(&_Swapper.TransactOpts, _swapChianId)
}

// SetSwapProxyHash is a paid mutator transaction binding the contract method 0x1f237452.
//
// Solidity: function setSwapProxyHash(bytes swapProxyAddr) returns()
func (_Swapper *SwapperTransactor) SetSwapProxyHash(opts *bind.TransactOpts, swapProxyAddr []byte) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "setSwapProxyHash", swapProxyAddr)
}

// SetSwapProxyHash is a paid mutator transaction binding the contract method 0x1f237452.
//
// Solidity: function setSwapProxyHash(bytes swapProxyAddr) returns()
func (_Swapper *SwapperSession) SetSwapProxyHash(swapProxyAddr []byte) (*types.Transaction, error) {
	return _Swapper.Contract.SetSwapProxyHash(&_Swapper.TransactOpts, swapProxyAddr)
}

// SetSwapProxyHash is a paid mutator transaction binding the contract method 0x1f237452.
//
// Solidity: function setSwapProxyHash(bytes swapProxyAddr) returns()
func (_Swapper *SwapperTransactorSession) SetSwapProxyHash(swapProxyAddr []byte) (*types.Transaction, error) {
	return _Swapper.Contract.SetSwapProxyHash(&_Swapper.TransactOpts, swapProxyAddr)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _weth) returns()
func (_Swapper *SwapperTransactor) SetWETH(opts *bind.TransactOpts, _weth common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "setWETH", _weth)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _weth) returns()
func (_Swapper *SwapperSession) SetWETH(_weth common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.SetWETH(&_Swapper.TransactOpts, _weth)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address _weth) returns()
func (_Swapper *SwapperTransactorSession) SetWETH(_weth common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.SetWETH(&_Swapper.TransactOpts, _weth)
}

// Swap is a paid mutator transaction binding the contract method 0x559ebe57.
//
// Solidity: function swap(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_Swapper *SwapperTransactor) Swap(opts *bind.TransactOpts, fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "swap", fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// Swap is a paid mutator transaction binding the contract method 0x559ebe57.
//
// Solidity: function swap(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_Swapper *SwapperSession) Swap(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap(&_Swapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// Swap is a paid mutator transaction binding the contract method 0x559ebe57.
//
// Solidity: function swap(address fromAssetHash, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 minOutAmount, uint256 fee, uint256 id) payable returns(bool)
func (_Swapper *SwapperTransactorSession) Swap(fromAssetHash common.Address, toPoolId uint64, toChainId uint64, toAssetHash []byte, toAddress []byte, amount *big.Int, minOutAmount *big.Int, fee *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap(&_Swapper.TransactOpts, fromAssetHash, toPoolId, toChainId, toAssetHash, toAddress, amount, minOutAmount, fee, id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapper *SwapperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapper *SwapperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.TransferOwnership(&_Swapper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapper *SwapperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.TransferOwnership(&_Swapper.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Swapper *SwapperTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Swapper *SwapperSession) Unpause() (*types.Transaction, error) {
	return _Swapper.Contract.Unpause(&_Swapper.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Swapper *SwapperTransactorSession) Unpause() (*types.Transaction, error) {
	return _Swapper.Contract.Unpause(&_Swapper.TransactOpts)
}

// SwapperAddLiquidityEventIterator is returned from FilterAddLiquidityEvent and is used to iterate over the raw logs and unpacked data for AddLiquidityEvent events raised by the Swapper contract.
type SwapperAddLiquidityEventIterator struct {
	Event *SwapperAddLiquidityEvent // Event containing the contract specifics and raw log

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
func (it *SwapperAddLiquidityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperAddLiquidityEvent)
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
		it.Event = new(SwapperAddLiquidityEvent)
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
func (it *SwapperAddLiquidityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperAddLiquidityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperAddLiquidityEvent represents a AddLiquidityEvent event raised by the Swapper contract.
type SwapperAddLiquidityEvent struct {
	FromAssetHash common.Address
	FromAddress   common.Address
	ToPoolId      uint64
	ToChainId     uint64
	ToAssetHash   []byte
	ToAddress     []byte
	Amount        *big.Int
	Fee           *big.Int
	Id            *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidityEvent is a free log retrieval operation binding the contract event 0x7b634860445c375b3604695e3d36b0ca94d7342cacaae46d96b8727e86522d32.
//
// Solidity: event AddLiquidityEvent(address fromAssetHash, address fromAddress, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 fee, uint256 id)
func (_Swapper *SwapperFilterer) FilterAddLiquidityEvent(opts *bind.FilterOpts) (*SwapperAddLiquidityEventIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "AddLiquidityEvent")
	if err != nil {
		return nil, err
	}
	return &SwapperAddLiquidityEventIterator{contract: _Swapper.contract, event: "AddLiquidityEvent", logs: logs, sub: sub}, nil
}

// WatchAddLiquidityEvent is a free log subscription operation binding the contract event 0x7b634860445c375b3604695e3d36b0ca94d7342cacaae46d96b8727e86522d32.
//
// Solidity: event AddLiquidityEvent(address fromAssetHash, address fromAddress, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 fee, uint256 id)
func (_Swapper *SwapperFilterer) WatchAddLiquidityEvent(opts *bind.WatchOpts, sink chan<- *SwapperAddLiquidityEvent) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "AddLiquidityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperAddLiquidityEvent)
				if err := _Swapper.contract.UnpackLog(event, "AddLiquidityEvent", log); err != nil {
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

// ParseAddLiquidityEvent is a log parse operation binding the contract event 0x7b634860445c375b3604695e3d36b0ca94d7342cacaae46d96b8727e86522d32.
//
// Solidity: event AddLiquidityEvent(address fromAssetHash, address fromAddress, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 fee, uint256 id)
func (_Swapper *SwapperFilterer) ParseAddLiquidityEvent(log types.Log) (*SwapperAddLiquidityEvent, error) {
	event := new(SwapperAddLiquidityEvent)
	if err := _Swapper.contract.UnpackLog(event, "AddLiquidityEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperLockEventIterator is returned from FilterLockEvent and is used to iterate over the raw logs and unpacked data for LockEvent events raised by the Swapper contract.
type SwapperLockEventIterator struct {
	Event *SwapperLockEvent // Event containing the contract specifics and raw log

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
func (it *SwapperLockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperLockEvent)
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
		it.Event = new(SwapperLockEvent)
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
func (it *SwapperLockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperLockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperLockEvent represents a LockEvent event raised by the Swapper contract.
type SwapperLockEvent struct {
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
func (_Swapper *SwapperFilterer) FilterLockEvent(opts *bind.FilterOpts) (*SwapperLockEventIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return &SwapperLockEventIterator{contract: _Swapper.contract, event: "LockEvent", logs: logs, sub: sub}, nil
}

// WatchLockEvent is a free log subscription operation binding the contract event 0x8636abd6d0e464fe725a13346c7ac779b73561c705506044a2e6b2cdb1295ea5.
//
// Solidity: event LockEvent(address fromAssetHash, address fromAddress, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount)
func (_Swapper *SwapperFilterer) WatchLockEvent(opts *bind.WatchOpts, sink chan<- *SwapperLockEvent) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperLockEvent)
				if err := _Swapper.contract.UnpackLog(event, "LockEvent", log); err != nil {
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
func (_Swapper *SwapperFilterer) ParseLockEvent(log types.Log) (*SwapperLockEvent, error) {
	event := new(SwapperLockEvent)
	if err := _Swapper.contract.UnpackLog(event, "LockEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Swapper contract.
type SwapperOwnershipTransferredIterator struct {
	Event *SwapperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwapperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperOwnershipTransferred)
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
		it.Event = new(SwapperOwnershipTransferred)
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
func (it *SwapperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperOwnershipTransferred represents a OwnershipTransferred event raised by the Swapper contract.
type SwapperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swapper *SwapperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapperOwnershipTransferredIterator{contract: _Swapper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swapper *SwapperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperOwnershipTransferred)
				if err := _Swapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Swapper *SwapperFilterer) ParseOwnershipTransferred(log types.Log) (*SwapperOwnershipTransferred, error) {
	event := new(SwapperOwnershipTransferred)
	if err := _Swapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Swapper contract.
type SwapperPausedIterator struct {
	Event *SwapperPaused // Event containing the contract specifics and raw log

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
func (it *SwapperPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperPaused)
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
		it.Event = new(SwapperPaused)
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
func (it *SwapperPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperPaused represents a Paused event raised by the Swapper contract.
type SwapperPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Swapper *SwapperFilterer) FilterPaused(opts *bind.FilterOpts) (*SwapperPausedIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SwapperPausedIterator{contract: _Swapper.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Swapper *SwapperFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SwapperPaused) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperPaused)
				if err := _Swapper.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Swapper *SwapperFilterer) ParsePaused(log types.Log) (*SwapperPaused, error) {
	event := new(SwapperPaused)
	if err := _Swapper.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperRemoveLiquidityEventIterator is returned from FilterRemoveLiquidityEvent and is used to iterate over the raw logs and unpacked data for RemoveLiquidityEvent events raised by the Swapper contract.
type SwapperRemoveLiquidityEventIterator struct {
	Event *SwapperRemoveLiquidityEvent // Event containing the contract specifics and raw log

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
func (it *SwapperRemoveLiquidityEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperRemoveLiquidityEvent)
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
		it.Event = new(SwapperRemoveLiquidityEvent)
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
func (it *SwapperRemoveLiquidityEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperRemoveLiquidityEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperRemoveLiquidityEvent represents a RemoveLiquidityEvent event raised by the Swapper contract.
type SwapperRemoveLiquidityEvent struct {
	FromAssetHash common.Address
	FromAddress   common.Address
	ToPoolId      uint64
	ToChainId     uint64
	ToAssetHash   []byte
	ToAddress     []byte
	Amount        *big.Int
	Fee           *big.Int
	Id            *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityEvent is a free log retrieval operation binding the contract event 0x7ee445799431a22b707efdb3f751a430c4f01f12d902e952200041d81255a41e.
//
// Solidity: event RemoveLiquidityEvent(address fromAssetHash, address fromAddress, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 fee, uint256 id)
func (_Swapper *SwapperFilterer) FilterRemoveLiquidityEvent(opts *bind.FilterOpts) (*SwapperRemoveLiquidityEventIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "RemoveLiquidityEvent")
	if err != nil {
		return nil, err
	}
	return &SwapperRemoveLiquidityEventIterator{contract: _Swapper.contract, event: "RemoveLiquidityEvent", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityEvent is a free log subscription operation binding the contract event 0x7ee445799431a22b707efdb3f751a430c4f01f12d902e952200041d81255a41e.
//
// Solidity: event RemoveLiquidityEvent(address fromAssetHash, address fromAddress, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 fee, uint256 id)
func (_Swapper *SwapperFilterer) WatchRemoveLiquidityEvent(opts *bind.WatchOpts, sink chan<- *SwapperRemoveLiquidityEvent) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "RemoveLiquidityEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperRemoveLiquidityEvent)
				if err := _Swapper.contract.UnpackLog(event, "RemoveLiquidityEvent", log); err != nil {
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

// ParseRemoveLiquidityEvent is a log parse operation binding the contract event 0x7ee445799431a22b707efdb3f751a430c4f01f12d902e952200041d81255a41e.
//
// Solidity: event RemoveLiquidityEvent(address fromAssetHash, address fromAddress, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 fee, uint256 id)
func (_Swapper *SwapperFilterer) ParseRemoveLiquidityEvent(log types.Log) (*SwapperRemoveLiquidityEvent, error) {
	event := new(SwapperRemoveLiquidityEvent)
	if err := _Swapper.contract.UnpackLog(event, "RemoveLiquidityEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperSwapEventIterator is returned from FilterSwapEvent and is used to iterate over the raw logs and unpacked data for SwapEvent events raised by the Swapper contract.
type SwapperSwapEventIterator struct {
	Event *SwapperSwapEvent // Event containing the contract specifics and raw log

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
func (it *SwapperSwapEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperSwapEvent)
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
		it.Event = new(SwapperSwapEvent)
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
func (it *SwapperSwapEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperSwapEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperSwapEvent represents a SwapEvent event raised by the Swapper contract.
type SwapperSwapEvent struct {
	FromAssetHash common.Address
	FromAddress   common.Address
	ToPoolId      uint64
	ToChainId     uint64
	ToAssetHash   []byte
	ToAddress     []byte
	Amount        *big.Int
	Fee           *big.Int
	Id            *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSwapEvent is a free log retrieval operation binding the contract event 0x9e37e0e96b266241aa70174d3c6d60151148a5b4181a57fb3d9475aa39ed0672.
//
// Solidity: event SwapEvent(address fromAssetHash, address fromAddress, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 fee, uint256 id)
func (_Swapper *SwapperFilterer) FilterSwapEvent(opts *bind.FilterOpts) (*SwapperSwapEventIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "SwapEvent")
	if err != nil {
		return nil, err
	}
	return &SwapperSwapEventIterator{contract: _Swapper.contract, event: "SwapEvent", logs: logs, sub: sub}, nil
}

// WatchSwapEvent is a free log subscription operation binding the contract event 0x9e37e0e96b266241aa70174d3c6d60151148a5b4181a57fb3d9475aa39ed0672.
//
// Solidity: event SwapEvent(address fromAssetHash, address fromAddress, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 fee, uint256 id)
func (_Swapper *SwapperFilterer) WatchSwapEvent(opts *bind.WatchOpts, sink chan<- *SwapperSwapEvent) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "SwapEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperSwapEvent)
				if err := _Swapper.contract.UnpackLog(event, "SwapEvent", log); err != nil {
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

// ParseSwapEvent is a log parse operation binding the contract event 0x9e37e0e96b266241aa70174d3c6d60151148a5b4181a57fb3d9475aa39ed0672.
//
// Solidity: event SwapEvent(address fromAssetHash, address fromAddress, uint64 toPoolId, uint64 toChainId, bytes toAssetHash, bytes toAddress, uint256 amount, uint256 fee, uint256 id)
func (_Swapper *SwapperFilterer) ParseSwapEvent(log types.Log) (*SwapperSwapEvent, error) {
	event := new(SwapperSwapEvent)
	if err := _Swapper.contract.UnpackLog(event, "SwapEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Swapper contract.
type SwapperUnpausedIterator struct {
	Event *SwapperUnpaused // Event containing the contract specifics and raw log

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
func (it *SwapperUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperUnpaused)
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
		it.Event = new(SwapperUnpaused)
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
func (it *SwapperUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperUnpaused represents a Unpaused event raised by the Swapper contract.
type SwapperUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Swapper *SwapperFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SwapperUnpausedIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SwapperUnpausedIterator{contract: _Swapper.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Swapper *SwapperFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SwapperUnpaused) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperUnpaused)
				if err := _Swapper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Swapper *SwapperFilterer) ParseUnpaused(log types.Log) (*SwapperUnpaused, error) {
	event := new(SwapperUnpaused)
	if err := _Swapper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
