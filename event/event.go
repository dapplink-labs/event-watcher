package event

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3/event-watcher/common/bigint"
	"github.com/the-web3/event-watcher/common/tasks"
	"github.com/the-web3/event-watcher/config"
	"github.com/the-web3/event-watcher/database"
	"github.com/the-web3/event-watcher/database/common"
	"github.com/the-web3/event-watcher/event/dapplink"
)

var blocksLimit = 10_000

type EventProcessor struct {
	db          *database.DB
	chainConfig config.ChainConfig

	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group

	LatestBlockHeader *common.BlockHeader
}

func NewEventProcessor(db *database.DB, chainConfig config.ChainConfig, shutdown context.CancelCauseFunc) (*EventProcessor, error) {
	LatestBlockHeader, err := db.Blocks.LatestBlockHeader()
	if err != nil {
		return nil, err
	}

	resCtx, resCancel := context.WithCancel(context.Background())

	return &EventProcessor{
		db:             db,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		chainConfig:    chainConfig,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		LatestBlockHeader: LatestBlockHeader,
	}, nil
}

func (ep *EventProcessor) Start() error {
	log.Info("starting bridge processor...")
	tickerL1Worker := time.NewTicker(time.Second * 5)
	ep.tasks.Go(func() error {
		for range tickerL1Worker.C {

		}
		return nil
	})
	return nil
}

func (ep *EventProcessor) Close() error {
	ep.resourceCancel()
	return ep.tasks.Wait()
}

func (ep *EventProcessor) processTreasureManagerEvents() error {
	log.Info("bridge", "l1", "kind", "initiated")
	lastBlockNumber := big.NewInt(int64(ep.chainConfig.StartingHeight))
	if ep.LatestBlockHeader != nil {
		lastBlockNumber = ep.LatestBlockHeader.Number
	}
	log.Info("Process init l1 event", "lastBlockNumber", lastBlockNumber)

	latestHeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common.BlockHeader{}).Where("number > ?", lastBlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}

	latestHeader, err := ep.db.Blocks.BlockHeaderWithScope(latestHeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query new L1 state: %w", err)
	} else if latestHeader == nil {
		log.Debug("no new L1 state found")
		return nil
	}
	fromHeight, toHeight := new(big.Int).Add(lastBlockNumber, bigint.One), latestHeader.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		log.Info("scanning for initiated bridge events", "fromHeight", fromHeight, "toHeight", toHeight)
		return dapplink.ProcessDepositEvents(tx, ep.chainConfig, fromHeight, toHeight)

	}); err != nil {
		return err
	}
	ep.LatestBlockHeader = latestHeader
	return nil
}
