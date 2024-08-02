package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env     string
	TZ      string
	Name    string
	Version string
	Api     ApiConfig
	Db      DatabaseConfig
	Log     LogConfig
}

type ApiConfig struct {
	Host        string
	Port        string
	Timeout     string
	HealthCheck string
	LogRequests bool
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

type LogConfig struct {
	Level    string
	Format   string
	Output   string
	FilePath string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env:     os.Getenv("ENV"),
		TZ:      os.Getenv("TZ"),
		Name:    os.Getenv("NAME"),
		Version: os.Getenv("VERSION"),
		Api: ApiConfig{
			Host:        os.Getenv("API_HOST"),
			Port:        os.Getenv("API_PORT"),
			Timeout:     os.Getenv("API_TIMEOUT"),
			HealthCheck: os.Getenv("API_HEALTH_CHECK"),
			LogRequests: os.Getenv("API_LOG_REQUESTS") == "true",
		},
		Db: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
		},
		Log: LogConfig{
			Level:    os.Getenv("LOG_LEVEL"),
			Format:   os.Getenv("LOG_FORMAT"),
			Output:   os.Getenv("LOG_OUTPUT"),
			FilePath: os.Getenv("LOG_FILE_PATH"),
		},
	}

	return cfg, nil
}
