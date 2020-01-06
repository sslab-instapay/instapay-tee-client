// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package instapay

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// InstapayABI is the input ABI used to generate the binding from.
const InstapayABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ejections\",\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\"},{\"name\":\"stage\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"create_channel\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"readme\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"owner_bal\",\"type\":\"uint256\"},{\"name\":\"receiver_bal\",\"type\":\"uint256\"}],\"name\":\"close_channel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pn\",\"type\":\"uint256\"},{\"name\":\"stage\",\"type\":\"uint8\"},{\"name\":\"ids\",\"type\":\"uint256[]\"},{\"name\":\"bals\",\"type\":\"uint256[]\"},{\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"eject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"channels\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"EventCreateChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ownerbal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"receiverbal\",\"type\":\"uint256\"}],\"name\":\"EventCloseChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"registeredstage\",\"type\":\"uint8\"}],\"name\":\"EventEject\",\"type\":\"event\"}]"

// InstapayBin is the compiled bytecode used for deploying new contracts.
var InstapayBin = "0x6080604052600080556110e160015534801561001a57600080fd5b50610e118061002a6000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630c934b831461007d57806310cae21a146100d757806382092f1d1461010d578063be4dab5814610138578063c0dba05114610179578063e5949b5d14610243575b600080fd5b34801561008957600080fd5b506100a8600480360381019080803590602001909291905050506102ff565b60405180831515151581526020018260018111156100c257fe5b60ff1681526020019250505060405180910390f35b61010b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061033d565b005b34801561011957600080fd5b50610122610548565b6040518082815260200191505060405180910390f35b34801561014457600080fd5b5061017760048036038101908080359060200190929190803590602001909291908035906020019092919050505061054e565b005b34801561018557600080fd5b5061024160048036038101908080359060200190929190803560ff169060200190929190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919291929080359060200190929190505050610807565b005b34801561024f57600080fd5b5061026e6004803603810190808035906020019092919050505061099c565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018260018111156102e857fe5b60ff16815260200194505050505060405180910390f35b60036020528060005260406000206000915090508060000160009054906101000a900460ff16908060000160019054906101000a900460ff16905082565b6000600181111561034a57fe5b60026000600160005401815260200190815260200160002060030160009054906101000a900460ff16600181111561037e57fe5b14151561038a57600080fd5b600080815480929190600101919050555033600260008054815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260008054815260200190815260200160002060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550346002600080548152602001908152602001600020600201819055506001600260008054815260200190815260200160002060030160006101000a81548160ff0219169083600181111561049757fe5b02179055507fad9d6cc8ba120612b7481ea698f5bb11e08fc633d7042eb1526d0efa3ce18855600054338334604051808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390a150565b60015481565b82338073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148061062157508073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b151561062c57600080fd5b60018081111561063857fe5b6002600087815260200190815260200160002060030160009054906101000a900460ff16600181111561066757fe5b14151561067357600080fd5b6002600086815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc670de0b6b3a764000086029081150290604051600060405180830381858888f193505050501580156106f9573d6000803e3d6000fd5b506002600086815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc670de0b6b3a764000085029081150290604051600060405180830381858888f19350505050158015610780573d6000803e3d6000fd5b5060006002600087815260200190815260200160002060030160006101000a81548160ff021916908360018111156107b457fe5b02179055507f7e25e653d41214b20f80c31e3ae0b8d5d51e00bde0c7206e469ba489299d3bc385858560405180848152602001838152602001828152602001935050505060405180910390a15050505050565b6000600115156003600088815260200190815260200160002060000160009054906101000a900460ff161515141561086e576108696003600088815260200190815260200160002060000160019054906101000a900460ff1686868686610a19565b610994565b60016003600088815260200190815260200160002060000160006101000a81548160ff021916908315150217905550846003600088815260200190815260200160002060000160016101000a81548160ff021916908360018111156108cf57fe5b0217905550600090505b83518110156109465761093984828151811015156108f357fe5b906020019060200201518483840181518110151561090d57fe5b906020019060200201518560018586010181518110151561092a57fe5b9060200190602002015161054e565b80806001019150506108d9565b7f9e40c6df849462cf9049ce149eef0782ae1be76a54193ab139b751d683ebd35386866040518083815260200182600181111561097f57fe5b60ff1681526020019250505060405180910390a15b505050505050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030160009054906101000a900460ff16905084565b60008060008060006001811115610a2c57fe5b896001811115610a3857fe5b1415610c0e5760006001811115610a4b57fe5b886001811115610a5757fe5b1415610b3357600093505b8651841015610b2e5760006001811115610a7857fe5b600260008987815181101515610a8a57fe5b90602001906020020151815260200190815260200160002060030160009054906101000a900460ff166001811115610abe57fe5b1415610ac957610b21565b610b208785815181101515610ada57fe5b9060200190602002015187868701815181101515610af457fe5b9060200190602002015188600188890101815181101515610b1157fe5b9060200190602002015161054e565b5b8380600101945050610a62565b610c09565b600092505b8651831015610c085760006001811115610b4e57fe5b600260008986815181101515610b6057fe5b90602001906020020151815260200190815260200160002060030160009054906101000a900460ff166001811115610b9457fe5b1415610b9f57610bfb565b610bfa8784815181101515610bb057fe5b906020019060200201518688868701815181101515610bcb57fe5b90602001906020020151018789600188890101815181101515610bea57fe5b906020019060200201510361054e565b5b8280600101935050610b38565b5b610dda565b60006001811115610c1b57fe5b886001811115610c2757fe5b1415610d0757600091505b8651821015610d025760006001811115610c4857fe5b600260008985815181101515610c5a57fe5b90602001906020020151815260200190815260200160002060030160009054906101000a900460ff166001811115610c8e57fe5b1415610c9957610cf5565b610cf48783815181101515610caa57fe5b906020019060200201518688858601815181101515610cc557fe5b90602001906020020151038789600187880101815181101515610ce457fe5b906020019060200201510161054e565b5b8180600101925050610c32565b610dd9565b600090505b8651811015610dd85760006001811115610d2257fe5b600260008984815181101515610d3457fe5b90602001906020020151815260200190815260200160002060030160009054906101000a900460ff166001811115610d6857fe5b1415610d7357610dcb565b610dca8782815181101515610d8457fe5b9060200190602002015187838401815181101515610d9e57fe5b9060200190602002015188600185860101815181101515610dbb57fe5b9060200190602002015161054e565b5b8080600101915050610d0c565b5b5b5050505050505050505600a165627a7a72305820fb5f6e82f251305630b60be14340a976cde73bf2bce825858d5e9e8bced9aa530029"

