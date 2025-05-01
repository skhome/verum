package matcher_test

import (
	"testing"

	"github.com/skhome/verum/assert"
	"github.com/skhome/verum/matcher"
)

func TestIsTrue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		actual bool
		ok     bool
	}{
		{name: "true", actual: true, ok: true},
		{name: "false", actual: false, ok: false},
	}
	message := "expected: value to be true"
	for _, tt := range tests {
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsTrue())
		})
	}
}

func TestIsFalse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		actual bool
		ok     bool
	}{
		{name: "false", actual: false, ok: true},
		{name: "true", actual: true, ok: false},
	}
	message := "expected: value to be false"
	for _, tt := range tests {
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsFalse())
		})
	}
}
