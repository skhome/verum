package message

// FailureMessageSupplier lazyly supplies a failure message for an assertion info.
type FailureMessageSupplier func() string
