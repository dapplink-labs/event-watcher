package worker

import (
	"gorm.io/gorm"
	"math/big"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"
)

type DepositTokens struct {
	GUID         uuid.UUID      `gorm:"primaryKey" json:"guid"`
	TokenAddress common.Address `json:"token_address" gorm:"serializer:bytes"`
	Sender       common.Address `json:"sender" gorm:"serializer:bytes"`
	Amount       *big.Int       `gorm:"serializer:u256"`
	Timestamp    uint64
}

type DepositTokensView interface {
	QueryDepositTokensList(page int, pageSize int, order string) ([]DepositTokens, uint64)
}

type DepositTokensDB interface {
	DepositTokensView

	StoreDepositTokens([]DepositTokens) error
}

type depositTokensDB struct {
	gorm *gorm.DB
}

func (db depositTokensDB) QueryDepositTokensList(page int, pageSize int, order string) ([]DepositTokens, uint64) {
	panic("implement me")
}

func (db depositTokensDB) StoreDepositTokens(depositTokensList []DepositTokens) error {
	result := db.gorm.CreateInBatches(&depositTokensList, len(depositTokensList))
	return result.Error
}

func NewDepositTokensDB(db *gorm.DB) DepositTokensDB {
	return &depositTokensDB{gorm: db}
}
