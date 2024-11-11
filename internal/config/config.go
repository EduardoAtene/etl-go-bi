package config

import (
	"log"
	"os"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	DBName   string
	Port     string
}

type Config struct {
	Database DatabaseConfig
}

func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "admin"),
			Host:     getEnv("DB_HOST", "localhost"),
			DBName:   getEnv("DB_NAME", "DW_Acidentes"),
			Port:     getEnv("DB_PORT", "3306"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Variável de ambiente %s não encontrada, usando valor padrão: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
