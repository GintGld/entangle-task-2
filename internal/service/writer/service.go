package service

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"task2/internal/lib/utils/sl"
	"task2/internal/models"
)

type Writer struct {
	log      *slog.Logger
	filename string
	readChan <-chan *models.ETHBlockReduced
}

func New(
	log *slog.Logger,
	filename string,
	readChan <-chan *models.ETHBlockReduced,
) *Writer {
	return &Writer{
		log:      log,
		filename: filename,
		readChan: readChan,
	}
}

// Run writes blocks from channel to file.
func (w *Writer) Run(ctx context.Context) error {
	const op = "Writer.Run"

	log := w.log.With(
		slog.String("op", op),
	)

	file, err := os.OpenFile(w.filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Error("failed to open file", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}
	defer file.Close()

read_loop:
	for block := range w.readChan {
		file.WriteString(block.String())

		select {
		case <-ctx.Done():
			break read_loop
		default:
		}
	}

	log.Info("finished writing blocks")

	return nil
}
