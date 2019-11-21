package cmdrunner

import "time"

type Cmd struct {
	Name string
	Args []string
	MeasureExecutionTime bool
	Timeout time.Duration
}