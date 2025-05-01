package message

import "fmt"

// Description is a description for an assertion.
type Description string

// DescriptionFormatter formats a description to be included in assertion errors.
type DescriptionFormatter func(description Description) string

// DefaultDescriptionFormatter formats a description by surrounding it in square brackets.
func DefaultDescriptionFormatter(description Description) string {
	if description != "" {
		return fmt.Sprintf("[%s] ", description)
	}
	return string(description)
}
