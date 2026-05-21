package main

import (
	"fmt"
	"log/slog"
	"os"
	_ "time/tzdata"

	"backend/auth"
	"backend/configs"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	configs.ConnectDB()
	slog.Info("database connected", "env", configs.AppEnv())

	app.POST("/register", auth.HandleUserRegistration)
	app.POST("/login", auth.HandleUserLogin)

	port := fmt.Sprintf(":%s", configs.EnvPort())
	slog.Info("server starting", "port", port)

	if err := app.Start(port); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
