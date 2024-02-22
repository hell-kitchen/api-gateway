package config

import (
	"context"
	"github.com/hell-kitchen/pkg/confita"
)

type Tags struct {
	Addr     string `config:"tags-addr"`
	UseTLS   bool   `config:"tags-use-tls"`
	CertFile string `config:"tags-cert-file"`
	KeyFile  string `config:"tags-key-file"`
}

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
