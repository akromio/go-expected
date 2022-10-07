package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToBeEmpty_IfEmptyStringThenWrapperMustBeReturned(t *testing.T) {
	// (1) act
	w := NewValueWrapper(new(test), "")
	out := w.ToBeEmpty()

	// (2) assessment
	assert := assert.New(t)
	assert.Equal(out, w)
}

func Test_ToBeEmpty_IfEmptySliceThenWrapperMustBeReturned(t *testing.T) {
	// (1) act
	w := NewValueWrapper(new(test), []string{})
	out := w.ToBeEmpty()

	// (2) assessment
	assert := assert.New(t)
	assert.Equal(out, w)
}

func Test_ToBeEmpty_IfFalseThenErrorMustBeCalled(t *testing.T) {
	// (1) pre
	tDouble := new(test)
	tDouble.On("Errorf", "'%v' should be empty.", []any{"ciao!"})

	// (2) act
	w := NewValueWrapper(tDouble, "ciao!")
	w.ToBeEmpty()

	// (3) assessment
	tDouble.AssertExpectations(t)
}

func Test_NotToBeEmpty_IfTrueThenWrapperMustBeReturned(t *testing.T) {
	// (1) act
	w := NewValueWrapper(new(test), "ciao!")
	out := w.NotToBeEmpty()

	// (2) assessment
	assert := assert.New(t)
	assert.Equal(out, w)
}

func Test_NotToBeEmpty_IfFalseThenErrofMustBeCalled(t *testing.T) {
	// (1) pre
	tDouble := new(test)
	tDouble.On("Errorf", "'%v' should not be empty.", []any{""})

	// (2) act
	w := NewValueWrapper(tDouble, "")
	w.NotToBeEmpty()

	// (3) assessment
	tDouble.AssertExpectations(t)
}
