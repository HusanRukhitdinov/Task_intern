package configs

import (
	"fmt"
	"github.com/spf13/cast"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string

	ServiceName string
	Environment string
	LoggerLevel string

	RedisHost string
	RedisPort string

	HTTPPort int
	HTTHost  string

}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env not found", err)
	}

	cfg := Config{}
	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "postgres"))
	cfg.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", 5432))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1111"))
	cfg.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DB", "task_intern"))


	cfg.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", "task_intern"))
	cfg.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "dev"))
	cfg.LoggerLevel = cast.ToString(getOrReturnDefault("LOGGER_LEVEL", "debug"))

	cfg.HTTHost = cast.ToString(getOrReturnDefault("HTTP_HOST", "localhost"))
	cfg.HTTPPort = cast.ToInt(getOrReturnDefault("HTTP_PORT", 8085))

	cfg.RedisHost = cast.ToString(getOrReturnDefault("REDIS_ADDRESS", "localhost"))
	cfg.RedisPort = cast.ToString(getOrReturnDefault("REDIS_PORT", 6379))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}
