package errors

import (
	"errors"
)

//  https://github.com/olezhek28/platform_common/blob/main/pkg/sys/error.go

type customError struct {
	msg  string
	code Code
}

func NewError(msg string, code Code) *customError {
	return &customError{msg, code}
}

func (r *customError) Error() string {
	return r.msg
}

func (r *customError) Code() Code {
	return r.code
}

func IsCustomError(err error) bool {
	var ce *customError
	return errors.As(err, &ce)
}

func GetCommonError(err error) *customError {
	var ce *customError
	if !errors.As(err, &ce) {
		return nil
	}

	return ce
}
