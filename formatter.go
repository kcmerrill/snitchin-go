package snitchin

import (
	"strings"
	"time"
)

var formats map[string]string = make(map[string]string)

func CreateFormatter(name string, formatter string) {
	name = strings.ToUpper(name)
	formats[name] = formatter
}

func Formatter(level_name string, c channel, msg string) string {
	f := Format(c.format)
	_, level, prefix := Level(level_name)

	/* The message */
	f = strings.Replace(f, "%%MSG%%", msg, 1)

	/* The time */
	f = strings.Replace(f, "%%TIME%%", time.Now().Format(time.RFC822), 1)

	/* The level */
	f = strings.Replace(f, "%%LEVEL%%", level_name, 1)

	/* The level */
	f = strings.Replace(f, "%%LEVEL_INT%%", string(level), 1)

	/* The channel name */
	f = strings.Replace(f, "%%CHANNEL%%", c.name, 1)

	/* Prefix */
	f = strings.Replace(f, "%%PREFIX%%", prefix, 1)

	/* Return the goods */
	return f
}

func Format(name string) string {
	name_upper := strings.ToUpper(name)
	if f, ok := formats[name_upper]; ok {
		return f
	} else {
		return name
	}
}
