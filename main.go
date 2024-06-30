package main1

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main1() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", getUsers)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":3000"))
}

// e.GET("/users/:id", getUsers)
func getUsers(c echo.Context) error {
	// userId
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

// /show?team=x-men&member=wolverine
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

