// Package errors provides base errors for application and
// custom error handler for HTTP-server.
package errors

import (
	goerrors "errors"
	"net/http"
)

var (
	ErrValidateData = goerrors.New("validate data") // HTTP code 400
	// ErrFilesLimit         = goerrors.New("files limit for task is reached") // HTTP code 400
	// ErrFileIsNotSupported = goerrors.New("file is not supported")           // HTTP code 400
	ErrNotFound = goerrors.New("record not found") // HTTP code 404
	// ErrServiceUnavailable = goerrors.New("server is busy")                  // HTTP code 503
)

// ErrorCode returns HTTP-code for given error.
// Given error is compared with the errors declared above.
func ErrorCode(err error) int {
	switch {
	case goerrors.Is(err, ErrValidateData):
		return http.StatusBadRequest
	// case goerrors.Is(err, ErrFilesLimit):
	// return http.StatusBadRequest
	// case goerrors.Is(err, ErrFileIsNotSupported):
	// return http.StatusBadRequest
	case goerrors.Is(err, ErrNotFound):
		// return http.StatusNotFound
		// case goerrors.Is(err, ErrServiceUnavailable):
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}
