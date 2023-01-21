package server

import (
	"context"
	"errors"

	"github.com/gorilla/mux"
)

type Server interface {
	Config() *Config
}

// NewSerer builder of subscriber Server Broker
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.port == "" {
		return nil, errors.New("port is required")
	}

	if config.dataBaseUrl == "" {
		return nil, errors.New("databaseurl is required")
	}
	return &Broker{
		config: config,
		router: mux.NewRouter(),
	}, nil
}
