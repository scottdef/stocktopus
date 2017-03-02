[![Go Report Card](https://goreportcard.com/badge/github.com/thourfor/stocktopus)](https://goreportcard.com/report/github.com/thourfor/stocktopus)
[![WTFPL licensed](https://img.shields.io/badge/license-WTFPL-blue.svg)]
(https://github.com/thourfor/stocktopus/blob/master/LICENSE)

<a href="https://slack.com/oauth/authorize?scope=commands&client_id=15348769670.121517816146"><img alt="Add to Slack" height="40" width="139" src="https://platform.slack-edge.com/img/add_to_slack.png" srcset="https://platform.slack-edge.com/img/add_to_slack.png 1x, https://platform.slack-edge.com/img/add_to_slack@2x.png 2x" /></a>

<img src="https://github.com/thourfor/stocktopus/blob/master/stocktopus_cropped.png" height="285" width="132"/>

#stocktopus
Simple Slack bot that posts stock prices. It can be build as an RTM Slack bot, or a slash command bot that loads into aws lambda

## Build
`go build -tags RTM`
or for aws
`go build`

## Run
`./stocktopus [slack-bot-token]`
or for aws
`./aws/zipit.sh`
and upload the stocktopus.zip to lambda

## Usage
The RTM bot will look for any direct messages sent to it and try to pase them as tickers, and respond with stock quotes.
> @stockbotname GOOGL

The aws slash command will respond to slash commands. Single tickers will be a quote and inline graph. 
> /stocktopus GOOGL




for a complete list of commands the bot supports.
> /stocktopus help 

# Privacy Policy

Stocktopus does not collect any personal identifying information. It does not store a history of your requests. The only data it does store is a unique ID received from Slack for a user if they opt to use the list or play money features. If you use the list or play money features it also stores the list of stocks a user has bought or added to their watch list. 

# Contact

If you have questions comments or concerns about the app please email thorvald@protonmail.ch



#### Special Thanks to /u/shirokarasu over at r/DrawForMe for creating the icon
