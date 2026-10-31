package config

import (
	"os"
)

type Config struct {
	Port      string
	LogLevel  string
	LogFormat string
}

var DefaultConfig = Config{
	Port:      "8080",
	LogLevel:  "INFO",
	LogFormat: "JSON",
}

// pointer for single unified memory reference
func LoadConfig() *Config {
	return &Config{
		Port:      getEnv("PORT", DefaultConfig.Port),
		LogLevel:  getEnv("LOG_LEVEL", DefaultConfig.LogLevel),
		LogFormat: getEnv("LOG_FORMAT", DefaultConfig.LogFormat),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
