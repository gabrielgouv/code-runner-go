package cmdrunner

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"strconv"
)

type CmdRunner struct {
	Dir string
}

type Cmd struct {
	Name string
	Args []string
}

type CmdOutput struct {
	Error error
	Output []string
}

func (c *CmdRunner) RunCommands(commands ...Cmd) (o CmdOutput) {

	for i := 0; i < len(commands); i++ {
		o = c.RunCommand(commands[i])
		if o.Error != nil {
			o.Error = errors.New("error running '" + commands[i].Name + "' command [" + strconv.Itoa(i) + "]:\n" + o.Error.Error())
			return
		}
	}

	return
}

func (c *CmdRunner) RunCommand(command Cmd) (o CmdOutput) {

	cmd := exec.Command(command.Name, command.Args...)

	if c.Dir != "" {
		cmd.Dir = c.Dir
	}

	cmdOut, errCmdOut := cmd.StdoutPipe()
	cmdErr, errCmdErr := cmd.StderrPipe()

	if errCmdOut != nil {
		return c.buildCmdOutput(errCmdErr, "Failed to connect stdout pipe. This is an internal Code Runner error.")
	}

	if errCmdErr != nil {
		return c.buildCmdOutput(errCmdErr, "Failed to connect stderr pipe. This is an internal Code Runner error.")
	}

	err := cmd.Start()

	if err != nil {
		return c.buildCmdOutput(err, "Failed to start the command. This is an internal Code Runner error.")
	}

	cmdOutBytes, errCmdOutBytes := ioutil.ReadAll(cmdOut)
	cmdErrBytes, errCmdErrBytes := ioutil.ReadAll(cmdErr)

	if errCmdOutBytes != nil {
		return c.buildCmdOutput(errCmdOutBytes, "Error reading stdout. This is an internal Code Runner error.")
	}

	if errCmdErrBytes != nil {
		return c.buildCmdOutput(errCmdErrBytes, "Error reading stderr. This is an internal Code Runner error.")
	}

	err = cmd.Wait()
	if err != nil {
		if cmdErrBytes != nil {
			return c.buildCmdOutput(errors.New(string(cmdErrBytes)), "A program execution error has occurred.")
		}
		return c.buildCmdOutput(errors.New("error waiting for process termination"), "Error waiting for process termination.")
	} else {

		if string(cmdErrBytes) != "" {
			return c.buildCmdOutput(errors.New(string(cmdErrBytes)), "A program execution error has occurred.")
		}

		return c.buildCmdOutput(nil, string(cmdOutBytes))
	}
}

func (c *CmdRunner) buildCmdOutput(error error, output ...string) (o CmdOutput) {

	if error != nil {
		o.Error = error
	}
	o.Output = output

	return
}
