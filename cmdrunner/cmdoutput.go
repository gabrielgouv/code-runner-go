package cmdrunner

type CmdOutput struct {
	Error error
	Output []string
	ExecutionTime int64
}
