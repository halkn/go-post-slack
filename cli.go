package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
)

const (
	msgVersion          = "Version: %s \n"
	msgURLNothing       = "url option(-u) is required"
	msgContentNothing   = "either message option(-m) or file option(-f) is required"
	msgDuplicateContent = "message option(-m) or file option(-f) should be only one"
)

// CLI : struct for cli tool
type CLI struct {
	outStream, errStream io.Writer
	client               Poster
}

// Run : Execute main logic.
func (cli *CLI) Run(args []string) int {
	var (
		showVersion bool
		url         string
		message     string
		file        string
	)

	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	f.SetOutput(cli.errStream)
	f.BoolVar(&showVersion, "v", false, "show version")
	f.StringVar(&url, "u", "", "POST  URL")
	f.StringVar(&message, "m", "", "Message Contents")
	f.StringVar(&file, "f", "", "Message file")
	if err := f.Parse(args[1:]); err != nil {
		fmt.Fprintf(cli.errStream, "failed flag parse")
	}

	if showVersion {
		fmt.Fprintf(cli.outStream, msgVersion, Version)
		return 0
	}

	if url == "" {
		fmt.Fprintln(cli.errStream, msgURLNothing)
		return 1
	}

	jsonstr, err := createJsonstr(message, file)
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return 1
	}

	if err := cli.client.PostRequest(url, jsonstr); err != nil {
		fmt.Fprintln(cli.errStream, err)
		return 1
	}

	return 0
}

// createJsonstr : return jsonstring { "text": "XXXXXX" }
func createJsonstr(message, file string) (string, error) {
	if message == "" && file == "" {
		return "", errors.New(msgContentNothing)
	}

	if message != "" && file != "" {
		return "", errors.New(msgDuplicateContent)
	}

	if message != "" {
		return `{"text":"` + message + `"}`, nil
	}

	if file != "" {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			return "", fmt.Errorf("can't read file %s", err)
		}

		return `{"text":"` + string(bytes) + `"}`, nil
	}

	return "", nil
}
