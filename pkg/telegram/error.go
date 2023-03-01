package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	errInvalidURL   = errors.New("Url is invalid")
	errUnauthorized = errors.New("User is not authorized")
	errUnableToSave = errors.New("Unable to save")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "Unrecognized Error :(")

	switch err {
	case errInvalidURL:
		msg.Text = "Error. Url is not valid"
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = "You are not authorized! Write /start "
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = "The link was not saved, try again"
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
