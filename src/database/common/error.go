package common

import (
	"errors"
)

var (
	UndefinedError error = errors.New("Undefined Error")
	DbDataError    error = errors.New("db data Error")
)
