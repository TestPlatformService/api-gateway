package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_SERVICE string
	API_ROUTER   string

	ACCES_KEY   string
	REFRESH_KEY string
	MINIO_URL   string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found?")
	}

	config := Config{}
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", ":50051"))
	config.API_ROUTER = cast.ToString(Coalesce("API_ROUTER", ":8080"))
	config.ACCES_KEY = cast.ToString(Coalesce("ACCES_KEY", "flashsalee"))
	config.REFRESH_KEY = cast.ToString(Coalesce("REFRESH_KEY", "OzNur"))
	config.MINIO_URL = cast.ToString(Coalesce("MINIO_URL", "localhost:9000"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
