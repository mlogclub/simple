package simple

import (
	"strconv"
)

var (
	ErrorNotLogin = NewError(1, "请先登录")
)

func NewError(code int, text string) *codeError {
	return &codeError{code, text, nil}
}

func NewErrorMsg(text string) *codeError {
	return &codeError{0, text, nil}
}

func NewErrorData(code int, text string, data interface{}) *codeError {
	return &codeError{code, text, data}
}

func FromError(err error) *codeError {
	if err == nil {
		return nil
	}
	return &codeError{0, err.Error(), nil}
}

type codeError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *codeError) Error() string {
	return strconv.Itoa(e.Code) + ": " + e.Message
}
