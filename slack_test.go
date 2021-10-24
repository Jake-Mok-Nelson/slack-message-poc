package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Systems/integration tests

func TestReplaceEmails(t *testing.T) {
	mo := SlackMessageOptions{
		TranslateEmails: true,
	}
	sm := SlackMessage{
		Message:        "A test message with an email jake.nelson@anz.com",
		MessageOptions: mo,
	}
	msg, err := buildMessage(sm)
	assert.Nil(t, err)

	replaceEmails(msg)

	//assert.Error(t, err)
}
