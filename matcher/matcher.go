package matcher

import "github.com/skhome/verum/message"

// Matcher can verify and describe a condition
type Matcher[T any] interface {
	// Matches returns true if the condition is met
	Matches(actual T) bool
	// Describe describes the matcher (e.g. "a value equal to")
	Describe(representation message.Representation) string
}

type (
	matcherFunc[T any] func(actual T) bool
	describeFunc       func(representation message.Representation) string
)

type matcher[T any] struct {
	matcherFn  matcherFunc[T]
	describeFn describeFunc
}

func (m *matcher[T]) Matches(actual T) bool {
	return m.matcherFn(actual)
}

func (m *matcher[T]) Describe(representation message.Representation) string {
	return m.describeFn(representation)
}
