// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package medicalDataStorage

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

// MedicalDataStorageMetaData contains all meta data concerning the MedicalDataStorage contract.
var MedicalDataStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedInfo\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"fillUSerInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"}],\"name\":\"getUSerInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"}],\"name\":\"testDecryption\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"}],\"name\":\"testEncryption\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610df0806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630528e7581461005157806308e562fd146100815780631a3796511461009d5780638c2e125e146100cd575b600080fd5b61006b600480360381019061006691906106ae565b6100fd565b60405161007891906107a5565b60405180910390f35b61009b60048036038101906100969190610825565b610127565b005b6100b760048036038101906100b29190610881565b610177565b6040516100c4919061094e565b60405180910390f35b6100e760048036038101906100e29190610970565b6101a7565b6040516100f491906107a5565b60405180910390f35b6060600061010a83610281565b905061011e6101198583610291565b61039a565b91505092915050565b816000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090816101729190610be2565b505050565b6060600061018484610281565b9050600061019184610281565b905061019d82826103aa565b9250505092915050565b60606102796000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546101f5906109fb565b80601f0160208091040260200160405190810160405280929190818152602001828054610221906109fb565b801561026e5780601f106102435761010080835404028352916020019161026e565b820191906000526020600020905b81548152906001019060200180831161025157829003601f168201915b5050505050836100fd565b905092915050565b6060600082905080915050919050565b60606000835167ffffffffffffffff8111156102b0576102af6104e2565b5b6040519080825280601f01601f1916602001820160405280156102e25781602001600182028036833780820191505090505b50905060005b845181101561038f57838451826102ff9190610ce3565b815181106103105761030f610d14565b5b602001015160f81c60f81b85828151811061032e5761032d610d14565b5b602001015160f81c60f81b1882828151811061034d5761034c610d14565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808061038790610d72565b9150506102e8565b508091505092915050565b6060600082905080915050919050565b60606000835167ffffffffffffffff8111156103c9576103c86104e2565b5b6040519080825280601f01601f1916602001820160405280156103fb5781602001600182028036833780820191505090505b50905060005b84518110156104a857838451826104189190610ce3565b8151811061042957610428610d14565b5b602001015160f81c60f81b85828151811061044757610446610d14565b5b602001015160f81c60f81b1882828151811061046657610465610d14565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806104a090610d72565b915050610401565b508091505092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61051a826104d1565b810181811067ffffffffffffffff82111715610539576105386104e2565b5b80604052505050565b600061054c6104b3565b90506105588282610511565b919050565b600067ffffffffffffffff821115610578576105776104e2565b5b610581826104d1565b9050602081019050919050565b82818337600083830152505050565b60006105b06105ab8461055d565b610542565b9050828152602081018484840111156105cc576105cb6104cc565b5b6105d784828561058e565b509392505050565b600082601f8301126105f4576105f36104c7565b5b813561060484826020860161059d565b91505092915050565b600067ffffffffffffffff821115610628576106276104e2565b5b610631826104d1565b9050602081019050919050565b600061065161064c8461060d565b610542565b90508281526020810184848401111561066d5761066c6104cc565b5b61067884828561058e565b509392505050565b600082601f830112610695576106946104c7565b5b81356106a584826020860161063e565b91505092915050565b600080604083850312156106c5576106c46104bd565b5b600083013567ffffffffffffffff8111156106e3576106e26104c2565b5b6106ef858286016105df565b925050602083013567ffffffffffffffff8111156107105761070f6104c2565b5b61071c85828601610680565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610760578082015181840152602081019050610745565b60008484015250505050565b600061077782610726565b6107818185610731565b9350610791818560208601610742565b61079a816104d1565b840191505092915050565b600060208201905081810360008301526107bf818461076c565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107f2826107c7565b9050919050565b610802816107e7565b811461080d57600080fd5b50565b60008135905061081f816107f9565b92915050565b6000806040838503121561083c5761083b6104bd565b5b600083013567ffffffffffffffff81111561085a576108596104c2565b5b610866858286016105df565b925050602061087785828601610810565b9150509250929050565b60008060408385031215610898576108976104bd565b5b600083013567ffffffffffffffff8111156108b6576108b56104c2565b5b6108c285828601610680565b925050602083013567ffffffffffffffff8111156108e3576108e26104c2565b5b6108ef85828601610680565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000610920826108f9565b61092a8185610904565b935061093a818560208601610742565b610943816104d1565b840191505092915050565b600060208201905081810360008301526109688184610915565b905092915050565b60008060408385031215610987576109866104bd565b5b600061099585828601610810565b925050602083013567ffffffffffffffff8111156109b6576109b56104c2565b5b6109c285828601610680565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610a1357607f821691505b602082108103610a2657610a256109cc565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610a8e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a51565b610a988683610a51565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610adf610ada610ad584610ab0565b610aba565b610ab0565b9050919050565b6000819050919050565b610af983610ac4565b610b0d610b0582610ae6565b848454610a5e565b825550505050565b600090565b610b22610b15565b610b2d818484610af0565b505050565b5b81811015610b5157610b46600082610b1a565b600181019050610b33565b5050565b601f821115610b9657610b6781610a2c565b610b7084610a41565b81016020851015610b7f578190505b610b93610b8b85610a41565b830182610b32565b50505b505050565b600082821c905092915050565b6000610bb960001984600802610b9b565b1980831691505092915050565b6000610bd28383610ba8565b9150826002028217905092915050565b610beb826108f9565b67ffffffffffffffff811115610c0457610c036104e2565b5b610c0e82546109fb565b610c19828285610b55565b600060209050601f831160018114610c4c5760008415610c3a578287015190505b610c448582610bc6565b865550610cac565b601f198416610c5a86610a2c565b60005b82811015610c8257848901518255600182019150602085019450602081019050610c5d565b86831015610c9f5784890151610c9b601f891682610ba8565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610cee82610ab0565b9150610cf983610ab0565b925082610d0957610d08610cb4565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d7d82610ab0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610daf57610dae610d43565b5b60018201905091905056fea2646970667358221220629a912273d2155f599080eec74bce7685a481b39a34c0c2979a7dcbb71c517964736f6c63430008130033",
}

// MedicalDataStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use MedicalDataStorageMetaData.ABI instead.
var MedicalDataStorageABI = MedicalDataStorageMetaData.ABI

// MedicalDataStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MedicalDataStorageMetaData.Bin instead.
var MedicalDataStorageBin = MedicalDataStorageMetaData.Bin

// DeployMedicalDataStorage deploys a new Ethereum contract, binding an instance of MedicalDataStorage to it.
func DeployMedicalDataStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MedicalDataStorage, error) {
	parsed, err := MedicalDataStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MedicalDataStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MedicalDataStorage{MedicalDataStorageCaller: MedicalDataStorageCaller{contract: contract}, MedicalDataStorageTransactor: MedicalDataStorageTransactor{contract: contract}, MedicalDataStorageFilterer: MedicalDataStorageFilterer{contract: contract}}, nil
}

// MedicalDataStorage is an auto generated Go binding around an Ethereum contract.
type MedicalDataStorage struct {
	MedicalDataStorageCaller     // Read-only binding to the contract
	MedicalDataStorageTransactor // Write-only binding to the contract
	MedicalDataStorageFilterer   // Log filterer for contract events
}

// MedicalDataStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type MedicalDataStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedicalDataStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MedicalDataStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedicalDataStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MedicalDataStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedicalDataStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MedicalDataStorageSession struct {
	Contract     *MedicalDataStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MedicalDataStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MedicalDataStorageCallerSession struct {
	Contract *MedicalDataStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MedicalDataStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MedicalDataStorageTransactorSession struct {
	Contract     *MedicalDataStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MedicalDataStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type MedicalDataStorageRaw struct {
	Contract *MedicalDataStorage // Generic contract binding to access the raw methods on
}

// MedicalDataStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MedicalDataStorageCallerRaw struct {
	Contract *MedicalDataStorageCaller // Generic read-only contract binding to access the raw methods on
}

// MedicalDataStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MedicalDataStorageTransactorRaw struct {
	Contract *MedicalDataStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMedicalDataStorage creates a new instance of MedicalDataStorage, bound to a specific deployed contract.
func NewMedicalDataStorage(address common.Address, backend bind.ContractBackend) (*MedicalDataStorage, error) {
	contract, err := bindMedicalDataStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MedicalDataStorage{MedicalDataStorageCaller: MedicalDataStorageCaller{contract: contract}, MedicalDataStorageTransactor: MedicalDataStorageTransactor{contract: contract}, MedicalDataStorageFilterer: MedicalDataStorageFilterer{contract: contract}}, nil
}

// NewMedicalDataStorageCaller creates a new read-only instance of MedicalDataStorage, bound to a specific deployed contract.
func NewMedicalDataStorageCaller(address common.Address, caller bind.ContractCaller) (*MedicalDataStorageCaller, error) {
	contract, err := bindMedicalDataStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MedicalDataStorageCaller{contract: contract}, nil
}

// NewMedicalDataStorageTransactor creates a new write-only instance of MedicalDataStorage, bound to a specific deployed contract.
func NewMedicalDataStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*MedicalDataStorageTransactor, error) {
	contract, err := bindMedicalDataStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MedicalDataStorageTransactor{contract: contract}, nil
}

