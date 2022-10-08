package internal

// ToBeNil checks whether the wrapped value is nil.
func (w *ValueWrapper) ToBeNil() *ValueWrapper {
	if w.value != nil {
		w.t.Error("Value should be nil.")
	}

	return w
}

// NotToBeNil checks whether the wrapped value is not nil.
func (w *ValueWrapper) NotToBeNil() *ValueWrapper {
	if w.value == nil {
		w.t.Error("Value should not be nil.")
	}

	return w
}
