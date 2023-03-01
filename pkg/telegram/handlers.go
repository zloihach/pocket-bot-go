package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/url"
)

const (
	commandStart           = "start"
	replyStartTemplate     = "Hello, I am Pocket Bot. For start using me, please, authorize me in Pocket. For authorize click to this link:\n%s"
	replyAlreadyAuthorized = "You are already authorized in the bot. Please send me the link and I'll save it"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	_, err := url.ParseRequestURI(message.Text)

	if err != nil {
		return errInvalidURL
	}

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return errUnauthorized
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		return errUnableToSave
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Url has been saved successful")
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, replyAlreadyAuthorized)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I dont know this command")
	_, err := b.bot.Send(msg)
	return err
}