// NewMedicalDataStorageFilterer creates a new log filterer instance of MedicalDataStorage, bound to a specific deployed contract.
func NewMedicalDataStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*MedicalDataStorageFilterer, error) {
	contract, err := bindMedicalDataStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MedicalDataStorageFilterer{contract: contract}, nil
}

// bindMedicalDataStorage binds a generic wrapper to an already deployed contract.
func bindMedicalDataStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MedicalDataStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MedicalDataStorage *MedicalDataStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MedicalDataStorage.Contract.MedicalDataStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MedicalDataStorage *MedicalDataStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MedicalDataStorage.Contract.MedicalDataStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MedicalDataStorage *MedicalDataStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MedicalDataStorage.Contract.MedicalDataStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MedicalDataStorage *MedicalDataStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MedicalDataStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MedicalDataStorage *MedicalDataStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MedicalDataStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MedicalDataStorage *MedicalDataStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MedicalDataStorage.Contract.contract.Transact(opts, method, params...)
}

// GetUSerInfo is a free data retrieval call binding the contract method 0x8c2e125e.
//
// Solidity: function getUSerInfo(address userAddress, string password) view returns(string)
func (_MedicalDataStorage *MedicalDataStorageCaller) GetUSerInfo(opts *bind.CallOpts, userAddress common.Address, password string) (string, error) {
	var out []interface{}
	err := _MedicalDataStorage.contract.Call(opts, &out, "getUSerInfo", userAddress, password)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUSerInfo is a free data retrieval call binding the contract method 0x8c2e125e.
//
// Solidity: function getUSerInfo(address userAddress, string password) view returns(string)
func (_MedicalDataStorage *MedicalDataStorageSession) GetUSerInfo(userAddress common.Address, password string) (string, error) {
	return _MedicalDataStorage.Contract.GetUSerInfo(&_MedicalDataStorage.CallOpts, userAddress, password)
}

// GetUSerInfo is a free data retrieval call binding the contract method 0x8c2e125e.
//
// Solidity: function getUSerInfo(address userAddress, string password) view returns(string)
func (_MedicalDataStorage *MedicalDataStorageCallerSession) GetUSerInfo(userAddress common.Address, password string) (string, error) {
	return _MedicalDataStorage.Contract.GetUSerInfo(&_MedicalDataStorage.CallOpts, userAddress, password)
}

// TestDecryption is a free data retrieval call binding the contract method 0x0528e758.
//
// Solidity: function testDecryption(bytes encryptedData, string _password) pure returns(string)
func (_MedicalDataStorage *MedicalDataStorageCaller) TestDecryption(opts *bind.CallOpts, encryptedData []byte, _password string) (string, error) {
	var out []interface{}
	err := _MedicalDataStorage.contract.Call(opts, &out, "testDecryption", encryptedData, _password)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TestDecryption is a free data retrieval call binding the contract method 0x0528e758.
//
// Solidity: function testDecryption(bytes encryptedData, string _password) pure returns(string)
func (_MedicalDataStorage *MedicalDataStorageSession) TestDecryption(encryptedData []byte, _password string) (string, error) {
	return _MedicalDataStorage.Contract.TestDecryption(&_MedicalDataStorage.CallOpts, encryptedData, _password)
}

// TestDecryption is a free data retrieval call binding the contract method 0x0528e758.
//
// Solidity: function testDecryption(bytes encryptedData, string _password) pure returns(string)
func (_MedicalDataStorage *MedicalDataStorageCallerSession) TestDecryption(encryptedData []byte, _password string) (string, error) {
	return _MedicalDataStorage.Contract.TestDecryption(&_MedicalDataStorage.CallOpts, encryptedData, _password)
}

// TestEncryption is a free data retrieval call binding the contract method 0x1a379651.
//
// Solidity: function testEncryption(string _data, string _password) pure returns(bytes)
func (_MedicalDataStorage *MedicalDataStorageCaller) TestEncryption(opts *bind.CallOpts, _data string, _password string) ([]byte, error) {
	var out []interface{}
	err := _MedicalDataStorage.contract.Call(opts, &out, "testEncryption", _data, _password)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestEncryption is a free data retrieval call binding the contract method 0x1a379651.
//
// Solidity: function testEncryption(string _data, string _password) pure returns(bytes)
func (_MedicalDataStorage *MedicalDataStorageSession) TestEncryption(_data string, _password string) ([]byte, error) {
	return _MedicalDataStorage.Contract.TestEncryption(&_MedicalDataStorage.CallOpts, _data, _password)
}

// TestEncryption is a free data retrieval call binding the contract method 0x1a379651.
//
// Solidity: function testEncryption(string _data, string _password) pure returns(bytes)
func (_MedicalDataStorage *MedicalDataStorageCallerSession) TestEncryption(_data string, _password string) ([]byte, error) {
	return _MedicalDataStorage.Contract.TestEncryption(&_MedicalDataStorage.CallOpts, _data, _password)
}

// FillUSerInfo is a paid mutator transaction binding the contract method 0x08e562fd.
//
// Solidity: function fillUSerInfo(bytes encryptedInfo, address userAddress) returns()
func (_MedicalDataStorage *MedicalDataStorageTransactor) FillUSerInfo(opts *bind.TransactOpts, encryptedInfo []byte, userAddress common.Address) (*types.Transaction, error) {
	return _MedicalDataStorage.contract.Transact(opts, "fillUSerInfo", encryptedInfo, userAddress)
}

// FillUSerInfo is a paid mutator transaction binding the contract method 0x08e562fd.
//
// Solidity: function fillUSerInfo(bytes encryptedInfo, address userAddress) returns()
func (_MedicalDataStorage *MedicalDataStorageSession) FillUSerInfo(encryptedInfo []byte, userAddress common.Address) (*types.Transaction, error) {
	return _MedicalDataStorage.Contract.FillUSerInfo(&_MedicalDataStorage.TransactOpts, encryptedInfo, userAddress)
}

// FillUSerInfo is a paid mutator transaction binding the contract method 0x08e562fd.
//
// Solidity: function fillUSerInfo(bytes encryptedInfo, address userAddress) returns()
func (_MedicalDataStorage *MedicalDataStorageTransactorSession) FillUSerInfo(encryptedInfo []byte, userAddress common.Address) (*types.Transaction, error) {
	return _MedicalDataStorage.Contract.FillUSerInfo(&_MedicalDataStorage.TransactOpts, encryptedInfo, userAddress)
}
