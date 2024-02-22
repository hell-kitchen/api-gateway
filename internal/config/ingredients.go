package config

import (
	"context"

	"github.com/hell-kitchen/pkg/confita"
)

// Ingredients is configuration of Ingredients Client.
type Ingredients struct {
	// Addr is full address to connect to Ingredients server.
	//
	// e.g. localhost:5000 or example.org
	Addr string `config:"ingredients-addr"`
	// UseTLS is used to use TLS transport credentials in work with service.
	UseTLS bool `config:"ingredients-use-tls"`
	// CertFile is location of CertFile.
	CertFile string `config:"ingredients-cert-file"`
	// KeyFile is location of KeyFile.
	KeyFile string `config:"ingredients-key-file"`
}

// NewIngredients creates new Ingredients config and Loads data from env/flag to it.
func NewIngredients() (*Ingredients, error) {
	cfg := &Ingredients{
		Addr:     ":5000",
		UseTLS:   false,
		CertFile: "",
		KeyFile:  "",
	}

	if err := confita.Get().Load(context.Background(), cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
