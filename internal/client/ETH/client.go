package client

import (
	"context"
	"fmt"
	"math/big"
	"task2/internal/models"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	url string
	c   *ethclient.Client
}

func New(url string) (*Client, error) {
	const op = "Client.New"

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Client{
		url: url,
		c:   client,
	}, nil
}

// Block returns ETH block by its number.
func (c *Client) Block(ctx context.Context, num *big.Int) (*models.ETHBlockReduced, error) {
	const op = "Client.Block"

	block, err := c.c.BlockByNumber(ctx, num)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return models.ReduceBlock(block), nil
}
