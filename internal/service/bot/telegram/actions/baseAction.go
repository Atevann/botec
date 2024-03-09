package actions

import "gopkg.in/telebot.v3"

// Базовый абстрактный экшен
type BaseAction struct {
	Ctx telebot.Context
}

func (action *BaseAction) SetContext(ctx telebot.Context) {
	action.Ctx = ctx
}
