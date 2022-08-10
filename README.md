# LegionTD 2 SDK for Go

This is an experimental SDK that I am writting in Go, contributions are welcome.

## Getting started

```console
$ go get github.com/syrull/ltdsdk
```

Example of getting an example unit

```go
package main

import (
	"fmt"

	"github.com/syrull/ltdsdk"
)

func main() {
	api := ltdsdk.NewLTDSDK("API_KEY", "https://apiv2.legiontd2.com/")
	unit, _ := api.GetUnit("Pollywog", "9.06.4")
	fmt.Println(unit.Name)
	fmt.Println(unit.DmgMax)
}
```

## Runing the tests

```console
$ go test
```