package config

import (
	"context"
	"sync"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
)

type iLoader interface {
	Load(ctx context.Context, to interface{}) error
}

var (
	_loader      iLoader
	_confitaMu   sync.Mutex
	_confitaOnce sync.Once
)

func initConfita() {
	_confitaMu.Lock()
	_loader = confita.NewLoader(
		env.NewBackend(),
		flags.NewBackend(),
	)
	_confitaMu.Unlock()
}

func getConfitaLoader() (l iLoader) {
	// init confita _loader if not initialized
	_confitaOnce.Do(initConfita)

	_confitaMu.Lock()
	l = _loader
	_confitaMu.Unlock()

	return l
}

// setNewLoader sets loader l as global.
//
// Return function on calling it will set old global loader back.
func setNewLoader(l iLoader) func() {
	_confitaMu.Lock()
	old := _loader
	_loader = l
	_confitaMu.Unlock()
	return func() {
		setNewLoader(old)
	}

}
