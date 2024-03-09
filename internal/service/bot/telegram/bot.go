package telegram

import (
	"botec/internal/models"
	"botec/internal/repositories"
	"encoding/json"
	"errors"
	"fmt"
	telebot "gopkg.in/telebot.v3"
	"log"
	"sync"
)

// InitBots инициализирует всех ботов
func InitBots(
	botRepo *repositories.BotRepository,
	actionsList *repositories.ActionsList,
	botActionsRepo *repositories.BotActionsRepository,
) error {
	bots, err := botRepo.GetAll()

	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, bot := range bots {
		wg.Add(1)
		runningBot, err := bot.Init()

		if err != nil {
			log.Printf("Ошибка запуска бота: %v", err)

			continue
		}

		go serve(&wg, runningBot, actionsList, botActionsRepo)
	}

	wg.Wait()

	return nil
}

// serve Запуск бота
func serve(
	wg *sync.WaitGroup,
	bot *models.Bot,
	actionsList *repositories.ActionsList,
	actionsRepo *repositories.BotActionsRepository,
) {
	defer func() {
		bot.Stop()
		wg.Done()
	}()

	bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
		botAction, err := actionsRepo.GetNextAction(bot.Id, 0, "")

		if err != nil {
			log.Printf("Ошибка при получении следующего экшена: %v", err)

			return err
		}

		action, isPresent := actionsList.GetOneByName(botAction.Action_name)

		if !isPresent {
			log.Printf("Экшен не существует: %v", botAction.Action_name)

			return errors.New(fmt.Sprintf("Экшен не существует: %v", botAction.Action_name))
		}

		action.SetContext(ctx)

		if json.Unmarshal([]byte(botAction.Action_data), &action) != nil {
			log.Printf("Ошибка декодирования экшена: %v", err)

			return err
		}

		return action.Execute()
	})

	bot.Start()
}
