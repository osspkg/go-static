package demo

import (
	"fmt"

	"github.com/osspkg/go-static"
)

//go:generate static . ui

var ui static.Reader

func run() {
	fmt.Println(ui.List())
}
