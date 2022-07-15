// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zk_abi

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

// IGettersABI is the input ABI used to generate the binding from.
const IGettersABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getTotalPriorityRequests\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalBlocksExecuted\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGovernor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"l2LogsRootHash\",\"outputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalBlocksVerified\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalBlocksCommitted\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IGettersFuncSigs maps the 4-byte function signature to its string representation.
var IGettersFuncSigs = map[string]string{
	"4fc07d75": "getGovernor()",
	"fe26699e": "getTotalBlocksCommitted()",
	"39607382": "getTotalBlocksExecuted()",
	"af6a2dcd": "getTotalBlocksVerified()",
	"056bd873": "getTotalPriorityRequests()",
	"46657fe9": "getVerifier()",
	"facd743b": "isValidator(address)",
	"54225f93": "l2LogsRootHash(uint32)",
}

// IGetters is an auto generated Go binding around an Ethereum contract.
type IGetters struct {
	IGettersCaller     // Read-only binding to the contract
	IGettersTransactor // Write-only binding to the contract
	IGettersFilterer   // Log filterer for contract events
}

// IGettersCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGettersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGettersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGettersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGettersSession struct {
	Contract     *IGetters         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGettersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGettersCallerSession struct {
	Contract *IGettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IGettersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGettersTransactorSession struct {
	Contract     *IGettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IGettersRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGettersRaw struct {
	Contract *IGetters // Generic contract binding to access the raw methods on
}

// IGettersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGettersCallerRaw struct {
	Contract *IGettersCaller // Generic read-only contract binding to access the raw methods on
}

// IGettersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGettersTransactorRaw struct {
	Contract *IGettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGetters creates a new instance of IGetters, bound to a specific deployed contract.
func NewIGetters(address common.Address, backend bind.ContractBackend) (*IGetters, error) {
	contract, err := bindIGetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGetters{IGettersCaller: IGettersCaller{contract: contract}, IGettersTransactor: IGettersTransactor{contract: contract}, IGettersFilterer: IGettersFilterer{contract: contract}}, nil
}

// NewIGettersCaller creates a new read-only instance of IGetters, bound to a specific deployed contract.
func NewIGettersCaller(address common.Address, caller bind.ContractCaller) (*IGettersCaller, error) {
	contract, err := bindIGetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGettersCaller{contract: contract}, nil
}

// NewIGettersTransactor creates a new write-only instance of IGetters, bound to a specific deployed contract.
func NewIGettersTransactor(address common.Address, transactor bind.ContractTransactor) (*IGettersTransactor, error) {
	contract, err := bindIGetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGettersTransactor{contract: contract}, nil
}

// NewIGettersFilterer creates a new log filterer instance of IGetters, bound to a specific deployed contract.
func NewIGettersFilterer(address common.Address, filterer bind.ContractFilterer) (*IGettersFilterer, error) {
	contract, err := bindIGetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGettersFilterer{contract: contract}, nil
}

// bindIGetters binds a generic wrapper to an already deployed contract.
func bindIGetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGettersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGetters *IGettersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGetters.Contract.IGettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGetters *IGettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGetters.Contract.IGettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGetters *IGettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGetters.Contract.IGettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGetters *IGettersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGetters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGetters *IGettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGetters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGetters *IGettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGetters.Contract.contract.Transact(opts, method, params...)
}

// GetGovernor is a free data retrieval call binding the contract method 0x4fc07d75.
//
// Solidity: function getGovernor() view returns(address)
func (_IGetters *IGettersCaller) GetGovernor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getGovernor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovernor is a free data retrieval call binding the contract method 0x4fc07d75.
//
// Solidity: function getGovernor() view returns(address)
func (_IGetters *IGettersSession) GetGovernor() (common.Address, error) {
	return _IGetters.Contract.GetGovernor(&_IGetters.CallOpts)
}

// GetGovernor is a free data retrieval call binding the contract method 0x4fc07d75.
//
// Solidity: function getGovernor() view returns(address)
func (_IGetters *IGettersCallerSession) GetGovernor() (common.Address, error) {
	return _IGetters.Contract.GetGovernor(&_IGetters.CallOpts)
}

