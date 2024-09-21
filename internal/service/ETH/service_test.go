package service

import (
	"context"
	"errors"
	"log/slog"
	"math/big"
	"os"
	"task2/internal/models"
	"task2/internal/service/ETH/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBlock(t *testing.T) {
	type args struct {
		ctx context.Context
		num *big.Int
	}
	type clientRes struct {
		resp *models.ETHBlockReduced
		err  error
	}
	type want struct {
		block *models.ETHBlockReduced
		err   error
	}
	tests := []struct {
		name      string
		args      args
		clientRes clientRes
		want      want
	}{
		{
			name: "main line",
			args: args{
				ctx: context.Background(),
				num: big.NewInt(123),
			},
			clientRes: clientRes{
				&models.ETHBlockReduced{
					Number:    big.NewInt(123),
					Hash:      "hash",
					TxCount:   10,
					Timestamp: time.Unix(100, 0),
				}, nil,
			},
			want: want{&models.ETHBlockReduced{
				Number:    big.NewInt(123),
				Hash:      "hash",
				TxCount:   10,
				Timestamp: time.Unix(100, 0),
			}, nil},
		},
		{
			name:      "client error",
			args:      args{context.Background(), big.NewInt(103)},
			clientRes: clientRes{nil, errors.New("some client error")},
			want:      want{nil, errors.New("ETH.Block: some client error")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.NewETHClient(t)

			client.
				On("Block", tt.args.ctx, tt.args.num).
				Return(tt.clientRes.resp, tt.clientRes.err)

			eth := ETH{
				log: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
				c:   client,
			}

			res, err := eth.Block(tt.args.ctx, tt.args.num)
			assert.Equal(t, tt.want.block, res)
			if tt.want.err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.want.err.Error())
			}
		})
	}
}
