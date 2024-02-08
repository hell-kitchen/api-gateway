package config

import (
	"github.com/heetch/confita"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_setNewLoader(t *testing.T) {
	loader := getConfitaLoader()
	otherLoader := confita.NewLoader()
	require.NotEqual(t, loader, otherLoader)
	assert.Equal(t, loader, _loader)
	restore := setNewLoader(otherLoader)
	now := getConfitaLoader()
	assert.Equal(t, now, otherLoader)
	restore()
	now = getConfitaLoader()
	assert.NotEqual(t, now, otherLoader)
	assert.Equal(t, now, loader)
}
