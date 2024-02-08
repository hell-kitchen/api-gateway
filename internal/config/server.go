package config

import (
	"context"
	"fmt"
	"time"
)

// Server is config to set on which port server will be running.
type Server struct {
	BindAddr   string        `config:"bind-addr,short=a"`
	BindPort   int           `config:"bind-port,short=p"`
	BindHost   string        `config:"bind-host"`
	Timeout    time.Duration `config:"server-timeout"`
	AllowHosts []string      `config:"allow-hosts"`
	fullAddr   string
}

// NewServer return new Server config.
func NewServer() (*Server, error) {
	ctx := context.Background()

	s := &Server{
		BindAddr: "localhost",
		BindPort: 8080,
		BindHost: "localhost",
		AllowHosts: []string{
			"localhost",
			"127.0.0.1",
			"0.0.0.0",
		},
		Timeout: time.Second,
	}
	if err := getConfitaLoader().Load(ctx, s); err != nil {
		return nil, fmt.Errorf("error while loading config: %w", err)
	}
	s.fullAddr = fmt.Sprintf("http://%s", s.BindHost)
	return s, nil
}

// Addr return address on which application will be started at.
//
// Use Addr while configuring HTTP server.
func (cfg Server) Addr() string {
	return fmt.Sprintf("%s:%d", cfg.BindAddr, cfg.BindPort)
}

func (cfg Server) FullAddr() string {
	return cfg.fullAddr
}
