package generator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewString(t *testing.T) {
	str := NewString(10)
	require.Len(t, str, 10)
}

func TestNewHash(t *testing.T) {
	hash := NewHash("test")
	require.NotEmpty(t, hash)
}
