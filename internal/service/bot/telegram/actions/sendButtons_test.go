package actions

import (
	"botec/internal/service/bot/telegram/actions/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendButtons_validate(t *testing.T) {
	testCases := []struct {
		name   string
		action *SendButtons
		error  error
	}{
		{
			name:   "Empty message text",
			action: &SendButtons{Text: "", Buttons: []string{"Button1", "Button2"}},
			error:  errors.ValidationError{Reason: "Message text is empty"},
		},
		{
			name:   "No buttons provided",
			action: &SendButtons{Text: "Some text", Buttons: nil},
			error:  errors.ValidationError{Reason: "No buttons provided"},
		},
		{
			name:   "Empty button text",
			action: &SendButtons{Text: "Some text", Buttons: []string{"Button1", ""}},
			error:  errors.ValidationError{Reason: "Button text is empty"},
		},
		{
			name:   "Valid case",
			action: &SendButtons{Text: "Some text", Buttons: []string{"Button1", "Button2"}},
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
