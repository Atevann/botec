package actions

import (
	"botec/internal/service/bot/telegram/actions/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendMessage_validate(t *testing.T) {
	testCases := []struct {
		name   string
		action *SendMessage
		error  error
	}{
		{
			name:   "Empty message text",
			action: &SendMessage{Text: ""},
			error:  errors.ValidationError{Reason: "Message text is empty"},
		},
		{
			name:   "Valid case",
			action: &SendMessage{Text: "Some text"},
			error:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.action.validate()

			if tc.error == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, err, tc.error)
			}
		})
	}
}
