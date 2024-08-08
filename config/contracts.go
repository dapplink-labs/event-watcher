package config

import "github.com/ethereum/go-ethereum/common"

const (
	TreasureManager = "0x4200000000000000000000000000000000000016"
)

var (
	TreasureManagerAddr = common.HexToAddress(TreasureManager)
)
