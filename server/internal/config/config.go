package config

import (
	"os"

	"github.com/joho/godotenv"
)


type Config struct{
	Port string
	DBUrl string
}

func Load() Config{
	_ = godotenv.Load()

	cnfg := Config{
		Port: getEnv("PORT", "8080"),
		DBUrl: getEnv("DB_URL", "postgres://hisham@localhost:5432/momentum?sslmode=disable"),
	}

	return cnfg
}


func getEnv(key, fallback string) string{
	value, exists := os.LookupEnv(key);
	if !exists {
		return fallback
	}
	return value
}
