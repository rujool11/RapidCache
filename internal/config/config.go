package config

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Config struct {
	LogLevel    string
	LogFormat   string
	HTTPEnabled string
	RESPPort    string
	HTTPPort    string
}

var DefaultConfig = Config{
	LogLevel:    "INFO",
	LogFormat:   "JSON",
	HTTPEnabled: "true",
	RESPPort:    "6379",
	HTTPPort:    "8080",
}

var (
	validEnvBools   = []string{"true", "false"}
	validLogFormats = []string{"TEXT", "JSON"}
	validLogLevels  = []string{"DEBUG", "INFO", "WARN", "ERROR"}
)

// pointer for single unified memory reference
func LoadConfig() (*Config, error) {

	cfg := &Config{
		LogLevel:    strings.ToUpper(getEnv("LOG_LEVEL", DefaultConfig.LogLevel)),
		LogFormat:   strings.ToUpper(getEnv("LOG_FORMAT", DefaultConfig.LogFormat)),
		HTTPEnabled: strings.ToLower(getEnv("HTTP_ENABLED", DefaultConfig.HTTPEnabled)),
		RESPPort:    getEnv("RESP_PORT", DefaultConfig.RESPPort),
		HTTPPort:    getEnv("HTTP_PORT", DefaultConfig.HTTPPort),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	var errs []error

	if !slices.Contains(validLogLevels, c.LogLevel) {
		errs = append(errs, fmt.Errorf("LOG_LEVEL: %q is invalid (allowed: %s)",
			c.LogLevel, strings.Join(validLogLevels, ", ")))
	}
	if !slices.Contains(validLogFormats, c.LogFormat) {
		errs = append(errs, fmt.Errorf("LOG_FORMAT: %q is invalid (allowed: %s)",
			c.LogFormat, strings.Join(validLogFormats, ", ")))
	}

	errs = append(errs, validatePort("RESP_PORT", c.RESPPort))

	if !slices.Contains(validEnvBools, c.HTTPEnabled) {
		errs = append(errs, errors.New("HTTP_ENABLED must be a boolean"))
	}

	if c.HTTPEnabled == "true" {
		errs = append(errs, validatePort("HTTP_PORT", c.HTTPPort))
		if c.RESPPort == c.HTTPPort {
			errs = append(errs, errors.New("RESP_PORT and HTTP_PORT must differ"))
		}
	}

	return errors.Join(errs...)
}

func validatePort(name, p string) error {
	port, err := strconv.Atoi(p)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("%s: %q is not a valid port (1-65535)", name, p)
	}
	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
