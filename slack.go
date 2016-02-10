package snitchin

import (
	"bytes"
	"encoding/json"
	"net/http"
)

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
		ERROR("There was an error posting to slack")
	}

	/* TODO: Come back to this in a bit */
	/* Not sure what to return here. I'm assuming it's how many bytes were written, then errors? */
	return 0, nil
}
