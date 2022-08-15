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

## Routes Covered

#### Player

[x] `players/byName/%s`
[x] `players/byId/%s`
[ ] `players/bestFriends/%s`
[ ] `players/matchHistory/%s`
[ ] `players/stats/%s`
[ ] `players/stats`

#### Units

[x] - `units/byName/%s`
[ ] - `units/byId/%s`
[ ] - `units/byVersion/%s`

#### Games

[x] - `games/`
[x] - `games/byId/%s`

#### Info

[x] - `info/abilities/byId/%s `
[x] - `info/descriptions/%s`
[x] - `info/spells/byId/%s`
[ ] - `info/legions/byId/%s`
[ ] - `info/legions/{offset}/{limit}`
[ ] - `info/waves/byId/%s`
[ ] - `info/waves/{offset}/{limit}`
[ ] - `info/spells/byId/%s`
[ ] - `info/spells/{offset}/{limit}`
[ ] - `info/abilities/byId/%s`
[ ] - `info/abilities/{offset}/{limi`t}
[ ] - `info/research/byId/%s`
[ ] - `info/research/{offset}/{limit`}
[ ] - `info/descriptions/%s`
```