package main

import (
	"encoding/json"
	"fmt"
	"github.com/gabrielgouv/code-runner-sd-ap2/cmdrunner"
	"github.com/gabrielgouv/code-runner-sd-ap2/crutil"
	"github.com/google/logger"
	"io/ioutil"
	"log"
)

func main() {
	args := crutil.ParseArgs()

	// Init logger
	defer logger.Init("CodeRunner", args.Log, false, ioutil.Discard).Close()
	logger.SetFlags(log.LstdFlags)

	runner := cmdrunner.CmdRunner{}
	runner.Dir = args.Dir
	//runner.MeasureTotalExecTime = true

	commandCopy := cmdrunner.Cmd{Name: "cp", Args:[]string{"/Users/gabrielgouv/Documents/codes/Main.java", "/Users/gabrielgouv/Documents/codes/java/"}}
	commandCompile := cmdrunner.Cmd{Name: "javac", Args:[]string{"Main.java"}}
	commandRun := cmdrunner.Cmd{Name: "java", Args:[]string{"Main"}, MeasureExecTime:true}
	commandRm := cmdrunner.Cmd{Name: "rm", Args:[]string{"Main.java", "Main.class"}}

	cmdOutput := runner.RunCommand(commandCopy, commandCompile, commandRun, commandRm)

	jsonOutput, _ := json.MarshalIndent(&cmdOutput, "", "  ")

	fmt.Printf("%s", string(jsonOutput))

}
