package assert

import (
	"fmt"

	"github.com/skhome/verum/message"
)

// AssertionInfo provides information about an assertion.
type AssertionInfo interface {
	// OverridingFailureMessage returns the message that, if specified, will
	// replace the default message of an assertion failure.
	OverridingFailureMessage() string

	// Description returns the description of an assertion.
	Description() message.Description

	// Representation returns the Representation of the actual and expected values.
	Representation() message.Representation
}

// WritableAssertionInfo implements a modifiable AssertionInfo.
type WritableAssertionInfo struct {
	overridingFailureMessage         string
	overridingFailureMessageSupplier message.FailureMessageSupplier
	description                      message.Description
	representation                   message.Representation
}

// NewWritableAssertionInfo creates a new WritableAssertionInfo with a default representation.
func NewWritableAssertionInfo() *WritableAssertionInfo {
	return &WritableAssertionInfo{
		representation: message.DefaultRepresentation,
	}
}

// OverridingFailureMessage returns the message that, if specified, will
// replace the default message of an assertion failure.
func (i *WritableAssertionInfo) OverridingFailureMessage() string {
	if i.overridingFailureMessageSupplier != nil {
		return i.overridingFailureMessageSupplier()
	}
	return i.overridingFailureMessage
}

// WithOverridingFailureMessage sets the failure message that replaces the default failure message.
func (i *WritableAssertionInfo) WithOverridingFailureMessage(message string, args ...any) {
	var formatted []any
	for _, arg := range args {
		formatted = append(formatted, i.representation(arg))
	}
	if len(args) > 0 {
		i.overridingFailureMessage = fmt.Sprintf(message, formatted...)
	} else {
		i.overridingFailureMessage = message
	}
}

// WithOverridingFailureMessageSupplier sets a lazy supplier for the failure message that replaces the default failure message.
func (i *WritableAssertionInfo) WithOverridingFailureMessageSupplier(supplier message.FailureMessageSupplier) {
	i.overridingFailureMessageSupplier = supplier
}

// Description returns the description of an assertion.
func (i *WritableAssertionInfo) Description() message.Description {
	return i.description
}

// HasDescription returns if the assertion has a description.
func (i *WritableAssertionInfo) HasDescription() bool {
	return i.description != ""
}

// WithDescription sets a description for an assertion.
func (i *WritableAssertionInfo) WithDescription(description message.Description) {
	i.description = description
}

// Representation returns the Representation of actual and expected values.
func (i *WritableAssertionInfo) Representation() message.Representation {
	return i.representation
}

// UsingHexadecimalRepresentation uses a hexadecimal representation for actual and expected values.
func (i *WritableAssertionInfo) UsingHexadecimalRepresentation() {
	i.representation = message.HexadecimalRepresentation
}

// UsingBinaryRepresentation uses binary representation for actual and expected values.
func (i *WritableAssertionInfo) UsingBinaryRepresentation() {
	i.representation = message.BinaryRepresentation
}

// UsingRepresentation uses the given representation for actual and expected values.
func (i *WritableAssertionInfo) UsingRepresentation(representation message.Representation) {
	i.representation = representation
}
