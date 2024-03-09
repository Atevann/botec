package models

import (
	"gopkg.in/telebot.v3"
	"time"
)

// Bot Сущность бота
type Bot struct {
	Id    uint
	Token string
	tgbot *telebot.Bot
}

// Init Инициализация бота
func (bot *Bot) Init() (*Bot, error) {
	tgbot, err := telebot.NewBot(
		telebot.Settings{
			Token:  bot.Token,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

	if err != nil {
		return nil, err
	}

	bot.tgbot = tgbot

	return bot, nil
}

// Handle Добавляет хэндлер
func (bot *Bot) Handle(endpoint interface{}, handlerFunc telebot.HandlerFunc) {
	bot.tgbot.Handle(endpoint, handlerFunc)
}

// Stop останавливает бота
func (bot *Bot) Start() {
	bot.tgbot.Start()
}

// Stop останавливает бота
func (bot *Bot) Stop() {
	bot.tgbot.Stop()
}
