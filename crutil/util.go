package crutil

import (
	"github.com/gabrielgouv/code-runner-sd-ap2/cmdrunner"
	"strings"
)

func StringToCommand(str string) (cmd cmdrunner.Cmd) {
	stringArr := strings.Split(str, " ")
	cmd = cmdrunner.Cmd{Name: stringArr[0], Args:stringArr[1:]}
	return
}

func StringsToCommands(strs ...string) (cmdArr []cmdrunner.Cmd) {
	for i := 0; i < len(strs); i++ {
		cmdArr = append(cmdArr, StringToCommand(strs[i]))
	}
	return
}