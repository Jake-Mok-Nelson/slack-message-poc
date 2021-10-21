# slack-message-poc
Sending a message to a channel or user with Slack API.

***Note: This is a POC/Spike only and is not production ready***

## Requirements

See the app manifest.yaml for the app as code.

You also need to export the Slack app Oauth token as an environment variable `export SLACK_API_TOKEN=xxxx-xxxx-xxxx-xxxx`

## Payload

Include a target (**either** a single email or channel name) as a string.

Include **either** a Message or Blocks (Message takes presedence)

![picture 1](images/9d48c586a2e691e6e65ed1a6b8a497bb3eb8d0c762ede76f3f66a3008e447e9a.png)  

![picture 2](images/b56909f8b632b95308ca0e4673e03e83e56e79675ce9b02da316d15ce34baee7.png)  

See slackmessage_test.go as an implementation example.
