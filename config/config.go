package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Port string
}

var (
	DBConfig *DatabaseConfig
	SvConfig *ServerConfig
)

// init Load environment variables
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load .env file")
	}

	DBConfig = &DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		Name:     getEnv("DB_NAME", ""),
	}

	if DBConfig.Password == "" {

	}

	SvConfig = &ServerConfig{
		Port: getEnv("SERVER_PORT", "9090"),
	}

	missingVars := []string{}

	if DBConfig.User == "" {
		missingVars = append(missingVars, "DB_USER")
	}
	if DBConfig.Password == "" {
		missingVars = append(missingVars, "DB_PASSWORD")
	}
	if DBConfig.Name == "" {
		missingVars = append(missingVars, "DB_NAME")
	}

	if len(missingVars) > 0 {
		log.Fatalf("The following environment variables are required and have not been set: %v", missingVars)
	}
}

// getEnv uses LookupEnv to check if the variable exists
func getEnv(key, defaultValue string) string {
	if value, found := os.LookupEnv(key); found && value != "" {
		return value
	}

	return defaultValue
}
