package actions

import (
	"botec/internal/service/bot/telegram/actions/errors"
	"gopkg.in/telebot.v3"
)

// Экшен отправки кнопок
type SendButtons struct {
	BaseAction
	Text    string
	Buttons []string
}

func (action *SendButtons) Execute() error {
	err := action.validate()
	if err != nil {
		return err
	}

	var rows []telebot.Row
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}

	for _, text := range action.Buttons {
		btn := menu.Text(text)
		row := menu.Row(btn)
		rows = append(rows, row)
	}

	menu.Reply(rows...)

	return action.Ctx.Send(action.Text, menu)
}

func (action *SendButtons) validate() error {
	if len(action.Text) == 0 {
		return errors.ValidationError{Reason: "Message text is empty"}
	}

	if len(action.Buttons) == 0 {
		return errors.ValidationError{Reason: "No buttons provided"}
	}

	for _, btn := range action.Buttons {
		if len(btn) == 0 {
			return errors.ValidationError{Reason: "Button text is empty"}
		}
	}

	return nil
}
