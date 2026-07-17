package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/spf13/pflag"
)

type Config struct {
	Address string `env:"SERVER_ADDRESS"`
	BaseURL string `env:"BASE_URL"`
}

func Load() *Config {

	var cfg Config

	// TODO: ensure env is
	env.Parse(&cfg)
	cfg.LoadFlags()

	fmt.Println("Config is ", cfg)

	return &cfg

}

func (cfg *Config) LoadFlags() {
	pflag.StringVarP(&cfg.Address,
		"address",
		"a",
		"localhost:8080",
		"Address of the listener. Defaults to localhost:8080",
	)

	// TODO: must default to the requested host
	pflag.StringVarP(&cfg.BaseURL,
		"baseurl",
		"b",
		"localhost:8080",
		"address of the baseURL of the returned hashes,response example: <BASEURL>/<hash>",
	)
	pflag.Parse()
}
