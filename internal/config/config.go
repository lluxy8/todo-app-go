package config

import (
	"fmt"
	"os"
)

type Config struct {
	App   AppConfig
	Mongo MongoConfig
}

type AppConfig struct {
	Env  string
	Port string
}

type MongoConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load() (*Config, error) {
	cfg := &Config{
		App: AppConfig{
			Env:  getEnv("APP_ENV", "local"),
			Port: getEnv("APP_PORT", "8080"),
		},
		Mongo: MongoConfig{
			Host:     getEnv("MONGO_HOST", "localhost"),
			Port:     getEnv("MONGO_PORT", "27017"),
			User:     getEnv("MONGO_USER", ""),
			Password: getEnv("MONGO_PASSWORD", ""),
			Database: getEnv("MONGO_DB", ""),
		},
	}

	return cfg, nil
}

func (m MongoConfig) URI() string {
	if m.User != "" {
		return fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s",
			m.User,
			m.Password,
			m.Host,
			m.Port,
			m.Database,
		)
	}

	return fmt.Sprintf(
		"mongodb://%s:%s/%s",
		m.Host,
		m.Port,
		m.Database,
	)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
