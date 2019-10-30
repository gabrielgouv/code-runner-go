package cmdrunner

type Cmd struct {
	Name string
	Args []string
	MeasureExecTime bool
}