package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

type SlackMessageOptions struct {
	TranslateEmails   bool `json:"translate-emails,omitempty"`   // translate any email addresses into users and tag them if they exist in the available teams
	TranslateChannels bool `json:"translate-channels,omitempty"` // translate any matches of '#<name>' to mention a channel if it exists
}
type SlackMessage struct {
	Target         string
	Message        string              `json:"message,omitempty"`
	Blocks         string              `json:"blocks,omitempty"`
	MessageOptions SlackMessageOptions `json:"message-options,omitempty"`
}

func (sm SlackMessage) Send() error {

	// Build the logger
	log.SetFormatter(&log.JSONFormatter{})
	switch logginglevel := os.Getenv("SLACK_MESSAGE_VERBOSITY"); logginglevel {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		// when verbosity isn't set, we'll use the default (info level)
		log.SetLevel(log.InfoLevel)
	}

	// Read the SLACK token from environment variables
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
