# snitchin-go
A multi channel log emitter(a clone of my php snitchin) written in GO

# Basics
Create as many or as few channels as you want, each with their own logging levels. The best way to demonstrate it is via code.

```
  package main

  import (
      log "github.com/kcmerrill/snitchin-go"
  )

  func main() {
      /* Log to a file */
      log.CreateChannel("TOAFILE", 100, log.File("/tmp/logger.txt"), "basic")
      /* Log to slack */
      log.CreateChannel("SLACK", 700, log.Slack("https://hooks.slack.com/services/customslackwebhookurl"), "%%MSG%%")

      log.DEBUG("HELLO DEBUG!!")
      log.INFO("hello worldzzzzz")
      log.NOTICE("THIS IS A NOTICE")
      log.WARNING("THIS IS A WARNING")
      log.ERROR("an error would go here")
      log.CRITICAL("this would be a critical, or would it be?")
      log.EMERGENCY("An emergency? Yes ... yes it is ..")
      log.ALERT("this is an alert!")

      /* If you only want to log to a specific channel, Also, notice the dynamic channel name */
      log.Channel("SECURITY").Log("ALERT", "My message would go here")

      /* Notice the custom log level */
      log.Channel("SECURITY").Log("CUSTOMLOGLEVEL", "My message would go here")
  }

```
