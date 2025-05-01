package matcher

import "github.com/skhome/verum/message"

// IsEmptySlice verifies that the actual slice is nil or empty.
//
//	// assertion will pass
//	assert.That(t, []string(nil)).Matches(matcher.IsEmptySlice())
//	assert.That(t, []string{}).Matches(matcher.IsEmptySlice())
//
//	// assertion will fail
//	assert.That(t, []string{"Hobbit"}).Matches(matcher.IsEmptySlice())
func IsEmptySlice[T any]() Matcher[[]T] {
	return &matcher[[]T]{
		matcherFn: func(actual []T) bool {
			return len(actual) == 0
		},
		describeFn: func(message.Representation) string {
			return "value to be true"
		},
	}
}
