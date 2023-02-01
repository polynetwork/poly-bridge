// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cross_chain_manager_abi

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
	MethodBlackChain = "BlackChain"

	MethodWhiteChain = "WhiteChain"

	MethodImportOuterTransfer = "importOuterTransfer"

	MethodMultiSignRipple = "multiSignRipple"

	MethodReconstructRippleTx = "reconstructRippleTx"

	MethodReplenish = "replenish"

	MethodCheckDone = "checkDone"

	MethodName = "name"

	EventMultiSign = "MultiSign"

	EventReplenishEvent = "ReplenishEvent"

	EventRippleTx = "RippleTx"

	EventMakeProof = "makeProof"
)

// CrossChainManagerAbiABI is the input ABI used to generate the binding from.
const CrossChainManagerAbiABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"payment\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sequence\",\"type\":\"uint32\"}],\"name\":\"MultiSign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"txHashes\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"ReplenishEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txJson\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sequence\",\"type\":\"uint32\"}],\"name\":\"RippleTx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"merkleValueHex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"BlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"makeProof\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"}],\"name\":\"BlackChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"}],\"name\":\"WhiteChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"crossChainID\",\"type\":\"bytes\"}],\"name\":\"checkDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"SourceChainID\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Height\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"Proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Extra\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Signature\",\"type\":\"bytes\"}],\"name\":\"importOuterTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ToChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"AssetAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"TxHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"TxJson\",\"type\":\"string\"}],\"name\":\"multiSignRipple\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"Name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"FromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"TxHash\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ToChainId\",\"type\":\"uint64\"}],\"name\":\"reconstructRippleTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"string[]\",\"name\":\"txHashes\",\"type\":\"string[]\"}],\"name\":\"replenish\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CrossChainManagerAbi is an auto generated Go binding around an Ethereum contract.
type CrossChainManagerAbi struct {
	CrossChainManagerAbiCaller     // Read-only binding to the contract
	CrossChainManagerAbiTransactor // Write-only binding to the contract
	CrossChainManagerAbiFilterer   // Log filterer for contract events
}

// CrossChainManagerAbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossChainManagerAbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainManagerAbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossChainManagerAbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainManagerAbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossChainManagerAbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainManagerAbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossChainManagerAbiSession struct {
	Contract     *CrossChainManagerAbi // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrossChainManagerAbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossChainManagerAbiCallerSession struct {
	Contract *CrossChainManagerAbiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CrossChainManagerAbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossChainManagerAbiTransactorSession struct {
	Contract     *CrossChainManagerAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CrossChainManagerAbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossChainManagerAbiRaw struct {
	Contract *CrossChainManagerAbi // Generic contract binding to access the raw methods on
}

// CrossChainManagerAbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossChainManagerAbiCallerRaw struct {
	Contract *CrossChainManagerAbiCaller // Generic read-only contract binding to access the raw methods on
}

// CrossChainManagerAbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossChainManagerAbiTransactorRaw struct {
	Contract *CrossChainManagerAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossChainManagerAbi creates a new instance of CrossChainManagerAbi, bound to a specific deployed contract.
func NewCrossChainManagerAbi(address common.Address, backend bind.ContractBackend) (*CrossChainManagerAbi, error) {
	contract, err := bindCrossChainManagerAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerAbi{CrossChainManagerAbiCaller: CrossChainManagerAbiCaller{contract: contract}, CrossChainManagerAbiTransactor: CrossChainManagerAbiTransactor{contract: contract}, CrossChainManagerAbiFilterer: CrossChainManagerAbiFilterer{contract: contract}}, nil
}

// NewCrossChainManagerAbiCaller creates a new read-only instance of CrossChainManagerAbi, bound to a specific deployed contract.
func NewCrossChainManagerAbiCaller(address common.Address, caller bind.ContractCaller) (*CrossChainManagerAbiCaller, error) {
	contract, err := bindCrossChainManagerAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerAbiCaller{contract: contract}, nil
}

