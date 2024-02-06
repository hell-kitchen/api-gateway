package config

import (
	"context"
	"fmt"
)

type Server struct {
	BindAddr string `config:"bind-addr"`
	BindPort int    `config:"bind-port"`
}

// NewServer return new Server config.
func NewServer() (*Server, error) {
	ctx := context.Background()

	s := &Server{
		BindAddr: "localhost",
		BindPort: 8080,
	}
	if err := getConfitaLoader().Load(ctx, s); err != nil {
		return nil, fmt.Errorf("error while loading config: %w", err)
	}
	return s, nil
}
