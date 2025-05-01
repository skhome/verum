package matcher

import (
	"errors"
	"fmt"

	"github.com/skhome/verum/message"
)

// ErrorIs verifies that at least one error in the error chain matches the given error.
//
//	err := errors.New("error")
//
//	// assertions will pass
//	assert.That(t, err).Matches(matcher.ErrorIs(err))
//
//	// assertions will fail
//	assert.That(t, errors.New("other")).Matches(matcher.ErrorIs(err))
func ErrorIs(target error) Matcher[error] {
	return &matcher[error]{
		matcherFn: func(actual error) bool {
			return errors.Is(actual, target)
		},
		describeFn: func(representation message.Representation) string {
			return fmt.Sprintf("error to have %s in its error chain", representation(target))
		},
	}
}

// HasMessage verifies that the actual error has the given error message.
//
//	err := errors.New("no such file")
//
//	// assertions will pass
//	assert.That(t, err).Matches(matcher.HasMessage("no such file"))
//
//	// assertions will fail
//	assert.That(t, err).Matches(matcher.HasMessage("file not found"))
func HasMessage(msg string) Matcher[error] {
	return &matcher[error]{
		matcherFn: func(actual error) bool {
			return actual.Error() == msg
		},
		describeFn: func(representation message.Representation) string {
			return "error to have message " + representation(msg)
		},
	}
}
