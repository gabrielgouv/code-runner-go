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
	logger.SetFlags(log.LUTC)

	runner := cmdrunner.CmdRunner{}
	runner.Dir = args.Dir
	//runner.MeasureGlobalExecutionTime = true

	//commandCopy := cmdrunner.Cmd{Name: "cp", Args:[]string{"/Users/gabrielgouv/Documents/codes/Main.java", "/Users/gabrielgouv/Documents/codes/java/"}}
	//commandCompile := cmdrunner.Cmd{Name: "javac", Args:[]string{"Main.java"}}
	//commandRun := cmdrunner.Cmd{Name: "java", Args:[]string{"Main"}, MeasureExecTime:true}
	//commandRm := cmdrunner.Cmd{Name: "rm", Args:[]string{"Main.java", "Main.class"}}

	commandCopy := crutil.StringToCommand("cp /Users/gabrielgouv/Documents/codes/Main.java /Users/gabrielgouv/Documents/codes/java/")
	commandCompile := crutil.StringToCommand("javac Main.java")
	commandRun := crutil.StringToCommand("java Main")
	commandRun.MeasureExecutionTime = true
	//commandRun.Timeout = 100 * time.Millisecond
	commandRm := crutil.StringToCommand("rm Main.java Main.class")

	cmdOutput := runner.RunCommand(commandCopy, commandCompile, commandRun, commandRm)

	jsonOutput, _ := json.MarshalIndent(&cmdOutput, "", "  ")

	fmt.Printf("%s", string(jsonOutput))

}
