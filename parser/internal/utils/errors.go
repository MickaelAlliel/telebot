package utils

import "github.com/pkg/errors"

type ErrorWithStackTrace interface {
	StackTrace() errors.StackTrace
}
