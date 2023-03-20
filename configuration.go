package main

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type ServiceConfig struct {
	Databasetype  DBTYPE `env:"DB_TYPE"`
	DBConnection  string `env:"DB_CONNECTION"`
	ServerAddress string `env:"SERVER_ADDRESSS" evnDefault:"3000"`
	IsProduction  string `env:"PRODUCTION"`
}

func ExtractConfiguration() (ServiceConfig, error) {
	// Loading the environment variables from '.env' file.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	cfg := ServiceConfig{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	return cfg, nil
}
