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

// ICrossChainManagerABI is the input ABI used to generate the binding from.
const ICrossChainManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"payment\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sequence\",\"type\":\"uint32\"}],\"name\":\"MultiSign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"txHashes\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"ReplenishEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txJson\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sequence\",\"type\":\"uint32\"}],\"name\":\"RippleTx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"merkleValueHex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"BlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"makeProof\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"}],\"name\":\"BlackChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ChainID\",\"type\":\"uint64\"}],\"name\":\"WhiteChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"crossChainID\",\"type\":\"bytes\"}],\"name\":\"checkDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"SourceChainID\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"Height\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"Proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Extra\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Signature\",\"type\":\"bytes\"}],\"name\":\"importOuterTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ToChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"AssetAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"TxHash\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"TxJson\",\"type\":\"string\"}],\"name\":\"multiSignRipple\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"Name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"FromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"TxHash\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ToChainId\",\"type\":\"uint64\"}],\"name\":\"reconstructRippleTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"string[]\",\"name\":\"txHashes\",\"type\":\"string[]\"}],\"name\":\"replenish\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ICrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var ICrossChainManagerFuncSigs = map[string]string{
	"8a449f03": "BlackChain(uint64)",
	"99d0e87a": "WhiteChain(uint64)",
	"1245f8d5": "checkDone(uint64,bytes)",
	"bbc2a76a": "importOuterTransfer(uint64,uint32,bytes,bytes,bytes)",
	"b7ef3989": "multiSignRipple(uint64,bytes,uint64,bytes,string)",
	"06fdde03": "name()",
	"3b178819": "reconstructRippleTx(uint64,bytes,uint64)",
	"f8bac498": "replenish(uint64,string[])",
}

// ICrossChainManager is an auto generated Go binding around an Ethereum contract.
type ICrossChainManager struct {
	ICrossChainManagerCaller     // Read-only binding to the contract
	ICrossChainManagerTransactor // Write-only binding to the contract
	ICrossChainManagerFilterer   // Log filterer for contract events
}

// ICrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICrossChainManagerSession struct {
	Contract     *ICrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ICrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICrossChainManagerCallerSession struct {
	Contract *ICrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ICrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICrossChainManagerTransactorSession struct {
	Contract     *ICrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ICrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICrossChainManagerRaw struct {
	Contract *ICrossChainManager // Generic contract binding to access the raw methods on
}

// ICrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICrossChainManagerCallerRaw struct {
	Contract *ICrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ICrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICrossChainManagerTransactorRaw struct {
	Contract *ICrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICrossChainManager creates a new instance of ICrossChainManager, bound to a specific deployed contract.
func NewICrossChainManager(address common.Address, backend bind.ContractBackend) (*ICrossChainManager, error) {
	contract, err := bindICrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICrossChainManager{ICrossChainManagerCaller: ICrossChainManagerCaller{contract: contract}, ICrossChainManagerTransactor: ICrossChainManagerTransactor{contract: contract}, ICrossChainManagerFilterer: ICrossChainManagerFilterer{contract: contract}}, nil
}

// NewICrossChainManagerCaller creates a new read-only instance of ICrossChainManager, bound to a specific deployed contract.
func NewICrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*ICrossChainManagerCaller, error) {
	contract, err := bindICrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossChainManagerCaller{contract: contract}, nil
}

// NewICrossChainManagerTransactor creates a new write-only instance of ICrossChainManager, bound to a specific deployed contract.
func NewICrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ICrossChainManagerTransactor, error) {
	contract, err := bindICrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossChainManagerTransactor{contract: contract}, nil
}

// NewICrossChainManagerFilterer creates a new log filterer instance of ICrossChainManager, bound to a specific deployed contract.
func NewICrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ICrossChainManagerFilterer, error) {
	contract, err := bindICrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICrossChainManagerFilterer{contract: contract}, nil
}

