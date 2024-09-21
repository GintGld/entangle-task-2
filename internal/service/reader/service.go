package service

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"task2/internal/lib/utils/sl"
	"task2/internal/models"
)

type Reader struct {
	log    *slog.Logger
	ethSrv ETH

	writeChan chan<- *models.ETHBlockReduced
}

type ETH interface {
	Block(context.Context, *big.Int) (*models.ETHBlockReduced, error)
}

func New(
	log *slog.Logger,
	ethSrv ETH,
	writechan chan<- *models.ETHBlockReduced,
) *Reader {
	return &Reader{
		log:       log,
		ethSrv:    ethSrv,
		writeChan: writechan,
	}
}

// Run setup reading blocks and sending them by channel.
func (r *Reader) Run(ctx context.Context, initNum *big.Int) error {
	const op = "Reader.Start"

	log := r.log.With(
		slog.String("op", op),
		slog.String("init num", initNum.String()),
	)

	one := big.NewInt(1)

write_loop:
	for num := initNum; ; num.Add(num, one) {
		block, err := r.ethSrv.Block(ctx, num)
		if err != nil {
			log.Error("failed to get block", sl.Err(err))
			return fmt.Errorf("%s: %w", op, err)
		}

		r.writeChan <- block

		select {
		case <-ctx.Done():
			break write_loop
		default:
		}
	}

	log.Info("finished reading blocks")

	close(r.writeChan)

	return nil
}
