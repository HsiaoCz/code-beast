package api

import "net/http"

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Error implements the Error interface
func (e Error) Error() string {
	return e.Msg
}

func NewError(code int, msg string) Error {
	return Error{
		Code: code,
		Msg:  msg,
	}
}

func ErrUnAuthorized() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Msg:  "unauthorized request",
	}
}

func ErrInvalidID() Error {
	return Error{
		Code: http.StatusBadRequest,
		Msg:  "invalid id given",
	}
}
