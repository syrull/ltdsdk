# LegionTD 2 SDK for Go
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
![ghactions](https://github.com/syrull/ltdsdk/actions/workflows/go.yml/badge.svg)
[Dev/API Chat](https://discord.gg/8h9tkPf6Sw)

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
	fmt.Println(unit.DmgMax)
}
```

## Runing the tests

```console
$ go test
```

## Routes Covered

```
Player
	[+] players/byName/%s
	[+] players/byId/%s
	[-] players/bestFriends/%s
	[-] players/matchHistory/%s
	[-] players/stats/%s
	[-] players/stats

Units

	[+] units/byName/%s
	[-] units/byId/%s
	[-] units/byVersion/%s

Games

	[+] games/
	[+] games/byId/%s

Info

	[+] info/abilities/byId/%s 
	[+] info/descriptions/%s
	[+] info/spells/byId/%s
	[+] info/waves/byId/%s
	[+] info/abilities/byId/%s
	[+] info/spells/byId/%s
	[-] info/abilities/{offset}/{limit}
	[-] info/legions/byId/%s
	[-] info/legions/{offset}/{limit}
	[-] info/waves/{offset}/{limit}
	[-] info/spells/{offset}/{limit}
	[-] info/research/byId/%s
	[-] info/research/{offset}/{limit}
	[-] info/descriptions/%s
```
