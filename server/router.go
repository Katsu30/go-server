package main

import (
  "github.com/labstack/echo"

	"github.com/Katsu30/go-server/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", handler.Hello)

	return e
}