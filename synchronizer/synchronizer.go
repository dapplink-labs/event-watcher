package synchronizer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3/event-watcher/common/tasks"
	"github.com/the-web3/event-watcher/config"
	"github.com/the-web3/event-watcher/database"
	common2 "github.com/the-web3/event-watcher/database/common"
	"github.com/the-web3/event-watcher/database/event"
	"github.com/the-web3/event-watcher/database/utils"
	"github.com/the-web3/event-watcher/synchronizer/node"
	"github.com/the-web3/event-watcher/synchronizer/retry"
)

type Synchronizer struct {
	ethClient node.EthClient
	db        *database.DB

	loopInterval     time.Duration
	headerBufferSize uint64
	headerTraversal  *node.HeaderTraversal

	headers      []types.Header
	latestHeader *types.Header

	startHeight       *big.Int
	confirmationDepth *big.Int
	chainCfg          *config.ChainConfig

	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
}

func NewSynchronizer(cfg *config.Config, db *database.DB, client node.EthClient, shutdown context.CancelCauseFunc) (*Synchronizer, error) {
	latestHeader, err := db.Blocks.LatestBlockHeader()
	if err != nil {
		return nil, err
	}
	var fromHeader *types.Header
	if latestHeader != nil {
		log.Info("sync detected last indexed block", "number", latestHeader.Number, "hash", latestHeader.Hash)
		fromHeader = latestHeader.RLPHeader.Header()
	} else if cfg.Chain.BlockStep > 0 {
		log.Info("no sync indexed state starting from supplied ethereum height", "height", cfg.Chain.StartingHeight)
		header, err := client.BlockHeaderByNumber(big.NewInt(int64(cfg.Chain.StartingHeight)))
		if err != nil {
			return nil, fmt.Errorf("could not fetch starting block header: %w", err)
		}
		fromHeader = header
	} else {
		log.Info("no eth wallet indexed state")
	}

	headerTraversal := node.NewHeaderTraversal(client, fromHeader, big.NewInt(0), cfg.Chain.ChainId)

	resCtx, resCancel := context.WithCancel(context.Background())
	return &Synchronizer{
		loopInterval:     time.Duration(cfg.Chain.LoopInterval) * time.Second,
		headerBufferSize: uint64(cfg.Chain.BlockStep),
		headerTraversal:  headerTraversal,
		ethClient:        client,
		latestHeader:     fromHeader,
		db:               db,
		chainCfg:         &cfg.Chain,
		resourceCtx:      resCtx,
		resourceCancel:   resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in Synchronizer: %w", err))
		}},
	}, nil
}

func (syncer *Synchronizer) Start() error {
	// tickerSyncer := time.NewTicker(syncer.loopInterval)
	tickerSyncer := time.NewTicker(time.Second * 3)
	syncer.tasks.Go(func() error {
		for range tickerSyncer.C {
			if len(syncer.headers) > 0 {
				log.Info("retrying previous batch")
			} else {
				newHeaders, err := syncer.headerTraversal.NextHeaders(uint64(syncer.chainCfg.BlockStep))
				if err != nil {
					log.Error("error querying for headers", "err", err)
					continue
				} else if len(newHeaders) == 0 {
					log.Warn("no new headers. syncer at head?")
				} else {
					syncer.headers = newHeaders
				}
				latestHeader := syncer.headerTraversal.LatestHeader()
				if latestHeader != nil {
					log.Info("Latest header", "latestHeader Number", latestHeader.Number)
				}
			}
			err := syncer.processBatch(syncer.headers, syncer.chainCfg)
			if err == nil {
				syncer.headers = nil
			}
		}
		return nil
	})
	return nil
}

func (syncer *Synchronizer) processBatch(headers []types.Header, chainCfg *config.ChainConfig) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHeader := headers[0], headers[len(headers)-1]
	log.Info("extracting batch", "size", len(headers), "startBlock", firstHeader.Number.String(), "endBlock", lastHeader.Number.String())

	headerMap := make(map[common.Hash]*types.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Hash()] = &header
	}
	log.Info("chainCfg Contracts", "contract address", chainCfg.Contracts[0])
	filterQuery := ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHeader.Number, Addresses: chainCfg.Contracts}
	logs, err := syncer.ethClient.FilterLogs(filterQuery)
	if err != nil {
		log.Info("failed to extract logs", "err", err)
		return err
	}

	if logs.ToBlockHeader.Number.Cmp(lastHeader.Number) != 0 {
		return fmt.Errorf("mismatch in FilterLog#ToBlock number")
	} else if logs.ToBlockHeader.Hash() != lastHeader.Hash() {
		return fmt.Errorf("mismatch in FitlerLog#ToBlock block hash")
	}

	if len(logs.Logs) > 0 {
		log.Info("detected logs", "size", len(logs.Logs))
	}

	blockHeaders := make([]common2.BlockHeader, 0, len(headers))
	for i := range headers {
		if headers[i].Number == nil {
			continue
		}
		bHeader := common2.BlockHeader{
			Hash:       headers[i].Hash(),
			ParentHash: headers[i].ParentHash,
			Number:     headers[i].Number,
			Timestamp:  headers[i].Time,
			RLPHeader:  (*utils.RLPHeader)(&headers[i]),
		}
		blockHeaders = append(blockHeaders, bHeader)
	}

	chainContractEvent := make([]event.ContractEvent, len(logs.Logs))
	for i := range logs.Logs {
		logEvent := logs.Logs[i]
		if _, ok := headerMap[logEvent.BlockHash]; !ok {
			continue
		}
		timestamp := headerMap[logEvent.BlockHash].Time
		chainContractEvent[i] = event.ContractEventFromLog(&logs.Logs[i], timestamp)
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](syncer.resourceCtx, 10, retryStrategy, func() (interface{}, error) {
		if err := syncer.db.Transaction(func(tx *database.DB) error {
			if err := tx.Blocks.StoreBlockHeaders(blockHeaders); err != nil {
				return err
			}
			if err := tx.ContractEvent.StoreContractEvents(chainContractEvent); err != nil {
				return err
			}
			return nil
		}); err != nil {
			log.Info("unable to persist batch", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func (syncer *Synchronizer) Close() error {
	return nil
}
