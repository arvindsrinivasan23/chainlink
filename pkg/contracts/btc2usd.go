// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_timeout\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"_minSubmissionValue\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_maxSubmissionValue\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AddedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AvailableFundsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CheckAccessDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CheckAccessEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"OracleAdminUpdateRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"OracleAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"OraclePermissionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RemovedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"name\":\"RequesterPermissionsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"paymentAmount\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"minSubmissionCount\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"maxSubmissionCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"restartDelay\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timeout\",\"type\":\"uint32\"}],\"name\":\"RoundDetailsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"submission\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"round\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"SubmissionReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"acceptAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocatedFunds\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableFunds\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_removed\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_added\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_addedAdmins\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"_minSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_restartDelay\",\"type\":\"uint32\"}],\"name\":\"changeOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableAccessCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableAccessCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkToken\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSubmissionCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSubmissionValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSubmissionCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSubmissionValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_queriedRoundId\",\"type\":\"uint32\"}],\"name\":\"oracleRoundState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_eligibleToSubmit\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_roundId\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"_latestSubmission\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"_startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_timeout\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"_availableFunds\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_oracleCount\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"removeAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNewRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restartDelay\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_authorized\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_delay\",\"type\":\"uint32\"}],\"name\":\"setRequesterPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newValidator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_submission\",\"type\":\"int256\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAvailableFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_minSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_restartDelay\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_timeout\",\"type\":\"uint32\"}],\"name\":\"updateFutureRounds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"withdrawablePayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() view returns(uint128)
func (_Contracts *ContractsCaller) AllocatedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allocatedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() view returns(uint128)
func (_Contracts *ContractsSession) AllocatedFunds() (*big.Int, error) {
	return _Contracts.Contract.AllocatedFunds(&_Contracts.CallOpts)
}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() view returns(uint128)
func (_Contracts *ContractsCallerSession) AllocatedFunds() (*big.Int, error) {
	return _Contracts.Contract.AllocatedFunds(&_Contracts.CallOpts)
}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() view returns(uint128)
func (_Contracts *ContractsCaller) AvailableFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "availableFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() view returns(uint128)
func (_Contracts *ContractsSession) AvailableFunds() (*big.Int, error) {
	return _Contracts.Contract.AvailableFunds(&_Contracts.CallOpts)
}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() view returns(uint128)
func (_Contracts *ContractsCallerSession) AvailableFunds() (*big.Int, error) {
	return _Contracts.Contract.AvailableFunds(&_Contracts.CallOpts)
}

// CheckEnabled is a free data retrieval call binding the contract method 0xdc7f0124.
//
// Solidity: function checkEnabled() view returns(bool)
func (_Contracts *ContractsCaller) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "checkEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckEnabled is a free data retrieval call binding the contract method 0xdc7f0124.
//
// Solidity: function checkEnabled() view returns(bool)
func (_Contracts *ContractsSession) CheckEnabled() (bool, error) {
	return _Contracts.Contract.CheckEnabled(&_Contracts.CallOpts)
}

