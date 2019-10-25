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

	output, err := runner.RunCommand("java", "Main")

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("Output: %s", output)
	}

}
