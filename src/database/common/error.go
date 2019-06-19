package common

import (
	"errors"
)

var (
	UndefinedError error = errors.New("Undefined Error")
	DbDataError    error = errors.New("db data Error")
	InserFailError error = errors.New("insert data fail")
	UpdateFailError error = errors.New("update data fail")
)
