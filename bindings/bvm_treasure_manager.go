// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// TreasureManagerMetaData contains all meta data concerning the TreasureManager contract.
var TreasureManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimAllTokens\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositERC20\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ethAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenWhiteList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRewards\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"granter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_treasureManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_withdrawManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryReward\",\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenWhiteList\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWithdrawManager\",\"inputs\":[{\"name\":\"_withdrawManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenWhiteList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasureManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userRewardAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GrantRewardTokenAmount\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"granter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawManagerUpdate\",\"inputs\":[{\"name\":\"withdrawManager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b506119b3806100206000396000f3fe6080604052600436106101bb5760003560e01c80634782f779116100ec578063a2672ad71161008a578063d547741f11610064578063d547741f1461054b578063ec3e9da51461056b578063f2fde38b1461058b578063f6326fb3146105ab57600080fd5b8063a2672ad7146104c8578063c0c53b8b146104e8578063c37220ea1461050857600080fd5b80638da5cb5b116100c65780638da5cb5b1461043657806391d148541461047357806397feb92614610493578063a217fddf146104b357600080fd5b80634782f779146103e1578063523fba7f146103f4578063715018a61461042157600080fd5b8063254c5d871161015957806336568abe1161013357806336568abe14610359578063410c9f9a1461037957806341398b151461039957806344004cc1146103c157600080fd5b8063254c5d87146102e15780632f2ff15d1461031957806332f289cf1461033957600080fd5b806317e3e2e81161019557806317e3e2e81461025c5780631e4bd42c1461027c57806323ecdbee14610291578063248a9ca3146102b357600080fd5b806301ffc9a7146101cf57806303186d22146102045780630ec08b001461023c57600080fd5b366101ca576101c86105af565b005b600080fd5b3480156101db57600080fd5b506101ef6101ea36600461169f565b610688565b60405190151581526020015b60405180910390f35b34801561021057600080fd5b50600054610224906001600160a01b031681565b6040516001600160a01b0390911681526020016101fb565b34801561024857600080fd5b506102246102573660046116c9565b6106bf565b34801561026857600080fd5b506101c86102773660046116f7565b6106e9565b34801561028857600080fd5b506101c861073b565b34801561029d57600080fd5b506102a66108b7565b6040516101fb9190611714565b3480156102bf57600080fd5b506102d36102ce3660046116c9565b610919565b6040519081526020016101fb565b3480156102ed57600080fd5b506102d36102fc366004611761565b600460209081526000928352604080842090915290825290205481565b34801561032557600080fd5b506101c861033436600461179a565b61093b565b34801561034557600080fd5b506101c86103543660046116f7565b61095d565b34801561036557600080fd5b506101c861037436600461179a565b610b25565b34801561038557600080fd5b506101c86103943660046117bf565b610b5d565b3480156103a557600080fd5b5061022473eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b3480156103cd57600080fd5b506101ef6103dc3660046117bf565b610c6c565b6101ef6103ef366004611800565b610d8d565b34801561040057600080fd5b506102d361040f3660046116f7565b60036020526000908152604090205481565b34801561042d57600080fd5b506101c8610f0a565b34801561044257600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610224565b34801561047f57600080fd5b506101ef61048e36600461179a565b610f1e565b34801561049f57600080fd5b506101ef6104ae366004611800565b610f56565b3480156104bf57600080fd5b506102d3600081565b3480156104d457600080fd5b506101c86104e33660046116f7565b610fe3565b3480156104f457600080fd5b506101c861050336600461182c565b611086565b34801561051457600080fd5b506102d36105233660046116f7565b3360009081526004602090815260408083206001600160a01b03949094168352929052205490565b34801561055757600080fd5b506101c861056636600461179a565b6111c8565b34801561057757600080fd5b50600154610224906001600160a01b031681565b34801561059757600080fd5b506101c86105a63660046116f7565b6111e4565b6101ef5b60006105b961121f565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee600090815260036020527f40c35f83c179e47de306c9fe55fdc60064f11dd52adb51ab61b5643ee626f98d805434929061060990849061188d565b9091555050604051348152339073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee907f4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd579060200160405180910390a350600161068560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b90565b60006001600160e01b03198216637965db0b60e01b14806106b957506301ffc9a760e01b6001600160e01b03198316145b92915050565b600281815481106106cf57600080fd5b6000918252602090912001546001600160a01b0316905081565b6106f1611269565b600180546001600160a01b0319166001600160a01b0383169081179091556040517f799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b990600090a250565b60005b6002548110156108b45760006002828154811061075d5761075d6118a0565b60009182526020808320909101543383526004825260408084206001600160a01b039092168085529190925291205490915080156108aa5773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b0383160161085357604051600090339083908381818185875af1925050503d80600081146107fc576040519150601f19603f3d011682016040523d82523d6000602084013e610801565b606091505b505090508061084d5760405162461bcd60e51b8152602060048201526013602482015272115512081d1c985b9cd9995c8819985a5b1959606a1b60448201526064015b60405180910390fd5b50610867565b6108676001600160a01b03831633836112c4565b3360009081526004602090815260408083206001600160a01b038616845282528083208390556003909152812080548392906108a49084906118b6565b90915550505b505060010161073e565b50565b6060600280548060200260200160405190810160405280929190818152602001828054801561090f57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116108f1575b5050505050905090565b600090815260008051602061195e833981519152602052604090206001015490565b61094482610919565b61094d81611323565b610957838361132d565b50505050565b6001600160a01b0381166109ab5760405162461bcd60e51b8152602060048201526015602482015274496e76616c696420746f6b656e206164647265737360581b6044820152606401610844565b3360009081526004602090815260408083206001600160a01b038516845290915290205480610a125760405162461bcd60e51b81526020600482015260136024820152724e6f2072657761726420617661696c61626c6560681b6044820152606401610844565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b03831601610acb57604051600090339083908381818185875af1925050503d8060008114610a79576040519150601f19603f3d011682016040523d82523d6000602084013e610a7e565b606091505b5050905080610ac55760405162461bcd60e51b8152602060048201526013602482015272115512081d1c985b9cd9995c8819985a5b1959606a1b6044820152606401610844565b50610adf565b610adf6001600160a01b03831633836112c4565b3360009081526004602090815260408083206001600160a01b03861684528252808320839055600390915281208054839290610b1c9084906118b6565b90915550505050565b6001600160a01b0381163314610b4e5760405163334bd91960e11b815260040160405180910390fd5b610b5882826113d2565b505050565b6000546001600160a01b03163314610b875760405162461bcd60e51b8152600401610844906118c9565b6001600160a01b03831615801590610ba757506001600160a01b03821615155b610be55760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b6044820152606401610844565b6001600160a01b03808316600090815260046020908152604080832093871683529290529081208054839290610c1c90849061188d565b9091555050604080516001600160a01b038481168252602082018490528516917f10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424910160405180910390a2505050565b6001546000906001600160a01b03163314610cc95760405162461bcd60e51b815260206004820152601e60248201527f54726561737572654d616e616765722e6f6e6c795769746864726177657200006044820152606401610844565b6001600160a01b038416600090815260036020526040902054821115610d405760405162461bcd60e51b815260206004820152602660248201527f496e73756666696369656e7420746f6b656e2062616c616e636520696e20636f6044820152651b9d1c9858dd60d21b6064820152608401610844565b610d546001600160a01b03851684846112c4565b6001600160a01b03841660009081526003602052604081208054849290610d7c9084906118b6565b9091555060019150505b9392505050565b6001546000906001600160a01b03163314610dea5760405162461bcd60e51b815260206004820152601e60248201527f54726561737572654d616e616765722e6f6e6c795769746864726177657200006044820152606401610844565b81471015610e465760405162461bcd60e51b8152602060048201526024808201527f496e73756666696369656e74204554482062616c616e636520696e20636f6e746044820152631c9858dd60e21b6064820152608401610844565b6000836001600160a01b03168360405160006040518083038185875af1925050503d8060008114610e93576040519150601f19603f3d011682016040523d82523d6000602084013e610e98565b606091505b5050905080610eab5760009150506106b9565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee600090815260036020527f40c35f83c179e47de306c9fe55fdc60064f11dd52adb51ab61b5643ee626f98d8054859290610efb9084906118b6565b90915550600195945050505050565b610f12611269565b610f1c600061144e565b565b600091825260008051602061195e833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6000610f6d6001600160a01b0384163330856114bf565b6001600160a01b03831660009081526003602052604081208054849290610f9590849061188d565b909155505060405182815233906001600160a01b038516907f4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd579060200160405180910390a350600192915050565b6000546001600160a01b0316331461100d5760405162461bcd60e51b8152600401610844906118c9565b6001600160a01b03811661103457604051631326d6d560e01b815260040160405180910390fd5b600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b0392909216919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156110cc5750825b905060008267ffffffffffffffff1660011480156110e95750303b155b9050811580156110f7575080155b156111155760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561113f57845460ff60401b1916600160401b1785555b600080546001600160a01b03808a166001600160a01b03199283161790925560018054928916929091169190911790556111788861144e565b83156111be57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6111d182610919565b6111da81611323565b61095783836113d2565b6111ec611269565b6001600160a01b03811661121657604051631e4fbdf760e01b815260006004820152602401610844565b6108b48161144e565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080546001190161126357604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b3361129b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610f1c5760405163118cdaa760e01b8152336004820152602401610844565b6040516001600160a01b03838116602483015260448201839052610b5891859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506114f8565b6108b4813361155b565b600060008051602061195e8339815191526113488484610f1e565b6113c8576000848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561137e3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506106b9565b60009150506106b9565b600060008051602061195e8339815191526113ed8484610f1e565b156113c8576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506106b9565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6040516001600160a01b0384811660248301528381166044830152606482018390526109579186918216906323b872dd906084016112f1565b600061150d6001600160a01b03841683611598565b90508051600014158015611532575080806020019051810190611530919061190c565b155b15610b5857604051635274afe760e01b81526001600160a01b0384166004820152602401610844565b6115658282610f1e565b6115945760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610844565b5050565b6060610d868383600084600080856001600160a01b031684866040516115be919061192e565b60006040518083038185875af1925050503d80600081146115fb576040519150601f19603f3d011682016040523d82523d6000602084013e611600565b606091505b509150915061161086838361161a565b9695505050505050565b60608261162f5761162a82611676565b610d86565b815115801561164657506001600160a01b0384163b155b1561166f57604051639996b31560e01b81526001600160a01b0385166004820152602401610844565b5080610d86565b8051156116865780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b6000602082840312156116b157600080fd5b81356001600160e01b031981168114610d8657600080fd5b6000602082840312156116db57600080fd5b5035919050565b6001600160a01b03811681146108b457600080fd5b60006020828403121561170957600080fd5b8135610d86816116e2565b6020808252825182820181905260009190848201906040850190845b818110156117555783516001600160a01b031683529284019291840191600101611730565b50909695505050505050565b6000806040838503121561177457600080fd5b823561177f816116e2565b9150602083013561178f816116e2565b809150509250929050565b600080604083850312156117ad57600080fd5b82359150602083013561178f816116e2565b6000806000606084860312156117d457600080fd5b83356117df816116e2565b925060208401356117ef816116e2565b929592945050506040919091013590565b6000806040838503121561181357600080fd5b823561181e816116e2565b946020939093013593505050565b60008060006060848603121561184157600080fd5b833561184c816116e2565b9250602084013561185c816116e2565b9150604084013561186c816116e2565b809150509250925092565b634e487b7160e01b600052601160045260246000fd5b808201808211156106b9576106b9611877565b634e487b7160e01b600052603260045260246000fd5b818103818111156106b9576106b9611877565b60208082526023908201527f54726561737572654d616e616765722e6f6e6c7954726561737572654d616e6160408201526233b2b960e91b606082015260800190565b60006020828403121561191e57600080fd5b81518015158114610d8657600080fd5b6000825160005b8181101561194f5760208186018101518583015201611935565b50600092019182525091905056fe02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a26469706673582212205bece7d125e29d8ae848aefd08d15c076430660f12d074ad8423a9deb03bf26464736f6c63430008180033",
}

// TreasureManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TreasureManagerMetaData.ABI instead.
var TreasureManagerABI = TreasureManagerMetaData.ABI

// TreasureManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TreasureManagerMetaData.Bin instead.
var TreasureManagerBin = TreasureManagerMetaData.Bin

// DeployTreasureManager deploys a new Ethereum contract, binding an instance of TreasureManager to it.
func DeployTreasureManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TreasureManager, error) {
	parsed, err := TreasureManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TreasureManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TreasureManager{TreasureManagerCaller: TreasureManagerCaller{contract: contract}, TreasureManagerTransactor: TreasureManagerTransactor{contract: contract}, TreasureManagerFilterer: TreasureManagerFilterer{contract: contract}}, nil
}

// TreasureManager is an auto generated Go binding around an Ethereum contract.
type TreasureManager struct {
	TreasureManagerCaller     // Read-only binding to the contract
	TreasureManagerTransactor // Write-only binding to the contract
	TreasureManagerFilterer   // Log filterer for contract events
}

// TreasureManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TreasureManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasureManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasureManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasureManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TreasureManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasureManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasureManagerSession struct {
	Contract     *TreasureManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasureManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasureManagerCallerSession struct {
	Contract *TreasureManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TreasureManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasureManagerTransactorSession struct {
	Contract     *TreasureManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TreasureManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TreasureManagerRaw struct {
	Contract *TreasureManager // Generic contract binding to access the raw methods on
}

// TreasureManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasureManagerCallerRaw struct {
	Contract *TreasureManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TreasureManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasureManagerTransactorRaw struct {
	Contract *TreasureManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasureManager creates a new instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManager(address common.Address, backend bind.ContractBackend) (*TreasureManager, error) {
	contract, err := bindTreasureManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TreasureManager{TreasureManagerCaller: TreasureManagerCaller{contract: contract}, TreasureManagerTransactor: TreasureManagerTransactor{contract: contract}, TreasureManagerFilterer: TreasureManagerFilterer{contract: contract}}, nil
}

// NewTreasureManagerCaller creates a new read-only instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManagerCaller(address common.Address, caller bind.ContractCaller) (*TreasureManagerCaller, error) {
	contract, err := bindTreasureManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerCaller{contract: contract}, nil
}

// NewTreasureManagerTransactor creates a new write-only instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasureManagerTransactor, error) {
	contract, err := bindTreasureManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerTransactor{contract: contract}, nil
}

// NewTreasureManagerFilterer creates a new log filterer instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TreasureManagerFilterer, error) {
	contract, err := bindTreasureManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerFilterer{contract: contract}, nil
}

