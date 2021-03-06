package error

import (
	"encoding/json"
	"errors"
)

type InternalServerError struct {
	status    int
	errorCode string
	codeType  string
	details   interface{}
}

func NewInternalServerError(errorCode string, data interface{}) error {
	err := &InternalServerError{
		status:    500,
		errorCode: errorCode,
		codeType:  "InternalServerError",
		details:   data,
	}

	jErr, _ := json.Marshal(err)
	return errors.New(string(jErr))
}
