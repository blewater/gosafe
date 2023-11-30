// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"approvedHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ApproveHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"ChangedFallbackHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"ChangedGuard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ChangedThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DisabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"EnabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemovedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"}],\"name\":\"SafeSetup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"SignMsg\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"addOwnerWithThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashToApprove\",\"type\":\"bytes32\"}],\"name\":\"approveHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignatures\",\"type\":\"uint256\"}],\"name\":\"checkNSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"encodeTransactionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModuleReturnData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getModulesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isModuleEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"requiredTxGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"setFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"setGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"signedMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"swapOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506001600481905550615e5780620000265f395ff3fe6080604052600436106101db575f3560e01c8063affed0e011610101578063e19a9dd911610094578063f08a032311610063578063f08a0323146107b9578063f698da25146107e1578063f8dc5dd91461080b578063ffa1ad741461083357610230565b8063e19a9dd914610703578063e318b52b1461072b578063e75235b814610753578063e86637db1461077d57610230565b8063cc2f8452116100d0578063cc2f84521461063a578063d4d9bdcd14610677578063d8d11f781461069f578063e009cfde146106db57610230565b8063affed0e014610584578063b4faba09146105ae578063b63e800d146105d6578063c4ca3a9c146105fe57610230565b80635624b25b116101795780636a761202116101485780636a761202146104c65780637d832974146104f6578063934f3a1114610532578063a0e67e2b1461055a57610230565b80635624b25b146103fe5780635ae6bd371461043a578063610b592514610476578063694e80c31461049e57610230565b80632f54bf6e116101b55780632f54bf6e1461031f5780633408e4701461035b578063468721a7146103855780635229073f146103c157610230565b80630d582f131461029357806312fb68e0146102bb5780632d9ad53d146102e357610230565b36610230573373ffffffffffffffffffffffffffffffffffffffff167f3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d346040516102269190613d5d565b60405180910390a2005b34801561023b575f80fd5b505f7f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d55f1b905080548061026d575f80f35b365f80373360601b36525f80601436015f80855af13d5f803e8061028f573d5ffd5b3d5ff35b34801561029e575f80fd5b506102b960048036038101906102b49190613e0b565b61085d565b005b3480156102c6575f80fd5b506102e160048036038101906102dc9190613fb8565b610bc8565b005b3480156102ee575f80fd5b5061030960048036038101906103049190614054565b6111a3565b6040516103169190614099565b60405180910390f35b34801561032a575f80fd5b5061034560048036038101906103409190614054565b611270565b6040516103529190614099565b60405180910390f35b348015610366575f80fd5b5061036f61133d565b60405161037c9190613d5d565b60405180910390f35b348015610390575f80fd5b506103ab60048036038101906103a691906140d5565b611349565b6040516103b89190614099565b60405180910390f35b3480156103cc575f80fd5b506103e760048036038101906103e291906140d5565b6114f7565b6040516103f59291906141cf565b60405180910390f35b348015610409575f80fd5b50610424600480360381019061041f91906141fd565b61152b565b604051610431919061423b565b60405180910390f35b348015610445575f80fd5b50610460600480360381019061045b919061425b565b6115c1565b60405161046d9190613d5d565b60405180910390f35b348015610481575f80fd5b5061049c60048036038101906104979190614054565b6115d6565b005b3480156104a9575f80fd5b506104c460048036038101906104bf9190614286565b6118de565b005b6104e060048036038101906104db9190614349565b6119b2565b6040516104ed9190614099565b60405180910390f35b348015610501575f80fd5b5061051c6004803603810190610517919061445e565b611d5d565b6040516105299190613d5d565b60405180910390f35b34801561053d575f80fd5b506105586004803603810190610553919061449c565b611d7d565b005b348015610565575f80fd5b5061056e611dd7565b60405161057b91906145db565b60405180910390f35b34801561058f575f80fd5b50610598611f87565b6040516105a59190613d5d565b60405180910390f35b3480156105b9575f80fd5b506105d460048036038101906105cf91906145fb565b611f8d565b005b3480156105e1575f80fd5b506105fc60048036038101906105f791906146aa565b611fab565b005b348015610609575f80fd5b50610624600480360381019061061f919061479b565b6120f9565b6040516106319190613d5d565b60405180910390f35b348015610645575f80fd5b50610660600480360381019061065b9190613e0b565b6121c1565b60405161066e92919061482e565b60405180910390f35b348015610682575f80fd5b5061069d6004803603810190610698919061425b565b6123b9565b005b3480156106aa575f80fd5b506106c560048036038101906106c0919061485c565b61251c565b6040516106d29190614964565b60405180910390f35b3480156106e6575f80fd5b5061070160048036038101906106fc919061497d565b612548565b005b34801561070e575f80fd5b5061072960048036038101906107249190614054565b61284f565b005b348015610736575f80fd5b50610751600480360381019061074c91906149bb565b6128bb565b005b34801561075e575f80fd5b50610767612e1f565b6040516107749190613d5d565b60405180910390f35b348015610788575f80fd5b506107a3600480360381019061079e919061485c565b612e28565b6040516107b0919061423b565b60405180910390f35b3480156107c4575f80fd5b506107df60048036038101906107da9190614054565b612ee8565b005b3480156107ec575f80fd5b506107f5612f33565b6040516108029190614964565b60405180910390f35b348015610816575f80fd5b50610831600480360381019061082c9190614a0b565b612f8d565b005b34801561083e575f80fd5b50610847613310565b6040516108549190614aad565b60405180910390f35b610865613349565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141580156108ce5750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b801561090657503073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b610945576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093c90614b17565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0690614b7f565b60405180910390fd5b60025f600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160025f600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060035f815480929190610b7590614bca565b91905055507f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea2682604051610ba99190614c11565b60405180910390a18060045414610bc457610bc3816118de565b5b5050565b610bdc6041826133b990919063ffffffff16565b82511015610c1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1690614c74565b60405180910390fd5b5f80805f805f5b8681101561119757610c3888826133f9565b8094508195508296505050505f8460ff1603610e8a57825f1c9450610c676041886133b990919063ffffffff16565b825f1c1015610cab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ca290614cdc565b60405180910390fd5b8751610cc36020845f1c61342690919063ffffffff16565b1115610d04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cfb90614d44565b60405180910390fd5b5f6020838a01015190508851610d3882610d2a6020875f1c61342690919063ffffffff16565b61342690919063ffffffff16565b1115610d79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7090614dac565b60405180910390fd5b60606020848b010190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168773ffffffffffffffffffffffffffffffffffffffff166320c13b0b8d846040518363ffffffff1660e01b8152600401610de5929190614dca565b602060405180830381865afa158015610e00573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e249190614e54565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610e83576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7a90614ec9565b60405180910390fd5b5050611046565b60018460ff1603610f6457825f1c94508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610f2057505f60085f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8c81526020019081526020015f205414155b610f5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5690614f31565b60405180910390fd5b611045565b601e8460ff161115610ff55760018a604051602001610f839190614fc3565b60405160208183030381529060405280519060200120600486610fa69190614ff4565b85856040515f8152602001604052604051610fc49493929190615037565b6020604051602081039080840390855afa158015610fe4573d5f803e3d5ffd5b505050602060405103519450611044565b60018a8585856040515f81526020016040526040516110179493929190615037565b6020604051602081039080840390855afa158015611037573d5f803e3d5ffd5b5050506020604051035194505b5b5b8573ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1611801561110957505f73ffffffffffffffffffffffffffffffffffffffff1660025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b80156111425750600173ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614155b611181576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611178906150c4565b60405180910390fd5b849550808061118f90614bca565b915050610c26565b50505050505050505050565b5f8173ffffffffffffffffffffffffffffffffffffffff16600173ffffffffffffffffffffffffffffffffffffffff161415801561126957505f73ffffffffffffffffffffffffffffffffffffffff1660015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b9050919050565b5f600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415801561133657505f73ffffffffffffffffffffffffffffffffffffffff1660025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b9050919050565b5f804690508091505090565b5f600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415801561140f57505f73ffffffffffffffffffffffffffffffffffffffff1660015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b61144e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114459061512c565b60405180910390fd5b61145b858585855a61344c565b905080156114ab573373ffffffffffffffffffffffffffffffffffffffff167f6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb860405160405180910390a26114ef565b3373ffffffffffffffffffffffffffffffffffffffff167facd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd37560405160405180910390a25b949350505050565b5f606061150686868686611349565b915060405160203d0181016040523d81523d5f602083013e8091505094509492505050565b60605f60208361153b919061514a565b67ffffffffffffffff81111561155457611553613e94565b5b6040519080825280601f01601f1916602001820160405280156115865781602001600182028036833780820191505090505b5090505f5b838110156115b6578085015480602083026020850101525080806115ae90614bca565b91505061158b565b508091505092915050565b6007602052805f5260405f205f915090505481565b6115de613349565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141580156116475750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b611686576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161167d906151d5565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611750576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117479061523d565b60405180910390fd5b60015f600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060015f600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440816040516118d39190614c11565b60405180910390a150565b6118e6613349565b60035481111561192b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611922906152a5565b60405180910390fd5b600181101561196f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119669061530d565b60405180910390fd5b806004819055507f610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c936004546040516119a79190613d5d565b60405180910390a150565b5f805f6119ca8e8e8e8e8e8e8e8e8e8e600554612e28565b905060055f8154809291906119de90614bca565b9190505550808051906020012091506119f8828286611d7d565b505f611a026134a2565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611ab4578073ffffffffffffffffffffffffffffffffffffffff166375f0bb528f8f8f8f8f8f8f8f8f8f8f336040518d63ffffffff1660e01b8152600401611a869c9b9a999897969594939291906153d9565b5f604051808303815f87803b158015611a9d575f80fd5b505af1158015611aaf573d5f803e3d5ffd5b505050505b6101f4611aef6109c48b611ac89190615492565b603f60408d611ad7919061514a565b611ae191906154f2565b6134d190919063ffffffff16565b611af99190615492565b5a1015611b3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b329061556c565b60405180910390fd5b5f5a9050611baa8f8f8f8f8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508e5f8d14611b96578e611ba5565b6109c45a611ba4919061558a565b5b61344c565b9350611bbf5a826134ea90919063ffffffff16565b90508380611bcd57505f8a14155b80611bd857505f8814155b611c17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c0e90615607565b60405180910390fd5b5f80891115611c3057611c2d828b8b8b8b613510565b90505b8415611c74577f442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e8482604051611c67929190615625565b60405180910390a1611cae565b7f23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d238482604051611ca5929190615625565b60405180910390a15b50505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611d4c578073ffffffffffffffffffffffffffffffffffffffff16639327136883856040518363ffffffff1660e01b8152600401611d1e92919061564c565b5f604051808303815f87803b158015611d35575f80fd5b505af1158015611d47573d5f803e3d5ffd5b505050505b50509b9a5050505050505050505050565b6008602052815f5260405f20602052805f5260405f205f91509150505481565b5f60045490505f8111611dc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dbc906156bd565b60405180910390fd5b611dd184848484610bc8565b50505050565b60605f60035467ffffffffffffffff811115611df657611df5613e94565b5b604051908082528060200260200182016040528015611e245781602001602082028036833780820191505090505b5090505f8060025f600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611f7e5780838381518110611ed257611ed16156db565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060025f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508180611f7690614bca565b925050611e8a565b82935050505090565b60055481565b5f80825160208401855af4805f523d6020523d5f60403e60403d015ffd5b611ff58a8a808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f82011690508083019250505050505050896136ab565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16146120325761203184613aaa565b5b61207f8787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050613ad7565b5f82111561209757612095825f60018685613510565b505b3373ffffffffffffffffffffffffffffffffffffffff167f141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a88b8b8b8b896040516120e595949392919061578e565b60405180910390a250505050505050505050565b5f805a905061214e878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050865a61344c565b612156575f80fd5b5f5a82612163919061558a565b90508060405160200161217691906157fa565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121b89190614aad565b60405180910390fd5b60605f8267ffffffffffffffff8111156121de576121dd613e94565b5b60405190808252806020026020018201604052801561220c5781602001602082028036833780820191505090505b5091505f8060015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141580156122da5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156122e557508482105b156123aa57808483815181106122fe576122fd6156db565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081806123a290614bca565b925050612271565b80925081845250509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff1660025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612483576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161247a9061585e565b60405180910390fd5b600160085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff16817ff2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c60405160405180910390a350565b5f6125308c8c8c8c8c8c8c8c8c8c8c612e28565b8051906020012090509b9a5050505050505050505050565b612550613349565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141580156125b95750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b6125f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125ef906151d5565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146126c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126b9906158c6565b60405180910390fd5b60015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507faab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276816040516128439190614c11565b60405180910390a15050565b612857613349565b5f7f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c85f1b90508181557f1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2826040516128af9190614c11565b60405180910390a15050565b6128c3613349565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561292c5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561296457503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b6129a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161299a90614b17565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612a6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a6490614b7f565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015612ad65750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b612b15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b0c90614b17565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612bdf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bd69061592e565b60405180910390fd5b60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf82604051612ddb9190614c11565b60405180910390a17f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea2681604051612e129190614c11565b60405180910390a1505050565b5f600454905090565b60605f7fbb8310d486368db6bd6f849402fdd73ad53d316b5a4b2644ad6efe0f941286d85f1b8d8d8d8d604051612e6092919061597a565b60405180910390208c8c8c8c8c8c8c604051602001612e899b9a99989796959493929190615992565b604051602081830303815290604052805190602001209050601960f81b600160f81b612eb3612f33565b83604051602001612ec79493929190615a86565b6040516020818303038152906040529150509b9a5050505050505050505050565b612ef0613349565b612ef981613aaa565b7f5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b081604051612f289190614c11565b60405180910390a150565b5f7f47e79534a245952e8b16893a336b85a3d9ea9fa8c573f3d803afb92a794692185f1b612f5f61133d565b30604051602001612f7293929190615b2e565b60405160208183030381529060405280519060200120905090565b612f95613349565b806001600354612fa5919061558a565b1015612fe6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612fdd906152a5565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415801561304f5750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b61308e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161308590614b17565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613158576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161314f9061592e565b60405180910390fd5b60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060035f8154809291906132bc90615b63565b91905055507ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf826040516132f09190614c11565b60405180910390a1806004541461330b5761330a816118de565b5b505050565b6040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146133b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133ae90615bd4565b60405180910390fd5b565b5f8083036133c9575f90506133f3565b5f82846133d6919061514a565b90508284826133e591906154f2565b146133ee575f80fd5b809150505b92915050565b5f805f8360410260208101860151925060408101860151915060ff60418201870151169350509250925092565b5f8082846134349190615492565b905083811015613442575f80fd5b8091505092915050565b5f6001808111156134605761345f615357565b5b83600181111561347357613472615357565b5b0361348a575f808551602087018986f49050613499565b5f80855160208701888a87f190505b95945050505050565b5f807f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c85f1b9050805491505090565b5f818310156134e057816134e2565b825b905092915050565b5f828211156134f7575f80fd5b5f8284613504919061558a565b90508091505092915050565b5f805f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161461354b578261354d565b325b90505f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff160361362f576135b53a8610613592573a613594565b855b6135a7888a61342690919063ffffffff16565b6133b990919063ffffffff16565b91508073ffffffffffffffffffffffffffffffffffffffff166108fc8390811502906040515f60405180830381858888f1935050505061362a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161362190615c3c565b60405180910390fd5b6136a1565b61365485613646888a61342690919063ffffffff16565b6133b990919063ffffffff16565b9150613661848284613ca3565b6136a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161369790615ca4565b60405180910390fd5b5b5095945050505050565b5f600454146136ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136e690615d0c565b60405180910390fd5b8151811115613733576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161372a906152a5565b60405180910390fd5b6001811015613777576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161376e9061530d565b60405180910390fd5b5f600190505f5b8351811015613a19575f84828151811061379b5761379a6156db565b5b602002602001015190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561380e5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561384657503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561387e57508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614155b6138bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138b490614b17565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613987576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161397e90614b7f565b60405180910390fd5b8060025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550809250508080613a1190614bca565b91505061377e565b50600160025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550825160038190555081600481905550505050565b5f7f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d55f1b90508181555050565b5f73ffffffffffffffffffffffffffffffffffffffff1660015f600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613ba2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b9990615d74565b60405180910390fd5b6001805f600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614613c9f57613c5f825f8360015a61344c565b613c9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c9590615ddc565b60405180910390fd5b5b5050565b5f8063a9059cbb8484604051602401613cbd929190615dfa565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060205f8251602084015f896127105a03f13d5f8114613d295760208114613d31575f9350613d3b565b819350613d3b565b5f51158215171593505b5050509392505050565b5f819050919050565b613d5781613d45565b82525050565b5f602082019050613d705f830184613d4e565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f613db082613d87565b9050919050565b613dc081613da6565b8114613dca575f80fd5b50565b5f81359050613ddb81613db7565b92915050565b613dea81613d45565b8114613df4575f80fd5b50565b5f81359050613e0581613de1565b92915050565b5f8060408385031215613e2157613e20613d7f565b5b5f613e2e85828601613dcd565b9250506020613e3f85828601613df7565b9150509250929050565b5f819050919050565b613e5b81613e49565b8114613e65575f80fd5b50565b5f81359050613e7681613e52565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b613eca82613e84565b810181811067ffffffffffffffff82111715613ee957613ee8613e94565b5b80604052505050565b5f613efb613d76565b9050613f078282613ec1565b919050565b5f67ffffffffffffffff821115613f2657613f25613e94565b5b613f2f82613e84565b9050602081019050919050565b828183375f83830152505050565b5f613f5c613f5784613f0c565b613ef2565b905082815260208101848484011115613f7857613f77613e80565b5b613f83848285613f3c565b509392505050565b5f82601f830112613f9f57613f9e613e7c565b5b8135613faf848260208601613f4a565b91505092915050565b5f805f8060808587031215613fd057613fcf613d7f565b5b5f613fdd87828801613e68565b945050602085013567ffffffffffffffff811115613ffe57613ffd613d83565b5b61400a87828801613f8b565b935050604085013567ffffffffffffffff81111561402b5761402a613d83565b5b61403787828801613f8b565b925050606061404887828801613df7565b91505092959194509250565b5f6020828403121561406957614068613d7f565b5b5f61407684828501613dcd565b91505092915050565b5f8115159050919050565b6140938161407f565b82525050565b5f6020820190506140ac5f83018461408a565b92915050565b600281106140be575f80fd5b50565b5f813590506140cf816140b2565b92915050565b5f805f80608085870312156140ed576140ec613d7f565b5b5f6140fa87828801613dcd565b945050602061410b87828801613df7565b935050604085013567ffffffffffffffff81111561412c5761412b613d83565b5b61413887828801613f8b565b9250506060614149878288016140c1565b91505092959194509250565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561418c578082015181840152602081019050614171565b5f8484015250505050565b5f6141a182614155565b6141ab818561415f565b93506141bb81856020860161416f565b6141c481613e84565b840191505092915050565b5f6040820190506141e25f83018561408a565b81810360208301526141f48184614197565b90509392505050565b5f806040838503121561421357614212613d7f565b5b5f61422085828601613df7565b925050602061423185828601613df7565b9150509250929050565b5f6020820190508181035f8301526142538184614197565b905092915050565b5f602082840312156142705761426f613d7f565b5b5f61427d84828501613e68565b91505092915050565b5f6020828403121561429b5761429a613d7f565b5b5f6142a884828501613df7565b91505092915050565b5f80fd5b5f80fd5b5f8083601f8401126142ce576142cd613e7c565b5b8235905067ffffffffffffffff8111156142eb576142ea6142b1565b5b602083019150836001820283011115614307576143066142b5565b5b9250929050565b5f61431882613d87565b9050919050565b6143288161430e565b8114614332575f80fd5b50565b5f813590506143438161431f565b92915050565b5f805f805f805f805f805f6101408c8e03121561436957614368613d7f565b5b5f6143768e828f01613dcd565b9b505060206143878e828f01613df7565b9a505060408c013567ffffffffffffffff8111156143a8576143a7613d83565b5b6143b48e828f016142b9565b995099505060606143c78e828f016140c1565b97505060806143d88e828f01613df7565b96505060a06143e98e828f01613df7565b95505060c06143fa8e828f01613df7565b94505060e061440b8e828f01613dcd565b93505061010061441d8e828f01614335565b9250506101208c013567ffffffffffffffff81111561443f5761443e613d83565b5b61444b8e828f01613f8b565b9150509295989b509295989b9093969950565b5f806040838503121561447457614473613d7f565b5b5f61448185828601613dcd565b925050602061449285828601613e68565b9150509250929050565b5f805f606084860312156144b3576144b2613d7f565b5b5f6144c086828701613e68565b935050602084013567ffffffffffffffff8111156144e1576144e0613d83565b5b6144ed86828701613f8b565b925050604084013567ffffffffffffffff81111561450e5761450d613d83565b5b61451a86828701613f8b565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61455681613da6565b82525050565b5f614567838361454d565b60208301905092915050565b5f602082019050919050565b5f61458982614524565b614593818561452e565b935061459e8361453e565b805f5b838110156145ce5781516145b5888261455c565b97506145c083614573565b9250506001810190506145a1565b5085935050505092915050565b5f6020820190508181035f8301526145f3818461457f565b905092915050565b5f806040838503121561461157614610613d7f565b5b5f61461e85828601613dcd565b925050602083013567ffffffffffffffff81111561463f5761463e613d83565b5b61464b85828601613f8b565b9150509250929050565b5f8083601f84011261466a57614669613e7c565b5b8235905067ffffffffffffffff811115614687576146866142b1565b5b6020830191508360208202830111156146a3576146a26142b5565b5b9250929050565b5f805f805f805f805f806101008b8d0312156146c9576146c8613d7f565b5b5f8b013567ffffffffffffffff8111156146e6576146e5613d83565b5b6146f28d828e01614655565b9a509a505060206147058d828e01613df7565b98505060406147168d828e01613dcd565b97505060608b013567ffffffffffffffff81111561473757614736613d83565b5b6147438d828e016142b9565b965096505060806147568d828e01613dcd565b94505060a06147678d828e01613dcd565b93505060c06147788d828e01613df7565b92505060e06147898d828e01614335565b9150509295989b9194979a5092959850565b5f805f805f608086880312156147b4576147b3613d7f565b5b5f6147c188828901613dcd565b95505060206147d288828901613df7565b945050604086013567ffffffffffffffff8111156147f3576147f2613d83565b5b6147ff888289016142b9565b93509350506060614812888289016140c1565b9150509295509295909350565b61482881613da6565b82525050565b5f6040820190508181035f830152614846818561457f565b9050614855602083018461481f565b9392505050565b5f805f805f805f805f805f6101408c8e03121561487c5761487b613d7f565b5b5f6148898e828f01613dcd565b9b5050602061489a8e828f01613df7565b9a505060408c013567ffffffffffffffff8111156148bb576148ba613d83565b5b6148c78e828f016142b9565b995099505060606148da8e828f016140c1565b97505060806148eb8e828f01613df7565b96505060a06148fc8e828f01613df7565b95505060c061490d8e828f01613df7565b94505060e061491e8e828f01613dcd565b9350506101006149308e828f01613dcd565b9250506101206149428e828f01613df7565b9150509295989b509295989b9093969950565b61495e81613e49565b82525050565b5f6020820190506149775f830184614955565b92915050565b5f806040838503121561499357614992613d7f565b5b5f6149a085828601613dcd565b92505060206149b185828601613dcd565b9150509250929050565b5f805f606084860312156149d2576149d1613d7f565b5b5f6149df86828701613dcd565b93505060206149f086828701613dcd565b9250506040614a0186828701613dcd565b9150509250925092565b5f805f60608486031215614a2257614a21613d7f565b5b5f614a2f86828701613dcd565b9350506020614a4086828701613dcd565b9250506040614a5186828701613df7565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f614a7f82614a5b565b614a898185614a65565b9350614a9981856020860161416f565b614aa281613e84565b840191505092915050565b5f6020820190508181035f830152614ac58184614a75565b905092915050565b7f47533230330000000000000000000000000000000000000000000000000000005f82015250565b5f614b01600583614a65565b9150614b0c82614acd565b602082019050919050565b5f6020820190508181035f830152614b2e81614af5565b9050919050565b7f47533230340000000000000000000000000000000000000000000000000000005f82015250565b5f614b69600583614a65565b9150614b7482614b35565b602082019050919050565b5f6020820190508181035f830152614b9681614b5d565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f614bd482613d45565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614c0657614c05614b9d565b5b600182019050919050565b5f602082019050614c245f83018461481f565b92915050565b7f47533032300000000000000000000000000000000000000000000000000000005f82015250565b5f614c5e600583614a65565b9150614c6982614c2a565b602082019050919050565b5f6020820190508181035f830152614c8b81614c52565b9050919050565b7f47533032310000000000000000000000000000000000000000000000000000005f82015250565b5f614cc6600583614a65565b9150614cd182614c92565b602082019050919050565b5f6020820190508181035f830152614cf381614cba565b9050919050565b7f47533032320000000000000000000000000000000000000000000000000000005f82015250565b5f614d2e600583614a65565b9150614d3982614cfa565b602082019050919050565b5f6020820190508181035f830152614d5b81614d22565b9050919050565b7f47533032330000000000000000000000000000000000000000000000000000005f82015250565b5f614d96600583614a65565b9150614da182614d62565b602082019050919050565b5f6020820190508181035f830152614dc381614d8a565b9050919050565b5f6040820190508181035f830152614de28185614197565b90508181036020830152614df68184614197565b90509392505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b614e3381614dff565b8114614e3d575f80fd5b50565b5f81519050614e4e81614e2a565b92915050565b5f60208284031215614e6957614e68613d7f565b5b5f614e7684828501614e40565b91505092915050565b7f47533032340000000000000000000000000000000000000000000000000000005f82015250565b5f614eb3600583614a65565b9150614ebe82614e7f565b602082019050919050565b5f6020820190508181035f830152614ee081614ea7565b9050919050565b7f47533032350000000000000000000000000000000000000000000000000000005f82015250565b5f614f1b600583614a65565b9150614f2682614ee7565b602082019050919050565b5f6020820190508181035f830152614f4881614f0f565b9050919050565b5f81905092915050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f82015250565b5f614f8d601c83614f4f565b9150614f9882614f59565b601c82019050919050565b5f819050919050565b614fbd614fb882613e49565b614fa3565b82525050565b5f614fcd82614f81565b9150614fd98284614fac565b60208201915081905092915050565b5f60ff82169050919050565b5f614ffe82614fe8565b915061500983614fe8565b9250828203905060ff81111561502257615021614b9d565b5b92915050565b61503181614fe8565b82525050565b5f60808201905061504a5f830187614955565b6150576020830186615028565b6150646040830185614955565b6150716060830184614955565b95945050505050565b7f47533032360000000000000000000000000000000000000000000000000000005f82015250565b5f6150ae600583614a65565b91506150b98261507a565b602082019050919050565b5f6020820190508181035f8301526150db816150a2565b9050919050565b7f47533130340000000000000000000000000000000000000000000000000000005f82015250565b5f615116600583614a65565b9150615121826150e2565b602082019050919050565b5f6020820190508181035f8301526151438161510a565b9050919050565b5f61515482613d45565b915061515f83613d45565b925082820261516d81613d45565b9150828204841483151761518457615183614b9d565b5b5092915050565b7f47533130310000000000000000000000000000000000000000000000000000005f82015250565b5f6151bf600583614a65565b91506151ca8261518b565b602082019050919050565b5f6020820190508181035f8301526151ec816151b3565b9050919050565b7f47533130320000000000000000000000000000000000000000000000000000005f82015250565b5f615227600583614a65565b9150615232826151f3565b602082019050919050565b5f6020820190508181035f8301526152548161521b565b9050919050565b7f47533230310000000000000000000000000000000000000000000000000000005f82015250565b5f61528f600583614a65565b915061529a8261525b565b602082019050919050565b5f6020820190508181035f8301526152bc81615283565b9050919050565b7f47533230320000000000000000000000000000000000000000000000000000005f82015250565b5f6152f7600583614a65565b9150615302826152c3565b602082019050919050565b5f6020820190508181035f830152615324816152eb565b9050919050565b5f615336838561415f565b9350615343838584613f3c565b61534c83613e84565b840190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6002811061539557615394615357565b5b50565b5f8190506153a582615384565b919050565b5f6153b482615398565b9050919050565b6153c4816153aa565b82525050565b6153d38161430e565b82525050565b5f610160820190506153ed5f83018f61481f565b6153fa602083018e613d4e565b818103604083015261540d818c8e61532b565b905061541c606083018b6153bb565b615429608083018a613d4e565b61543660a0830189613d4e565b61544360c0830188613d4e565b61545060e083018761481f565b61545e6101008301866153ca565b8181036101208301526154718185614197565b905061548161014083018461481f565b9d9c50505050505050505050505050565b5f61549c82613d45565b91506154a783613d45565b92508282019050808211156154bf576154be614b9d565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6154fc82613d45565b915061550783613d45565b925082615517576155166154c5565b5b828204905092915050565b7f47533031300000000000000000000000000000000000000000000000000000005f82015250565b5f615556600583614a65565b915061556182615522565b602082019050919050565b5f6020820190508181035f8301526155838161554a565b9050919050565b5f61559482613d45565b915061559f83613d45565b92508282039050818111156155b7576155b6614b9d565b5b92915050565b7f47533031330000000000000000000000000000000000000000000000000000005f82015250565b5f6155f1600583614a65565b91506155fc826155bd565b602082019050919050565b5f6020820190508181035f83015261561e816155e5565b9050919050565b5f6040820190506156385f830185614955565b6156456020830184613d4e565b9392505050565b5f60408201905061565f5f830185614955565b61566c602083018461408a565b9392505050565b7f47533030310000000000000000000000000000000000000000000000000000005f82015250565b5f6156a7600583614a65565b91506156b282615673565b602082019050919050565b5f6020820190508181035f8301526156d48161569b565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050919050565b5f61571f6020840184613dcd565b905092915050565b5f602082019050919050565b5f61573e838561452e565b935061574982615708565b805f5b858110156157815761575e8284615711565b615768888261455c565b975061577383615727565b92505060018101905061574c565b5085925050509392505050565b5f6080820190508181035f8301526157a7818789615733565b90506157b66020830186613d4e565b6157c3604083018561481f565b6157d0606083018461481f565b9695505050505050565b5f819050919050565b6157f46157ef82613d45565b6157da565b82525050565b5f61580582846157e3565b60208201915081905092915050565b7f47533033300000000000000000000000000000000000000000000000000000005f82015250565b5f615848600583614a65565b915061585382615814565b602082019050919050565b5f6020820190508181035f8301526158758161583c565b9050919050565b7f47533130330000000000000000000000000000000000000000000000000000005f82015250565b5f6158b0600583614a65565b91506158bb8261587c565b602082019050919050565b5f6020820190508181035f8301526158dd816158a4565b9050919050565b7f47533230350000000000000000000000000000000000000000000000000000005f82015250565b5f615918600583614a65565b9150615923826158e4565b602082019050919050565b5f6020820190508181035f8301526159458161590c565b9050919050565b5f81905092915050565b5f615961838561594c565b935061596e838584613f3c565b82840190509392505050565b5f615986828486615956565b91508190509392505050565b5f610160820190506159a65f83018e614955565b6159b3602083018d61481f565b6159c0604083018c613d4e565b6159cd606083018b614955565b6159da608083018a6153bb565b6159e760a0830189613d4e565b6159f460c0830188613d4e565b615a0160e0830187613d4e565b615a0f61010083018661481f565b615a1d61012083018561481f565b615a2b610140830184613d4e565b9c9b505050505050505050505050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b5f819050919050565b615a80615a7b82615a3b565b615a66565b82525050565b5f615a918287615a6f565b600182019150615aa18286615a6f565b600182019150615ab18285614fac565b602082019150615ac18284614fac565b60208201915081905095945050505050565b5f819050919050565b5f615af6615af1615aec84613d87565b615ad3565b613d87565b9050919050565b5f615b0782615adc565b9050919050565b5f615b1882615afd565b9050919050565b615b2881615b0e565b82525050565b5f606082019050615b415f830186614955565b615b4e6020830185613d4e565b615b5b6040830184615b1f565b949350505050565b5f615b6d82613d45565b91505f8203615b7f57615b7e614b9d565b5b600182039050919050565b7f47533033310000000000000000000000000000000000000000000000000000005f82015250565b5f615bbe600583614a65565b9150615bc982615b8a565b602082019050919050565b5f6020820190508181035f830152615beb81615bb2565b9050919050565b7f47533031310000000000000000000000000000000000000000000000000000005f82015250565b5f615c26600583614a65565b9150615c3182615bf2565b602082019050919050565b5f6020820190508181035f830152615c5381615c1a565b9050919050565b7f47533031320000000000000000000000000000000000000000000000000000005f82015250565b5f615c8e600583614a65565b9150615c9982615c5a565b602082019050919050565b5f6020820190508181035f830152615cbb81615c82565b9050919050565b7f47533230300000000000000000000000000000000000000000000000000000005f82015250565b5f615cf6600583614a65565b9150615d0182615cc2565b602082019050919050565b5f6020820190508181035f830152615d2381615cea565b9050919050565b7f47533130300000000000000000000000000000000000000000000000000000005f82015250565b5f615d5e600583614a65565b9150615d6982615d2a565b602082019050919050565b5f6020820190508181035f830152615d8b81615d52565b9050919050565b7f47533030300000000000000000000000000000000000000000000000000000005f82015250565b5f615dc6600583614a65565b9150615dd182615d92565b602082019050919050565b5f6020820190508181035f830152615df381615dba565b9050919050565b5f604082019050615e0d5f83018561481f565b615e1a6020830184613d4e565b939250505056fea26469706673582212200eba17a72cdc4801ed8ccdd305e1a632ccc4ce04ffcea0a9f33e4e424fee857964736f6c63430008150033",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainMetaData.Bin instead.
var MainBin = MainMetaData.Bin

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Main *MainCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Main *MainSession) VERSION() (string, error) {
	return _Main.Contract.VERSION(&_Main.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Main *MainCallerSession) VERSION() (string, error) {
	return _Main.Contract.VERSION(&_Main.CallOpts)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_Main *MainCaller) ApprovedHashes(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "approvedHashes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_Main *MainSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _Main.Contract.ApprovedHashes(&_Main.CallOpts, arg0, arg1)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_Main *MainCallerSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _Main.Contract.ApprovedHashes(&_Main.CallOpts, arg0, arg1)
}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_Main *MainCaller) CheckNSignatures(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "checkNSignatures", dataHash, data, signatures, requiredSignatures)

	if err != nil {
		return err
	}

	return err

}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_Main *MainSession) CheckNSignatures(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _Main.Contract.CheckNSignatures(&_Main.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_Main *MainCallerSession) CheckNSignatures(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _Main.Contract.CheckNSignatures(&_Main.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_Main *MainCaller) CheckSignatures(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte) error {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "checkSignatures", dataHash, data, signatures)

	if err != nil {
		return err
	}

	return err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_Main *MainSession) CheckSignatures(dataHash [32]byte, data []byte, signatures []byte) error {
	return _Main.Contract.CheckSignatures(&_Main.CallOpts, dataHash, data, signatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_Main *MainCallerSession) CheckSignatures(dataHash [32]byte, data []byte, signatures []byte) error {
	return _Main.Contract.CheckSignatures(&_Main.CallOpts, dataHash, data, signatures)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Main *MainCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Main *MainSession) DomainSeparator() ([32]byte, error) {
	return _Main.Contract.DomainSeparator(&_Main.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Main *MainCallerSession) DomainSeparator() ([32]byte, error) {
	return _Main.Contract.DomainSeparator(&_Main.CallOpts)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_Main *MainCaller) EncodeTransactionData(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "encodeTransactionData", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_Main *MainSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _Main.Contract.EncodeTransactionData(&_Main.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_Main *MainCallerSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _Main.Contract.EncodeTransactionData(&_Main.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Main *MainCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Main *MainSession) GetChainId() (*big.Int, error) {
	return _Main.Contract.GetChainId(&_Main.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Main *MainCallerSession) GetChainId() (*big.Int, error) {
	return _Main.Contract.GetChainId(&_Main.CallOpts)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_Main *MainCaller) GetModulesPaginated(opts *bind.CallOpts, start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getModulesPaginated", start, pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_Main *MainSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _Main.Contract.GetModulesPaginated(&_Main.CallOpts, start, pageSize)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_Main *MainCallerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _Main.Contract.GetModulesPaginated(&_Main.CallOpts, start, pageSize)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Main *MainCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Main *MainSession) GetOwners() ([]common.Address, error) {
	return _Main.Contract.GetOwners(&_Main.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Main *MainCallerSession) GetOwners() ([]common.Address, error) {
	return _Main.Contract.GetOwners(&_Main.CallOpts)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_Main *MainCaller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_Main *MainSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _Main.Contract.GetStorageAt(&_Main.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_Main *MainCallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _Main.Contract.GetStorageAt(&_Main.CallOpts, offset, length)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Main *MainCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Main *MainSession) GetThreshold() (*big.Int, error) {
	return _Main.Contract.GetThreshold(&_Main.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Main *MainCallerSession) GetThreshold() (*big.Int, error) {
	return _Main.Contract.GetThreshold(&_Main.CallOpts)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_Main *MainCaller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getTransactionHash", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_Main *MainSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _Main.Contract.GetTransactionHash(&_Main.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_Main *MainCallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _Main.Contract.GetTransactionHash(&_Main.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_Main *MainCaller) IsModuleEnabled(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "isModuleEnabled", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_Main *MainSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _Main.Contract.IsModuleEnabled(&_Main.CallOpts, module)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_Main *MainCallerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _Main.Contract.IsModuleEnabled(&_Main.CallOpts, module)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_Main *MainCaller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "isOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_Main *MainSession) IsOwner(owner common.Address) (bool, error) {
	return _Main.Contract.IsOwner(&_Main.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_Main *MainCallerSession) IsOwner(owner common.Address) (bool, error) {
	return _Main.Contract.IsOwner(&_Main.CallOpts, owner)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Main *MainCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Main *MainSession) Nonce() (*big.Int, error) {
	return _Main.Contract.Nonce(&_Main.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Main *MainCallerSession) Nonce() (*big.Int, error) {
	return _Main.Contract.Nonce(&_Main.CallOpts)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_Main *MainCaller) SignedMessages(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "signedMessages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_Main *MainSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _Main.Contract.SignedMessages(&_Main.CallOpts, arg0)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_Main *MainCallerSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _Main.Contract.SignedMessages(&_Main.CallOpts, arg0)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_Main *MainTransactor) AddOwnerWithThreshold(opts *bind.TransactOpts, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addOwnerWithThreshold", owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_Main *MainSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AddOwnerWithThreshold(&_Main.TransactOpts, owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_Main *MainTransactorSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AddOwnerWithThreshold(&_Main.TransactOpts, owner, _threshold)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_Main *MainTransactor) ApproveHash(opts *bind.TransactOpts, hashToApprove [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approveHash", hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_Main *MainSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _Main.Contract.ApproveHash(&_Main.TransactOpts, hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_Main *MainTransactorSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _Main.Contract.ApproveHash(&_Main.TransactOpts, hashToApprove)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_Main *MainTransactor) ChangeThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "changeThreshold", _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_Main *MainSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ChangeThreshold(&_Main.TransactOpts, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_Main *MainTransactorSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ChangeThreshold(&_Main.TransactOpts, _threshold)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_Main *MainTransactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_Main *MainSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _Main.Contract.DisableModule(&_Main.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_Main *MainTransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _Main.Contract.DisableModule(&_Main.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_Main *MainTransactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_Main *MainSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _Main.Contract.EnableModule(&_Main.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_Main *MainTransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _Main.Contract.EnableModule(&_Main.TransactOpts, module)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_Main *MainTransactor) ExecTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "execTransaction", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_Main *MainSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _Main.Contract.ExecTransaction(&_Main.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_Main *MainTransactorSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _Main.Contract.ExecTransaction(&_Main.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_Main *MainTransactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_Main *MainSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Main.Contract.ExecTransactionFromModule(&_Main.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_Main *MainTransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Main.Contract.ExecTransactionFromModule(&_Main.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_Main *MainTransactor) ExecTransactionFromModuleReturnData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "execTransactionFromModuleReturnData", to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_Main *MainSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Main.Contract.ExecTransactionFromModuleReturnData(&_Main.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_Main *MainTransactorSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Main.Contract.ExecTransactionFromModuleReturnData(&_Main.TransactOpts, to, value, data, operation)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_Main *MainTransactor) RemoveOwner(opts *bind.TransactOpts, prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeOwner", prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_Main *MainSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveOwner(&_Main.TransactOpts, prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_Main *MainTransactorSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RemoveOwner(&_Main.TransactOpts, prevOwner, owner, _threshold)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_Main *MainTransactor) RequiredTxGas(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "requiredTxGas", to, value, data, operation)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_Main *MainSession) RequiredTxGas(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Main.Contract.RequiredTxGas(&_Main.TransactOpts, to, value, data, operation)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_Main *MainTransactorSession) RequiredTxGas(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Main.Contract.RequiredTxGas(&_Main.TransactOpts, to, value, data, operation)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_Main *MainTransactor) SetFallbackHandler(opts *bind.TransactOpts, handler common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setFallbackHandler", handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_Main *MainSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetFallbackHandler(&_Main.TransactOpts, handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_Main *MainTransactorSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetFallbackHandler(&_Main.TransactOpts, handler)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_Main *MainTransactor) SetGuard(opts *bind.TransactOpts, guard common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setGuard", guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_Main *MainSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetGuard(&_Main.TransactOpts, guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_Main *MainTransactorSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetGuard(&_Main.TransactOpts, guard)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_Main *MainTransactor) Setup(opts *bind.TransactOpts, _owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setup", _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_Main *MainSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _Main.Contract.Setup(&_Main.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_Main *MainTransactorSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _Main.Contract.Setup(&_Main.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_Main *MainTransactor) SimulateAndRevert(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "simulateAndRevert", targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_Main *MainSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Main.Contract.SimulateAndRevert(&_Main.TransactOpts, targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_Main *MainTransactorSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Main.Contract.SimulateAndRevert(&_Main.TransactOpts, targetContract, calldataPayload)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_Main *MainTransactor) SwapOwner(opts *bind.TransactOpts, prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "swapOwner", prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_Main *MainSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.SwapOwner(&_Main.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_Main *MainTransactorSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.SwapOwner(&_Main.TransactOpts, prevOwner, oldOwner, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Main *MainTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Main.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Main *MainSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Main.Contract.Fallback(&_Main.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Main *MainTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Main.Contract.Fallback(&_Main.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Main *MainTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Main *MainSession) Receive() (*types.Transaction, error) {
	return _Main.Contract.Receive(&_Main.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Main *MainTransactorSession) Receive() (*types.Transaction, error) {
	return _Main.Contract.Receive(&_Main.TransactOpts)
}

// MainAddedOwnerIterator is returned from FilterAddedOwner and is used to iterate over the raw logs and unpacked data for AddedOwner events raised by the Main contract.
type MainAddedOwnerIterator struct {
	Event *MainAddedOwner // Event containing the contract specifics and raw log

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
func (it *MainAddedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAddedOwner)
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
		it.Event = new(MainAddedOwner)
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
func (it *MainAddedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAddedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAddedOwner represents a AddedOwner event raised by the Main contract.
type MainAddedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedOwner is a free log retrieval operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_Main *MainFilterer) FilterAddedOwner(opts *bind.FilterOpts) (*MainAddedOwnerIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return &MainAddedOwnerIterator{contract: _Main.contract, event: "AddedOwner", logs: logs, sub: sub}, nil
}

// WatchAddedOwner is a free log subscription operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_Main *MainFilterer) WatchAddedOwner(opts *bind.WatchOpts, sink chan<- *MainAddedOwner) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAddedOwner)
				if err := _Main.contract.UnpackLog(event, "AddedOwner", log); err != nil {
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

// ParseAddedOwner is a log parse operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_Main *MainFilterer) ParseAddedOwner(log types.Log) (*MainAddedOwner, error) {
	event := new(MainAddedOwner)
	if err := _Main.contract.UnpackLog(event, "AddedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainApproveHashIterator is returned from FilterApproveHash and is used to iterate over the raw logs and unpacked data for ApproveHash events raised by the Main contract.
type MainApproveHashIterator struct {
	Event *MainApproveHash // Event containing the contract specifics and raw log

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
func (it *MainApproveHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainApproveHash)
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
		it.Event = new(MainApproveHash)
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
func (it *MainApproveHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainApproveHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainApproveHash represents a ApproveHash event raised by the Main contract.
type MainApproveHash struct {
	ApprovedHash [32]byte
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApproveHash is a free log retrieval operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_Main *MainFilterer) FilterApproveHash(opts *bind.FilterOpts, approvedHash [][32]byte, owner []common.Address) (*MainApproveHashIterator, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MainApproveHashIterator{contract: _Main.contract, event: "ApproveHash", logs: logs, sub: sub}, nil
}

// WatchApproveHash is a free log subscription operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_Main *MainFilterer) WatchApproveHash(opts *bind.WatchOpts, sink chan<- *MainApproveHash, approvedHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainApproveHash)
				if err := _Main.contract.UnpackLog(event, "ApproveHash", log); err != nil {
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

// ParseApproveHash is a log parse operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_Main *MainFilterer) ParseApproveHash(log types.Log) (*MainApproveHash, error) {
	event := new(MainApproveHash)
	if err := _Main.contract.UnpackLog(event, "ApproveHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainChangedFallbackHandlerIterator is returned from FilterChangedFallbackHandler and is used to iterate over the raw logs and unpacked data for ChangedFallbackHandler events raised by the Main contract.
type MainChangedFallbackHandlerIterator struct {
	Event *MainChangedFallbackHandler // Event containing the contract specifics and raw log

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
func (it *MainChangedFallbackHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainChangedFallbackHandler)
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
		it.Event = new(MainChangedFallbackHandler)
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
func (it *MainChangedFallbackHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainChangedFallbackHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainChangedFallbackHandler represents a ChangedFallbackHandler event raised by the Main contract.
type MainChangedFallbackHandler struct {
	Handler common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangedFallbackHandler is a free log retrieval operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_Main *MainFilterer) FilterChangedFallbackHandler(opts *bind.FilterOpts) (*MainChangedFallbackHandlerIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ChangedFallbackHandler")
	if err != nil {
		return nil, err
	}
	return &MainChangedFallbackHandlerIterator{contract: _Main.contract, event: "ChangedFallbackHandler", logs: logs, sub: sub}, nil
}

// WatchChangedFallbackHandler is a free log subscription operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_Main *MainFilterer) WatchChangedFallbackHandler(opts *bind.WatchOpts, sink chan<- *MainChangedFallbackHandler) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ChangedFallbackHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainChangedFallbackHandler)
				if err := _Main.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
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

// ParseChangedFallbackHandler is a log parse operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_Main *MainFilterer) ParseChangedFallbackHandler(log types.Log) (*MainChangedFallbackHandler, error) {
	event := new(MainChangedFallbackHandler)
	if err := _Main.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainChangedGuardIterator is returned from FilterChangedGuard and is used to iterate over the raw logs and unpacked data for ChangedGuard events raised by the Main contract.
type MainChangedGuardIterator struct {
	Event *MainChangedGuard // Event containing the contract specifics and raw log

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
func (it *MainChangedGuardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainChangedGuard)
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
		it.Event = new(MainChangedGuard)
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
func (it *MainChangedGuardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainChangedGuardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainChangedGuard represents a ChangedGuard event raised by the Main contract.
type MainChangedGuard struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangedGuard is a free log retrieval operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_Main *MainFilterer) FilterChangedGuard(opts *bind.FilterOpts) (*MainChangedGuardIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ChangedGuard")
	if err != nil {
		return nil, err
	}
	return &MainChangedGuardIterator{contract: _Main.contract, event: "ChangedGuard", logs: logs, sub: sub}, nil
}

// WatchChangedGuard is a free log subscription operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_Main *MainFilterer) WatchChangedGuard(opts *bind.WatchOpts, sink chan<- *MainChangedGuard) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ChangedGuard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainChangedGuard)
				if err := _Main.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
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

// ParseChangedGuard is a log parse operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_Main *MainFilterer) ParseChangedGuard(log types.Log) (*MainChangedGuard, error) {
	event := new(MainChangedGuard)
	if err := _Main.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainChangedThresholdIterator is returned from FilterChangedThreshold and is used to iterate over the raw logs and unpacked data for ChangedThreshold events raised by the Main contract.
type MainChangedThresholdIterator struct {
	Event *MainChangedThreshold // Event containing the contract specifics and raw log

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
func (it *MainChangedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainChangedThreshold)
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
		it.Event = new(MainChangedThreshold)
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
func (it *MainChangedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainChangedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainChangedThreshold represents a ChangedThreshold event raised by the Main contract.
type MainChangedThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedThreshold is a free log retrieval operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_Main *MainFilterer) FilterChangedThreshold(opts *bind.FilterOpts) (*MainChangedThresholdIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return &MainChangedThresholdIterator{contract: _Main.contract, event: "ChangedThreshold", logs: logs, sub: sub}, nil
}

// WatchChangedThreshold is a free log subscription operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_Main *MainFilterer) WatchChangedThreshold(opts *bind.WatchOpts, sink chan<- *MainChangedThreshold) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainChangedThreshold)
				if err := _Main.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
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

// ParseChangedThreshold is a log parse operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_Main *MainFilterer) ParseChangedThreshold(log types.Log) (*MainChangedThreshold, error) {
	event := new(MainChangedThreshold)
	if err := _Main.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainDisabledModuleIterator is returned from FilterDisabledModule and is used to iterate over the raw logs and unpacked data for DisabledModule events raised by the Main contract.
type MainDisabledModuleIterator struct {
	Event *MainDisabledModule // Event containing the contract specifics and raw log

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
func (it *MainDisabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainDisabledModule)
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
		it.Event = new(MainDisabledModule)
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
func (it *MainDisabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainDisabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainDisabledModule represents a DisabledModule event raised by the Main contract.
type MainDisabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisabledModule is a free log retrieval operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_Main *MainFilterer) FilterDisabledModule(opts *bind.FilterOpts) (*MainDisabledModuleIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return &MainDisabledModuleIterator{contract: _Main.contract, event: "DisabledModule", logs: logs, sub: sub}, nil
}

// WatchDisabledModule is a free log subscription operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_Main *MainFilterer) WatchDisabledModule(opts *bind.WatchOpts, sink chan<- *MainDisabledModule) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainDisabledModule)
				if err := _Main.contract.UnpackLog(event, "DisabledModule", log); err != nil {
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

// ParseDisabledModule is a log parse operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_Main *MainFilterer) ParseDisabledModule(log types.Log) (*MainDisabledModule, error) {
	event := new(MainDisabledModule)
	if err := _Main.contract.UnpackLog(event, "DisabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainEnabledModuleIterator is returned from FilterEnabledModule and is used to iterate over the raw logs and unpacked data for EnabledModule events raised by the Main contract.
type MainEnabledModuleIterator struct {
	Event *MainEnabledModule // Event containing the contract specifics and raw log

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
func (it *MainEnabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainEnabledModule)
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
		it.Event = new(MainEnabledModule)
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
func (it *MainEnabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainEnabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainEnabledModule represents a EnabledModule event raised by the Main contract.
type MainEnabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledModule is a free log retrieval operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_Main *MainFilterer) FilterEnabledModule(opts *bind.FilterOpts) (*MainEnabledModuleIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return &MainEnabledModuleIterator{contract: _Main.contract, event: "EnabledModule", logs: logs, sub: sub}, nil
}

// WatchEnabledModule is a free log subscription operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_Main *MainFilterer) WatchEnabledModule(opts *bind.WatchOpts, sink chan<- *MainEnabledModule) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainEnabledModule)
				if err := _Main.contract.UnpackLog(event, "EnabledModule", log); err != nil {
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

// ParseEnabledModule is a log parse operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_Main *MainFilterer) ParseEnabledModule(log types.Log) (*MainEnabledModule, error) {
	event := new(MainEnabledModule)
	if err := _Main.contract.UnpackLog(event, "EnabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the Main contract.
type MainExecutionFailureIterator struct {
	Event *MainExecutionFailure // Event containing the contract specifics and raw log

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
func (it *MainExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainExecutionFailure)
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
		it.Event = new(MainExecutionFailure)
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
func (it *MainExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainExecutionFailure represents a ExecutionFailure event raised by the Main contract.
type MainExecutionFailure struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_Main *MainFilterer) FilterExecutionFailure(opts *bind.FilterOpts) (*MainExecutionFailureIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return &MainExecutionFailureIterator{contract: _Main.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_Main *MainFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MainExecutionFailure) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainExecutionFailure)
				if err := _Main.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// ParseExecutionFailure is a log parse operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_Main *MainFilterer) ParseExecutionFailure(log types.Log) (*MainExecutionFailure, error) {
	event := new(MainExecutionFailure)
	if err := _Main.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainExecutionFromModuleFailureIterator is returned from FilterExecutionFromModuleFailure and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleFailure events raised by the Main contract.
type MainExecutionFromModuleFailureIterator struct {
	Event *MainExecutionFromModuleFailure // Event containing the contract specifics and raw log

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
func (it *MainExecutionFromModuleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainExecutionFromModuleFailure)
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
		it.Event = new(MainExecutionFromModuleFailure)
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
func (it *MainExecutionFromModuleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainExecutionFromModuleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainExecutionFromModuleFailure represents a ExecutionFromModuleFailure event raised by the Main contract.
type MainExecutionFromModuleFailure struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleFailure is a free log retrieval operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_Main *MainFilterer) FilterExecutionFromModuleFailure(opts *bind.FilterOpts, module []common.Address) (*MainExecutionFromModuleFailureIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return &MainExecutionFromModuleFailureIterator{contract: _Main.contract, event: "ExecutionFromModuleFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleFailure is a free log subscription operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_Main *MainFilterer) WatchExecutionFromModuleFailure(opts *bind.WatchOpts, sink chan<- *MainExecutionFromModuleFailure, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainExecutionFromModuleFailure)
				if err := _Main.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
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

// ParseExecutionFromModuleFailure is a log parse operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_Main *MainFilterer) ParseExecutionFromModuleFailure(log types.Log) (*MainExecutionFromModuleFailure, error) {
	event := new(MainExecutionFromModuleFailure)
	if err := _Main.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainExecutionFromModuleSuccessIterator is returned from FilterExecutionFromModuleSuccess and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleSuccess events raised by the Main contract.
type MainExecutionFromModuleSuccessIterator struct {
	Event *MainExecutionFromModuleSuccess // Event containing the contract specifics and raw log

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
func (it *MainExecutionFromModuleSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainExecutionFromModuleSuccess)
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
		it.Event = new(MainExecutionFromModuleSuccess)
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
func (it *MainExecutionFromModuleSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainExecutionFromModuleSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainExecutionFromModuleSuccess represents a ExecutionFromModuleSuccess event raised by the Main contract.
type MainExecutionFromModuleSuccess struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleSuccess is a free log retrieval operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_Main *MainFilterer) FilterExecutionFromModuleSuccess(opts *bind.FilterOpts, module []common.Address) (*MainExecutionFromModuleSuccessIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return &MainExecutionFromModuleSuccessIterator{contract: _Main.contract, event: "ExecutionFromModuleSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleSuccess is a free log subscription operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_Main *MainFilterer) WatchExecutionFromModuleSuccess(opts *bind.WatchOpts, sink chan<- *MainExecutionFromModuleSuccess, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainExecutionFromModuleSuccess)
				if err := _Main.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
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

// ParseExecutionFromModuleSuccess is a log parse operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_Main *MainFilterer) ParseExecutionFromModuleSuccess(log types.Log) (*MainExecutionFromModuleSuccess, error) {
	event := new(MainExecutionFromModuleSuccess)
	if err := _Main.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the Main contract.
type MainExecutionSuccessIterator struct {
	Event *MainExecutionSuccess // Event containing the contract specifics and raw log

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
func (it *MainExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainExecutionSuccess)
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
		it.Event = new(MainExecutionSuccess)
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
func (it *MainExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainExecutionSuccess represents a ExecutionSuccess event raised by the Main contract.
type MainExecutionSuccess struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_Main *MainFilterer) FilterExecutionSuccess(opts *bind.FilterOpts) (*MainExecutionSuccessIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return &MainExecutionSuccessIterator{contract: _Main.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_Main *MainFilterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *MainExecutionSuccess) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainExecutionSuccess)
				if err := _Main.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
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

// ParseExecutionSuccess is a log parse operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_Main *MainFilterer) ParseExecutionSuccess(log types.Log) (*MainExecutionSuccess, error) {
	event := new(MainExecutionSuccess)
	if err := _Main.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRemovedOwnerIterator is returned from FilterRemovedOwner and is used to iterate over the raw logs and unpacked data for RemovedOwner events raised by the Main contract.
type MainRemovedOwnerIterator struct {
	Event *MainRemovedOwner // Event containing the contract specifics and raw log

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
func (it *MainRemovedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRemovedOwner)
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
		it.Event = new(MainRemovedOwner)
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
func (it *MainRemovedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRemovedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRemovedOwner represents a RemovedOwner event raised by the Main contract.
type MainRemovedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedOwner is a free log retrieval operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_Main *MainFilterer) FilterRemovedOwner(opts *bind.FilterOpts) (*MainRemovedOwnerIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return &MainRemovedOwnerIterator{contract: _Main.contract, event: "RemovedOwner", logs: logs, sub: sub}, nil
}

// WatchRemovedOwner is a free log subscription operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_Main *MainFilterer) WatchRemovedOwner(opts *bind.WatchOpts, sink chan<- *MainRemovedOwner) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRemovedOwner)
				if err := _Main.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
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

// ParseRemovedOwner is a log parse operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_Main *MainFilterer) ParseRemovedOwner(log types.Log) (*MainRemovedOwner, error) {
	event := new(MainRemovedOwner)
	if err := _Main.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSafeReceivedIterator is returned from FilterSafeReceived and is used to iterate over the raw logs and unpacked data for SafeReceived events raised by the Main contract.
type MainSafeReceivedIterator struct {
	Event *MainSafeReceived // Event containing the contract specifics and raw log

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
func (it *MainSafeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSafeReceived)
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
		it.Event = new(MainSafeReceived)
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
func (it *MainSafeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSafeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSafeReceived represents a SafeReceived event raised by the Main contract.
type MainSafeReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeReceived is a free log retrieval operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_Main *MainFilterer) FilterSafeReceived(opts *bind.FilterOpts, sender []common.Address) (*MainSafeReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &MainSafeReceivedIterator{contract: _Main.contract, event: "SafeReceived", logs: logs, sub: sub}, nil
}

// WatchSafeReceived is a free log subscription operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_Main *MainFilterer) WatchSafeReceived(opts *bind.WatchOpts, sink chan<- *MainSafeReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSafeReceived)
				if err := _Main.contract.UnpackLog(event, "SafeReceived", log); err != nil {
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

// ParseSafeReceived is a log parse operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_Main *MainFilterer) ParseSafeReceived(log types.Log) (*MainSafeReceived, error) {
	event := new(MainSafeReceived)
	if err := _Main.contract.UnpackLog(event, "SafeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSafeSetupIterator is returned from FilterSafeSetup and is used to iterate over the raw logs and unpacked data for SafeSetup events raised by the Main contract.
type MainSafeSetupIterator struct {
	Event *MainSafeSetup // Event containing the contract specifics and raw log

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
func (it *MainSafeSetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSafeSetup)
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
		it.Event = new(MainSafeSetup)
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
func (it *MainSafeSetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSafeSetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSafeSetup represents a SafeSetup event raised by the Main contract.
type MainSafeSetup struct {
	Initiator       common.Address
	Owners          []common.Address
	Threshold       *big.Int
	Initializer     common.Address
	FallbackHandler common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSafeSetup is a free log retrieval operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_Main *MainFilterer) FilterSafeSetup(opts *bind.FilterOpts, initiator []common.Address) (*MainSafeSetupIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &MainSafeSetupIterator{contract: _Main.contract, event: "SafeSetup", logs: logs, sub: sub}, nil
}

// WatchSafeSetup is a free log subscription operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_Main *MainFilterer) WatchSafeSetup(opts *bind.WatchOpts, sink chan<- *MainSafeSetup, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSafeSetup)
				if err := _Main.contract.UnpackLog(event, "SafeSetup", log); err != nil {
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

// ParseSafeSetup is a log parse operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_Main *MainFilterer) ParseSafeSetup(log types.Log) (*MainSafeSetup, error) {
	event := new(MainSafeSetup)
	if err := _Main.contract.UnpackLog(event, "SafeSetup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSignMsgIterator is returned from FilterSignMsg and is used to iterate over the raw logs and unpacked data for SignMsg events raised by the Main contract.
type MainSignMsgIterator struct {
	Event *MainSignMsg // Event containing the contract specifics and raw log

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
func (it *MainSignMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSignMsg)
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
		it.Event = new(MainSignMsg)
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
func (it *MainSignMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSignMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSignMsg represents a SignMsg event raised by the Main contract.
type MainSignMsg struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignMsg is a free log retrieval operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_Main *MainFilterer) FilterSignMsg(opts *bind.FilterOpts, msgHash [][32]byte) (*MainSignMsgIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &MainSignMsgIterator{contract: _Main.contract, event: "SignMsg", logs: logs, sub: sub}, nil
}

// WatchSignMsg is a free log subscription operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_Main *MainFilterer) WatchSignMsg(opts *bind.WatchOpts, sink chan<- *MainSignMsg, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSignMsg)
				if err := _Main.contract.UnpackLog(event, "SignMsg", log); err != nil {
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

// ParseSignMsg is a log parse operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_Main *MainFilterer) ParseSignMsg(log types.Log) (*MainSignMsg, error) {
	event := new(MainSignMsg)
	if err := _Main.contract.UnpackLog(event, "SignMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
