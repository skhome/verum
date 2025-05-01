package matcher_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/skhome/verum/assert"
	"github.com/skhome/verum/matcher"
)

func TestIsEmptyString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		ok     bool
	}{
		{actual: "", ok: true},
		{actual: "Frodo", ok: false},
	}
	messageFormat := "expected: empty string, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsEmptyString())
		})
	}
}

func TestIsBlankString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		ok     bool
	}{
		{actual: "", ok: true},
		{actual: " ", ok: true},
		{actual: "\t", ok: true},
		{actual: "Frodo", ok: false},
	}
	messageFormat := "expected: blank string, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsBlank())
		})
	}
}

func TestHasLineCount(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		count  int
		ok     bool
	}{
		{actual: "Frodo", count: 1, ok: true},
		{actual: "Frodo\nSam", count: 2, ok: true},
		{actual: "Frodo\nSam", count: 3, ok: false},
	}
	messageFormat := "expected: string with %d lines, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.count, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.HasLineCount(tt.count))
		})
	}
}

func TestIsEqualToIgnoringCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual   string
		expected string
		ok       bool
	}{
		{actual: "Frodo", expected: "Frodo", ok: true},
		{actual: "frodo", expected: "Frodo", ok: true},
		{actual: "Sam", expected: "Frodo", ok: false},
	}
	messageFormat := "expected: string equal to <%v> ignoring case, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.expected, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsEqualToIgnoringCase(tt.expected))
		})
	}
}

func TestContainsWhitespace(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		ok     bool
	}{
		{actual: "", ok: false},
		{actual: " ", ok: true},
		{actual: "\t", ok: true},
		{actual: "Frodo", ok: false},
	}
	messageFormat := "expected: string containing whitespace characters, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.ContainsWhitespace())
		})
	}
}

func TestContainsDigit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		ok     bool
	}{
		{actual: "3", ok: true},
		{actual: "3 rings of power", ok: true},
		{actual: "Nazgul", ok: false},
	}
	messageFormat := "expected: string containing a digit, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.ContainsDigit())
		})
	}
}

func TestContainsOnlyDigits(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		ok     bool
	}{
		{actual: "1", ok: true},
		{actual: "", ok: false},
		{actual: "10a", ok: false},
	}
	messageFormat := "expected: string containing only digits, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.ContainsOnlyDigits())
		})
	}
}

func TestContainsSubstring(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual  string
		substrs []string
		ok      bool
	}{
		{actual: "Frodo", substrs: []string{"do"}, ok: true},
		{actual: "Frodo", substrs: []string{"am"}, ok: false},
	}
	messageFormat := "expected: string to contain <%v>, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.substrs, tt.actual)
		run(t, strings.Join(tt.substrs, ""), tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.ContainsSubstring(tt.substrs...))
		})
	}
}

func TestContainsOnlyOnce(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		substr string
		ok     bool
	}{
		{actual: "Frodo", substr: "do", ok: true},
		{actual: "Frodo", substr: "o", ok: false},
		{actual: "Frodo", substr: "y", ok: false},
	}
	messageFormat := "expected: string to contain <%v> only once, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.substr, tt.actual)
		run(t, tt.substr, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.ContainsSubstringOnlyOnce(tt.substr))
		})
	}
}

func TestContainsSubstringIgnoringCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual  string
		substrs []string
		ok      bool
	}{
		{actual: "Frodo", substrs: []string{"do", "Fro"}, ok: true},
		{actual: "Frodo", substrs: []string{"DO", "fro"}, ok: true},
		{actual: "Frodo", substrs: []string{"am"}, ok: false},
	}
	messageFormat := "expected: string to contain <%v> ignoring case, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.substrs, tt.actual)
		run(t, strings.Join(tt.substrs, ""), tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.ContainsSubstringIgnoringCase(tt.substrs...))
		})
	}
}

func TestContainsSubstringIgnoringWhitespace(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual  string
		substrs []string
		ok      bool
	}{
		{actual: "Gandalf the grey", substrs: []string{"alf"}, ok: true},
		{actual: "Gandalf the grey", substrs: []string{"dalf", "grey"}, ok: true},
		{actual: "Gandalf the grey", substrs: []string{"thegrey"}, ok: true},
		{actual: "Gandalf the grey", substrs: []string{"the  grey"}, ok: true},
		{actual: "Gandalf the grey", substrs: []string{"t h e grey"}, ok: true},
		{actual: "Gandalf the grey", substrs: []string{"Alf"}, ok: false},
	}
	messageFormat := "expected: string to contain <%v> ignoring whitespace, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.substrs, tt.actual)
		run(t, strings.Join(tt.substrs, ""), tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.ContainsSubstringIgnoringWhitespace(tt.substrs...))
		})
	}
}

func TestContainsAnySubstring(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual  string
		substrs []string
		ok      bool
	}{
		{actual: "Frodo", substrs: []string{"do", "am"}, ok: true},
		{actual: "Frodo", substrs: []string{"am"}, ok: false},
	}
	messageFormat := "expected: string to contain any of <%v>, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.substrs, tt.actual)
		run(t, strings.Join(tt.substrs, ""), tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.ContainsAnySubstring(tt.substrs...))
		})
	}
}

func TestStartsWith(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		substr string
		ok     bool
	}{
		{actual: "Frodo", substr: "Fro", ok: true},
		{actual: "Sam", substr: "Fro", ok: false},
	}
	messageFormat := "expected: string to start with <%v>, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.substr, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.StartsWith(tt.substr))
		})
	}
}

func TestStartsWithIgnoringCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		substr string
		ok     bool
	}{
		{actual: "Frodo", substr: "Fro", ok: true},
		{actual: "frodo", substr: "Fro", ok: true},
		{actual: "Sam", substr: "Fro", ok: false},
	}
	messageFormat := "expected: string to start with <%v> ignoring case, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.substr, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.StartsWithIgnoringCase(tt.substr))
		})
	}
}

func TestEndsWith(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		substr string
		ok     bool
	}{
		{actual: "Frodo", substr: "do", ok: true},
		{actual: "Sam", substr: "do", ok: false},
	}
	messageFormat := "expected: string to end with <%v>, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.substr, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.EndsWith(tt.substr))
		})
	}
}

func TestEndsWithIgnoringCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		actual string
		substr string
		ok     bool
	}{
		{actual: "Frodo", substr: "do", ok: true},
		{actual: "FRODO", substr: "do", ok: true},
		{actual: "Sam", substr: "do", ok: false},
	}
	messageFormat := "expected: string to end with <%v> ignoring case, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.substr, tt.actual)
		run(t, tt.actual, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.EndsWithIgnoringCase(tt.substr))
		})
	}
}
