package fileoperation

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/NViktorovich/ktsnet/configures"
)

func InterfacesReplace(filePath, netInterface, newSubnet string) {
	re := regexp.MustCompile(configures.IPPattern)
	var correctInterfaceBlock bool = false

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Print(err)
	}

	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		if strings.Contains(line, netInterface) {
			correctInterfaceBlock = true
		}
		if correctInterfaceBlock && strings.Contains(line, "address") {
			match := re.FindString(line)
			ipAddress := strings.Split(match, ".")
			newIP := replaceSubnet(newSubnet, ipAddress)
			jNewIP := strings.Join(newIP, ".")
			lines[i] = fmt.Sprintf("address %s # %s", jNewIP, "IP для работы в сети МПК")

			correctInterfaceBlock = false
			break
		}
	}
	newData := strings.Join(lines, "\n")
	err = ioutil.WriteFile("interfaces", []byte(newData), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func replaceSubnet(newSubnet string, fulNet []string) []string {
	fulNet[2] = newSubnet
	return fulNet
}
