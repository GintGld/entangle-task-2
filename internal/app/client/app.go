package app

import (
	"fmt"
	client "task2/internal/client/ETH"
)

type App struct {
	C *client.Client
}

func New(url string) (*App, error) {
	const op = "App.Client.New"

	c, err := client.New(url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &App{
		C: c,
	}, nil
}
