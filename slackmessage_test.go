package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMessageSent(t *testing.T) {
	payload := SlackMessage{
		Target:  "myemail.email.com",
		Message: "This is a test message",
	}

	err := payload.Send()
	assert.Nil(t, err)
}

func TestBlockKitMessageSent(t *testing.T) {
	blockJson := `[
			{
				"type": "header",
				"text": {
					"type": "plain_text",
					"text": "New request",
					"emoji": true
				}
			},
			{
				"type": "section",
				"fields": [
					{
						"type": "mrkdwn",
						"text": "*Type:*\nPaid Time Off"
					},
					{
						"type": "mrkdwn",
						"text": "*Created by:*\n<example.com|Fred Enriquez>"
					}
				]
			},
			{
				"type": "section",
				"fields": [
					{
						"type": "mrkdwn",
						"text": "*When:*\nAug 10 - Aug 13"
					},
					{
						"type": "mrkdwn",
						"text": "*Type:*\nPaid time off"
					}
				]
			},
			{
				"type": "section",
				"fields": [
					{
						"type": "mrkdwn",
						"text": "*Hours:*\n16.0 (2 days)"
					},
					{
						"type": "mrkdwn",
						"text": "*Remaining balance:*\n32.0 hours (4 days)"
					}
				]
			},
			{
				"type": "section",
				"text": {
					"type": "mrkdwn",
					"text": "<https://example.com|View request>"
				}
			}
		]`
	payload := SlackMessage{
		Target: "myemail.email.com",
		Blocks: blockJson,
	}

	err := payload.Send()
	assert.Nil(t, err)
}
