package routes

import (
	"onlysync/controllers"

	"github.com/labstack/echo/v4"
)

func LoginRoutes(e *echo.Echo) {
	// User signup route
	e.POST("/signup", controllers.Signup)

}
