package main

import (
	"fmt"
	"net/mail"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

const (

	// Public Examples
	target     = "git-notifications-test-314141" // public channel example
	userToPing = "jakenelson@myemail.com"
	sender     = "My Notification App"
)

// isEmail checks if the string is an email address
func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		log.Fatal("Is your SLACK_API_TOKEN set?")
	}

	log.Info("Authenticating to Slack API...")
	client := slack.New(token)
	_, err := client.AuthTest()
	if err != nil {
		log.Fatal(err)
	}

	// Determine if the targer is a user or a channel.
	var targetID string
	var targetFound = false
	userTag, err := client.GetUserByEmail(userToPing)
	if err != nil {
		log.Fatal(err)
	}

	if isEmail(target) {
		log.Info("Target is an a user. Email: ", target)
		user, err := client.GetUserByEmail(target)
		targetID = user.ID
		if err != nil {
			log.Fatal(err)
		}
		targetFound = true
	} else {
		log.Info("Target is a channel: ", target)
		groups, _, err := client.GetConversations(&slack.GetConversationsParameters{
			Limit: 50000,
		})
		if err != nil {
			fmt.Printf("%s\n", err)
			log.Fatal(err)
		}
		for _, group := range groups {
			fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
			if group.Name == target {
				log.Infof("Found group/channel %v", target)
				targetFound = true
				targetID = group.ID

				// log.Debug("Joining the conversation for ", group.ID)
				// _, _, _, err = client.OpenConversation(&slack.OpenConversationParameters{
				// 	ChannelID: group.ID,
				// })
				// if err != nil {
				// 	log.Fatal(err)
				// }
				break
			}
		}
	}

	if !targetFound {
		log.Fatalf("We couldn't find the target (%v) you requested.")
	}

	// Get UserID or Channel ID.

	// Compose a message
	log.Infof("Building message for target: %v", targetID)
	message := fmt.Sprintf("Hi, this is a test from %v\n\nOkay, it's not really, I'm actually just a slack app that's pretending.\n\n Here, have an emoji? 💡\n\nHow about a code block?\n```Step 1. Build Github App\nStep 2. ?\nStep 3. Profit!```\nNow let's try pinging someone like @%v", sender, userTag.Name)
	msgText := slack.NewTextBlockObject("mrkdwn", message, false, false)
	msgSection := slack.NewSectionBlock(msgText, nil, nil)
	msg := slack.MsgOptionBlocks(
		msgSection,
	)

	// Send private message
	log.Debug("Sending message...")
	_, _, err = client.PostMessage(targetID, msg)
	if err != nil {
		panic(err)
	}
	println("Message sent.")
}
