package main

import (
	"github.com/boltdb/bolt"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	"pocket-bot-go/pkg/repository"
	"pocket-bot-go/pkg/repository/boltdb"
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

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "https://localhost/")

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		return err
	}); err != nil {
		return nil, err
	}
	return db, err
}
