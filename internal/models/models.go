package models

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

// Block contains only neccesary
// information about ETH block.
type ETHBlockReduced struct {
	Number    *big.Int
	Hash      string
	TxCount   int
	Timestamp time.Time
}

// ReduceBlock converts ethereum block
// to reduced block.
func ReduceBlock(b *types.Block) *ETHBlockReduced {
	return &ETHBlockReduced{
		Number:    b.Number(),
		Hash:      b.Hash().Hex(),
		TxCount:   len(b.Transactions()),
		Timestamp: time.Unix(int64(b.Time()), 0),
	}
}

// String formats reduces block to string.
func (e *ETHBlockReduced) String() string {
	return fmt.Sprintf(
		"Number: %s Hash: %s TxCount: %d Timestamp: %s\n",
		e.Number.String(),
		e.Hash,
		e.TxCount,
		e.Timestamp.UTC().Format(time.UnixDate),
	)
}
