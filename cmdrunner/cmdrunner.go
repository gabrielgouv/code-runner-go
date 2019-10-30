package cmdrunner

import (
	"errors"
	"github.com/google/logger"
	"io/ioutil"
	"os/exec"
	"strconv"
	"time"
)

type CmdRunner struct {
	Dir string
	MeasureTotalExecTime bool
}

func (c *CmdRunner) RunCommand(commands ...Cmd) (o CmdOutput) {

	var outputs []string
	var totalElapsed int64

	for i := 0; i < len(commands); i++ {
		logger.Infof("Step %d/%d : Running '%s %v'...", i + 1, len(commands), commands[i].Name, commands[i].Args)
		output, elapsed, err := c.runCommand(commands[i])

		if output != "" {
			logger.Infof(" └──> Output: %s", output)
			outputs = append(outputs, output)
		}

		if err != nil {
			o.Error = errors.New("Error running '" + commands[i].Name + "' command [" + strconv.Itoa(i + 1) + "]:\n" + o.Error.Error())
			logger.Fatalf(" └──> Error: %s", o.Error.Error())
			break
		}

		logger.Infof(" └──> Finished in %s", elapsed)

		if c.MeasureTotalExecTime || (!c.MeasureTotalExecTime && commands[i].MeasureExecTime) {
			totalElapsed += elapsed.Milliseconds()
		}

	}

	o.Output = outputs
	o.ExecTime = totalElapsed

	return
}

func (c *CmdRunner) runCommand(command Cmd) (string, time.Duration, error) {

	startTime := time.Now()

	cmd := exec.Command(command.Name, command.Args...)

	if c.Dir != "" {
		cmd.Dir = c.Dir
	}

	cmdOut, errCmdOut := cmd.StdoutPipe()
	cmdErr, errCmdErr := cmd.StderrPipe()

	if errCmdOut != nil {
		return "Failed to connect stdout pipe. This is an internal Code Runner error.", elapsedTime(startTime), errCmdErr
	}

	if errCmdErr != nil {
		return "Failed to connect stderr pipe. This is an internal Code Runner error.", elapsedTime(startTime), errCmdErr
	}

	err := cmd.Start()

	if err != nil {
		return "Failed to start the command. This is an internal Code Runner error.", elapsedTime(startTime), err
	}

	cmdOutBytes, errCmdOutBytes := ioutil.ReadAll(cmdOut)
	cmdErrBytes, errCmdErrBytes := ioutil.ReadAll(cmdErr)

	if errCmdOutBytes != nil {
		return "Error reading stdout. This is an internal Code Runner error.", elapsedTime(startTime), errCmdOutBytes
	}

	if errCmdErrBytes != nil {
		return "Error reading stderr. This is an internal Code Runner error.", elapsedTime(startTime), errCmdErrBytes
	}

	err = cmd.Wait()
	if err != nil {
		if cmdErrBytes != nil {
			return "A program execution error has occurred.", elapsedTime(startTime), errors.New(string(cmdErrBytes))
		}
		return "Error waiting for process termination.", elapsedTime(startTime), errors.New("error waiting for process termination")
	} else {

		if string(cmdErrBytes) != "" {
			return "A program execution error has occurred.", elapsedTime(startTime), errors.New(string(cmdErrBytes))
		}

		return string(cmdOutBytes), elapsedTime(startTime), nil
	}
}

func elapsedTime(start time.Time) time.Duration {
	return time.Since(start)
}
