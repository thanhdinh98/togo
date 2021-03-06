package error

import (
	"encoding/json"
	"errors"
)

type NotFoundError struct {
	status    int
	errorCode string
	codeType  string
	details   interface{}
}

func NewNotFoundError(errorCode string, data interface{}) error {
	err := &NotFoundError{
		status:    404,
		errorCode: errorCode,
		codeType:  "NotFoundError",
		details:   data,
	}

	jErr, _ := json.Marshal(err)
	return errors.New(string(jErr))
}