// bindTreasureManager binds a generic wrapper to an already deployed contract.
func bindTreasureManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TreasureManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasureManager *TreasureManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasureManager.Contract.TreasureManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasureManager *TreasureManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.Contract.TreasureManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasureManager *TreasureManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasureManager.Contract.TreasureManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasureManager *TreasureManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasureManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasureManager *TreasureManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasureManager *TreasureManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasureManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasureManager *TreasureManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasureManager *TreasureManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TreasureManager.Contract.DEFAULTADMINROLE(&_TreasureManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasureManager *TreasureManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TreasureManager.Contract.DEFAULTADMINROLE(&_TreasureManager.CallOpts)
}

// EthAddress is a free data retrieval call binding the contract method 0x41398b15.
//
// Solidity: function ethAddress() view returns(address)
func (_TreasureManager *TreasureManagerCaller) EthAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "ethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthAddress is a free data retrieval call binding the contract method 0x41398b15.
//
// Solidity: function ethAddress() view returns(address)
func (_TreasureManager *TreasureManagerSession) EthAddress() (common.Address, error) {
	return _TreasureManager.Contract.EthAddress(&_TreasureManager.CallOpts)
}

// EthAddress is a free data retrieval call binding the contract method 0x41398b15.
//
// Solidity: function ethAddress() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) EthAddress() (common.Address, error) {
	return _TreasureManager.Contract.EthAddress(&_TreasureManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasureManager *TreasureManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasureManager *TreasureManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TreasureManager.Contract.GetRoleAdmin(&_TreasureManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasureManager *TreasureManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TreasureManager.Contract.GetRoleAdmin(&_TreasureManager.CallOpts, role)
}

// GetTokenWhiteList is a free data retrieval call binding the contract method 0x23ecdbee.
//
// Solidity: function getTokenWhiteList() view returns(address[])
func (_TreasureManager *TreasureManagerCaller) GetTokenWhiteList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "getTokenWhiteList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenWhiteList is a free data retrieval call binding the contract method 0x23ecdbee.
//
// Solidity: function getTokenWhiteList() view returns(address[])
func (_TreasureManager *TreasureManagerSession) GetTokenWhiteList() ([]common.Address, error) {
	return _TreasureManager.Contract.GetTokenWhiteList(&_TreasureManager.CallOpts)
}

// GetTokenWhiteList is a free data retrieval call binding the contract method 0x23ecdbee.
//
// Solidity: function getTokenWhiteList() view returns(address[])
func (_TreasureManager *TreasureManagerCallerSession) GetTokenWhiteList() ([]common.Address, error) {
	return _TreasureManager.Contract.GetTokenWhiteList(&_TreasureManager.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasureManager *TreasureManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasureManager *TreasureManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TreasureManager.Contract.HasRole(&_TreasureManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasureManager *TreasureManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TreasureManager.Contract.HasRole(&_TreasureManager.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasureManager *TreasureManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasureManager *TreasureManagerSession) Owner() (common.Address, error) {
	return _TreasureManager.Contract.Owner(&_TreasureManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) Owner() (common.Address, error) {
	return _TreasureManager.Contract.Owner(&_TreasureManager.CallOpts)
}

// QueryReward is a free data retrieval call binding the contract method 0xc37220ea.
//
// Solidity: function queryReward(address _tokenAddress) view returns(uint256)
func (_TreasureManager *TreasureManagerCaller) QueryReward(opts *bind.CallOpts, _tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "queryReward", _tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryReward is a free data retrieval call binding the contract method 0xc37220ea.
//
// Solidity: function queryReward(address _tokenAddress) view returns(uint256)
func (_TreasureManager *TreasureManagerSession) QueryReward(_tokenAddress common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.QueryReward(&_TreasureManager.CallOpts, _tokenAddress)
}

// QueryReward is a free data retrieval call binding the contract method 0xc37220ea.
//
// Solidity: function queryReward(address _tokenAddress) view returns(uint256)
func (_TreasureManager *TreasureManagerCallerSession) QueryReward(_tokenAddress common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.QueryReward(&_TreasureManager.CallOpts, _tokenAddress)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasureManager *TreasureManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasureManager *TreasureManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TreasureManager.Contract.SupportsInterface(&_TreasureManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasureManager *TreasureManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TreasureManager.Contract.SupportsInterface(&_TreasureManager.CallOpts, interfaceId)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_TreasureManager *TreasureManagerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.TokenBalances(&_TreasureManager.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.TokenBalances(&_TreasureManager.CallOpts, arg0)
}

// TokenWhiteList is a free data retrieval call binding the contract method 0x0ec08b00.
//
// Solidity: function tokenWhiteList(uint256 ) view returns(address)
func (_TreasureManager *TreasureManagerCaller) TokenWhiteList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "tokenWhiteList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenWhiteList is a free data retrieval call binding the contract method 0x0ec08b00.
//
// Solidity: function tokenWhiteList(uint256 ) view returns(address)
func (_TreasureManager *TreasureManagerSession) TokenWhiteList(arg0 *big.Int) (common.Address, error) {
	return _TreasureManager.Contract.TokenWhiteList(&_TreasureManager.CallOpts, arg0)
}

// TokenWhiteList is a free data retrieval call binding the contract method 0x0ec08b00.
//
// Solidity: function tokenWhiteList(uint256 ) view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) TokenWhiteList(arg0 *big.Int) (common.Address, error) {
	return _TreasureManager.Contract.TokenWhiteList(&_TreasureManager.CallOpts, arg0)
}

// TreasureManager is a free data retrieval call binding the contract method 0x03186d22.
//
// Solidity: function treasureManager() view returns(address)
func (_TreasureManager *TreasureManagerCaller) TreasureManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "treasureManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasureManager is a free data retrieval call binding the contract method 0x03186d22.
//
// Solidity: function treasureManager() view returns(address)
func (_TreasureManager *TreasureManagerSession) TreasureManager() (common.Address, error) {
	return _TreasureManager.Contract.TreasureManager(&_TreasureManager.CallOpts)
}

// TreasureManager is a free data retrieval call binding the contract method 0x03186d22.
//
// Solidity: function treasureManager() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) TreasureManager() (common.Address, error) {
	return _TreasureManager.Contract.TreasureManager(&_TreasureManager.CallOpts)
}

// UserRewardAmounts is a free data retrieval call binding the contract method 0x254c5d87.
//
// Solidity: function userRewardAmounts(address , address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCaller) UserRewardAmounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "userRewardAmounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardAmounts is a free data retrieval call binding the contract method 0x254c5d87.
//
// Solidity: function userRewardAmounts(address , address ) view returns(uint256)
func (_TreasureManager *TreasureManagerSession) UserRewardAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.UserRewardAmounts(&_TreasureManager.CallOpts, arg0, arg1)
}

// UserRewardAmounts is a free data retrieval call binding the contract method 0x254c5d87.
//
// Solidity: function userRewardAmounts(address , address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCallerSession) UserRewardAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.UserRewardAmounts(&_TreasureManager.CallOpts, arg0, arg1)
}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_TreasureManager *TreasureManagerCaller) WithdrawManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "withdrawManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_TreasureManager *TreasureManagerSession) WithdrawManager() (common.Address, error) {
	return _TreasureManager.Contract.WithdrawManager(&_TreasureManager.CallOpts)
}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) WithdrawManager() (common.Address, error) {
	return _TreasureManager.Contract.WithdrawManager(&_TreasureManager.CallOpts)
}

// ClaimAllTokens is a paid mutator transaction binding the contract method 0x1e4bd42c.
//
// Solidity: function claimAllTokens() returns()
func (_TreasureManager *TreasureManagerTransactor) ClaimAllTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "claimAllTokens")
}

