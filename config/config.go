package config

import (
	"github.com/ethereum/go-ethereum/common"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3/event-watcher/flags"
)

const (
	defaultConfirmations = 64
	defaultLoopInterval  = 5000
)

type Config struct {
	Migrations     string
	Chain          ChainConfig
	MasterDB       DBConfig
	SlaveDB        DBConfig
	SlaveDbEnable  bool
	ApiCacheEnable bool
	HTTPServer     ServerConfig
}

type ChainConfig struct {
	ChainRpcUrl    string
	ChainId        uint
	StartingHeight uint64
	Confirmations  uint64
	BlockStep      uint64
	Contracts      []common.Address
	LoopInterval   time.Duration
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type ServerConfig struct {
	Host string
	Port int
}

func LoadConfig(cliCtx *cli.Context) (Config, error) {
	var cfg Config
	cfg = NewConfig(cliCtx)

	if cfg.Chain.Confirmations == 0 {
		cfg.Chain.Confirmations = defaultConfirmations
	}

	if cfg.Chain.LoopInterval == 0 {
		cfg.Chain.LoopInterval = defaultLoopInterval
	}

	log.Info("loaded chain config", "config", cfg.Chain)
	return cfg, nil
}

func LoadContracts() []common.Address {
	var Contracts []common.Address
	Contracts = append(Contracts, TreasureManagerAddr)
	return Contracts
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		Migrations: ctx.String(flags.MigrationsFlag.Name),
		Chain: ChainConfig{
			ChainId:        ctx.Uint(flags.ChainIdFlag.Name),
			ChainRpcUrl:    ctx.String(flags.ChainRpcFlag.Name),
			StartingHeight: ctx.Uint64(flags.StartingHeightFlag.Name),
			Confirmations:  ctx.Uint64(flags.ConfirmationsFlag.Name),
			BlockStep:      ctx.Uint64(flags.BlocksStepFlag.Name),
			Contracts:      LoadContracts(),
			LoopInterval:   ctx.Duration(flags.LoopIntervalFlag.Name),
		},
		MasterDB: DBConfig{
			Host:     ctx.String(flags.MasterDbHostFlag.Name),
			Port:     ctx.Int(flags.MasterDbPortFlag.Name),
			Name:     ctx.String(flags.MasterDbNameFlag.Name),
			User:     ctx.String(flags.MasterDbUserFlag.Name),
			Password: ctx.String(flags.MasterDbPasswordFlag.Name),
		},
		SlaveDB: DBConfig{
			Host:     ctx.String(flags.SlaveDbHostFlag.Name),
			Port:     ctx.Int(flags.SlaveDbPortFlag.Name),
			Name:     ctx.String(flags.SlaveDbNameFlag.Name),
			User:     ctx.String(flags.SlaveDbUserFlag.Name),
			Password: ctx.String(flags.SlaveDbPasswordFlag.Name),
		},
		SlaveDbEnable: ctx.Bool(flags.SlaveDbEnableFlag.Name),
		HTTPServer: ServerConfig{
			Host: ctx.String(flags.HttpHostFlag.Name),
			Port: ctx.Int(flags.HttpPortFlag.Name),
		},
	}
}
