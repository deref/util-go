package cmdutil

import (
	"fmt"
	"os"
)

func Warnf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "%v\n", fmt.Errorf(format, v...))
}

func Fatalf(format string, v ...interface{}) {
	Warnf(format, v...)
	os.Exit(1)
}

func Fatal(err error) {
	Fatalf("%v", err)
}

func MustGetwd() string {
	wd, err := os.Getwd()
	if err != nil {
		Fatalf("unable to get working directory: %w", err)
	}
	return wd
}
