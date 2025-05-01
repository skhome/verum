package matcher

import (
	"github.com/skhome/verum/message"
)

const descriptionNil = "value to be nil"

// IsNilError verifies that the actual error is nil.
//
//	// assertions will pass
//	assert.That(t, nil).Matches(matcher.IsNilError())
//
//	// assertions will fail
//	assert.That(t, errors.New("other")).Matches(matcher.IsNilError())
func IsNilError() Matcher[error] {
	return &matcher[error]{
		matcherFn: func(actual error) bool {
			return actual == nil
		},
		describeFn: func(message.Representation) string {
			return descriptionNil
		},
	}
}

// IsNilMap verifies that the actual map is nil.
//
//	// assertions will pass
//	assert.That(t, nil).Matches(matcher.IsNilMap())
//
//	// assertions will fail
//	assert.That(t, map[string]string{"Nenya": "Galadriel"}).Matches(matcher.IsNilMap())
func IsNilMap[K comparable, V any]() Matcher[map[K]V] {
	return &matcher[map[K]V]{
		matcherFn: func(actual map[K]V) bool {
			return actual == nil
		},
		describeFn: func(message.Representation) string {
			return descriptionNil
		},
	}
}

// IsNil verifies that the actual pointer is nil.
//
//	// assertions will pass
//	assert.That(t, nil).Matches(matcher.IsNil())
//
//	// assertions will fail
//	assert.That(t, ptr).Matches(matcher.IsNil())
func IsNil[T any]() Matcher[*T] {
	return &matcher[*T]{
		describeFn: func(message.Representation) string {
			return descriptionNil
		},
		matcherFn: func(actual *T) bool {
			return actual == nil
		},
	}
}