// NewCrossChainManagerAbiTransactor creates a new write-only instance of CrossChainManagerAbi, bound to a specific deployed contract.
func NewCrossChainManagerAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainManagerAbiTransactor, error) {
	contract, err := bindCrossChainManagerAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerAbiTransactor{contract: contract}, nil
}

// NewCrossChainManagerAbiFilterer creates a new log filterer instance of CrossChainManagerAbi, bound to a specific deployed contract.
func NewCrossChainManagerAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainManagerAbiFilterer, error) {
	contract, err := bindCrossChainManagerAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerAbiFilterer{contract: contract}, nil
}

// bindCrossChainManagerAbi binds a generic wrapper to an already deployed contract.
func bindCrossChainManagerAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossChainManagerAbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainManagerAbi *CrossChainManagerAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainManagerAbi.Contract.CrossChainManagerAbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainManagerAbi *CrossChainManagerAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.CrossChainManagerAbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainManagerAbi *CrossChainManagerAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.CrossChainManagerAbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainManagerAbi *CrossChainManagerAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainManagerAbi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.contract.Transact(opts, method, params...)
}

// CheckDone is a free data retrieval call binding the contract method 0x1245f8d5.
//
// Solidity: function checkDone(uint64 chainID, bytes crossChainID) view returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiCaller) CheckDone(opts *bind.CallOpts, chainID uint64, crossChainID []byte) (bool, error) {
	var out []interface{}
	err := _CrossChainManagerAbi.contract.Call(opts, &out, "checkDone", chainID, crossChainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDone is a free data retrieval call binding the contract method 0x1245f8d5.
//
// Solidity: function checkDone(uint64 chainID, bytes crossChainID) view returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiSession) CheckDone(chainID uint64, crossChainID []byte) (bool, error) {
	return _CrossChainManagerAbi.Contract.CheckDone(&_CrossChainManagerAbi.CallOpts, chainID, crossChainID)
}

// CheckDone is a free data retrieval call binding the contract method 0x1245f8d5.
//
// Solidity: function checkDone(uint64 chainID, bytes crossChainID) view returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiCallerSession) CheckDone(chainID uint64, crossChainID []byte) (bool, error) {
	return _CrossChainManagerAbi.Contract.CheckDone(&_CrossChainManagerAbi.CallOpts, chainID, crossChainID)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string Name)
func (_CrossChainManagerAbi *CrossChainManagerAbiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossChainManagerAbi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string Name)
func (_CrossChainManagerAbi *CrossChainManagerAbiSession) Name() (string, error) {
	return _CrossChainManagerAbi.Contract.Name(&_CrossChainManagerAbi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string Name)
func (_CrossChainManagerAbi *CrossChainManagerAbiCallerSession) Name() (string, error) {
	return _CrossChainManagerAbi.Contract.Name(&_CrossChainManagerAbi.CallOpts)
}

// BlackChain is a paid mutator transaction binding the contract method 0x8a449f03.
//
// Solidity: function BlackChain(uint64 ChainID) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactor) BlackChain(opts *bind.TransactOpts, ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManagerAbi.contract.Transact(opts, "BlackChain", ChainID)
}

// BlackChain is a paid mutator transaction binding the contract method 0x8a449f03.
//
// Solidity: function BlackChain(uint64 ChainID) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiSession) BlackChain(ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.BlackChain(&_CrossChainManagerAbi.TransactOpts, ChainID)
}

// BlackChain is a paid mutator transaction binding the contract method 0x8a449f03.
//
// Solidity: function BlackChain(uint64 ChainID) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactorSession) BlackChain(ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.BlackChain(&_CrossChainManagerAbi.TransactOpts, ChainID)
}

// WhiteChain is a paid mutator transaction binding the contract method 0x99d0e87a.
//
// Solidity: function WhiteChain(uint64 ChainID) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactor) WhiteChain(opts *bind.TransactOpts, ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManagerAbi.contract.Transact(opts, "WhiteChain", ChainID)
}

// WhiteChain is a paid mutator transaction binding the contract method 0x99d0e87a.
//
// Solidity: function WhiteChain(uint64 ChainID) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiSession) WhiteChain(ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.WhiteChain(&_CrossChainManagerAbi.TransactOpts, ChainID)
}

