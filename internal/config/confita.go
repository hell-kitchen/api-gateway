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

// setNewLoader sets loader l as global.
//
// Return function on calling it will set old global loader back.
func setNewLoader(l *confita.Loader) func() {
	_confitaMu.Lock()
	old := _loader
	_loader = l
	_confitaMu.Unlock()
	return func() {
		setNewLoader(old)
	}

}
