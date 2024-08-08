package dapplink

import (
	"github.com/google/uuid"
	"math/big"

	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3/event-watcher/config"
	"github.com/the-web3/event-watcher/database"
	"github.com/the-web3/event-watcher/database/worker"
	"github.com/the-web3/event-watcher/event/contracts"
)

func ProcessDepositEvents(db *database.DB, chainCfg config.ChainConfig, fromHeight, toHeight *big.Int) error {
	txDepositTokens, err := contracts.DepositTokensEvents(chainCfg.Contracts[0], db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(txDepositTokens) > 0 {
		log.Info("detected transaction deposits", "size", len(txDepositTokens))
	}

	depositTk := make([]worker.DepositTokens, len(txDepositTokens))
	for i := range txDepositTokens {
		depositTk[i] = worker.DepositTokens{
			GUID:         uuid.New(),
			TokenAddress: depositTk[i].TokenAddress,
			Sender:       txDepositTokens[i].Sender,
			Amount:       big.NewInt(0),
			Timestamp:    txDepositTokens[i].Timestamp,
		}
	}
	if len(depositTk) > 0 {
		if err := db.DepositTokens.StoreDepositTokens(depositTk); err != nil {
			return err
		}
	}
	return nil
}
