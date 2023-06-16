package main

import (
  "github.com/labstack/echo"

	"go-server/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", handler.Hello)

	return e
}