// GetTotalBlocksCommitted is a free data retrieval call binding the contract method 0xfe26699e.
//
// Solidity: function getTotalBlocksCommitted() view returns(uint32)
func (_IGetters *IGettersCaller) GetTotalBlocksCommitted(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getTotalBlocksCommitted")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTotalBlocksCommitted is a free data retrieval call binding the contract method 0xfe26699e.
//
// Solidity: function getTotalBlocksCommitted() view returns(uint32)
func (_IGetters *IGettersSession) GetTotalBlocksCommitted() (uint32, error) {
	return _IGetters.Contract.GetTotalBlocksCommitted(&_IGetters.CallOpts)
}

// GetTotalBlocksCommitted is a free data retrieval call binding the contract method 0xfe26699e.
//
// Solidity: function getTotalBlocksCommitted() view returns(uint32)
func (_IGetters *IGettersCallerSession) GetTotalBlocksCommitted() (uint32, error) {
	return _IGetters.Contract.GetTotalBlocksCommitted(&_IGetters.CallOpts)
}

// GetTotalBlocksExecuted is a free data retrieval call binding the contract method 0x39607382.
//
// Solidity: function getTotalBlocksExecuted() view returns(uint32)
func (_IGetters *IGettersCaller) GetTotalBlocksExecuted(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getTotalBlocksExecuted")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTotalBlocksExecuted is a free data retrieval call binding the contract method 0x39607382.
//
// Solidity: function getTotalBlocksExecuted() view returns(uint32)
func (_IGetters *IGettersSession) GetTotalBlocksExecuted() (uint32, error) {
	return _IGetters.Contract.GetTotalBlocksExecuted(&_IGetters.CallOpts)
}

// GetTotalBlocksExecuted is a free data retrieval call binding the contract method 0x39607382.
//
// Solidity: function getTotalBlocksExecuted() view returns(uint32)
func (_IGetters *IGettersCallerSession) GetTotalBlocksExecuted() (uint32, error) {
	return _IGetters.Contract.GetTotalBlocksExecuted(&_IGetters.CallOpts)
}

// GetTotalBlocksVerified is a free data retrieval call binding the contract method 0xaf6a2dcd.
//
// Solidity: function getTotalBlocksVerified() view returns(uint32)
func (_IGetters *IGettersCaller) GetTotalBlocksVerified(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getTotalBlocksVerified")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTotalBlocksVerified is a free data retrieval call binding the contract method 0xaf6a2dcd.
//
// Solidity: function getTotalBlocksVerified() view returns(uint32)
func (_IGetters *IGettersSession) GetTotalBlocksVerified() (uint32, error) {
	return _IGetters.Contract.GetTotalBlocksVerified(&_IGetters.CallOpts)
}

// GetTotalBlocksVerified is a free data retrieval call binding the contract method 0xaf6a2dcd.
//
// Solidity: function getTotalBlocksVerified() view returns(uint32)
func (_IGetters *IGettersCallerSession) GetTotalBlocksVerified() (uint32, error) {
	return _IGetters.Contract.GetTotalBlocksVerified(&_IGetters.CallOpts)
}

// GetTotalPriorityRequests is a free data retrieval call binding the contract method 0x056bd873.
//
// Solidity: function getTotalPriorityRequests() view returns(uint64)
func (_IGetters *IGettersCaller) GetTotalPriorityRequests(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getTotalPriorityRequests")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTotalPriorityRequests is a free data retrieval call binding the contract method 0x056bd873.
//
// Solidity: function getTotalPriorityRequests() view returns(uint64)
func (_IGetters *IGettersSession) GetTotalPriorityRequests() (uint64, error) {
	return _IGetters.Contract.GetTotalPriorityRequests(&_IGetters.CallOpts)
}

// GetTotalPriorityRequests is a free data retrieval call binding the contract method 0x056bd873.
//
// Solidity: function getTotalPriorityRequests() view returns(uint64)
func (_IGetters *IGettersCallerSession) GetTotalPriorityRequests() (uint64, error) {
	return _IGetters.Contract.GetTotalPriorityRequests(&_IGetters.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IGetters *IGettersCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IGetters *IGettersSession) GetVerifier() (common.Address, error) {
	return _IGetters.Contract.GetVerifier(&_IGetters.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IGetters *IGettersCallerSession) GetVerifier() (common.Address, error) {
	return _IGetters.Contract.GetVerifier(&_IGetters.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IGetters *IGettersCaller) IsValidator(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "isValidator", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IGetters *IGettersSession) IsValidator(_address common.Address) (bool, error) {
	return _IGetters.Contract.IsValidator(&_IGetters.CallOpts, _address)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IGetters *IGettersCallerSession) IsValidator(_address common.Address) (bool, error) {
	return _IGetters.Contract.IsValidator(&_IGetters.CallOpts, _address)
}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x54225f93.
//
// Solidity: function l2LogsRootHash(uint32 blockNumber) view returns(bytes32 hash)
func (_IGetters *IGettersCaller) L2LogsRootHash(opts *bind.CallOpts, blockNumber uint32) ([32]byte, error) {
	var out []interface{}
	err := _IGetters.contract.Call(opts, &out, "l2LogsRootHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x54225f93.
//
// Solidity: function l2LogsRootHash(uint32 blockNumber) view returns(bytes32 hash)
func (_IGetters *IGettersSession) L2LogsRootHash(blockNumber uint32) ([32]byte, error) {
	return _IGetters.Contract.L2LogsRootHash(&_IGetters.CallOpts, blockNumber)
}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x54225f93.
//
// Solidity: function l2LogsRootHash(uint32 blockNumber) view returns(bytes32 hash)
func (_IGetters *IGettersCallerSession) L2LogsRootHash(blockNumber uint32) ([32]byte, error) {
	return _IGetters.Contract.L2LogsRootHash(&_IGetters.CallOpts, blockNumber)
}