// WhiteChain is a paid mutator transaction binding the contract method 0x99d0e87a.
//
// Solidity: function WhiteChain(uint64 ChainID) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactorSession) WhiteChain(ChainID uint64) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.WhiteChain(&_CrossChainManagerAbi.TransactOpts, ChainID)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0xbbc2a76a.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes Extra, bytes Signature) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactor) ImportOuterTransfer(opts *bind.TransactOpts, SourceChainID uint64, Height uint32, Proof []byte, Extra []byte, Signature []byte) (*types.Transaction, error) {
	return _CrossChainManagerAbi.contract.Transact(opts, "importOuterTransfer", SourceChainID, Height, Proof, Extra, Signature)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0xbbc2a76a.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes Extra, bytes Signature) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiSession) ImportOuterTransfer(SourceChainID uint64, Height uint32, Proof []byte, Extra []byte, Signature []byte) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.ImportOuterTransfer(&_CrossChainManagerAbi.TransactOpts, SourceChainID, Height, Proof, Extra, Signature)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0xbbc2a76a.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes Extra, bytes Signature) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactorSession) ImportOuterTransfer(SourceChainID uint64, Height uint32, Proof []byte, Extra []byte, Signature []byte) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.ImportOuterTransfer(&_CrossChainManagerAbi.TransactOpts, SourceChainID, Height, Proof, Extra, Signature)
}

// MultiSignRipple is a paid mutator transaction binding the contract method 0xb7ef3989.
//
// Solidity: function multiSignRipple(uint64 ToChainId, bytes AssetAddress, uint64 FromChainId, bytes TxHash, string TxJson) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactor) MultiSignRipple(opts *bind.TransactOpts, ToChainId uint64, AssetAddress []byte, FromChainId uint64, TxHash []byte, TxJson string) (*types.Transaction, error) {
	return _CrossChainManagerAbi.contract.Transact(opts, "multiSignRipple", ToChainId, AssetAddress, FromChainId, TxHash, TxJson)
}

// MultiSignRipple is a paid mutator transaction binding the contract method 0xb7ef3989.
//
// Solidity: function multiSignRipple(uint64 ToChainId, bytes AssetAddress, uint64 FromChainId, bytes TxHash, string TxJson) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiSession) MultiSignRipple(ToChainId uint64, AssetAddress []byte, FromChainId uint64, TxHash []byte, TxJson string) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.MultiSignRipple(&_CrossChainManagerAbi.TransactOpts, ToChainId, AssetAddress, FromChainId, TxHash, TxJson)
}

// MultiSignRipple is a paid mutator transaction binding the contract method 0xb7ef3989.
//
// Solidity: function multiSignRipple(uint64 ToChainId, bytes AssetAddress, uint64 FromChainId, bytes TxHash, string TxJson) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactorSession) MultiSignRipple(ToChainId uint64, AssetAddress []byte, FromChainId uint64, TxHash []byte, TxJson string) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.MultiSignRipple(&_CrossChainManagerAbi.TransactOpts, ToChainId, AssetAddress, FromChainId, TxHash, TxJson)
}

// ReconstructRippleTx is a paid mutator transaction binding the contract method 0x3b178819.
//
// Solidity: function reconstructRippleTx(uint64 FromChainId, bytes TxHash, uint64 ToChainId) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactor) ReconstructRippleTx(opts *bind.TransactOpts, FromChainId uint64, TxHash []byte, ToChainId uint64) (*types.Transaction, error) {
	return _CrossChainManagerAbi.contract.Transact(opts, "reconstructRippleTx", FromChainId, TxHash, ToChainId)
}

// ReconstructRippleTx is a paid mutator transaction binding the contract method 0x3b178819.
//
// Solidity: function reconstructRippleTx(uint64 FromChainId, bytes TxHash, uint64 ToChainId) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiSession) ReconstructRippleTx(FromChainId uint64, TxHash []byte, ToChainId uint64) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.ReconstructRippleTx(&_CrossChainManagerAbi.TransactOpts, FromChainId, TxHash, ToChainId)
}

