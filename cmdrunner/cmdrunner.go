package cmdrunner

import (
	"context"
	"errors"
	"github.com/google/logger"
	"os/exec"
	"strconv"
	"time"
)

type CmdRunner struct {
	Dir string
	MeasureGlobalExecutionTime bool
	GlobalTimeout time.Duration
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

		if elapsed.Milliseconds() == -1 {
			break
		}

		logger.Infof(" └──> Finished in %s", elapsed)

		if c.MeasureGlobalExecutionTime || (!c.MeasureGlobalExecutionTime && commands[i].MeasureExecutionTime) {
			totalElapsed += elapsed.Milliseconds()
		}

	}

	o.Output = outputs
	o.ExecutionTime = totalElapsed

	return
}

func (c *CmdRunner) runCommand(command Cmd) (string, time.Duration, error) {

	// Max limit is 1 minute if nothing is set
	timeout := 1 * time.Minute

	if c.GlobalTimeout.Milliseconds() != 0 {
		timeout = c.GlobalTimeout
	} else if c.GlobalTimeout.Milliseconds() == 0 && command.Timeout.Milliseconds() != 0 {
		timeout = command.Timeout
	}

	startTime := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, command.Name, command.Args...)

	if c.Dir != "" {
		cmd.Dir = c.Dir
	}

	out, err := cmd.Output()

	if ctx.Err() == context.DeadlineExceeded {
		return "Execution timed out", -1, nil
	}

	if err != nil {
		return string(out), elapsedTime(startTime), errors.New("Non-zero exit code: " + err.Error())
	} else {
		return string(out), elapsedTime(startTime), nil
	}

}

func elapsedTime(start time.Time) time.Duration {
	return time.Since(start)
}
