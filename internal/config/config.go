package config

import (
	"math/big"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	RPC     string
	InitNum *big.Int
	Out     string
}

// LoadConfig parses command line argument.
func LoadConfig() *Config {
	fs := pflag.NewFlagSet("indexer", pflag.PanicOnError)
	fs.String("rpc", "", "blockchain RPC URL")
	fs.String("start", "", "number of block to start")
	fs.String("out", "", "output file")

	viper.BindPFlags(fs)

	fs.Parse(os.Args[1:])

	var c Config

	viper.AutomaticEnv()

	c.RPC = viper.GetString("rpc")
	if c.RPC == "" {
		panic("rpc not set")
	}

	c.Out = viper.GetString("out")
	if c.Out == "" {
		panic("out file not set")
	}

	var i big.Int
	if _, ok := i.SetString(viper.GetString("start"), 10); !ok {
		panic("can't parse start num: \"" + viper.GetString("start") + "\"")
	}
	c.InitNum = &i

	return &c
}
