package env

import (
	"os"

	"github.com/joho/godotenv"
)

var Env map[string]string

func GetEnv(key, def string) string {

	val := os.Getenv(key)
	if val != "" {
		return val
	}

	return def
}

func SetupEnvFile() {

	_ = godotenv.Load(".env")
}
