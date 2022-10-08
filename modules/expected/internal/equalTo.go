package internal

import "reflect"

// EqualTo checks whether the wrapped value is equal to another.
func (w *ValueWrapper) EqualTo(v any) *ValueWrapper {
	if !reflect.DeepEqual(w.value, v) {
		w.t.Errorf("%v should be equal to %v.", w.value, v)
	}

	return w
}

// NotEqualTo checks whether the wrapped value is not equal to another.
func (w *ValueWrapper) NotEqualTo(v any) *ValueWrapper {
	if reflect.DeepEqual(w.value, v) {
		w.t.Errorf("Values should not be equal to '%v'.", v)
	}

	return w
}
