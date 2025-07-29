package errors

import (
	goerrors "errors"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// CustomErrorHandler is a handler for http server errors.
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	msg := err.Error()
	// get http error code
	errStatusCode := ErrorCode(err)

	// if resource was not found
	var fiberErr *fiber.Error
	if goerrors.As(err, &fiberErr) {
		errStatusCode = fiberErr.Code
		if strings.HasPrefix(fiberErr.Message, "Cannot GET") {
			msg = "resource not found"
		}
	}

	// log error
	logrus.Errorf("%d: %s", errStatusCode, msg)
	// send error response
	return ctx.Status(errStatusCode).JSON(msg)
}
