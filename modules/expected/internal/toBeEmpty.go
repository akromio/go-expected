package internal

import "reflect"

// isEmpty checks whether a value is empty.
func isEmpty(v any) bool {
	return reflect.ValueOf(v).Len() == 0
}

// ToBeEmpty checks whether the wrapper value is empty.
func (w *ValueWrapper) ToBeEmpty() *ValueWrapper {
	if !isEmpty(w.value) {
		w.t.Errorf("'%v' should be empty.", w.value)
	}

	return w
}

// NotToBeEmpty checks whether the wrapper value is not empty.
func (w *ValueWrapper) NotToBeEmpty() *ValueWrapper {
	if isEmpty(w.value) {
		w.t.Errorf("'%v' should not be empty.", w.value)
	}

	return w
}
