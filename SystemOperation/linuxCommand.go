package systemoperation

import (
	"fmt"
	"log"
	"os/exec"
)

func TerminalCommand(cmd string, args ...string) {
	linuxCommand := exec.Command(cmd, args...)
	out, err := linuxCommand.Output()
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(out))
}
