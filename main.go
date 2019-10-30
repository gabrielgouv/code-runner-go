package main

import (
	"fmt"
	"github.com/gabrielgouv/code-runner-sd-ap2/cmdrunner"
	"github.com/gabrielgouv/code-runner-sd-ap2/crutil"
	"github.com/google/logger"
	"io/ioutil"
)

func main() {
	args := crutil.ParseArgs()

	// Init logger
	defer logger.Init("CodeRunner", args.Log, false, ioutil.Discard).Close()

	runner := cmdrunner.CmdRunner{}
	runner.Dir = args.Dir

	commandCopy := cmdrunner.Cmd{Name: "cp", Args:[]string{"/Users/gabrielgouv/Documents/codes/Main.java", "/Users/gabrielgouv/Documents/codes/java/"}}
	commandCompile := cmdrunner.Cmd{Name: "javac", Args:[]string{"Main.java"}}
	commandRun := cmdrunner.Cmd{Name: "java", Args:[]string{"Main"}}
	commandRm := cmdrunner.Cmd{Name: "rm", Args:[]string{"Main.java", "Main.class"}}

	cmdOutput := runner.RunCommand(commandCopy, commandCompile, commandRun, commandRm)

	if cmdOutput.Error != nil {
		fmt.Printf("%s", cmdOutput.Error)
	} else {
		fmt.Printf("%v", cmdOutput.Output)
	}

}
