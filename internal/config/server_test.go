package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewServer(t *testing.T) {
	s, err := NewServer()
	require.NoError(t, err)
	assert.NotNil(t, s)
}
