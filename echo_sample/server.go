package main

import (
	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// echo のインスタンスを作る
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello/:username", handler.MainPage())

	e.Start(":1323")
}
