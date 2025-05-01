package matcher

import "github.com/skhome/verum/message"

// IsTrue verifies that the actual value is true.
//
//	// assertions will pass
//	assert.That(t, true).Matches(IsTrue())
//
//	// assertions will fail
//	assert.That(t, false).Matches(IsTrue())
func IsTrue() Matcher[bool] {
	return &matcher[bool]{
		matcherFn: func(actual bool) bool {
			return actual
		},
		describeFn: func(_ message.Representation) string {
			return "value to be true"
		},
	}
}

// IsFalse verifies that the actual value is false.
//
//	// assertions will pass
//	assert.That(t, false).Matches(IsFalse())
//
//	// assertions will fail
//	assert.That(t, true).Matches(IsFalse())
func IsFalse() Matcher[bool] {
	return &matcher[bool]{
		matcherFn: func(actual bool) bool {
			return !actual
		},
		describeFn: func(_ message.Representation) string {
			return "value to be false"
		},
	}
}
