package main

import (
	"onlysync/configs"
	"onlysync/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS default
	//Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	e.Use(middleware.CORS())

	//CORS restricted
	//Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	//wth GET, PUT, POST or DELETE method.

	//run database
	configs.ConnectDB()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello from OnlySync API", "status": "success"})
	})
	//routes
	routes.LoginRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
