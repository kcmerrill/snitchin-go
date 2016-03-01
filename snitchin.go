package snitchin

import (
	"os"
)

func DEBUG(msg string) {
	log("DEBUG", msg)
}

func INFO(msg string) {
	log("INFO", msg)
}

func NOTICE(msg string) {
	log("NOTICE", msg)
}

func WARNING(msg string) {
	log("WARNING", msg)
}

func ERROR(msg string) {
	log("ERROR", msg)
}

func CRITICAL(msg string) {
	log("CRITICAL", msg)
}

func ALERT(msg string) {
	log("ALERT", msg)
}

func EMERGENCY(msg string) {
	log("EMERGENCY", msg)
}

func log(level_name string, msg string) {
	for _, channel := range channels {
		channel.Log(level_name, msg)
	}
}

func init() {
	//Setup our custom log levels
	CustomLevel("DEBUG", 100, "\033[38;5;0m")
	CustomLevel("INFO", 200, "\033[38;5;231m")
	CustomLevel("NOTICE", 250, "\033[38;5;42m")
	CustomLevel("WARNING", 300, "\033[38;5;220m")
	CustomLevel("ERROR", 400, "\033[38;5;1m")
	CustomLevel("CRITICAL", 500, "\033[38;5;88m")
	CustomLevel("EMERGENCY", 600, "\033[38;5;160m")
	CustomLevel("ALERT", 700, "\033[38;5;123m")
	CustomLevel("SUCCESS", 700, "\033]02;36\]")

	//Lets create a basic formatter
	CreateFormatter("basic-color", "%%PREFIX%%[%%TIME%%] [%%LEVEL%%] %%MSG%%\033[0m\n")
	CreateFormatter("basic", "[%%TIME%%] [%%LEVEL%%] %%MSG%%\n")

	//Lets create a default channel
	CreateChannel("DEFAULT", 400, os.Stdout, "basic-color")
}