// DeployInstapay deploys a new Ethereum contract, binding an instance of Instapay to it.
func DeployInstapay(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Instapay, error) {
	parsed, err := abi.JSON(strings.NewReader(InstapayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InstapayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Instapay{InstapayCaller: InstapayCaller{contract: contract}, InstapayTransactor: InstapayTransactor{contract: contract}, InstapayFilterer: InstapayFilterer{contract: contract}}, nil
}

// Instapay is an auto generated Go binding around an Ethereum contract.
type Instapay struct {
	InstapayCaller     // Read-only binding to the contract
	InstapayTransactor // Write-only binding to the contract
	InstapayFilterer   // Log filterer for contract events
}

// InstapayCaller is an auto generated read-only Go binding around an Ethereum contract.
type InstapayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstapayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InstapayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstapayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InstapayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstapaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InstapaySession struct {
	Contract     *Instapay         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InstapayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InstapayCallerSession struct {
	Contract *InstapayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// InstapayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InstapayTransactorSession struct {
	Contract     *InstapayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InstapayRaw is an auto generated low-level Go binding around an Ethereum contract.
type InstapayRaw struct {
	Contract *Instapay // Generic contract binding to access the raw methods on
}

// InstapayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InstapayCallerRaw struct {
	Contract *InstapayCaller // Generic read-only contract binding to access the raw methods on
}

// InstapayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InstapayTransactorRaw struct {
	Contract *InstapayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInstapay creates a new instance of Instapay, bound to a specific deployed contract.
func NewInstapay(address common.Address, backend bind.ContractBackend) (*Instapay, error) {
	contract, err := bindInstapay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Instapay{InstapayCaller: InstapayCaller{contract: contract}, InstapayTransactor: InstapayTransactor{contract: contract}, InstapayFilterer: InstapayFilterer{contract: contract}}, nil
}

// NewInstapayCaller creates a new read-only instance of Instapay, bound to a specific deployed contract.
func NewInstapayCaller(address common.Address, caller bind.ContractCaller) (*InstapayCaller, error) {
	contract, err := bindInstapay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InstapayCaller{contract: contract}, nil
}

// NewInstapayTransactor creates a new write-only instance of Instapay, bound to a specific deployed contract.
func NewInstapayTransactor(address common.Address, transactor bind.ContractTransactor) (*InstapayTransactor, error) {
	contract, err := bindInstapay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InstapayTransactor{contract: contract}, nil
}

// NewInstapayFilterer creates a new log filterer instance of Instapay, bound to a specific deployed contract.
func NewInstapayFilterer(address common.Address, filterer bind.ContractFilterer) (*InstapayFilterer, error) {
	contract, err := bindInstapay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InstapayFilterer{contract: contract}, nil
}

// bindInstapay binds a generic wrapper to an already deployed contract.
func bindInstapay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InstapayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Instapay *InstapayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Instapay.Contract.InstapayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Instapay *InstapayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Instapay.Contract.InstapayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Instapay *InstapayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Instapay.Contract.InstapayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Instapay *InstapayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Instapay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Instapay *InstapayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Instapay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Instapay *InstapayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Instapay.Contract.contract.Transact(opts, method, params...)
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels(uint256 ) constant returns(address owner, address receiver, uint256 deposit, uint8 status)
func (_Instapay *InstapayCaller) Channels(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner    common.Address
	Receiver common.Address
	Deposit  *big.Int
	Status   uint8
}, error) {
	ret := new(struct {
		Owner    common.Address
		Receiver common.Address
		Deposit  *big.Int
		Status   uint8
	})
	out := ret
	err := _Instapay.contract.Call(opts, out, "channels", arg0)
	return *ret, err
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels(uint256 ) constant returns(address owner, address receiver, uint256 deposit, uint8 status)
func (_Instapay *InstapaySession) Channels(arg0 *big.Int) (struct {
	Owner    common.Address
	Receiver common.Address
	Deposit  *big.Int
	Status   uint8
}, error) {
	return _Instapay.Contract.Channels(&_Instapay.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels(uint256 ) constant returns(address owner, address receiver, uint256 deposit, uint8 status)
func (_Instapay *InstapayCallerSession) Channels(arg0 *big.Int) (struct {
	Owner    common.Address
	Receiver common.Address
	Deposit  *big.Int
	Status   uint8
}, error) {
	return _Instapay.Contract.Channels(&_Instapay.CallOpts, arg0)
}

// Ejections is a free data retrieval call binding the contract method 0x0c934b83.
//
// Solidity: function ejections(uint256 ) constant returns(bool registered, uint8 stage)
func (_Instapay *InstapayCaller) Ejections(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Registered bool
	Stage      uint8
}, error) {
	ret := new(struct {
		Registered bool
		Stage      uint8
	})
	out := ret
	err := _Instapay.contract.Call(opts, out, "ejections", arg0)
	return *ret, err
}

// Ejections is a free data retrieval call binding the contract method 0x0c934b83.
//
// Solidity: function ejections(uint256 ) constant returns(bool registered, uint8 stage)
func (_Instapay *InstapaySession) Ejections(arg0 *big.Int) (struct {
	Registered bool
	Stage      uint8
}, error) {
	return _Instapay.Contract.Ejections(&_Instapay.CallOpts, arg0)
}

// Ejections is a free data retrieval call binding the contract method 0x0c934b83.
//
// Solidity: function ejections(uint256 ) constant returns(bool registered, uint8 stage)
func (_Instapay *InstapayCallerSession) Ejections(arg0 *big.Int) (struct {
	Registered bool
	Stage      uint8
}, error) {
	return _Instapay.Contract.Ejections(&_Instapay.CallOpts, arg0)
}

// Readme is a free data retrieval call binding the contract method 0x82092f1d.
//
// Solidity: function readme() constant returns(uint256)
func (_Instapay *InstapayCaller) Readme(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Instapay.contract.Call(opts, out, "readme")
	return *ret0, err
}

// Readme is a free data retrieval call binding the contract method 0x82092f1d.
//
// Solidity: function readme() constant returns(uint256)
func (_Instapay *InstapaySession) Readme() (*big.Int, error) {
	return _Instapay.Contract.Readme(&_Instapay.CallOpts)
}

// Readme is a free data retrieval call binding the contract method 0x82092f1d.
//
// Solidity: function readme() constant returns(uint256)
func (_Instapay *InstapayCallerSession) Readme() (*big.Int, error) {
	return _Instapay.Contract.Readme(&_Instapay.CallOpts)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xbe4dab58.
//
// Solidity: function close_channel(uint256 id, uint256 owner_bal, uint256 receiver_bal) returns()
func (_Instapay *InstapayTransactor) CloseChannel(opts *bind.TransactOpts, id *big.Int, owner_bal *big.Int, receiver_bal *big.Int) (*types.Transaction, error) {
	return _Instapay.contract.Transact(opts, "close_channel", id, owner_bal, receiver_bal)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xbe4dab58.
//
// Solidity: function close_channel(uint256 id, uint256 owner_bal, uint256 receiver_bal) returns()
func (_Instapay *InstapaySession) CloseChannel(id *big.Int, owner_bal *big.Int, receiver_bal *big.Int) (*types.Transaction, error) {
	return _Instapay.Contract.CloseChannel(&_Instapay.TransactOpts, id, owner_bal, receiver_bal)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xbe4dab58.
//
// Solidity: function close_channel(uint256 id, uint256 owner_bal, uint256 receiver_bal) returns()
func (_Instapay *InstapayTransactorSession) CloseChannel(id *big.Int, owner_bal *big.Int, receiver_bal *big.Int) (*types.Transaction, error) {
	return _Instapay.Contract.CloseChannel(&_Instapay.TransactOpts, id, owner_bal, receiver_bal)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x10cae21a.
//
// Solidity: function create_channel(address receiver) returns()
func (_Instapay *InstapayTransactor) CreateChannel(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Instapay.contract.Transact(opts, "create_channel", receiver)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x10cae21a.
//
// Solidity: function create_channel(address receiver) returns()
func (_Instapay *InstapaySession) CreateChannel(receiver common.Address) (*types.Transaction, error) {
	return _Instapay.Contract.CreateChannel(&_Instapay.TransactOpts, receiver)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x10cae21a.
//
// Solidity: function create_channel(address receiver) returns()
func (_Instapay *InstapayTransactorSession) CreateChannel(receiver common.Address) (*types.Transaction, error) {
	return _Instapay.Contract.CreateChannel(&_Instapay.TransactOpts, receiver)
}

// Eject is a paid mutator transaction binding the contract method 0xc0dba051.
//
// Solidity: function eject(uint256 pn, uint8 stage, uint256[] ids, uint256[] bals, uint256 v) returns()
func (_Instapay *InstapayTransactor) Eject(opts *bind.TransactOpts, pn *big.Int, stage uint8, ids []*big.Int, bals []*big.Int, v *big.Int) (*types.Transaction, error) {
	return _Instapay.contract.Transact(opts, "eject", pn, stage, ids, bals, v)
}

// Eject is a paid mutator transaction binding the contract method 0xc0dba051.
//
// Solidity: function eject(uint256 pn, uint8 stage, uint256[] ids, uint256[] bals, uint256 v) returns()
func (_Instapay *InstapaySession) Eject(pn *big.Int, stage uint8, ids []*big.Int, bals []*big.Int, v *big.Int) (*types.Transaction, error) {
	return _Instapay.Contract.Eject(&_Instapay.TransactOpts, pn, stage, ids, bals, v)
}

// Eject is a paid mutator transaction binding the contract method 0xc0dba051.
//
// Solidity: function eject(uint256 pn, uint8 stage, uint256[] ids, uint256[] bals, uint256 v) returns()
func (_Instapay *InstapayTransactorSession) Eject(pn *big.Int, stage uint8, ids []*big.Int, bals []*big.Int, v *big.Int) (*types.Transaction, error) {
	return _Instapay.Contract.Eject(&_Instapay.TransactOpts, pn, stage, ids, bals, v)
}

// InstapayEventCloseChannelIterator is returned from FilterEventCloseChannel and is used to iterate over the raw logs and unpacked data for EventCloseChannel events raised by the Instapay contract.
type InstapayEventCloseChannelIterator struct {
	Event *InstapayEventCloseChannel // Event containing the contract specifics and raw log

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
func (it *InstapayEventCloseChannelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstapayEventCloseChannel)
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
		it.Event = new(InstapayEventCloseChannel)
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
func (it *InstapayEventCloseChannelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstapayEventCloseChannelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstapayEventCloseChannel represents a EventCloseChannel event raised by the Instapay contract.
type InstapayEventCloseChannel struct {
	Id          *big.Int
	Ownerbal    *big.Int
	Receiverbal *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEventCloseChannel is a free log retrieval operation binding the contract event 0x7e25e653d41214b20f80c31e3ae0b8d5d51e00bde0c7206e469ba489299d3bc3.
//
// Solidity: event EventCloseChannel(uint256 id, uint256 ownerbal, uint256 receiverbal)
func (_Instapay *InstapayFilterer) FilterEventCloseChannel(opts *bind.FilterOpts) (*InstapayEventCloseChannelIterator, error) {

	logs, sub, err := _Instapay.contract.FilterLogs(opts, "EventCloseChannel")
	if err != nil {
		return nil, err
	}
	return &InstapayEventCloseChannelIterator{contract: _Instapay.contract, event: "EventCloseChannel", logs: logs, sub: sub}, nil
}

// WatchEventCloseChannel is a free log subscription operation binding the contract event 0x7e25e653d41214b20f80c31e3ae0b8d5d51e00bde0c7206e469ba489299d3bc3.
//
// Solidity: event EventCloseChannel(uint256 id, uint256 ownerbal, uint256 receiverbal)
func (_Instapay *InstapayFilterer) WatchEventCloseChannel(opts *bind.WatchOpts, sink chan<- *InstapayEventCloseChannel) (event.Subscription, error) {

	logs, sub, err := _Instapay.contract.WatchLogs(opts, "EventCloseChannel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstapayEventCloseChannel)
				if err := _Instapay.contract.UnpackLog(event, "EventCloseChannel", log); err != nil {
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

// ParseEventCloseChannel is a log parse operation binding the contract event 0x7e25e653d41214b20f80c31e3ae0b8d5d51e00bde0c7206e469ba489299d3bc3.
//
// Solidity: event EventCloseChannel(uint256 id, uint256 ownerbal, uint256 receiverbal)
func (_Instapay *InstapayFilterer) ParseEventCloseChannel(log types.Log) (*InstapayEventCloseChannel, error) {
	event := new(InstapayEventCloseChannel)
	if err := _Instapay.contract.UnpackLog(event, "EventCloseChannel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstapayEventCreateChannelIterator is returned from FilterEventCreateChannel and is used to iterate over the raw logs and unpacked data for EventCreateChannel events raised by the Instapay contract.
type InstapayEventCreateChannelIterator struct {
	Event *InstapayEventCreateChannel // Event containing the contract specifics and raw log

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
func (it *InstapayEventCreateChannelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstapayEventCreateChannel)
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
		it.Event = new(InstapayEventCreateChannel)
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
func (it *InstapayEventCreateChannelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstapayEventCreateChannelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstapayEventCreateChannel represents a EventCreateChannel event raised by the Instapay contract.
type InstapayEventCreateChannel struct {
	Id       *big.Int
	Owner    common.Address
	Receiver common.Address
	Deposit  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventCreateChannel is a free log retrieval operation binding the contract event 0xad9d6cc8ba120612b7481ea698f5bb11e08fc633d7042eb1526d0efa3ce18855.
//
// Solidity: event EventCreateChannel(uint256 id, address owner, address receiver, uint256 deposit)
func (_Instapay *InstapayFilterer) FilterEventCreateChannel(opts *bind.FilterOpts) (*InstapayEventCreateChannelIterator, error) {

	logs, sub, err := _Instapay.contract.FilterLogs(opts, "EventCreateChannel")
	if err != nil {
		return nil, err
	}
	return &InstapayEventCreateChannelIterator{contract: _Instapay.contract, event: "EventCreateChannel", logs: logs, sub: sub}, nil
}

// WatchEventCreateChannel is a free log subscription operation binding the contract event 0xad9d6cc8ba120612b7481ea698f5bb11e08fc633d7042eb1526d0efa3ce18855.
//
// Solidity: event EventCreateChannel(uint256 id, address owner, address receiver, uint256 deposit)
func (_Instapay *InstapayFilterer) WatchEventCreateChannel(opts *bind.WatchOpts, sink chan<- *InstapayEventCreateChannel) (event.Subscription, error) {

	logs, sub, err := _Instapay.contract.WatchLogs(opts, "EventCreateChannel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstapayEventCreateChannel)
				if err := _Instapay.contract.UnpackLog(event, "EventCreateChannel", log); err != nil {
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

// ParseEventCreateChannel is a log parse operation binding the contract event 0xad9d6cc8ba120612b7481ea698f5bb11e08fc633d7042eb1526d0efa3ce18855.
//
// Solidity: event EventCreateChannel(uint256 id, address owner, address receiver, uint256 deposit)
func (_Instapay *InstapayFilterer) ParseEventCreateChannel(log types.Log) (*InstapayEventCreateChannel, error) {
	event := new(InstapayEventCreateChannel)
	if err := _Instapay.contract.UnpackLog(event, "EventCreateChannel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InstapayEventEjectIterator is returned from FilterEventEject and is used to iterate over the raw logs and unpacked data for EventEject events raised by the Instapay contract.
type InstapayEventEjectIterator struct {
	Event *InstapayEventEject // Event containing the contract specifics and raw log

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
func (it *InstapayEventEjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstapayEventEject)
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
		it.Event = new(InstapayEventEject)
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
func (it *InstapayEventEjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstapayEventEjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstapayEventEject represents a EventEject event raised by the Instapay contract.
type InstapayEventEject struct {
	Pn              *big.Int
	Registeredstage uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEventEject is a free log retrieval operation binding the contract event 0x9e40c6df849462cf9049ce149eef0782ae1be76a54193ab139b751d683ebd353.
//
// Solidity: event EventEject(uint256 pn, uint8 registeredstage)
func (_Instapay *InstapayFilterer) FilterEventEject(opts *bind.FilterOpts) (*InstapayEventEjectIterator, error) {

	logs, sub, err := _Instapay.contract.FilterLogs(opts, "EventEject")
	if err != nil {
		return nil, err
	}
	return &InstapayEventEjectIterator{contract: _Instapay.contract, event: "EventEject", logs: logs, sub: sub}, nil
}

// WatchEventEject is a free log subscription operation binding the contract event 0x9e40c6df849462cf9049ce149eef0782ae1be76a54193ab139b751d683ebd353.
//
// Solidity: event EventEject(uint256 pn, uint8 registeredstage)
func (_Instapay *InstapayFilterer) WatchEventEject(opts *bind.WatchOpts, sink chan<- *InstapayEventEject) (event.Subscription, error) {

	logs, sub, err := _Instapay.contract.WatchLogs(opts, "EventEject")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstapayEventEject)
				if err := _Instapay.contract.UnpackLog(event, "EventEject", log); err != nil {
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

// ParseEventEject is a log parse operation binding the contract event 0x9e40c6df849462cf9049ce149eef0782ae1be76a54193ab139b751d683ebd353.
//
// Solidity: event EventEject(uint256 pn, uint8 registeredstage)
func (_Instapay *InstapayFilterer) ParseEventEject(log types.Log) (*InstapayEventEject, error) {
	event := new(InstapayEventEject)
	if err := _Instapay.contract.UnpackLog(event, "EventEject", log); err != nil {
		return nil, err
	}
	return event, nil
}
