package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppContainer struct {
	App  *App
	HTTP *HTTP
}

func New() (*AppContainer, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	app := &App{
		Name: getEnv("APP_NAME", "TTA Svelte Go"),
		Env:  getEnv("APP_ENV", "development"),
	}

	http := &HTTP{
		Host: getEnv("HTTP_HOST", "localhost"),
		Port: getEnv("HTTP_PORT", "8080"),
	}

	return &AppContainer{
		App:  app,
		HTTP: http,
	}, nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val == "" {
		return val
	}
	return fallback
}
