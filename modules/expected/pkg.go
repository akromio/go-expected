package expected

import (
	"github.com/akromio/go-expected/expected/internal"
)

type ExpectedFunc = func(value any) *internal.ValueWrapper

// New returns a function for creating value wrappers
// during a test.
func New(t internal.TestingT) ExpectedFunc {
	return func(value any) *internal.ValueWrapper {
		return internal.NewValueWrapper(t, value)
	}
}
