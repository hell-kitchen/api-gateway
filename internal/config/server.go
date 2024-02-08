package config

import (
	"context"
	"fmt"
)

// Server is config to set on which port server will be running.
type Server struct {
	BindAddr string `config:"bind-addr,short=a"`
	BindPort int    `config:"bind-port,short=p"`
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

func (cfg Server) Addr() string {
	return fmt.Sprintf("%s:%d", cfg.BindAddr, cfg.BindPort)
}
