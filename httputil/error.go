package httputil

import (
	"errors"
	"fmt"
	"net/http"
)

type Error interface {
	error
	HTTPStatus() int
}

func StatusOf(err error) int {
	if err, ok := err.(Error); ok {
		return err.HTTPStatus()
	}
	return http.StatusInternalServerError
}

type statusError struct {
	status  int
	wrapped error
}

func (err *statusError) HTTPStatus() int {
	return err.status
}

func (err *statusError) Error() string {
	return err.wrapped.Error()
}

func (err *statusError) Unwrap() error {
	return err.wrapped
}

func WithStatus(status int, err error) Error {
	return &statusError{
		status:  status,
		wrapped: err,
	}
}

func NewError(status int, message string) Error {
	return WithStatus(status, errors.New(message))
}

func Errorf(status int, format string, v ...interface{}) Error {
	return WithStatus(status, fmt.Errorf(format, v...))
}
