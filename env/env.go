package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/prinick96/elog"
	"os"
)

const DEFAULT_PORT = "8000"

type EnvApp struct {
	// Server config
	SERVER_PORT string

	// Database config
	DB_ENGINE   string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_OPTIONS  string
}

// GetEnv loads the environment configuration from the provided file and assigns
// the loaded value to the EnvApp struct and returns its value
func GetEnv(fileName string) EnvApp {
	err := godotenv.Load(fileName)
	elog.New(elog.PANIC, fmt.Sprintf("Unable to load the file %s", fileName), err)
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	return EnvApp{
		SERVER_PORT: PORT,
		DB_ENGINE:   os.Getenv("DB_ENGINE"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_OPTIONS:  os.Getenv("DB_OPTIONS"),
	}
}
