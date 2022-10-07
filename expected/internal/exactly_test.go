package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Exactly_IfTrueThenWrapperMustBeReturned(t *testing.T) {
	// (1) act
	w := NewValueWrapper(new(test), "ciao!")
	out := w.Exactly("ci" + "ao!")

	// (2) assessment
	assert := assert.New(t)
	assert.Equal(out, w)
}

func Test_Exactly_IfFalseThenErrorfMustBeCalled(t *testing.T) {
	// (1) pre
	tDouble := new(test)
	tDouble.On("Errorf", "%v should be exactly %v.", []any{"ciao!", "hello!"})

	// (2) act
	w := NewValueWrapper(tDouble, "ciao!")
	w.Exactly("hello!")

	// (3) assessment
	tDouble.AssertExpectations(t)
}
