package util

import (
	"echo-mysql-default/internal/domain"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

func GetEnvironmentConfig() *domain.Config {
	config := new(domain.Config)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Missing the .env file:\n%v", GetStackTrace(err))
	}

	err = env.Parse(config)
	if err != nil {
		log.Fatalf("Failed to read .env variables:\n%s", GetStackTrace(err))
	}
	return config
}
