package app

import (
	"log/slog"
	client "task2/internal/app/client"
	indexer "task2/internal/app/indexer"
)

type App struct {
	Indexer *indexer.App
}

func New(
	log *slog.Logger,
	url string,
	filename string,
) *App {
	client, err := client.New(url)
	if err != nil {
		panic("failed to create client " + err.Error())
	}

	indexer := indexer.New(
		log,
		filename,
		client.C,
	)

	return &App{
		Indexer: indexer,
	}
}
