package service

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"task2/internal/lib/utils/sl"
	"task2/internal/models"
)

type ETH struct {
	log *slog.Logger
	c   ETHClient
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name ETHClient
type ETHClient interface {
	Block(ctx context.Context, num *big.Int) (*models.ETHBlockReduced, error)
}

func New(
	log *slog.Logger,
	client ETHClient,
) *ETH {
	return &ETH{
		log: log,
		c:   client,
	}
}

// Block returnes reduced block by its number.
func (e *ETH) Block(ctx context.Context, num *big.Int) (*models.ETHBlockReduced, error) {
	const op = "ETH.Block"

	log := e.log.With(
		slog.String("op", op),
		slog.String("block numer", num.String()),
	)

	block, err := e.c.Block(ctx, num)
	if err != nil {
		log.Error("failed to get block", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return block, nil
}
