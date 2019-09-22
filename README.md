# go-post-slack

CLI tool for sending messages to slack written in Golang.  
Use the [Incoming Webhooks](https://api.slack.com/incoming-webhooks) Slack API.

## Usage

```sh
Usage of ./go-post-slack:
  -f string
        Message file
  -m string
        Message Contents
  -u string
        POST  URL
  -v    show version
```

### Example

* -m option

```sh
go-post-slack -u $WebhookURL -m message
```

* -f option

```sh
go-post-slack -u $WebhookURL -f path/to/messagefile
```

## Installation

```sh
go get github.com/halkn/go-post-slack
cd $GOPATH/github.com/halkn/go-post-slack
go install
```

## License

MIT