// ReconstructRippleTx is a paid mutator transaction binding the contract method 0x3b178819.
//
// Solidity: function reconstructRippleTx(uint64 FromChainId, bytes TxHash, uint64 ToChainId) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactorSession) ReconstructRippleTx(FromChainId uint64, TxHash []byte, ToChainId uint64) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.ReconstructRippleTx(&_CrossChainManagerAbi.TransactOpts, FromChainId, TxHash, ToChainId)
}

// Replenish is a paid mutator transaction binding the contract method 0xf8bac498.
//
// Solidity: function replenish(uint64 chainID, string[] txHashes) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactor) Replenish(opts *bind.TransactOpts, chainID uint64, txHashes []string) (*types.Transaction, error) {
	return _CrossChainManagerAbi.contract.Transact(opts, "replenish", chainID, txHashes)
}

// Replenish is a paid mutator transaction binding the contract method 0xf8bac498.
//
// Solidity: function replenish(uint64 chainID, string[] txHashes) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiSession) Replenish(chainID uint64, txHashes []string) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.Replenish(&_CrossChainManagerAbi.TransactOpts, chainID, txHashes)
}

// Replenish is a paid mutator transaction binding the contract method 0xf8bac498.
//
// Solidity: function replenish(uint64 chainID, string[] txHashes) returns(bool success)
func (_CrossChainManagerAbi *CrossChainManagerAbiTransactorSession) Replenish(chainID uint64, txHashes []string) (*types.Transaction, error) {
	return _CrossChainManagerAbi.Contract.Replenish(&_CrossChainManagerAbi.TransactOpts, chainID, txHashes)
}

