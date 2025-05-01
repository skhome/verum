package matcher

import (
	"github.com/skhome/verum/message"
	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Integer | constraints.Float
}

// IsZero verifies that the actual value is zero.
//
//	// assertions will pass
//	assert.That(t, 0).Matches(matcher.IsZero())
//
//	// assertions will fail
//	assert.That(t, 42).Matches(matcher.IsZero())
func IsZero[T number]() Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual == 0
		},
		describeFn: func(message.Representation) string {
			return "value to be zero"
		},
	}
}

// IsPositive verifies that the actual value is positive.
//
//	// assertions will pass
//	assert.That(t, 1).Matches(matcher.IsPositive())
//
//	// assertions will fail
//	assert.That(t, 0).Matches(matcher.IsPositive())
func IsPositive[T number]() Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual > T(0)
		},
		describeFn: func(message.Representation) string {
			return "value to be positive"
		},
	}
}

// IsNegative verifies that the actual value is negative.
//
//	// assertions will pass
//	assert.That(t, -1).Matches(matcher.IsNegative())
//
//	// assertions will fail
//	assert.That(t, 0).Matches(matcher.IsNegative())
func IsNegative[T number]() Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual < T(0)
		},
		describeFn: func(message.Representation) string {
			return "value to be negative"
		},
	}
}

// IsEven verifies that the actual value is even.
//
//	// assertions will pass
//	assert.That(t, 2).Matches(matcher.IsEven())
//
//	// assertions will fail
//	assert.That(t, 3).Matches(matcher.IsEven())
func IsEven[T constraints.Integer]() Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual%2 == 0
		},
		describeFn: func(message.Representation) string {
			return "value to be even"
		},
	}
}

// IsOdd verifies that the actual value is odd.
//
//	// assertions will pass
//	assert.That(t, 1).Matches(matcher.IsOdd())
//
//	// assertions will fail
//	assert.That(t, 2).Matches(matcher.IsOdd())
func IsOdd[T constraints.Integer]() Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual%2 != 0
		},
		describeFn: func(message.Representation) string {
			return "value to be odd"
		},
	}
}
