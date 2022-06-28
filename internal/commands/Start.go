package commands

import (
	"context"
)

func (botCommand BotCommandFactory) Start(ctx context.Context) (text string, err error) {
	text = "Welcome!\n" +
		"This bot was developed as part of a training course Route256.\n" +
		"The functionality of the bot is to issue the share price by ticker.\n" +
		"Please enter stock ticker to get the price."
	return
}
