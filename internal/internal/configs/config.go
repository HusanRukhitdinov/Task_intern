package configs

import (
	"fmt"
	"os"

	"github.com/spf13/cast"

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

	SuperUserID string

	// Email configuration
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	FromEmail    string
	FromName     string
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

	cfg.SuperUserID = cast.ToString(getOrReturnDefault("SUPER_USER_ID", "super_user_id"))

	// Email configuration
	cfg.SMTPHost = cast.ToString(getOrReturnDefault("SMTP_HOST", "smtp.gmail.com"))
	cfg.SMTPPort = cast.ToString(getOrReturnDefault("SMTP_PORT", "587"))
	cfg.SMTPUsername = cast.ToString(getOrReturnDefault("SMTP_USERNAME", ""))
	cfg.SMTPPassword = cast.ToString(getOrReturnDefault("SMTP_PASSWORD", ""))
	cfg.FromEmail = cast.ToString(getOrReturnDefault("FROM_EMAIL", ""))
	cfg.FromName = cast.ToString(getOrReturnDefault("FROM_NAME", "Intern App"))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}
