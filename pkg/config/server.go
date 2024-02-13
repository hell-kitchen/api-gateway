package config

import (
	"context"
	"fmt"
	"time"

	"github.com/hell-kitchen/pkg/confita"
)

// Server is config of HTTP Server.
type Server struct {
	// BindAddr is Host address
	// Default: localhost
	BindAddr string `config:"bind-addr,short=a"`
	// BindPort is port on which application will be running at.
	// By default is 8080
	BindPort int `config:"bind-port,short=p"`
	// BindHost is external address of system.
	//
	// Do not add protocol in front of this parameter.
	// Bad: https://example.org http://example.org
	// Good: example.org
	//
	// If you set it to outer domain address and set UseHTTPS to true
	// then application will autogenerate cert file to use LetsEncrypt encryption.
	BindHost string `config:"bind-host"`
	// UseHTTPS is bool attribute to run HTTPS server instea of HTTP.
	// If setted to true then it will generate autocert file for you.
	UseHTTPS bool `config:"https"`
	// Timeout is timeout of server to wait for response.
	Timeout time.Duration `config:"server-timeout"`
	// AllowHosts is used in CORS configurations.
	AllowHosts []string `config:"allow-hosts"`
	// fullAddr is internal attribute to store address on which Server
	// will be running at.
	// Creating autimaticly on creation by BindAddr and BindHost.
	//
	// By default is "localhost:8080"
	fullAddr string
}

// NewServer return new Server config.
func NewServer() (*Server, error) {
	ctx := context.Background()

	s := &Server{
		BindAddr: "localhost",
		BindPort: 8080,
		BindHost: "localhost",
		UseHTTPS: false,
		Timeout:  time.Second,
		AllowHosts: []string{
			"localhost",
			"127.0.0.1",
			"0.0.0.0",
		},
		fullAddr: "",
	}
	if err := confita.Get().Load(ctx, s); err != nil {
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

// FullAddr return external address of server.
func (cfg Server) FullAddr() string {
	return cfg.fullAddr
}
