package example

import (
	"fmt"

	"github.com/deweppro/go-static"
)

//go:generate static ./.. UI

func run() {
	cache := static.New()
	if err := cache.FromBase64TarGZ(UI); err != nil {
		panic(err)
	}

	fmt.Println(cache.List())
}
