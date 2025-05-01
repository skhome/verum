package matcher_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/skhome/verum/assert"
	"github.com/skhome/verum/matcher"
)

func TestIsNilError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		actual error
		ok     bool
	}{
		{name: "nil", actual: nil, ok: true},
		{name: "not nil", actual: errors.New("abc"), ok: false},
	}
	messageFormat := "expected: value to be nil, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsNilError())
		})
	}
}

func TestIsNilMap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		actual map[string]int
		ok     bool
	}{
		{name: "nil", actual: nil, ok: true},
		{name: "not nil", actual: map[string]int{}, ok: false},
		{name: "has elements", actual: map[string]int{"Nazgul": 9}, ok: false},
	}
	messageFormat := "expected: value to be nil, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsNilMap[string, int]())
		})
	}
}

func TestIsNilPtr(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		actual *bytes.Buffer
		ok     bool
	}{
		{name: "nil", actual: nil, ok: true},
		{name: "not nil", actual: &bytes.Buffer{}, ok: false},
	}
	messageFormat := "expected: value to be nil, but got: <%v>"
	for _, tt := range tests {
		message := fmt.Sprintf(messageFormat, tt.actual)
		run(t, tt.name, tt.ok, message, func(t *fixtureT) {
			assert.That(t, tt.actual).Matches(matcher.IsNil[bytes.Buffer]())
		})
	}
}
