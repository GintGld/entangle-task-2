package app

import (
	"context"
	"log/slog"
	"math/big"

	"task2/internal/models"
	eth "task2/internal/service/ETH"
	reader "task2/internal/service/reader"
	writer "task2/internal/service/writer"
)

type App struct {
	reader *reader.Reader
	writer *writer.Writer

	cancel context.CancelFunc
}

func New(
	log *slog.Logger,
	filename string,
	client reader.ETH,
) *App {
	eth := eth.New(
		log,
		client,
	)

	// Use channel without buffer
	// to prevent queue of blocks at shutdown.
	ch := make(chan *models.ETHBlockReduced)

	reader := reader.New(
		log,
		eth,
		ch,
	)
	writer := writer.New(
		log,
		filename,
		ch,
	)

	return &App{
		reader: reader,
		writer: writer,
	}
}

// Run starts indexer in parallel mode.
func (a *App) Run(initNum *big.Int) {
	ctx, cancel := context.WithCancel(context.Background())

	a.cancel = cancel

	go func() {
		if err := a.reader.Run(ctx, initNum); err != nil {
			a.cancel()
		}
	}()
	go func() {
		if err := a.writer.Run(ctx); err != nil {
			a.cancel()
		}
	}()
}

// Close closes internal context.
func (a *App) Close() {
	a.cancel()
}
