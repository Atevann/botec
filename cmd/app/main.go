package main

import (
	"Atevann/botec/config"
	"Atevann/botec/internal/service/bot/telegram"
	telebot "gopkg.in/telebot.v3"
	"log"
)

func main() {
	Config := config.NewConfig()

	b, err := telegram.NewBot(Config.Bot.Token)

	if err != nil {
		log.Fatal("Ошибка создания бота")
	}

	b.Bot.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send(c.Text())
	})

	b.Bot.Start()
}
