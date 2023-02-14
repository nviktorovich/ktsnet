package main

import (
	"github.com/NViktorovich/ktsnet/configures"
	"github.com/NViktorovich/ktsnet/fileoperation"
	"github.com/NViktorovich/ktsnet/systemoperation"
)

func main() {
	subNet, netInterface, filePath := configures.CheckUserCase()
	fileoperation.InterfacesReplace(filePath, netInterface, subNet)
	systemoperation.TerminalCommand(configures.NetworkingServiceCall, configures.RestartArg)

}
