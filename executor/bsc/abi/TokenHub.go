// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// TokenHubABI is the input ABI used to generate the binding from.
const TokenHubABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"peggyAmount\",\"type\":\"uint256\"}],\"name\":\"LogBindRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"peggyAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"LogBindSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"}],\"name\":\"LogBindRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"}],\"name\":\"LogBindTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"}],\"name\":\"LogBindInvalidParameter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refundAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"expireTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"relayFee\",\"type\":\"uint256\"}],\"name\":\"LogTransferOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"expireTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"relayFee\",\"type\":\"uint256\"}],\"name\":\"LogBatchTransferOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"recipientAddrs\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"refundAddrs\",\"type\":\"address[]\"}],\"name\":\"LogBatchTransferOutAddrs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refundAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"expireTime\",\"type\":\"uint256\"}],\"name\":\"LogTransferInFailureTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refundAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"auctualBalance\",\"type\":\"uint256\"}],\"name\":\"LogTransferInFailureInsufficientBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refundAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"}],\"name\":\"LogTransferInFailureUnboundToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refundAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bep2TokenSymbol\",\"type\":\"bytes32\"}],\"name\":\"LogTransferInFailureUnknownReason\",\"type\":\"event\"}]"

// TokenHub is an auto generated Go binding around an Ethereum contract.
type TokenHub struct {
	TokenHubCaller     // Read-only binding to the contract
	TokenHubTransactor // Write-only binding to the contract
	TokenHubFilterer   // Log filterer for contract events
}

// TokenHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenHubSession struct {
	Contract     *TokenHub         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenHubCallerSession struct {
	Contract *TokenHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TokenHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenHubTransactorSession struct {
	Contract     *TokenHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TokenHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenHubRaw struct {
	Contract *TokenHub // Generic contract binding to access the raw methods on
}

// TokenHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenHubCallerRaw struct {
	Contract *TokenHubCaller // Generic read-only contract binding to access the raw methods on
}

// TokenHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenHubTransactorRaw struct {
	Contract *TokenHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenHub creates a new instance of TokenHub, bound to a specific deployed contract.
func NewTokenHub(address common.Address, backend bind.ContractBackend) (*TokenHub, error) {
	contract, err := bindTokenHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenHub{TokenHubCaller: TokenHubCaller{contract: contract}, TokenHubTransactor: TokenHubTransactor{contract: contract}, TokenHubFilterer: TokenHubFilterer{contract: contract}}, nil
}

// NewTokenHubCaller creates a new read-only instance of TokenHub, bound to a specific deployed contract.
func NewTokenHubCaller(address common.Address, caller bind.ContractCaller) (*TokenHubCaller, error) {
	contract, err := bindTokenHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenHubCaller{contract: contract}, nil
}

// NewTokenHubTransactor creates a new write-only instance of TokenHub, bound to a specific deployed contract.
func NewTokenHubTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenHubTransactor, error) {
	contract, err := bindTokenHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenHubTransactor{contract: contract}, nil
}

// NewTokenHubFilterer creates a new log filterer instance of TokenHub, bound to a specific deployed contract.
func NewTokenHubFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenHubFilterer, error) {
	contract, err := bindTokenHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenHubFilterer{contract: contract}, nil
}

// bindTokenHub binds a generic wrapper to an already deployed contract.
func bindTokenHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenHub *TokenHubRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenHub.Contract.TokenHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenHub *TokenHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenHub.Contract.TokenHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenHub *TokenHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenHub.Contract.TokenHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenHub *TokenHubCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenHub *TokenHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenHub *TokenHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenHub.Contract.contract.Transact(opts, method, params...)
}

