package auth

import (
	"net/mail"
	"strings"

	"backend/responses"

	"github.com/labstack/echo/v4"
)

type loginResult struct {
	Token  string `json:"token"`
	Status bool   `json:"status"`
}

func HandleUserLogin(ctx echo.Context) error {
	email := strings.TrimSpace(ctx.FormValue("email"))
	password := ctx.FormValue("password")

	if email == "" || password == "" {
		return responses.SendError(ctx, 400, "email and password are required")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return responses.SendError(ctx, 400, "invalid email format")
	}

	token, status, err := processUserLogin(email, password)
	if err != nil {
		return responses.SendError(ctx, 401, "invalid credentials")
	}

	return responses.SendOK(ctx, loginResult{
		Token:  token,
		Status: status,
	})
}

func HandleUserRegistration(ctx echo.Context) error {
	email := strings.TrimSpace(ctx.FormValue("email"))
	password := ctx.FormValue("password")

	if email == "" || password == "" {
		return responses.SendError(ctx, 400, "email and password are required")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return responses.SendError(ctx, 400, "invalid email format")
	}

	if len(password) < 8 {
		return responses.SendError(ctx, 400, "password must be at least 8 characters")
	}

	user, err := createNewUser(email, []byte(password))
	if err != nil {
		return responses.SendError(ctx, 500, "failed to create user")
	}

	return responses.SendCreated(ctx, user, "user registered successfully")
}