// bindICrossChainManager binds a generic wrapper to an already deployed contract.
func bindICrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossChainManager *ICrossChainManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossChainManager.Contract.ICrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossChainManager *ICrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.ICrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossChainManager *ICrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.ICrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossChainManager *ICrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossChainManager *ICrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossChainManager *ICrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// CheckDone is a free data retrieval call binding the contract method 0x1245f8d5.
//
// Solidity: function checkDone(uint64 chainID, bytes crossChainID) view returns(bool success)
func (_ICrossChainManager *ICrossChainManagerCaller) CheckDone(opts *bind.CallOpts, chainID uint64, crossChainID []byte) (bool, error) {
	var out []interface{}
	err := _ICrossChainManager.contract.Call(opts, &out, "checkDone", chainID, crossChainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDone is a free data retrieval call binding the contract method 0x1245f8d5.
//
// Solidity: function checkDone(uint64 chainID, bytes crossChainID) view returns(bool success)
func (_ICrossChainManager *ICrossChainManagerSession) CheckDone(chainID uint64, crossChainID []byte) (bool, error) {
	return _ICrossChainManager.Contract.CheckDone(&_ICrossChainManager.CallOpts, chainID, crossChainID)
}

// CheckDone is a free data retrieval call binding the contract method 0x1245f8d5.
//
// Solidity: function checkDone(uint64 chainID, bytes crossChainID) view returns(bool success)
func (_ICrossChainManager *ICrossChainManagerCallerSession) CheckDone(chainID uint64, crossChainID []byte) (bool, error) {
	return _ICrossChainManager.Contract.CheckDone(&_ICrossChainManager.CallOpts, chainID, crossChainID)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string Name)
func (_ICrossChainManager *ICrossChainManagerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ICrossChainManager.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string Name)
func (_ICrossChainManager *ICrossChainManagerSession) Name() (string, error) {
	return _ICrossChainManager.Contract.Name(&_ICrossChainManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string Name)
func (_ICrossChainManager *ICrossChainManagerCallerSession) Name() (string, error) {
	return _ICrossChainManager.Contract.Name(&_ICrossChainManager.CallOpts)
}

// BlackChain is a paid mutator transaction binding the contract method 0x8a449f03.
//
// Solidity: function BlackChain(uint64 ChainID) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactor) BlackChain(opts *bind.TransactOpts, ChainID uint64) (*types.Transaction, error) {
	return _ICrossChainManager.contract.Transact(opts, "BlackChain", ChainID)
}

// BlackChain is a paid mutator transaction binding the contract method 0x8a449f03.
//
// Solidity: function BlackChain(uint64 ChainID) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerSession) BlackChain(ChainID uint64) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.BlackChain(&_ICrossChainManager.TransactOpts, ChainID)
}

// BlackChain is a paid mutator transaction binding the contract method 0x8a449f03.
//
// Solidity: function BlackChain(uint64 ChainID) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactorSession) BlackChain(ChainID uint64) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.BlackChain(&_ICrossChainManager.TransactOpts, ChainID)
}

// WhiteChain is a paid mutator transaction binding the contract method 0x99d0e87a.
//
// Solidity: function WhiteChain(uint64 ChainID) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactor) WhiteChain(opts *bind.TransactOpts, ChainID uint64) (*types.Transaction, error) {
	return _ICrossChainManager.contract.Transact(opts, "WhiteChain", ChainID)
}

// WhiteChain is a paid mutator transaction binding the contract method 0x99d0e87a.
//
// Solidity: function WhiteChain(uint64 ChainID) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerSession) WhiteChain(ChainID uint64) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.WhiteChain(&_ICrossChainManager.TransactOpts, ChainID)
}

// WhiteChain is a paid mutator transaction binding the contract method 0x99d0e87a.
//
// Solidity: function WhiteChain(uint64 ChainID) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactorSession) WhiteChain(ChainID uint64) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.WhiteChain(&_ICrossChainManager.TransactOpts, ChainID)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0xbbc2a76a.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes Extra, bytes Signature) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactor) ImportOuterTransfer(opts *bind.TransactOpts, SourceChainID uint64, Height uint32, Proof []byte, Extra []byte, Signature []byte) (*types.Transaction, error) {
	return _ICrossChainManager.contract.Transact(opts, "importOuterTransfer", SourceChainID, Height, Proof, Extra, Signature)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0xbbc2a76a.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes Extra, bytes Signature) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerSession) ImportOuterTransfer(SourceChainID uint64, Height uint32, Proof []byte, Extra []byte, Signature []byte) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.ImportOuterTransfer(&_ICrossChainManager.TransactOpts, SourceChainID, Height, Proof, Extra, Signature)
}

// ImportOuterTransfer is a paid mutator transaction binding the contract method 0xbbc2a76a.
//
// Solidity: function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes Proof, bytes Extra, bytes Signature) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactorSession) ImportOuterTransfer(SourceChainID uint64, Height uint32, Proof []byte, Extra []byte, Signature []byte) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.ImportOuterTransfer(&_ICrossChainManager.TransactOpts, SourceChainID, Height, Proof, Extra, Signature)
}

