// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ECVerifyABI is the input ABI used to generate the binding from.
const ECVerifyABI = "[]"

// ECVerifyBin is the compiled bytecode used for deploying new contracts.
const ECVerifyBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820805f31bab2756bde4263c739f76cf8fa2eadb30232a45307bc83b2113a0b27230029`

// DeployECVerify deploys a new Ethereum contract, binding an instance of ECVerify to it.
func DeployECVerify(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECVerify, error) {
	parsed, err := abi.JSON(strings.NewReader(ECVerifyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECVerifyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECVerify{ECVerifyCaller: ECVerifyCaller{contract: contract}, ECVerifyTransactor: ECVerifyTransactor{contract: contract}, ECVerifyFilterer: ECVerifyFilterer{contract: contract}}, nil
}

// ECVerify is an auto generated Go binding around an Ethereum contract.
type ECVerify struct {
	ECVerifyCaller     // Read-only binding to the contract
	ECVerifyTransactor // Write-only binding to the contract
	ECVerifyFilterer   // Log filterer for contract events
}

// ECVerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECVerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECVerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECVerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECVerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECVerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECVerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECVerifySession struct {
	Contract     *ECVerify         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECVerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECVerifyCallerSession struct {
	Contract *ECVerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ECVerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECVerifyTransactorSession struct {
	Contract     *ECVerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ECVerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECVerifyRaw struct {
	Contract *ECVerify // Generic contract binding to access the raw methods on
}

// ECVerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECVerifyCallerRaw struct {
	Contract *ECVerifyCaller // Generic read-only contract binding to access the raw methods on
}

// ECVerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECVerifyTransactorRaw struct {
	Contract *ECVerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECVerify creates a new instance of ECVerify, bound to a specific deployed contract.
func NewECVerify(address common.Address, backend bind.ContractBackend) (*ECVerify, error) {
	contract, err := bindECVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECVerify{ECVerifyCaller: ECVerifyCaller{contract: contract}, ECVerifyTransactor: ECVerifyTransactor{contract: contract}, ECVerifyFilterer: ECVerifyFilterer{contract: contract}}, nil
}

// NewECVerifyCaller creates a new read-only instance of ECVerify, bound to a specific deployed contract.
func NewECVerifyCaller(address common.Address, caller bind.ContractCaller) (*ECVerifyCaller, error) {
	contract, err := bindECVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECVerifyCaller{contract: contract}, nil
}

// NewECVerifyTransactor creates a new write-only instance of ECVerify, bound to a specific deployed contract.
func NewECVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*ECVerifyTransactor, error) {
	contract, err := bindECVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECVerifyTransactor{contract: contract}, nil
}

// NewECVerifyFilterer creates a new log filterer instance of ECVerify, bound to a specific deployed contract.
func NewECVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*ECVerifyFilterer, error) {
	contract, err := bindECVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECVerifyFilterer{contract: contract}, nil
}

// bindECVerify binds a generic wrapper to an already deployed contract.
func bindECVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECVerifyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECVerify *ECVerifyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECVerify.Contract.ECVerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECVerify *ECVerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECVerify.Contract.ECVerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECVerify *ECVerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECVerify.Contract.ECVerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECVerify *ECVerifyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECVerify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECVerify *ECVerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECVerify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECVerify *ECVerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECVerify.Contract.contract.Transact(opts, method, params...)
}

// SecretRegistryABI is the input ABI used to generate the binding from.
const SecretRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"registerSecret\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"secrethash_to_block\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contract_version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"secrethash\",\"type\":\"bytes32\"}],\"name\":\"getSecretRevealBlockHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"SecretRevealed\",\"type\":\"event\"}]"

// SecretRegistryBin is the compiled bytecode used for deploying new contracts.
const SecretRegistryBin = `0x608060405234801561001057600080fd5b5061032f806100206000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312ad8bfc81146100665780639734030914610092578063b32c65c8146100bc578063c1f6294614610146575b600080fd5b34801561007257600080fd5b5061007e60043561015e565b604080519115158252519081900360200190f35b34801561009e57600080fd5b506100aa6004356102a8565b60408051918252519081900360200190f35b3480156100c857600080fd5b506100d16102ba565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561010b5781810151838201526020016100f3565b50505050905090810190601f1680156101385780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015257600080fd5b506100aa6004356102f1565b6040805160208082018490528251808303820181529183019283905281516000938493600293909282918401908083835b602083106101cc57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161018f565b51815160209384036101000a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018019909216911617905260405191909301945091925050808303816000865af115801561022b573d6000803e3d6000fd5b5050506040513d602081101561024057600080fd5b5051905082158061025d5750600081815260208190526040812054115b1561026757600080fd5b6000818152602081905260408082204390555184917f9b7ddc883342824bd7ddbff103e7a69f8f2e60b96c075cd1b8b8b9713ecc75a491a250600192915050565b60006020819052908152604090205481565b60408051808201909152600581527f302e332e5f000000000000000000000000000000000000000000000000000000602082015281565b600090815260208190526040902054905600a165627a7a7230582083e075d6d414060311b6f7f9be579a691f4ae80c351e642125fa7ea6c98add3a0029`

// DeploySecretRegistry deploys a new Ethereum contract, binding an instance of SecretRegistry to it.
func DeploySecretRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SecretRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(SecretRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SecretRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SecretRegistry{SecretRegistryCaller: SecretRegistryCaller{contract: contract}, SecretRegistryTransactor: SecretRegistryTransactor{contract: contract}, SecretRegistryFilterer: SecretRegistryFilterer{contract: contract}}, nil
}

// SecretRegistry is an auto generated Go binding around an Ethereum contract.
type SecretRegistry struct {
	SecretRegistryCaller     // Read-only binding to the contract
	SecretRegistryTransactor // Write-only binding to the contract
	SecretRegistryFilterer   // Log filterer for contract events
}

// SecretRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SecretRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecretRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SecretRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecretRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SecretRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecretRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SecretRegistrySession struct {
	Contract     *SecretRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SecretRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SecretRegistryCallerSession struct {
	Contract *SecretRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SecretRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SecretRegistryTransactorSession struct {
	Contract     *SecretRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SecretRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SecretRegistryRaw struct {
	Contract *SecretRegistry // Generic contract binding to access the raw methods on
}

// SecretRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SecretRegistryCallerRaw struct {
	Contract *SecretRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// SecretRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SecretRegistryTransactorRaw struct {
	Contract *SecretRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSecretRegistry creates a new instance of SecretRegistry, bound to a specific deployed contract.
func NewSecretRegistry(address common.Address, backend bind.ContractBackend) (*SecretRegistry, error) {
	contract, err := bindSecretRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SecretRegistry{SecretRegistryCaller: SecretRegistryCaller{contract: contract}, SecretRegistryTransactor: SecretRegistryTransactor{contract: contract}, SecretRegistryFilterer: SecretRegistryFilterer{contract: contract}}, nil
}

// NewSecretRegistryCaller creates a new read-only instance of SecretRegistry, bound to a specific deployed contract.
func NewSecretRegistryCaller(address common.Address, caller bind.ContractCaller) (*SecretRegistryCaller, error) {
	contract, err := bindSecretRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SecretRegistryCaller{contract: contract}, nil
}

// NewSecretRegistryTransactor creates a new write-only instance of SecretRegistry, bound to a specific deployed contract.
func NewSecretRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SecretRegistryTransactor, error) {
	contract, err := bindSecretRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SecretRegistryTransactor{contract: contract}, nil
}

// NewSecretRegistryFilterer creates a new log filterer instance of SecretRegistry, bound to a specific deployed contract.
func NewSecretRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SecretRegistryFilterer, error) {
	contract, err := bindSecretRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SecretRegistryFilterer{contract: contract}, nil
}

// bindSecretRegistry binds a generic wrapper to an already deployed contract.
func bindSecretRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SecretRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecretRegistry *SecretRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SecretRegistry.Contract.SecretRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecretRegistry *SecretRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecretRegistry.Contract.SecretRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecretRegistry *SecretRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecretRegistry.Contract.SecretRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecretRegistry *SecretRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SecretRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecretRegistry *SecretRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecretRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecretRegistry *SecretRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecretRegistry.Contract.contract.Transact(opts, method, params...)
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_SecretRegistry *SecretRegistryCaller) ContractVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SecretRegistry.contract.Call(opts, out, "contract_version")
	return *ret0, err
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_SecretRegistry *SecretRegistrySession) ContractVersion() (string, error) {
	return _SecretRegistry.Contract.ContractVersion(&_SecretRegistry.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_SecretRegistry *SecretRegistryCallerSession) ContractVersion() (string, error) {
	return _SecretRegistry.Contract.ContractVersion(&_SecretRegistry.CallOpts)
}

// GetSecretRevealBlockHeight is a free data retrieval call binding the contract method 0xc1f62946.
//
// Solidity: function getSecretRevealBlockHeight(secrethash bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistryCaller) GetSecretRevealBlockHeight(opts *bind.CallOpts, secrethash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SecretRegistry.contract.Call(opts, out, "getSecretRevealBlockHeight", secrethash)
	return *ret0, err
}

// GetSecretRevealBlockHeight is a free data retrieval call binding the contract method 0xc1f62946.
//
// Solidity: function getSecretRevealBlockHeight(secrethash bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistrySession) GetSecretRevealBlockHeight(secrethash [32]byte) (*big.Int, error) {
	return _SecretRegistry.Contract.GetSecretRevealBlockHeight(&_SecretRegistry.CallOpts, secrethash)
}

// GetSecretRevealBlockHeight is a free data retrieval call binding the contract method 0xc1f62946.
//
// Solidity: function getSecretRevealBlockHeight(secrethash bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistryCallerSession) GetSecretRevealBlockHeight(secrethash [32]byte) (*big.Int, error) {
	return _SecretRegistry.Contract.GetSecretRevealBlockHeight(&_SecretRegistry.CallOpts, secrethash)
}

// SecrethashToBlock is a free data retrieval call binding the contract method 0x97340309.
//
// Solidity: function secrethash_to_block( bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistryCaller) SecrethashToBlock(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SecretRegistry.contract.Call(opts, out, "secrethash_to_block", arg0)
	return *ret0, err
}

// SecrethashToBlock is a free data retrieval call binding the contract method 0x97340309.
//
// Solidity: function secrethash_to_block( bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistrySession) SecrethashToBlock(arg0 [32]byte) (*big.Int, error) {
	return _SecretRegistry.Contract.SecrethashToBlock(&_SecretRegistry.CallOpts, arg0)
}

// SecrethashToBlock is a free data retrieval call binding the contract method 0x97340309.
//
// Solidity: function secrethash_to_block( bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistryCallerSession) SecrethashToBlock(arg0 [32]byte) (*big.Int, error) {
	return _SecretRegistry.Contract.SecrethashToBlock(&_SecretRegistry.CallOpts, arg0)
}

// RegisterSecret is a paid mutator transaction binding the contract method 0x12ad8bfc.
//
// Solidity: function registerSecret(secret bytes32) returns(bool)
func (_SecretRegistry *SecretRegistryTransactor) RegisterSecret(opts *bind.TransactOpts, secret [32]byte) (*types.Transaction, error) {
	return _SecretRegistry.contract.Transact(opts, "registerSecret", secret)
}

// RegisterSecret is a paid mutator transaction binding the contract method 0x12ad8bfc.
//
// Solidity: function registerSecret(secret bytes32) returns(bool)
func (_SecretRegistry *SecretRegistrySession) RegisterSecret(secret [32]byte) (*types.Transaction, error) {
	return _SecretRegistry.Contract.RegisterSecret(&_SecretRegistry.TransactOpts, secret)
}

// RegisterSecret is a paid mutator transaction binding the contract method 0x12ad8bfc.
//
// Solidity: function registerSecret(secret bytes32) returns(bool)
func (_SecretRegistry *SecretRegistryTransactorSession) RegisterSecret(secret [32]byte) (*types.Transaction, error) {
	return _SecretRegistry.Contract.RegisterSecret(&_SecretRegistry.TransactOpts, secret)
}

// SecretRegistrySecretRevealedIterator is returned from FilterSecretRevealed and is used to iterate over the raw logs and unpacked data for SecretRevealed events raised by the SecretRegistry contract.
type SecretRegistrySecretRevealedIterator struct {
	Event *SecretRegistrySecretRevealed // Event containing the contract specifics and raw log

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
func (it *SecretRegistrySecretRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecretRegistrySecretRevealed)
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
		it.Event = new(SecretRegistrySecretRevealed)
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
func (it *SecretRegistrySecretRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecretRegistrySecretRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecretRegistrySecretRevealed represents a SecretRevealed event raised by the SecretRegistry contract.
type SecretRegistrySecretRevealed struct {
	Secret [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSecretRevealed is a free log retrieval operation binding the contract event 0x9b7ddc883342824bd7ddbff103e7a69f8f2e60b96c075cd1b8b8b9713ecc75a4.
//
// Solidity: e SecretRevealed(secret indexed bytes32)
func (_SecretRegistry *SecretRegistryFilterer) FilterSecretRevealed(opts *bind.FilterOpts, secret [][32]byte) (*SecretRegistrySecretRevealedIterator, error) {

	var secretRule []interface{}
	for _, secretItem := range secret {
		secretRule = append(secretRule, secretItem)
	}

	logs, sub, err := _SecretRegistry.contract.FilterLogs(opts, "SecretRevealed", secretRule)
	if err != nil {
		return nil, err
	}
	return &SecretRegistrySecretRevealedIterator{contract: _SecretRegistry.contract, event: "SecretRevealed", logs: logs, sub: sub}, nil
}

// WatchSecretRevealed is a free log subscription operation binding the contract event 0x9b7ddc883342824bd7ddbff103e7a69f8f2e60b96c075cd1b8b8b9713ecc75a4.
//
// Solidity: e SecretRevealed(secret indexed bytes32)
func (_SecretRegistry *SecretRegistryFilterer) WatchSecretRevealed(opts *bind.WatchOpts, sink chan<- *SecretRegistrySecretRevealed, secret [][32]byte) (event.Subscription, error) {

	var secretRule []interface{}
	for _, secretItem := range secret {
		secretRule = append(secretRule, secretItem)
	}

	logs, sub, err := _SecretRegistry.contract.WatchLogs(opts, "SecretRevealed", secretRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecretRegistrySecretRevealed)
				if err := _SecretRegistry.contract.UnpackLog(event, "SecretRevealed", log); err != nil {
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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _amount uint256, _extraData bytes) returns(success bool)
func (_Token *TokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approveAndCall", _spender, _amount, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _amount uint256, _extraData bytes) returns(success bool)
func (_Token *TokenSession) ApproveAndCall(_spender common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.Contract.ApproveAndCall(&_Token.TransactOpts, _spender, _amount, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _amount uint256, _extraData bytes) returns(success bool)
func (_Token *TokenTransactorSession) ApproveAndCall(_spender common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.Contract.ApproveAndCall(&_Token.TransactOpts, _spender, _amount, _extraData)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(to address, value uint256, data bytes) returns()
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", to, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(to address, value uint256, data bytes) returns()
func (_Token *TokenSession) Transfer(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(to address, value uint256, data bytes) returns()
func (_Token *TokenTransactorSession) Transfer(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// TokenNetworkABI is the input ABI used to generate the binding from.
const TokenNetworkABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"secret_registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chain_id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"partner\",\"type\":\"address\"},{\"name\":\"transferred_amount\",\"type\":\"uint256\"},{\"name\":\"expiration\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"secret_hash\",\"type\":\"bytes32\"},{\"name\":\"merkle_proof\",\"type\":\"bytes\"}],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contract_address\",\"type\":\"address\"}],\"name\":\"contractExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"channels\",\"outputs\":[{\"name\":\"settle_timeout\",\"type\":\"uint64\"},{\"name\":\"settle_block_number\",\"type\":\"uint64\"},{\"name\":\"open_block_number\",\"type\":\"uint64\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participant\",\"type\":\"address\"},{\"name\":\"partner\",\"type\":\"address\"},{\"name\":\"participant_balance\",\"type\":\"uint256\"},{\"name\":\"participant_withdraw\",\"type\":\"uint256\"},{\"name\":\"participant_signature\",\"type\":\"bytes\"},{\"name\":\"partner_signature\",\"type\":\"bytes\"}],\"name\":\"withDraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participant\",\"type\":\"address\"},{\"name\":\"partner\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"cheater\",\"type\":\"address\"},{\"name\":\"lockhash\",\"type\":\"bytes32\"},{\"name\":\"additional_hash\",\"type\":\"bytes32\"},{\"name\":\"cheater_signature\",\"type\":\"bytes\"}],\"name\":\"punishObsoleteUnlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participant1\",\"type\":\"address\"},{\"name\":\"participant1_balance\",\"type\":\"uint256\"},{\"name\":\"participant2\",\"type\":\"address\"},{\"name\":\"participant2_balance\",\"type\":\"uint256\"},{\"name\":\"participant1_signature\",\"type\":\"bytes\"},{\"name\":\"participant2_signature\",\"type\":\"bytes\"}],\"name\":\"cooperativeSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signature_prefix\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"participant\",\"type\":\"address\"},{\"name\":\"partner\",\"type\":\"address\"},{\"name\":\"lockhash\",\"type\":\"bytes32\"}],\"name\":\"queryUnlockedLocks\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"token_\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"punish_block_number\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"channel_identifier\",\"type\":\"bytes32\"}],\"name\":\"getChannelInfoByChannelIdentifier\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"partner\",\"type\":\"address\"},{\"name\":\"participant\",\"type\":\"address\"},{\"name\":\"transferred_amount\",\"type\":\"uint256\"},{\"name\":\"expiration\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"secret_hash\",\"type\":\"bytes32\"},{\"name\":\"merkle_proof\",\"type\":\"bytes\"},{\"name\":\"participant_signature\",\"type\":\"bytes\"}],\"name\":\"unlockDelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"partner\",\"type\":\"address\"},{\"name\":\"transferred_amount\",\"type\":\"uint256\"},{\"name\":\"locksroot\",\"type\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\"},{\"name\":\"additional_hash\",\"type\":\"bytes32\"},{\"name\":\"partner_signature\",\"type\":\"bytes\"}],\"name\":\"updateBalanceProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"participant\",\"type\":\"address\"},{\"name\":\"partner\",\"type\":\"address\"}],\"name\":\"getChannelParticipantInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes24\"},{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participant1\",\"type\":\"address\"},{\"name\":\"participant2\",\"type\":\"address\"},{\"name\":\"settle_timeout\",\"type\":\"uint64\"}],\"name\":\"openChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contract_version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"partner\",\"type\":\"address\"},{\"name\":\"transferred_amount\",\"type\":\"uint256\"},{\"name\":\"locksroot\",\"type\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\"},{\"name\":\"additional_hash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"closeChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participant1\",\"type\":\"address\"},{\"name\":\"participant1_transferred_amount\",\"type\":\"uint256\"},{\"name\":\"participant1_locksroot\",\"type\":\"bytes32\"},{\"name\":\"participant2\",\"type\":\"address\"},{\"name\":\"participant2_transferred_amount\",\"type\":\"uint256\"},{\"name\":\"participant2_locksroot\",\"type\":\"bytes32\"}],\"name\":\"settleChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"partner\",\"type\":\"address\"},{\"name\":\"participant\",\"type\":\"address\"},{\"name\":\"transferred_amount\",\"type\":\"uint256\"},{\"name\":\"locksroot\",\"type\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\"},{\"name\":\"additional_hash\",\"type\":\"bytes32\"},{\"name\":\"partner_signature\",\"type\":\"bytes\"},{\"name\":\"participant_signature\",\"type\":\"bytes\"}],\"name\":\"updateBalanceProofDelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"participant1\",\"type\":\"address\"},{\"name\":\"participant2\",\"type\":\"address\"}],\"name\":\"getChannelInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participant\",\"type\":\"address\"},{\"name\":\"partner\",\"type\":\"address\"},{\"name\":\"settle_timeout\",\"type\":\"uint64\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"openChannelWithDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token_address\",\"type\":\"address\"},{\"name\":\"_secret_registry\",\"type\":\"address\"},{\"name\":\"_chain_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"participant1\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"participant2\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"settle_timeout\",\"type\":\"uint64\"}],\"name\":\"ChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"participant1\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"participant2\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"settle_timeout\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"participant1_deposit\",\"type\":\"uint256\"}],\"name\":\"ChannelOpenedAndDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"total_deposit\",\"type\":\"uint256\"}],\"name\":\"ChannelNewDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"closing_participant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"locksroot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"transferred_amount\",\"type\":\"uint256\"}],\"name\":\"ChannelClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"payer_participant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"lockhash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"transferred_amount\",\"type\":\"uint256\"}],\"name\":\"ChannelUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"locksroot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"transferred_amount\",\"type\":\"uint256\"}],\"name\":\"BalanceProofUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"ChannelPunished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"participant1_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"participant2_amount\",\"type\":\"uint256\"}],\"name\":\"ChannelSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"participant1_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"participant2_amount\",\"type\":\"uint256\"}],\"name\":\"ChannelCooperativeSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"participant1\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"participant1_balance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"participant2\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"participant2_balance\",\"type\":\"uint256\"}],\"name\":\"ChannelWithdraw\",\"type\":\"event\"}]"

// TokenNetworkBin is the compiled bytecode used for deploying new contracts.
const TokenNetworkBin = `0x60806040523480156200001157600080fd5b5060405160608062003f49833981016040908152815160208301519190920151600160a060020a03831615156200004757600080fd5b600160a060020a03821615156200005d57600080fd5b600081116200006b57600080fd5b6200007f8364010000000062000177810204565b15156200008b57600080fd5b6200009f8264010000000062000177810204565b1515620000ab57600080fd5b60008054600160a060020a03808616600160a060020a031992831617808455600180548784169416939093179092556002849055604080517f18160ddd000000000000000000000000000000000000000000000000000000008152905192909116916318160ddd9160048082019260209290919082900301818787803b1580156200013557600080fd5b505af11580156200014a573d6000803e3d6000fd5b505050506040513d60208110156200016157600080fd5b5051116200016e57600080fd5b5050506200017f565b6000903b1190565b613dba806200018f6000396000f3006080604052600436106101535763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166324d73a9381146101585780633af973b1146101965780634aaf2b54146101bd5780637709bc78146102415780637a7ebd7b146102835780637c090e4b146102d35780638340f54914610395578063837536b9146103cc5780638568536a1461045057806387234237146105145780638b1ddc531461059e5780638f4ffcb1146105d55780639375cff21461061a5780639fe5b1871461064c578063a570b7d5146106a0578063aaa3dbcc1461076d578063ac133709146107f9578063aef914411461085f578063b32c65c8146108a0578063b9eec014146108b5578063c0ee0b8a14610941578063e11cbf991461097f578063f8658b25146109c1578063f94c9e1314610a98578063fc0c546a14610acc578063fc65697014610ae1575b600080fd5b34801561016457600080fd5b5061016d610b25565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156101a257600080fd5b506101ab610b41565b60408051918252519081900360200190f35b3480156101c957600080fd5b50604080516020600460a43581810135601f810184900484028501840190955284845261023f94823573ffffffffffffffffffffffffffffffffffffffff169460248035956044359560643595608435953695929460c4949201918190840183828082843750949750610b479650505050505050565b005b34801561024d57600080fd5b5061026f73ffffffffffffffffffffffffffffffffffffffff60043516610b5e565b604080519115158252519081900360200190f35b34801561028f57600080fd5b5061029b600435610b66565b6040805167ffffffffffffffff95861681529385166020850152919093168282015260ff909216606082015290519081900360800190f35b3480156102df57600080fd5b50604080516020601f60843560048181013592830184900484028501840190955281845261023f9473ffffffffffffffffffffffffffffffffffffffff81358116956024803590921695604435956064359536959460a4949391019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750610bc99650505050505050565b3480156103a157600080fd5b5061023f73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610f0b565b3480156103d857600080fd5b50604080516020601f60843560048181013592830184900484028501840190955281845261023f9473ffffffffffffffffffffffffffffffffffffffff81358116956024803590921695604435956064359536959460a49493910191908190840183828082843750949750610f1e9650505050505050565b34801561045c57600080fd5b50604080516020601f60843560048181013592830184900484028501840190955281845261023f9473ffffffffffffffffffffffffffffffffffffffff81358116956024803596604435909316956064359536959460a49493919091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506111d19650505050505050565b34801561052057600080fd5b5061052961155f565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561056357818101518382015260200161054b565b50505050905090810190601f1680156105905780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105aa57600080fd5b5061026f73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435611596565b3480156105e157600080fd5b5061026f6004803573ffffffffffffffffffffffffffffffffffffffff90811691602480359260443516916064359182019101356116ac565b34801561062657600080fd5b5061062f61171e565b6040805167ffffffffffffffff9092168252519081900360200190f35b34801561065857600080fd5b50610664600435611723565b6040805195865267ffffffffffffffff94851660208701529284168584015260ff90911660608501529091166080830152519081900360a00190f35b3480156106ac57600080fd5b50604080516020601f60c43560048181013592830184900484028501840190955281845261023f9473ffffffffffffffffffffffffffffffffffffffff8135811695602480359092169560443595606435956084359560a435953695919460e49492939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506117899650505050505050565b34801561077957600080fd5b50604080516020600460a43581810135601f810184900484028501840190955284845261023f94823573ffffffffffffffffffffffffffffffffffffffff169460248035956044359560643567ffffffffffffffff1695608435953695929460c49492019181908401838280828437509497506119969650505050505050565b34801561080557600080fd5b5061082d73ffffffffffffffffffffffffffffffffffffffff60043581169060243516611ba4565b6040805193845267ffffffffffffffff19909216602084015267ffffffffffffffff1682820152519081900360600190f35b34801561086b57600080fd5b5061023f73ffffffffffffffffffffffffffffffffffffffff6004358116906024351667ffffffffffffffff60443516611c32565b3480156108ac57600080fd5b50610529611e16565b3480156108c157600080fd5b50604080516020600460a43581810135601f810184900484028501840190955284845261023f94823573ffffffffffffffffffffffffffffffffffffffff169460248035956044359560643567ffffffffffffffff1695608435953695929460c4949201918190840183828082843750949750611e4d9650505050505050565b34801561094d57600080fd5b5061026f6004803573ffffffffffffffffffffffffffffffffffffffff16906024803591604435918201910135612085565b34801561098b57600080fd5b5061023f73ffffffffffffffffffffffffffffffffffffffff60043581169060243590604435906064351660843560a4356120f4565b3480156109cd57600080fd5b50604080516020601f60c43560048181013592830184900484028501840190955281845261023f9473ffffffffffffffffffffffffffffffffffffffff81358116956024803590921695604435956064359567ffffffffffffffff608435169560a435953695919460e49492939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506124b69650505050505050565b348015610aa457600080fd5b5061066473ffffffffffffffffffffffffffffffffffffffff60043581169060243516612752565b348015610ad857600080fd5b5061016d6127d7565b348015610aed57600080fd5b5061023f73ffffffffffffffffffffffffffffffffffffffff6004358116906024351667ffffffffffffffff604435166064356127f3565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b610b5686338787878787612808565b505050505050565b6000903b1190565b60036020526000908152604090205467ffffffffffffffff80821691680100000000000000008104821691700100000000000000000000000000000000820416907801000000000000000000000000000000000000000000000000900460ff1684565b6000806000806000806000610bde8d8d612c44565b60008181526003602052604090208054919750700100000000000000000000000000000000820467ffffffffffffffff16965093507801000000000000000000000000000000000000000000000000900460ff16600114610c3e57600080fd5b610c4c868e8d8d898e612dc3565b73ffffffffffffffffffffffffffffffffffffffff8e8116911614610c7057600080fd5b610c7e868e8d8d898d612dc3565b73ffffffffffffffffffffffffffffffffffffffff8d8116911614610ca257600080fd5b505073ffffffffffffffffffffffffffffffffffffffff808c166000908152600183016020526040808220928d1682528120805483540197508b88039450908a11610cec57600080fd5b8a8a1115610cf957600080fd5b8a871015610d0657600080fd5b83871015610d1357600080fd5b898b039a508a8260000181905550838160000181905550438360000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8e8c6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610e1957600080fd5b505af1158015610e2d573d6000803e3d6000fd5b505050506040513d6020811015610e4357600080fd5b50511515610e5057600080fd5b85600019167fdc5ff4ab383e66679a382f376c0e80534f51f3f3a398add646422cd81f5f815d8e8d8f88604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390a250505050505050505050505050565b610f19838383336001612f7a565b505050565b600080600080600080610f318b8b612c44565b6000818152600360205260409020805491975093507801000000000000000000000000000000000000000000000000900460ff16600214610f7157600080fd5b73ffffffffffffffffffffffffffffffffffffffff8b1660009081526001848101602052604090912090810154680100000000000000000267ffffffffffffffff191695509150841515610fc457600080fd5b8254610ff39087908b90700100000000000000000000000000000000900467ffffffffffffffff168b8b613127565b73ffffffffffffffffffffffffffffffffffffffff8b811691161461101757600080fd5b5073ffffffffffffffffffffffffffffffffffffffff891660009081526001808401602090815260409283902091840154835167ffffffffffffffff780100000000000000000000000000000000000000000000000092839004169091028183015260288082018d90528451808303909101815260489091019384905280519293909290918291908401908083835b602083106110c55780518252601f1990920191602091820191016110a6565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120600081815260028901909252929020549197505060ff1615159150611113905057600080fd5b6000848152600283016020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557fffffffffffffffff000000000000000000000000000000000000000000000000600186015583548554018555918355815173ffffffffffffffffffffffffffffffffffffffff8e168152915188927fa913b8478dcdecf113bad71030afc079c268eb9abc88e45615f438824127ae0092908290030190a25050505050505050505050565b6000806000806000806111e48c8b612c44565b6000818152600360205260409020805491965093507801000000000000000000000000000000000000000000000000900460ff1660011461122457600080fd5b8254700100000000000000000000000000000000900467ffffffffffffffff169350611255858d8d8d8d898e613287565b73ffffffffffffffffffffffffffffffffffffffff8d811691161461127957600080fd5b611288858d8d8d8d898d613287565b73ffffffffffffffffffffffffffffffffffffffff8b81169116146112ac57600080fd5b505073ffffffffffffffffffffffffffffffffffffffff808b166000908152600180840160209081526040808420948d16845280842080548654868855878601879055868355948201869055898652600390935290842080547fffffffffffffff0000000000000000000000000000000000000000000000000016905591019650908b1115611430576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8d8d6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156113f957600080fd5b505af115801561140d573d6000803e3d6000fd5b505050506040513d602081101561142357600080fd5b5051151561143057600080fd5b60008911156114ef5760008054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e81166004830152602482018e90529151919092169263a9059cbb92604480820193602093909283900390910190829087803b1580156114b857600080fd5b505af11580156114cc573d6000803e3d6000fd5b505050506040513d60208110156114e257600080fd5b505115156114ef57600080fd5b8a890186146114fd57600080fd5b8a86101561150a57600080fd5b8886101561151757600080fd5b604080518c8152602081018b9052815187927ffb2f4bc0fb2e0f1001f78d15e81a2e1981f262d31e8bd72309e26cc63bf7bb02928290030190a2505050505050505050505050565b60408051808201909152601d81527f19536d61727452616964656e205369676e6564204d6573736167653a0a000000602082015281565b60008060008060006115a88888612c44565b600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8d168452600180820184529382902093840154825167ffffffffffffffff780100000000000000000000000000000000000000000000000092839004169091028185015260288082018d9052835180830390910181526048909101928390528051959850909650929450919282918401908083835b602083106116625780518252601f199092019160209182019101611643565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060009081526002969096019052509092205460ff169998505050505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff8581169116146116d457600080fd5b611712868685858080601f01602080910402602001604051908101604052809392919081815260200183838082843750600194506134269350505050565b50600195945050505050565b600581565b600081815260036020526040902054909167ffffffffffffffff680100000000000000008304811692700100000000000000000000000000000000810482169260ff78010000000000000000000000000000000000000000000000008304169290911690565b60008060006117988b8b612c44565b60008181526003602090815260409182902082518084018452601d8082527f19536d61727452616964656e205369676e6564204d6573736167653a0a00000082850190815283546002549651979950939750919533958f958f958f958c9570010000000000000000000000000000000090920467ffffffffffffffff16949092019182918083835b6020831061183f5780518252601f199092019160209182019101611820565b51815160209384036101000a600019018019909216911617905273ffffffffffffffffffffffffffffffffffffffff9b909b166c010000000000000000000000000292019182525060148101979097525060348601949094526054850192909252607484015267ffffffffffffffff167801000000000000000000000000000000000000000000000000026094830152609c8083019190915260408051808403909201825260bc90920191829052805190935090918291908401908083835b6020831061191d5780518252601f1990920191602091820191016118fe565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902092506119568385613489565b73ffffffffffffffffffffffffffffffffffffffff8b811691161461197a57600080fd5b6119898b8b8b8b8b8b8b612808565b5050505050505050505050565b60008060006119a58933612c44565b600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8e16845260018101909252909120815492955090935091507801000000000000000000000000000000000000000000000000900460ff16600214611a0e57600080fd5b8154436801000000000000000090910467ffffffffffffffff161015611a3357600080fd5b600181015467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690871611611a6d57600080fd5b611a94838989898660000160109054906101000a900467ffffffffffffffff168a8a613576565b73ffffffffffffffffffffffffffffffffffffffff8a8116911614611ab857600080fd5b611ac288886136a5565b60018201805467ffffffffffffffff8916780100000000000000000000000000000000000000000000000002680100000000000000009093047fffffffffffffffff0000000000000000000000000000000000000000000000009091161777ffffffffffffffffffffffffffffffffffffffffffffffff169190911790556040805173ffffffffffffffffffffffffffffffffffffffff8b168152602081018990528082018a9052905184917f910c9237f4197a18340110a181e8fb775496506a007a94b46f9f80f2a35918f9919081900360600190a2505050505050505050565b600080600080600080611bb78888612c44565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff9b909b16835260019a8b01909152902080549801549798680100000000000000008902987801000000000000000000000000000000000000000000000000900467ffffffffffffffff16975095505050505050565b6000808260068167ffffffffffffffff1610158015611c5e5750622932e08167ffffffffffffffff1611155b1515611c6957600080fd5b73ffffffffffffffffffffffffffffffffffffffff86161515611c8b57600080fd5b73ffffffffffffffffffffffffffffffffffffffff85161515611cad57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8681169086161415611cd357600080fd5b611cdd8686612c44565b6000818152600360205260409020805491945092507801000000000000000000000000000000000000000000000000900460ff1615611d1b57600080fd5b81547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff4367ffffffffffffffff908116700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff91881667ffffffffffffffff19909416841791909116171678010000000000000000000000000000000000000000000000001783556040805173ffffffffffffffffffffffffffffffffffffffff808a16825288166020820152808201929092525184917f4d4097deeecde59dede1bb370eb147fc3fa969b7b6a6f89f95526635328e86df919081900360600190a2505050505050565b60408051808201909152600581527f302e342e5f000000000000000000000000000000000000000000000000000000602082015281565b600080600080611e5d338b612c44565b6000818152600360205260409020805491955092507801000000000000000000000000000000000000000000000000900460ff16600114611e9d57600080fd5b81547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff7fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff909116780200000000000000000000000000000000000000000000000017908116680100000000000000004367ffffffffffffffff9384160183160217835560009088161115612036575073ffffffffffffffffffffffffffffffffffffffff8916600090815260018201602052604090208154611f869085908b908b908b90700100000000000000000000000000000000900467ffffffffffffffff168b8b613576565b925073ffffffffffffffffffffffffffffffffffffffff8a811690841614611fad57600080fd5b611fb789896136a5565b60018201805467ffffffffffffffff8a16780100000000000000000000000000000000000000000000000002680100000000000000009093047fffffffffffffffff0000000000000000000000000000000000000000000000009091161777ffffffffffffffffffffffffffffffffffffffffffffffff169190911790555b60408051338152602081018a90528082018b9052905185917f69610baaace24c039f891a11b42c0b1df1496ab0db38b0c4ee4ed33d6d53da1a919081900360600190a250505050505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff1633146120aa57600080fd5b6120e960008585858080601f01602080910402602001604051908101604052809392919081815260200183838082843750600094506134269350505050565b506001949350505050565b6000806000806000806121078c8a612c44565b6000818152600360205260409020805491955093507801000000000000000000000000000000000000000000000000900460ff1660021461214757600080fd5b82544367ffffffffffffffff6801000000000000000090920482166005019091161061217257600080fd5b505073ffffffffffffffffffffffffffffffffffffffff808b166000908152600183016020526040808220928a16825290206121ae8b8b6136a5565b6001830154680100000000000000000267ffffffffffffffff199081169116146121d757600080fd5b6121e188886136a5565b6001820154680100000000000000000267ffffffffffffffff1990811691161461220a57600080fd5b805482548981018d81039850910195508b111561222657600095505b6122308686613744565b73ffffffffffffffffffffffffffffffffffffffff808e1660009081526001808701602090815260408084208481558301849055938e1683528383208381559091018290558782526003905290812080547fffffffffffffff0000000000000000000000000000000000000000000000000016905581870399509096508611156123af576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8d886040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561237857600080fd5b505af115801561238c573d6000803e3d6000fd5b505050506040513d60208110156123a257600080fd5b505115156123af57600080fd5b600088111561246e5760008054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d81166004830152602482018d90529151919092169263a9059cbb92604480820193602093909283900390910190829087803b15801561243757600080fd5b505af115801561244b573d6000803e3d6000fd5b505050506040513d602081101561246157600080fd5b5051151561246e57600080fd5b60408051878152602081018a9052815186927ff94fb5c0628a82dc90648e8dc5e983f632633b0d26603d64e8cc042ca0790aa4928290030190a2505050505050505050505050565b6000806000806124c68c8c612c44565b935060036000856000191660001916815260200190815260200160002091508160010160008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508160000160189054906101000a900460ff1660ff16600214151561254b57600080fd5b815468010000000000000000900467ffffffffffffffff1692504383101561257257600080fd5b8154600267ffffffffffffffff9182160484031643101561259257600080fd5b600181015467ffffffffffffffff78010000000000000000000000000000000000000000000000009091048116908916116125cc57600080fd5b6125f4848b8b8b8660000160109054906101000a900467ffffffffffffffff168c8c8c61375c565b73ffffffffffffffffffffffffffffffffffffffff8c811691161461261857600080fd5b61263f848b8b8b8660000160109054906101000a900467ffffffffffffffff168c8c613576565b73ffffffffffffffffffffffffffffffffffffffff8d811691161461266357600080fd5b61266d8a8a6136a5565b60018201805467ffffffffffffffff8b16780100000000000000000000000000000000000000000000000002680100000000000000009093047fffffffffffffffff0000000000000000000000000000000000000000000000009091161777ffffffffffffffffffffffffffffffffffffffffffffffff169190911790556040805173ffffffffffffffffffffffffffffffffffffffff8e168152602081018b90528082018c9052905185917f910c9237f4197a18340110a181e8fb775496506a007a94b46f9f80f2a35918f9919081900360600190a2505050505050505050505050565b60008060008060008060006127678989612c44565b600081815260036020526040902054909a67ffffffffffffffff68010000000000000000830481169b50700100000000000000000000000000000000830481169a5060ff780100000000000000000000000000000000000000000000000084041699509091169650945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b61280284848484336001613923565b50505050565b600080600080600080600061281d8e8e612c44565b965060036000886000191660001916815260200190815260200160002091508160010160008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050438260000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16101515156128b057600080fd5b81547801000000000000000000000000000000000000000000000000900460ff166002146128dd57600080fd5b600154604080517fc1f62946000000000000000000000000000000000000000000000000000000008152600481018c9052905173ffffffffffffffffffffffffffffffffffffffff9092169163c1f62946916024808201926020929091908290030181600087803b15801561295157600080fd5b505af1158015612965573d6000803e3d6000fd5b505050506040513d602081101561297b57600080fd5b5051935060008411801561298f57508a8411155b151561299a57600080fd5b6040805160208082018e90528183018d905260608083018d905283518084039091018152608090920192839052815191929182918401908083835b602083106129f45780518252601f1990920191602091820191016129d5565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209450612a2d8589613c1a565b9250612a398c846136a5565b6001820154680100000000000000000267ffffffffffffffff19908116911614612a6257600080fd5b60018101546040805167ffffffffffffffff78010000000000000000000000000000000000000000000000009384900416909202602080840191909152602880840189905282518085039091018152604890930191829052825182918401908083835b60208310612ae45780518252601f199092019160209182019101612ac5565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120600081815260028801909252929020549199505060ff16159150612b31905057600080fd5b6000868152600282016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790559a89019a612b788c846136a5565b8160010160006101000a81548177ffffffffffffffffffffffffffffffffffffffffffffffff0219169083680100000000000000009004021790555086600019167f9e3b094fde58f3a83bd8b77d0a995fdb71f3169c6fa7e6d386e9f5902841e5ff8f878f604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018360001916600019168152602001828152602001935050505060405180910390a25050505050505050505050505050565b60008173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161015612d35576040805173ffffffffffffffffffffffffffffffffffffffff8581166c0100000000000000000000000090810260208085019190915291861681026034840152300260488301528251808303603c018152605c90920192839052815191929182918401908083835b60208310612d015780518252601f199092019160209182019101612ce2565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050612dbd565b604080516c0100000000000000000000000073ffffffffffffffffffffffffffffffffffffffff808616820260208085019190915290871682026034840152309190910260488301528251603c818403018152605c909201928390528151919291829184019080838360208310612d015780518252601f199092019160209182019101612ce2565b92915050565b6000806040805190810160405280601d81526020017f19536d61727452616964656e205369676e6564204d6573736167653a0a0000008152508787878b886002546040516020018088805190602001908083835b60208310612e365780518252601f199092019160209182019101612e17565b6001836020036101000a0380198251168184511680821785525050505050509050018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140186815260200185815260200184600019166000191681526020018367ffffffffffffffff1667ffffffffffffffff1678010000000000000000000000000000000000000000000000000281526008018281526020019750505050505050506040516020818303038152906040526040518082805190602001908083835b60208310612f355780518252601f199092019160209182019101612f16565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050612f6e8184613489565b98975050505050505050565b6000808080808711612f8b57600080fd5b612f958989612c44565b600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8e16845260018101909252909120805496509194509250905084156130965760008054604080517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152306024830152604482018c9052915191909216926323b872dd92606480820193602093909283900390910190829087803b15801561305f57600080fd5b505af1158015613073573d6000803e3d6000fd5b505050506040513d602081101561308957600080fd5b5051151561309657600080fd5b81547801000000000000000000000000000000000000000000000000900460ff166001146130c357600080fd5b9286018084556040805173ffffffffffffffffffffffffffffffffffffffff8b16815260208101839052815192959285927f0346e981e2bfa2366dc2307a8f1fa24779830a01121b1275fe565c6b98bb4d34928290030190a2505050505050505050565b6000806040805190810160405280601d81526020017f19536d61727452616964656e205369676e6564204d6573736167653a0a000000815250868887600254886040516020018087805190602001908083835b602083106131995780518252601f19909201916020918201910161317a565b51815160209384036101000a6000190180199092169116179052920197885250868101959095525067ffffffffffffffff92909216780100000000000000000000000000000000000000000000000002604080860191909152604885019190915260688085019290925280518085039092018252608890930192839052805190935082918401908083835b602083106132435780518252601f199092019160209182019101613224565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061327c8184613489565b979650505050505050565b6000806040805190810160405280601d81526020017f19536d61727452616964656e205369676e6564204d6573736167653a0a000000815250888888888d896002546040516020018089805190602001908083835b602083106132fb5780518252601f1990920191602091820191016132dc565b51815160209384036101000a600019018019909216911617905273ffffffffffffffffffffffffffffffffffffffff9b8c166c0100000000000000000000000090810292909401918252601482019a909a5297909916026034870152506048850193909352606884019190915267ffffffffffffffff16780100000000000000000000000000000000000000000000000002608883015260908083019190915260408051808403909201825260b090920191829052805190945090925082918401908083835b602083106133e05780518252601f1990920191602091820191016133c1565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090506134198184613489565b9998505050505050505050565b602082015160008080600184141561345b5761344186613d6b565b919450925090506134568383838a8c8a613923565b61347f565b83600214156101535761346d86613d7f565b90935091506134568383898b89612f7a565b5050505050505050565b6000806000808451604114151561349f57600080fd5b50505060208201516040830151606084015160001a601b60ff821610156134c457601b015b8060ff16601b14806134d957508060ff16601c145b15156134e457600080fd5b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af115801561353e573d6000803e3d6000fd5b5050604051601f19015194505073ffffffffffffffffffffffffffffffffffffffff8416151561356d57600080fd5b50505092915050565b6000806040805190810160405280601d81526020017f19536d61727452616964656e205369676e6564204d6573736167653a0a000000815250888888878d8a6002546040516020018089805190602001908083835b602083106135ea5780518252601f1990920191602091820191016135cb565b51815160001960209485036101000a019081169019919091161790529201998a52508881019790975250780100000000000000000000000000000000000000000000000067ffffffffffffffff95861681026040808a01919091526048890195909552606888019390935293160260888501526090808501929092528051808503909201825260b09093019283905280519093508291840190808383602083106133e05780518252601f1990920191602091820191016133c1565b6000811580156136b3575082155b156136c057506000612dbd565b604080516020808201859052818301869052825180830384018152606090920192839052815191929182918401908083835b602083106137115780518252601f1990920191602091820191016136f2565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091209695505050505050565b60008183116137535782613755565b815b9392505050565b6000806040805190810160405280601d81526020017f19536d61727452616964656e205369676e6564204d6573736167653a0a000000815250898989888e8b6002548b604051602001808a805190602001908083835b602083106137d15780518252601f1990920191602091820191016137b2565b51815160001960209485036101000a0190811690199190911617905292018b81528083018b9052780100000000000000000000000000000000000000000000000067ffffffffffffffff808c1682026040840152604883018b9052606883018a9052881602608882015260908101869052845160b090910192850191508083835b602083106138715780518252601f199092019160209182019101613852565b6001836020036101000a03801982511681845116808217855250505050505090500199505050505050505050506040516020818303038152906040526040518082805190602001908083835b602083106138dc5780518252601f1990920191602091820191016138bd565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090506139158184613489565b9a9950505050505050505050565b60008060008660068167ffffffffffffffff16101580156139515750622932e08167ffffffffffffffff1611155b151561395c57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8a16151561397e57600080fd5b73ffffffffffffffffffffffffffffffffffffffff891615156139a057600080fd5b73ffffffffffffffffffffffffffffffffffffffff8a8116908a1614156139c657600080fd5b600087116139d357600080fd5b6139dd8a8a612c44565b600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8f16845260018101909252909120815492965090945092507801000000000000000000000000000000000000000000000000900460ff1615613a4457600080fd5b82547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff4367ffffffffffffffff908116700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff918c1667ffffffffffffffff199094169390931716919091171678010000000000000000000000000000000000000000000000001783558415613ba15760008054604080517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152306024830152604482018c9052915191909216926323b872dd92606480820193602093909283900390910190829087803b158015613b6a57600080fd5b505af1158015613b7e573d6000803e3d6000fd5b505050506040513d6020811015613b9457600080fd5b50511515613ba157600080fd5b8682556040805173ffffffffffffffffffffffffffffffffffffffff808d1682528b16602082015267ffffffffffffffff8a168183015260608101899052905185917fcac76648b0a531becb6e54db5fe838853fdc47ef130aab3566114ee7c739d0a0919081900360800190a250505050505050505050565b600080600060208451811515613c2c57fe5b0615613c3757600080fd5b602091505b83518211613d6257508281015180851015613cd657604080516020808201889052818301849052825180830384018152606090920192839052815191929182918401908083835b60208310613ca25780518252601f199092019160209182019101613c83565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209450613d57565b604080516020808201849052818301889052825180830384018152606090920192839052815191929182918401908083835b60208310613d275780518252601f199092019160209182019101613d08565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902094505b602082019150613c3c565b50929392505050565b604081015160608201516080909201519092565b604081015160608201519150915600a165627a7a723058208976b94ce2133c9760307de2a94560bcac6dd4a4d450a369447563767b1f45920029`

// DeployTokenNetwork deploys a new Ethereum contract, binding an instance of TokenNetwork to it.
func DeployTokenNetwork(auth *bind.TransactOpts, backend bind.ContractBackend, _token_address common.Address, _secret_registry common.Address, _chain_id *big.Int) (common.Address, *types.Transaction, *TokenNetwork, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenNetworkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenNetworkBin), backend, _token_address, _secret_registry, _chain_id)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenNetwork{TokenNetworkCaller: TokenNetworkCaller{contract: contract}, TokenNetworkTransactor: TokenNetworkTransactor{contract: contract}, TokenNetworkFilterer: TokenNetworkFilterer{contract: contract}}, nil
}

// TokenNetwork is an auto generated Go binding around an Ethereum contract.
type TokenNetwork struct {
	TokenNetworkCaller     // Read-only binding to the contract
	TokenNetworkTransactor // Write-only binding to the contract
	TokenNetworkFilterer   // Log filterer for contract events
}

// TokenNetworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenNetworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenNetworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenNetworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenNetworkSession struct {
	Contract     *TokenNetwork     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenNetworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenNetworkCallerSession struct {
	Contract *TokenNetworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenNetworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenNetworkTransactorSession struct {
	Contract     *TokenNetworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenNetworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenNetworkRaw struct {
	Contract *TokenNetwork // Generic contract binding to access the raw methods on
}

// TokenNetworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenNetworkCallerRaw struct {
	Contract *TokenNetworkCaller // Generic read-only contract binding to access the raw methods on
}

// TokenNetworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenNetworkTransactorRaw struct {
	Contract *TokenNetworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenNetwork creates a new instance of TokenNetwork, bound to a specific deployed contract.
func NewTokenNetwork(address common.Address, backend bind.ContractBackend) (*TokenNetwork, error) {
	contract, err := bindTokenNetwork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenNetwork{TokenNetworkCaller: TokenNetworkCaller{contract: contract}, TokenNetworkTransactor: TokenNetworkTransactor{contract: contract}, TokenNetworkFilterer: TokenNetworkFilterer{contract: contract}}, nil
}

// NewTokenNetworkCaller creates a new read-only instance of TokenNetwork, bound to a specific deployed contract.
func NewTokenNetworkCaller(address common.Address, caller bind.ContractCaller) (*TokenNetworkCaller, error) {
	contract, err := bindTokenNetwork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkCaller{contract: contract}, nil
}

// NewTokenNetworkTransactor creates a new write-only instance of TokenNetwork, bound to a specific deployed contract.
func NewTokenNetworkTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenNetworkTransactor, error) {
	contract, err := bindTokenNetwork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkTransactor{contract: contract}, nil
}

// NewTokenNetworkFilterer creates a new log filterer instance of TokenNetwork, bound to a specific deployed contract.
func NewTokenNetworkFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenNetworkFilterer, error) {
	contract, err := bindTokenNetwork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkFilterer{contract: contract}, nil
}

// bindTokenNetwork binds a generic wrapper to an already deployed contract.
func bindTokenNetwork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenNetworkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenNetwork *TokenNetworkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenNetwork.Contract.TokenNetworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenNetwork *TokenNetworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenNetwork.Contract.TokenNetworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenNetwork *TokenNetworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenNetwork.Contract.TokenNetworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenNetwork *TokenNetworkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenNetwork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenNetwork *TokenNetworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenNetwork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenNetwork *TokenNetworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenNetwork.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetwork *TokenNetworkCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "chain_id")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetwork *TokenNetworkSession) ChainId() (*big.Int, error) {
	return _TokenNetwork.Contract.ChainId(&_TokenNetwork.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetwork *TokenNetworkCallerSession) ChainId() (*big.Int, error) {
	return _TokenNetwork.Contract.ChainId(&_TokenNetwork.CallOpts)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels( bytes32) constant returns(settle_timeout uint64, settle_block_number uint64, open_block_number uint64, state uint8)
func (_TokenNetwork *TokenNetworkCaller) Channels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	SettleTimeout     uint64
	SettleBlockNumber uint64
	OpenBlockNumber   uint64
	State             uint8
}, error) {
	ret := new(struct {
		SettleTimeout     uint64
		SettleBlockNumber uint64
		OpenBlockNumber   uint64
		State             uint8
	})
	out := ret
	err := _TokenNetwork.contract.Call(opts, out, "channels", arg0)
	return *ret, err
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels( bytes32) constant returns(settle_timeout uint64, settle_block_number uint64, open_block_number uint64, state uint8)
func (_TokenNetwork *TokenNetworkSession) Channels(arg0 [32]byte) (struct {
	SettleTimeout     uint64
	SettleBlockNumber uint64
	OpenBlockNumber   uint64
	State             uint8
}, error) {
	return _TokenNetwork.Contract.Channels(&_TokenNetwork.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels( bytes32) constant returns(settle_timeout uint64, settle_block_number uint64, open_block_number uint64, state uint8)
func (_TokenNetwork *TokenNetworkCallerSession) Channels(arg0 [32]byte) (struct {
	SettleTimeout     uint64
	SettleBlockNumber uint64
	OpenBlockNumber   uint64
	State             uint8
}, error) {
	return _TokenNetwork.Contract.Channels(&_TokenNetwork.CallOpts, arg0)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetwork *TokenNetworkCaller) ContractExists(opts *bind.CallOpts, contract_address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "contractExists", contract_address)
	return *ret0, err
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetwork *TokenNetworkSession) ContractExists(contract_address common.Address) (bool, error) {
	return _TokenNetwork.Contract.ContractExists(&_TokenNetwork.CallOpts, contract_address)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetwork *TokenNetworkCallerSession) ContractExists(contract_address common.Address) (bool, error) {
	return _TokenNetwork.Contract.ContractExists(&_TokenNetwork.CallOpts, contract_address)
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetwork *TokenNetworkCaller) ContractVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "contract_version")
	return *ret0, err
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetwork *TokenNetworkSession) ContractVersion() (string, error) {
	return _TokenNetwork.Contract.ContractVersion(&_TokenNetwork.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetwork *TokenNetworkCallerSession) ContractVersion() (string, error) {
	return _TokenNetwork.Contract.ContractVersion(&_TokenNetwork.CallOpts)
}

// GetChannelInfo is a free data retrieval call binding the contract method 0xf94c9e13.
//
// Solidity: function getChannelInfo(participant1 address, participant2 address) constant returns(bytes32, uint64, uint64, uint8, uint64)
func (_TokenNetwork *TokenNetworkCaller) GetChannelInfo(opts *bind.CallOpts, participant1 common.Address, participant2 common.Address) ([32]byte, uint64, uint64, uint8, uint64, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(uint64)
		ret2 = new(uint64)
		ret3 = new(uint8)
		ret4 = new(uint64)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _TokenNetwork.contract.Call(opts, out, "getChannelInfo", participant1, participant2)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetChannelInfo is a free data retrieval call binding the contract method 0xf94c9e13.
//
// Solidity: function getChannelInfo(participant1 address, participant2 address) constant returns(bytes32, uint64, uint64, uint8, uint64)
func (_TokenNetwork *TokenNetworkSession) GetChannelInfo(participant1 common.Address, participant2 common.Address) ([32]byte, uint64, uint64, uint8, uint64, error) {
	return _TokenNetwork.Contract.GetChannelInfo(&_TokenNetwork.CallOpts, participant1, participant2)
}

// GetChannelInfo is a free data retrieval call binding the contract method 0xf94c9e13.
//
// Solidity: function getChannelInfo(participant1 address, participant2 address) constant returns(bytes32, uint64, uint64, uint8, uint64)
func (_TokenNetwork *TokenNetworkCallerSession) GetChannelInfo(participant1 common.Address, participant2 common.Address) ([32]byte, uint64, uint64, uint8, uint64, error) {
	return _TokenNetwork.Contract.GetChannelInfo(&_TokenNetwork.CallOpts, participant1, participant2)
}

// GetChannelInfoByChannelIdentifier is a free data retrieval call binding the contract method 0x9fe5b187.
//
// Solidity: function getChannelInfoByChannelIdentifier(channel_identifier bytes32) constant returns(bytes32, uint64, uint64, uint8, uint64)
func (_TokenNetwork *TokenNetworkCaller) GetChannelInfoByChannelIdentifier(opts *bind.CallOpts, channel_identifier [32]byte) ([32]byte, uint64, uint64, uint8, uint64, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(uint64)
		ret2 = new(uint64)
		ret3 = new(uint8)
		ret4 = new(uint64)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _TokenNetwork.contract.Call(opts, out, "getChannelInfoByChannelIdentifier", channel_identifier)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetChannelInfoByChannelIdentifier is a free data retrieval call binding the contract method 0x9fe5b187.
//
// Solidity: function getChannelInfoByChannelIdentifier(channel_identifier bytes32) constant returns(bytes32, uint64, uint64, uint8, uint64)
func (_TokenNetwork *TokenNetworkSession) GetChannelInfoByChannelIdentifier(channel_identifier [32]byte) ([32]byte, uint64, uint64, uint8, uint64, error) {
	return _TokenNetwork.Contract.GetChannelInfoByChannelIdentifier(&_TokenNetwork.CallOpts, channel_identifier)
}

// GetChannelInfoByChannelIdentifier is a free data retrieval call binding the contract method 0x9fe5b187.
//
// Solidity: function getChannelInfoByChannelIdentifier(channel_identifier bytes32) constant returns(bytes32, uint64, uint64, uint8, uint64)
func (_TokenNetwork *TokenNetworkCallerSession) GetChannelInfoByChannelIdentifier(channel_identifier [32]byte) ([32]byte, uint64, uint64, uint8, uint64, error) {
	return _TokenNetwork.Contract.GetChannelInfoByChannelIdentifier(&_TokenNetwork.CallOpts, channel_identifier)
}

// GetChannelParticipantInfo is a free data retrieval call binding the contract method 0xac133709.
//
// Solidity: function getChannelParticipantInfo(participant address, partner address) constant returns(uint256, bytes24, uint64)
func (_TokenNetwork *TokenNetworkCaller) GetChannelParticipantInfo(opts *bind.CallOpts, participant common.Address, partner common.Address) (*big.Int, [24]byte, uint64, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([24]byte)
		ret2 = new(uint64)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _TokenNetwork.contract.Call(opts, out, "getChannelParticipantInfo", participant, partner)
	return *ret0, *ret1, *ret2, err
}

// GetChannelParticipantInfo is a free data retrieval call binding the contract method 0xac133709.
//
// Solidity: function getChannelParticipantInfo(participant address, partner address) constant returns(uint256, bytes24, uint64)
func (_TokenNetwork *TokenNetworkSession) GetChannelParticipantInfo(participant common.Address, partner common.Address) (*big.Int, [24]byte, uint64, error) {
	return _TokenNetwork.Contract.GetChannelParticipantInfo(&_TokenNetwork.CallOpts, participant, partner)
}

// GetChannelParticipantInfo is a free data retrieval call binding the contract method 0xac133709.
//
// Solidity: function getChannelParticipantInfo(participant address, partner address) constant returns(uint256, bytes24, uint64)
func (_TokenNetwork *TokenNetworkCallerSession) GetChannelParticipantInfo(participant common.Address, partner common.Address) (*big.Int, [24]byte, uint64, error) {
	return _TokenNetwork.Contract.GetChannelParticipantInfo(&_TokenNetwork.CallOpts, participant, partner)
}

// PunishBlockNumber is a free data retrieval call binding the contract method 0x9375cff2.
//
// Solidity: function punish_block_number() constant returns(uint64)
func (_TokenNetwork *TokenNetworkCaller) PunishBlockNumber(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "punish_block_number")
	return *ret0, err
}

// PunishBlockNumber is a free data retrieval call binding the contract method 0x9375cff2.
//
// Solidity: function punish_block_number() constant returns(uint64)
func (_TokenNetwork *TokenNetworkSession) PunishBlockNumber() (uint64, error) {
	return _TokenNetwork.Contract.PunishBlockNumber(&_TokenNetwork.CallOpts)
}

// PunishBlockNumber is a free data retrieval call binding the contract method 0x9375cff2.
//
// Solidity: function punish_block_number() constant returns(uint64)
func (_TokenNetwork *TokenNetworkCallerSession) PunishBlockNumber() (uint64, error) {
	return _TokenNetwork.Contract.PunishBlockNumber(&_TokenNetwork.CallOpts)
}

// QueryUnlockedLocks is a free data retrieval call binding the contract method 0x8b1ddc53.
//
// Solidity: function queryUnlockedLocks(participant address, partner address, lockhash bytes32) constant returns(bool)
func (_TokenNetwork *TokenNetworkCaller) QueryUnlockedLocks(opts *bind.CallOpts, participant common.Address, partner common.Address, lockhash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "queryUnlockedLocks", participant, partner, lockhash)
	return *ret0, err
}

// QueryUnlockedLocks is a free data retrieval call binding the contract method 0x8b1ddc53.
//
// Solidity: function queryUnlockedLocks(participant address, partner address, lockhash bytes32) constant returns(bool)
func (_TokenNetwork *TokenNetworkSession) QueryUnlockedLocks(participant common.Address, partner common.Address, lockhash [32]byte) (bool, error) {
	return _TokenNetwork.Contract.QueryUnlockedLocks(&_TokenNetwork.CallOpts, participant, partner, lockhash)
}

// QueryUnlockedLocks is a free data retrieval call binding the contract method 0x8b1ddc53.
//
// Solidity: function queryUnlockedLocks(participant address, partner address, lockhash bytes32) constant returns(bool)
func (_TokenNetwork *TokenNetworkCallerSession) QueryUnlockedLocks(participant common.Address, partner common.Address, lockhash [32]byte) (bool, error) {
	return _TokenNetwork.Contract.QueryUnlockedLocks(&_TokenNetwork.CallOpts, participant, partner, lockhash)
}

// SecretRegistry is a free data retrieval call binding the contract method 0x24d73a93.
//
// Solidity: function secret_registry() constant returns(address)
func (_TokenNetwork *TokenNetworkCaller) SecretRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "secret_registry")
	return *ret0, err
}

// SecretRegistry is a free data retrieval call binding the contract method 0x24d73a93.
//
// Solidity: function secret_registry() constant returns(address)
func (_TokenNetwork *TokenNetworkSession) SecretRegistry() (common.Address, error) {
	return _TokenNetwork.Contract.SecretRegistry(&_TokenNetwork.CallOpts)
}

// SecretRegistry is a free data retrieval call binding the contract method 0x24d73a93.
//
// Solidity: function secret_registry() constant returns(address)
func (_TokenNetwork *TokenNetworkCallerSession) SecretRegistry() (common.Address, error) {
	return _TokenNetwork.Contract.SecretRegistry(&_TokenNetwork.CallOpts)
}

// SignaturePrefix is a free data retrieval call binding the contract method 0x87234237.
//
// Solidity: function signature_prefix() constant returns(string)
func (_TokenNetwork *TokenNetworkCaller) SignaturePrefix(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "signature_prefix")
	return *ret0, err
}

// SignaturePrefix is a free data retrieval call binding the contract method 0x87234237.
//
// Solidity: function signature_prefix() constant returns(string)
func (_TokenNetwork *TokenNetworkSession) SignaturePrefix() (string, error) {
	return _TokenNetwork.Contract.SignaturePrefix(&_TokenNetwork.CallOpts)
}

// SignaturePrefix is a free data retrieval call binding the contract method 0x87234237.
//
// Solidity: function signature_prefix() constant returns(string)
func (_TokenNetwork *TokenNetworkCallerSession) SignaturePrefix() (string, error) {
	return _TokenNetwork.Contract.SignaturePrefix(&_TokenNetwork.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenNetwork *TokenNetworkCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenNetwork *TokenNetworkSession) Token() (common.Address, error) {
	return _TokenNetwork.Contract.Token(&_TokenNetwork.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenNetwork *TokenNetworkCallerSession) Token() (common.Address, error) {
	return _TokenNetwork.Contract.Token(&_TokenNetwork.CallOpts)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xb9eec014.
//
// Solidity: function closeChannel(partner address, transferred_amount uint256, locksroot bytes32, nonce uint64, additional_hash bytes32, signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) CloseChannel(opts *bind.TransactOpts, partner common.Address, transferred_amount *big.Int, locksroot [32]byte, nonce uint64, additional_hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "closeChannel", partner, transferred_amount, locksroot, nonce, additional_hash, signature)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xb9eec014.
//
// Solidity: function closeChannel(partner address, transferred_amount uint256, locksroot bytes32, nonce uint64, additional_hash bytes32, signature bytes) returns()
func (_TokenNetwork *TokenNetworkSession) CloseChannel(partner common.Address, transferred_amount *big.Int, locksroot [32]byte, nonce uint64, additional_hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.CloseChannel(&_TokenNetwork.TransactOpts, partner, transferred_amount, locksroot, nonce, additional_hash, signature)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xb9eec014.
//
// Solidity: function closeChannel(partner address, transferred_amount uint256, locksroot bytes32, nonce uint64, additional_hash bytes32, signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) CloseChannel(partner common.Address, transferred_amount *big.Int, locksroot [32]byte, nonce uint64, additional_hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.CloseChannel(&_TokenNetwork.TransactOpts, partner, transferred_amount, locksroot, nonce, additional_hash, signature)
}

// CooperativeSettle is a paid mutator transaction binding the contract method 0x8568536a.
//
// Solidity: function cooperativeSettle(participant1 address, participant1_balance uint256, participant2 address, participant2_balance uint256, participant1_signature bytes, participant2_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) CooperativeSettle(opts *bind.TransactOpts, participant1 common.Address, participant1_balance *big.Int, participant2 common.Address, participant2_balance *big.Int, participant1_signature []byte, participant2_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "cooperativeSettle", participant1, participant1_balance, participant2, participant2_balance, participant1_signature, participant2_signature)
}

// CooperativeSettle is a paid mutator transaction binding the contract method 0x8568536a.
//
// Solidity: function cooperativeSettle(participant1 address, participant1_balance uint256, participant2 address, participant2_balance uint256, participant1_signature bytes, participant2_signature bytes) returns()
func (_TokenNetwork *TokenNetworkSession) CooperativeSettle(participant1 common.Address, participant1_balance *big.Int, participant2 common.Address, participant2_balance *big.Int, participant1_signature []byte, participant2_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.CooperativeSettle(&_TokenNetwork.TransactOpts, participant1, participant1_balance, participant2, participant2_balance, participant1_signature, participant2_signature)
}

// CooperativeSettle is a paid mutator transaction binding the contract method 0x8568536a.
//
// Solidity: function cooperativeSettle(participant1 address, participant1_balance uint256, participant2 address, participant2_balance uint256, participant1_signature bytes, participant2_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) CooperativeSettle(participant1 common.Address, participant1_balance *big.Int, participant2 common.Address, participant2_balance *big.Int, participant1_signature []byte, participant2_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.CooperativeSettle(&_TokenNetwork.TransactOpts, participant1, participant1_balance, participant2, participant2_balance, participant1_signature, participant2_signature)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(participant address, partner address, amount uint256) returns()
func (_TokenNetwork *TokenNetworkTransactor) Deposit(opts *bind.TransactOpts, participant common.Address, partner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "deposit", participant, partner, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(participant address, partner address, amount uint256) returns()
func (_TokenNetwork *TokenNetworkSession) Deposit(participant common.Address, partner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.Contract.Deposit(&_TokenNetwork.TransactOpts, participant, partner, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(participant address, partner address, amount uint256) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) Deposit(participant common.Address, partner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.Contract.Deposit(&_TokenNetwork.TransactOpts, participant, partner, amount)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xaef91441.
//
// Solidity: function openChannel(participant1 address, participant2 address, settle_timeout uint64) returns()
func (_TokenNetwork *TokenNetworkTransactor) OpenChannel(opts *bind.TransactOpts, participant1 common.Address, participant2 common.Address, settle_timeout uint64) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "openChannel", participant1, participant2, settle_timeout)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xaef91441.
//
// Solidity: function openChannel(participant1 address, participant2 address, settle_timeout uint64) returns()
func (_TokenNetwork *TokenNetworkSession) OpenChannel(participant1 common.Address, participant2 common.Address, settle_timeout uint64) (*types.Transaction, error) {
	return _TokenNetwork.Contract.OpenChannel(&_TokenNetwork.TransactOpts, participant1, participant2, settle_timeout)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xaef91441.
//
// Solidity: function openChannel(participant1 address, participant2 address, settle_timeout uint64) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) OpenChannel(participant1 common.Address, participant2 common.Address, settle_timeout uint64) (*types.Transaction, error) {
	return _TokenNetwork.Contract.OpenChannel(&_TokenNetwork.TransactOpts, participant1, participant2, settle_timeout)
}

// OpenChannelWithDeposit is a paid mutator transaction binding the contract method 0xfc656970.
//
// Solidity: function openChannelWithDeposit(participant address, partner address, settle_timeout uint64, deposit uint256) returns()
func (_TokenNetwork *TokenNetworkTransactor) OpenChannelWithDeposit(opts *bind.TransactOpts, participant common.Address, partner common.Address, settle_timeout uint64, deposit *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "openChannelWithDeposit", participant, partner, settle_timeout, deposit)
}

// OpenChannelWithDeposit is a paid mutator transaction binding the contract method 0xfc656970.
//
// Solidity: function openChannelWithDeposit(participant address, partner address, settle_timeout uint64, deposit uint256) returns()
func (_TokenNetwork *TokenNetworkSession) OpenChannelWithDeposit(participant common.Address, partner common.Address, settle_timeout uint64, deposit *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.Contract.OpenChannelWithDeposit(&_TokenNetwork.TransactOpts, participant, partner, settle_timeout, deposit)
}

// OpenChannelWithDeposit is a paid mutator transaction binding the contract method 0xfc656970.
//
// Solidity: function openChannelWithDeposit(participant address, partner address, settle_timeout uint64, deposit uint256) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) OpenChannelWithDeposit(participant common.Address, partner common.Address, settle_timeout uint64, deposit *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.Contract.OpenChannelWithDeposit(&_TokenNetwork.TransactOpts, participant, partner, settle_timeout, deposit)
}

// PunishObsoleteUnlock is a paid mutator transaction binding the contract method 0x837536b9.
//
// Solidity: function punishObsoleteUnlock(beneficiary address, cheater address, lockhash bytes32, additional_hash bytes32, cheater_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) PunishObsoleteUnlock(opts *bind.TransactOpts, beneficiary common.Address, cheater common.Address, lockhash [32]byte, additional_hash [32]byte, cheater_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "punishObsoleteUnlock", beneficiary, cheater, lockhash, additional_hash, cheater_signature)
}

// PunishObsoleteUnlock is a paid mutator transaction binding the contract method 0x837536b9.
//
// Solidity: function punishObsoleteUnlock(beneficiary address, cheater address, lockhash bytes32, additional_hash bytes32, cheater_signature bytes) returns()
func (_TokenNetwork *TokenNetworkSession) PunishObsoleteUnlock(beneficiary common.Address, cheater common.Address, lockhash [32]byte, additional_hash [32]byte, cheater_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.PunishObsoleteUnlock(&_TokenNetwork.TransactOpts, beneficiary, cheater, lockhash, additional_hash, cheater_signature)
}

// PunishObsoleteUnlock is a paid mutator transaction binding the contract method 0x837536b9.
//
// Solidity: function punishObsoleteUnlock(beneficiary address, cheater address, lockhash bytes32, additional_hash bytes32, cheater_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) PunishObsoleteUnlock(beneficiary common.Address, cheater common.Address, lockhash [32]byte, additional_hash [32]byte, cheater_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.PunishObsoleteUnlock(&_TokenNetwork.TransactOpts, beneficiary, cheater, lockhash, additional_hash, cheater_signature)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(from address, value uint256, token_ address, data bytes) returns(success bool)
func (_TokenNetwork *TokenNetworkTransactor) ReceiveApproval(opts *bind.TransactOpts, from common.Address, value *big.Int, token_ common.Address, data []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "receiveApproval", from, value, token_, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(from address, value uint256, token_ address, data bytes) returns(success bool)
func (_TokenNetwork *TokenNetworkSession) ReceiveApproval(from common.Address, value *big.Int, token_ common.Address, data []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.ReceiveApproval(&_TokenNetwork.TransactOpts, from, value, token_, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(from address, value uint256, token_ address, data bytes) returns(success bool)
func (_TokenNetwork *TokenNetworkTransactorSession) ReceiveApproval(from common.Address, value *big.Int, token_ common.Address, data []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.ReceiveApproval(&_TokenNetwork.TransactOpts, from, value, token_, data)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xe11cbf99.
//
// Solidity: function settleChannel(participant1 address, participant1_transferred_amount uint256, participant1_locksroot bytes32, participant2 address, participant2_transferred_amount uint256, participant2_locksroot bytes32) returns()
func (_TokenNetwork *TokenNetworkTransactor) SettleChannel(opts *bind.TransactOpts, participant1 common.Address, participant1_transferred_amount *big.Int, participant1_locksroot [32]byte, participant2 common.Address, participant2_transferred_amount *big.Int, participant2_locksroot [32]byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "settleChannel", participant1, participant1_transferred_amount, participant1_locksroot, participant2, participant2_transferred_amount, participant2_locksroot)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xe11cbf99.
//
// Solidity: function settleChannel(participant1 address, participant1_transferred_amount uint256, participant1_locksroot bytes32, participant2 address, participant2_transferred_amount uint256, participant2_locksroot bytes32) returns()
func (_TokenNetwork *TokenNetworkSession) SettleChannel(participant1 common.Address, participant1_transferred_amount *big.Int, participant1_locksroot [32]byte, participant2 common.Address, participant2_transferred_amount *big.Int, participant2_locksroot [32]byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.SettleChannel(&_TokenNetwork.TransactOpts, participant1, participant1_transferred_amount, participant1_locksroot, participant2, participant2_transferred_amount, participant2_locksroot)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xe11cbf99.
//
// Solidity: function settleChannel(participant1 address, participant1_transferred_amount uint256, participant1_locksroot bytes32, participant2 address, participant2_transferred_amount uint256, participant2_locksroot bytes32) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) SettleChannel(participant1 common.Address, participant1_transferred_amount *big.Int, participant1_locksroot [32]byte, participant2 common.Address, participant2_transferred_amount *big.Int, participant2_locksroot [32]byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.SettleChannel(&_TokenNetwork.TransactOpts, participant1, participant1_transferred_amount, participant1_locksroot, participant2, participant2_transferred_amount, participant2_locksroot)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback( address, value uint256, data bytes) returns(success bool)
func (_TokenNetwork *TokenNetworkTransactor) TokenFallback(opts *bind.TransactOpts, arg0 common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "tokenFallback", arg0, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback( address, value uint256, data bytes) returns(success bool)
func (_TokenNetwork *TokenNetworkSession) TokenFallback(arg0 common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.TokenFallback(&_TokenNetwork.TransactOpts, arg0, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback( address, value uint256, data bytes) returns(success bool)
func (_TokenNetwork *TokenNetworkTransactorSession) TokenFallback(arg0 common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.TokenFallback(&_TokenNetwork.TransactOpts, arg0, value, data)
}

// Unlock is a paid mutator transaction binding the contract method 0x4aaf2b54.
//
// Solidity: function unlock(partner address, transferred_amount uint256, expiration uint256, amount uint256, secret_hash bytes32, merkle_proof bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) Unlock(opts *bind.TransactOpts, partner common.Address, transferred_amount *big.Int, expiration *big.Int, amount *big.Int, secret_hash [32]byte, merkle_proof []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "unlock", partner, transferred_amount, expiration, amount, secret_hash, merkle_proof)
}

// Unlock is a paid mutator transaction binding the contract method 0x4aaf2b54.
//
// Solidity: function unlock(partner address, transferred_amount uint256, expiration uint256, amount uint256, secret_hash bytes32, merkle_proof bytes) returns()
func (_TokenNetwork *TokenNetworkSession) Unlock(partner common.Address, transferred_amount *big.Int, expiration *big.Int, amount *big.Int, secret_hash [32]byte, merkle_proof []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.Unlock(&_TokenNetwork.TransactOpts, partner, transferred_amount, expiration, amount, secret_hash, merkle_proof)
}

// Unlock is a paid mutator transaction binding the contract method 0x4aaf2b54.
//
// Solidity: function unlock(partner address, transferred_amount uint256, expiration uint256, amount uint256, secret_hash bytes32, merkle_proof bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) Unlock(partner common.Address, transferred_amount *big.Int, expiration *big.Int, amount *big.Int, secret_hash [32]byte, merkle_proof []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.Unlock(&_TokenNetwork.TransactOpts, partner, transferred_amount, expiration, amount, secret_hash, merkle_proof)
}

// UnlockDelegate is a paid mutator transaction binding the contract method 0xa570b7d5.
//
// Solidity: function unlockDelegate(partner address, participant address, transferred_amount uint256, expiration uint256, amount uint256, secret_hash bytes32, merkle_proof bytes, participant_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) UnlockDelegate(opts *bind.TransactOpts, partner common.Address, participant common.Address, transferred_amount *big.Int, expiration *big.Int, amount *big.Int, secret_hash [32]byte, merkle_proof []byte, participant_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "unlockDelegate", partner, participant, transferred_amount, expiration, amount, secret_hash, merkle_proof, participant_signature)
}

// UnlockDelegate is a paid mutator transaction binding the contract method 0xa570b7d5.
//
// Solidity: function unlockDelegate(partner address, participant address, transferred_amount uint256, expiration uint256, amount uint256, secret_hash bytes32, merkle_proof bytes, participant_signature bytes) returns()
func (_TokenNetwork *TokenNetworkSession) UnlockDelegate(partner common.Address, participant common.Address, transferred_amount *big.Int, expiration *big.Int, amount *big.Int, secret_hash [32]byte, merkle_proof []byte, participant_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.UnlockDelegate(&_TokenNetwork.TransactOpts, partner, participant, transferred_amount, expiration, amount, secret_hash, merkle_proof, participant_signature)
}

// UnlockDelegate is a paid mutator transaction binding the contract method 0xa570b7d5.
//
// Solidity: function unlockDelegate(partner address, participant address, transferred_amount uint256, expiration uint256, amount uint256, secret_hash bytes32, merkle_proof bytes, participant_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) UnlockDelegate(partner common.Address, participant common.Address, transferred_amount *big.Int, expiration *big.Int, amount *big.Int, secret_hash [32]byte, merkle_proof []byte, participant_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.UnlockDelegate(&_TokenNetwork.TransactOpts, partner, participant, transferred_amount, expiration, amount, secret_hash, merkle_proof, participant_signature)
}

// UpdateBalanceProof is a paid mutator transaction binding the contract method 0xaaa3dbcc.
//
// Solidity: function updateBalanceProof(partner address, transferred_amount uint256, locksroot bytes32, nonce uint64, additional_hash bytes32, partner_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) UpdateBalanceProof(opts *bind.TransactOpts, partner common.Address, transferred_amount *big.Int, locksroot [32]byte, nonce uint64, additional_hash [32]byte, partner_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "updateBalanceProof", partner, transferred_amount, locksroot, nonce, additional_hash, partner_signature)
}

// UpdateBalanceProof is a paid mutator transaction binding the contract method 0xaaa3dbcc.
//
// Solidity: function updateBalanceProof(partner address, transferred_amount uint256, locksroot bytes32, nonce uint64, additional_hash bytes32, partner_signature bytes) returns()
func (_TokenNetwork *TokenNetworkSession) UpdateBalanceProof(partner common.Address, transferred_amount *big.Int, locksroot [32]byte, nonce uint64, additional_hash [32]byte, partner_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.UpdateBalanceProof(&_TokenNetwork.TransactOpts, partner, transferred_amount, locksroot, nonce, additional_hash, partner_signature)
}

// UpdateBalanceProof is a paid mutator transaction binding the contract method 0xaaa3dbcc.
//
// Solidity: function updateBalanceProof(partner address, transferred_amount uint256, locksroot bytes32, nonce uint64, additional_hash bytes32, partner_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) UpdateBalanceProof(partner common.Address, transferred_amount *big.Int, locksroot [32]byte, nonce uint64, additional_hash [32]byte, partner_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.UpdateBalanceProof(&_TokenNetwork.TransactOpts, partner, transferred_amount, locksroot, nonce, additional_hash, partner_signature)
}

// UpdateBalanceProofDelegate is a paid mutator transaction binding the contract method 0xf8658b25.
//
// Solidity: function updateBalanceProofDelegate(partner address, participant address, transferred_amount uint256, locksroot bytes32, nonce uint64, additional_hash bytes32, partner_signature bytes, participant_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) UpdateBalanceProofDelegate(opts *bind.TransactOpts, partner common.Address, participant common.Address, transferred_amount *big.Int, locksroot [32]byte, nonce uint64, additional_hash [32]byte, partner_signature []byte, participant_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "updateBalanceProofDelegate", partner, participant, transferred_amount, locksroot, nonce, additional_hash, partner_signature, participant_signature)
}

// UpdateBalanceProofDelegate is a paid mutator transaction binding the contract method 0xf8658b25.
//
// Solidity: function updateBalanceProofDelegate(partner address, participant address, transferred_amount uint256, locksroot bytes32, nonce uint64, additional_hash bytes32, partner_signature bytes, participant_signature bytes) returns()
func (_TokenNetwork *TokenNetworkSession) UpdateBalanceProofDelegate(partner common.Address, participant common.Address, transferred_amount *big.Int, locksroot [32]byte, nonce uint64, additional_hash [32]byte, partner_signature []byte, participant_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.UpdateBalanceProofDelegate(&_TokenNetwork.TransactOpts, partner, participant, transferred_amount, locksroot, nonce, additional_hash, partner_signature, participant_signature)
}

// UpdateBalanceProofDelegate is a paid mutator transaction binding the contract method 0xf8658b25.
//
// Solidity: function updateBalanceProofDelegate(partner address, participant address, transferred_amount uint256, locksroot bytes32, nonce uint64, additional_hash bytes32, partner_signature bytes, participant_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) UpdateBalanceProofDelegate(partner common.Address, participant common.Address, transferred_amount *big.Int, locksroot [32]byte, nonce uint64, additional_hash [32]byte, partner_signature []byte, participant_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.UpdateBalanceProofDelegate(&_TokenNetwork.TransactOpts, partner, participant, transferred_amount, locksroot, nonce, additional_hash, partner_signature, participant_signature)
}

// WithDraw is a paid mutator transaction binding the contract method 0x7c090e4b.
//
// Solidity: function withDraw(participant address, partner address, participant_balance uint256, participant_withdraw uint256, participant_signature bytes, partner_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) WithDraw(opts *bind.TransactOpts, participant common.Address, partner common.Address, participant_balance *big.Int, participant_withdraw *big.Int, participant_signature []byte, partner_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "withDraw", participant, partner, participant_balance, participant_withdraw, participant_signature, partner_signature)
}

// WithDraw is a paid mutator transaction binding the contract method 0x7c090e4b.
//
// Solidity: function withDraw(participant address, partner address, participant_balance uint256, participant_withdraw uint256, participant_signature bytes, partner_signature bytes) returns()
func (_TokenNetwork *TokenNetworkSession) WithDraw(participant common.Address, partner common.Address, participant_balance *big.Int, participant_withdraw *big.Int, participant_signature []byte, partner_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.WithDraw(&_TokenNetwork.TransactOpts, participant, partner, participant_balance, participant_withdraw, participant_signature, partner_signature)
}

// WithDraw is a paid mutator transaction binding the contract method 0x7c090e4b.
//
// Solidity: function withDraw(participant address, partner address, participant_balance uint256, participant_withdraw uint256, participant_signature bytes, partner_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) WithDraw(participant common.Address, partner common.Address, participant_balance *big.Int, participant_withdraw *big.Int, participant_signature []byte, partner_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.WithDraw(&_TokenNetwork.TransactOpts, participant, partner, participant_balance, participant_withdraw, participant_signature, partner_signature)
}

// TokenNetworkBalanceProofUpdatedIterator is returned from FilterBalanceProofUpdated and is used to iterate over the raw logs and unpacked data for BalanceProofUpdated events raised by the TokenNetwork contract.
type TokenNetworkBalanceProofUpdatedIterator struct {
	Event *TokenNetworkBalanceProofUpdated // Event containing the contract specifics and raw log

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
func (it *TokenNetworkBalanceProofUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkBalanceProofUpdated)
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
		it.Event = new(TokenNetworkBalanceProofUpdated)
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
func (it *TokenNetworkBalanceProofUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkBalanceProofUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkBalanceProofUpdated represents a BalanceProofUpdated event raised by the TokenNetwork contract.
type TokenNetworkBalanceProofUpdated struct {
	ChannelIdentifier [32]byte
	Participant       common.Address
	Locksroot         [32]byte
	TransferredAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBalanceProofUpdated is a free log retrieval operation binding the contract event 0x910c9237f4197a18340110a181e8fb775496506a007a94b46f9f80f2a35918f9.
//
// Solidity: e BalanceProofUpdated(channel_identifier indexed bytes32, participant address, locksroot bytes32, transferred_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterBalanceProofUpdated(opts *bind.FilterOpts, channel_identifier [][32]byte) (*TokenNetworkBalanceProofUpdatedIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "BalanceProofUpdated", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkBalanceProofUpdatedIterator{contract: _TokenNetwork.contract, event: "BalanceProofUpdated", logs: logs, sub: sub}, nil
}

// WatchBalanceProofUpdated is a free log subscription operation binding the contract event 0x910c9237f4197a18340110a181e8fb775496506a007a94b46f9f80f2a35918f9.
//
// Solidity: e BalanceProofUpdated(channel_identifier indexed bytes32, participant address, locksroot bytes32, transferred_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchBalanceProofUpdated(opts *bind.WatchOpts, sink chan<- *TokenNetworkBalanceProofUpdated, channel_identifier [][32]byte) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "BalanceProofUpdated", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkBalanceProofUpdated)
				if err := _TokenNetwork.contract.UnpackLog(event, "BalanceProofUpdated", log); err != nil {
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

// TokenNetworkChannelClosedIterator is returned from FilterChannelClosed and is used to iterate over the raw logs and unpacked data for ChannelClosed events raised by the TokenNetwork contract.
type TokenNetworkChannelClosedIterator struct {
	Event *TokenNetworkChannelClosed // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelClosed)
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
		it.Event = new(TokenNetworkChannelClosed)
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
func (it *TokenNetworkChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelClosed represents a ChannelClosed event raised by the TokenNetwork contract.
type TokenNetworkChannelClosed struct {
	ChannelIdentifier  [32]byte
	ClosingParticipant common.Address
	Locksroot          [32]byte
	TransferredAmount  *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChannelClosed is a free log retrieval operation binding the contract event 0x69610baaace24c039f891a11b42c0b1df1496ab0db38b0c4ee4ed33d6d53da1a.
//
// Solidity: e ChannelClosed(channel_identifier indexed bytes32, closing_participant address, locksroot bytes32, transferred_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelClosed(opts *bind.FilterOpts, channel_identifier [][32]byte) (*TokenNetworkChannelClosedIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelClosed", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelClosedIterator{contract: _TokenNetwork.contract, event: "ChannelClosed", logs: logs, sub: sub}, nil
}

// WatchChannelClosed is a free log subscription operation binding the contract event 0x69610baaace24c039f891a11b42c0b1df1496ab0db38b0c4ee4ed33d6d53da1a.
//
// Solidity: e ChannelClosed(channel_identifier indexed bytes32, closing_participant address, locksroot bytes32, transferred_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelClosed(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelClosed, channel_identifier [][32]byte) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelClosed", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelClosed)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
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

// TokenNetworkChannelCooperativeSettledIterator is returned from FilterChannelCooperativeSettled and is used to iterate over the raw logs and unpacked data for ChannelCooperativeSettled events raised by the TokenNetwork contract.
type TokenNetworkChannelCooperativeSettledIterator struct {
	Event *TokenNetworkChannelCooperativeSettled // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelCooperativeSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelCooperativeSettled)
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
		it.Event = new(TokenNetworkChannelCooperativeSettled)
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
func (it *TokenNetworkChannelCooperativeSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelCooperativeSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelCooperativeSettled represents a ChannelCooperativeSettled event raised by the TokenNetwork contract.
type TokenNetworkChannelCooperativeSettled struct {
	ChannelIdentifier  [32]byte
	Participant1Amount *big.Int
	Participant2Amount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChannelCooperativeSettled is a free log retrieval operation binding the contract event 0xfb2f4bc0fb2e0f1001f78d15e81a2e1981f262d31e8bd72309e26cc63bf7bb02.
//
// Solidity: e ChannelCooperativeSettled(channel_identifier indexed bytes32, participant1_amount uint256, participant2_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelCooperativeSettled(opts *bind.FilterOpts, channel_identifier [][32]byte) (*TokenNetworkChannelCooperativeSettledIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelCooperativeSettled", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelCooperativeSettledIterator{contract: _TokenNetwork.contract, event: "ChannelCooperativeSettled", logs: logs, sub: sub}, nil
}

// WatchChannelCooperativeSettled is a free log subscription operation binding the contract event 0xfb2f4bc0fb2e0f1001f78d15e81a2e1981f262d31e8bd72309e26cc63bf7bb02.
//
// Solidity: e ChannelCooperativeSettled(channel_identifier indexed bytes32, participant1_amount uint256, participant2_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelCooperativeSettled(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelCooperativeSettled, channel_identifier [][32]byte) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelCooperativeSettled", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelCooperativeSettled)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelCooperativeSettled", log); err != nil {
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

// TokenNetworkChannelNewDepositIterator is returned from FilterChannelNewDeposit and is used to iterate over the raw logs and unpacked data for ChannelNewDeposit events raised by the TokenNetwork contract.
type TokenNetworkChannelNewDepositIterator struct {
	Event *TokenNetworkChannelNewDeposit // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelNewDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelNewDeposit)
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
		it.Event = new(TokenNetworkChannelNewDeposit)
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
func (it *TokenNetworkChannelNewDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelNewDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelNewDeposit represents a ChannelNewDeposit event raised by the TokenNetwork contract.
type TokenNetworkChannelNewDeposit struct {
	ChannelIdentifier [32]byte
	Participant       common.Address
	TotalDeposit      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChannelNewDeposit is a free log retrieval operation binding the contract event 0x0346e981e2bfa2366dc2307a8f1fa24779830a01121b1275fe565c6b98bb4d34.
//
// Solidity: e ChannelNewDeposit(channel_identifier indexed bytes32, participant address, total_deposit uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelNewDeposit(opts *bind.FilterOpts, channel_identifier [][32]byte) (*TokenNetworkChannelNewDepositIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelNewDeposit", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelNewDepositIterator{contract: _TokenNetwork.contract, event: "ChannelNewDeposit", logs: logs, sub: sub}, nil
}

// WatchChannelNewDeposit is a free log subscription operation binding the contract event 0x0346e981e2bfa2366dc2307a8f1fa24779830a01121b1275fe565c6b98bb4d34.
//
// Solidity: e ChannelNewDeposit(channel_identifier indexed bytes32, participant address, total_deposit uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelNewDeposit(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelNewDeposit, channel_identifier [][32]byte) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelNewDeposit", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelNewDeposit)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelNewDeposit", log); err != nil {
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

// TokenNetworkChannelOpenedIterator is returned from FilterChannelOpened and is used to iterate over the raw logs and unpacked data for ChannelOpened events raised by the TokenNetwork contract.
type TokenNetworkChannelOpenedIterator struct {
	Event *TokenNetworkChannelOpened // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelOpened)
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
		it.Event = new(TokenNetworkChannelOpened)
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
func (it *TokenNetworkChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelOpened represents a ChannelOpened event raised by the TokenNetwork contract.
type TokenNetworkChannelOpened struct {
	ChannelIdentifier [32]byte
	Participant1      common.Address
	Participant2      common.Address
	SettleTimeout     uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChannelOpened is a free log retrieval operation binding the contract event 0x4d4097deeecde59dede1bb370eb147fc3fa969b7b6a6f89f95526635328e86df.
//
// Solidity: e ChannelOpened(channel_identifier indexed bytes32, participant1 address, participant2 address, settle_timeout uint64)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelOpened(opts *bind.FilterOpts, channel_identifier [][32]byte) (*TokenNetworkChannelOpenedIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelOpened", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelOpenedIterator{contract: _TokenNetwork.contract, event: "ChannelOpened", logs: logs, sub: sub}, nil
}

// WatchChannelOpened is a free log subscription operation binding the contract event 0x4d4097deeecde59dede1bb370eb147fc3fa969b7b6a6f89f95526635328e86df.
//
// Solidity: e ChannelOpened(channel_identifier indexed bytes32, participant1 address, participant2 address, settle_timeout uint64)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelOpened(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelOpened, channel_identifier [][32]byte) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelOpened", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelOpened)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
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

// TokenNetworkChannelOpenedAndDepositIterator is returned from FilterChannelOpenedAndDeposit and is used to iterate over the raw logs and unpacked data for ChannelOpenedAndDeposit events raised by the TokenNetwork contract.
type TokenNetworkChannelOpenedAndDepositIterator struct {
	Event *TokenNetworkChannelOpenedAndDeposit // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelOpenedAndDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelOpenedAndDeposit)
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
		it.Event = new(TokenNetworkChannelOpenedAndDeposit)
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
func (it *TokenNetworkChannelOpenedAndDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelOpenedAndDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelOpenedAndDeposit represents a ChannelOpenedAndDeposit event raised by the TokenNetwork contract.
type TokenNetworkChannelOpenedAndDeposit struct {
	ChannelIdentifier   [32]byte
	Participant1        common.Address
	Participant2        common.Address
	SettleTimeout       uint64
	Participant1Deposit *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenedAndDeposit is a free log retrieval operation binding the contract event 0xcac76648b0a531becb6e54db5fe838853fdc47ef130aab3566114ee7c739d0a0.
//
// Solidity: e ChannelOpenedAndDeposit(channel_identifier indexed bytes32, participant1 address, participant2 address, settle_timeout uint64, participant1_deposit uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelOpenedAndDeposit(opts *bind.FilterOpts, channel_identifier [][32]byte) (*TokenNetworkChannelOpenedAndDepositIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelOpenedAndDeposit", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelOpenedAndDepositIterator{contract: _TokenNetwork.contract, event: "ChannelOpenedAndDeposit", logs: logs, sub: sub}, nil
}

// WatchChannelOpenedAndDeposit is a free log subscription operation binding the contract event 0xcac76648b0a531becb6e54db5fe838853fdc47ef130aab3566114ee7c739d0a0.
//
// Solidity: e ChannelOpenedAndDeposit(channel_identifier indexed bytes32, participant1 address, participant2 address, settle_timeout uint64, participant1_deposit uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelOpenedAndDeposit(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelOpenedAndDeposit, channel_identifier [][32]byte) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelOpenedAndDeposit", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelOpenedAndDeposit)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelOpenedAndDeposit", log); err != nil {
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

// TokenNetworkChannelPunishedIterator is returned from FilterChannelPunished and is used to iterate over the raw logs and unpacked data for ChannelPunished events raised by the TokenNetwork contract.
type TokenNetworkChannelPunishedIterator struct {
	Event *TokenNetworkChannelPunished // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelPunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelPunished)
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
		it.Event = new(TokenNetworkChannelPunished)
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
func (it *TokenNetworkChannelPunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelPunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelPunished represents a ChannelPunished event raised by the TokenNetwork contract.
type TokenNetworkChannelPunished struct {
	ChannelIdentifier [32]byte
	Beneficiary       common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChannelPunished is a free log retrieval operation binding the contract event 0xa913b8478dcdecf113bad71030afc079c268eb9abc88e45615f438824127ae00.
//
// Solidity: e ChannelPunished(channel_identifier indexed bytes32, beneficiary address)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelPunished(opts *bind.FilterOpts, channel_identifier [][32]byte) (*TokenNetworkChannelPunishedIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelPunished", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelPunishedIterator{contract: _TokenNetwork.contract, event: "ChannelPunished", logs: logs, sub: sub}, nil
}

// WatchChannelPunished is a free log subscription operation binding the contract event 0xa913b8478dcdecf113bad71030afc079c268eb9abc88e45615f438824127ae00.
//
// Solidity: e ChannelPunished(channel_identifier indexed bytes32, beneficiary address)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelPunished(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelPunished, channel_identifier [][32]byte) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelPunished", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelPunished)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelPunished", log); err != nil {
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

// TokenNetworkChannelSettledIterator is returned from FilterChannelSettled and is used to iterate over the raw logs and unpacked data for ChannelSettled events raised by the TokenNetwork contract.
type TokenNetworkChannelSettledIterator struct {
	Event *TokenNetworkChannelSettled // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelSettled)
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
		it.Event = new(TokenNetworkChannelSettled)
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
func (it *TokenNetworkChannelSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelSettled represents a ChannelSettled event raised by the TokenNetwork contract.
type TokenNetworkChannelSettled struct {
	ChannelIdentifier  [32]byte
	Participant1Amount *big.Int
	Participant2Amount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChannelSettled is a free log retrieval operation binding the contract event 0xf94fb5c0628a82dc90648e8dc5e983f632633b0d26603d64e8cc042ca0790aa4.
//
// Solidity: e ChannelSettled(channel_identifier indexed bytes32, participant1_amount uint256, participant2_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelSettled(opts *bind.FilterOpts, channel_identifier [][32]byte) (*TokenNetworkChannelSettledIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelSettled", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelSettledIterator{contract: _TokenNetwork.contract, event: "ChannelSettled", logs: logs, sub: sub}, nil
}

// WatchChannelSettled is a free log subscription operation binding the contract event 0xf94fb5c0628a82dc90648e8dc5e983f632633b0d26603d64e8cc042ca0790aa4.
//
// Solidity: e ChannelSettled(channel_identifier indexed bytes32, participant1_amount uint256, participant2_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelSettled(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelSettled, channel_identifier [][32]byte) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelSettled", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelSettled)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelSettled", log); err != nil {
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

// TokenNetworkChannelUnlockedIterator is returned from FilterChannelUnlocked and is used to iterate over the raw logs and unpacked data for ChannelUnlocked events raised by the TokenNetwork contract.
type TokenNetworkChannelUnlockedIterator struct {
	Event *TokenNetworkChannelUnlocked // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelUnlocked)
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
		it.Event = new(TokenNetworkChannelUnlocked)
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
func (it *TokenNetworkChannelUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelUnlocked represents a ChannelUnlocked event raised by the TokenNetwork contract.
type TokenNetworkChannelUnlocked struct {
	ChannelIdentifier [32]byte
	PayerParticipant  common.Address
	Lockhash          [32]byte
	TransferredAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChannelUnlocked is a free log retrieval operation binding the contract event 0x9e3b094fde58f3a83bd8b77d0a995fdb71f3169c6fa7e6d386e9f5902841e5ff.
//
// Solidity: e ChannelUnlocked(channel_identifier indexed bytes32, payer_participant address, lockhash bytes32, transferred_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelUnlocked(opts *bind.FilterOpts, channel_identifier [][32]byte) (*TokenNetworkChannelUnlockedIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelUnlocked", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelUnlockedIterator{contract: _TokenNetwork.contract, event: "ChannelUnlocked", logs: logs, sub: sub}, nil
}

// WatchChannelUnlocked is a free log subscription operation binding the contract event 0x9e3b094fde58f3a83bd8b77d0a995fdb71f3169c6fa7e6d386e9f5902841e5ff.
//
// Solidity: e ChannelUnlocked(channel_identifier indexed bytes32, payer_participant address, lockhash bytes32, transferred_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelUnlocked(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelUnlocked, channel_identifier [][32]byte) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelUnlocked", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelUnlocked)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelUnlocked", log); err != nil {
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

// TokenNetworkChannelWithdrawIterator is returned from FilterChannelWithdraw and is used to iterate over the raw logs and unpacked data for ChannelWithdraw events raised by the TokenNetwork contract.
type TokenNetworkChannelWithdrawIterator struct {
	Event *TokenNetworkChannelWithdraw // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelWithdraw)
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
		it.Event = new(TokenNetworkChannelWithdraw)
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
func (it *TokenNetworkChannelWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelWithdraw represents a ChannelWithdraw event raised by the TokenNetwork contract.
type TokenNetworkChannelWithdraw struct {
	ChannelIdentifier   [32]byte
	Participant1        common.Address
	Participant1Balance *big.Int
	Participant2        common.Address
	Participant2Balance *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterChannelWithdraw is a free log retrieval operation binding the contract event 0xdc5ff4ab383e66679a382f376c0e80534f51f3f3a398add646422cd81f5f815d.
//
// Solidity: e ChannelWithdraw(channel_identifier indexed bytes32, participant1 address, participant1_balance uint256, participant2 address, participant2_balance uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelWithdraw(opts *bind.FilterOpts, channel_identifier [][32]byte) (*TokenNetworkChannelWithdrawIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelWithdraw", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelWithdrawIterator{contract: _TokenNetwork.contract, event: "ChannelWithdraw", logs: logs, sub: sub}, nil
}

// WatchChannelWithdraw is a free log subscription operation binding the contract event 0xdc5ff4ab383e66679a382f376c0e80534f51f3f3a398add646422cd81f5f815d.
//
// Solidity: e ChannelWithdraw(channel_identifier indexed bytes32, participant1 address, participant1_balance uint256, participant2 address, participant2_balance uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelWithdraw(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelWithdraw, channel_identifier [][32]byte) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelWithdraw", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelWithdraw)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelWithdraw", log); err != nil {
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

// TokenNetworkRegistryABI is the input ABI used to generate the binding from.
const TokenNetworkRegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"token_to_token_networks\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chain_id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token_address\",\"type\":\"address\"}],\"name\":\"createERC20TokenNetwork\",\"outputs\":[{\"name\":\"token_network_address\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contract_address\",\"type\":\"address\"}],\"name\":\"contractExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contract_version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secret_registry_address\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_secret_registry_address\",\"type\":\"address\"},{\"name\":\"_chain_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token_address\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token_network_address\",\"type\":\"address\"}],\"name\":\"TokenNetworkCreated\",\"type\":\"event\"}]"

// TokenNetworkRegistryBin is the compiled bytecode used for deploying new contracts.
const TokenNetworkRegistryBin = `0x608060405234801561001057600080fd5b506040516040806143e98339810160405280516020909101516000811161003657600080fd5b600160a060020a038216151561004b57600080fd5b61005d82640100000000610091810204565b151561006857600080fd5b60008054600160a060020a031916600160a060020a039390931692909217909155600155610099565b6000903b1190565b614341806100a86000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630fabd9e7811461007c5780633af973b1146100d35780634cf71a04146100fa5780637709bc7814610128578063b32c65c81461016a578063d0ad4bec146101f4575b600080fd5b34801561008857600080fd5b506100aa73ffffffffffffffffffffffffffffffffffffffff60043516610209565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100df57600080fd5b506100e8610231565b60408051918252519081900360200190f35b34801561010657600080fd5b506100aa73ffffffffffffffffffffffffffffffffffffffff60043516610237565b34801561013457600080fd5b5061015673ffffffffffffffffffffffffffffffffffffffff60043516610361565b604080519115158252519081900360200190f35b34801561017657600080fd5b5061017f610369565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101b95781810151838201526020016101a1565b50505050905090810190601f1680156101e65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020057600080fd5b506100aa6103a0565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b73ffffffffffffffffffffffffffffffffffffffff8082166000908152600260205260408120549091161561026b57600080fd5b600054600154839173ffffffffffffffffffffffffffffffffffffffff16906102926103bc565b73ffffffffffffffffffffffffffffffffffffffff9384168152919092166020820152604080820192909252905190819003606001906000f0801580156102dd573d6000803e3d6000fd5b5073ffffffffffffffffffffffffffffffffffffffff83811660008181526002602052604080822080547fffffffffffffffffffffffff000000000000000000000000000000000000000016948616948517905551939450919290917ff11a7558a113d9627989c5edf26cbd19143b7375248e621c8e30ac9e0847dc3f91a3919050565b6000903b1190565b60408051808201909152600581527f302e332e5f000000000000000000000000000000000000000000000000000000602082015281565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b604051613f49806103cd83390190560060806040523480156200001157600080fd5b5060405160608062003f49833981016040908152815160208301519190920151600160a060020a03831615156200004757600080fd5b600160a060020a03821615156200005d57600080fd5b600081116200006b57600080fd5b6200007f8364010000000062000177810204565b15156200008b57600080fd5b6200009f8264010000000062000177810204565b1515620000ab57600080fd5b60008054600160a060020a03808616600160a060020a031992831617808455600180548784169416939093179092556002849055604080517f18160ddd000000000000000000000000000000000000000000000000000000008152905192909116916318160ddd9160048082019260209290919082900301818787803b1580156200013557600080fd5b505af11580156200014a573d6000803e3d6000fd5b505050506040513d60208110156200016157600080fd5b5051116200016e57600080fd5b5050506200017f565b6000903b1190565b613dba806200018f6000396000f3006080604052600436106101535763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166324d73a9381146101585780633af973b1146101965780634aaf2b54146101bd5780637709bc78146102415780637a7ebd7b146102835780637c090e4b146102d35780638340f54914610395578063837536b9146103cc5780638568536a1461045057806387234237146105145780638b1ddc531461059e5780638f4ffcb1146105d55780639375cff21461061a5780639fe5b1871461064c578063a570b7d5146106a0578063aaa3dbcc1461076d578063ac133709146107f9578063aef914411461085f578063b32c65c8146108a0578063b9eec014146108b5578063c0ee0b8a14610941578063e11cbf991461097f578063f8658b25146109c1578063f94c9e1314610a98578063fc0c546a14610acc578063fc65697014610ae1575b600080fd5b34801561016457600080fd5b5061016d610b25565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156101a257600080fd5b506101ab610b41565b60408051918252519081900360200190f35b3480156101c957600080fd5b50604080516020600460a43581810135601f810184900484028501840190955284845261023f94823573ffffffffffffffffffffffffffffffffffffffff169460248035956044359560643595608435953695929460c4949201918190840183828082843750949750610b479650505050505050565b005b34801561024d57600080fd5b5061026f73ffffffffffffffffffffffffffffffffffffffff60043516610b5e565b604080519115158252519081900360200190f35b34801561028f57600080fd5b5061029b600435610b66565b6040805167ffffffffffffffff95861681529385166020850152919093168282015260ff909216606082015290519081900360800190f35b3480156102df57600080fd5b50604080516020601f60843560048181013592830184900484028501840190955281845261023f9473ffffffffffffffffffffffffffffffffffffffff81358116956024803590921695604435956064359536959460a4949391019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750610bc99650505050505050565b3480156103a157600080fd5b5061023f73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610f0b565b3480156103d857600080fd5b50604080516020601f60843560048181013592830184900484028501840190955281845261023f9473ffffffffffffffffffffffffffffffffffffffff81358116956024803590921695604435956064359536959460a49493910191908190840183828082843750949750610f1e9650505050505050565b34801561045c57600080fd5b50604080516020601f60843560048181013592830184900484028501840190955281845261023f9473ffffffffffffffffffffffffffffffffffffffff81358116956024803596604435909316956064359536959460a49493919091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506111d19650505050505050565b34801561052057600080fd5b5061052961155f565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561056357818101518382015260200161054b565b50505050905090810190601f1680156105905780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105aa57600080fd5b5061026f73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435611596565b3480156105e157600080fd5b5061026f6004803573ffffffffffffffffffffffffffffffffffffffff90811691602480359260443516916064359182019101356116ac565b34801561062657600080fd5b5061062f61171e565b6040805167ffffffffffffffff9092168252519081900360200190f35b34801561065857600080fd5b50610664600435611723565b6040805195865267ffffffffffffffff94851660208701529284168584015260ff90911660608501529091166080830152519081900360a00190f35b3480156106ac57600080fd5b50604080516020601f60c43560048181013592830184900484028501840190955281845261023f9473ffffffffffffffffffffffffffffffffffffffff8135811695602480359092169560443595606435956084359560a435953695919460e49492939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506117899650505050505050565b34801561077957600080fd5b50604080516020600460a43581810135601f810184900484028501840190955284845261023f94823573ffffffffffffffffffffffffffffffffffffffff169460248035956044359560643567ffffffffffffffff1695608435953695929460c49492019181908401838280828437509497506119969650505050505050565b34801561080557600080fd5b5061082d73ffffffffffffffffffffffffffffffffffffffff60043581169060243516611ba4565b6040805193845267ffffffffffffffff19909216602084015267ffffffffffffffff1682820152519081900360600190f35b34801561086b57600080fd5b5061023f73ffffffffffffffffffffffffffffffffffffffff6004358116906024351667ffffffffffffffff60443516611c32565b3480156108ac57600080fd5b50610529611e16565b3480156108c157600080fd5b50604080516020600460a43581810135601f810184900484028501840190955284845261023f94823573ffffffffffffffffffffffffffffffffffffffff169460248035956044359560643567ffffffffffffffff1695608435953695929460c4949201918190840183828082843750949750611e4d9650505050505050565b34801561094d57600080fd5b5061026f6004803573ffffffffffffffffffffffffffffffffffffffff16906024803591604435918201910135612085565b34801561098b57600080fd5b5061023f73ffffffffffffffffffffffffffffffffffffffff60043581169060243590604435906064351660843560a4356120f4565b3480156109cd57600080fd5b50604080516020601f60c43560048181013592830184900484028501840190955281845261023f9473ffffffffffffffffffffffffffffffffffffffff81358116956024803590921695604435956064359567ffffffffffffffff608435169560a435953695919460e49492939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506124b69650505050505050565b348015610aa457600080fd5b5061066473ffffffffffffffffffffffffffffffffffffffff60043581169060243516612752565b348015610ad857600080fd5b5061016d6127d7565b348015610aed57600080fd5b5061023f73ffffffffffffffffffffffffffffffffffffffff6004358116906024351667ffffffffffffffff604435166064356127f3565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b610b5686338787878787612808565b505050505050565b6000903b1190565b60036020526000908152604090205467ffffffffffffffff80821691680100000000000000008104821691700100000000000000000000000000000000820416907801000000000000000000000000000000000000000000000000900460ff1684565b6000806000806000806000610bde8d8d612c44565b60008181526003602052604090208054919750700100000000000000000000000000000000820467ffffffffffffffff16965093507801000000000000000000000000000000000000000000000000900460ff16600114610c3e57600080fd5b610c4c868e8d8d898e612dc3565b73ffffffffffffffffffffffffffffffffffffffff8e8116911614610c7057600080fd5b610c7e868e8d8d898d612dc3565b73ffffffffffffffffffffffffffffffffffffffff8d8116911614610ca257600080fd5b505073ffffffffffffffffffffffffffffffffffffffff808c166000908152600183016020526040808220928d1682528120805483540197508b88039450908a11610cec57600080fd5b8a8a1115610cf957600080fd5b8a871015610d0657600080fd5b83871015610d1357600080fd5b898b039a508a8260000181905550838160000181905550438360000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8e8c6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610e1957600080fd5b505af1158015610e2d573d6000803e3d6000fd5b505050506040513d6020811015610e4357600080fd5b50511515610e5057600080fd5b85600019167fdc5ff4ab383e66679a382f376c0e80534f51f3f3a398add646422cd81f5f815d8e8d8f88604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390a250505050505050505050505050565b610f19838383336001612f7a565b505050565b600080600080600080610f318b8b612c44565b6000818152600360205260409020805491975093507801000000000000000000000000000000000000000000000000900460ff16600214610f7157600080fd5b73ffffffffffffffffffffffffffffffffffffffff8b1660009081526001848101602052604090912090810154680100000000000000000267ffffffffffffffff191695509150841515610fc457600080fd5b8254610ff39087908b90700100000000000000000000000000000000900467ffffffffffffffff168b8b613127565b73ffffffffffffffffffffffffffffffffffffffff8b811691161461101757600080fd5b5073ffffffffffffffffffffffffffffffffffffffff891660009081526001808401602090815260409283902091840154835167ffffffffffffffff780100000000000000000000000000000000000000000000000092839004169091028183015260288082018d90528451808303909101815260489091019384905280519293909290918291908401908083835b602083106110c55780518252601f1990920191602091820191016110a6565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120600081815260028901909252929020549197505060ff1615159150611113905057600080fd5b6000848152600283016020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557fffffffffffffffff000000000000000000000000000000000000000000000000600186015583548554018555918355815173ffffffffffffffffffffffffffffffffffffffff8e168152915188927fa913b8478dcdecf113bad71030afc079c268eb9abc88e45615f438824127ae0092908290030190a25050505050505050505050565b6000806000806000806111e48c8b612c44565b6000818152600360205260409020805491965093507801000000000000000000000000000000000000000000000000900460ff1660011461122457600080fd5b8254700100000000000000000000000000000000900467ffffffffffffffff169350611255858d8d8d8d898e613287565b73ffffffffffffffffffffffffffffffffffffffff8d811691161461127957600080fd5b611288858d8d8d8d898d613287565b73ffffffffffffffffffffffffffffffffffffffff8b81169116146112ac57600080fd5b505073ffffffffffffffffffffffffffffffffffffffff808b166000908152600180840160209081526040808420948d16845280842080548654868855878601879055868355948201869055898652600390935290842080547fffffffffffffff0000000000000000000000000000000000000000000000000016905591019650908b1115611430576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8d8d6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156113f957600080fd5b505af115801561140d573d6000803e3d6000fd5b505050506040513d602081101561142357600080fd5b5051151561143057600080fd5b60008911156114ef5760008054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e81166004830152602482018e90529151919092169263a9059cbb92604480820193602093909283900390910190829087803b1580156114b857600080fd5b505af11580156114cc573d6000803e3d6000fd5b505050506040513d60208110156114e257600080fd5b505115156114ef57600080fd5b8a890186146114fd57600080fd5b8a86101561150a57600080fd5b8886101561151757600080fd5b604080518c8152602081018b9052815187927ffb2f4bc0fb2e0f1001f78d15e81a2e1981f262d31e8bd72309e26cc63bf7bb02928290030190a2505050505050505050505050565b60408051808201909152601d81527f19536d61727452616964656e205369676e6564204d6573736167653a0a000000602082015281565b60008060008060006115a88888612c44565b600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8d168452600180820184529382902093840154825167ffffffffffffffff780100000000000000000000000000000000000000000000000092839004169091028185015260288082018d9052835180830390910181526048909101928390528051959850909650929450919282918401908083835b602083106116625780518252601f199092019160209182019101611643565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060009081526002969096019052509092205460ff169998505050505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff8581169116146116d457600080fd5b611712868685858080601f01602080910402602001604051908101604052809392919081815260200183838082843750600194506134269350505050565b50600195945050505050565b600581565b600081815260036020526040902054909167ffffffffffffffff680100000000000000008304811692700100000000000000000000000000000000810482169260ff78010000000000000000000000000000000000000000000000008304169290911690565b60008060006117988b8b612c44565b60008181526003602090815260409182902082518084018452601d8082527f19536d61727452616964656e205369676e6564204d6573736167653a0a00000082850190815283546002549651979950939750919533958f958f958f958c9570010000000000000000000000000000000090920467ffffffffffffffff16949092019182918083835b6020831061183f5780518252601f199092019160209182019101611820565b51815160209384036101000a600019018019909216911617905273ffffffffffffffffffffffffffffffffffffffff9b909b166c010000000000000000000000000292019182525060148101979097525060348601949094526054850192909252607484015267ffffffffffffffff167801000000000000000000000000000000000000000000000000026094830152609c8083019190915260408051808403909201825260bc90920191829052805190935090918291908401908083835b6020831061191d5780518252601f1990920191602091820191016118fe565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902092506119568385613489565b73ffffffffffffffffffffffffffffffffffffffff8b811691161461197a57600080fd5b6119898b8b8b8b8b8b8b612808565b5050505050505050505050565b60008060006119a58933612c44565b600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8e16845260018101909252909120815492955090935091507801000000000000000000000000000000000000000000000000900460ff16600214611a0e57600080fd5b8154436801000000000000000090910467ffffffffffffffff161015611a3357600080fd5b600181015467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690871611611a6d57600080fd5b611a94838989898660000160109054906101000a900467ffffffffffffffff168a8a613576565b73ffffffffffffffffffffffffffffffffffffffff8a8116911614611ab857600080fd5b611ac288886136a5565b60018201805467ffffffffffffffff8916780100000000000000000000000000000000000000000000000002680100000000000000009093047fffffffffffffffff0000000000000000000000000000000000000000000000009091161777ffffffffffffffffffffffffffffffffffffffffffffffff169190911790556040805173ffffffffffffffffffffffffffffffffffffffff8b168152602081018990528082018a9052905184917f910c9237f4197a18340110a181e8fb775496506a007a94b46f9f80f2a35918f9919081900360600190a2505050505050505050565b600080600080600080611bb78888612c44565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff9b909b16835260019a8b01909152902080549801549798680100000000000000008902987801000000000000000000000000000000000000000000000000900467ffffffffffffffff16975095505050505050565b6000808260068167ffffffffffffffff1610158015611c5e5750622932e08167ffffffffffffffff1611155b1515611c6957600080fd5b73ffffffffffffffffffffffffffffffffffffffff86161515611c8b57600080fd5b73ffffffffffffffffffffffffffffffffffffffff85161515611cad57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8681169086161415611cd357600080fd5b611cdd8686612c44565b6000818152600360205260409020805491945092507801000000000000000000000000000000000000000000000000900460ff1615611d1b57600080fd5b81547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff4367ffffffffffffffff908116700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff91881667ffffffffffffffff19909416841791909116171678010000000000000000000000000000000000000000000000001783556040805173ffffffffffffffffffffffffffffffffffffffff808a16825288166020820152808201929092525184917f4d4097deeecde59dede1bb370eb147fc3fa969b7b6a6f89f95526635328e86df919081900360600190a2505050505050565b60408051808201909152600581527f302e342e5f000000000000000000000000000000000000000000000000000000602082015281565b600080600080611e5d338b612c44565b6000818152600360205260409020805491955092507801000000000000000000000000000000000000000000000000900460ff16600114611e9d57600080fd5b81547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff7fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff909116780200000000000000000000000000000000000000000000000017908116680100000000000000004367ffffffffffffffff9384160183160217835560009088161115612036575073ffffffffffffffffffffffffffffffffffffffff8916600090815260018201602052604090208154611f869085908b908b908b90700100000000000000000000000000000000900467ffffffffffffffff168b8b613576565b925073ffffffffffffffffffffffffffffffffffffffff8a811690841614611fad57600080fd5b611fb789896136a5565b60018201805467ffffffffffffffff8a16780100000000000000000000000000000000000000000000000002680100000000000000009093047fffffffffffffffff0000000000000000000000000000000000000000000000009091161777ffffffffffffffffffffffffffffffffffffffffffffffff169190911790555b60408051338152602081018a90528082018b9052905185917f69610baaace24c039f891a11b42c0b1df1496ab0db38b0c4ee4ed33d6d53da1a919081900360600190a250505050505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff1633146120aa57600080fd5b6120e960008585858080601f01602080910402602001604051908101604052809392919081815260200183838082843750600094506134269350505050565b506001949350505050565b6000806000806000806121078c8a612c44565b6000818152600360205260409020805491955093507801000000000000000000000000000000000000000000000000900460ff1660021461214757600080fd5b82544367ffffffffffffffff6801000000000000000090920482166005019091161061217257600080fd5b505073ffffffffffffffffffffffffffffffffffffffff808b166000908152600183016020526040808220928a16825290206121ae8b8b6136a5565b6001830154680100000000000000000267ffffffffffffffff199081169116146121d757600080fd5b6121e188886136a5565b6001820154680100000000000000000267ffffffffffffffff1990811691161461220a57600080fd5b805482548981018d81039850910195508b111561222657600095505b6122308686613744565b73ffffffffffffffffffffffffffffffffffffffff808e1660009081526001808701602090815260408084208481558301849055938e1683528383208381559091018290558782526003905290812080547fffffffffffffff0000000000000000000000000000000000000000000000000016905581870399509096508611156123af576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8d886040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561237857600080fd5b505af115801561238c573d6000803e3d6000fd5b505050506040513d60208110156123a257600080fd5b505115156123af57600080fd5b600088111561246e5760008054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d81166004830152602482018d90529151919092169263a9059cbb92604480820193602093909283900390910190829087803b15801561243757600080fd5b505af115801561244b573d6000803e3d6000fd5b505050506040513d602081101561246157600080fd5b5051151561246e57600080fd5b60408051878152602081018a9052815186927ff94fb5c0628a82dc90648e8dc5e983f632633b0d26603d64e8cc042ca0790aa4928290030190a2505050505050505050505050565b6000806000806124c68c8c612c44565b935060036000856000191660001916815260200190815260200160002091508160010160008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508160000160189054906101000a900460ff1660ff16600214151561254b57600080fd5b815468010000000000000000900467ffffffffffffffff1692504383101561257257600080fd5b8154600267ffffffffffffffff9182160484031643101561259257600080fd5b600181015467ffffffffffffffff78010000000000000000000000000000000000000000000000009091048116908916116125cc57600080fd5b6125f4848b8b8b8660000160109054906101000a900467ffffffffffffffff168c8c8c61375c565b73ffffffffffffffffffffffffffffffffffffffff8c811691161461261857600080fd5b61263f848b8b8b8660000160109054906101000a900467ffffffffffffffff168c8c613576565b73ffffffffffffffffffffffffffffffffffffffff8d811691161461266357600080fd5b61266d8a8a6136a5565b60018201805467ffffffffffffffff8b16780100000000000000000000000000000000000000000000000002680100000000000000009093047fffffffffffffffff0000000000000000000000000000000000000000000000009091161777ffffffffffffffffffffffffffffffffffffffffffffffff169190911790556040805173ffffffffffffffffffffffffffffffffffffffff8e168152602081018b90528082018c9052905185917f910c9237f4197a18340110a181e8fb775496506a007a94b46f9f80f2a35918f9919081900360600190a2505050505050505050505050565b60008060008060008060006127678989612c44565b600081815260036020526040902054909a67ffffffffffffffff68010000000000000000830481169b50700100000000000000000000000000000000830481169a5060ff780100000000000000000000000000000000000000000000000084041699509091169650945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b61280284848484336001613923565b50505050565b600080600080600080600061281d8e8e612c44565b965060036000886000191660001916815260200190815260200160002091508160010160008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050438260000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16101515156128b057600080fd5b81547801000000000000000000000000000000000000000000000000900460ff166002146128dd57600080fd5b600154604080517fc1f62946000000000000000000000000000000000000000000000000000000008152600481018c9052905173ffffffffffffffffffffffffffffffffffffffff9092169163c1f62946916024808201926020929091908290030181600087803b15801561295157600080fd5b505af1158015612965573d6000803e3d6000fd5b505050506040513d602081101561297b57600080fd5b5051935060008411801561298f57508a8411155b151561299a57600080fd5b6040805160208082018e90528183018d905260608083018d905283518084039091018152608090920192839052815191929182918401908083835b602083106129f45780518252601f1990920191602091820191016129d5565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209450612a2d8589613c1a565b9250612a398c846136a5565b6001820154680100000000000000000267ffffffffffffffff19908116911614612a6257600080fd5b60018101546040805167ffffffffffffffff78010000000000000000000000000000000000000000000000009384900416909202602080840191909152602880840189905282518085039091018152604890930191829052825182918401908083835b60208310612ae45780518252601f199092019160209182019101612ac5565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120600081815260028801909252929020549199505060ff16159150612b31905057600080fd5b6000868152600282016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790559a89019a612b788c846136a5565b8160010160006101000a81548177ffffffffffffffffffffffffffffffffffffffffffffffff0219169083680100000000000000009004021790555086600019167f9e3b094fde58f3a83bd8b77d0a995fdb71f3169c6fa7e6d386e9f5902841e5ff8f878f604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018360001916600019168152602001828152602001935050505060405180910390a25050505050505050505050505050565b60008173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161015612d35576040805173ffffffffffffffffffffffffffffffffffffffff8581166c0100000000000000000000000090810260208085019190915291861681026034840152300260488301528251808303603c018152605c90920192839052815191929182918401908083835b60208310612d015780518252601f199092019160209182019101612ce2565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050612dbd565b604080516c0100000000000000000000000073ffffffffffffffffffffffffffffffffffffffff808616820260208085019190915290871682026034840152309190910260488301528251603c818403018152605c909201928390528151919291829184019080838360208310612d015780518252601f199092019160209182019101612ce2565b92915050565b6000806040805190810160405280601d81526020017f19536d61727452616964656e205369676e6564204d6573736167653a0a0000008152508787878b886002546040516020018088805190602001908083835b60208310612e365780518252601f199092019160209182019101612e17565b6001836020036101000a0380198251168184511680821785525050505050509050018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140186815260200185815260200184600019166000191681526020018367ffffffffffffffff1667ffffffffffffffff1678010000000000000000000000000000000000000000000000000281526008018281526020019750505050505050506040516020818303038152906040526040518082805190602001908083835b60208310612f355780518252601f199092019160209182019101612f16565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050612f6e8184613489565b98975050505050505050565b6000808080808711612f8b57600080fd5b612f958989612c44565b600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8e16845260018101909252909120805496509194509250905084156130965760008054604080517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152306024830152604482018c9052915191909216926323b872dd92606480820193602093909283900390910190829087803b15801561305f57600080fd5b505af1158015613073573d6000803e3d6000fd5b505050506040513d602081101561308957600080fd5b5051151561309657600080fd5b81547801000000000000000000000000000000000000000000000000900460ff166001146130c357600080fd5b9286018084556040805173ffffffffffffffffffffffffffffffffffffffff8b16815260208101839052815192959285927f0346e981e2bfa2366dc2307a8f1fa24779830a01121b1275fe565c6b98bb4d34928290030190a2505050505050505050565b6000806040805190810160405280601d81526020017f19536d61727452616964656e205369676e6564204d6573736167653a0a000000815250868887600254886040516020018087805190602001908083835b602083106131995780518252601f19909201916020918201910161317a565b51815160209384036101000a6000190180199092169116179052920197885250868101959095525067ffffffffffffffff92909216780100000000000000000000000000000000000000000000000002604080860191909152604885019190915260688085019290925280518085039092018252608890930192839052805190935082918401908083835b602083106132435780518252601f199092019160209182019101613224565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061327c8184613489565b979650505050505050565b6000806040805190810160405280601d81526020017f19536d61727452616964656e205369676e6564204d6573736167653a0a000000815250888888888d896002546040516020018089805190602001908083835b602083106132fb5780518252601f1990920191602091820191016132dc565b51815160209384036101000a600019018019909216911617905273ffffffffffffffffffffffffffffffffffffffff9b8c166c0100000000000000000000000090810292909401918252601482019a909a5297909916026034870152506048850193909352606884019190915267ffffffffffffffff16780100000000000000000000000000000000000000000000000002608883015260908083019190915260408051808403909201825260b090920191829052805190945090925082918401908083835b602083106133e05780518252601f1990920191602091820191016133c1565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090506134198184613489565b9998505050505050505050565b602082015160008080600184141561345b5761344186613d6b565b919450925090506134568383838a8c8a613923565b61347f565b83600214156101535761346d86613d7f565b90935091506134568383898b89612f7a565b5050505050505050565b6000806000808451604114151561349f57600080fd5b50505060208201516040830151606084015160001a601b60ff821610156134c457601b015b8060ff16601b14806134d957508060ff16601c145b15156134e457600080fd5b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af115801561353e573d6000803e3d6000fd5b5050604051601f19015194505073ffffffffffffffffffffffffffffffffffffffff8416151561356d57600080fd5b50505092915050565b6000806040805190810160405280601d81526020017f19536d61727452616964656e205369676e6564204d6573736167653a0a000000815250888888878d8a6002546040516020018089805190602001908083835b602083106135ea5780518252601f1990920191602091820191016135cb565b51815160001960209485036101000a019081169019919091161790529201998a52508881019790975250780100000000000000000000000000000000000000000000000067ffffffffffffffff95861681026040808a01919091526048890195909552606888019390935293160260888501526090808501929092528051808503909201825260b09093019283905280519093508291840190808383602083106133e05780518252601f1990920191602091820191016133c1565b6000811580156136b3575082155b156136c057506000612dbd565b604080516020808201859052818301869052825180830384018152606090920192839052815191929182918401908083835b602083106137115780518252601f1990920191602091820191016136f2565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091209695505050505050565b60008183116137535782613755565b815b9392505050565b6000806040805190810160405280601d81526020017f19536d61727452616964656e205369676e6564204d6573736167653a0a000000815250898989888e8b6002548b604051602001808a805190602001908083835b602083106137d15780518252601f1990920191602091820191016137b2565b51815160001960209485036101000a0190811690199190911617905292018b81528083018b9052780100000000000000000000000000000000000000000000000067ffffffffffffffff808c1682026040840152604883018b9052606883018a9052881602608882015260908101869052845160b090910192850191508083835b602083106138715780518252601f199092019160209182019101613852565b6001836020036101000a03801982511681845116808217855250505050505090500199505050505050505050506040516020818303038152906040526040518082805190602001908083835b602083106138dc5780518252601f1990920191602091820191016138bd565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090506139158184613489565b9a9950505050505050505050565b60008060008660068167ffffffffffffffff16101580156139515750622932e08167ffffffffffffffff1611155b151561395c57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8a16151561397e57600080fd5b73ffffffffffffffffffffffffffffffffffffffff891615156139a057600080fd5b73ffffffffffffffffffffffffffffffffffffffff8a8116908a1614156139c657600080fd5b600087116139d357600080fd5b6139dd8a8a612c44565b600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8f16845260018101909252909120815492965090945092507801000000000000000000000000000000000000000000000000900460ff1615613a4457600080fd5b82547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff4367ffffffffffffffff908116700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff918c1667ffffffffffffffff199094169390931716919091171678010000000000000000000000000000000000000000000000001783558415613ba15760008054604080517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152306024830152604482018c9052915191909216926323b872dd92606480820193602093909283900390910190829087803b158015613b6a57600080fd5b505af1158015613b7e573d6000803e3d6000fd5b505050506040513d6020811015613b9457600080fd5b50511515613ba157600080fd5b8682556040805173ffffffffffffffffffffffffffffffffffffffff808d1682528b16602082015267ffffffffffffffff8a168183015260608101899052905185917fcac76648b0a531becb6e54db5fe838853fdc47ef130aab3566114ee7c739d0a0919081900360800190a250505050505050505050565b600080600060208451811515613c2c57fe5b0615613c3757600080fd5b602091505b83518211613d6257508281015180851015613cd657604080516020808201889052818301849052825180830384018152606090920192839052815191929182918401908083835b60208310613ca25780518252601f199092019160209182019101613c83565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209450613d57565b604080516020808201849052818301889052825180830384018152606090920192839052815191929182918401908083835b60208310613d275780518252601f199092019160209182019101613d08565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902094505b602082019150613c3c565b50929392505050565b604081015160608201516080909201519092565b604081015160608201519150915600a165627a7a723058208976b94ce2133c9760307de2a94560bcac6dd4a4d450a369447563767b1f45920029a165627a7a72305820db19656226d0bc9fe16e5fedd3c387fde8234882f15fd4252a85a9284d7686970029`

// DeployTokenNetworkRegistry deploys a new Ethereum contract, binding an instance of TokenNetworkRegistry to it.
func DeployTokenNetworkRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _secret_registry_address common.Address, _chain_id *big.Int) (common.Address, *types.Transaction, *TokenNetworkRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenNetworkRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenNetworkRegistryBin), backend, _secret_registry_address, _chain_id)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenNetworkRegistry{TokenNetworkRegistryCaller: TokenNetworkRegistryCaller{contract: contract}, TokenNetworkRegistryTransactor: TokenNetworkRegistryTransactor{contract: contract}, TokenNetworkRegistryFilterer: TokenNetworkRegistryFilterer{contract: contract}}, nil
}

// TokenNetworkRegistry is an auto generated Go binding around an Ethereum contract.
type TokenNetworkRegistry struct {
	TokenNetworkRegistryCaller     // Read-only binding to the contract
	TokenNetworkRegistryTransactor // Write-only binding to the contract
	TokenNetworkRegistryFilterer   // Log filterer for contract events
}

// TokenNetworkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenNetworkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenNetworkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenNetworkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenNetworkRegistrySession struct {
	Contract     *TokenNetworkRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenNetworkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenNetworkRegistryCallerSession struct {
	Contract *TokenNetworkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TokenNetworkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenNetworkRegistryTransactorSession struct {
	Contract     *TokenNetworkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TokenNetworkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenNetworkRegistryRaw struct {
	Contract *TokenNetworkRegistry // Generic contract binding to access the raw methods on
}

// TokenNetworkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenNetworkRegistryCallerRaw struct {
	Contract *TokenNetworkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenNetworkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenNetworkRegistryTransactorRaw struct {
	Contract *TokenNetworkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenNetworkRegistry creates a new instance of TokenNetworkRegistry, bound to a specific deployed contract.
func NewTokenNetworkRegistry(address common.Address, backend bind.ContractBackend) (*TokenNetworkRegistry, error) {
	contract, err := bindTokenNetworkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkRegistry{TokenNetworkRegistryCaller: TokenNetworkRegistryCaller{contract: contract}, TokenNetworkRegistryTransactor: TokenNetworkRegistryTransactor{contract: contract}, TokenNetworkRegistryFilterer: TokenNetworkRegistryFilterer{contract: contract}}, nil
}

// NewTokenNetworkRegistryCaller creates a new read-only instance of TokenNetworkRegistry, bound to a specific deployed contract.
func NewTokenNetworkRegistryCaller(address common.Address, caller bind.ContractCaller) (*TokenNetworkRegistryCaller, error) {
	contract, err := bindTokenNetworkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkRegistryCaller{contract: contract}, nil
}

// NewTokenNetworkRegistryTransactor creates a new write-only instance of TokenNetworkRegistry, bound to a specific deployed contract.
func NewTokenNetworkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenNetworkRegistryTransactor, error) {
	contract, err := bindTokenNetworkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkRegistryTransactor{contract: contract}, nil
}

// NewTokenNetworkRegistryFilterer creates a new log filterer instance of TokenNetworkRegistry, bound to a specific deployed contract.
func NewTokenNetworkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenNetworkRegistryFilterer, error) {
	contract, err := bindTokenNetworkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkRegistryFilterer{contract: contract}, nil
}

// bindTokenNetworkRegistry binds a generic wrapper to an already deployed contract.
func bindTokenNetworkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenNetworkRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenNetworkRegistry *TokenNetworkRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenNetworkRegistry.Contract.TokenNetworkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenNetworkRegistry *TokenNetworkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.TokenNetworkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenNetworkRegistry *TokenNetworkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.TokenNetworkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenNetworkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenNetworkRegistry *TokenNetworkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenNetworkRegistry *TokenNetworkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetworkRegistry *TokenNetworkRegistryCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenNetworkRegistry.contract.Call(opts, out, "chain_id")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) ChainId() (*big.Int, error) {
	return _TokenNetworkRegistry.Contract.ChainId(&_TokenNetworkRegistry.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerSession) ChainId() (*big.Int, error) {
	return _TokenNetworkRegistry.Contract.ChainId(&_TokenNetworkRegistry.CallOpts)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetworkRegistry *TokenNetworkRegistryCaller) ContractExists(opts *bind.CallOpts, contract_address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenNetworkRegistry.contract.Call(opts, out, "contractExists", contract_address)
	return *ret0, err
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) ContractExists(contract_address common.Address) (bool, error) {
	return _TokenNetworkRegistry.Contract.ContractExists(&_TokenNetworkRegistry.CallOpts, contract_address)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerSession) ContractExists(contract_address common.Address) (bool, error) {
	return _TokenNetworkRegistry.Contract.ContractExists(&_TokenNetworkRegistry.CallOpts, contract_address)
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetworkRegistry *TokenNetworkRegistryCaller) ContractVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenNetworkRegistry.contract.Call(opts, out, "contract_version")
	return *ret0, err
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) ContractVersion() (string, error) {
	return _TokenNetworkRegistry.Contract.ContractVersion(&_TokenNetworkRegistry.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerSession) ContractVersion() (string, error) {
	return _TokenNetworkRegistry.Contract.ContractVersion(&_TokenNetworkRegistry.CallOpts)
}

// SecretRegistryAddress is a free data retrieval call binding the contract method 0xd0ad4bec.
//
// Solidity: function secret_registry_address() constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistryCaller) SecretRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenNetworkRegistry.contract.Call(opts, out, "secret_registry_address")
	return *ret0, err
}

// SecretRegistryAddress is a free data retrieval call binding the contract method 0xd0ad4bec.
//
// Solidity: function secret_registry_address() constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) SecretRegistryAddress() (common.Address, error) {
	return _TokenNetworkRegistry.Contract.SecretRegistryAddress(&_TokenNetworkRegistry.CallOpts)
}

// SecretRegistryAddress is a free data retrieval call binding the contract method 0xd0ad4bec.
//
// Solidity: function secret_registry_address() constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerSession) SecretRegistryAddress() (common.Address, error) {
	return _TokenNetworkRegistry.Contract.SecretRegistryAddress(&_TokenNetworkRegistry.CallOpts)
}

// TokenToTokenNetworks is a free data retrieval call binding the contract method 0x0fabd9e7.
//
// Solidity: function token_to_token_networks( address) constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistryCaller) TokenToTokenNetworks(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenNetworkRegistry.contract.Call(opts, out, "token_to_token_networks", arg0)
	return *ret0, err
}

// TokenToTokenNetworks is a free data retrieval call binding the contract method 0x0fabd9e7.
//
// Solidity: function token_to_token_networks( address) constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) TokenToTokenNetworks(arg0 common.Address) (common.Address, error) {
	return _TokenNetworkRegistry.Contract.TokenToTokenNetworks(&_TokenNetworkRegistry.CallOpts, arg0)
}

// TokenToTokenNetworks is a free data retrieval call binding the contract method 0x0fabd9e7.
//
// Solidity: function token_to_token_networks( address) constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerSession) TokenToTokenNetworks(arg0 common.Address) (common.Address, error) {
	return _TokenNetworkRegistry.Contract.TokenToTokenNetworks(&_TokenNetworkRegistry.CallOpts, arg0)
}

// CreateERC20TokenNetwork is a paid mutator transaction binding the contract method 0x4cf71a04.
//
// Solidity: function createERC20TokenNetwork(_token_address address) returns(token_network_address address)
func (_TokenNetworkRegistry *TokenNetworkRegistryTransactor) CreateERC20TokenNetwork(opts *bind.TransactOpts, _token_address common.Address) (*types.Transaction, error) {
	return _TokenNetworkRegistry.contract.Transact(opts, "createERC20TokenNetwork", _token_address)
}

// CreateERC20TokenNetwork is a paid mutator transaction binding the contract method 0x4cf71a04.
//
// Solidity: function createERC20TokenNetwork(_token_address address) returns(token_network_address address)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) CreateERC20TokenNetwork(_token_address common.Address) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.CreateERC20TokenNetwork(&_TokenNetworkRegistry.TransactOpts, _token_address)
}

// CreateERC20TokenNetwork is a paid mutator transaction binding the contract method 0x4cf71a04.
//
// Solidity: function createERC20TokenNetwork(_token_address address) returns(token_network_address address)
func (_TokenNetworkRegistry *TokenNetworkRegistryTransactorSession) CreateERC20TokenNetwork(_token_address common.Address) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.CreateERC20TokenNetwork(&_TokenNetworkRegistry.TransactOpts, _token_address)
}

// TokenNetworkRegistryTokenNetworkCreatedIterator is returned from FilterTokenNetworkCreated and is used to iterate over the raw logs and unpacked data for TokenNetworkCreated events raised by the TokenNetworkRegistry contract.
type TokenNetworkRegistryTokenNetworkCreatedIterator struct {
	Event *TokenNetworkRegistryTokenNetworkCreated // Event containing the contract specifics and raw log

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
func (it *TokenNetworkRegistryTokenNetworkCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkRegistryTokenNetworkCreated)
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
		it.Event = new(TokenNetworkRegistryTokenNetworkCreated)
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
func (it *TokenNetworkRegistryTokenNetworkCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkRegistryTokenNetworkCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkRegistryTokenNetworkCreated represents a TokenNetworkCreated event raised by the TokenNetworkRegistry contract.
type TokenNetworkRegistryTokenNetworkCreated struct {
	TokenAddress        common.Address
	TokenNetworkAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokenNetworkCreated is a free log retrieval operation binding the contract event 0xf11a7558a113d9627989c5edf26cbd19143b7375248e621c8e30ac9e0847dc3f.
//
// Solidity: e TokenNetworkCreated(token_address indexed address, token_network_address indexed address)
func (_TokenNetworkRegistry *TokenNetworkRegistryFilterer) FilterTokenNetworkCreated(opts *bind.FilterOpts, token_address []common.Address, token_network_address []common.Address) (*TokenNetworkRegistryTokenNetworkCreatedIterator, error) {

	var token_addressRule []interface{}
	for _, token_addressItem := range token_address {
		token_addressRule = append(token_addressRule, token_addressItem)
	}
	var token_network_addressRule []interface{}
	for _, token_network_addressItem := range token_network_address {
		token_network_addressRule = append(token_network_addressRule, token_network_addressItem)
	}

	logs, sub, err := _TokenNetworkRegistry.contract.FilterLogs(opts, "TokenNetworkCreated", token_addressRule, token_network_addressRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkRegistryTokenNetworkCreatedIterator{contract: _TokenNetworkRegistry.contract, event: "TokenNetworkCreated", logs: logs, sub: sub}, nil
}

// WatchTokenNetworkCreated is a free log subscription operation binding the contract event 0xf11a7558a113d9627989c5edf26cbd19143b7375248e621c8e30ac9e0847dc3f.
//
// Solidity: e TokenNetworkCreated(token_address indexed address, token_network_address indexed address)
func (_TokenNetworkRegistry *TokenNetworkRegistryFilterer) WatchTokenNetworkCreated(opts *bind.WatchOpts, sink chan<- *TokenNetworkRegistryTokenNetworkCreated, token_address []common.Address, token_network_address []common.Address) (event.Subscription, error) {

	var token_addressRule []interface{}
	for _, token_addressItem := range token_address {
		token_addressRule = append(token_addressRule, token_addressItem)
	}
	var token_network_addressRule []interface{}
	for _, token_network_addressItem := range token_network_address {
		token_network_addressRule = append(token_network_addressRule, token_network_addressItem)
	}

	logs, sub, err := _TokenNetworkRegistry.contract.WatchLogs(opts, "TokenNetworkCreated", token_addressRule, token_network_addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkRegistryTokenNetworkCreated)
				if err := _TokenNetworkRegistry.contract.UnpackLog(event, "TokenNetworkCreated", log); err != nil {
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

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"contract_address\",\"type\":\"address\"}],\"name\":\"contractExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contract_version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
const UtilsBin = `0x608060405234801561001057600080fd5b50610187806100206000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637709bc788114610050578063b32c65c814610092575b600080fd5b34801561005c57600080fd5b5061007e73ffffffffffffffffffffffffffffffffffffffff6004351661011c565b604080519115158252519081900360200190f35b34801561009e57600080fd5b506100a7610124565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e15781810151838201526020016100c9565b50505050905090810190601f16801561010e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6000903b1190565b60408051808201909152600581527f302e332e5f0000000000000000000000000000000000000000000000000000006020820152815600a165627a7a723058207f85681e0290e75ebbbf8e43ecfa0fb5d3b60a75b5c995b9ae631d13f9ca3d7b0029`

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
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_Utils *UtilsCaller) ContractExists(opts *bind.CallOpts, contract_address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Utils.contract.Call(opts, out, "contractExists", contract_address)
	return *ret0, err
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_Utils *UtilsSession) ContractExists(contract_address common.Address) (bool, error) {
	return _Utils.Contract.ContractExists(&_Utils.CallOpts, contract_address)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_Utils *UtilsCallerSession) ContractExists(contract_address common.Address) (bool, error) {
	return _Utils.Contract.ContractExists(&_Utils.CallOpts, contract_address)
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_Utils *UtilsCaller) ContractVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Utils.contract.Call(opts, out, "contract_version")
	return *ret0, err
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_Utils *UtilsSession) ContractVersion() (string, error) {
	return _Utils.Contract.ContractVersion(&_Utils.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_Utils *UtilsCallerSession) ContractVersion() (string, error) {
	return _Utils.Contract.ContractVersion(&_Utils.CallOpts)
}