// ClaimAllTokens is a paid mutator transaction binding the contract method 0x1e4bd42c.
//
// Solidity: function claimAllTokens() returns()
func (_TreasureManager *TreasureManagerSession) ClaimAllTokens() (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimAllTokens(&_TreasureManager.TransactOpts)
}

// ClaimAllTokens is a paid mutator transaction binding the contract method 0x1e4bd42c.
//
// Solidity: function claimAllTokens() returns()
func (_TreasureManager *TreasureManagerTransactorSession) ClaimAllTokens() (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimAllTokens(&_TreasureManager.TransactOpts)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x32f289cf.
//
// Solidity: function claimToken(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactor) ClaimToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "claimToken", tokenAddress)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x32f289cf.
//
// Solidity: function claimToken(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerSession) ClaimToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimToken(&_TreasureManager.TransactOpts, tokenAddress)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x32f289cf.
//
// Solidity: function claimToken(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactorSession) ClaimToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimToken(&_TreasureManager.TransactOpts, tokenAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address tokenAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerTransactor) DepositERC20(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "depositERC20", tokenAddress, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address tokenAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerSession) DepositERC20(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositERC20(&_TreasureManager.TransactOpts, tokenAddress, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address tokenAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) DepositERC20(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositERC20(&_TreasureManager.TransactOpts, tokenAddress, amount)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns(bool)
func (_TreasureManager *TreasureManagerTransactor) DepositETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "depositETH")
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns(bool)
func (_TreasureManager *TreasureManagerSession) DepositETH() (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositETH(&_TreasureManager.TransactOpts)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) DepositETH() (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositETH(&_TreasureManager.TransactOpts)
}

