package config

import (
	"context"
	"github.com/hell-kitchen/pkg/confita"
)

type Ingredients struct {
	Addr     string `config:"ingredients-addr"`
	UseTLS   bool   `config:"ingredients-use-tls"`
	CertFile string `config:"ingredients-cert-file"`
	KeyFile  string `config:"ingredients-key-file"`
}

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
