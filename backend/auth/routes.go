package auth

import (
	"github.com/labstack/echo/v4"
)

type loginResult struct {
	token string
	status bool
	error error
}

func HandleUserLogin(ctx echo.Context) error {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	token, status, err := processUserLogin(email, password)
	if err != nil {
		return ctx.JSON(401, err)
	}
	rslt := loginResult{
		token: token,
		status: status,
		error: nil,
	}
	return ctx.JSON(200, rslt)
}

func HandleUserRegistration(ctx echo.Context) error {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	_, err := createNewUser(email, []byte(password))
	if err != nil {
		return ctx.JSON(401, err)
	}

	return ctx.JSON(201, "Success")
}
