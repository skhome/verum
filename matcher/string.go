package matcher

import (
	"fmt"
	"strings"

	"github.com/skhome/verum/check"
	"github.com/skhome/verum/message"
)

// IsEmptyString verifies that the actual string is empty.
//
//	// assertions will pass
//	assert.That(t, "").Matches(matcher.IsEmptyString())
//
//	// assertions will fail
//	assert.That(t, "Frodo").Matches(matcher.IsEmptyString())
func IsEmptyString() Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return len(actual) == 0
		},
		describeFn: func(message.Representation) string {
			return "empty string"
		},
	}
}

// IsBlank verifies that the actual string is empty or contains only whitespace characters.
//
//	// assertions will pass
//	assert.That(t, "").Matches(matcher.IsBlank())
//	assert.That(t, " ").Matches(matcher.IsBlank())
//
//	// assertions will fail
//	assert.That(t, "Frodo").Matches(matcher.IsBlank())
func IsBlank() Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return strings.TrimSpace(actual) == ""
		},
		describeFn: func(message.Representation) string {
			return "blank string"
		},
	}
}

// HasLineCount verifies that the actual string has the specified number of lines.
//
//	// assertions will pass
//	assert.That(t, "Frodo\nSam").Matches(matcher.HasLineCount(2))
//
//	// assertions will fail
//	assert.That(t, "Frodo").Matches(matcher.HasLineCount(2))
func HasLineCount(count int) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return check.StringLineCount(actual) == count
		},
		describeFn: func(message.Representation) string {
			return fmt.Sprintf("string with %d lines", count)
		},
	}
}

// IsEqualToIgnoringCase verifies that the actual string equals the given one ignoring case.
//
//	// assertions will pass
//	assert.That(t, "Frodo").Matches(matcher.IsEqualTo("Frodo"))
//	assert.That(t, "frodo").Matches(matcher.IsEqualTo("Frodo"))
//
//	// assertions will fail
//	assert.That(t, "Sam").Matches(matcher.IsEqualTo("Frodo"))
func IsEqualToIgnoringCase(expected string) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return strings.EqualFold(actual, expected)
		},
		describeFn: func(representation message.Representation) string {
			return fmt.Sprintf("string equal to %s ignoring case", representation(expected))
		},
	}
}

// ContainsWhitespace verifies that the actual string contains at least one whitespace character.
//
//	// assertions will pass
//	assert.That(t, "Frodo Baggins").Matches(matcher.ContainsWhitespace())
//
//	// assertions will fail
//	assert.That(t, "Frodo").Matches(matcher.ContainsWhitespace())
func ContainsWhitespace() Matcher[string] {
	return &matcher[string]{
		matcherFn: check.StringContainsWhitespace,
		describeFn: func(message.Representation) string {
			return "string containing whitespace characters"
		},
	}
}

// ContainsDigit verifies that the actual string contains at least one digit.
//
//	// assertion will pass
//	assert.That(t, "3 elven rings").Matches(matcher.ContainsDigit())
//
//	// assertion will fail
//	assert.That(t, "Nazgul").Matches(matcher.ContainsDigit())
func ContainsDigit() Matcher[string] {
	return &matcher[string]{
		matcherFn: check.StringContainsDigit,
		describeFn: func(message.Representation) string {
			return "string containing a digit"
		},
	}
}

// ContainsOnlyDigits verifies that the actual string contains only digits.
//
//	// assertion will pass
//	assert.That(t, "1234567890").Matches(matcher.ContainsOnlyDigits())
//
//	// assertion will fail
//	assert.That(t, "123abc").Matches(matcher.ContainsOnlyDigits())
func ContainsOnlyDigits() Matcher[string] {
	return &matcher[string]{
		matcherFn: check.StringContainsOnlyDigits,
		describeFn: func(message.Representation) string {
			return "string containing only digits"
		},
	}
}

// ContainsSubstring verifies that the actual string contains the specified substring.
//
//	// assertion will pass
//	assret.That(t, "Frodo").Matches(matcher.ContainsSubstring("do"))
//
//	// assertion will fail
//	assert.That(t, "Frodo").Matches(matcher.ContainsSubstring("am"))
func ContainsSubstring(substrs ...string) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return check.StringContains(actual, substrs)
		},
		describeFn: func(representation message.Representation) string {
			return "string to contain " + representation(substrs)
		},
	}
}

