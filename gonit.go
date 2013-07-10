package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var configFile string = "./config.json"
var config *Config

type Config struct {
	RepositoryDir string
	WebDir        string
	Port          string
}

func main() {
	var err error
	config, err = loadConfig()
	if err != nil {
		fmt.Printf("Error while loading configuration file %s: %s\n", configFile, err.Error())
		return
	}
	printConfig()
	setupHandlers()
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("Error starting web server: %s\n", err.Error())
	}
}

func printConfig() {
	fmt.Printf("Loaded configuration: %s\n", configFile)
	fmt.Printf("Using repository directory: %s\n", config.RepositoryDir)
	fmt.Printf("Serving web content from folder %s using port %s\n", config.WebDir, config.Port)
}

func loadConfig() (config *Config, err error) {
	var b []byte
	b, err = ioutil.ReadFile(configFile)
	if err != nil {
		return
	}
	var conf Config
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return
	}
	config = &conf
	return
}