// MultiSignRipple is a paid mutator transaction binding the contract method 0xb7ef3989.
//
// Solidity: function multiSignRipple(uint64 ToChainId, bytes AssetAddress, uint64 FromChainId, bytes TxHash, string TxJson) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactor) MultiSignRipple(opts *bind.TransactOpts, ToChainId uint64, AssetAddress []byte, FromChainId uint64, TxHash []byte, TxJson string) (*types.Transaction, error) {
	return _ICrossChainManager.contract.Transact(opts, "multiSignRipple", ToChainId, AssetAddress, FromChainId, TxHash, TxJson)
}

// MultiSignRipple is a paid mutator transaction binding the contract method 0xb7ef3989.
//
// Solidity: function multiSignRipple(uint64 ToChainId, bytes AssetAddress, uint64 FromChainId, bytes TxHash, string TxJson) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerSession) MultiSignRipple(ToChainId uint64, AssetAddress []byte, FromChainId uint64, TxHash []byte, TxJson string) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.MultiSignRipple(&_ICrossChainManager.TransactOpts, ToChainId, AssetAddress, FromChainId, TxHash, TxJson)
}

// MultiSignRipple is a paid mutator transaction binding the contract method 0xb7ef3989.
//
// Solidity: function multiSignRipple(uint64 ToChainId, bytes AssetAddress, uint64 FromChainId, bytes TxHash, string TxJson) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactorSession) MultiSignRipple(ToChainId uint64, AssetAddress []byte, FromChainId uint64, TxHash []byte, TxJson string) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.MultiSignRipple(&_ICrossChainManager.TransactOpts, ToChainId, AssetAddress, FromChainId, TxHash, TxJson)
}

// ReconstructRippleTx is a paid mutator transaction binding the contract method 0x3b178819.
//
// Solidity: function reconstructRippleTx(uint64 FromChainId, bytes TxHash, uint64 ToChainId) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactor) ReconstructRippleTx(opts *bind.TransactOpts, FromChainId uint64, TxHash []byte, ToChainId uint64) (*types.Transaction, error) {
	return _ICrossChainManager.contract.Transact(opts, "reconstructRippleTx", FromChainId, TxHash, ToChainId)
}

// ReconstructRippleTx is a paid mutator transaction binding the contract method 0x3b178819.
//
// Solidity: function reconstructRippleTx(uint64 FromChainId, bytes TxHash, uint64 ToChainId) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerSession) ReconstructRippleTx(FromChainId uint64, TxHash []byte, ToChainId uint64) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.ReconstructRippleTx(&_ICrossChainManager.TransactOpts, FromChainId, TxHash, ToChainId)
}

// ReconstructRippleTx is a paid mutator transaction binding the contract method 0x3b178819.
//
// Solidity: function reconstructRippleTx(uint64 FromChainId, bytes TxHash, uint64 ToChainId) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactorSession) ReconstructRippleTx(FromChainId uint64, TxHash []byte, ToChainId uint64) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.ReconstructRippleTx(&_ICrossChainManager.TransactOpts, FromChainId, TxHash, ToChainId)
}

// Replenish is a paid mutator transaction binding the contract method 0xf8bac498.
//
// Solidity: function replenish(uint64 chainID, string[] txHashes) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactor) Replenish(opts *bind.TransactOpts, chainID uint64, txHashes []string) (*types.Transaction, error) {
	return _ICrossChainManager.contract.Transact(opts, "replenish", chainID, txHashes)
}

// Replenish is a paid mutator transaction binding the contract method 0xf8bac498.
//
// Solidity: function replenish(uint64 chainID, string[] txHashes) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerSession) Replenish(chainID uint64, txHashes []string) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.Replenish(&_ICrossChainManager.TransactOpts, chainID, txHashes)
}

// Replenish is a paid mutator transaction binding the contract method 0xf8bac498.
//
// Solidity: function replenish(uint64 chainID, string[] txHashes) returns(bool success)
func (_ICrossChainManager *ICrossChainManagerTransactorSession) Replenish(chainID uint64, txHashes []string) (*types.Transaction, error) {
	return _ICrossChainManager.Contract.Replenish(&_ICrossChainManager.TransactOpts, chainID, txHashes)
}

