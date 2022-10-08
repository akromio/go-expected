package internal

import "github.com/stretchr/testify/mock"

type test struct {
	mock.Mock
}

func (t *test) Error(args ...any) {
	t.Called(args)
}

func (t *test) Errorf(format string, args ...any) {
	t.Called(format, args)
}
