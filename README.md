# LegionTD 2 SDK for Go

This is an experimental SDK that I am writting in Go, contributions are welcome.

## Getting started

```go
package main

import (
    "fmt"
    
    "github.com/syrull/ltdsdk"
)

func main() {
    api := NewLTDSDK("test_api_key", "https://apiv2.legiontd2.com/")
    unit, _ := api.GetUnit("Pollywog", "9.06.4")
    fmt.Println(unit.Name)
    fmt.Println(unit.Damage)
}

```

## Runing the tests

```console
$ go test
```