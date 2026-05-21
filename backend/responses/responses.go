package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

type SuccessResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func SendError(c echo.Context, status int, msg string) error {
	return c.JSON(status, ErrorResponse{
		Error:   http.StatusText(status),
		Message: msg,
	})
}

func SendCreated(c echo.Context, data interface{}, msg string) error {
	return c.JSON(http.StatusCreated, SuccessResponse{
		Data:    data,
		Message: msg,
	})
}

func SendOK(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}
