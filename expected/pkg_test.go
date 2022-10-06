package expected

import (
	"testing"
)

func Test_EqualTo_IfTrueThenWrapperMustBeReturned(t *testing.T) {
	expected := New(t)
	expected("ciao!").EqualTo("ciao!")
}
