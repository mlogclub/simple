package simple

import (
	"strconv"
)

var (
	ErrorNotLogin = NewError(1, "请先登录")
)

// New returns an error that formats as the given text.
func NewError(code int, text string) *CodeError {
	return &CodeError{code, text}
}

func NewErrorMsg(text string) *CodeError {
	return &CodeError{0, text}
}

func NewError2(err error) *CodeError {
	if err == nil {
		return nil
	}
	return &CodeError{0, err.Error()}
}

// errorString is a trivial implementation of error.
type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *CodeError) Error() string {
	return strconv.Itoa(e.Code) + ": " + e.Message
}
