package matcher_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/skhome/verum/assert"
	"github.com/skhome/verum/matcher"
)

func TestErrorIs(t *testing.T) {
	t.Parallel()
	cause := errors.New("cause")
	tests := []struct {
		name   string
		actual error
		target error
		ok     bool
	}{
		{name: "same", actual: cause, target: cause, ok: true},
		{name: "wrapped", actual: fmt.Errorf("because: %w", cause), target: cause, ok: true},
		{name: "different", actual: errors.New("another"), target: cause, ok: false},
	}
	messageFormat := "expected: error to have <%v> in its error chain"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.target)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.ErrorIs(tt.target))
		})
	}
}

func TestHasMessage(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		actual  error
		message string
		ok      bool
	}{
		{name: "same", actual: errors.New("file not found"), message: "file not found", ok: true},
		{name: "different", actual: errors.New("file not found"), message: "no such file", ok: false},
	}
	messageFormat := "expected: error to have message <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.message)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.HasMessage(tt.message))
		})
	}
}
