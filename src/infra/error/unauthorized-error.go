package error

import (
	"encoding/json"
	"errors"
)

type UnauthorizedError struct {
	status    int
	errorCode string
	codeType  string
	details   interface{}
}

func NewUnauthorizedError(errorCode string, data interface{}) error {
	err := &UnauthorizedError{
		status:    400,
		errorCode: errorCode,
		codeType:  "UnauthorizedError",
		details:   data,
	}

	jErr, _ := json.Marshal(err)
	return errors.New(string(jErr))
}
