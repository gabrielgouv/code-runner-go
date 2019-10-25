package cmdrunner

import (
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

	cmdOut, err := cmd.StdoutPipe()

	if err != nil {
		return "Failed to connect stdout pipe. This is an internal Code Runner error.", err

	} else {

		err = cmd.Start()

		if err != nil {
			return "Failed to start the command. This is an internal Code Runner error.", err
		}

		cmdBytes, err := ioutil.ReadAll(cmdOut)

		if err != nil {
			return "Error reading output. This is an internal Code Runner error.", err
		}

		err = cmd.Wait()

		if err != nil {
			return "Error waiting for process termination. This is an internal Code Runner error.", err
		} else {
			return string(cmdBytes), nil
		}

	}

}


