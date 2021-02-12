package static_test

import (
	"testing"

	"github.com/deweppro/go-static"
	"github.com/stretchr/testify/require"
)

func TestUnit_ReadDir(t *testing.T) {
	c := static.New()

	require.NoError(t, c.FromDir("./cmd"))
	require.Equal(t, c.List(), []string{"/cmd/static/main.go"})
}
