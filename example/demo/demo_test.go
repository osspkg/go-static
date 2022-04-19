package demo

import (
	"fmt"
	"testing"

	"github.com/deweppro/go-static"
)

//go:generate static . ui

var ui static.Reader

func TestDemo(t *testing.T) {
	fmt.Println(ui.List())
}
