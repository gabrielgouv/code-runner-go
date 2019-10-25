package argsparser

import "flag"

type Args struct {
	Dir string
	Test string
	Main string
	Output string
	Location string
	EnableAsyncTests bool
}

// ParseArgs returns a struct with all arguments passed to program
func ParseArgs() (a Args) {
	dir := flag.String("dir", "", "Main directory.")
	test := flag.String("test", "", "Tests directory.")
	main := flag.String("main", "", "Main file.")
	output := flag.String("output", "", "Output type (json or text). Default is text.")
	location := flag.String("location", "", "Output location path.")
	enableAsyncTests := flag.Bool("enable-async-tests", false, "Enables async tests execution. Default is false.")

	flag.Parse()

	a.Dir = *dir
	a.Test = *test
	a.Main = *main
	a.Output = *output
	a.Location = *location
	a.EnableAsyncTests = *enableAsyncTests

	return
}