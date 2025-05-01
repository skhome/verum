package assert

import (
	"github.com/skhome/verum/matcher"
	"github.com/skhome/verum/message"
)

// TestingT is an interface wrapper for *testing.T
type TestingT interface {
	Errorf(format string, args ...any)
}

// tHelper is a helper interface to signal that this function is a test helper.
type tHelper interface {
	Helper()
}

// Assertion represents an assertion
type Assertion[T any] struct {
	t         TestingT
	info      *WritableAssertionInfo
	formatter message.Formatter
	actual    T
}

// DescribedAs sets an optional description for the following assertion.
func (a *Assertion[T]) DescribedAs(description string) *Assertion[T] {
	a.info.WithDescription(message.Description(description))
	return a
}

// WithFailMessage overrides the default error message for the following assertions.
func (a *Assertion[T]) WithFailMessage(message string, args ...any) *Assertion[T] {
	a.info.WithOverridingFailureMessage(message, args...)
	return a
}

// WithFailMessageSupplier overrides the default error message for the following assertions.
// The new error message is built if the assertion fails by consuming the given Supplier function.
func (a *Assertion[T]) WithFailMessageSupplier(supplier message.FailureMessageSupplier) *Assertion[T] {
	a.info.WithOverridingFailureMessageSupplier(supplier)
	return a
}

// WithRepresentation uses the given representation to describe values in error messages.
func (a *Assertion[T]) WithRepresentation(representation message.Representation) *Assertion[T] {
	a.info.UsingRepresentation(representation)
	return a
}

// InHexadecimal uses hexadecimal representation to describe values in error messages.
func (a *Assertion[T]) InHexadecimal() *Assertion[T] {
	a.info.UsingHexadecimalRepresentation()
	return a
}

// InBinary uses binary representation to describe values in error messages.
func (a *Assertion[T]) InBinary() *Assertion[T] {
	a.info.UsingBinaryRepresentation()
	return a
}

// FailWithMessage records an assertion error
func (a *Assertion[T]) FailWithMessage(message string) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	description := a.info.Description()
	overridingErrorMessage := a.info.OverridingFailureMessage()
	representation := a.info.Representation()
	if overridingErrorMessage != "" {
		a.t.Errorf(a.formatter.Format(description, overridingErrorMessage, representation, a.actual))
	} else {
		a.t.Errorf(a.formatter.Format(description, message, representation, a.actual))
	}
}

// Matches runs the provided matchers against the actual value.
func (a *Assertion[T]) Matches(matchers ...matcher.Matcher[T]) {
	for _, matcher := range matchers {
		if !matcher.Matches(a.actual) {
			expected := matcher.Describe(a.info.representation)
			a.FailWithMessage(expected)
		}
	}
}

// That creates a new assertion for the given value.
func That[T any](t TestingT, actual T) *Assertion[T] {
	return &Assertion[T]{
		t:         t,
		info:      NewWritableAssertionInfo(),
		formatter: message.NewCompactMessageFormatter(),
		actual:    actual,
	}
}
