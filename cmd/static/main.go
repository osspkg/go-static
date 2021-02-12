package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/deweppro/go-static"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		panic("bad args count")
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	spath := strings.Split(path, "/")
	pack := spath[len(spath)-1]

	cache := static.New()
	cache.FromDir(args[0])

	b64, err := cache.ToBase64TarGZ()
	if err != nil {
		panic(err)
	}

	data := fmt.Sprintf("package %s\n\n//%s static archive\nvar %s = \"%s\"", pack, args[1], args[1], b64)
	err = ioutil.WriteFile(pack+"_static.go", []byte(data), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
