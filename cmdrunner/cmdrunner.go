package cmdrunner

import (
	"errors"
	"io/ioutil"
	"os/exec"
)

type CmdRunner struct {
	Dir string
}

func (c *CmdRunner) RunCommand(name string, args... string) (string, error) {

	cmd := exec.Command(name, args...)

	if c.Dir != "" {
		cmd.Dir = c.Dir
	}

	cmdOut, errCmdOut := cmd.StdoutPipe()
	cmdErr, errCmdErr := cmd.StderrPipe()

	if errCmdOut != nil {
		return "Failed to connect stdout pipe. This is an internal Code Runner error.", errCmdErr
	}

	if errCmdErr != nil {
		return "Failed to connect stderr pipe. This is an internal Code Runner error.", errCmdErr
	}

	err := cmd.Start()

	if err != nil {
		return "Failed to start the command. This is an internal Code Runner error.", err
	}

	cmdOutBytes, errCmdOutBytes := ioutil.ReadAll(cmdOut)
	cmdErrBytes, errCmdErrBytes := ioutil.ReadAll(cmdErr)

	if errCmdOutBytes != nil {
		return "Error reading stdout. This is an internal Code Runner error.", errCmdOutBytes
	}

	if errCmdErrBytes != nil {
		return "Error reading stderr. This is an internal Code Runner error.", errCmdErrBytes
	}

	err = cmd.Wait()
	if err != nil {
		if cmdErrBytes != nil {
			return "A program execution error has occurred.", errors.New(string(cmdErrBytes))
		}
		return "Error waiting for process termination.", errors.New("error waiting for process termination")
	} else {

		if string(cmdErrBytes) != "" {
			return "A program execution error has occurred.",  errors.New(string(cmdErrBytes))
		}

		return string(cmdOutBytes), nil
	}


}


