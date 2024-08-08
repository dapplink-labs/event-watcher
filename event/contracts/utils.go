package contracts

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func DecodeVersionedNonce(nonce *big.Int) (uint16, *big.Int) {
	nonceBytes := nonce.Bytes()
	nonceByteLen := len(nonceBytes)
	if nonceByteLen < 30 {
		return 0, nonce
	} else if nonceByteLen == 31 {
		return uint16(nonceBytes[0]), new(big.Int).SetBytes(nonceBytes[1:])
	} else {
		version := binary.BigEndian.Uint16(nonceBytes[:2])
		return version, new(big.Int).SetBytes(nonceBytes[2:])
	}
}

func UnpackLog(out interface{}, log *types.Log, name string, contractAbi *abi.ABI) error {
	eventAbi, ok := contractAbi.Events[name]
	if !ok {
		return fmt.Errorf("event %s not present in supplied ABI", name)
	} else if len(log.Topics) == 0 {
		return errors.New("anonymous events are not supported")
	} else if log.Topics[0] != eventAbi.ID {
		return errors.New("event signature mismatch")
	}

	err := contractAbi.UnpackIntoInterface(out, name, log.Data)
	if err != nil {
		return err
	}

	if len(log.Topics) > 1 {
		var indexedArgs abi.Arguments
		for _, arg := range eventAbi.Inputs {
			if arg.Indexed {
				indexedArgs = append(indexedArgs, arg)
			}
		}

		err := abi.ParseTopics(out, indexedArgs, log.Topics[1:])
		if err != nil {
			return err
		}
	}

	return nil
}
