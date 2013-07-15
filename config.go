package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	RepositoryDir string
	WebDir        string
	Port          string
}

func (config *Config) Load(b []byte) error {
	if err := json.Unmarshal(b, config); err != nil {
		return err
	}
	if config.Port[0:0] != ":" {
		config.Port = fmt.Sprintf(":%s", config.Port)
	}
	if config.RepositoryDir[len(config.RepositoryDir)-1:len(config.RepositoryDir)] == "/" {
		config.RepositoryDir = config.RepositoryDir[:len(config.RepositoryDir)-1]
	}
	return nil
}
