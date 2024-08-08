package event_watcher

import (
	"context"

	"sync/atomic"

	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3/event-watcher/config"
	"github.com/the-web3/event-watcher/database"
	"github.com/the-web3/event-watcher/event"
	"github.com/the-web3/event-watcher/synchronizer"
	"github.com/the-web3/event-watcher/synchronizer/node"
)

type EventWatcher struct {
	synchronizer   *synchronizer.Synchronizer
	eventProcessor *event.EventProcessor

	shutdown context.CancelCauseFunc
	stopped  atomic.Bool
}

func NewEventWatcher(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*EventWatcher, error) {
	ethClient, err := node.DialEthClient(ctx, cfg.Chain.ChainRpcUrl)
	if err != nil {
		return nil, err
	}

	db, err := database.NewDB(ctx, cfg.MasterDB)
	if err != nil {
		log.Error("init database fail", err)
		return nil, err
	}

	syncer, err := synchronizer.NewSynchronizer(cfg, db, ethClient, shutdown)
	if err != nil {
		log.Error("new synchronizer fail", "err", err)
		return nil, err
	}

	eventProcessor, err := event.NewEventProcessor(db, cfg.Chain, shutdown)
	if err != nil {
		log.Error("new evet processor fail", "err", err)
		return nil, err
	}

	out := &EventWatcher{
		synchronizer:   syncer,
		eventProcessor: eventProcessor,
		shutdown:       shutdown,
	}
	return out, nil
}

func (ew *EventWatcher) Start(ctx context.Context) error {
	err := ew.synchronizer.Start()
	if err != nil {
		return err
	}
	err = ew.eventProcessor.Start()
	if err != nil {
		return err
	}
	return nil
}

func (ew *EventWatcher) Stop(ctx context.Context) error {
	err := ew.synchronizer.Close()
	if err != nil {
		return err
	}
	err = ew.eventProcessor.Close()
	if err != nil {
		return err
	}
	return nil
}

func (ew *EventWatcher) Stopped() bool {
	return ew.stopped.Load()
}
