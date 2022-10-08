package internal

// ValueWrapper contains a value for performing assertions.
type ValueWrapper struct {
	t     TestingT
	value any // the value wrapped
}

// NewValueWrapper returns a new valueWrapper.
func NewValueWrapper(t TestingT, value any) *ValueWrapper {
	return &ValueWrapper{t, value}
}
