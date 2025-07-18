// Package errors provides base errors for application and
// custom error handler for HTTP-server.
package errors

import (
	goerrors "errors"
	"net/http"
)

var (
	ErrValidateData = goerrors.New("validate data")    // HTTP code 400
	ErrNotFound     = goerrors.New("record not found") // HTTP code 404
)

// ErrorCode returns HTTP-code for given error.
// Given error is compared with the errors declared above.
func ErrorCode(err error) int {
	switch {
	case goerrors.Is(err, ErrValidateData):
		return http.StatusBadRequest
	case goerrors.Is(err, ErrNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
