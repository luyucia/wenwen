package main

import (
	"github.com/labstack/echo"
	"wenwen/src/controllers"
)

func initRouter(r *echo.Echo) {

	r.GET("/ping", controllers.Ping)

}
