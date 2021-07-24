package main

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

func routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		posts := []string{
			"Foobar",
			"Barfoo",
			"Foobaz",
		}

		return c.Render(http.StatusOK, "templates/index.html",
			pongo2.Context{"title": "hello echo fw", "posts": posts})
	})
}
