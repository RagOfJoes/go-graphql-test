package config

import (
	"os"
)

type MongoConfig struct {
	Uri string
}

type Config struct {
	Port  string
	Mongo MongoConfig
}

func New() *Config {
	return &Config{
		Port: getEnv("PORT", "8081"),
		Mongo: MongoConfig{
			Uri: getEnv("MONGO_URI", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
