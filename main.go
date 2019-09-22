package main

import (
	"os"
)

const (
	// Version is constant that is tool's version.
	Version string = "1.0.0"
)

func main() {
	cli := &CLI{
		outStream: os.Stdout,
		errStream: os.Stderr,
		client:    &SlackClient{},
	}
	os.Exit(cli.Run(os.Args))
}
