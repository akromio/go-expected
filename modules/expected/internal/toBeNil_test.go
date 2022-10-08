package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToBeNil_IfNilThenWrapperMustBeReturned(t *testing.T) {
	// (1) act
	w := NewValueWrapper(new(test), nil)
	out := w.ToBeNil()

	// (2) assessment
	assert := assert.New(t)
	assert.Equal(out, w)
}

func Test_ToBeNil_IfNotNilThenErrorMustBeCalled(t *testing.T) {
	// (1) pre
	tDouble := new(test)
	tDouble.On("Error", []any{"Value should be nil."})

	// (2) act
	w := NewValueWrapper(tDouble, "ciao!")
	w.ToBeNil()

	// (3) assessment
	tDouble.AssertExpectations(t)
}

func Test_NotToBeNil_IfNilThenErrorMustBeCalled(t *testing.T) {
	// (1) pre
	tDouble := new(test)
	tDouble.On("Error", []any{"Value should not be nil."})

	// (2) act
	w := NewValueWrapper(tDouble, nil)
	w.NotToBeNil()

	// (3) assessment
	tDouble.AssertExpectations(t)
}

func Test_NotToBeNil_IfNotNil_ThenWrapperMustBeReturned(t *testing.T) {
	// (1) act
	w := NewValueWrapper(new(test), "ciao!")
	out := w.NotToBeNil()

	// (2) assessment
	assert := assert.New(t)
	assert.Equal(out, w)
}
