package commands

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework-2/pkg/api"
)

var IncompleteCommandError = fmt.Errorf("Please, enter the ticker after command. Like this: \n /command <ticker> (without brackets)")

type CurrentCommand struct {
	Command, SubCommand string
}

func (currentCommand *CurrentCommand) Clear() {
	currentCommand.Command = ""
	currentCommand.SubCommand = ""
}

type BotCommandFactory struct {
	Client investapi.SharesServiceClient
	Bot    *tgbotapi.BotAPI
}

type BotCommand interface {
	Start(ctx context.Context) (text string, err error)
	GetShare(ctx context.Context, update tgbotapi.Update) (text string, err error)
}
