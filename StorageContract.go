// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdcToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"QuotaPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"QuotaUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"pricePerQuotaChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getQuota\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerQuota\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"purchasedQuota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"quotas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPricePerQuota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"useQuota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

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
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contract *ContractCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contract *ContractSession) Admin() (common.Address, error) {
	return _Contract.Contract.Admin(&_Contract.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contract *ContractCallerSession) Admin() (common.Address, error) {
	return _Contract.Contract.Admin(&_Contract.CallOpts)
}

// GetQuota is a free data retrieval call binding the contract method 0x43a2a302.
//
// Solidity: function getQuota(address user) view returns(uint256)
func (_Contract *ContractCaller) GetQuota(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getQuota", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuota is a free data retrieval call binding the contract method 0x43a2a302.
//
// Solidity: function getQuota(address user) view returns(uint256)
func (_Contract *ContractSession) GetQuota(user common.Address) (*big.Int, error) {
	return _Contract.Contract.GetQuota(&_Contract.CallOpts, user)
}

// GetQuota is a free data retrieval call binding the contract method 0x43a2a302.
//
// Solidity: function getQuota(address user) view returns(uint256)
func (_Contract *ContractCallerSession) GetQuota(user common.Address) (*big.Int, error) {
	return _Contract.Contract.GetQuota(&_Contract.CallOpts, user)
}

// PricePerQuota is a free data retrieval call binding the contract method 0x0fccbbf1.
//
// Solidity: function pricePerQuota() view returns(uint256)
func (_Contract *ContractCaller) PricePerQuota(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pricePerQuota")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerQuota is a free data retrieval call binding the contract method 0x0fccbbf1.
//
// Solidity: function pricePerQuota() view returns(uint256)
func (_Contract *ContractSession) PricePerQuota() (*big.Int, error) {
	return _Contract.Contract.PricePerQuota(&_Contract.CallOpts)
}

// PricePerQuota is a free data retrieval call binding the contract method 0x0fccbbf1.
//
// Solidity: function pricePerQuota() view returns(uint256)
func (_Contract *ContractCallerSession) PricePerQuota() (*big.Int, error) {
	return _Contract.Contract.PricePerQuota(&_Contract.CallOpts)
}

// Quotas is a free data retrieval call binding the contract method 0xc33342e9.
//
// Solidity: function quotas(address ) view returns(uint256)
func (_Contract *ContractCaller) Quotas(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "quotas", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quotas is a free data retrieval call binding the contract method 0xc33342e9.
//
// Solidity: function quotas(address ) view returns(uint256)
func (_Contract *ContractSession) Quotas(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Quotas(&_Contract.CallOpts, arg0)
}

// Quotas is a free data retrieval call binding the contract method 0xc33342e9.
//
// Solidity: function quotas(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Quotas(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Quotas(&_Contract.CallOpts, arg0)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Contract *ContractCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Contract *ContractSession) UsdcToken() (common.Address, error) {
	return _Contract.Contract.UsdcToken(&_Contract.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Contract *ContractCallerSession) UsdcToken() (common.Address, error) {
	return _Contract.Contract.UsdcToken(&_Contract.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Contract *ContractTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Contract *ContractSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeAdmin(&_Contract.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Contract *ContractTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeAdmin(&_Contract.TransactOpts, newAdmin)
}

// PurchasedQuota is a paid mutator transaction binding the contract method 0x04ce078a.
//
// Solidity: function purchasedQuota(uint256 amount) returns()
func (_Contract *ContractTransactor) PurchasedQuota(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "purchasedQuota", amount)
}

// PurchasedQuota is a paid mutator transaction binding the contract method 0x04ce078a.
//
// Solidity: function purchasedQuota(uint256 amount) returns()
func (_Contract *ContractSession) PurchasedQuota(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PurchasedQuota(&_Contract.TransactOpts, amount)
}

// PurchasedQuota is a paid mutator transaction binding the contract method 0x04ce078a.
//
// Solidity: function purchasedQuota(uint256 amount) returns()
func (_Contract *ContractTransactorSession) PurchasedQuota(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PurchasedQuota(&_Contract.TransactOpts, amount)
}

// SetPricePerQuota is a paid mutator transaction binding the contract method 0xfd72fa08.
//
// Solidity: function setPricePerQuota(uint256 _price) returns()
func (_Contract *ContractTransactor) SetPricePerQuota(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPricePerQuota", _price)
}

// SetPricePerQuota is a paid mutator transaction binding the contract method 0xfd72fa08.
//
// Solidity: function setPricePerQuota(uint256 _price) returns()
func (_Contract *ContractSession) SetPricePerQuota(_price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPricePerQuota(&_Contract.TransactOpts, _price)
}

// SetPricePerQuota is a paid mutator transaction binding the contract method 0xfd72fa08.
//
// Solidity: function setPricePerQuota(uint256 _price) returns()
func (_Contract *ContractTransactorSession) SetPricePerQuota(_price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPricePerQuota(&_Contract.TransactOpts, _price)
}

// UseQuota is a paid mutator transaction binding the contract method 0x75012172.
//
// Solidity: function useQuota(address user, uint256 amount) returns()
func (_Contract *ContractTransactor) UseQuota(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "useQuota", user, amount)
}

// UseQuota is a paid mutator transaction binding the contract method 0x75012172.
//
// Solidity: function useQuota(address user, uint256 amount) returns()
func (_Contract *ContractSession) UseQuota(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UseQuota(&_Contract.TransactOpts, user, amount)
}

// UseQuota is a paid mutator transaction binding the contract method 0x75012172.
//
// Solidity: function useQuota(address user, uint256 amount) returns()
func (_Contract *ContractTransactorSession) UseQuota(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UseQuota(&_Contract.TransactOpts, user, amount)
}

// ContractAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Contract contract.
type ContractAdminChangedIterator struct {
	Event *ContractAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAdminChanged)
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
		it.Event = new(ContractAdminChanged)
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
func (it *ContractAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAdminChanged represents a AdminChanged event raised by the Contract contract.
type ContractAdminChanged struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7ce7ec0b50378fb6c0186ffb5f48325f6593fcb4ca4386f21861af3129188f5c.
//
// Solidity: event AdminChanged(address indexed newAdmin)
func (_Contract *ContractFilterer) FilterAdminChanged(opts *bind.FilterOpts, newAdmin []common.Address) (*ContractAdminChangedIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AdminChanged", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &ContractAdminChangedIterator{contract: _Contract.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7ce7ec0b50378fb6c0186ffb5f48325f6593fcb4ca4386f21861af3129188f5c.
//
// Solidity: event AdminChanged(address indexed newAdmin)
func (_Contract *ContractFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractAdminChanged, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AdminChanged", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAdminChanged)
				if err := _Contract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7ce7ec0b50378fb6c0186ffb5f48325f6593fcb4ca4386f21861af3129188f5c.
//
// Solidity: event AdminChanged(address indexed newAdmin)
func (_Contract *ContractFilterer) ParseAdminChanged(log types.Log) (*ContractAdminChanged, error) {
	event := new(ContractAdminChanged)
	if err := _Contract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractQuotaPurchasedIterator is returned from FilterQuotaPurchased and is used to iterate over the raw logs and unpacked data for QuotaPurchased events raised by the Contract contract.
type ContractQuotaPurchasedIterator struct {
	Event *ContractQuotaPurchased // Event containing the contract specifics and raw log

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
func (it *ContractQuotaPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractQuotaPurchased)
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
		it.Event = new(ContractQuotaPurchased)
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
func (it *ContractQuotaPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractQuotaPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractQuotaPurchased represents a QuotaPurchased event raised by the Contract contract.
type ContractQuotaPurchased struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterQuotaPurchased is a free log retrieval operation binding the contract event 0x459d0b85bc62ccfd34d3f5e2787efde5bdfa38862497479098c2526a97f72fd1.
//
// Solidity: event QuotaPurchased(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) FilterQuotaPurchased(opts *bind.FilterOpts, user []common.Address) (*ContractQuotaPurchasedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "QuotaPurchased", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractQuotaPurchasedIterator{contract: _Contract.contract, event: "QuotaPurchased", logs: logs, sub: sub}, nil
}

// WatchQuotaPurchased is a free log subscription operation binding the contract event 0x459d0b85bc62ccfd34d3f5e2787efde5bdfa38862497479098c2526a97f72fd1.
//
// Solidity: event QuotaPurchased(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) WatchQuotaPurchased(opts *bind.WatchOpts, sink chan<- *ContractQuotaPurchased, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "QuotaPurchased", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractQuotaPurchased)
				if err := _Contract.contract.UnpackLog(event, "QuotaPurchased", log); err != nil {
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

// ParseQuotaPurchased is a log parse operation binding the contract event 0x459d0b85bc62ccfd34d3f5e2787efde5bdfa38862497479098c2526a97f72fd1.
//
// Solidity: event QuotaPurchased(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) ParseQuotaPurchased(log types.Log) (*ContractQuotaPurchased, error) {
	event := new(ContractQuotaPurchased)
	if err := _Contract.contract.UnpackLog(event, "QuotaPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractQuotaUsedIterator is returned from FilterQuotaUsed and is used to iterate over the raw logs and unpacked data for QuotaUsed events raised by the Contract contract.
type ContractQuotaUsedIterator struct {
	Event *ContractQuotaUsed // Event containing the contract specifics and raw log

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
func (it *ContractQuotaUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractQuotaUsed)
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
		it.Event = new(ContractQuotaUsed)
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
func (it *ContractQuotaUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractQuotaUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractQuotaUsed represents a QuotaUsed event raised by the Contract contract.
type ContractQuotaUsed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterQuotaUsed is a free log retrieval operation binding the contract event 0xa4c312ab2850e31593ed2e6492d2f35592ffcd845401078cc7a91ca8a82576ef.
//
// Solidity: event QuotaUsed(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) FilterQuotaUsed(opts *bind.FilterOpts, user []common.Address) (*ContractQuotaUsedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "QuotaUsed", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractQuotaUsedIterator{contract: _Contract.contract, event: "QuotaUsed", logs: logs, sub: sub}, nil
}

// WatchQuotaUsed is a free log subscription operation binding the contract event 0xa4c312ab2850e31593ed2e6492d2f35592ffcd845401078cc7a91ca8a82576ef.
//
// Solidity: event QuotaUsed(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) WatchQuotaUsed(opts *bind.WatchOpts, sink chan<- *ContractQuotaUsed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "QuotaUsed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractQuotaUsed)
				if err := _Contract.contract.UnpackLog(event, "QuotaUsed", log); err != nil {
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

// ParseQuotaUsed is a log parse operation binding the contract event 0xa4c312ab2850e31593ed2e6492d2f35592ffcd845401078cc7a91ca8a82576ef.
//
// Solidity: event QuotaUsed(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) ParseQuotaUsed(log types.Log) (*ContractQuotaUsed, error) {
	event := new(ContractQuotaUsed)
	if err := _Contract.contract.UnpackLog(event, "QuotaUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPricePerQuotaChangedIterator is returned from FilterPricePerQuotaChanged and is used to iterate over the raw logs and unpacked data for PricePerQuotaChanged events raised by the Contract contract.
type ContractPricePerQuotaChangedIterator struct {
	Event *ContractPricePerQuotaChanged // Event containing the contract specifics and raw log

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
func (it *ContractPricePerQuotaChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPricePerQuotaChanged)
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
		it.Event = new(ContractPricePerQuotaChanged)
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
func (it *ContractPricePerQuotaChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPricePerQuotaChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPricePerQuotaChanged represents a PricePerQuotaChanged event raised by the Contract contract.
type ContractPricePerQuotaChanged struct {
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPricePerQuotaChanged is a free log retrieval operation binding the contract event 0xc0cbd58778448e2decd5a8bf641e7610c6681c7a85656ce8717aec6ac4d4e900.
//
// Solidity: event pricePerQuotaChanged(uint256 newPrice)
func (_Contract *ContractFilterer) FilterPricePerQuotaChanged(opts *bind.FilterOpts) (*ContractPricePerQuotaChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "pricePerQuotaChanged")
	if err != nil {
		return nil, err
	}
	return &ContractPricePerQuotaChangedIterator{contract: _Contract.contract, event: "pricePerQuotaChanged", logs: logs, sub: sub}, nil
}

// WatchPricePerQuotaChanged is a free log subscription operation binding the contract event 0xc0cbd58778448e2decd5a8bf641e7610c6681c7a85656ce8717aec6ac4d4e900.
//
// Solidity: event pricePerQuotaChanged(uint256 newPrice)
func (_Contract *ContractFilterer) WatchPricePerQuotaChanged(opts *bind.WatchOpts, sink chan<- *ContractPricePerQuotaChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "pricePerQuotaChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPricePerQuotaChanged)
				if err := _Contract.contract.UnpackLog(event, "pricePerQuotaChanged", log); err != nil {
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

// ParsePricePerQuotaChanged is a log parse operation binding the contract event 0xc0cbd58778448e2decd5a8bf641e7610c6681c7a85656ce8717aec6ac4d4e900.
//
// Solidity: event pricePerQuotaChanged(uint256 newPrice)
func (_Contract *ContractFilterer) ParsePricePerQuotaChanged(log types.Log) (*ContractPricePerQuotaChanged, error) {
	event := new(ContractPricePerQuotaChanged)
	if err := _Contract.contract.UnpackLog(event, "pricePerQuotaChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
