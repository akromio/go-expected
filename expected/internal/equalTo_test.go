package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EqualTo_IfTrueThenWrapperMustBeReturned(t *testing.T) {
	// (1) act
	w := NewValueWrapper(new(test), "ciao!")
	out := w.EqualTo("ciao!")

	// (2) assessment
	assert := assert.New(t)
	assert.Equal(out, w)
}

func Test_EqualTo_IfFalseThenErrorfMustBeCalled(t *testing.T) {
	// (1) pre
	tDouble := new(test)
	tDouble.On("Errorf", "%v should be equal to %v.", []any{"hello!", "ciao!"})

	// (2) act
	w := NewValueWrapper(tDouble, "hello!")
	w.EqualTo("ciao!")

	// (3) assessment
	tDouble.AssertExpectations(t)
}

func Test_NotEqualTo_IfTrueThenWrapperMustBeReturned(t *testing.T) {
	// (1) act
	w := NewValueWrapper(new(test), "ciao!")
	out := w.NotEqualTo("hello!")

	// (2) assessment
	assert := assert.New(t)
	assert.Equal(out, w)
}

func Test_NotEqualTo_IfFalseThenErrorfMustBeCalled(t *testing.T) {
	// (1) pre
	tDouble := new(test)
	tDouble.On("Errorf", "Values should not be equal to '%v'.", []any{"ciao!"})

	// (2) act
	w := NewValueWrapper(tDouble, "ciao!")
	w.NotEqualTo("ciao!")

	// (3) assessment
	tDouble.AssertExpectations(t)
}
