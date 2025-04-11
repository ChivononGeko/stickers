package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	PortServer  string
	PosterToken string
	SumatraPDF  string
	PrinterName string
	AllowedIDs  map[string]struct{}
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("ошибка загрузки .env файла: %v", err)
	}

	config := &Config{
		PortServer:  getEnv("PORT_SERVER", ""),
		PosterToken: getEnv("POSTER_TOKEN", ""),
		SumatraPDF:  getEnv("SUMATRA_PATCH", ""),
		PrinterName: getEnv("PRINTER_NAME", ""),
		AllowedIDs:  parseAllowedIDs(getEnv("ALLOWED_IDS", "")),
	}

	if config.PortServer == "" {
		return nil, fmt.Errorf("PORT_SERVER не задан")
	}
	if config.SumatraPDF == "" {
		return nil, fmt.Errorf("SUMATRA_PATCH не задан")
	}
	if config.PosterToken == "" {
		return nil, fmt.Errorf("POSTER_TOKEN не задан")
	}
	if config.PrinterName == "" {
		return nil, fmt.Errorf("PRINTER_NAME не задан")
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func parseAllowedIDs(ids string) map[string]struct{} {
	allowedIDs := make(map[string]struct{})
	if ids == "" {
		return allowedIDs
	}

	for _, id := range strings.Split(ids, ",") {
		allowedIDs[id] = struct{}{}
	}

	return allowedIDs
}
