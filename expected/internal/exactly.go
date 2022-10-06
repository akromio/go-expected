package internal

import (
	"reflect"
)

// Exactly checks whether the wrapped value has the same type as another value
// and their values are equal.
func (w *ValueWrapper) Exactly(v any) *ValueWrapper {
	wType := reflect.TypeOf(w.value)
	vType := reflect.TypeOf(v)

	if wType != vType || w.value != v {
		w.t.Errorf("%v should be exactly %v.", w.value, v)
	}

	return w
}
