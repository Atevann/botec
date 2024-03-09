package repositories

import (
	"botec/internal/service/bot/telegram/actions"
)

// ActionsList Мапа существующих экшенов
type ActionsList struct {
	actions map[string]actions.ActionInterface
}

// NewActionsList Регистрация экшенов
func NewActionsList() *ActionsList {
	return &ActionsList{
		actions: map[string]actions.ActionInterface{
			"SendMessage": new(actions.SendMessage),
			"SendButtons": new(actions.SendButtons),
		},
	}
}

// GetOneByName Возвращает экшен по имени
func (al *ActionsList) GetOneByName(name string) (actions.ActionInterface, bool) {
	action, isPresent := al.actions[name]
	return action, isPresent
}
