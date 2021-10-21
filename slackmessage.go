package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

type SlackMessage struct {
	Target  string
	Message string `json:"message,omitempty"`
	Blocks  string `json:"blocks,omitempty"`
}

func (sm SlackMessage) Send() error {
	log.SetFormatter(&log.JSONFormatter{})
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		return fmt.Errorf("missing SLACK_API_TOKEN variable")
	}

	log.Info("Authenticating to Slack API...")
	client := slack.New(token)
	_, err := client.AuthTest()
	if err != nil {
		return err
	}

	targId, err := getTargetId(client, sm.Target)
	if err != nil {
		return err
	}

	messageOptions, err := buildMessage(sm)
	if err != nil {
		return err
	}

	log.Info("Sending message...")
	_, _, _, err = client.SendMessage(targId, messageOptions)
	if err != nil {
		return err
	}
	log.Info("Message sent.")
	return nil

}
