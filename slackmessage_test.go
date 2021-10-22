package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Systems/integration tests

func TestMissingMessageFailure(t *testing.T) {

	payload := SlackMessage{
		Target: "someemail@somedomain.com",
	}

	err := payload.Send()
	assert.EqualErrorf(t, err, "both Message and Blocks were undefined, you must provide one", "")
	//assert.Error(t, err)
}

func TestTargetMissing(t *testing.T) {

	payload := SlackMessage{
		Message: "A message that should fail because there's no target",
	}

	err := payload.Send()
	assert.EqualErrorf(t, err, "we didn't manage to find the group/user you asked for ()", "")
	//assert.Error(t, err)
}

// func TestSimpleMessageSent(t *testing.T) {
// 	payload := SlackMessage{
// 		Target:  "someemail@somedomain.com",
// 		Message: "This is a test message",
// 	}

// 	err := payload.Send()
// 	assert.Nil(t, err)
// }

// func TestBlockKitMessageSent(t *testing.T) {
// 	blockJson := `[
// 			{
// 				"type": "header",
// 				"text": {
// 					"type": "plain_text",
// 					"text": "New request",
// 					"emoji": true
// 				}
// 			},
// 			{
// 				"type": "section",
// 				"fields": [
// 					{
// 						"type": "mrkdwn",
// 						"text": "*Type:*\nPaid Time Off"
// 					},
// 					{
// 						"type": "mrkdwn",
// 						"text": "*Created by:*\n<example.com|Fred Enriquez>"
// 					}
// 				]
// 			},
// 			{
// 				"type": "section",
// 				"fields": [
// 					{
// 						"type": "mrkdwn",
// 						"text": "*When:*\nAug 10 - Aug 13"
// 					},
// 					{
// 						"type": "mrkdwn",
// 						"text": "*Type:*\nPaid time off"
// 					}
// 				]
// 			},
// 			{
// 				"type": "section",
// 				"fields": [
// 					{
// 						"type": "mrkdwn",
// 						"text": "*Hours:*\n16.0 (2 days)"
// 					},
// 					{
// 						"type": "mrkdwn",
// 						"text": "*Remaining balance:*\n32.0 hours (4 days)"
// 					}
// 				]
// 			},
// 			{
// 				"type": "section",
// 				"text": {
// 					"type": "mrkdwn",
// 					"text": "<https://example.com|View request>"
// 				}
// 			}
// 		]`
// 	payload := SlackMessage{
// 		Target: "someemail@somedomain.com",
// 		Blocks: blockJson,
// 	}

// 	err := payload.Send()
// 	assert.Nil(t, err)
// }
