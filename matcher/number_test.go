package matcher_test

import (
	"fmt"
	"testing"

	"github.com/skhome/verum/assert"
	"github.com/skhome/verum/matcher"
)

func TestIsZero(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		actual int
		ok     bool
	}{
		{name: "zero", actual: 0, ok: true},
		{name: "not", actual: 42, ok: false},
	}
	messageFormat := "expected: value to be zero, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsZero[int]())
		})
	}
}

func TestIsPositive(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		actual int
		ok     bool
	}{
		{name: "positive", actual: 1, ok: true},
		{name: "zero", actual: 0, ok: false},
	}
	messageFormat := "expected: value to be positive, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsPositive[int]())
		})
	}
}

func TestIsNegative(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		actual int
		ok     bool
	}{
		{name: "negative", actual: -1, ok: true},
		{name: "zero", actual: 0, ok: false},
	}
	messageFormat := "expected: value to be negative, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsNegative[int]())
		})
	}
}

func TestIsEven(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		actual int
		ok     bool
	}{
		{name: "even", actual: 2, ok: true},
		{name: "odd", actual: 1, ok: false},
	}
	messageFormat := "expected: value to be even, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsEven[int]())
		})
	}
}

func TestIsOdd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		actual int
		ok     bool
	}{
		{name: "odd", actual: 1, ok: true},
		{name: "even", actual: 2, ok: false},
	}
	messageFormat := "expected: value to be odd, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsOdd[int]())
		})
	}
}
