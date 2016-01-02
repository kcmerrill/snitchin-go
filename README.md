# snitchin-go
A multi channel log emitter(a clone of my php snitchin) written in GO

# Basics
Create as many or as few channels as you want, each with their own logging levels. The best way to demonstrate it is via code.

```
package main

import (
        log "github.com/kcmerrill/snitchin-go"
        "os"
)

func main() {
        /* Create a new logger */
        s := log.New()

        /* Define some channels, and their log levels */
        s.AddChannel("CRIT-CHANNEL", log.CRITICAL, os.Stdout)
        s.AddChannel("DEFAULT", log.DEBUG, os.Stderr)
        s.AddChannel("TOFILE", log.WARNING, log.File("/tmp/log.txt"))
        s.AddChannel("TOSLACK", log.INFO, log.Slack("https://hooks.slack.com/services/yourwebhookurlhere"))

        /* Need to call out to a specific channel? Even if it doesn't exist? We'll create one with stdio for you */
        s.Channel("DOESNOTEXIST").Log(log.INFO, "this channel does not exist")

        /* Show some basic logs, goes to every channel */
        s.Log(log.CRITICAL, "this is a message to all channels")
        s.Log(log.DEBUG, "This should only be one log line")
        
        /* Log to a specific channel onyl */
        s.Channel("CRIT-CHANNEL").Log(log.CRITICAL, "My critical message here")
}

```
