// Package to read configuration files/params
package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var ConnectionString string

// ReadConfig reads environment variables from config.env file
func ReadConfig(cfg string) error {
	if err := godotenv.Load(cfg); err != nil {
		return err
	}

	fmt.Println(os.Getenv("PGHOST"))
	return nil
}
