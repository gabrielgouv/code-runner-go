package cmdrunner

type CmdOutput struct {
	Error error
	Output []string
	ExecTime int64
}
