package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"homework-2/internal/commands"
	"homework-2/internal/dict"
	"homework-2/pkg/api"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Подключение к серверу
	var addr = "localhost:8082"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	client := investapi.NewSharesServiceClient(conn)
	ctx := context.Background()

	// Инициализация бота
	log.Printf(os.Getenv("TELEGRAM_APITOKEN"))
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	var command = commands.BotCommandFactory{
		Client: client,
		Bot:    bot,
	}

	// Обработка команд
	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message == nil || !update.Message.IsCommand() {
			msg.Text = "Please submit your request like this:\n /command <ticker> (without brackets)"
		} else {
			switch update.Message.Command() {
			case dict.CommandStart:
				msg.Text, err = command.Start(ctx)
			case dict.CommandGetShare:
				msg.Text, err = command.GetShare(ctx, update)
			case dict.CommandGetLastPrice:
				msg.Text, err = command.GetSharePrice(ctx, update)
			default:
				msg.Text = "I don't know that command"
			}
		}
		if err != nil {
			msg.Text = err.Error()
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
