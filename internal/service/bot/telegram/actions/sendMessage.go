package actions

import (
	"botec/internal/service/bot/telegram/actions/errors"
)

// SendMessage экшен отправки сообщения
type SendMessage struct {
	BaseAction
	Text string
}

func (action *SendMessage) Execute() error {
	err := action.validate()
	if err != nil {
		return err
	}

	return action.Ctx.Send(action.Text)
}

func (action *SendMessage) validate() error {
	if len(action.Text) == 0 {
		return errors.ValidationError{Reason: "Message text is empty"}
	}

	return nil
}
