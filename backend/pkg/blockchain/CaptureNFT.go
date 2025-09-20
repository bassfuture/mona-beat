// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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
	_ = abi.ConvertType
)

// CaptureNFTMetaData contains all meta data concerning the CaptureNFT contract.
var CaptureNFTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"attemptCapture\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"captureId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"difficulty\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rarity\",\"type\":\"uint8\",\"internalType\":\"enumCaptureNFT.Rarity\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorizeMinter\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorizedMinters\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"captureEvents\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"player\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"captureId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"rarity\",\"type\":\"uint8\",\"internalType\":\"enumCaptureNFT.Rarity\"},{\"name\":\"difficulty\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"difficultySuccessRate\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentTokenId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNFTDetails\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rarity\",\"type\":\"uint8\",\"internalType\":\"enumCaptureNFT.Rarity\"},{\"name\":\"difficulty\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"captureId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlayerNFTs\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlayerStats\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalCaptures\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"successfulCaptures\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nftCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"successRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRarityStats\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nftMetadata\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"rarity\",\"type\":\"uint8\",\"internalType\":\"enumCaptureNFT.Rarity\"},{\"name\":\"difficulty\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"captureId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributes\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"playerSuccessfulCaptures\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"playerTotalCaptures\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rarityCount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumCaptureNFT.Rarity\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeMinter\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDifficultySuccessRate\",\"inputs\":[{\"name\":\"difficulty\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"successRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenByIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenOfOwnerByIndex\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"name\":\"_fromTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CaptureAttempted\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"captureId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"rarity\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumCaptureNFT.Rarity\"},{\"name\":\"difficulty\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdate\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinterAuthorized\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinterRevoked\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTMinted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rarity\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumCaptureNFT.Rarity\"},{\"name\":\"captureId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721OutOfBoundsIndex\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// CaptureNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use CaptureNFTMetaData.ABI instead.
var CaptureNFTABI = CaptureNFTMetaData.ABI

// CaptureNFT is an auto generated Go binding around an Ethereum contract.
type CaptureNFT struct {
	CaptureNFTCaller     // Read-only binding to the contract
	CaptureNFTTransactor // Write-only binding to the contract
	CaptureNFTFilterer   // Log filterer for contract events
}

// CaptureNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type CaptureNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CaptureNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CaptureNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CaptureNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CaptureNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CaptureNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CaptureNFTSession struct {
	Contract     *CaptureNFT       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CaptureNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CaptureNFTCallerSession struct {
	Contract *CaptureNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CaptureNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CaptureNFTTransactorSession struct {
	Contract     *CaptureNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CaptureNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type CaptureNFTRaw struct {
	Contract *CaptureNFT // Generic contract binding to access the raw methods on
}

// CaptureNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CaptureNFTCallerRaw struct {
	Contract *CaptureNFTCaller // Generic read-only contract binding to access the raw methods on
}

// CaptureNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CaptureNFTTransactorRaw struct {
	Contract *CaptureNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCaptureNFT creates a new instance of CaptureNFT, bound to a specific deployed contract.
func NewCaptureNFT(address common.Address, backend bind.ContractBackend) (*CaptureNFT, error) {
	contract, err := bindCaptureNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CaptureNFT{CaptureNFTCaller: CaptureNFTCaller{contract: contract}, CaptureNFTTransactor: CaptureNFTTransactor{contract: contract}, CaptureNFTFilterer: CaptureNFTFilterer{contract: contract}}, nil
}

// NewCaptureNFTCaller creates a new read-only instance of CaptureNFT, bound to a specific deployed contract.
func NewCaptureNFTCaller(address common.Address, caller bind.ContractCaller) (*CaptureNFTCaller, error) {
	contract, err := bindCaptureNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTCaller{contract: contract}, nil
}

// NewCaptureNFTTransactor creates a new write-only instance of CaptureNFT, bound to a specific deployed contract.
func NewCaptureNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*CaptureNFTTransactor, error) {
	contract, err := bindCaptureNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTTransactor{contract: contract}, nil
}

// NewCaptureNFTFilterer creates a new log filterer instance of CaptureNFT, bound to a specific deployed contract.
func NewCaptureNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*CaptureNFTFilterer, error) {
	contract, err := bindCaptureNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTFilterer{contract: contract}, nil
}

// bindCaptureNFT binds a generic wrapper to an already deployed contract.
func bindCaptureNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CaptureNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CaptureNFT *CaptureNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CaptureNFT.Contract.CaptureNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CaptureNFT *CaptureNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CaptureNFT.Contract.CaptureNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CaptureNFT *CaptureNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CaptureNFT.Contract.CaptureNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CaptureNFT *CaptureNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CaptureNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CaptureNFT *CaptureNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CaptureNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CaptureNFT *CaptureNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CaptureNFT.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedMinters is a free data retrieval call binding the contract method 0xaa2fe91b.
//
// Solidity: function authorizedMinters(address ) view returns(bool)
func (_CaptureNFT *CaptureNFTCaller) AuthorizedMinters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "authorizedMinters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedMinters is a free data retrieval call binding the contract method 0xaa2fe91b.
//
// Solidity: function authorizedMinters(address ) view returns(bool)
func (_CaptureNFT *CaptureNFTSession) AuthorizedMinters(arg0 common.Address) (bool, error) {
	return _CaptureNFT.Contract.AuthorizedMinters(&_CaptureNFT.CallOpts, arg0)
}