// CrossChainManagerAbiMultiSignIterator is returned from FilterMultiSign and is used to iterate over the raw logs and unpacked data for MultiSign events raised by the CrossChainManagerAbi contract.
type CrossChainManagerAbiMultiSignIterator struct {
	Event *CrossChainManagerAbiMultiSign // Event containing the contract specifics and raw log

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
func (it *CrossChainManagerAbiMultiSignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainManagerAbiMultiSign)
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
		it.Event = new(CrossChainManagerAbiMultiSign)
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
func (it *CrossChainManagerAbiMultiSignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainManagerAbiMultiSignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainManagerAbiMultiSign represents a MultiSign event raised by the CrossChainManagerAbi contract.
type CrossChainManagerAbiMultiSign struct {
	FromChainId uint64
	ToChainId   uint64
	TxHash      string
	Payment     string
	Sequence    uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMultiSign is a free log retrieval operation binding the contract event 0x162a93de0c236723115fd6139780a23fb76844208fcc6a7a51d803138cbe11a0.
//
// Solidity: event MultiSign(uint64 fromChainId, uint64 toChainId, string txHash, string payment, uint32 sequence)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) FilterMultiSign(opts *bind.FilterOpts) (*CrossChainManagerAbiMultiSignIterator, error) {

	logs, sub, err := _CrossChainManagerAbi.contract.FilterLogs(opts, "MultiSign")
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerAbiMultiSignIterator{contract: _CrossChainManagerAbi.contract, event: "MultiSign", logs: logs, sub: sub}, nil
}

// WatchMultiSign is a free log subscription operation binding the contract event 0x162a93de0c236723115fd6139780a23fb76844208fcc6a7a51d803138cbe11a0.
//
// Solidity: event MultiSign(uint64 fromChainId, uint64 toChainId, string txHash, string payment, uint32 sequence)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) WatchMultiSign(opts *bind.WatchOpts, sink chan<- *CrossChainManagerAbiMultiSign) (event.Subscription, error) {

	logs, sub, err := _CrossChainManagerAbi.contract.WatchLogs(opts, "MultiSign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainManagerAbiMultiSign)
				if err := _CrossChainManagerAbi.contract.UnpackLog(event, "MultiSign", log); err != nil {
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

// ParseMultiSign is a log parse operation binding the contract event 0x162a93de0c236723115fd6139780a23fb76844208fcc6a7a51d803138cbe11a0.
//
// Solidity: event MultiSign(uint64 fromChainId, uint64 toChainId, string txHash, string payment, uint32 sequence)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) ParseMultiSign(log types.Log) (*CrossChainManagerAbiMultiSign, error) {
	event := new(CrossChainManagerAbiMultiSign)
	if err := _CrossChainManagerAbi.contract.UnpackLog(event, "MultiSign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainManagerAbiReplenishEventIterator is returned from FilterReplenishEvent and is used to iterate over the raw logs and unpacked data for ReplenishEvent events raised by the CrossChainManagerAbi contract.
type CrossChainManagerAbiReplenishEventIterator struct {
	Event *CrossChainManagerAbiReplenishEvent // Event containing the contract specifics and raw log

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
func (it *CrossChainManagerAbiReplenishEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainManagerAbiReplenishEvent)
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
		it.Event = new(CrossChainManagerAbiReplenishEvent)
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
func (it *CrossChainManagerAbiReplenishEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainManagerAbiReplenishEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainManagerAbiReplenishEvent represents a ReplenishEvent event raised by the CrossChainManagerAbi contract.
type CrossChainManagerAbiReplenishEvent struct {
	TxHashes []string
	ChainID  uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReplenishEvent is a free log retrieval operation binding the contract event 0xac3e52c0a7de47fbd0f9a52b8f205485cd725235d94d678f638e16d02404fb38.
//
// Solidity: event ReplenishEvent(string[] txHashes, uint64 chainID)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) FilterReplenishEvent(opts *bind.FilterOpts) (*CrossChainManagerAbiReplenishEventIterator, error) {

	logs, sub, err := _CrossChainManagerAbi.contract.FilterLogs(opts, "ReplenishEvent")
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerAbiReplenishEventIterator{contract: _CrossChainManagerAbi.contract, event: "ReplenishEvent", logs: logs, sub: sub}, nil
}

// WatchReplenishEvent is a free log subscription operation binding the contract event 0xac3e52c0a7de47fbd0f9a52b8f205485cd725235d94d678f638e16d02404fb38.
//
// Solidity: event ReplenishEvent(string[] txHashes, uint64 chainID)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) WatchReplenishEvent(opts *bind.WatchOpts, sink chan<- *CrossChainManagerAbiReplenishEvent) (event.Subscription, error) {

	logs, sub, err := _CrossChainManagerAbi.contract.WatchLogs(opts, "ReplenishEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainManagerAbiReplenishEvent)
				if err := _CrossChainManagerAbi.contract.UnpackLog(event, "ReplenishEvent", log); err != nil {
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

// ParseReplenishEvent is a log parse operation binding the contract event 0xac3e52c0a7de47fbd0f9a52b8f205485cd725235d94d678f638e16d02404fb38.
//
// Solidity: event ReplenishEvent(string[] txHashes, uint64 chainID)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) ParseReplenishEvent(log types.Log) (*CrossChainManagerAbiReplenishEvent, error) {
	event := new(CrossChainManagerAbiReplenishEvent)
	if err := _CrossChainManagerAbi.contract.UnpackLog(event, "ReplenishEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainManagerAbiRippleTxIterator is returned from FilterRippleTx and is used to iterate over the raw logs and unpacked data for RippleTx events raised by the CrossChainManagerAbi contract.
type CrossChainManagerAbiRippleTxIterator struct {
	Event *CrossChainManagerAbiRippleTx // Event containing the contract specifics and raw log

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
func (it *CrossChainManagerAbiRippleTxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainManagerAbiRippleTx)
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
		it.Event = new(CrossChainManagerAbiRippleTx)
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
func (it *CrossChainManagerAbiRippleTxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainManagerAbiRippleTxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainManagerAbiRippleTx represents a RippleTx event raised by the CrossChainManagerAbi contract.
type CrossChainManagerAbiRippleTx struct {
	FromChainId uint64
	ToChainId   uint64
	TxHash      string
	TxJson      string
	Sequence    uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRippleTx is a free log retrieval operation binding the contract event 0xdc5a7a51ad95eb87bc70191f6b497dec6619c6669a89b2edfed188b271e948ae.
//
// Solidity: event RippleTx(uint64 fromChainId, uint64 toChainId, string txHash, string txJson, uint32 sequence)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) FilterRippleTx(opts *bind.FilterOpts) (*CrossChainManagerAbiRippleTxIterator, error) {

	logs, sub, err := _CrossChainManagerAbi.contract.FilterLogs(opts, "RippleTx")
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerAbiRippleTxIterator{contract: _CrossChainManagerAbi.contract, event: "RippleTx", logs: logs, sub: sub}, nil
}

// WatchRippleTx is a free log subscription operation binding the contract event 0xdc5a7a51ad95eb87bc70191f6b497dec6619c6669a89b2edfed188b271e948ae.
//
// Solidity: event RippleTx(uint64 fromChainId, uint64 toChainId, string txHash, string txJson, uint32 sequence)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) WatchRippleTx(opts *bind.WatchOpts, sink chan<- *CrossChainManagerAbiRippleTx) (event.Subscription, error) {

	logs, sub, err := _CrossChainManagerAbi.contract.WatchLogs(opts, "RippleTx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainManagerAbiRippleTx)
				if err := _CrossChainManagerAbi.contract.UnpackLog(event, "RippleTx", log); err != nil {
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

// ParseRippleTx is a log parse operation binding the contract event 0xdc5a7a51ad95eb87bc70191f6b497dec6619c6669a89b2edfed188b271e948ae.
//
// Solidity: event RippleTx(uint64 fromChainId, uint64 toChainId, string txHash, string txJson, uint32 sequence)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) ParseRippleTx(log types.Log) (*CrossChainManagerAbiRippleTx, error) {
	event := new(CrossChainManagerAbiRippleTx)
	if err := _CrossChainManagerAbi.contract.UnpackLog(event, "RippleTx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainManagerAbiMakeProofIterator is returned from FilterMakeProof and is used to iterate over the raw logs and unpacked data for MakeProof events raised by the CrossChainManagerAbi contract.
type CrossChainManagerAbiMakeProofIterator struct {
	Event *CrossChainManagerAbiMakeProof // Event containing the contract specifics and raw log

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
func (it *CrossChainManagerAbiMakeProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainManagerAbiMakeProof)
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
		it.Event = new(CrossChainManagerAbiMakeProof)
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
func (it *CrossChainManagerAbiMakeProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainManagerAbiMakeProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainManagerAbiMakeProof represents a MakeProof event raised by the CrossChainManagerAbi contract.
type CrossChainManagerAbiMakeProof struct {
	MerkleValueHex string
	BlockHeight    uint64
	Key            string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMakeProof is a free log retrieval operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) FilterMakeProof(opts *bind.FilterOpts) (*CrossChainManagerAbiMakeProofIterator, error) {

	logs, sub, err := _CrossChainManagerAbi.contract.FilterLogs(opts, "makeProof")
	if err != nil {
		return nil, err
	}
	return &CrossChainManagerAbiMakeProofIterator{contract: _CrossChainManagerAbi.contract, event: "makeProof", logs: logs, sub: sub}, nil
}

// WatchMakeProof is a free log subscription operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) WatchMakeProof(opts *bind.WatchOpts, sink chan<- *CrossChainManagerAbiMakeProof) (event.Subscription, error) {

	logs, sub, err := _CrossChainManagerAbi.contract.WatchLogs(opts, "makeProof")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainManagerAbiMakeProof)
				if err := _CrossChainManagerAbi.contract.UnpackLog(event, "makeProof", log); err != nil {
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

// ParseMakeProof is a log parse operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_CrossChainManagerAbi *CrossChainManagerAbiFilterer) ParseMakeProof(log types.Log) (*CrossChainManagerAbiMakeProof, error) {
	event := new(CrossChainManagerAbiMakeProof)
	if err := _CrossChainManagerAbi.contract.UnpackLog(event, "makeProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
