package matcher_test

import (
	"fmt"
	"strings"
	"testing"
)

// function that runs a single test
type fixtureT struct {
	message string
}

func (f *fixtureT) Errorf(format string, args ...any) {
	f.message = fmt.Sprintf(format, args...)
}

func (f *fixtureT) Helper() {}

func run(t *testing.T, name string, ok bool, message string, exec func(t *fixtureT)) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		t.Parallel()
		fixture := new(fixtureT)
		exec(fixture)
		if ok {
			if len(fixture.message) > 0 {
				t.Errorf("expected not to fail, but got %#v", fixture.message)
			}
		} else {
			if !strings.Contains(fixture.message, message) {
				t.Errorf("expected to fail with error message %q, but got %#v", message, fixture.message)
			}
		}
	})
}
