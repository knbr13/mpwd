package main

import (
	"encoding/json"
	"net/url"
	"os"
	"sync"
)

type Config struct {
	HealthCheckInterval string   `json:"healthCheckInterval"`
	Servers             []string `json:"servers"`
	ListenPort          string   `json:"listenPort"`
}

type Server struct {
	URL               *url.URL   // URL of the backend server.
	ActiveConnections int        // Count of active connections
	Mutex             sync.Mutex // A mutex for safe concurrency
	Healthy           bool
}

func loadConfig(file string) (Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	var config Config
	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
