package matcher_test

import (
	"fmt"
	"testing"

	"github.com/skhome/verum/assert"
	"github.com/skhome/verum/matcher"
)

func TestIsEqualToString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		actual   string
		expected string
		ok       bool
	}{
		{name: "ok", actual: "Frodo", expected: "Frodo", ok: true},
		{name: "not equal", actual: "Frodo", expected: "Sam", ok: false},
	}
	messageFormat := "expected: value equal to <%s>, but got: <%s>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.expected, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsEqualTo(tt.expected))
		})
	}
}

func TestIsEqualToBool(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		actual   bool
		expected bool
		ok       bool
	}{
		{name: "ok", actual: true, expected: true, ok: true},
		{name: "not equal", actual: true, expected: false, ok: false},
	}
	messageFormat := "expected: value equal to <%t>, but got: <%t>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.expected, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsEqualTo(tt.expected))
		})
	}
}

func TestIsLessThan(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		actual   int
		expected int
		ok       bool
	}{
		{name: "less", actual: 41, expected: 42, ok: true},
		{name: "not", actual: 42, expected: 42, ok: false},
	}
	messageFormat := "expected: value less than <%d>, but got: <%d>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.expected, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsLessThan(tt.expected))
		})
	}
}

func TestIsLessThanOrEqualTo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		actual   int
		expected int
		ok       bool
	}{
		{name: "less", actual: 41, expected: 42, ok: true},
		{name: "equal", actual: 42, expected: 42, ok: true},
		{name: "not", actual: 43, expected: 42, ok: false},
	}
	messageFormat := "expected: value less than or equal to <%d>, but got: <%d>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.expected, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsLessThanOrEqualTo(tt.expected))
		})
	}
}

func TestIsGreaterThan(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		actual   int
		expected int
		ok       bool
	}{
		{name: "greater", actual: 43, expected: 42, ok: true},
		{name: "not", actual: 42, expected: 42, ok: false},
	}
	messageFormat := "expected: value greater than <%d>, but got: <%d>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.expected, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsGreaterThan(tt.expected))
		})
	}
}

func TestIsGreaterThanOrEqualTo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		actual   int
		expected int
		ok       bool
	}{
		{name: "greater", actual: 43, expected: 42, ok: true},
		{name: "equal", actual: 42, expected: 42, ok: true},
		{name: "not", actual: 41, expected: 42, ok: false},
	}
	messageFormat := "expected: value greater than or equal to <%d>, but got: <%d>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.expected, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsGreaterThanOrEqualTo(tt.expected))
		})
	}
}

func TestIsBetween(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		actual         int
		startInclusive int
		endInclusive   int
		ok             bool
	}{
		{name: "start", actual: 42, startInclusive: 42, endInclusive: 43, ok: true},
		{name: "end", actual: 42, startInclusive: 41, endInclusive: 42, ok: true},
		{name: "not", actual: 42, startInclusive: 43, endInclusive: 44, ok: false},
	}
	messageFormat := "expected: value between <%d> and <%d>, but got: <%d>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.startInclusive, tt.endInclusive, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsBetween(tt.startInclusive, tt.endInclusive))
		})
	}
}
