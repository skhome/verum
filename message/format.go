package message

import "fmt"

// Formatter formats a message to be included in an assertion error.
type Formatter interface {
	// Format formats the message to be included in an assertion error.
	Format(description Description, message string, representation Representation, actual any) string
}

// CompactMessageFormatter produces a compact assertion error message.
type CompactMessageFormatter struct {
	descriptionFormatter DescriptionFormatter
}

// NewCompactMessageFormatter creates and returns a new CompactMessageFormatter
func NewCompactMessageFormatter() CompactMessageFormatter {
	return CompactMessageFormatter{DefaultDescriptionFormatter}
}

// Format formats the message to be included in an assertion error.
func (f CompactMessageFormatter) Format(description Description, message string, representation Representation, actual any) string {
	des := f.descriptionFormatter(description)
	act := f.asText(representation, actual)
	return fmt.Sprintf("%sexpected: %s, but got: %s", des, message, act)
}

func (f CompactMessageFormatter) asText(representation Representation, value any) string {
	return representation(value)
}
