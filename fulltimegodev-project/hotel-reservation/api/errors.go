package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}

func (e Error) Error() string {
	return e.Message
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		return ctx.Status(apiError.Code).JSON(apiError)
	}
	apiError := NewError(http.StatusInternalServerError, err.Error())
	return ctx.Status(apiError.Code).JSON(apiError)
}

func ErrInvalidID() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid ID",
	}
}

func ErrUnAuthorized() Error {
	return Error{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized User",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid JSON request",
	}
}

func ErrResourceNotFound(res string) Error {
	return Error{
		Code:    http.StatusNotFound,
		Message: res + "Resource Not Found",
	}
}
