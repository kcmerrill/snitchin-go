package snitchin

import (
	"flag"
	"fmt"
	"strings"
)

var verbosity map[int]*bool

type level struct {
	name   string
	value  int
	prefix string
}

var levels map[string]*level = make(map[string]*level)

func CLIFlags() {
	verbosity = make(map[int]*bool)
	vees := ""
	/* Setup our --vvvvvv array */
	for x := 700; x >= 0; x -= 100 {
		vees += "v"
		verbosity[x] = flag.Bool(vees, false, fmt.Sprintf("Log verbosity level set to %d", x))
	}
}

func CLILogLevel() int {
	for level, f := range verbosity {
		if *f {
			return level
		}
	}
	return 0
}

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
