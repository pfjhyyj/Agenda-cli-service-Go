package utils

import (
	"fmt"
	"net/http"

	errors "github.com/go-errors/errors"
)

// Panic panic with error stack
func Panic(err error) {
	panic(errors.Wrap(err.Error()+"\n"+string(errors.Wrap(err, 1).Stack()), 1))
}

// Panicf formats according to a format specifier and panic the err
func Panicf(format string, i ...interface{}) {
	err := fmt.Errorf(format, i...)
	panic(errors.Wrap(err.Error()+"\n"+string(errors.Wrap(err, 1).Stack()), 1))
}

// HTTPErrorHandler handle the status code from http request
func HTTPErrorHandler(code int, msg string) (err error) {
	if code == http.StatusOK {
		return nil
	}
	if code == http.StatusCreated {
		return nil
	}
	if code == http.StatusUnauthorized {
		return fmt.Errorf("Not Logged in")
	}
	return fmt.Errorf("%s", msg)
}
