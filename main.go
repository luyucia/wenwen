package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	//实例化echo对象。
	e := echo.New()

	InitRouter(e)

	e.Start(":80")
}
