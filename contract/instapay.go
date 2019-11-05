// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ejections\",\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\"},{\"name\":\"stage\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"create_channel\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"readme\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"owner_bal\",\"type\":\"uint256\"},{\"name\":\"receiver_bal\",\"type\":\"uint256\"}],\"name\":\"close_channel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pn\",\"type\":\"uint256\"},{\"name\":\"stage\",\"type\":\"uint8\"},{\"name\":\"ids\",\"type\":\"uint256[]\"},{\"name\":\"bals\",\"type\":\"uint256[]\"},{\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"eject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"channels\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"EventCreateChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"owner_bal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"receiver_bal\",\"type\":\"uint256\"}],\"name\":\"EventCloseChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"registered_stage\",\"type\":\"uint8\"}],\"name\":\"EventEject\",\"type\":\"event\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels(uint256 ) constant returns(address owner, address receiver, uint256 deposit, uint8 status)
func (_Contract *ContractCaller) Channels(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Contract.contract.Call(opts, out, "channels", arg0)
	return *ret, err
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels(uint256 ) constant returns(address owner, address receiver, uint256 deposit, uint8 status)
func (_Contract *ContractSession) Channels(arg0 *big.Int) (struct {
	Owner    common.Address
	Receiver common.Address
	Deposit  *big.Int
	Status   uint8
}, error) {
	return _Contract.Contract.Channels(&_Contract.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels(uint256 ) constant returns(address owner, address receiver, uint256 deposit, uint8 status)
func (_Contract *ContractCallerSession) Channels(arg0 *big.Int) (struct {
	Owner    common.Address
	Receiver common.Address
	Deposit  *big.Int
	Status   uint8
}, error) {
	return _Contract.Contract.Channels(&_Contract.CallOpts, arg0)
}

// Ejections is a free data retrieval call binding the contract method 0x0c934b83.
//
// Solidity: function ejections(uint256 ) constant returns(bool registered, uint8 stage)
func (_Contract *ContractCaller) Ejections(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Registered bool
	Stage      uint8
}, error) {
	ret := new(struct {
		Registered bool
		Stage      uint8
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "ejections", arg0)
	return *ret, err
}

// Ejections is a free data retrieval call binding the contract method 0x0c934b83.
//
// Solidity: function ejections(uint256 ) constant returns(bool registered, uint8 stage)
func (_Contract *ContractSession) Ejections(arg0 *big.Int) (struct {
	Registered bool
	Stage      uint8
}, error) {
	return _Contract.Contract.Ejections(&_Contract.CallOpts, arg0)
}

// Ejections is a free data retrieval call binding the contract method 0x0c934b83.
//
// Solidity: function ejections(uint256 ) constant returns(bool registered, uint8 stage)
func (_Contract *ContractCallerSession) Ejections(arg0 *big.Int) (struct {
	Registered bool
	Stage      uint8
}, error) {
	return _Contract.Contract.Ejections(&_Contract.CallOpts, arg0)
}

// Readme is a free data retrieval call binding the contract method 0x82092f1d.
//
// Solidity: function readme() constant returns(uint256)
func (_Contract *ContractCaller) Readme(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "readme")
	return *ret0, err
}

// Readme is a free data retrieval call binding the contract method 0x82092f1d.
//
// Solidity: function readme() constant returns(uint256)
func (_Contract *ContractSession) Readme() (*big.Int, error) {
	return _Contract.Contract.Readme(&_Contract.CallOpts)
}

// Readme is a free data retrieval call binding the contract method 0x82092f1d.
//
// Solidity: function readme() constant returns(uint256)
func (_Contract *ContractCallerSession) Readme() (*big.Int, error) {
	return _Contract.Contract.Readme(&_Contract.CallOpts)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xbe4dab58.
//
// Solidity: function close_channel(uint256 id, uint256 owner_bal, uint256 receiver_bal) returns()
func (_Contract *ContractTransactor) CloseChannel(opts *bind.TransactOpts, id *big.Int, owner_bal *big.Int, receiver_bal *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "close_channel", id, owner_bal, receiver_bal)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xbe4dab58.
//
// Solidity: function close_channel(uint256 id, uint256 owner_bal, uint256 receiver_bal) returns()
func (_Contract *ContractSession) CloseChannel(id *big.Int, owner_bal *big.Int, receiver_bal *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CloseChannel(&_Contract.TransactOpts, id, owner_bal, receiver_bal)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xbe4dab58.
//
// Solidity: function close_channel(uint256 id, uint256 owner_bal, uint256 receiver_bal) returns()
func (_Contract *ContractTransactorSession) CloseChannel(id *big.Int, owner_bal *big.Int, receiver_bal *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CloseChannel(&_Contract.TransactOpts, id, owner_bal, receiver_bal)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x10cae21a.
//
// Solidity: function create_channel(address receiver) returns()
func (_Contract *ContractTransactor) CreateChannel(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "create_channel", receiver)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x10cae21a.
//
// Solidity: function create_channel(address receiver) returns()
func (_Contract *ContractSession) CreateChannel(receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CreateChannel(&_Contract.TransactOpts, receiver)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x10cae21a.
//
// Solidity: function create_channel(address receiver) returns()
func (_Contract *ContractTransactorSession) CreateChannel(receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CreateChannel(&_Contract.TransactOpts, receiver)
}

// Eject is a paid mutator transaction binding the contract method 0xc0dba051.
//
// Solidity: function eject(uint256 pn, uint8 stage, uint256[] ids, uint256[] bals, uint256 v) returns()
func (_Contract *ContractTransactor) Eject(opts *bind.TransactOpts, pn *big.Int, stage uint8, ids []*big.Int, bals []*big.Int, v *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "eject", pn, stage, ids, bals, v)
}

// Eject is a paid mutator transaction binding the contract method 0xc0dba051.
//
// Solidity: function eject(uint256 pn, uint8 stage, uint256[] ids, uint256[] bals, uint256 v) returns()
func (_Contract *ContractSession) Eject(pn *big.Int, stage uint8, ids []*big.Int, bals []*big.Int, v *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Eject(&_Contract.TransactOpts, pn, stage, ids, bals, v)
}

// Eject is a paid mutator transaction binding the contract method 0xc0dba051.
//
// Solidity: function eject(uint256 pn, uint8 stage, uint256[] ids, uint256[] bals, uint256 v) returns()
func (_Contract *ContractTransactorSession) Eject(pn *big.Int, stage uint8, ids []*big.Int, bals []*big.Int, v *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Eject(&_Contract.TransactOpts, pn, stage, ids, bals, v)
}

// ContractEventCloseChannelIterator is returned from FilterEventCloseChannel and is used to iterate over the raw logs and unpacked data for EventCloseChannel events raised by the Contract contract.
type ContractEventCloseChannelIterator struct {
	Event *ContractEventCloseChannel // Event containing the contract specifics and raw log

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
func (it *ContractEventCloseChannelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventCloseChannel)
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
		it.Event = new(ContractEventCloseChannel)
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
func (it *ContractEventCloseChannelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventCloseChannelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventCloseChannel represents a EventCloseChannel event raised by the Contract contract.
type ContractEventCloseChannel struct {
	Id          *big.Int
	OwnerBal    *big.Int
	ReceiverBal *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEventCloseChannel is a free log retrieval operation binding the contract event 0x7e25e653d41214b20f80c31e3ae0b8d5d51e00bde0c7206e469ba489299d3bc3.
//
// Solidity: event EventCloseChannel(uint256 id, uint256 owner_bal, uint256 receiver_bal)
func (_Contract *ContractFilterer) FilterEventCloseChannel(opts *bind.FilterOpts) (*ContractEventCloseChannelIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EventCloseChannel")
	if err != nil {
		return nil, err
	}
	return &ContractEventCloseChannelIterator{contract: _Contract.contract, event: "EventCloseChannel", logs: logs, sub: sub}, nil
}

// WatchEventCloseChannel is a free log subscription operation binding the contract event 0x7e25e653d41214b20f80c31e3ae0b8d5d51e00bde0c7206e469ba489299d3bc3.
//
// Solidity: event EventCloseChannel(uint256 id, uint256 owner_bal, uint256 receiver_bal)
func (_Contract *ContractFilterer) WatchEventCloseChannel(opts *bind.WatchOpts, sink chan<- *ContractEventCloseChannel) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EventCloseChannel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventCloseChannel)
				if err := _Contract.contract.UnpackLog(event, "EventCloseChannel", log); err != nil {
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
// Solidity: event EventCloseChannel(uint256 id, uint256 owner_bal, uint256 receiver_bal)
func (_Contract *ContractFilterer) ParseEventCloseChannel(log types.Log) (*ContractEventCloseChannel, error) {
	event := new(ContractEventCloseChannel)
	if err := _Contract.contract.UnpackLog(event, "EventCloseChannel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractEventCreateChannelIterator is returned from FilterEventCreateChannel and is used to iterate over the raw logs and unpacked data for EventCreateChannel events raised by the Contract contract.
type ContractEventCreateChannelIterator struct {
	Event *ContractEventCreateChannel // Event containing the contract specifics and raw log

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
func (it *ContractEventCreateChannelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventCreateChannel)
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
		it.Event = new(ContractEventCreateChannel)
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
func (it *ContractEventCreateChannelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventCreateChannelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventCreateChannel represents a EventCreateChannel event raised by the Contract contract.
type ContractEventCreateChannel struct {
	Id       *big.Int
	Owner    common.Address
	Receiver common.Address
	Deposit  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventCreateChannel is a free log retrieval operation binding the contract event 0xad9d6cc8ba120612b7481ea698f5bb11e08fc633d7042eb1526d0efa3ce18855.
//
// Solidity: event EventCreateChannel(uint256 id, address owner, address receiver, uint256 deposit)
func (_Contract *ContractFilterer) FilterEventCreateChannel(opts *bind.FilterOpts) (*ContractEventCreateChannelIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EventCreateChannel")
	if err != nil {
		return nil, err
	}
	return &ContractEventCreateChannelIterator{contract: _Contract.contract, event: "EventCreateChannel", logs: logs, sub: sub}, nil
}

// WatchEventCreateChannel is a free log subscription operation binding the contract event 0xad9d6cc8ba120612b7481ea698f5bb11e08fc633d7042eb1526d0efa3ce18855.
//
// Solidity: event EventCreateChannel(uint256 id, address owner, address receiver, uint256 deposit)
func (_Contract *ContractFilterer) WatchEventCreateChannel(opts *bind.WatchOpts, sink chan<- *ContractEventCreateChannel) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EventCreateChannel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventCreateChannel)
				if err := _Contract.contract.UnpackLog(event, "EventCreateChannel", log); err != nil {
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
func (_Contract *ContractFilterer) ParseEventCreateChannel(log types.Log) (*ContractEventCreateChannel, error) {
	event := new(ContractEventCreateChannel)
	if err := _Contract.contract.UnpackLog(event, "EventCreateChannel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractEventEjectIterator is returned from FilterEventEject and is used to iterate over the raw logs and unpacked data for EventEject events raised by the Contract contract.
type ContractEventEjectIterator struct {
	Event *ContractEventEject // Event containing the contract specifics and raw log

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
func (it *ContractEventEjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventEject)
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
		it.Event = new(ContractEventEject)
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
func (it *ContractEventEjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventEjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventEject represents a EventEject event raised by the Contract contract.
type ContractEventEject struct {
	Pn              *big.Int
	RegisteredStage uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEventEject is a free log retrieval operation binding the contract event 0x9e40c6df849462cf9049ce149eef0782ae1be76a54193ab139b751d683ebd353.
//
// Solidity: event EventEject(uint256 pn, uint8 registered_stage)
func (_Contract *ContractFilterer) FilterEventEject(opts *bind.FilterOpts) (*ContractEventEjectIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EventEject")
	if err != nil {
		return nil, err
	}
	return &ContractEventEjectIterator{contract: _Contract.contract, event: "EventEject", logs: logs, sub: sub}, nil
}

// WatchEventEject is a free log subscription operation binding the contract event 0x9e40c6df849462cf9049ce149eef0782ae1be76a54193ab139b751d683ebd353.
//
// Solidity: event EventEject(uint256 pn, uint8 registered_stage)
func (_Contract *ContractFilterer) WatchEventEject(opts *bind.WatchOpts, sink chan<- *ContractEventEject) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EventEject")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventEject)
				if err := _Contract.contract.UnpackLog(event, "EventEject", log); err != nil {
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
// Solidity: event EventEject(uint256 pn, uint8 registered_stage)
func (_Contract *ContractFilterer) ParseEventEject(log types.Log) (*ContractEventEject, error) {
	event := new(ContractEventEject)
	if err := _Contract.contract.UnpackLog(event, "EventEject", log); err != nil {
		return nil, err
	}
	return event, nil
}
