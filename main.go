package main

import (
	"github.com/NViktorovich/ktsnet/configures"
	"github.com/NViktorovich/ktsnet/fileoperation"
)

func main() {
	subNet, netInterface, filePath := configures.CheckUserCase()
	fileoperation.InterfacesReplace(filePath, netInterface, subNet)
}
