// Package to read configuration files/params
package config

import (
	"github.com/joho/godotenv"
)

type Environment struct {
	DBHostName string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBPort     string
}

var ConnectionString string

// ReadConfig reads environment variables from config.env file
func ReadConfig(cfg string) error {
	if err := godotenv.Load(cfg); err != nil {
		return err
	}

	return nil
}
