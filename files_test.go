package static_test

import (
	"testing"

	"github.com/osspkg/go-static"
	"github.com/stretchr/testify/require"
)

func TestUnit_ReadDir(t *testing.T) {
	c := static.New()

	require.NoError(t, c.FromDir("./cmd"))
	require.Equal(t, c.List(), []string{"/static/main.go", "/static/main_test.go"})
}
