package config

import (
	"context"

	"github.com/hell-kitchen/pkg/confita"
)

// Tags is config of Tags Client.
type Tags struct {
	// Addr is full address to connect to Tags server.
	//
	// e.g. localhost:5000 or example.org
	Addr string `config:"tags-addr"`
	// UseTLS is used to use TLS transport credentials in work with service.
	UseTLS bool `config:"tags-use-tls"`
	// CertFile is location of CertFile.
	CertFile string `config:"tags-cert-file"`
	// KeyFile is location of KeyFile.
	KeyFile string `config:"tags-key-file"`
}

// NewTags creates new Tags config and Loads data from env/flag to it.
func NewTags() (*Tags, error) {
	cfg := &Tags{
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
