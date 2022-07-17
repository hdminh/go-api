package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_code/configs"
	"go_code/routes"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(configs.GetJWTKey()),
	}))
	//route config
	routes.PublicRoute(e)
	routes.UserRoute(api)
	//mongo connection
	configs.ConnectDB()
	e.Logger.Fatal(e.Start(":6000"))
}
