/*
 *  Copyright (c) 2021-2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"go.osspkg.com/static"
)

var template = `// Code generated by go.osspkg.com/static. DO NOT EDIT.
package %s

import "go.osspkg.com/static"

func init() {
	//%s static archive
	%s = func(s string) static.Reader {
		c := static.New()
		if err := c.FromBase64TarGZ(s); err != nil {
			panic(err.Error())
		}
		return c
	}("%s")
}

/* FILES:
%s
*/
`

var (
	validRex = regexp.MustCompile(`(?i)[^a-z0-9]`)
	packRex  = regexp.MustCompile(`package ([a-z0-9_]+)`)
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		panic("bad args count")
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	spath := strings.Split(path, "/")

	pack, err := extractFromFiles(path)
	if err != nil {
		pack = spath[len(spath)-1]
	}
	pack = validRex.ReplaceAllString(pack, `_`)

	cache := static.New()
	err = cache.FromDir(path + "/" + args[0])
	if err != nil {
		panic(err.Error())
	}

	b64, err := cache.ToBase64TarGZ()
	if err != nil {
		panic(err.Error())
	}

	data := fmt.Sprintf(template, pack, args[1], args[1], b64, strings.Join(cache.List(), "\n"))
	filename := strings.ToLower(args[1])
	err = os.WriteFile(filename+"_static.go", []byte(data), os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
}

func extractFromFiles(path string) (string, error) {
	list, err := filepath.Glob(path + "/*.go")
	if err != nil {
		return "", err
	}

	for _, s := range list {
		if strings.Contains(s, "_static.go") {
			continue
		}
		if strings.Contains(s, "_test.go") {
			continue
		}

		b, err0 := os.ReadFile(s)
		if err0 != nil {
			return "", err0
		}

		result := packRex.FindAllStringSubmatch(string(b), -1)
		if len(result) == 1 && len(result[0]) == 2 {
			return result[0][1], nil
		}
	}

	return "", fmt.Errorf("empty folder")
}
