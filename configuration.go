package main

import (
	"fmt"

	"github.com/caarlos0/env"
)

type ServiceConfig struct {
	Databasetype  DBTYPE `env:"DB_TYPE"`
	DBConnection  string `env:"DB_CONNECTION"`
	ServerAddress string `env:"SERVER_ADDRESSS" evnDefault:"3000"`
	IsProduction  bool   `env:"PRODUCTION"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{}

	if err := env.Parse(&conf); err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	return conf, nil
}
