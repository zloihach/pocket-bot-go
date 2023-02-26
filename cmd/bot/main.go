package main

import (
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	"pocket-bot-go/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6185260694:AAGuxr60fcxfFEZiwtb-wL2-2eksr90Cvl8")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("105911-f13979eb79848bf83a8fd8d")

	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, pocketClient, "https://localhost/")

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
