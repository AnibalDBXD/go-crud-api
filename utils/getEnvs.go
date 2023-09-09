package utils

import (
	"fmt"
	"os"
)

type AppEnv struct {
	Port   string
	DB_URL string
}

func GetEnvs() (AppEnv, error) {

	portString := os.Getenv("PORT")
	dbUrlString := os.Getenv("DB_URL")

	if portString == "" {
		return AppEnv{}, fmt.Errorf("PORT si not found in environment variables")
	}

	appEnv := AppEnv{
		Port:   portString,
		DB_URL: dbUrlString,
	}

	return appEnv, nil
}
