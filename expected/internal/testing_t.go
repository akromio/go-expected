package internal

// TestingT is the interface for testing.T.
type TestingT interface {
	Error(args ...any)
	Errorf(format string, args ...any)
}
