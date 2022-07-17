package routes

import (
    "go_code/controllers"
    "github.com/labstack/echo/v4"
)

func PublicRoute(e *echo.Echo) {

	auth := e.Group("/auth")

	auth.GET("/login", controllers.Login)
    auth.POST("/signup", controllers.SignUp)
}