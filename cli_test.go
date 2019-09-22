package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

type dummyClient struct{}

func (dmy *dummyClient) PostRequest(url, jsonstr string) (err error) {
	if url == "NGURL" {
		return errors.New("dummy for NG Test")
	}
	return nil
}

func TestRun(t *testing.T) {

	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)

	dummyCli := &CLI{
		outStream: outStream,
		errStream: errStream,
		client:    &dummyClient{},
	}
	type pattern struct {
		args      []string
		expectMsg string
	}

	// OK test.
	func() {
		patterns := []pattern{
			{strings.Split("go-post-slack -u testurl -m testmessage", " "), ""},
			{strings.Split("go-post-slack -u testurl -f testdata/testMessage.dat", " "), ""},
			{strings.Split("go-post-slack -v", " "), fmt.Sprintf(msgVersion, Version)},
		}
		for _, p := range patterns {
			status := dummyCli.Run(p.args)
			if status != 0 {
				t.Errorf("expected %d but it was actually %d", status, 0)
			}
			if outStream.String() != p.expectMsg {
				t.Errorf("expected %s but it was actually %s", outStream.String(), p.expectMsg)
			}
			outStream.Reset()
		}
	}()

	// NG test.
	func() {
		patterns := []pattern{
			{
				strings.Split("go-post-slack ", " "),
				fmt.Sprintln(msgURLNothing),
			},
			{
				strings.Split("go-post-slack -m test", " "),
				fmt.Sprintln(msgURLNothing),
			},
			{
				strings.Split("go-post-slack -f testdata/testMessage.dat", " "),
				fmt.Sprintln(msgURLNothing),
			},
			{
				strings.Split("go-post-slack -u testurl", " "),
				fmt.Sprintln(msgContentNothing),
			},
			{
				strings.Split("go-post-slack -u testurl -m testmessage -f testdata/testMessage.dat", " "),
				fmt.Sprintln(msgDuplicateContent),
			},
		}
		for _, p := range patterns {
			status := dummyCli.Run(p.args)
			if status != 1 {
				t.Errorf("expected %d but it was actually %d", status, 1)
			}
			if errStream.String() != p.expectMsg {
				t.Errorf("expected %s but it was actually %s", errStream.String(), p.expectMsg)
			}
			errStream.Reset()
		}
	}()

	// NG file not found pattern with -f opt.
	func() {
		args := strings.Split("go-post-slack -u testurl -f testdata/nothing.dat", " ")
		if status := dummyCli.Run(args); status != 1 {
			t.Errorf("expected %d but it was actually %d", status, 1)
		}
	}()

	// NG httpclient return error
	func() {
		args := strings.Split("go-post-slack -u NGURL -m NGMSG", " ")
		if status := dummyCli.Run(args); status != 1 {
			t.Errorf("expected %d but it was actually %d", status, 1)
		}
	}()
}
