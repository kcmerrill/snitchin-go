package snitchin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	DEBUG = iota
	INFO
	CRITICAL
	WARNING
	FATAL
)

/* Create main struct */
type snitchin struct {
	channels map[string]channel
}

/* Struct to create channel */
type channel struct {
	name   string
	level  int
	writer io.Writer
	msg    string
}

/* Create the logging method */
func (c channel) Log(level int, msg string) {
	/* Is the channel level ok to be logged to? */
	if level >= c.level {
		msg = fmt.Sprintf("%s [%s] %s", time.Now().Format(time.RFC822), c.name, msg+"\n")
		io.WriteString(c.writer, msg)
	}
}

/* Create a new snitcher */
func New() *snitchin {
	/* Return a new snitchin */
	s := &snitchin{}
	/* Init the map */
	s.channels = make(map[string]channel)
	/* Return the goods */
	return s
}

/* Create a channel */
func (s *snitchin) AddChannel(name string, level int, writer io.Writer) {
	s.channels[name] = channel{name: name, level: level, writer: writer}
}

/* Return a specific channel */
func (s *snitchin) Channel(name string) channel {
	if _, exists := s.channels[name]; !exists {
		/* If the channel doesn't exist, create it with a basic stdout on the default level */
		s.AddChannel(name, DEBUG, os.Stdout)
	}

	/* Return a channel */
	return s.channels[name]
}

/* Send the log to every channel */
func (s *snitchin) Log(level int, msg string) {
	for _, channel := range s.channels {
		channel.Log(level, msg)
	}
}

/* Uses slack webhook integration. where webhook is the webhook url provided by slack */
func Slack(webhook string) *slack {
	return &slack{webhook: webhook}
}

type slack struct {
	webhook string
}

func (s slack) Write(msg []byte) (int, error) {
	/* Create a json struct to be sent to slack */
	slack_msg, err := json.Marshal(struct {
		Text string `json:"text"`
	}{
		string(msg),
	})

	if err == nil {
		http.Post(s.webhook, "application/json", bytes.NewBuffer(slack_msg))
	} else {
		fmt.Println("ERROR! There was an error posting to slack")
	}

	/* TODO: Come back to this in a bit */
	/* Not sure what to return here. I'm assuming it's how many bytes were written, then errors? */
	return 0, nil
}

func File(name string) *os.File {
	file, error := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if error != nil {
		/* Oh dear ..., we could panic, or just default to stdout. Hmmmm .... */
		panic("Unable to open the file: " + name)
	}

	/* Return the file writer */
	return file
}
