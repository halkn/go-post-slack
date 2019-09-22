package main

import (
	"bytes"
	"net/http"
)

// Poster is the interface to post to slack-api.
type Poster interface {
	PostRequest(url, jsonstr string) error
}

// SlackClient is a HTTP Client to post to slack.
type SlackClient struct{}

// PostRequest : Use your Incoming Webhook URL to post a message
func (c *SlackClient) PostRequest(url, jsonstr string) (err error) {
	// generate http-reqest
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonstr)),
	)
	if err != nil {
		return
	}

	// Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Do post
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() { err = res.Body.Close() }()

	return
}
