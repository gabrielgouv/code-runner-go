package main

import (
	"fmt"
	"github.com/gabrielgouv/code-runner-sd-ap2/argsparser"
	"github.com/gabrielgouv/code-runner-sd-ap2/cmdrunner"
)

func main() {

	args := argsparser.ParseArgs()

	runner := cmdrunner.CmdRunner{}
	runner.Dir = args.Dir

	commandCompile := cmdrunner.Cmd{Name:"javac", Args:[]string{"Main.java"}}
	commandRun := cmdrunner.Cmd{Name:"java", Args:[]string{"Main"}}

	cmdOutput := runner.RunCommands(commandCompile, commandRun)

	if cmdOutput.Error != nil {
		fmt.Printf("%s", cmdOutput.Error)
	} else {
		fmt.Printf("Output: %v", cmdOutput.Output)
	}

}
