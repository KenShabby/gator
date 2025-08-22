// Package config handles JSON file structures
package config

import (
	"errors"
	"os"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg Config) Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, errors.New("could not open user's home directory")
	}
	gatorConfigLoc := homeDir + ".gatorconfig.json"

	gatorConfig := Config{DBUrl: gatorConfigLoc}

	return gatorConfig, nil
}

func (cfg Config) SetUser() error {

	return nil
}
