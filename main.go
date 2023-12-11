package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pharrisee/templ-example/content/views"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		// call the templ generated code
		return views.Layout("Home", views.Index()).Render(c.Request().Context(), c.Response())
	})

	e.Start(":1323")
}
