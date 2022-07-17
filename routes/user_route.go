package routes

import (
    "go_code/controllers"

    "github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Group) {

	e.PUT("/user/:userId", controllers.EditUser)
	e.GET("/user/:userId", controllers.GetUserInfo)
	e.DELETE("/user/:userId", controllers.DeleteUser)
	e.GET("/users", controllers.GetAllUsers)

}