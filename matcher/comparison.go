package matcher

import (
	"cmp"
	"fmt"

	"github.com/skhome/verum/message"
)

// IsEqualTo verifies that the actual value is equal to the given one.
//
//	// assertions will pass
//	assert.That(t, "Frodo").Matches(IsEqualTo("Frodo"))
//
//	// assertions will fail
//	assert.That(t, "Frodo").Matches(IsEqualTo("Sam"))
func IsEqualTo[T comparable](expected T) Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual == expected
		},
		describeFn: func(representation message.Representation) string {
			return "value equal to " + representation(expected)
		},
	}
}

// IsLessThan verifies that the actual value is less than the given one.
//
//	// assertions will pass
//	assert.That(t, 41).Matches(IsLessThan(42))
//
//	// assertions will fail
//	assert.That(t, 42).Matches(IsLessThan(42))
func IsLessThan[T cmp.Ordered](expected T) Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual < expected
		},
		describeFn: func(representation message.Representation) string {
			return "value less than " + representation(expected)
		},
	}
}

// IsLessThanOrEqualTo verifies that the actual value is less than or equal to the given one.
//
//	// assertions will pass
//	assert.That(t, 42).Matches(IsLessThanOrEqualTo(42))
//	assert.That(t, 41).Matches(IsLessThanOrEqualTo(42))
//
//	// assertions will fail
//	assert.That(t, 43).Matches(IsLessThanOrEqualTo(42))
func IsLessThanOrEqualTo[T cmp.Ordered](expected T) Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual <= expected
		},
		describeFn: func(representation message.Representation) string {
			return "value less than or equal to " + representation(expected)
		},
	}
}

// IsGreaterThan verifies that the actual value is greater than the given one.
//
//	// assertions will pass
//	assert.That(t, 43).Matches(IsGreaterThan(42))
//
//	// assertions will fail
//	assert.That(t, 42).Matches(IsGreaterThan(42))
func IsGreaterThan[T cmp.Ordered](expected T) Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual > expected
		},
		describeFn: func(representation message.Representation) string {
			return "value greater than " + representation(expected)
		},
	}
}

// IsGreaterThanOrEqualTo verifies that the actual value is greater than or equal to the given one.
//
//	// assertions will pass
//	assert.That(t, 42).Matches(IsGreaterThanOrEqualTo(42))
//	assert.That(t, 43).Matches(IsGreaterThanOrEqualTo(42))
//
//	// assertions will fail
//	assert.That(t, 41).Matches(IsGreaterThanOrEqualTo(42))
func IsGreaterThanOrEqualTo[T cmp.Ordered](expected T) Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual >= expected
		},
		describeFn: func(representation message.Representation) string {
			return "value greater than or equal to " + representation(expected)
		},
	}
}

// IsBetween verifies that the actual value is between the given ones.
//
//	// assertions will pass
//	assert.That(t, 42).Matches(IsBetween(42, 43))
//	assert.That(t, 42).Matches(IsBetween(41, 42))
//
//	// assertions will fail
//	assert.That(t, 41).Matches(IsGreaterThanOrEqualTo(42))
func IsBetween[T cmp.Ordered](startInclusive T, endInclusive T) Matcher[T] {
	return &matcher[T]{
		matcherFn: func(actual T) bool {
			return actual >= startInclusive && actual <= endInclusive
		},
		describeFn: func(representation message.Representation) string {
			return fmt.Sprintf("value between %s and %s", representation(startInclusive), representation(endInclusive))
		},
	}
}
