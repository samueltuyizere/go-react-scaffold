package main

import (
	"fmt"
	"log"
	_ "time/tzdata"

	"backend/auth"
	"backend/configs"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var app = echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	configs.ConnectDB()

	// public routes
	app.POST("/register", auth.HandleUserRegistration)
	app.POST("/login", auth.HandleUserLogin)

	port := fmt.Sprintf(":%s", configs.EnvPort())
	log.Fatal(app.Start(port))

}
