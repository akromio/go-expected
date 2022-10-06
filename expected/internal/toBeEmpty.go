package internal

import "reflect"

// isEmpty checks whether a value is empty.
func isEmpty(v any) bool {
	// (1) check
	empty := v == nil

	if !empty {
		value := reflect.ValueOf(v)

		switch value.Kind() {
		case reflect.Slice, reflect.Map, reflect.Chan:
			empty = value.Len() == 0
		case reflect.Ptr:
			if empty = value.IsNil(); !empty {
				empty = isEmpty(value.Elem().Interface())
			}
		default:
			zero := reflect.Zero(value.Type())
			empty = reflect.DeepEqual(v, zero.Interface())
		}
	}

	// (2) return result
	return empty
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