// CheckEnabled is a free data retrieval call binding the contract method 0xdc7f0124.
//
// Solidity: function checkEnabled() view returns(bool)
func (_Contracts *ContractsCallerSession) CheckEnabled() (bool, error) {
	return _Contracts.Contract.CheckEnabled(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCallerSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Contracts *ContractsCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Contracts *ContractsSession) Description() (string, error) {
	return _Contracts.Contract.Description(&_Contracts.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Contracts *ContractsCallerSession) Description() (string, error) {
	return _Contracts.Contract.Description(&_Contracts.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) view returns(address)
func (_Contracts *ContractsCaller) GetAdmin(opts *bind.CallOpts, _oracle common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAdmin", _oracle)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) view returns(address)
func (_Contracts *ContractsSession) GetAdmin(_oracle common.Address) (common.Address, error) {
	return _Contracts.Contract.GetAdmin(&_Contracts.CallOpts, _oracle)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) view returns(address)
func (_Contracts *ContractsCallerSession) GetAdmin(_oracle common.Address) (common.Address, error) {
	return _Contracts.Contract.GetAdmin(&_Contracts.CallOpts, _oracle)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_Contracts *ContractsCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAnswer", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_Contracts *ContractsSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetAnswer(&_Contracts.CallOpts, _roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_Contracts *ContractsCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetAnswer(&_Contracts.CallOpts, _roundId)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Contracts *ContractsCaller) GetOracles(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOracles")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Contracts *ContractsSession) GetOracles() ([]common.Address, error) {
	return _Contracts.Contract.GetOracles(&_Contracts.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Contracts *ContractsCallerSession) GetOracles() ([]common.Address, error) {
	return _Contracts.Contract.GetOracles(&_Contracts.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Contracts *ContractsCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Contracts *ContractsSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Contracts.Contract.GetRoundData(&_Contracts.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Contracts *ContractsCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Contracts.Contract.GetRoundData(&_Contracts.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_Contracts *ContractsCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getTimestamp", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_Contracts *ContractsSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetTimestamp(&_Contracts.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetTimestamp(&_Contracts.CallOpts, _roundId)
}

// HasAccess is a free data retrieval call binding the contract method 0x6b14daf8.
//
// Solidity: function hasAccess(address _user, bytes _calldata) view returns(bool)
func (_Contracts *ContractsCaller) HasAccess(opts *bind.CallOpts, _user common.Address, _calldata []byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "hasAccess", _user, _calldata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAccess is a free data retrieval call binding the contract method 0x6b14daf8.
//
// Solidity: function hasAccess(address _user, bytes _calldata) view returns(bool)
func (_Contracts *ContractsSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _Contracts.Contract.HasAccess(&_Contracts.CallOpts, _user, _calldata)
}

// HasAccess is a free data retrieval call binding the contract method 0x6b14daf8.
//
// Solidity: function hasAccess(address _user, bytes _calldata) view returns(bool)
func (_Contracts *ContractsCallerSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _Contracts.Contract.HasAccess(&_Contracts.CallOpts, _user, _calldata)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_Contracts *ContractsCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_Contracts *ContractsSession) LatestAnswer() (*big.Int, error) {
	return _Contracts.Contract.LatestAnswer(&_Contracts.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_Contracts *ContractsCallerSession) LatestAnswer() (*big.Int, error) {
	return _Contracts.Contract.LatestAnswer(&_Contracts.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_Contracts *ContractsCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_Contracts *ContractsSession) LatestRound() (*big.Int, error) {
	return _Contracts.Contract.LatestRound(&_Contracts.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_Contracts *ContractsCallerSession) LatestRound() (*big.Int, error) {
	return _Contracts.Contract.LatestRound(&_Contracts.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Contracts *ContractsCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Contracts *ContractsSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Contracts.Contract.LatestRoundData(&_Contracts.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Contracts *ContractsCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Contracts.Contract.LatestRoundData(&_Contracts.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_Contracts *ContractsCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_Contracts *ContractsSession) LatestTimestamp() (*big.Int, error) {
	return _Contracts.Contract.LatestTimestamp(&_Contracts.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_Contracts *ContractsCallerSession) LatestTimestamp() (*big.Int, error) {
	return _Contracts.Contract.LatestTimestamp(&_Contracts.CallOpts)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_Contracts *ContractsCaller) LinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "linkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_Contracts *ContractsSession) LinkToken() (common.Address, error) {
	return _Contracts.Contract.LinkToken(&_Contracts.CallOpts)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_Contracts *ContractsCallerSession) LinkToken() (common.Address, error) {
	return _Contracts.Contract.LinkToken(&_Contracts.CallOpts)
}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() view returns(uint32)
func (_Contracts *ContractsCaller) MaxSubmissionCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "maxSubmissionCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() view returns(uint32)
func (_Contracts *ContractsSession) MaxSubmissionCount() (uint32, error) {
	return _Contracts.Contract.MaxSubmissionCount(&_Contracts.CallOpts)
}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() view returns(uint32)
func (_Contracts *ContractsCallerSession) MaxSubmissionCount() (uint32, error) {
	return _Contracts.Contract.MaxSubmissionCount(&_Contracts.CallOpts)
}

// MaxSubmissionValue is a free data retrieval call binding the contract method 0x23ca2903.
//
// Solidity: function maxSubmissionValue() view returns(int256)
func (_Contracts *ContractsCaller) MaxSubmissionValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "maxSubmissionValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSubmissionValue is a free data retrieval call binding the contract method 0x23ca2903.
//
// Solidity: function maxSubmissionValue() view returns(int256)
func (_Contracts *ContractsSession) MaxSubmissionValue() (*big.Int, error) {
	return _Contracts.Contract.MaxSubmissionValue(&_Contracts.CallOpts)
}

// MaxSubmissionValue is a free data retrieval call binding the contract method 0x23ca2903.
//
// Solidity: function maxSubmissionValue() view returns(int256)
func (_Contracts *ContractsCallerSession) MaxSubmissionValue() (*big.Int, error) {
	return _Contracts.Contract.MaxSubmissionValue(&_Contracts.CallOpts)
}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() view returns(uint32)
func (_Contracts *ContractsCaller) MinSubmissionCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "minSubmissionCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() view returns(uint32)
func (_Contracts *ContractsSession) MinSubmissionCount() (uint32, error) {
	return _Contracts.Contract.MinSubmissionCount(&_Contracts.CallOpts)
}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() view returns(uint32)
func (_Contracts *ContractsCallerSession) MinSubmissionCount() (uint32, error) {
	return _Contracts.Contract.MinSubmissionCount(&_Contracts.CallOpts)
}

// MinSubmissionValue is a free data retrieval call binding the contract method 0x7c2b0b21.
//
// Solidity: function minSubmissionValue() view returns(int256)
func (_Contracts *ContractsCaller) MinSubmissionValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "minSubmissionValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSubmissionValue is a free data retrieval call binding the contract method 0x7c2b0b21.
//
// Solidity: function minSubmissionValue() view returns(int256)
func (_Contracts *ContractsSession) MinSubmissionValue() (*big.Int, error) {
	return _Contracts.Contract.MinSubmissionValue(&_Contracts.CallOpts)
}

// MinSubmissionValue is a free data retrieval call binding the contract method 0x7c2b0b21.
//
// Solidity: function minSubmissionValue() view returns(int256)
func (_Contracts *ContractsCallerSession) MinSubmissionValue() (*big.Int, error) {
	return _Contracts.Contract.MinSubmissionValue(&_Contracts.CallOpts)
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_Contracts *ContractsCaller) OracleCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "oracleCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_Contracts *ContractsSession) OracleCount() (uint8, error) {
	return _Contracts.Contract.OracleCount(&_Contracts.CallOpts)
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_Contracts *ContractsCallerSession) OracleCount() (uint8, error) {
	return _Contracts.Contract.OracleCount(&_Contracts.CallOpts)
}

// OracleRoundState is a free data retrieval call binding the contract method 0x88aa80e7.
//
// Solidity: function oracleRoundState(address _oracle, uint32 _queriedRoundId) view returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint8 _oracleCount, uint128 _paymentAmount)
func (_Contracts *ContractsCaller) OracleRoundState(opts *bind.CallOpts, _oracle common.Address, _queriedRoundId uint32) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint8
	PaymentAmount    *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "oracleRoundState", _oracle, _queriedRoundId)

	outstruct := new(struct {
		EligibleToSubmit bool
		RoundId          uint32
		LatestSubmission *big.Int
		StartedAt        uint64
		Timeout          uint64
		AvailableFunds   *big.Int
		OracleCount      uint8
		PaymentAmount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EligibleToSubmit = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RoundId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.LatestSubmission = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Timeout = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.AvailableFunds = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OracleCount = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.PaymentAmount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OracleRoundState is a free data retrieval call binding the contract method 0x88aa80e7.
//
// Solidity: function oracleRoundState(address _oracle, uint32 _queriedRoundId) view returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint8 _oracleCount, uint128 _paymentAmount)
func (_Contracts *ContractsSession) OracleRoundState(_oracle common.Address, _queriedRoundId uint32) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint8
	PaymentAmount    *big.Int
}, error) {
	return _Contracts.Contract.OracleRoundState(&_Contracts.CallOpts, _oracle, _queriedRoundId)
}

// OracleRoundState is a free data retrieval call binding the contract method 0x88aa80e7.
//
// Solidity: function oracleRoundState(address _oracle, uint32 _queriedRoundId) view returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint8 _oracleCount, uint128 _paymentAmount)
func (_Contracts *ContractsCallerSession) OracleRoundState(_oracle common.Address, _queriedRoundId uint32) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint8
	PaymentAmount    *big.Int
}, error) {
	return _Contracts.Contract.OracleRoundState(&_Contracts.CallOpts, _oracle, _queriedRoundId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_Contracts *ContractsCaller) PaymentAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "paymentAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_Contracts *ContractsSession) PaymentAmount() (*big.Int, error) {
	return _Contracts.Contract.PaymentAmount(&_Contracts.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_Contracts *ContractsCallerSession) PaymentAmount() (*big.Int, error) {
	return _Contracts.Contract.PaymentAmount(&_Contracts.CallOpts)
}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() view returns(uint32)
func (_Contracts *ContractsCaller) RestartDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "restartDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() view returns(uint32)
func (_Contracts *ContractsSession) RestartDelay() (uint32, error) {
	return _Contracts.Contract.RestartDelay(&_Contracts.CallOpts)
}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() view returns(uint32)
func (_Contracts *ContractsCallerSession) RestartDelay() (uint32, error) {
	return _Contracts.Contract.RestartDelay(&_Contracts.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint32)
func (_Contracts *ContractsCaller) Timeout(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "timeout")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint32)
func (_Contracts *ContractsSession) Timeout() (uint32, error) {
	return _Contracts.Contract.Timeout(&_Contracts.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint32)
func (_Contracts *ContractsCallerSession) Timeout() (uint32, error) {
	return _Contracts.Contract.Timeout(&_Contracts.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Contracts *ContractsCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Contracts *ContractsSession) Validator() (common.Address, error) {
	return _Contracts.Contract.Validator(&_Contracts.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Contracts *ContractsCallerSession) Validator() (common.Address, error) {
	return _Contracts.Contract.Validator(&_Contracts.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Contracts *ContractsCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Contracts *ContractsSession) Version() (*big.Int, error) {
	return _Contracts.Contract.Version(&_Contracts.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Contracts *ContractsCallerSession) Version() (*big.Int, error) {
	return _Contracts.Contract.Version(&_Contracts.CallOpts)
}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) view returns(uint256)
func (_Contracts *ContractsCaller) WithdrawablePayment(opts *bind.CallOpts, _oracle common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "withdrawablePayment", _oracle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) view returns(uint256)
func (_Contracts *ContractsSession) WithdrawablePayment(_oracle common.Address) (*big.Int, error) {
	return _Contracts.Contract.WithdrawablePayment(&_Contracts.CallOpts, _oracle)
}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) view returns(uint256)
func (_Contracts *ContractsCallerSession) WithdrawablePayment(_oracle common.Address) (*big.Int, error) {
	return _Contracts.Contract.WithdrawablePayment(&_Contracts.CallOpts, _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_Contracts *ContractsTransactor) AcceptAdmin(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "acceptAdmin", _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_Contracts *ContractsSession) AcceptAdmin(_oracle common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AcceptAdmin(&_Contracts.TransactOpts, _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_Contracts *ContractsTransactorSession) AcceptAdmin(_oracle common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AcceptAdmin(&_Contracts.TransactOpts, _oracle)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contracts *ContractsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contracts *ContractsSession) AcceptOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.AcceptOwnership(&_Contracts.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contracts *ContractsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.AcceptOwnership(&_Contracts.TransactOpts)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address _user) returns()
func (_Contracts *ContractsTransactor) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addAccess", _user)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address _user) returns()
func (_Contracts *ContractsSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddAccess(&_Contracts.TransactOpts, _user)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address _user) returns()
func (_Contracts *ContractsTransactorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddAccess(&_Contracts.TransactOpts, _user)
}

// ChangeOracles is a paid mutator transaction binding the contract method 0x3969c20f.
//
// Solidity: function changeOracles(address[] _removed, address[] _added, address[] _addedAdmins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_Contracts *ContractsTransactor) ChangeOracles(opts *bind.TransactOpts, _removed []common.Address, _added []common.Address, _addedAdmins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "changeOracles", _removed, _added, _addedAdmins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// ChangeOracles is a paid mutator transaction binding the contract method 0x3969c20f.
//
// Solidity: function changeOracles(address[] _removed, address[] _added, address[] _addedAdmins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_Contracts *ContractsSession) ChangeOracles(_removed []common.Address, _added []common.Address, _addedAdmins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeOracles(&_Contracts.TransactOpts, _removed, _added, _addedAdmins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// ChangeOracles is a paid mutator transaction binding the contract method 0x3969c20f.
//
// Solidity: function changeOracles(address[] _removed, address[] _added, address[] _addedAdmins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_Contracts *ContractsTransactorSession) ChangeOracles(_removed []common.Address, _added []common.Address, _addedAdmins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeOracles(&_Contracts.TransactOpts, _removed, _added, _addedAdmins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// DisableAccessCheck is a paid mutator transaction binding the contract method 0x0a756983.
//
// Solidity: function disableAccessCheck() returns()
func (_Contracts *ContractsTransactor) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "disableAccessCheck")
}

// DisableAccessCheck is a paid mutator transaction binding the contract method 0x0a756983.
//
// Solidity: function disableAccessCheck() returns()
func (_Contracts *ContractsSession) DisableAccessCheck() (*types.Transaction, error) {
	return _Contracts.Contract.DisableAccessCheck(&_Contracts.TransactOpts)
}

// DisableAccessCheck is a paid mutator transaction binding the contract method 0x0a756983.
//
// Solidity: function disableAccessCheck() returns()
func (_Contracts *ContractsTransactorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _Contracts.Contract.DisableAccessCheck(&_Contracts.TransactOpts)
}

// EnableAccessCheck is a paid mutator transaction binding the contract method 0x8038e4a1.
//
// Solidity: function enableAccessCheck() returns()
func (_Contracts *ContractsTransactor) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "enableAccessCheck")
}

// EnableAccessCheck is a paid mutator transaction binding the contract method 0x8038e4a1.
//
// Solidity: function enableAccessCheck() returns()
func (_Contracts *ContractsSession) EnableAccessCheck() (*types.Transaction, error) {
	return _Contracts.Contract.EnableAccessCheck(&_Contracts.TransactOpts)
}

// EnableAccessCheck is a paid mutator transaction binding the contract method 0x8038e4a1.
//
// Solidity: function enableAccessCheck() returns()
func (_Contracts *ContractsTransactorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _Contracts.Contract.EnableAccessCheck(&_Contracts.TransactOpts)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_Contracts *ContractsTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "onTokenTransfer", arg0, arg1, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_Contracts *ContractsSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnTokenTransfer(&_Contracts.TransactOpts, arg0, arg1, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_Contracts *ContractsTransactorSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.OnTokenTransfer(&_Contracts.TransactOpts, arg0, arg1, _data)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address _user) returns()
func (_Contracts *ContractsTransactor) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeAccess", _user)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address _user) returns()
func (_Contracts *ContractsSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveAccess(&_Contracts.TransactOpts, _user)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address _user) returns()
func (_Contracts *ContractsTransactorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveAccess(&_Contracts.TransactOpts, _user)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_Contracts *ContractsTransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "requestNewRound")
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_Contracts *ContractsSession) RequestNewRound() (*types.Transaction, error) {
	return _Contracts.Contract.RequestNewRound(&_Contracts.TransactOpts)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_Contracts *ContractsTransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _Contracts.Contract.RequestNewRound(&_Contracts.TransactOpts)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_Contracts *ContractsTransactor) SetRequesterPermissions(opts *bind.TransactOpts, _requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setRequesterPermissions", _requester, _authorized, _delay)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_Contracts *ContractsSession) SetRequesterPermissions(_requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _Contracts.Contract.SetRequesterPermissions(&_Contracts.TransactOpts, _requester, _authorized, _delay)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_Contracts *ContractsTransactorSession) SetRequesterPermissions(_requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _Contracts.Contract.SetRequesterPermissions(&_Contracts.TransactOpts, _requester, _authorized, _delay)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _newValidator) returns()
func (_Contracts *ContractsTransactor) SetValidator(opts *bind.TransactOpts, _newValidator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setValidator", _newValidator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _newValidator) returns()
func (_Contracts *ContractsSession) SetValidator(_newValidator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetValidator(&_Contracts.TransactOpts, _newValidator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _newValidator) returns()
func (_Contracts *ContractsTransactorSession) SetValidator(_newValidator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetValidator(&_Contracts.TransactOpts, _newValidator)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_Contracts *ContractsTransactor) Submit(opts *bind.TransactOpts, _roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "submit", _roundId, _submission)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_Contracts *ContractsSession) Submit(_roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Submit(&_Contracts.TransactOpts, _roundId, _submission)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_Contracts *ContractsTransactorSession) Submit(_roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Submit(&_Contracts.TransactOpts, _roundId, _submission)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_Contracts *ContractsTransactor) TransferAdmin(opts *bind.TransactOpts, _oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferAdmin", _oracle, _newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_Contracts *ContractsSession) TransferAdmin(_oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferAdmin(&_Contracts.TransactOpts, _oracle, _newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_Contracts *ContractsTransactorSession) TransferAdmin(_oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferAdmin(&_Contracts.TransactOpts, _oracle, _newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_Contracts *ContractsSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, _to)
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_Contracts *ContractsTransactor) UpdateAvailableFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateAvailableFunds")
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_Contracts *ContractsSession) UpdateAvailableFunds() (*types.Transaction, error) {
	return _Contracts.Contract.UpdateAvailableFunds(&_Contracts.TransactOpts)
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_Contracts *ContractsTransactorSession) UpdateAvailableFunds() (*types.Transaction, error) {
	return _Contracts.Contract.UpdateAvailableFunds(&_Contracts.TransactOpts)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_Contracts *ContractsTransactor) UpdateFutureRounds(opts *bind.TransactOpts, _paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateFutureRounds", _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_Contracts *ContractsSession) UpdateFutureRounds(_paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateFutureRounds(&_Contracts.TransactOpts, _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_Contracts *ContractsTransactorSession) UpdateFutureRounds(_paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateFutureRounds(&_Contracts.TransactOpts, _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_Contracts *ContractsTransactor) WithdrawFunds(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawFunds", _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_Contracts *ContractsSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawFunds(&_Contracts.TransactOpts, _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawFunds(&_Contracts.TransactOpts, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_Contracts *ContractsTransactor) WithdrawPayment(opts *bind.TransactOpts, _oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawPayment", _oracle, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_Contracts *ContractsSession) WithdrawPayment(_oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawPayment(&_Contracts.TransactOpts, _oracle, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_Contracts *ContractsTransactorSession) WithdrawPayment(_oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawPayment(&_Contracts.TransactOpts, _oracle, _recipient, _amount)
}

// ContractsAddedAccessIterator is returned from FilterAddedAccess and is used to iterate over the raw logs and unpacked data for AddedAccess events raised by the Contracts contract.
type ContractsAddedAccessIterator struct {
	Event *ContractsAddedAccess // Event containing the contract specifics and raw log

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
func (it *ContractsAddedAccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAddedAccess)
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
		it.Event = new(ContractsAddedAccess)
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
func (it *ContractsAddedAccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAddedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAddedAccess represents a AddedAccess event raised by the Contracts contract.
type ContractsAddedAccess struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedAccess is a free log retrieval operation binding the contract event 0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4.
//
// Solidity: event AddedAccess(address user)
func (_Contracts *ContractsFilterer) FilterAddedAccess(opts *bind.FilterOpts) (*ContractsAddedAccessIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return &ContractsAddedAccessIterator{contract: _Contracts.contract, event: "AddedAccess", logs: logs, sub: sub}, nil
}

// WatchAddedAccess is a free log subscription operation binding the contract event 0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4.
//
// Solidity: event AddedAccess(address user)
func (_Contracts *ContractsFilterer) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *ContractsAddedAccess) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAddedAccess)
				if err := _Contracts.contract.UnpackLog(event, "AddedAccess", log); err != nil {
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

// ParseAddedAccess is a log parse operation binding the contract event 0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4.
//
// Solidity: event AddedAccess(address user)
func (_Contracts *ContractsFilterer) ParseAddedAccess(log types.Log) (*ContractsAddedAccess, error) {
	event := new(ContractsAddedAccess)
	if err := _Contracts.contract.UnpackLog(event, "AddedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the Contracts contract.
type ContractsAnswerUpdatedIterator struct {
	Event *ContractsAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAnswerUpdated)
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
		it.Event = new(ContractsAnswerUpdated)
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
func (it *ContractsAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAnswerUpdated represents a AnswerUpdated event raised by the Contracts contract.
type ContractsAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_Contracts *ContractsFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*ContractsAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAnswerUpdatedIterator{contract: _Contracts.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_Contracts *ContractsFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *ContractsAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAnswerUpdated)
				if err := _Contracts.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_Contracts *ContractsFilterer) ParseAnswerUpdated(log types.Log) (*ContractsAnswerUpdated, error) {
	event := new(ContractsAnswerUpdated)
	if err := _Contracts.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAvailableFundsUpdatedIterator is returned from FilterAvailableFundsUpdated and is used to iterate over the raw logs and unpacked data for AvailableFundsUpdated events raised by the Contracts contract.
type ContractsAvailableFundsUpdatedIterator struct {
	Event *ContractsAvailableFundsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsAvailableFundsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAvailableFundsUpdated)
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
		it.Event = new(ContractsAvailableFundsUpdated)
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
func (it *ContractsAvailableFundsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAvailableFundsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAvailableFundsUpdated represents a AvailableFundsUpdated event raised by the Contracts contract.
type ContractsAvailableFundsUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAvailableFundsUpdated is a free log retrieval operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_Contracts *ContractsFilterer) FilterAvailableFundsUpdated(opts *bind.FilterOpts, amount []*big.Int) (*ContractsAvailableFundsUpdatedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AvailableFundsUpdated", amountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAvailableFundsUpdatedIterator{contract: _Contracts.contract, event: "AvailableFundsUpdated", logs: logs, sub: sub}, nil
}

// WatchAvailableFundsUpdated is a free log subscription operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_Contracts *ContractsFilterer) WatchAvailableFundsUpdated(opts *bind.WatchOpts, sink chan<- *ContractsAvailableFundsUpdated, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AvailableFundsUpdated", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAvailableFundsUpdated)
				if err := _Contracts.contract.UnpackLog(event, "AvailableFundsUpdated", log); err != nil {
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

// ParseAvailableFundsUpdated is a log parse operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_Contracts *ContractsFilterer) ParseAvailableFundsUpdated(log types.Log) (*ContractsAvailableFundsUpdated, error) {
	event := new(ContractsAvailableFundsUpdated)
	if err := _Contracts.contract.UnpackLog(event, "AvailableFundsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCheckAccessDisabledIterator is returned from FilterCheckAccessDisabled and is used to iterate over the raw logs and unpacked data for CheckAccessDisabled events raised by the Contracts contract.
type ContractsCheckAccessDisabledIterator struct {
	Event *ContractsCheckAccessDisabled // Event containing the contract specifics and raw log

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
func (it *ContractsCheckAccessDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCheckAccessDisabled)
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
		it.Event = new(ContractsCheckAccessDisabled)
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
func (it *ContractsCheckAccessDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCheckAccessDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCheckAccessDisabled represents a CheckAccessDisabled event raised by the Contracts contract.
type ContractsCheckAccessDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCheckAccessDisabled is a free log retrieval operation binding the contract event 0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638.
//
// Solidity: event CheckAccessDisabled()
func (_Contracts *ContractsFilterer) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*ContractsCheckAccessDisabledIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return &ContractsCheckAccessDisabledIterator{contract: _Contracts.contract, event: "CheckAccessDisabled", logs: logs, sub: sub}, nil
}

// WatchCheckAccessDisabled is a free log subscription operation binding the contract event 0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638.
//
// Solidity: event CheckAccessDisabled()
func (_Contracts *ContractsFilterer) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *ContractsCheckAccessDisabled) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCheckAccessDisabled)
				if err := _Contracts.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
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

// ParseCheckAccessDisabled is a log parse operation binding the contract event 0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638.
//
// Solidity: event CheckAccessDisabled()
func (_Contracts *ContractsFilterer) ParseCheckAccessDisabled(log types.Log) (*ContractsCheckAccessDisabled, error) {
	event := new(ContractsCheckAccessDisabled)
	if err := _Contracts.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCheckAccessEnabledIterator is returned from FilterCheckAccessEnabled and is used to iterate over the raw logs and unpacked data for CheckAccessEnabled events raised by the Contracts contract.
type ContractsCheckAccessEnabledIterator struct {
	Event *ContractsCheckAccessEnabled // Event containing the contract specifics and raw log

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
func (it *ContractsCheckAccessEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCheckAccessEnabled)
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
		it.Event = new(ContractsCheckAccessEnabled)
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
func (it *ContractsCheckAccessEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCheckAccessEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCheckAccessEnabled represents a CheckAccessEnabled event raised by the Contracts contract.
type ContractsCheckAccessEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCheckAccessEnabled is a free log retrieval operation binding the contract event 0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480.
//
// Solidity: event CheckAccessEnabled()
func (_Contracts *ContractsFilterer) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*ContractsCheckAccessEnabledIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return &ContractsCheckAccessEnabledIterator{contract: _Contracts.contract, event: "CheckAccessEnabled", logs: logs, sub: sub}, nil
}

// WatchCheckAccessEnabled is a free log subscription operation binding the contract event 0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480.
//
// Solidity: event CheckAccessEnabled()
func (_Contracts *ContractsFilterer) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *ContractsCheckAccessEnabled) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCheckAccessEnabled)
				if err := _Contracts.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
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

// ParseCheckAccessEnabled is a log parse operation binding the contract event 0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480.
//
// Solidity: event CheckAccessEnabled()
func (_Contracts *ContractsFilterer) ParseCheckAccessEnabled(log types.Log) (*ContractsCheckAccessEnabled, error) {
	event := new(ContractsCheckAccessEnabled)
	if err := _Contracts.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the Contracts contract.
type ContractsNewRoundIterator struct {
	Event *ContractsNewRound // Event containing the contract specifics and raw log

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
func (it *ContractsNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsNewRound)
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
		it.Event = new(ContractsNewRound)
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
func (it *ContractsNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsNewRound represents a NewRound event raised by the Contracts contract.
type ContractsNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_Contracts *ContractsFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*ContractsNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &ContractsNewRoundIterator{contract: _Contracts.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_Contracts *ContractsFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *ContractsNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsNewRound)
				if err := _Contracts.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_Contracts *ContractsFilterer) ParseNewRound(log types.Log) (*ContractsNewRound, error) {
	event := new(ContractsNewRound)
	if err := _Contracts.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOracleAdminUpdateRequestedIterator is returned from FilterOracleAdminUpdateRequested and is used to iterate over the raw logs and unpacked data for OracleAdminUpdateRequested events raised by the Contracts contract.
type ContractsOracleAdminUpdateRequestedIterator struct {
	Event *ContractsOracleAdminUpdateRequested // Event containing the contract specifics and raw log

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
func (it *ContractsOracleAdminUpdateRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOracleAdminUpdateRequested)
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
		it.Event = new(ContractsOracleAdminUpdateRequested)
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
func (it *ContractsOracleAdminUpdateRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOracleAdminUpdateRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOracleAdminUpdateRequested represents a OracleAdminUpdateRequested event raised by the Contracts contract.
type ContractsOracleAdminUpdateRequested struct {
	Oracle   common.Address
	Admin    common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOracleAdminUpdateRequested is a free log retrieval operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_Contracts *ContractsFilterer) FilterOracleAdminUpdateRequested(opts *bind.FilterOpts, oracle []common.Address) (*ContractsOracleAdminUpdateRequestedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OracleAdminUpdateRequested", oracleRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOracleAdminUpdateRequestedIterator{contract: _Contracts.contract, event: "OracleAdminUpdateRequested", logs: logs, sub: sub}, nil
}

// WatchOracleAdminUpdateRequested is a free log subscription operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_Contracts *ContractsFilterer) WatchOracleAdminUpdateRequested(opts *bind.WatchOpts, sink chan<- *ContractsOracleAdminUpdateRequested, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OracleAdminUpdateRequested", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOracleAdminUpdateRequested)
				if err := _Contracts.contract.UnpackLog(event, "OracleAdminUpdateRequested", log); err != nil {
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

// ParseOracleAdminUpdateRequested is a log parse operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_Contracts *ContractsFilterer) ParseOracleAdminUpdateRequested(log types.Log) (*ContractsOracleAdminUpdateRequested, error) {
	event := new(ContractsOracleAdminUpdateRequested)
	if err := _Contracts.contract.UnpackLog(event, "OracleAdminUpdateRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOracleAdminUpdatedIterator is returned from FilterOracleAdminUpdated and is used to iterate over the raw logs and unpacked data for OracleAdminUpdated events raised by the Contracts contract.
type ContractsOracleAdminUpdatedIterator struct {
	Event *ContractsOracleAdminUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsOracleAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOracleAdminUpdated)
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
		it.Event = new(ContractsOracleAdminUpdated)
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
func (it *ContractsOracleAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOracleAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOracleAdminUpdated represents a OracleAdminUpdated event raised by the Contracts contract.
type ContractsOracleAdminUpdated struct {
	Oracle   common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOracleAdminUpdated is a free log retrieval operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_Contracts *ContractsFilterer) FilterOracleAdminUpdated(opts *bind.FilterOpts, oracle []common.Address, newAdmin []common.Address) (*ContractsOracleAdminUpdatedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OracleAdminUpdated", oracleRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOracleAdminUpdatedIterator{contract: _Contracts.contract, event: "OracleAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleAdminUpdated is a free log subscription operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_Contracts *ContractsFilterer) WatchOracleAdminUpdated(opts *bind.WatchOpts, sink chan<- *ContractsOracleAdminUpdated, oracle []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OracleAdminUpdated", oracleRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOracleAdminUpdated)
				if err := _Contracts.contract.UnpackLog(event, "OracleAdminUpdated", log); err != nil {
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

// ParseOracleAdminUpdated is a log parse operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_Contracts *ContractsFilterer) ParseOracleAdminUpdated(log types.Log) (*ContractsOracleAdminUpdated, error) {
	event := new(ContractsOracleAdminUpdated)
	if err := _Contracts.contract.UnpackLog(event, "OracleAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOraclePermissionsUpdatedIterator is returned from FilterOraclePermissionsUpdated and is used to iterate over the raw logs and unpacked data for OraclePermissionsUpdated events raised by the Contracts contract.
type ContractsOraclePermissionsUpdatedIterator struct {
	Event *ContractsOraclePermissionsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsOraclePermissionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOraclePermissionsUpdated)
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
		it.Event = new(ContractsOraclePermissionsUpdated)
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
func (it *ContractsOraclePermissionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOraclePermissionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOraclePermissionsUpdated represents a OraclePermissionsUpdated event raised by the Contracts contract.
type ContractsOraclePermissionsUpdated struct {
	Oracle      common.Address
	Whitelisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOraclePermissionsUpdated is a free log retrieval operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_Contracts *ContractsFilterer) FilterOraclePermissionsUpdated(opts *bind.FilterOpts, oracle []common.Address, whitelisted []bool) (*ContractsOraclePermissionsUpdatedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OraclePermissionsUpdated", oracleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOraclePermissionsUpdatedIterator{contract: _Contracts.contract, event: "OraclePermissionsUpdated", logs: logs, sub: sub}, nil
}

// WatchOraclePermissionsUpdated is a free log subscription operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_Contracts *ContractsFilterer) WatchOraclePermissionsUpdated(opts *bind.WatchOpts, sink chan<- *ContractsOraclePermissionsUpdated, oracle []common.Address, whitelisted []bool) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OraclePermissionsUpdated", oracleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOraclePermissionsUpdated)
				if err := _Contracts.contract.UnpackLog(event, "OraclePermissionsUpdated", log); err != nil {
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

// ParseOraclePermissionsUpdated is a log parse operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_Contracts *ContractsFilterer) ParseOraclePermissionsUpdated(log types.Log) (*ContractsOraclePermissionsUpdated, error) {
	event := new(ContractsOraclePermissionsUpdated)
	if err := _Contracts.contract.UnpackLog(event, "OraclePermissionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the Contracts contract.
type ContractsOwnershipTransferRequestedIterator struct {
	Event *ContractsOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferRequested)
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
		it.Event = new(ContractsOwnershipTransferRequested)
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
func (it *ContractsOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the Contracts contract.
type ContractsOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferRequestedIterator{contract: _Contracts.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferRequested)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Contracts *ContractsFilterer) ParseOwnershipTransferRequested(log types.Log) (*ContractsOwnershipTransferRequested, error) {
	event := new(ContractsOwnershipTransferRequested)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRemovedAccessIterator is returned from FilterRemovedAccess and is used to iterate over the raw logs and unpacked data for RemovedAccess events raised by the Contracts contract.
type ContractsRemovedAccessIterator struct {
	Event *ContractsRemovedAccess // Event containing the contract specifics and raw log

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
func (it *ContractsRemovedAccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRemovedAccess)
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
		it.Event = new(ContractsRemovedAccess)
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
func (it *ContractsRemovedAccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRemovedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRemovedAccess represents a RemovedAccess event raised by the Contracts contract.
type ContractsRemovedAccess struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedAccess is a free log retrieval operation binding the contract event 0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1.
//
// Solidity: event RemovedAccess(address user)
func (_Contracts *ContractsFilterer) FilterRemovedAccess(opts *bind.FilterOpts) (*ContractsRemovedAccessIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return &ContractsRemovedAccessIterator{contract: _Contracts.contract, event: "RemovedAccess", logs: logs, sub: sub}, nil
}

// WatchRemovedAccess is a free log subscription operation binding the contract event 0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1.
//
// Solidity: event RemovedAccess(address user)
func (_Contracts *ContractsFilterer) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *ContractsRemovedAccess) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRemovedAccess)
				if err := _Contracts.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
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

// ParseRemovedAccess is a log parse operation binding the contract event 0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1.
//
// Solidity: event RemovedAccess(address user)
func (_Contracts *ContractsFilterer) ParseRemovedAccess(log types.Log) (*ContractsRemovedAccess, error) {
	event := new(ContractsRemovedAccess)
	if err := _Contracts.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRequesterPermissionsSetIterator is returned from FilterRequesterPermissionsSet and is used to iterate over the raw logs and unpacked data for RequesterPermissionsSet events raised by the Contracts contract.
type ContractsRequesterPermissionsSetIterator struct {
	Event *ContractsRequesterPermissionsSet // Event containing the contract specifics and raw log

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
func (it *ContractsRequesterPermissionsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRequesterPermissionsSet)
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
		it.Event = new(ContractsRequesterPermissionsSet)
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
func (it *ContractsRequesterPermissionsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRequesterPermissionsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRequesterPermissionsSet represents a RequesterPermissionsSet event raised by the Contracts contract.
type ContractsRequesterPermissionsSet struct {
	Requester  common.Address
	Authorized bool
	Delay      uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequesterPermissionsSet is a free log retrieval operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_Contracts *ContractsFilterer) FilterRequesterPermissionsSet(opts *bind.FilterOpts, requester []common.Address) (*ContractsRequesterPermissionsSetIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RequesterPermissionsSet", requesterRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRequesterPermissionsSetIterator{contract: _Contracts.contract, event: "RequesterPermissionsSet", logs: logs, sub: sub}, nil
}

// WatchRequesterPermissionsSet is a free log subscription operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_Contracts *ContractsFilterer) WatchRequesterPermissionsSet(opts *bind.WatchOpts, sink chan<- *ContractsRequesterPermissionsSet, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RequesterPermissionsSet", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRequesterPermissionsSet)
				if err := _Contracts.contract.UnpackLog(event, "RequesterPermissionsSet", log); err != nil {
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

// ParseRequesterPermissionsSet is a log parse operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_Contracts *ContractsFilterer) ParseRequesterPermissionsSet(log types.Log) (*ContractsRequesterPermissionsSet, error) {
	event := new(ContractsRequesterPermissionsSet)
	if err := _Contracts.contract.UnpackLog(event, "RequesterPermissionsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRoundDetailsUpdatedIterator is returned from FilterRoundDetailsUpdated and is used to iterate over the raw logs and unpacked data for RoundDetailsUpdated events raised by the Contracts contract.
type ContractsRoundDetailsUpdatedIterator struct {
	Event *ContractsRoundDetailsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsRoundDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRoundDetailsUpdated)
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
		it.Event = new(ContractsRoundDetailsUpdated)
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
func (it *ContractsRoundDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRoundDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRoundDetailsUpdated represents a RoundDetailsUpdated event raised by the Contracts contract.
type ContractsRoundDetailsUpdated struct {
	PaymentAmount      *big.Int
	MinSubmissionCount uint32
	MaxSubmissionCount uint32
	RestartDelay       uint32
	Timeout            uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRoundDetailsUpdated is a free log retrieval operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_Contracts *ContractsFilterer) FilterRoundDetailsUpdated(opts *bind.FilterOpts, paymentAmount []*big.Int, minSubmissionCount []uint32, maxSubmissionCount []uint32) (*ContractsRoundDetailsUpdatedIterator, error) {

	var paymentAmountRule []interface{}
	for _, paymentAmountItem := range paymentAmount {
		paymentAmountRule = append(paymentAmountRule, paymentAmountItem)
	}
	var minSubmissionCountRule []interface{}
	for _, minSubmissionCountItem := range minSubmissionCount {
		minSubmissionCountRule = append(minSubmissionCountRule, minSubmissionCountItem)
	}
	var maxSubmissionCountRule []interface{}
	for _, maxSubmissionCountItem := range maxSubmissionCount {
		maxSubmissionCountRule = append(maxSubmissionCountRule, maxSubmissionCountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RoundDetailsUpdated", paymentAmountRule, minSubmissionCountRule, maxSubmissionCountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRoundDetailsUpdatedIterator{contract: _Contracts.contract, event: "RoundDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchRoundDetailsUpdated is a free log subscription operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_Contracts *ContractsFilterer) WatchRoundDetailsUpdated(opts *bind.WatchOpts, sink chan<- *ContractsRoundDetailsUpdated, paymentAmount []*big.Int, minSubmissionCount []uint32, maxSubmissionCount []uint32) (event.Subscription, error) {

	var paymentAmountRule []interface{}
	for _, paymentAmountItem := range paymentAmount {
		paymentAmountRule = append(paymentAmountRule, paymentAmountItem)
	}
	var minSubmissionCountRule []interface{}
	for _, minSubmissionCountItem := range minSubmissionCount {
		minSubmissionCountRule = append(minSubmissionCountRule, minSubmissionCountItem)
	}
	var maxSubmissionCountRule []interface{}
	for _, maxSubmissionCountItem := range maxSubmissionCount {
		maxSubmissionCountRule = append(maxSubmissionCountRule, maxSubmissionCountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RoundDetailsUpdated", paymentAmountRule, minSubmissionCountRule, maxSubmissionCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRoundDetailsUpdated)
				if err := _Contracts.contract.UnpackLog(event, "RoundDetailsUpdated", log); err != nil {
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

// ParseRoundDetailsUpdated is a log parse operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_Contracts *ContractsFilterer) ParseRoundDetailsUpdated(log types.Log) (*ContractsRoundDetailsUpdated, error) {
	event := new(ContractsRoundDetailsUpdated)
	if err := _Contracts.contract.UnpackLog(event, "RoundDetailsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSubmissionReceivedIterator is returned from FilterSubmissionReceived and is used to iterate over the raw logs and unpacked data for SubmissionReceived events raised by the Contracts contract.
type ContractsSubmissionReceivedIterator struct {
	Event *ContractsSubmissionReceived // Event containing the contract specifics and raw log

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
func (it *ContractsSubmissionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSubmissionReceived)
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
		it.Event = new(ContractsSubmissionReceived)
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
func (it *ContractsSubmissionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSubmissionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSubmissionReceived represents a SubmissionReceived event raised by the Contracts contract.
type ContractsSubmissionReceived struct {
	Submission *big.Int
	Round      uint32
	Oracle     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubmissionReceived is a free log retrieval operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_Contracts *ContractsFilterer) FilterSubmissionReceived(opts *bind.FilterOpts, submission []*big.Int, round []uint32, oracle []common.Address) (*ContractsSubmissionReceivedIterator, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SubmissionReceived", submissionRule, roundRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSubmissionReceivedIterator{contract: _Contracts.contract, event: "SubmissionReceived", logs: logs, sub: sub}, nil
}

// WatchSubmissionReceived is a free log subscription operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_Contracts *ContractsFilterer) WatchSubmissionReceived(opts *bind.WatchOpts, sink chan<- *ContractsSubmissionReceived, submission []*big.Int, round []uint32, oracle []common.Address) (event.Subscription, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SubmissionReceived", submissionRule, roundRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSubmissionReceived)
				if err := _Contracts.contract.UnpackLog(event, "SubmissionReceived", log); err != nil {
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

// ParseSubmissionReceived is a log parse operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_Contracts *ContractsFilterer) ParseSubmissionReceived(log types.Log) (*ContractsSubmissionReceived, error) {
	event := new(ContractsSubmissionReceived)
	if err := _Contracts.contract.UnpackLog(event, "SubmissionReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsValidatorUpdatedIterator is returned from FilterValidatorUpdated and is used to iterate over the raw logs and unpacked data for ValidatorUpdated events raised by the Contracts contract.
type ContractsValidatorUpdatedIterator struct {
	Event *ContractsValidatorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsValidatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsValidatorUpdated)
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
		it.Event = new(ContractsValidatorUpdated)
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
func (it *ContractsValidatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsValidatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsValidatorUpdated represents a ValidatorUpdated event raised by the Contracts contract.
type ContractsValidatorUpdated struct {
	Previous common.Address
	Current  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValidatorUpdated is a free log retrieval operation binding the contract event 0xcfac5dc75b8d9a7e074162f59d9adcd33da59f0fe8dfb21580db298fc0fdad0d.
//
// Solidity: event ValidatorUpdated(address indexed previous, address indexed current)
func (_Contracts *ContractsFilterer) FilterValidatorUpdated(opts *bind.FilterOpts, previous []common.Address, current []common.Address) (*ContractsValidatorUpdatedIterator, error) {

	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ValidatorUpdated", previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &ContractsValidatorUpdatedIterator{contract: _Contracts.contract, event: "ValidatorUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorUpdated is a free log subscription operation binding the contract event 0xcfac5dc75b8d9a7e074162f59d9adcd33da59f0fe8dfb21580db298fc0fdad0d.
//
// Solidity: event ValidatorUpdated(address indexed previous, address indexed current)
func (_Contracts *ContractsFilterer) WatchValidatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractsValidatorUpdated, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ValidatorUpdated", previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsValidatorUpdated)
				if err := _Contracts.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
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

// ParseValidatorUpdated is a log parse operation binding the contract event 0xcfac5dc75b8d9a7e074162f59d9adcd33da59f0fe8dfb21580db298fc0fdad0d.
//
// Solidity: event ValidatorUpdated(address indexed previous, address indexed current)
func (_Contracts *ContractsFilterer) ParseValidatorUpdated(log types.Log) (*ContractsValidatorUpdated, error) {
	event := new(ContractsValidatorUpdated)
	if err := _Contracts.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
