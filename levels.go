package snitchin

import (
	"strings"
)

type level struct {
	name   string
	value  int
	prefix string
}

var levels map[string]*level = make(map[string]*level)

func Level(name string) (string, int, string) {
	name = strings.ToUpper(name)
	if level, ok := levels[name]; ok {
		return level.name, level.value, level.prefix
	} else {
		CustomLevel(name, 700, "\033[38;5;231m")
	}
	return Level(name)
}

func CustomLevel(name string, value int, prefix string) {
	name = strings.ToUpper(name)
	levels[name] = &level{name, value, prefix}
}
