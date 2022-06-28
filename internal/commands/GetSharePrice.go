package commands

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework-2/external/investapi"
	_ "homework-2/pkg/api"
	"strings"
)

func (botCommand *BotCommandFactory) GetSharePrice(ctx context.Context, update tgbotapi.Update) (text string, err error) {
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
	lastPrice, err := investapi.GetSharePrice(&share)
	if err != nil {
		return err.Error(), err
	}
	price := float64(lastPrice[0].Price.Units) + float64(lastPrice[0].Price.Nano)/1000000000
	text = fmt.Sprintf("Last price of %s: %.2f %s", share.Name, price, strings.ToUpper(share.Currency))
	return
}
