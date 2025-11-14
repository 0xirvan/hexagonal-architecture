package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Container for all configurations
type AppContainer struct {
	App  *App
	HTTP *HTTP
}

// New loads environment variables and returns the application configuration
func New() (*AppContainer, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	app := &App{
		Name: getEnv("APP_NAME", "TTA Svelte Go"),
		Env:  getEnv("APP_ENV", "development"),
	}

	http := &HTTP{
		Host:           getEnv("HTTP_HOST", "localhost"),
		Port:           getEnv("HTTP_PORT", "8080"),
		AllowedOrigins: getEnv("HTTP_ALLOWED_ORIGINS", "*"),
	}

	return &AppContainer{
		App:  app,
		HTTP: http,
	}, nil
}

// getEnv retrieves the vaulue of the envirment variable named by the key
// if the variable is empty, it returns the fallback value
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val == "" {
		return val
	}
	return fallback
}
