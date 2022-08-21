# LegionTD 2 SDK for Go
![Coverage](https://img.shields.io/badge/Coverage-98.0%25-brightgreen)
![ghactions](https://github.com/syrull/ltdsdk/actions/workflows/go.yml/badge.svg)

[Dev/API Discord](https://discord.gg/8h9tkPf6Sw)

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
	unit, _ := api.GetUnit("Pollywog")
	fmt.Println(unit.Name)
	fmt.Println(unit.DamagePerSecond)
}
```

## Runing the tests

```console
$ go test
```

## Examples

- [Getting all of the spells into a JSON files](examples/get_all_spells/main.go)
- [Getting all of the units into a JSON files](examples/get_all_units/main.go)