// TokenHubLogBatchTransferOutIterator is returned from FilterLogBatchTransferOut and is used to iterate over the raw logs and unpacked data for LogBatchTransferOut events raised by the TokenHub contract.
type TokenHubLogBatchTransferOutIterator struct {
	Event *TokenHubLogBatchTransferOut // Event containing the contract specifics and raw log

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
func (it *TokenHubLogBatchTransferOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogBatchTransferOut)
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
		it.Event = new(TokenHubLogBatchTransferOut)
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
func (it *TokenHubLogBatchTransferOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogBatchTransferOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogBatchTransferOut represents a LogBatchTransferOut event raised by the TokenHub contract.
type TokenHubLogBatchTransferOut struct {
	Sequence        *big.Int
	Amounts         []*big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	ExpireTime      *big.Int
	RelayFee        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogBatchTransferOut is a free log retrieval operation binding the contract event 0x00a18f0343865824d1375c23f5dd79fdf32a12f50400ef2591e52276f8378e31.
//
// Solidity: event LogBatchTransferOut(uint256 sequence, uint256[] amounts, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime, uint256 relayFee)
func (_TokenHub *TokenHubFilterer) FilterLogBatchTransferOut(opts *bind.FilterOpts) (*TokenHubLogBatchTransferOutIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogBatchTransferOut")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogBatchTransferOutIterator{contract: _TokenHub.contract, event: "LogBatchTransferOut", logs: logs, sub: sub}, nil
}

// WatchLogBatchTransferOut is a free log subscription operation binding the contract event 0x00a18f0343865824d1375c23f5dd79fdf32a12f50400ef2591e52276f8378e31.
//
// Solidity: event LogBatchTransferOut(uint256 sequence, uint256[] amounts, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime, uint256 relayFee)
func (_TokenHub *TokenHubFilterer) WatchLogBatchTransferOut(opts *bind.WatchOpts, sink chan<- *TokenHubLogBatchTransferOut) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogBatchTransferOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogBatchTransferOut)
				if err := _TokenHub.contract.UnpackLog(event, "LogBatchTransferOut", log); err != nil {
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

// ParseLogBatchTransferOut is a log parse operation binding the contract event 0x00a18f0343865824d1375c23f5dd79fdf32a12f50400ef2591e52276f8378e31.
//
// Solidity: event LogBatchTransferOut(uint256 sequence, uint256[] amounts, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime, uint256 relayFee)
func (_TokenHub *TokenHubFilterer) ParseLogBatchTransferOut(log types.Log) (*TokenHubLogBatchTransferOut, error) {
	event := new(TokenHubLogBatchTransferOut)
	if err := _TokenHub.contract.UnpackLog(event, "LogBatchTransferOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogBatchTransferOutAddrsIterator is returned from FilterLogBatchTransferOutAddrs and is used to iterate over the raw logs and unpacked data for LogBatchTransferOutAddrs events raised by the TokenHub contract.
type TokenHubLogBatchTransferOutAddrsIterator struct {
	Event *TokenHubLogBatchTransferOutAddrs // Event containing the contract specifics and raw log

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
func (it *TokenHubLogBatchTransferOutAddrsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogBatchTransferOutAddrs)
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
		it.Event = new(TokenHubLogBatchTransferOutAddrs)
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
func (it *TokenHubLogBatchTransferOutAddrsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogBatchTransferOutAddrsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogBatchTransferOutAddrs represents a LogBatchTransferOutAddrs event raised by the TokenHub contract.
type TokenHubLogBatchTransferOutAddrs struct {
	Sequence       *big.Int
	RecipientAddrs []common.Address
	RefundAddrs    []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogBatchTransferOutAddrs is a free log retrieval operation binding the contract event 0x8740bbd4e1a2505bf2908481adbf1056fb52f762152b702f6c65468f63c55cf8.
//
// Solidity: event LogBatchTransferOutAddrs(uint256 sequence, address[] recipientAddrs, address[] refundAddrs)
func (_TokenHub *TokenHubFilterer) FilterLogBatchTransferOutAddrs(opts *bind.FilterOpts) (*TokenHubLogBatchTransferOutAddrsIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogBatchTransferOutAddrs")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogBatchTransferOutAddrsIterator{contract: _TokenHub.contract, event: "LogBatchTransferOutAddrs", logs: logs, sub: sub}, nil
}

// WatchLogBatchTransferOutAddrs is a free log subscription operation binding the contract event 0x8740bbd4e1a2505bf2908481adbf1056fb52f762152b702f6c65468f63c55cf8.
//
// Solidity: event LogBatchTransferOutAddrs(uint256 sequence, address[] recipientAddrs, address[] refundAddrs)
func (_TokenHub *TokenHubFilterer) WatchLogBatchTransferOutAddrs(opts *bind.WatchOpts, sink chan<- *TokenHubLogBatchTransferOutAddrs) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogBatchTransferOutAddrs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogBatchTransferOutAddrs)
				if err := _TokenHub.contract.UnpackLog(event, "LogBatchTransferOutAddrs", log); err != nil {
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

// ParseLogBatchTransferOutAddrs is a log parse operation binding the contract event 0x8740bbd4e1a2505bf2908481adbf1056fb52f762152b702f6c65468f63c55cf8.
//
// Solidity: event LogBatchTransferOutAddrs(uint256 sequence, address[] recipientAddrs, address[] refundAddrs)
func (_TokenHub *TokenHubFilterer) ParseLogBatchTransferOutAddrs(log types.Log) (*TokenHubLogBatchTransferOutAddrs, error) {
	event := new(TokenHubLogBatchTransferOutAddrs)
	if err := _TokenHub.contract.UnpackLog(event, "LogBatchTransferOutAddrs", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogBindInvalidParameterIterator is returned from FilterLogBindInvalidParameter and is used to iterate over the raw logs and unpacked data for LogBindInvalidParameter events raised by the TokenHub contract.
type TokenHubLogBindInvalidParameterIterator struct {
	Event *TokenHubLogBindInvalidParameter // Event containing the contract specifics and raw log

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
func (it *TokenHubLogBindInvalidParameterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogBindInvalidParameter)
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
		it.Event = new(TokenHubLogBindInvalidParameter)
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
func (it *TokenHubLogBindInvalidParameterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogBindInvalidParameterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogBindInvalidParameter represents a LogBindInvalidParameter event raised by the TokenHub contract.
type TokenHubLogBindInvalidParameter struct {
	Sequence        *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogBindInvalidParameter is a free log retrieval operation binding the contract event 0x2117f993c9cc877c531b4e6bd55d822cb48b529fd003c80e5bd6c27b7c1c1702.
//
// Solidity: event LogBindInvalidParameter(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) FilterLogBindInvalidParameter(opts *bind.FilterOpts) (*TokenHubLogBindInvalidParameterIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogBindInvalidParameter")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogBindInvalidParameterIterator{contract: _TokenHub.contract, event: "LogBindInvalidParameter", logs: logs, sub: sub}, nil
}

// WatchLogBindInvalidParameter is a free log subscription operation binding the contract event 0x2117f993c9cc877c531b4e6bd55d822cb48b529fd003c80e5bd6c27b7c1c1702.
//
// Solidity: event LogBindInvalidParameter(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) WatchLogBindInvalidParameter(opts *bind.WatchOpts, sink chan<- *TokenHubLogBindInvalidParameter) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogBindInvalidParameter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogBindInvalidParameter)
				if err := _TokenHub.contract.UnpackLog(event, "LogBindInvalidParameter", log); err != nil {
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

// ParseLogBindInvalidParameter is a log parse operation binding the contract event 0x2117f993c9cc877c531b4e6bd55d822cb48b529fd003c80e5bd6c27b7c1c1702.
//
// Solidity: event LogBindInvalidParameter(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) ParseLogBindInvalidParameter(log types.Log) (*TokenHubLogBindInvalidParameter, error) {
	event := new(TokenHubLogBindInvalidParameter)
	if err := _TokenHub.contract.UnpackLog(event, "LogBindInvalidParameter", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogBindRejectedIterator is returned from FilterLogBindRejected and is used to iterate over the raw logs and unpacked data for LogBindRejected events raised by the TokenHub contract.
type TokenHubLogBindRejectedIterator struct {
	Event *TokenHubLogBindRejected // Event containing the contract specifics and raw log

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
func (it *TokenHubLogBindRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogBindRejected)
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
		it.Event = new(TokenHubLogBindRejected)
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
func (it *TokenHubLogBindRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogBindRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogBindRejected represents a LogBindRejected event raised by the TokenHub contract.
type TokenHubLogBindRejected struct {
	Sequence        *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogBindRejected is a free log retrieval operation binding the contract event 0x341e20b0b6b62cb3990e2d1f8bcb0a15e7d7fd446355a7be807face162285254.
//
// Solidity: event LogBindRejected(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) FilterLogBindRejected(opts *bind.FilterOpts) (*TokenHubLogBindRejectedIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogBindRejected")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogBindRejectedIterator{contract: _TokenHub.contract, event: "LogBindRejected", logs: logs, sub: sub}, nil
}

// WatchLogBindRejected is a free log subscription operation binding the contract event 0x341e20b0b6b62cb3990e2d1f8bcb0a15e7d7fd446355a7be807face162285254.
//
// Solidity: event LogBindRejected(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) WatchLogBindRejected(opts *bind.WatchOpts, sink chan<- *TokenHubLogBindRejected) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogBindRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogBindRejected)
				if err := _TokenHub.contract.UnpackLog(event, "LogBindRejected", log); err != nil {
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

// ParseLogBindRejected is a log parse operation binding the contract event 0x341e20b0b6b62cb3990e2d1f8bcb0a15e7d7fd446355a7be807face162285254.
//
// Solidity: event LogBindRejected(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) ParseLogBindRejected(log types.Log) (*TokenHubLogBindRejected, error) {
	event := new(TokenHubLogBindRejected)
	if err := _TokenHub.contract.UnpackLog(event, "LogBindRejected", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogBindRequestIterator is returned from FilterLogBindRequest and is used to iterate over the raw logs and unpacked data for LogBindRequest events raised by the TokenHub contract.
type TokenHubLogBindRequestIterator struct {
	Event *TokenHubLogBindRequest // Event containing the contract specifics and raw log

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
func (it *TokenHubLogBindRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogBindRequest)
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
		it.Event = new(TokenHubLogBindRequest)
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
func (it *TokenHubLogBindRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogBindRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogBindRequest represents a LogBindRequest event raised by the TokenHub contract.
type TokenHubLogBindRequest struct {
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	TotalSupply     *big.Int
	PeggyAmount     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogBindRequest is a free log retrieval operation binding the contract event 0xf8608cf3c27350e6aa0beaa6723ea6911e3d7353e8b22a69bb112c15f93867ca.
//
// Solidity: event LogBindRequest(address contractAddr, bytes32 bep2TokenSymbol, uint256 totalSupply, uint256 peggyAmount)
func (_TokenHub *TokenHubFilterer) FilterLogBindRequest(opts *bind.FilterOpts) (*TokenHubLogBindRequestIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogBindRequest")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogBindRequestIterator{contract: _TokenHub.contract, event: "LogBindRequest", logs: logs, sub: sub}, nil
}

// WatchLogBindRequest is a free log subscription operation binding the contract event 0xf8608cf3c27350e6aa0beaa6723ea6911e3d7353e8b22a69bb112c15f93867ca.
//
// Solidity: event LogBindRequest(address contractAddr, bytes32 bep2TokenSymbol, uint256 totalSupply, uint256 peggyAmount)
func (_TokenHub *TokenHubFilterer) WatchLogBindRequest(opts *bind.WatchOpts, sink chan<- *TokenHubLogBindRequest) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogBindRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogBindRequest)
				if err := _TokenHub.contract.UnpackLog(event, "LogBindRequest", log); err != nil {
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

// ParseLogBindRequest is a log parse operation binding the contract event 0xf8608cf3c27350e6aa0beaa6723ea6911e3d7353e8b22a69bb112c15f93867ca.
//
// Solidity: event LogBindRequest(address contractAddr, bytes32 bep2TokenSymbol, uint256 totalSupply, uint256 peggyAmount)
func (_TokenHub *TokenHubFilterer) ParseLogBindRequest(log types.Log) (*TokenHubLogBindRequest, error) {
	event := new(TokenHubLogBindRequest)
	if err := _TokenHub.contract.UnpackLog(event, "LogBindRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogBindSuccessIterator is returned from FilterLogBindSuccess and is used to iterate over the raw logs and unpacked data for LogBindSuccess events raised by the TokenHub contract.
type TokenHubLogBindSuccessIterator struct {
	Event *TokenHubLogBindSuccess // Event containing the contract specifics and raw log

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
func (it *TokenHubLogBindSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogBindSuccess)
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
		it.Event = new(TokenHubLogBindSuccess)
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
func (it *TokenHubLogBindSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogBindSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogBindSuccess represents a LogBindSuccess event raised by the TokenHub contract.
type TokenHubLogBindSuccess struct {
	Sequence        *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	TotalSupply     *big.Int
	PeggyAmount     *big.Int
	Decimals        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogBindSuccess is a free log retrieval operation binding the contract event 0x8005b9354dd0ca4c5593805bcd00ea12b5fce8a2cc9bc15252f50fb2d17c09d2.
//
// Solidity: event LogBindSuccess(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol, uint256 totalSupply, uint256 peggyAmount, uint256 decimals)
func (_TokenHub *TokenHubFilterer) FilterLogBindSuccess(opts *bind.FilterOpts) (*TokenHubLogBindSuccessIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogBindSuccess")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogBindSuccessIterator{contract: _TokenHub.contract, event: "LogBindSuccess", logs: logs, sub: sub}, nil
}

// WatchLogBindSuccess is a free log subscription operation binding the contract event 0x8005b9354dd0ca4c5593805bcd00ea12b5fce8a2cc9bc15252f50fb2d17c09d2.
//
// Solidity: event LogBindSuccess(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol, uint256 totalSupply, uint256 peggyAmount, uint256 decimals)
func (_TokenHub *TokenHubFilterer) WatchLogBindSuccess(opts *bind.WatchOpts, sink chan<- *TokenHubLogBindSuccess) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogBindSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogBindSuccess)
				if err := _TokenHub.contract.UnpackLog(event, "LogBindSuccess", log); err != nil {
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

// ParseLogBindSuccess is a log parse operation binding the contract event 0x8005b9354dd0ca4c5593805bcd00ea12b5fce8a2cc9bc15252f50fb2d17c09d2.
//
// Solidity: event LogBindSuccess(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol, uint256 totalSupply, uint256 peggyAmount, uint256 decimals)
func (_TokenHub *TokenHubFilterer) ParseLogBindSuccess(log types.Log) (*TokenHubLogBindSuccess, error) {
	event := new(TokenHubLogBindSuccess)
	if err := _TokenHub.contract.UnpackLog(event, "LogBindSuccess", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogBindTimeoutIterator is returned from FilterLogBindTimeout and is used to iterate over the raw logs and unpacked data for LogBindTimeout events raised by the TokenHub contract.
type TokenHubLogBindTimeoutIterator struct {
	Event *TokenHubLogBindTimeout // Event containing the contract specifics and raw log

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
func (it *TokenHubLogBindTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogBindTimeout)
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
		it.Event = new(TokenHubLogBindTimeout)
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
func (it *TokenHubLogBindTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogBindTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogBindTimeout represents a LogBindTimeout event raised by the TokenHub contract.
type TokenHubLogBindTimeout struct {
	Sequence        *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogBindTimeout is a free log retrieval operation binding the contract event 0x4781c2d0a33124fb32083581f5b48c93a59b71fd567ce2d4a56c89196baa2ccd.
//
// Solidity: event LogBindTimeout(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) FilterLogBindTimeout(opts *bind.FilterOpts) (*TokenHubLogBindTimeoutIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogBindTimeout")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogBindTimeoutIterator{contract: _TokenHub.contract, event: "LogBindTimeout", logs: logs, sub: sub}, nil
}

// WatchLogBindTimeout is a free log subscription operation binding the contract event 0x4781c2d0a33124fb32083581f5b48c93a59b71fd567ce2d4a56c89196baa2ccd.
//
// Solidity: event LogBindTimeout(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) WatchLogBindTimeout(opts *bind.WatchOpts, sink chan<- *TokenHubLogBindTimeout) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogBindTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogBindTimeout)
				if err := _TokenHub.contract.UnpackLog(event, "LogBindTimeout", log); err != nil {
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

// ParseLogBindTimeout is a log parse operation binding the contract event 0x4781c2d0a33124fb32083581f5b48c93a59b71fd567ce2d4a56c89196baa2ccd.
//
// Solidity: event LogBindTimeout(uint256 sequence, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) ParseLogBindTimeout(log types.Log) (*TokenHubLogBindTimeout, error) {
	event := new(TokenHubLogBindTimeout)
	if err := _TokenHub.contract.UnpackLog(event, "LogBindTimeout", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogTransferInFailureInsufficientBalanceIterator is returned from FilterLogTransferInFailureInsufficientBalance and is used to iterate over the raw logs and unpacked data for LogTransferInFailureInsufficientBalance events raised by the TokenHub contract.
type TokenHubLogTransferInFailureInsufficientBalanceIterator struct {
	Event *TokenHubLogTransferInFailureInsufficientBalance // Event containing the contract specifics and raw log

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
func (it *TokenHubLogTransferInFailureInsufficientBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogTransferInFailureInsufficientBalance)
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
		it.Event = new(TokenHubLogTransferInFailureInsufficientBalance)
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
func (it *TokenHubLogTransferInFailureInsufficientBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogTransferInFailureInsufficientBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogTransferInFailureInsufficientBalance represents a LogTransferInFailureInsufficientBalance event raised by the TokenHub contract.
type TokenHubLogTransferInFailureInsufficientBalance struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	AuctualBalance  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogTransferInFailureInsufficientBalance is a free log retrieval operation binding the contract event 0x1de400dfa3e72ba83f12c6f1d8b9b85dc3d2aedc6eacc27b481267826aec7422.
//
// Solidity: event LogTransferInFailureInsufficientBalance(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 auctualBalance)
func (_TokenHub *TokenHubFilterer) FilterLogTransferInFailureInsufficientBalance(opts *bind.FilterOpts) (*TokenHubLogTransferInFailureInsufficientBalanceIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogTransferInFailureInsufficientBalance")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogTransferInFailureInsufficientBalanceIterator{contract: _TokenHub.contract, event: "LogTransferInFailureInsufficientBalance", logs: logs, sub: sub}, nil
}

// WatchLogTransferInFailureInsufficientBalance is a free log subscription operation binding the contract event 0x1de400dfa3e72ba83f12c6f1d8b9b85dc3d2aedc6eacc27b481267826aec7422.
//
// Solidity: event LogTransferInFailureInsufficientBalance(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 auctualBalance)
func (_TokenHub *TokenHubFilterer) WatchLogTransferInFailureInsufficientBalance(opts *bind.WatchOpts, sink chan<- *TokenHubLogTransferInFailureInsufficientBalance) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogTransferInFailureInsufficientBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogTransferInFailureInsufficientBalance)
				if err := _TokenHub.contract.UnpackLog(event, "LogTransferInFailureInsufficientBalance", log); err != nil {
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

// ParseLogTransferInFailureInsufficientBalance is a log parse operation binding the contract event 0x1de400dfa3e72ba83f12c6f1d8b9b85dc3d2aedc6eacc27b481267826aec7422.
//
// Solidity: event LogTransferInFailureInsufficientBalance(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 auctualBalance)
func (_TokenHub *TokenHubFilterer) ParseLogTransferInFailureInsufficientBalance(log types.Log) (*TokenHubLogTransferInFailureInsufficientBalance, error) {
	event := new(TokenHubLogTransferInFailureInsufficientBalance)
	if err := _TokenHub.contract.UnpackLog(event, "LogTransferInFailureInsufficientBalance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogTransferInFailureTimeoutIterator is returned from FilterLogTransferInFailureTimeout and is used to iterate over the raw logs and unpacked data for LogTransferInFailureTimeout events raised by the TokenHub contract.
type TokenHubLogTransferInFailureTimeoutIterator struct {
	Event *TokenHubLogTransferInFailureTimeout // Event containing the contract specifics and raw log

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
func (it *TokenHubLogTransferInFailureTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogTransferInFailureTimeout)
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
		it.Event = new(TokenHubLogTransferInFailureTimeout)
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
func (it *TokenHubLogTransferInFailureTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogTransferInFailureTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogTransferInFailureTimeout represents a LogTransferInFailureTimeout event raised by the TokenHub contract.
type TokenHubLogTransferInFailureTimeout struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	ExpireTime      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogTransferInFailureTimeout is a free log retrieval operation binding the contract event 0x8090e98e190cb0b05412d5c1a8cd5ee9af5d40da935335cef5d4179c7da63d79.
//
// Solidity: event LogTransferInFailureTimeout(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime)
func (_TokenHub *TokenHubFilterer) FilterLogTransferInFailureTimeout(opts *bind.FilterOpts) (*TokenHubLogTransferInFailureTimeoutIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogTransferInFailureTimeout")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogTransferInFailureTimeoutIterator{contract: _TokenHub.contract, event: "LogTransferInFailureTimeout", logs: logs, sub: sub}, nil
}

// WatchLogTransferInFailureTimeout is a free log subscription operation binding the contract event 0x8090e98e190cb0b05412d5c1a8cd5ee9af5d40da935335cef5d4179c7da63d79.
//
// Solidity: event LogTransferInFailureTimeout(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime)
func (_TokenHub *TokenHubFilterer) WatchLogTransferInFailureTimeout(opts *bind.WatchOpts, sink chan<- *TokenHubLogTransferInFailureTimeout) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogTransferInFailureTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogTransferInFailureTimeout)
				if err := _TokenHub.contract.UnpackLog(event, "LogTransferInFailureTimeout", log); err != nil {
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

// ParseLogTransferInFailureTimeout is a log parse operation binding the contract event 0x8090e98e190cb0b05412d5c1a8cd5ee9af5d40da935335cef5d4179c7da63d79.
//
// Solidity: event LogTransferInFailureTimeout(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime)
func (_TokenHub *TokenHubFilterer) ParseLogTransferInFailureTimeout(log types.Log) (*TokenHubLogTransferInFailureTimeout, error) {
	event := new(TokenHubLogTransferInFailureTimeout)
	if err := _TokenHub.contract.UnpackLog(event, "LogTransferInFailureTimeout", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogTransferInFailureUnboundTokenIterator is returned from FilterLogTransferInFailureUnboundToken and is used to iterate over the raw logs and unpacked data for LogTransferInFailureUnboundToken events raised by the TokenHub contract.
type TokenHubLogTransferInFailureUnboundTokenIterator struct {
	Event *TokenHubLogTransferInFailureUnboundToken // Event containing the contract specifics and raw log

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
func (it *TokenHubLogTransferInFailureUnboundTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogTransferInFailureUnboundToken)
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
		it.Event = new(TokenHubLogTransferInFailureUnboundToken)
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
func (it *TokenHubLogTransferInFailureUnboundTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogTransferInFailureUnboundTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogTransferInFailureUnboundToken represents a LogTransferInFailureUnboundToken event raised by the TokenHub contract.
type TokenHubLogTransferInFailureUnboundToken struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogTransferInFailureUnboundToken is a free log retrieval operation binding the contract event 0x055f2adbd109a4e99b3821af55571cccb4981551d10e3846b21574d348572a59.
//
// Solidity: event LogTransferInFailureUnboundToken(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) FilterLogTransferInFailureUnboundToken(opts *bind.FilterOpts) (*TokenHubLogTransferInFailureUnboundTokenIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogTransferInFailureUnboundToken")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogTransferInFailureUnboundTokenIterator{contract: _TokenHub.contract, event: "LogTransferInFailureUnboundToken", logs: logs, sub: sub}, nil
}

// WatchLogTransferInFailureUnboundToken is a free log subscription operation binding the contract event 0x055f2adbd109a4e99b3821af55571cccb4981551d10e3846b21574d348572a59.
//
// Solidity: event LogTransferInFailureUnboundToken(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) WatchLogTransferInFailureUnboundToken(opts *bind.WatchOpts, sink chan<- *TokenHubLogTransferInFailureUnboundToken) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogTransferInFailureUnboundToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogTransferInFailureUnboundToken)
				if err := _TokenHub.contract.UnpackLog(event, "LogTransferInFailureUnboundToken", log); err != nil {
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

// ParseLogTransferInFailureUnboundToken is a log parse operation binding the contract event 0x055f2adbd109a4e99b3821af55571cccb4981551d10e3846b21574d348572a59.
//
// Solidity: event LogTransferInFailureUnboundToken(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) ParseLogTransferInFailureUnboundToken(log types.Log) (*TokenHubLogTransferInFailureUnboundToken, error) {
	event := new(TokenHubLogTransferInFailureUnboundToken)
	if err := _TokenHub.contract.UnpackLog(event, "LogTransferInFailureUnboundToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogTransferInFailureUnknownReasonIterator is returned from FilterLogTransferInFailureUnknownReason and is used to iterate over the raw logs and unpacked data for LogTransferInFailureUnknownReason events raised by the TokenHub contract.
type TokenHubLogTransferInFailureUnknownReasonIterator struct {
	Event *TokenHubLogTransferInFailureUnknownReason // Event containing the contract specifics and raw log

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
func (it *TokenHubLogTransferInFailureUnknownReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogTransferInFailureUnknownReason)
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
		it.Event = new(TokenHubLogTransferInFailureUnknownReason)
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
func (it *TokenHubLogTransferInFailureUnknownReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogTransferInFailureUnknownReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogTransferInFailureUnknownReason represents a LogTransferInFailureUnknownReason event raised by the TokenHub contract.
type TokenHubLogTransferInFailureUnknownReason struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Bep2TokenAmount *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogTransferInFailureUnknownReason is a free log retrieval operation binding the contract event 0xcb6ddd4a252f58c1ff32f31fbb529dc35e8f6a81908f6211bbe7dfa94ef52f1f.
//
// Solidity: event LogTransferInFailureUnknownReason(uint256 sequence, address refundAddr, address recipient, uint256 bep2TokenAmount, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) FilterLogTransferInFailureUnknownReason(opts *bind.FilterOpts) (*TokenHubLogTransferInFailureUnknownReasonIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogTransferInFailureUnknownReason")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogTransferInFailureUnknownReasonIterator{contract: _TokenHub.contract, event: "LogTransferInFailureUnknownReason", logs: logs, sub: sub}, nil
}

// WatchLogTransferInFailureUnknownReason is a free log subscription operation binding the contract event 0xcb6ddd4a252f58c1ff32f31fbb529dc35e8f6a81908f6211bbe7dfa94ef52f1f.
//
// Solidity: event LogTransferInFailureUnknownReason(uint256 sequence, address refundAddr, address recipient, uint256 bep2TokenAmount, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) WatchLogTransferInFailureUnknownReason(opts *bind.WatchOpts, sink chan<- *TokenHubLogTransferInFailureUnknownReason) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogTransferInFailureUnknownReason")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogTransferInFailureUnknownReason)
				if err := _TokenHub.contract.UnpackLog(event, "LogTransferInFailureUnknownReason", log); err != nil {
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

// ParseLogTransferInFailureUnknownReason is a log parse operation binding the contract event 0xcb6ddd4a252f58c1ff32f31fbb529dc35e8f6a81908f6211bbe7dfa94ef52f1f.
//
// Solidity: event LogTransferInFailureUnknownReason(uint256 sequence, address refundAddr, address recipient, uint256 bep2TokenAmount, address contractAddr, bytes32 bep2TokenSymbol)
func (_TokenHub *TokenHubFilterer) ParseLogTransferInFailureUnknownReason(log types.Log) (*TokenHubLogTransferInFailureUnknownReason, error) {
	event := new(TokenHubLogTransferInFailureUnknownReason)
	if err := _TokenHub.contract.UnpackLog(event, "LogTransferInFailureUnknownReason", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenHubLogTransferOutIterator is returned from FilterLogTransferOut and is used to iterate over the raw logs and unpacked data for LogTransferOut events raised by the TokenHub contract.
type TokenHubLogTransferOutIterator struct {
	Event *TokenHubLogTransferOut // Event containing the contract specifics and raw log

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
func (it *TokenHubLogTransferOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenHubLogTransferOut)
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
		it.Event = new(TokenHubLogTransferOut)
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
func (it *TokenHubLogTransferOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenHubLogTransferOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenHubLogTransferOut represents a LogTransferOut event raised by the TokenHub contract.
type TokenHubLogTransferOut struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol [32]byte
	ExpireTime      *big.Int
	RelayFee        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogTransferOut is a free log retrieval operation binding the contract event 0x5bd451c53ab05abd9855ceb52a469590655af1d732a4cfd67f1f9b53d74dc613.
//
// Solidity: event LogTransferOut(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime, uint256 relayFee)
func (_TokenHub *TokenHubFilterer) FilterLogTransferOut(opts *bind.FilterOpts) (*TokenHubLogTransferOutIterator, error) {

	logs, sub, err := _TokenHub.contract.FilterLogs(opts, "LogTransferOut")
	if err != nil {
		return nil, err
	}
	return &TokenHubLogTransferOutIterator{contract: _TokenHub.contract, event: "LogTransferOut", logs: logs, sub: sub}, nil
}

// WatchLogTransferOut is a free log subscription operation binding the contract event 0x5bd451c53ab05abd9855ceb52a469590655af1d732a4cfd67f1f9b53d74dc613.
//
// Solidity: event LogTransferOut(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime, uint256 relayFee)
func (_TokenHub *TokenHubFilterer) WatchLogTransferOut(opts *bind.WatchOpts, sink chan<- *TokenHubLogTransferOut) (event.Subscription, error) {

	logs, sub, err := _TokenHub.contract.WatchLogs(opts, "LogTransferOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenHubLogTransferOut)
				if err := _TokenHub.contract.UnpackLog(event, "LogTransferOut", log); err != nil {
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

// ParseLogTransferOut is a log parse operation binding the contract event 0x5bd451c53ab05abd9855ceb52a469590655af1d732a4cfd67f1f9b53d74dc613.
//
// Solidity: event LogTransferOut(uint256 sequence, address refundAddr, address recipient, uint256 amount, address contractAddr, bytes32 bep2TokenSymbol, uint256 expireTime, uint256 relayFee)
func (_TokenHub *TokenHubFilterer) ParseLogTransferOut(log types.Log) (*TokenHubLogTransferOut, error) {
	event := new(TokenHubLogTransferOut)
	if err := _TokenHub.contract.UnpackLog(event, "LogTransferOut", log); err != nil {
		return nil, err
	}
	return event, nil
}