// ICrossChainManagerMultiSignIterator is returned from FilterMultiSign and is used to iterate over the raw logs and unpacked data for MultiSign events raised by the ICrossChainManager contract.
type ICrossChainManagerMultiSignIterator struct {
	Event *ICrossChainManagerMultiSign // Event containing the contract specifics and raw log

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
func (it *ICrossChainManagerMultiSignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainManagerMultiSign)
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
		it.Event = new(ICrossChainManagerMultiSign)
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
func (it *ICrossChainManagerMultiSignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainManagerMultiSignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainManagerMultiSign represents a MultiSign event raised by the ICrossChainManager contract.
type ICrossChainManagerMultiSign struct {
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
func (_ICrossChainManager *ICrossChainManagerFilterer) FilterMultiSign(opts *bind.FilterOpts) (*ICrossChainManagerMultiSignIterator, error) {

	logs, sub, err := _ICrossChainManager.contract.FilterLogs(opts, "MultiSign")
	if err != nil {
		return nil, err
	}
	return &ICrossChainManagerMultiSignIterator{contract: _ICrossChainManager.contract, event: "MultiSign", logs: logs, sub: sub}, nil
}

// WatchMultiSign is a free log subscription operation binding the contract event 0x162a93de0c236723115fd6139780a23fb76844208fcc6a7a51d803138cbe11a0.
//
// Solidity: event MultiSign(uint64 fromChainId, uint64 toChainId, string txHash, string payment, uint32 sequence)
func (_ICrossChainManager *ICrossChainManagerFilterer) WatchMultiSign(opts *bind.WatchOpts, sink chan<- *ICrossChainManagerMultiSign) (event.Subscription, error) {

	logs, sub, err := _ICrossChainManager.contract.WatchLogs(opts, "MultiSign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainManagerMultiSign)
				if err := _ICrossChainManager.contract.UnpackLog(event, "MultiSign", log); err != nil {
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
func (_ICrossChainManager *ICrossChainManagerFilterer) ParseMultiSign(log types.Log) (*ICrossChainManagerMultiSign, error) {
	event := new(ICrossChainManagerMultiSign)
	if err := _ICrossChainManager.contract.UnpackLog(event, "MultiSign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainManagerReplenishEventIterator is returned from FilterReplenishEvent and is used to iterate over the raw logs and unpacked data for ReplenishEvent events raised by the ICrossChainManager contract.
type ICrossChainManagerReplenishEventIterator struct {
	Event *ICrossChainManagerReplenishEvent // Event containing the contract specifics and raw log

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
func (it *ICrossChainManagerReplenishEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainManagerReplenishEvent)
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
		it.Event = new(ICrossChainManagerReplenishEvent)
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
func (it *ICrossChainManagerReplenishEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainManagerReplenishEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainManagerReplenishEvent represents a ReplenishEvent event raised by the ICrossChainManager contract.
type ICrossChainManagerReplenishEvent struct {
	TxHashes []string
	ChainID  uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReplenishEvent is a free log retrieval operation binding the contract event 0xac3e52c0a7de47fbd0f9a52b8f205485cd725235d94d678f638e16d02404fb38.
//
// Solidity: event ReplenishEvent(string[] txHashes, uint64 chainID)
func (_ICrossChainManager *ICrossChainManagerFilterer) FilterReplenishEvent(opts *bind.FilterOpts) (*ICrossChainManagerReplenishEventIterator, error) {

	logs, sub, err := _ICrossChainManager.contract.FilterLogs(opts, "ReplenishEvent")
	if err != nil {
		return nil, err
	}
	return &ICrossChainManagerReplenishEventIterator{contract: _ICrossChainManager.contract, event: "ReplenishEvent", logs: logs, sub: sub}, nil
}

// WatchReplenishEvent is a free log subscription operation binding the contract event 0xac3e52c0a7de47fbd0f9a52b8f205485cd725235d94d678f638e16d02404fb38.
//
// Solidity: event ReplenishEvent(string[] txHashes, uint64 chainID)
func (_ICrossChainManager *ICrossChainManagerFilterer) WatchReplenishEvent(opts *bind.WatchOpts, sink chan<- *ICrossChainManagerReplenishEvent) (event.Subscription, error) {

	logs, sub, err := _ICrossChainManager.contract.WatchLogs(opts, "ReplenishEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainManagerReplenishEvent)
				if err := _ICrossChainManager.contract.UnpackLog(event, "ReplenishEvent", log); err != nil {
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
func (_ICrossChainManager *ICrossChainManagerFilterer) ParseReplenishEvent(log types.Log) (*ICrossChainManagerReplenishEvent, error) {
	event := new(ICrossChainManagerReplenishEvent)
	if err := _ICrossChainManager.contract.UnpackLog(event, "ReplenishEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainManagerRippleTxIterator is returned from FilterRippleTx and is used to iterate over the raw logs and unpacked data for RippleTx events raised by the ICrossChainManager contract.
type ICrossChainManagerRippleTxIterator struct {
	Event *ICrossChainManagerRippleTx // Event containing the contract specifics and raw log

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
func (it *ICrossChainManagerRippleTxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainManagerRippleTx)
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
		it.Event = new(ICrossChainManagerRippleTx)
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
func (it *ICrossChainManagerRippleTxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainManagerRippleTxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainManagerRippleTx represents a RippleTx event raised by the ICrossChainManager contract.
type ICrossChainManagerRippleTx struct {
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
func (_ICrossChainManager *ICrossChainManagerFilterer) FilterRippleTx(opts *bind.FilterOpts) (*ICrossChainManagerRippleTxIterator, error) {

	logs, sub, err := _ICrossChainManager.contract.FilterLogs(opts, "RippleTx")
	if err != nil {
		return nil, err
	}
	return &ICrossChainManagerRippleTxIterator{contract: _ICrossChainManager.contract, event: "RippleTx", logs: logs, sub: sub}, nil
}

// WatchRippleTx is a free log subscription operation binding the contract event 0xdc5a7a51ad95eb87bc70191f6b497dec6619c6669a89b2edfed188b271e948ae.
//
// Solidity: event RippleTx(uint64 fromChainId, uint64 toChainId, string txHash, string txJson, uint32 sequence)
func (_ICrossChainManager *ICrossChainManagerFilterer) WatchRippleTx(opts *bind.WatchOpts, sink chan<- *ICrossChainManagerRippleTx) (event.Subscription, error) {

	logs, sub, err := _ICrossChainManager.contract.WatchLogs(opts, "RippleTx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainManagerRippleTx)
				if err := _ICrossChainManager.contract.UnpackLog(event, "RippleTx", log); err != nil {
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
func (_ICrossChainManager *ICrossChainManagerFilterer) ParseRippleTx(log types.Log) (*ICrossChainManagerRippleTx, error) {
	event := new(ICrossChainManagerRippleTx)
	if err := _ICrossChainManager.contract.UnpackLog(event, "RippleTx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrossChainManagerMakeProofIterator is returned from FilterMakeProof and is used to iterate over the raw logs and unpacked data for MakeProof events raised by the ICrossChainManager contract.
type ICrossChainManagerMakeProofIterator struct {
	Event *ICrossChainManagerMakeProof // Event containing the contract specifics and raw log

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
func (it *ICrossChainManagerMakeProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICrossChainManagerMakeProof)
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
		it.Event = new(ICrossChainManagerMakeProof)
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
func (it *ICrossChainManagerMakeProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICrossChainManagerMakeProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICrossChainManagerMakeProof represents a MakeProof event raised by the ICrossChainManager contract.
type ICrossChainManagerMakeProof struct {
	MerkleValueHex string
	BlockHeight    uint64
	Key            string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMakeProof is a free log retrieval operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_ICrossChainManager *ICrossChainManagerFilterer) FilterMakeProof(opts *bind.FilterOpts) (*ICrossChainManagerMakeProofIterator, error) {

	logs, sub, err := _ICrossChainManager.contract.FilterLogs(opts, "makeProof")
	if err != nil {
		return nil, err
	}
	return &ICrossChainManagerMakeProofIterator{contract: _ICrossChainManager.contract, event: "makeProof", logs: logs, sub: sub}, nil
}

// WatchMakeProof is a free log subscription operation binding the contract event 0x25680d41ae78d1188140c6547c9b1890e26bbfa2e0c5b5f1d81aef8985f4d49d.
//
// Solidity: event makeProof(string merkleValueHex, uint64 BlockHeight, string key)
func (_ICrossChainManager *ICrossChainManagerFilterer) WatchMakeProof(opts *bind.WatchOpts, sink chan<- *ICrossChainManagerMakeProof) (event.Subscription, error) {

	logs, sub, err := _ICrossChainManager.contract.WatchLogs(opts, "makeProof")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICrossChainManagerMakeProof)
				if err := _ICrossChainManager.contract.UnpackLog(event, "makeProof", log); err != nil {
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
func (_ICrossChainManager *ICrossChainManagerFilterer) ParseMakeProof(log types.Log) (*ICrossChainManagerMakeProof, error) {
	event := new(ICrossChainManagerMakeProof)
	if err := _ICrossChainManager.contract.UnpackLog(event, "makeProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

