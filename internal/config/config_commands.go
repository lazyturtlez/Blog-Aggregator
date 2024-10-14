package config

import (
	"encoding/json"
	"errors"
	"os"
)

const configFile = "/.gatorconfig.json"

func Read() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	filePath := homeDir + configFile

	dat, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(dat, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Config) SetUser(username string) error {
	if len(username) < 1 {
		return errors.New("no username provided")
	}
	c.CurrentUserName = username

	return Write(c)
}

func Write(cfg *Config) error {
	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := homeDir + configFile
	return os.WriteFile(filePath, dat, 0644)
}
