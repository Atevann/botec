package actions

import "gopkg.in/telebot.v3"

// ActionInterface Интерфейс экшенов
type ActionInterface interface {
	// SetContext Устанавливает контекст экшену
	SetContext(ctx telebot.Context)
	// Execute Выполняет экшен
	Execute() error
	// validate Валидирует переданные в экшен данные
	validate() error
}