// GrantRewards is a paid mutator transaction binding the contract method 0x410c9f9a.
//
// Solidity: function grantRewards(address tokenAddress, address granter, uint256 amount) returns()
func (_TreasureManager *TreasureManagerTransactor) GrantRewards(opts *bind.TransactOpts, tokenAddress common.Address, granter common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "grantRewards", tokenAddress, granter, amount)
}

// GrantRewards is a paid mutator transaction binding the contract method 0x410c9f9a.
//
// Solidity: function grantRewards(address tokenAddress, address granter, uint256 amount) returns()
func (_TreasureManager *TreasureManagerSession) GrantRewards(tokenAddress common.Address, granter common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantRewards(&_TreasureManager.TransactOpts, tokenAddress, granter, amount)
}

// GrantRewards is a paid mutator transaction binding the contract method 0x410c9f9a.
//
// Solidity: function grantRewards(address tokenAddress, address granter, uint256 amount) returns()
func (_TreasureManager *TreasureManagerTransactorSession) GrantRewards(tokenAddress common.Address, granter common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantRewards(&_TreasureManager.TransactOpts, tokenAddress, granter, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantRole(&_TreasureManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantRole(&_TreasureManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _initialOwner, address _treasureManager, address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _treasureManager common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "initialize", _initialOwner, _treasureManager, _withdrawManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _initialOwner, address _treasureManager, address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerSession) Initialize(_initialOwner common.Address, _treasureManager common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.Initialize(&_TreasureManager.TransactOpts, _initialOwner, _treasureManager, _withdrawManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _initialOwner, address _treasureManager, address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactorSession) Initialize(_initialOwner common.Address, _treasureManager common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.Initialize(&_TreasureManager.TransactOpts, _initialOwner, _treasureManager, _withdrawManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasureManager *TreasureManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasureManager *TreasureManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceOwnership(&_TreasureManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasureManager *TreasureManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceOwnership(&_TreasureManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasureManager *TreasureManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasureManager *TreasureManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceRole(&_TreasureManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasureManager *TreasureManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceRole(&_TreasureManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RevokeRole(&_TreasureManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RevokeRole(&_TreasureManager.TransactOpts, role, account)
}

// SetTokenWhiteList is a paid mutator transaction binding the contract method 0xa2672ad7.
//
// Solidity: function setTokenWhiteList(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactor) SetTokenWhiteList(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "setTokenWhiteList", tokenAddress)
}

// SetTokenWhiteList is a paid mutator transaction binding the contract method 0xa2672ad7.
//
// Solidity: function setTokenWhiteList(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerSession) SetTokenWhiteList(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetTokenWhiteList(&_TreasureManager.TransactOpts, tokenAddress)
}

// SetTokenWhiteList is a paid mutator transaction binding the contract method 0xa2672ad7.
//
// Solidity: function setTokenWhiteList(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactorSession) SetTokenWhiteList(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetTokenWhiteList(&_TreasureManager.TransactOpts, tokenAddress)
}

// SetWithdrawManager is a paid mutator transaction binding the contract method 0x17e3e2e8.
//
// Solidity: function setWithdrawManager(address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactor) SetWithdrawManager(opts *bind.TransactOpts, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "setWithdrawManager", _withdrawManager)
}

// SetWithdrawManager is a paid mutator transaction binding the contract method 0x17e3e2e8.
//
// Solidity: function setWithdrawManager(address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerSession) SetWithdrawManager(_withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetWithdrawManager(&_TreasureManager.TransactOpts, _withdrawManager)
}

// SetWithdrawManager is a paid mutator transaction binding the contract method 0x17e3e2e8.
//
// Solidity: function setWithdrawManager(address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactorSession) SetWithdrawManager(_withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetWithdrawManager(&_TreasureManager.TransactOpts, _withdrawManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasureManager *TreasureManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasureManager *TreasureManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.TransferOwnership(&_TreasureManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasureManager *TreasureManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.TransferOwnership(&_TreasureManager.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerTransactor) WithdrawERC20(opts *bind.TransactOpts, tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "withdrawERC20", tokenAddress, withdrawAddress, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerSession) WithdrawERC20(tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawERC20(&_TreasureManager.TransactOpts, tokenAddress, withdrawAddress, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) WithdrawERC20(tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawERC20(&_TreasureManager.TransactOpts, tokenAddress, withdrawAddress, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address withdrawAddress, uint256 amount) payable returns(bool)
func (_TreasureManager *TreasureManagerTransactor) WithdrawETH(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "withdrawETH", withdrawAddress, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address withdrawAddress, uint256 amount) payable returns(bool)
func (_TreasureManager *TreasureManagerSession) WithdrawETH(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawETH(&_TreasureManager.TransactOpts, withdrawAddress, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address withdrawAddress, uint256 amount) payable returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) WithdrawETH(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawETH(&_TreasureManager.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasureManager *TreasureManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasureManager *TreasureManagerSession) Receive() (*types.Transaction, error) {
	return _TreasureManager.Contract.Receive(&_TreasureManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasureManager *TreasureManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _TreasureManager.Contract.Receive(&_TreasureManager.TransactOpts)
}

// TreasureManagerDepositTokenIterator is returned from FilterDepositToken and is used to iterate over the raw logs and unpacked data for DepositToken events raised by the TreasureManager contract.
type TreasureManagerDepositTokenIterator struct {
	Event *TreasureManagerDepositToken // Event containing the contract specifics and raw log

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
func (it *TreasureManagerDepositTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerDepositToken)
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
		it.Event = new(TreasureManagerDepositToken)
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
func (it *TreasureManagerDepositTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerDepositTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerDepositToken represents a DepositToken event raised by the TreasureManager contract.
type TreasureManagerDepositToken struct {
	TokenAddress common.Address
	Sender       common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositToken is a free log retrieval operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) FilterDepositToken(opts *bind.FilterOpts, tokenAddress []common.Address, sender []common.Address) (*TreasureManagerDepositTokenIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "DepositToken", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerDepositTokenIterator{contract: _TreasureManager.contract, event: "DepositToken", logs: logs, sub: sub}, nil
}

// WatchDepositToken is a free log subscription operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) WatchDepositToken(opts *bind.WatchOpts, sink chan<- *TreasureManagerDepositToken, tokenAddress []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "DepositToken", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerDepositToken)
				if err := _TreasureManager.contract.UnpackLog(event, "DepositToken", log); err != nil {
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

// ParseDepositToken is a log parse operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) ParseDepositToken(log types.Log) (*TreasureManagerDepositToken, error) {
	event := new(TreasureManagerDepositToken)
	if err := _TreasureManager.contract.UnpackLog(event, "DepositToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerGrantRewardTokenAmountIterator is returned from FilterGrantRewardTokenAmount and is used to iterate over the raw logs and unpacked data for GrantRewardTokenAmount events raised by the TreasureManager contract.
type TreasureManagerGrantRewardTokenAmountIterator struct {
	Event *TreasureManagerGrantRewardTokenAmount // Event containing the contract specifics and raw log

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
func (it *TreasureManagerGrantRewardTokenAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerGrantRewardTokenAmount)
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
		it.Event = new(TreasureManagerGrantRewardTokenAmount)
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
func (it *TreasureManagerGrantRewardTokenAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerGrantRewardTokenAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerGrantRewardTokenAmount represents a GrantRewardTokenAmount event raised by the TreasureManager contract.
type TreasureManagerGrantRewardTokenAmount struct {
	TokenAddress common.Address
	Granter      common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGrantRewardTokenAmount is a free log retrieval operation binding the contract event 0x10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424.
//
// Solidity: event GrantRewardTokenAmount(address indexed tokenAddress, address granter, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) FilterGrantRewardTokenAmount(opts *bind.FilterOpts, tokenAddress []common.Address) (*TreasureManagerGrantRewardTokenAmountIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "GrantRewardTokenAmount", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerGrantRewardTokenAmountIterator{contract: _TreasureManager.contract, event: "GrantRewardTokenAmount", logs: logs, sub: sub}, nil
}

// WatchGrantRewardTokenAmount is a free log subscription operation binding the contract event 0x10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424.
//
// Solidity: event GrantRewardTokenAmount(address indexed tokenAddress, address granter, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) WatchGrantRewardTokenAmount(opts *bind.WatchOpts, sink chan<- *TreasureManagerGrantRewardTokenAmount, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "GrantRewardTokenAmount", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerGrantRewardTokenAmount)
				if err := _TreasureManager.contract.UnpackLog(event, "GrantRewardTokenAmount", log); err != nil {
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

// ParseGrantRewardTokenAmount is a log parse operation binding the contract event 0x10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424.
//
// Solidity: event GrantRewardTokenAmount(address indexed tokenAddress, address granter, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) ParseGrantRewardTokenAmount(log types.Log) (*TreasureManagerGrantRewardTokenAmount, error) {
	event := new(TreasureManagerGrantRewardTokenAmount)
	if err := _TreasureManager.contract.UnpackLog(event, "GrantRewardTokenAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TreasureManager contract.
type TreasureManagerInitializedIterator struct {
	Event *TreasureManagerInitialized // Event containing the contract specifics and raw log

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
func (it *TreasureManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerInitialized)
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
		it.Event = new(TreasureManagerInitialized)
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
func (it *TreasureManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerInitialized represents a Initialized event raised by the TreasureManager contract.
type TreasureManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TreasureManager *TreasureManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*TreasureManagerInitializedIterator, error) {

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TreasureManagerInitializedIterator{contract: _TreasureManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TreasureManager *TreasureManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TreasureManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerInitialized)
				if err := _TreasureManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TreasureManager *TreasureManagerFilterer) ParseInitialized(log types.Log) (*TreasureManagerInitialized, error) {
	event := new(TreasureManagerInitialized)
	if err := _TreasureManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TreasureManager contract.
type TreasureManagerOwnershipTransferredIterator struct {
	Event *TreasureManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TreasureManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerOwnershipTransferred)
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
		it.Event = new(TreasureManagerOwnershipTransferred)
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
func (it *TreasureManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerOwnershipTransferred represents a OwnershipTransferred event raised by the TreasureManager contract.
type TreasureManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TreasureManager *TreasureManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TreasureManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerOwnershipTransferredIterator{contract: _TreasureManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TreasureManager *TreasureManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TreasureManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerOwnershipTransferred)
				if err := _TreasureManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TreasureManager *TreasureManagerFilterer) ParseOwnershipTransferred(log types.Log) (*TreasureManagerOwnershipTransferred, error) {
	event := new(TreasureManagerOwnershipTransferred)
	if err := _TreasureManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TreasureManager contract.
type TreasureManagerRoleAdminChangedIterator struct {
	Event *TreasureManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TreasureManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerRoleAdminChanged)
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
		it.Event = new(TreasureManagerRoleAdminChanged)
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
func (it *TreasureManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerRoleAdminChanged represents a RoleAdminChanged event raised by the TreasureManager contract.
type TreasureManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasureManager *TreasureManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TreasureManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerRoleAdminChangedIterator{contract: _TreasureManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasureManager *TreasureManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TreasureManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerRoleAdminChanged)
				if err := _TreasureManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasureManager *TreasureManagerFilterer) ParseRoleAdminChanged(log types.Log) (*TreasureManagerRoleAdminChanged, error) {
	event := new(TreasureManagerRoleAdminChanged)
	if err := _TreasureManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TreasureManager contract.
type TreasureManagerRoleGrantedIterator struct {
	Event *TreasureManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *TreasureManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerRoleGranted)
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
		it.Event = new(TreasureManagerRoleGranted)
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
func (it *TreasureManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerRoleGranted represents a RoleGranted event raised by the TreasureManager contract.
type TreasureManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TreasureManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerRoleGrantedIterator{contract: _TreasureManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TreasureManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerRoleGranted)
				if err := _TreasureManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) ParseRoleGranted(log types.Log) (*TreasureManagerRoleGranted, error) {
	event := new(TreasureManagerRoleGranted)
	if err := _TreasureManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TreasureManager contract.
type TreasureManagerRoleRevokedIterator struct {
	Event *TreasureManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TreasureManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerRoleRevoked)
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
		it.Event = new(TreasureManagerRoleRevoked)
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
func (it *TreasureManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerRoleRevoked represents a RoleRevoked event raised by the TreasureManager contract.
type TreasureManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TreasureManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerRoleRevokedIterator{contract: _TreasureManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TreasureManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerRoleRevoked)
				if err := _TreasureManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) ParseRoleRevoked(log types.Log) (*TreasureManagerRoleRevoked, error) {
	event := new(TreasureManagerRoleRevoked)
	if err := _TreasureManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerWithdrawManagerUpdateIterator is returned from FilterWithdrawManagerUpdate and is used to iterate over the raw logs and unpacked data for WithdrawManagerUpdate events raised by the TreasureManager contract.
type TreasureManagerWithdrawManagerUpdateIterator struct {
	Event *TreasureManagerWithdrawManagerUpdate // Event containing the contract specifics and raw log

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
func (it *TreasureManagerWithdrawManagerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerWithdrawManagerUpdate)
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
		it.Event = new(TreasureManagerWithdrawManagerUpdate)
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
func (it *TreasureManagerWithdrawManagerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerWithdrawManagerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerWithdrawManagerUpdate represents a WithdrawManagerUpdate event raised by the TreasureManager contract.
type TreasureManagerWithdrawManagerUpdate struct {
	WithdrawManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawManagerUpdate is a free log retrieval operation binding the contract event 0x799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b9.
//
// Solidity: event WithdrawManagerUpdate(address indexed withdrawManager)
func (_TreasureManager *TreasureManagerFilterer) FilterWithdrawManagerUpdate(opts *bind.FilterOpts, withdrawManager []common.Address) (*TreasureManagerWithdrawManagerUpdateIterator, error) {

	var withdrawManagerRule []interface{}
	for _, withdrawManagerItem := range withdrawManager {
		withdrawManagerRule = append(withdrawManagerRule, withdrawManagerItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "WithdrawManagerUpdate", withdrawManagerRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerWithdrawManagerUpdateIterator{contract: _TreasureManager.contract, event: "WithdrawManagerUpdate", logs: logs, sub: sub}, nil
}

// WatchWithdrawManagerUpdate is a free log subscription operation binding the contract event 0x799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b9.
//
// Solidity: event WithdrawManagerUpdate(address indexed withdrawManager)
func (_TreasureManager *TreasureManagerFilterer) WatchWithdrawManagerUpdate(opts *bind.WatchOpts, sink chan<- *TreasureManagerWithdrawManagerUpdate, withdrawManager []common.Address) (event.Subscription, error) {

	var withdrawManagerRule []interface{}
	for _, withdrawManagerItem := range withdrawManager {
		withdrawManagerRule = append(withdrawManagerRule, withdrawManagerItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "WithdrawManagerUpdate", withdrawManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerWithdrawManagerUpdate)
				if err := _TreasureManager.contract.UnpackLog(event, "WithdrawManagerUpdate", log); err != nil {
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

// ParseWithdrawManagerUpdate is a log parse operation binding the contract event 0x799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b9.
//
// Solidity: event WithdrawManagerUpdate(address indexed withdrawManager)
func (_TreasureManager *TreasureManagerFilterer) ParseWithdrawManagerUpdate(log types.Log) (*TreasureManagerWithdrawManagerUpdate, error) {
	event := new(TreasureManagerWithdrawManagerUpdate)
	if err := _TreasureManager.contract.UnpackLog(event, "WithdrawManagerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
