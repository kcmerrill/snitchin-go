package snitchin

import (
	"io"
	"os"
)

var channels map[string]*channel = make(map[string]*channel)

type channel struct {
	name   string
	level  int
	writer io.Writer
	format string
}

func Channel(name string) *channel {
	/* Does the channel exist? */
	if _, exists := channels[name]; !exists {
		CreateChannel(name, 700, os.Stdout, "basic")
	}
	return channels[name]
}

func (c channel) Log(level_name string, msg string) {
	_, level, _ := Level(level_name)
	if level >= c.level {
		msg = Formatter(level_name, c, msg)
		io.WriteString(c.writer, msg)
	}
}

func CreateChannel(name string, level int, writer io.Writer, format string) {
	channels[name] = &channel{name, level, writer, format}
}
