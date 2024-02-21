package config

import (
	"context"
	"github.com/hell-kitchen/pkg/confita"
)

type Tags struct {
	Addr     string `config:"tags-addr"`
	UseTLS   bool   `config:"use-tls"`
	CertFile string `config:"cert-file"`
	KeyFile  string `config:"key-file"`
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
