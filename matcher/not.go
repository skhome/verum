package matcher

import (
	"github.com/skhome/verum/message"
)

// Not negates the condition of the given matcher.
func Not[A any](m Matcher[A]) Matcher[A] {
	return &matcher[A]{
		matcherFn: func(actual A) bool {
			return !m.Matches(actual)
		},
		describeFn: func(r message.Representation) string {
			return "not " + m.Describe(r)
		},
	}
}
