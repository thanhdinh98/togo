package error

import (
	"encoding/json"
	"errors"
)

type ForbiddenError struct {
	status    int
	errorCode string
	codeType  string
	details   interface{}
}

func NewForbiddenError(errorCode string, data interface{}) error {
	err := &ForbiddenError{
		status:    503,
		errorCode: errorCode,
		codeType:  "ForbiddenError",
		details:   data,
	}

	jErr, _ := json.Marshal(err)
	return errors.New(string(jErr))
}
