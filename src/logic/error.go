package logic

import (
	"errors"
)

var (
	ParamError error = errors.New("Wrong Param")
	NoDataError error = errors.New("No Data")
)
