package main

import (
	"net/url"
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
