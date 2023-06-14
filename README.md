# go-static

[![Coverage Status](https://coveralls.io/repos/github/osspkg/go-static/badge.svg?branch=master)](https://coveralls.io/github/osspkg/go-static?branch=master)
[![Release](https://img.shields.io/github/release/osspkg/go-static.svg?style=flat-square)](https://github.com/osspkg/go-static/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/osspkg/go-static)](https://goreportcard.com/report/github.com/osspkg/go-static)
[![CI](https://github.com/osspkg/go-static/actions/workflows/ci.yml/badge.svg)](https://github.com/osspkg/go-static/actions/workflows/ci.yml)

_Library for embedding static files inside an application_

# Install as tool

```bash
go install github.com/osspkg/go-static/cmd/static@latest
```

# Packaging

```go
//go:generate static <DIR> <VAR>
```

* DIR - Path to the static folder
* VAR - A variable containing `static.Reader` interface

## Example go code

```go
package example

import (
	"fmt"

	"github.com/osspkg/go-static"
)

//go:generate static ./.. ui

var ui static.Reader

func run() {
	fmt.Println(ui.List())
}
```
## License

BSD-3-Clause License. See the LICENSE file for details.