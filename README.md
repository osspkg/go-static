# go-static

_Library for embedding static files inside an application_

# Install as tool

```bash
go get -u github.com/deweppro/go-static/...
```

# Packaging

```go
//go:generate static <DIR> <VAR>
```

* DIR - Path to the static folder
* VAR - A variable containing an archive in base64 format

## Example

```go
//go:generate static ./../../ui UI
```

## Example go code

```go
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
```
