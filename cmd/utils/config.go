package utils

import "github.com/joho/godotenv"

type Config struct {
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig() {
	godotenv.Load(".env")
}
