package telegram

import (
	telebot "gopkg.in/telebot.v3"
	"time"
)

// Bot структура телеграм бота
type Bot struct {
	Bot *telebot.Bot
}

// NewBot Создание нового бота по токену
func NewBot(token string) (*Bot, error) {
	bot, err := telebot.NewBot(
		telebot.Settings{
			Token:  token,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

	if err != nil {
		return nil, err
	}

	return &Bot{Bot: bot}, nil
}
