package config

import (
	"sync"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
)

var (
	_loader      *confita.Loader
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

func getConfitaLoader() (l *confita.Loader) {
	// init confita _loader if not initialized
	_confitaOnce.Do(initConfita)

	_confitaMu.Lock()
	l = _loader
	_confitaMu.Unlock()

	return l
}
