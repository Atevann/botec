package telegram

import (
	"botec/internal/repositories"
	telebot "gopkg.in/telebot.v3"
	"log"
	"sync"
	"time"
)

// Bot структура телеграм бота
type Bot struct {
	Bot *telebot.Bot
}

// InitBots инициализирует всех ботов
func InitBots(botRepo *repositories.BotRepository) error {
	bots, err := botRepo.GetAll()

	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, bot := range bots {
		wg.Add(1)
		runningBot, err := NewBot(bot.Token)

		if err != nil {
			log.Println("Ошибка запуска бота")
			continue
		}

		go runningBot.serve(&wg)
	}

	wg.Wait()

	return nil
}

// serve Запуск бота
func (Bot *Bot) serve(wg *sync.WaitGroup) {
	defer func() {
		Bot.Bot.Stop()
		wg.Done()
	}()

	Bot.Bot.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send(c.Text())
	})

	Bot.Bot.Start()
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