// AuthorizedMinters is a free data retrieval call binding the contract method 0xaa2fe91b.
//
// Solidity: function authorizedMinters(address ) view returns(bool)
func (_CaptureNFT *CaptureNFTCallerSession) AuthorizedMinters(arg0 common.Address) (bool, error) {
	return _CaptureNFT.Contract.AuthorizedMinters(&_CaptureNFT.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CaptureNFT *CaptureNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CaptureNFT *CaptureNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CaptureNFT.Contract.BalanceOf(&_CaptureNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CaptureNFT *CaptureNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CaptureNFT.Contract.BalanceOf(&_CaptureNFT.CallOpts, owner)
}

// CaptureEvents is a free data retrieval call binding the contract method 0x15dd7c19.
//
// Solidity: function captureEvents(string ) view returns(address player, string captureId, bool success, uint8 rarity, uint256 difficulty, uint256 timestamp, uint256 tokenId)
func (_CaptureNFT *CaptureNFTCaller) CaptureEvents(opts *bind.CallOpts, arg0 string) (struct {
	Player     common.Address
	CaptureId  string
	Success    bool
	Rarity     uint8
	Difficulty *big.Int
	Timestamp  *big.Int
	TokenId    *big.Int
}, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "captureEvents", arg0)

	outstruct := new(struct {
		Player     common.Address
		CaptureId  string
		Success    bool
		Rarity     uint8
		Difficulty *big.Int
		Timestamp  *big.Int
		TokenId    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Player = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CaptureId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Success = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Rarity = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Difficulty = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CaptureEvents is a free data retrieval call binding the contract method 0x15dd7c19.
//
// Solidity: function captureEvents(string ) view returns(address player, string captureId, bool success, uint8 rarity, uint256 difficulty, uint256 timestamp, uint256 tokenId)
func (_CaptureNFT *CaptureNFTSession) CaptureEvents(arg0 string) (struct {
	Player     common.Address
	CaptureId  string
	Success    bool
	Rarity     uint8
	Difficulty *big.Int
	Timestamp  *big.Int
	TokenId    *big.Int
}, error) {
	return _CaptureNFT.Contract.CaptureEvents(&_CaptureNFT.CallOpts, arg0)
}

// CaptureEvents is a free data retrieval call binding the contract method 0x15dd7c19.
//
// Solidity: function captureEvents(string ) view returns(address player, string captureId, bool success, uint8 rarity, uint256 difficulty, uint256 timestamp, uint256 tokenId)
func (_CaptureNFT *CaptureNFTCallerSession) CaptureEvents(arg0 string) (struct {
	Player     common.Address
	CaptureId  string
	Success    bool
	Rarity     uint8
	Difficulty *big.Int
	Timestamp  *big.Int
	TokenId    *big.Int
}, error) {
	return _CaptureNFT.Contract.CaptureEvents(&_CaptureNFT.CallOpts, arg0)
}

// DifficultySuccessRate is a free data retrieval call binding the contract method 0xe5ea0d60.
//
// Solidity: function difficultySuccessRate(uint256 ) view returns(uint256)
func (_CaptureNFT *CaptureNFTCaller) DifficultySuccessRate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "difficultySuccessRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DifficultySuccessRate is a free data retrieval call binding the contract method 0xe5ea0d60.
//
// Solidity: function difficultySuccessRate(uint256 ) view returns(uint256)
func (_CaptureNFT *CaptureNFTSession) DifficultySuccessRate(arg0 *big.Int) (*big.Int, error) {
	return _CaptureNFT.Contract.DifficultySuccessRate(&_CaptureNFT.CallOpts, arg0)
}

// DifficultySuccessRate is a free data retrieval call binding the contract method 0xe5ea0d60.
//
// Solidity: function difficultySuccessRate(uint256 ) view returns(uint256)
func (_CaptureNFT *CaptureNFTCallerSession) DifficultySuccessRate(arg0 *big.Int) (*big.Int, error) {
	return _CaptureNFT.Contract.DifficultySuccessRate(&_CaptureNFT.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CaptureNFT *CaptureNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CaptureNFT *CaptureNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CaptureNFT.Contract.GetApproved(&_CaptureNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CaptureNFT *CaptureNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CaptureNFT.Contract.GetApproved(&_CaptureNFT.CallOpts, tokenId)
}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_CaptureNFT *CaptureNFTCaller) GetCurrentTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "getCurrentTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_CaptureNFT *CaptureNFTSession) GetCurrentTokenId() (*big.Int, error) {
	return _CaptureNFT.Contract.GetCurrentTokenId(&_CaptureNFT.CallOpts)
}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_CaptureNFT *CaptureNFTCallerSession) GetCurrentTokenId() (*big.Int, error) {
	return _CaptureNFT.Contract.GetCurrentTokenId(&_CaptureNFT.CallOpts)
}

// GetNFTDetails is a free data retrieval call binding the contract method 0x82260f35.
//
// Solidity: function getNFTDetails(uint256 tokenId) view returns(address owner, uint8 rarity, uint256 difficulty, string captureId, uint256 timestamp, string tokenURI)
func (_CaptureNFT *CaptureNFTCaller) GetNFTDetails(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Owner      common.Address
	Rarity     uint8
	Difficulty *big.Int
	CaptureId  string
	Timestamp  *big.Int
	TokenURI   string
}, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "getNFTDetails", tokenId)

	outstruct := new(struct {
		Owner      common.Address
		Rarity     uint8
		Difficulty *big.Int
		CaptureId  string
		Timestamp  *big.Int
		TokenURI   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Rarity = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Difficulty = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CaptureId = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TokenURI = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// GetNFTDetails is a free data retrieval call binding the contract method 0x82260f35.
//
// Solidity: function getNFTDetails(uint256 tokenId) view returns(address owner, uint8 rarity, uint256 difficulty, string captureId, uint256 timestamp, string tokenURI)
func (_CaptureNFT *CaptureNFTSession) GetNFTDetails(tokenId *big.Int) (struct {
	Owner      common.Address
	Rarity     uint8
	Difficulty *big.Int
	CaptureId  string
	Timestamp  *big.Int
	TokenURI   string
}, error) {
	return _CaptureNFT.Contract.GetNFTDetails(&_CaptureNFT.CallOpts, tokenId)
}

// GetNFTDetails is a free data retrieval call binding the contract method 0x82260f35.
//
// Solidity: function getNFTDetails(uint256 tokenId) view returns(address owner, uint8 rarity, uint256 difficulty, string captureId, uint256 timestamp, string tokenURI)
func (_CaptureNFT *CaptureNFTCallerSession) GetNFTDetails(tokenId *big.Int) (struct {
	Owner      common.Address
	Rarity     uint8
	Difficulty *big.Int
	CaptureId  string
	Timestamp  *big.Int
	TokenURI   string
}, error) {
	return _CaptureNFT.Contract.GetNFTDetails(&_CaptureNFT.CallOpts, tokenId)
}

// GetPlayerNFTs is a free data retrieval call binding the contract method 0x00570ab0.
//
// Solidity: function getPlayerNFTs(address player) view returns(uint256[])
func (_CaptureNFT *CaptureNFTCaller) GetPlayerNFTs(opts *bind.CallOpts, player common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "getPlayerNFTs", player)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPlayerNFTs is a free data retrieval call binding the contract method 0x00570ab0.
//
// Solidity: function getPlayerNFTs(address player) view returns(uint256[])
func (_CaptureNFT *CaptureNFTSession) GetPlayerNFTs(player common.Address) ([]*big.Int, error) {
	return _CaptureNFT.Contract.GetPlayerNFTs(&_CaptureNFT.CallOpts, player)
}

// GetPlayerNFTs is a free data retrieval call binding the contract method 0x00570ab0.
//
// Solidity: function getPlayerNFTs(address player) view returns(uint256[])
func (_CaptureNFT *CaptureNFTCallerSession) GetPlayerNFTs(player common.Address) ([]*big.Int, error) {
	return _CaptureNFT.Contract.GetPlayerNFTs(&_CaptureNFT.CallOpts, player)
}

// GetPlayerStats is a free data retrieval call binding the contract method 0x4fd66eae.
//
// Solidity: function getPlayerStats(address player) view returns(uint256 totalCaptures, uint256 successfulCaptures, uint256 nftCount, uint256 successRate)
func (_CaptureNFT *CaptureNFTCaller) GetPlayerStats(opts *bind.CallOpts, player common.Address) (struct {
	TotalCaptures      *big.Int
	SuccessfulCaptures *big.Int
	NftCount           *big.Int
	SuccessRate        *big.Int
}, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "getPlayerStats", player)

	outstruct := new(struct {
		TotalCaptures      *big.Int
		SuccessfulCaptures *big.Int
		NftCount           *big.Int
		SuccessRate        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCaptures = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SuccessfulCaptures = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NftCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SuccessRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPlayerStats is a free data retrieval call binding the contract method 0x4fd66eae.
//
// Solidity: function getPlayerStats(address player) view returns(uint256 totalCaptures, uint256 successfulCaptures, uint256 nftCount, uint256 successRate)
func (_CaptureNFT *CaptureNFTSession) GetPlayerStats(player common.Address) (struct {
	TotalCaptures      *big.Int
	SuccessfulCaptures *big.Int
	NftCount           *big.Int
	SuccessRate        *big.Int
}, error) {
	return _CaptureNFT.Contract.GetPlayerStats(&_CaptureNFT.CallOpts, player)
}

// GetPlayerStats is a free data retrieval call binding the contract method 0x4fd66eae.
//
// Solidity: function getPlayerStats(address player) view returns(uint256 totalCaptures, uint256 successfulCaptures, uint256 nftCount, uint256 successRate)
func (_CaptureNFT *CaptureNFTCallerSession) GetPlayerStats(player common.Address) (struct {
	TotalCaptures      *big.Int
	SuccessfulCaptures *big.Int
	NftCount           *big.Int
	SuccessRate        *big.Int
}, error) {
	return _CaptureNFT.Contract.GetPlayerStats(&_CaptureNFT.CallOpts, player)
}

// GetRarityStats is a free data retrieval call binding the contract method 0x8b3c88b4.
//
// Solidity: function getRarityStats() view returns(uint256[5])
func (_CaptureNFT *CaptureNFTCaller) GetRarityStats(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "getRarityStats")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetRarityStats is a free data retrieval call binding the contract method 0x8b3c88b4.
//
// Solidity: function getRarityStats() view returns(uint256[5])
func (_CaptureNFT *CaptureNFTSession) GetRarityStats() ([5]*big.Int, error) {
	return _CaptureNFT.Contract.GetRarityStats(&_CaptureNFT.CallOpts)
}

// GetRarityStats is a free data retrieval call binding the contract method 0x8b3c88b4.
//
// Solidity: function getRarityStats() view returns(uint256[5])
func (_CaptureNFT *CaptureNFTCallerSession) GetRarityStats() ([5]*big.Int, error) {
	return _CaptureNFT.Contract.GetRarityStats(&_CaptureNFT.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CaptureNFT *CaptureNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CaptureNFT *CaptureNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CaptureNFT.Contract.IsApprovedForAll(&_CaptureNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CaptureNFT *CaptureNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CaptureNFT.Contract.IsApprovedForAll(&_CaptureNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CaptureNFT *CaptureNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CaptureNFT *CaptureNFTSession) Name() (string, error) {
	return _CaptureNFT.Contract.Name(&_CaptureNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CaptureNFT *CaptureNFTCallerSession) Name() (string, error) {
	return _CaptureNFT.Contract.Name(&_CaptureNFT.CallOpts)
}

// NftMetadata is a free data retrieval call binding the contract method 0x49541015.
//
// Solidity: function nftMetadata(uint256 ) view returns(uint8 rarity, uint256 difficulty, string captureId, uint256 timestamp, string attributes)
func (_CaptureNFT *CaptureNFTCaller) NftMetadata(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Rarity     uint8
	Difficulty *big.Int
	CaptureId  string
	Timestamp  *big.Int
	Attributes string
}, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "nftMetadata", arg0)

	outstruct := new(struct {
		Rarity     uint8
		Difficulty *big.Int
		CaptureId  string
		Timestamp  *big.Int
		Attributes string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rarity = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Difficulty = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CaptureId = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Attributes = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// NftMetadata is a free data retrieval call binding the contract method 0x49541015.
//
// Solidity: function nftMetadata(uint256 ) view returns(uint8 rarity, uint256 difficulty, string captureId, uint256 timestamp, string attributes)
func (_CaptureNFT *CaptureNFTSession) NftMetadata(arg0 *big.Int) (struct {
	Rarity     uint8
	Difficulty *big.Int
	CaptureId  string
	Timestamp  *big.Int
	Attributes string
}, error) {
	return _CaptureNFT.Contract.NftMetadata(&_CaptureNFT.CallOpts, arg0)
}

// NftMetadata is a free data retrieval call binding the contract method 0x49541015.
//
// Solidity: function nftMetadata(uint256 ) view returns(uint8 rarity, uint256 difficulty, string captureId, uint256 timestamp, string attributes)
func (_CaptureNFT *CaptureNFTCallerSession) NftMetadata(arg0 *big.Int) (struct {
	Rarity     uint8
	Difficulty *big.Int
	CaptureId  string
	Timestamp  *big.Int
	Attributes string
}, error) {
	return _CaptureNFT.Contract.NftMetadata(&_CaptureNFT.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CaptureNFT *CaptureNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CaptureNFT *CaptureNFTSession) Owner() (common.Address, error) {
	return _CaptureNFT.Contract.Owner(&_CaptureNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CaptureNFT *CaptureNFTCallerSession) Owner() (common.Address, error) {
	return _CaptureNFT.Contract.Owner(&_CaptureNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CaptureNFT *CaptureNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CaptureNFT *CaptureNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CaptureNFT.Contract.OwnerOf(&_CaptureNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CaptureNFT *CaptureNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CaptureNFT.Contract.OwnerOf(&_CaptureNFT.CallOpts, tokenId)
}

// PlayerSuccessfulCaptures is a free data retrieval call binding the contract method 0xc4ac7025.
//
// Solidity: function playerSuccessfulCaptures(address ) view returns(uint256)
func (_CaptureNFT *CaptureNFTCaller) PlayerSuccessfulCaptures(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "playerSuccessfulCaptures", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlayerSuccessfulCaptures is a free data retrieval call binding the contract method 0xc4ac7025.
//
// Solidity: function playerSuccessfulCaptures(address ) view returns(uint256)
func (_CaptureNFT *CaptureNFTSession) PlayerSuccessfulCaptures(arg0 common.Address) (*big.Int, error) {
	return _CaptureNFT.Contract.PlayerSuccessfulCaptures(&_CaptureNFT.CallOpts, arg0)
}

// PlayerSuccessfulCaptures is a free data retrieval call binding the contract method 0xc4ac7025.
//
// Solidity: function playerSuccessfulCaptures(address ) view returns(uint256)
func (_CaptureNFT *CaptureNFTCallerSession) PlayerSuccessfulCaptures(arg0 common.Address) (*big.Int, error) {
	return _CaptureNFT.Contract.PlayerSuccessfulCaptures(&_CaptureNFT.CallOpts, arg0)
}

// PlayerTotalCaptures is a free data retrieval call binding the contract method 0xdd510823.
//
// Solidity: function playerTotalCaptures(address ) view returns(uint256)
func (_CaptureNFT *CaptureNFTCaller) PlayerTotalCaptures(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "playerTotalCaptures", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlayerTotalCaptures is a free data retrieval call binding the contract method 0xdd510823.
//
// Solidity: function playerTotalCaptures(address ) view returns(uint256)
func (_CaptureNFT *CaptureNFTSession) PlayerTotalCaptures(arg0 common.Address) (*big.Int, error) {
	return _CaptureNFT.Contract.PlayerTotalCaptures(&_CaptureNFT.CallOpts, arg0)
}

// PlayerTotalCaptures is a free data retrieval call binding the contract method 0xdd510823.
//
// Solidity: function playerTotalCaptures(address ) view returns(uint256)
func (_CaptureNFT *CaptureNFTCallerSession) PlayerTotalCaptures(arg0 common.Address) (*big.Int, error) {
	return _CaptureNFT.Contract.PlayerTotalCaptures(&_CaptureNFT.CallOpts, arg0)
}

// RarityCount is a free data retrieval call binding the contract method 0x22da8b9e.
//
// Solidity: function rarityCount(uint8 ) view returns(uint256)
func (_CaptureNFT *CaptureNFTCaller) RarityCount(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "rarityCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RarityCount is a free data retrieval call binding the contract method 0x22da8b9e.
//
// Solidity: function rarityCount(uint8 ) view returns(uint256)
func (_CaptureNFT *CaptureNFTSession) RarityCount(arg0 uint8) (*big.Int, error) {
	return _CaptureNFT.Contract.RarityCount(&_CaptureNFT.CallOpts, arg0)
}

// RarityCount is a free data retrieval call binding the contract method 0x22da8b9e.
//
// Solidity: function rarityCount(uint8 ) view returns(uint256)
func (_CaptureNFT *CaptureNFTCallerSession) RarityCount(arg0 uint8) (*big.Int, error) {
	return _CaptureNFT.Contract.RarityCount(&_CaptureNFT.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CaptureNFT *CaptureNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CaptureNFT *CaptureNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CaptureNFT.Contract.SupportsInterface(&_CaptureNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CaptureNFT *CaptureNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CaptureNFT.Contract.SupportsInterface(&_CaptureNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CaptureNFT *CaptureNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CaptureNFT *CaptureNFTSession) Symbol() (string, error) {
	return _CaptureNFT.Contract.Symbol(&_CaptureNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CaptureNFT *CaptureNFTCallerSession) Symbol() (string, error) {
	return _CaptureNFT.Contract.Symbol(&_CaptureNFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_CaptureNFT *CaptureNFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_CaptureNFT *CaptureNFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _CaptureNFT.Contract.TokenByIndex(&_CaptureNFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_CaptureNFT *CaptureNFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _CaptureNFT.Contract.TokenByIndex(&_CaptureNFT.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_CaptureNFT *CaptureNFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_CaptureNFT *CaptureNFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _CaptureNFT.Contract.TokenOfOwnerByIndex(&_CaptureNFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_CaptureNFT *CaptureNFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _CaptureNFT.Contract.TokenOfOwnerByIndex(&_CaptureNFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CaptureNFT *CaptureNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CaptureNFT *CaptureNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CaptureNFT.Contract.TokenURI(&_CaptureNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CaptureNFT *CaptureNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CaptureNFT.Contract.TokenURI(&_CaptureNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CaptureNFT *CaptureNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CaptureNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CaptureNFT *CaptureNFTSession) TotalSupply() (*big.Int, error) {
	return _CaptureNFT.Contract.TotalSupply(&_CaptureNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CaptureNFT *CaptureNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _CaptureNFT.Contract.TotalSupply(&_CaptureNFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CaptureNFT *CaptureNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CaptureNFT *CaptureNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.Contract.Approve(&_CaptureNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CaptureNFT *CaptureNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.Contract.Approve(&_CaptureNFT.TransactOpts, to, tokenId)
}

// AttemptCapture is a paid mutator transaction binding the contract method 0x79243d71.
//
// Solidity: function attemptCapture(address player, string captureId, uint256 difficulty, uint8 rarity, bool success, string tokenURI) returns(uint256)
func (_CaptureNFT *CaptureNFTTransactor) AttemptCapture(opts *bind.TransactOpts, player common.Address, captureId string, difficulty *big.Int, rarity uint8, success bool, tokenURI string) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "attemptCapture", player, captureId, difficulty, rarity, success, tokenURI)
}

// AttemptCapture is a paid mutator transaction binding the contract method 0x79243d71.
//
// Solidity: function attemptCapture(address player, string captureId, uint256 difficulty, uint8 rarity, bool success, string tokenURI) returns(uint256)
func (_CaptureNFT *CaptureNFTSession) AttemptCapture(player common.Address, captureId string, difficulty *big.Int, rarity uint8, success bool, tokenURI string) (*types.Transaction, error) {
	return _CaptureNFT.Contract.AttemptCapture(&_CaptureNFT.TransactOpts, player, captureId, difficulty, rarity, success, tokenURI)
}

// AttemptCapture is a paid mutator transaction binding the contract method 0x79243d71.
//
// Solidity: function attemptCapture(address player, string captureId, uint256 difficulty, uint8 rarity, bool success, string tokenURI) returns(uint256)
func (_CaptureNFT *CaptureNFTTransactorSession) AttemptCapture(player common.Address, captureId string, difficulty *big.Int, rarity uint8, success bool, tokenURI string) (*types.Transaction, error) {
	return _CaptureNFT.Contract.AttemptCapture(&_CaptureNFT.TransactOpts, player, captureId, difficulty, rarity, success, tokenURI)
}

// AuthorizeMinter is a paid mutator transaction binding the contract method 0x0c984832.
//
// Solidity: function authorizeMinter(address minter) returns()
func (_CaptureNFT *CaptureNFTTransactor) AuthorizeMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "authorizeMinter", minter)
}

// AuthorizeMinter is a paid mutator transaction binding the contract method 0x0c984832.
//
// Solidity: function authorizeMinter(address minter) returns()
func (_CaptureNFT *CaptureNFTSession) AuthorizeMinter(minter common.Address) (*types.Transaction, error) {
	return _CaptureNFT.Contract.AuthorizeMinter(&_CaptureNFT.TransactOpts, minter)
}

// AuthorizeMinter is a paid mutator transaction binding the contract method 0x0c984832.
//
// Solidity: function authorizeMinter(address minter) returns()
func (_CaptureNFT *CaptureNFTTransactorSession) AuthorizeMinter(minter common.Address) (*types.Transaction, error) {
	return _CaptureNFT.Contract.AuthorizeMinter(&_CaptureNFT.TransactOpts, minter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CaptureNFT *CaptureNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CaptureNFT *CaptureNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _CaptureNFT.Contract.RenounceOwnership(&_CaptureNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CaptureNFT *CaptureNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CaptureNFT.Contract.RenounceOwnership(&_CaptureNFT.TransactOpts)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address minter) returns()
func (_CaptureNFT *CaptureNFTTransactor) RevokeMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "revokeMinter", minter)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address minter) returns()
func (_CaptureNFT *CaptureNFTSession) RevokeMinter(minter common.Address) (*types.Transaction, error) {
	return _CaptureNFT.Contract.RevokeMinter(&_CaptureNFT.TransactOpts, minter)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address minter) returns()
func (_CaptureNFT *CaptureNFTTransactorSession) RevokeMinter(minter common.Address) (*types.Transaction, error) {
	return _CaptureNFT.Contract.RevokeMinter(&_CaptureNFT.TransactOpts, minter)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CaptureNFT *CaptureNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CaptureNFT *CaptureNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.Contract.SafeTransferFrom(&_CaptureNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CaptureNFT *CaptureNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.Contract.SafeTransferFrom(&_CaptureNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CaptureNFT *CaptureNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CaptureNFT *CaptureNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CaptureNFT.Contract.SafeTransferFrom0(&_CaptureNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CaptureNFT *CaptureNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CaptureNFT.Contract.SafeTransferFrom0(&_CaptureNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CaptureNFT *CaptureNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CaptureNFT *CaptureNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CaptureNFT.Contract.SetApprovalForAll(&_CaptureNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CaptureNFT *CaptureNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CaptureNFT.Contract.SetApprovalForAll(&_CaptureNFT.TransactOpts, operator, approved)
}

// SetDifficultySuccessRate is a paid mutator transaction binding the contract method 0x4d88bd2e.
//
// Solidity: function setDifficultySuccessRate(uint256 difficulty, uint256 successRate) returns()
func (_CaptureNFT *CaptureNFTTransactor) SetDifficultySuccessRate(opts *bind.TransactOpts, difficulty *big.Int, successRate *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "setDifficultySuccessRate", difficulty, successRate)
}

// SetDifficultySuccessRate is a paid mutator transaction binding the contract method 0x4d88bd2e.
//
// Solidity: function setDifficultySuccessRate(uint256 difficulty, uint256 successRate) returns()
func (_CaptureNFT *CaptureNFTSession) SetDifficultySuccessRate(difficulty *big.Int, successRate *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.Contract.SetDifficultySuccessRate(&_CaptureNFT.TransactOpts, difficulty, successRate)
}

// SetDifficultySuccessRate is a paid mutator transaction binding the contract method 0x4d88bd2e.
//
// Solidity: function setDifficultySuccessRate(uint256 difficulty, uint256 successRate) returns()
func (_CaptureNFT *CaptureNFTTransactorSession) SetDifficultySuccessRate(difficulty *big.Int, successRate *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.Contract.SetDifficultySuccessRate(&_CaptureNFT.TransactOpts, difficulty, successRate)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CaptureNFT *CaptureNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CaptureNFT *CaptureNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.Contract.TransferFrom(&_CaptureNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CaptureNFT *CaptureNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CaptureNFT.Contract.TransferFrom(&_CaptureNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CaptureNFT *CaptureNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CaptureNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CaptureNFT *CaptureNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CaptureNFT.Contract.TransferOwnership(&_CaptureNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CaptureNFT *CaptureNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CaptureNFT.Contract.TransferOwnership(&_CaptureNFT.TransactOpts, newOwner)
}

// CaptureNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CaptureNFT contract.
type CaptureNFTApprovalIterator struct {
	Event *CaptureNFTApproval // Event containing the contract specifics and raw log

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
func (it *CaptureNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaptureNFTApproval)
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
		it.Event = new(CaptureNFTApproval)
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
func (it *CaptureNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaptureNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaptureNFTApproval represents a Approval event raised by the CaptureNFT contract.
type CaptureNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CaptureNFT *CaptureNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CaptureNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CaptureNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTApprovalIterator{contract: _CaptureNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CaptureNFT *CaptureNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CaptureNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CaptureNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaptureNFTApproval)
				if err := _CaptureNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CaptureNFT *CaptureNFTFilterer) ParseApproval(log types.Log) (*CaptureNFTApproval, error) {
	event := new(CaptureNFTApproval)
	if err := _CaptureNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaptureNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CaptureNFT contract.
type CaptureNFTApprovalForAllIterator struct {
	Event *CaptureNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CaptureNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaptureNFTApprovalForAll)
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
		it.Event = new(CaptureNFTApprovalForAll)
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
func (it *CaptureNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaptureNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaptureNFTApprovalForAll represents a ApprovalForAll event raised by the CaptureNFT contract.
type CaptureNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CaptureNFT *CaptureNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CaptureNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CaptureNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTApprovalForAllIterator{contract: _CaptureNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CaptureNFT *CaptureNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CaptureNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CaptureNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaptureNFTApprovalForAll)
				if err := _CaptureNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CaptureNFT *CaptureNFTFilterer) ParseApprovalForAll(log types.Log) (*CaptureNFTApprovalForAll, error) {
	event := new(CaptureNFTApprovalForAll)
	if err := _CaptureNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaptureNFTBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the CaptureNFT contract.
type CaptureNFTBatchMetadataUpdateIterator struct {
	Event *CaptureNFTBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *CaptureNFTBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaptureNFTBatchMetadataUpdate)
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
		it.Event = new(CaptureNFTBatchMetadataUpdate)
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
func (it *CaptureNFTBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaptureNFTBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaptureNFTBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the CaptureNFT contract.
type CaptureNFTBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_CaptureNFT *CaptureNFTFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*CaptureNFTBatchMetadataUpdateIterator, error) {

	logs, sub, err := _CaptureNFT.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &CaptureNFTBatchMetadataUpdateIterator{contract: _CaptureNFT.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_CaptureNFT *CaptureNFTFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *CaptureNFTBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _CaptureNFT.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaptureNFTBatchMetadataUpdate)
				if err := _CaptureNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_CaptureNFT *CaptureNFTFilterer) ParseBatchMetadataUpdate(log types.Log) (*CaptureNFTBatchMetadataUpdate, error) {
	event := new(CaptureNFTBatchMetadataUpdate)
	if err := _CaptureNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaptureNFTCaptureAttemptedIterator is returned from FilterCaptureAttempted and is used to iterate over the raw logs and unpacked data for CaptureAttempted events raised by the CaptureNFT contract.
type CaptureNFTCaptureAttemptedIterator struct {
	Event *CaptureNFTCaptureAttempted // Event containing the contract specifics and raw log

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
func (it *CaptureNFTCaptureAttemptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaptureNFTCaptureAttempted)
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
		it.Event = new(CaptureNFTCaptureAttempted)
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
func (it *CaptureNFTCaptureAttemptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaptureNFTCaptureAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaptureNFTCaptureAttempted represents a CaptureAttempted event raised by the CaptureNFT contract.
type CaptureNFTCaptureAttempted struct {
	Player     common.Address
	CaptureId  common.Hash
	Success    bool
	Rarity     uint8
	Difficulty *big.Int
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCaptureAttempted is a free log retrieval operation binding the contract event 0x22c64a0f9995bb2e2b95b6707c7eca921ff65a0d6fcd325311b722c6329d17da.
//
// Solidity: event CaptureAttempted(address indexed player, string indexed captureId, bool success, uint8 rarity, uint256 difficulty, uint256 tokenId)
func (_CaptureNFT *CaptureNFTFilterer) FilterCaptureAttempted(opts *bind.FilterOpts, player []common.Address, captureId []string) (*CaptureNFTCaptureAttemptedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var captureIdRule []interface{}
	for _, captureIdItem := range captureId {
		captureIdRule = append(captureIdRule, captureIdItem)
	}

	logs, sub, err := _CaptureNFT.contract.FilterLogs(opts, "CaptureAttempted", playerRule, captureIdRule)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTCaptureAttemptedIterator{contract: _CaptureNFT.contract, event: "CaptureAttempted", logs: logs, sub: sub}, nil
}

// WatchCaptureAttempted is a free log subscription operation binding the contract event 0x22c64a0f9995bb2e2b95b6707c7eca921ff65a0d6fcd325311b722c6329d17da.
//
// Solidity: event CaptureAttempted(address indexed player, string indexed captureId, bool success, uint8 rarity, uint256 difficulty, uint256 tokenId)
func (_CaptureNFT *CaptureNFTFilterer) WatchCaptureAttempted(opts *bind.WatchOpts, sink chan<- *CaptureNFTCaptureAttempted, player []common.Address, captureId []string) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var captureIdRule []interface{}
	for _, captureIdItem := range captureId {
		captureIdRule = append(captureIdRule, captureIdItem)
	}

	logs, sub, err := _CaptureNFT.contract.WatchLogs(opts, "CaptureAttempted", playerRule, captureIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaptureNFTCaptureAttempted)
				if err := _CaptureNFT.contract.UnpackLog(event, "CaptureAttempted", log); err != nil {
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

// ParseCaptureAttempted is a log parse operation binding the contract event 0x22c64a0f9995bb2e2b95b6707c7eca921ff65a0d6fcd325311b722c6329d17da.
//
// Solidity: event CaptureAttempted(address indexed player, string indexed captureId, bool success, uint8 rarity, uint256 difficulty, uint256 tokenId)
func (_CaptureNFT *CaptureNFTFilterer) ParseCaptureAttempted(log types.Log) (*CaptureNFTCaptureAttempted, error) {
	event := new(CaptureNFTCaptureAttempted)
	if err := _CaptureNFT.contract.UnpackLog(event, "CaptureAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaptureNFTMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the CaptureNFT contract.
type CaptureNFTMetadataUpdateIterator struct {
	Event *CaptureNFTMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *CaptureNFTMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaptureNFTMetadataUpdate)
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
		it.Event = new(CaptureNFTMetadataUpdate)
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
func (it *CaptureNFTMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaptureNFTMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaptureNFTMetadataUpdate represents a MetadataUpdate event raised by the CaptureNFT contract.
type CaptureNFTMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_CaptureNFT *CaptureNFTFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*CaptureNFTMetadataUpdateIterator, error) {

	logs, sub, err := _CaptureNFT.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &CaptureNFTMetadataUpdateIterator{contract: _CaptureNFT.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_CaptureNFT *CaptureNFTFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *CaptureNFTMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _CaptureNFT.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaptureNFTMetadataUpdate)
				if err := _CaptureNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_CaptureNFT *CaptureNFTFilterer) ParseMetadataUpdate(log types.Log) (*CaptureNFTMetadataUpdate, error) {
	event := new(CaptureNFTMetadataUpdate)
	if err := _CaptureNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaptureNFTMinterAuthorizedIterator is returned from FilterMinterAuthorized and is used to iterate over the raw logs and unpacked data for MinterAuthorized events raised by the CaptureNFT contract.
type CaptureNFTMinterAuthorizedIterator struct {
	Event *CaptureNFTMinterAuthorized // Event containing the contract specifics and raw log

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
func (it *CaptureNFTMinterAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaptureNFTMinterAuthorized)
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
		it.Event = new(CaptureNFTMinterAuthorized)
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
func (it *CaptureNFTMinterAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaptureNFTMinterAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaptureNFTMinterAuthorized represents a MinterAuthorized event raised by the CaptureNFT contract.
type CaptureNFTMinterAuthorized struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterAuthorized is a free log retrieval operation binding the contract event 0x83b05b6735acd4b85e3bded8e72c851d1a87718f81e3c8e6f0c9d9a2baa88e46.
//
// Solidity: event MinterAuthorized(address indexed minter)
func (_CaptureNFT *CaptureNFTFilterer) FilterMinterAuthorized(opts *bind.FilterOpts, minter []common.Address) (*CaptureNFTMinterAuthorizedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _CaptureNFT.contract.FilterLogs(opts, "MinterAuthorized", minterRule)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTMinterAuthorizedIterator{contract: _CaptureNFT.contract, event: "MinterAuthorized", logs: logs, sub: sub}, nil
}

// WatchMinterAuthorized is a free log subscription operation binding the contract event 0x83b05b6735acd4b85e3bded8e72c851d1a87718f81e3c8e6f0c9d9a2baa88e46.
//
// Solidity: event MinterAuthorized(address indexed minter)
func (_CaptureNFT *CaptureNFTFilterer) WatchMinterAuthorized(opts *bind.WatchOpts, sink chan<- *CaptureNFTMinterAuthorized, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _CaptureNFT.contract.WatchLogs(opts, "MinterAuthorized", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaptureNFTMinterAuthorized)
				if err := _CaptureNFT.contract.UnpackLog(event, "MinterAuthorized", log); err != nil {
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

// ParseMinterAuthorized is a log parse operation binding the contract event 0x83b05b6735acd4b85e3bded8e72c851d1a87718f81e3c8e6f0c9d9a2baa88e46.
//
// Solidity: event MinterAuthorized(address indexed minter)
func (_CaptureNFT *CaptureNFTFilterer) ParseMinterAuthorized(log types.Log) (*CaptureNFTMinterAuthorized, error) {
	event := new(CaptureNFTMinterAuthorized)
	if err := _CaptureNFT.contract.UnpackLog(event, "MinterAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaptureNFTMinterRevokedIterator is returned from FilterMinterRevoked and is used to iterate over the raw logs and unpacked data for MinterRevoked events raised by the CaptureNFT contract.
type CaptureNFTMinterRevokedIterator struct {
	Event *CaptureNFTMinterRevoked // Event containing the contract specifics and raw log

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
func (it *CaptureNFTMinterRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaptureNFTMinterRevoked)
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
		it.Event = new(CaptureNFTMinterRevoked)
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
func (it *CaptureNFTMinterRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaptureNFTMinterRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaptureNFTMinterRevoked represents a MinterRevoked event raised by the CaptureNFT contract.
type CaptureNFTMinterRevoked struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterRevoked is a free log retrieval operation binding the contract event 0x44f4322f8daa225d5f4877ad0f7d3dfba248a774396f3ca99405ed40a044fe81.
//
// Solidity: event MinterRevoked(address indexed minter)
func (_CaptureNFT *CaptureNFTFilterer) FilterMinterRevoked(opts *bind.FilterOpts, minter []common.Address) (*CaptureNFTMinterRevokedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _CaptureNFT.contract.FilterLogs(opts, "MinterRevoked", minterRule)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTMinterRevokedIterator{contract: _CaptureNFT.contract, event: "MinterRevoked", logs: logs, sub: sub}, nil
}

// WatchMinterRevoked is a free log subscription operation binding the contract event 0x44f4322f8daa225d5f4877ad0f7d3dfba248a774396f3ca99405ed40a044fe81.
//
// Solidity: event MinterRevoked(address indexed minter)
func (_CaptureNFT *CaptureNFTFilterer) WatchMinterRevoked(opts *bind.WatchOpts, sink chan<- *CaptureNFTMinterRevoked, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _CaptureNFT.contract.WatchLogs(opts, "MinterRevoked", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaptureNFTMinterRevoked)
				if err := _CaptureNFT.contract.UnpackLog(event, "MinterRevoked", log); err != nil {
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

// ParseMinterRevoked is a log parse operation binding the contract event 0x44f4322f8daa225d5f4877ad0f7d3dfba248a774396f3ca99405ed40a044fe81.
//
// Solidity: event MinterRevoked(address indexed minter)
func (_CaptureNFT *CaptureNFTFilterer) ParseMinterRevoked(log types.Log) (*CaptureNFTMinterRevoked, error) {
	event := new(CaptureNFTMinterRevoked)
	if err := _CaptureNFT.contract.UnpackLog(event, "MinterRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaptureNFTNFTMintedIterator is returned from FilterNFTMinted and is used to iterate over the raw logs and unpacked data for NFTMinted events raised by the CaptureNFT contract.
type CaptureNFTNFTMintedIterator struct {
	Event *CaptureNFTNFTMinted // Event containing the contract specifics and raw log

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
func (it *CaptureNFTNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaptureNFTNFTMinted)
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
		it.Event = new(CaptureNFTNFTMinted)
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
func (it *CaptureNFTNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaptureNFTNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaptureNFTNFTMinted represents a NFTMinted event raised by the CaptureNFT contract.
type CaptureNFTNFTMinted struct {
	To        common.Address
	TokenId   *big.Int
	Rarity    uint8
	CaptureId string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNFTMinted is a free log retrieval operation binding the contract event 0xa08b3bf5b7af36b46fffc856a242e512024f345d8981c820b40ed15e1a7b5cae.
//
// Solidity: event NFTMinted(address indexed to, uint256 indexed tokenId, uint8 rarity, string captureId)
func (_CaptureNFT *CaptureNFTFilterer) FilterNFTMinted(opts *bind.FilterOpts, to []common.Address, tokenId []*big.Int) (*CaptureNFTNFTMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CaptureNFT.contract.FilterLogs(opts, "NFTMinted", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTNFTMintedIterator{contract: _CaptureNFT.contract, event: "NFTMinted", logs: logs, sub: sub}, nil
}

// WatchNFTMinted is a free log subscription operation binding the contract event 0xa08b3bf5b7af36b46fffc856a242e512024f345d8981c820b40ed15e1a7b5cae.
//
// Solidity: event NFTMinted(address indexed to, uint256 indexed tokenId, uint8 rarity, string captureId)
func (_CaptureNFT *CaptureNFTFilterer) WatchNFTMinted(opts *bind.WatchOpts, sink chan<- *CaptureNFTNFTMinted, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CaptureNFT.contract.WatchLogs(opts, "NFTMinted", toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaptureNFTNFTMinted)
				if err := _CaptureNFT.contract.UnpackLog(event, "NFTMinted", log); err != nil {
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

// ParseNFTMinted is a log parse operation binding the contract event 0xa08b3bf5b7af36b46fffc856a242e512024f345d8981c820b40ed15e1a7b5cae.
//
// Solidity: event NFTMinted(address indexed to, uint256 indexed tokenId, uint8 rarity, string captureId)
func (_CaptureNFT *CaptureNFTFilterer) ParseNFTMinted(log types.Log) (*CaptureNFTNFTMinted, error) {
	event := new(CaptureNFTNFTMinted)
	if err := _CaptureNFT.contract.UnpackLog(event, "NFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaptureNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CaptureNFT contract.
type CaptureNFTOwnershipTransferredIterator struct {
	Event *CaptureNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CaptureNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaptureNFTOwnershipTransferred)
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
		it.Event = new(CaptureNFTOwnershipTransferred)
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
func (it *CaptureNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaptureNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaptureNFTOwnershipTransferred represents a OwnershipTransferred event raised by the CaptureNFT contract.
type CaptureNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CaptureNFT *CaptureNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CaptureNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CaptureNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTOwnershipTransferredIterator{contract: _CaptureNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CaptureNFT *CaptureNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CaptureNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CaptureNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaptureNFTOwnershipTransferred)
				if err := _CaptureNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CaptureNFT *CaptureNFTFilterer) ParseOwnershipTransferred(log types.Log) (*CaptureNFTOwnershipTransferred, error) {
	event := new(CaptureNFTOwnershipTransferred)
	if err := _CaptureNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaptureNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CaptureNFT contract.
type CaptureNFTTransferIterator struct {
	Event *CaptureNFTTransfer // Event containing the contract specifics and raw log

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
func (it *CaptureNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaptureNFTTransfer)
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
		it.Event = new(CaptureNFTTransfer)
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
func (it *CaptureNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaptureNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaptureNFTTransfer represents a Transfer event raised by the CaptureNFT contract.
type CaptureNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CaptureNFT *CaptureNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CaptureNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CaptureNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CaptureNFTTransferIterator{contract: _CaptureNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CaptureNFT *CaptureNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CaptureNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CaptureNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaptureNFTTransfer)
				if err := _CaptureNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CaptureNFT *CaptureNFTFilterer) ParseTransfer(log types.Log) (*CaptureNFTTransfer, error) {
	event := new(CaptureNFTTransfer)
	if err := _CaptureNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
