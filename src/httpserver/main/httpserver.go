package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/users/:name", getUserName)
	e.GET("/show", show)
	e.POST("/save", save)

	e.Logger.Fatal(e.Start(":1323"))
	
}

func save(c echo.Context) error {

	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.String(http.StatusOK, "name: "+name+", email:"+email)

}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team :"+team+", member :"+member)
}

func getUserName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}