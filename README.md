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

To provide the Blocks structure you can build it using the BlockKit [here](https://app.slack.com/block-kit-builder/TMR2N9HHC#%7B%22blocks%22:%5B%7B%22type%22:%22header%22,%22text%22:%7B%22type%22:%22plain_text%22,%22text%22:%22New%20request%22,%22emoji%22:true%7D%7D,%7B%22type%22:%22section%22,%22fields%22:%5B%7B%22type%22:%22mrkdwn%22,%22text%22:%22*Type:*%5CnPaid%20Time%20Off%22%7D,%7B%22type%22:%22mrkdwn%22,%22text%22:%22*Created%20by:*%5Cn%3Cexample.com%7CFred%20Enriquez%3E%22%7D%5D%7D,%7B%22type%22:%22section%22,%22fields%22:%5B%7B%22type%22:%22mrkdwn%22,%22text%22:%22*When:*%5CnAug%2010%20-%20Aug%2013%22%7D,%7B%22type%22:%22mrkdwn%22,%22text%22:%22*Type:*%5CnPaid%20time%20off%22%7D%5D%7D,%7B%22type%22:%22section%22,%22fields%22:%5B%7B%22type%22:%22mrkdwn%22,%22text%22:%22*Hours:*%5Cn16.0%20(2%20days)%22%7D,%7B%22type%22:%22mrkdwn%22,%22text%22:%22*Remaining%20balance:*%5Cn32.0%20hours%20(4%20days)%22%7D%5D%7D,%7B%22type%22:%22section%22,%22text%22:%7B%22type%22:%22mrkdwn%22,%22text%22:%22%3Chttps://example.com%7CView%20request%3E%22%7D%7D%5D%7D) and grab pass the blocks array into `SlackMessage.Blocks`.