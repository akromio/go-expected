package internal

import "reflect"

// Length returns the size of a given value.
func length(v any) int {
	return reflect.ValueOf(v).Len()
}

// ToHaveLen checks whether the length of the wrapped value
// is a given one.
func (w *ValueWrapper) ToHaveLen(size int) *ValueWrapper {
	if length(w.value) != size {
		w.t.Errorf("%v should have a length of %v.", w.value, size)
	}

	return w
}

// NotToHaveLen checks whether the length of the wrapped value
// is not a given one.
func (w *ValueWrapper) NotToHaveLen(size int) *ValueWrapper {
	if length(w.value) == size {
		w.t.Errorf("%v should not have a length of %v.", w.value, size)
	}

	return w
}
