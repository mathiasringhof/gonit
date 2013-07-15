package main

import (
	"fmt"
	"net/http"
)

const configFile = "./config.json"

var config Config

func main() {
	config = Config{}
	if err := config.Load(nil); err != nil {
		fmt.Printf("Error while loading configuration file %s: %s\n", configFile, err.Error())
		return
	}
	printConfig(&config)
	serveHttp(&config)
}

func printConfig(config *Config) {
	fmt.Printf("Configuration loaded: %s\n", configFile)
	fmt.Printf("Using repository directory: %s\n", config.RepositoryDir)
	fmt.Printf("Serving web content from folder %s using port %s\n", config.WebDir, config.Port)
}

func serveHttp(config *Config) {
	setupHandlers()
	if err := http.ListenAndServe(config.Port, nil); err != nil {
		fmt.Printf("Error starting web server: %s\n", err.Error())
	}
}
