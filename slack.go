package main

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

func buildMessage(sm SlackMessage) (options slack.MsgOption, err error) {

	// Check the Message value to see if it is provided, if it is, it wins over BlockKit
	if len(sm.Message) > 0 {
		log.Info("Building message...")
		msgText := slack.NewTextBlockObject("mrkdwn", sm.Message, false, false)
		msgSection := slack.NewSectionBlock(msgText, nil, nil)
		options := slack.MsgOptionBlocks(
			msgSection,
		)

		return options, nil
	} else {
		log.Info("No message provided, assuming it's in the Blocks and proceeding...")
	}

	// Read the input blocks if they exist into capturedBlocks, which we can use to build the
	// the message if the Message value wasn't passed.
	if sm.Blocks != "" {
		log.Debug("Found Blocks (blockit) being passed in so I'm trying to unmarshall it...")
		var capturedBlocks slack.Blocks
		err = json.Unmarshal([]byte(sm.Blocks), &capturedBlocks)
		if err != nil {
			return nil, err
		}
		options := slack.MsgOptionBlocks(capturedBlocks.BlockSet...)

		return options, nil
	} else {
		return nil, fmt.Errorf("both Message and Blocks were undefined, you must provide one")
	}
}

func getTargetId(client *slack.Client, target string) (id string, err error) {
	if isEmail(target) {
		log.Info("We received an email address, trying to match against a Slack ID...")
		user, err := client.GetUserByEmail(target)
		if err != nil {
			return "", err
		}

		return user.ID, nil
	}

	log.Info("Target is not an email so we're assuming it's a channel name: ", target)
	groups, _, err := client.GetConversations(&slack.GetConversationsParameters{
		Limit: 500000,
	})
	if err != nil {
		return "", err
	}

	// We've got a list of public channels so we'll scan them for the correct name
	// this is so that we can find the ID.
	log.Info("Searching for the channel...")
	for _, group := range groups {
		if group.Name == target {
			log.Infof("Found group/channel %v", target)
			return group.ID, nil
		}
	}

	return "", fmt.Errorf("we didn't manage to find the group you asked for (%v)", target)
}