// ContainsSubstringOnlyOnce verifies that the actual string contains the specified substring only once.
//
//	// assertion will pass
//	assert.That(t, "Frodo").Matches(matcher.ContainsSubstringOnlyOnce("do"))
//
//	// assertion will fail
//	assert.That(t, "Frodo").Matches(matcher.ContainsSubstringOnlyOnce("o"))
func ContainsSubstringOnlyOnce(substr string) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return strings.Count(actual, substr) == 1
		},
		describeFn: func(representation message.Representation) string {
			return fmt.Sprintf("string to contain %s only once", representation(substr))
		},
	}
}

// ContainsSubstringIgnoringCase verifies that the actual string contains the specified substring, ignoring case.
//
//	// assertion will pass
//	assert.That(t, "Frodo").Matches(matcher.ContainsSubstringIgnoringCase("fro"))
//
//	// assertion will fail
//	assert.That(t, "Frodo").Matches(matcher.ContainsSubstringIgnoringCase("am"))
func ContainsSubstringIgnoringCase(substrs ...string) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return check.StringContainsIgnoringCase(actual, substrs)
		},
		describeFn: func(representation message.Representation) string {
			return fmt.Sprintf("string to contain %s ignoring case", representation(substrs))
		},
	}
}

// ContainsSubstringIgnoringWhitespace verifies that the actual string contains the specified substrings, ignoring whitespace.
//
//	// assertion will pass
//	assert.That(t, "Gandalf the grey").Matches(matcher.ContainsSubstringIgnoringWhitespace("thegrey"))
//
//	// assertion will fail
//	assert.That(t, "Gandalf the grey").Matches(matcher.ContainsSubstringIgnoringWhitespace("gan"))
func ContainsSubstringIgnoringWhitespace(substrs ...string) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return check.StringContainsIgnoringWhitespace(actual, substrs)
		},
		describeFn: func(representation message.Representation) string {
			return fmt.Sprintf("string to contain %s ignoring whitespace", representation(substrs))
		},
	}
}

// ContainsAnySubstring verifies that the actual string contains any of the specified substrings.
//
//	// assertion will pass
//	assert.That(t, "Frodo").Matches(matcher.ContainsAnySubstring("do", "am"))
//
//	// assertion will fail
//	assert.That(t, "Frodo").Matches(matcher.ContainsAnySubstring("am"))
func ContainsAnySubstring(substrs ...string) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return check.StringContainsAny(actual, substrs)
		},
		describeFn: func(representation message.Representation) string {
			return "string to contain any of " + representation(substrs)
		},
	}
}

// StartsWith verifies that the actual string starts with the specified substring.
//
//	// assertion will pass
//	assert.That(t, "Frodo").Matches(matcher.StartsWith("Fro"))
//
//	// assertion will fail
//	assert.That(t, "Frodo").Matches(matcher.StartsWith("fro"))
func StartsWith(prefix string) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return check.StringStartsWith(actual, prefix)
		},
		describeFn: func(representation message.Representation) string {
			return "string to start with " + representation(prefix)
		},
	}
}

// StartsWithIgnoringCase verifies that the actual string starts with the specified substring, ignoring case.
//
//	// assertion will pass
//	assert.That(t, "Frodo").Matches(matcher.StartsWithIgnoringCase("fro"))
//
//	// assertion will fail
//	assert.That(t, "Sam").Matches(matcher.StartsWithIgnoringCase("fro"))
func StartsWithIgnoringCase(prefix string) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return check.StringStartsWithIgnoringCase(actual, prefix)
		},
		describeFn: func(representation message.Representation) string {
			return fmt.Sprintf("string to start with %s ignoring case", representation(prefix))
		},
	}
}

// EndsWith verifies that the actual string ends with the specified substring.
//
//	// assertion will pass
//	assert.That(t, "Frodo").Matches(matcher.EndsWith("do"))
//
//	// assertion will fail
//	assert.That(t, "Sam").Matches(matcher.EndsWith("do"))
func EndsWith(suffix string) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return check.StringEndsWith(actual, suffix)
		},
		describeFn: func(representation message.Representation) string {
			return "string to end with " + representation(suffix)
		},
	}
}

// EndsWithIgnoringCase verifies that the actual string ends with the specified substring, ignoring case.
//
//	// assertion will pass
//	assert.That(t, "FRODO").Matches(matcher.EndsWithIgnoringCase("do"))
//
//	// assertion will fail
//	assert.That(t, "Sam").Matches(matcher.EndsWithIgnoringCase("do"))
func EndsWithIgnoringCase(suffix string) Matcher[string] {
	return &matcher[string]{
		matcherFn: func(actual string) bool {
			return check.StringEndsWithIgnoringCase(actual, suffix)
		},
		describeFn: func(representation message.Representation) string {
			return fmt.Sprintf("string to end with %s ignoring case", representation(suffix))
		},
	}
}
