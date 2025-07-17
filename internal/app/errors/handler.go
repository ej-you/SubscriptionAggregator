package errors

import (
	goerrors "errors"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
)

// CustomErrorHandler is a handler for http server errors.
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	msg := err.Error()

	// if resource was not found
	var fiberErr *fiber.Error
	if goerrors.As(err, &fiberErr) && strings.HasPrefix(fiberErr.Message, "Cannot GET") {
		msg = "resource not found"
	}
	// get http error code
	errStatusCode := ErrorCode(err)
	// send error response
	return ctx.Status(errStatusCode).JSON(msg)
}
