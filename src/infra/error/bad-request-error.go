package error

import (
	"encoding/json"
	"errors"
)

type BadRequestError struct {
	status    int
	errorCode string
	codeType  string
	details   interface{}
}

func NewBadRequestError(errorCode string, data interface{}) error {
	err := &BadRequestError{
		status:    400,
		errorCode: errorCode,
		codeType:  "BadRequestError",
		details:   data,
	}

	jErr, _ := json.Marshal(err)
	return errors.New(string(jErr))
}
