package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/the-web3/event-watcher/bindings"
	"github.com/the-web3/event-watcher/database"
	"github.com/the-web3/event-watcher/database/event"
)

type DepositTokensEvent struct {
	TokenAddress common.Address
	Sender       common.Address
	amount       *big.Int
	Timestamp    uint64
	Raw          types.Log
}

func DepositTokensEvents(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]DepositTokensEvent, error) {
	treasureManagerAbi, err := bindings.TreasureManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	transactionDepositTokenEventAbi := treasureManagerAbi.Events["DepositToken"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: transactionDepositTokenEventAbi.ID}
	transactionDepositTokenEvents, err := db.ContractEvent.ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	txDepositTokens := make([]DepositTokensEvent, len(transactionDepositTokenEvents))
	for i := range transactionDepositTokenEvents {
		depositTokens := bindings.TreasureManagerDepositToken{Raw: *transactionDepositTokenEvents[i].RLPLog}
		err := UnpackLog(&depositTokens, transactionDepositTokenEvents[i].RLPLog, transactionDepositTokenEventAbi.Name, treasureManagerAbi)
		if err != nil {
			return nil, err
		}
		txDepositTokens[i] = DepositTokensEvent{
			TokenAddress: depositTokens.TokenAddress,
			Sender:       depositTokens.Sender,
			amount:       depositTokens.Amount,
			Timestamp:    transactionDepositTokenEvents[i].Timestamp,
		}
	}
	return txDepositTokens, nil
}
