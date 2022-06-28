package commands

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework-2/external/investapi"
	"strconv"
	"strings"
)

func (botCommand *BotCommandFactory) GetShare(ctx context.Context, update tgbotapi.Update) (text string, err error) {
	commandTicker := strings.Split(update.Message.Text, " ")
	if len(commandTicker) < 2 {
		return "", IncompleteCommandError
	}
	ticker := commandTicker[1]
	share, err := investapi.GetShare(ticker)
	if err != nil {
		text = err.Error()
		return text, err
	}
	_, err = botCommand.Client.AddShare(ctx, &share)
	if err != nil {
		return err.Error(), err
	}
	text = "Name: " + share.Name + "\n" +
		"Ticker: " + share.Ticker + "\n" +
		"FIGI: " + share.Figi + "\n" +
		"Currency: " + share.Currency + "\n" +
		"Exchange: " + share.ClassCode + "\n" +
		"Country of risk: " + share.CountryOfRiskName + "\n" +
		"Issue size: " + strconv.FormatInt(share.IssueSize, 10)
	return
}
