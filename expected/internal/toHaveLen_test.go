package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToHaveLen_IfTrueThenWrapperMustBeReturned(t *testing.T) {
	// (1) act
	w := NewValueWrapper(new(test), "ciao!")
	out := w.ToHaveLen(5)

	// (2) assessment
	assert := assert.New(t)
	assert.Equal(out, w)
}

func Test_ToHaveLen_IfFalseThenErrorfMustBeCalled(t *testing.T) {
	// (1) pre
	tDouble := new(test)
	tDouble.On("Errorf", "%v should have a length of %v.", []any{"ciao!", 10})

	// (2) act
	w := NewValueWrapper(tDouble, "ciao!")
	w.ToHaveLen(10)

	// (3) assessment
	tDouble.AssertExpectations(t)
}

func Test_NotToHaveLen_IfTrueThenWrapperMustBeReturned(t *testing.T) {
	// (1) act
	w := NewValueWrapper(new(test), "ciao!")
	out := w.NotToHaveLen(15)

	// (2) assessment
	assert := assert.New(t)
	assert.Equal(out, w)
}

func Test_NotToHaveLen_IfFalseThenErrorfMustBeCalled(t *testing.T) {
	// (1) pre
	tDouble := new(test)
	tDouble.On("Errorf", "%v should not have a length of %v.", []any{"ciao!", 5})

	// (2) act
	w := NewValueWrapper(tDouble, "ciao!")
	w.NotToHaveLen(5)

	// (3) assessment
	tDouble.AssertExpectations(t)
}
