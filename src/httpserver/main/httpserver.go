package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name string `json: "name"`
	Email string `json: "email"`
}

type Message struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Message string `json: "message"`
}

type Response struct {
	Name string `json: "name"`
	Email string `json: "email"`
	Message string `json:"message"`
	Stusts string `json: "status"`
}

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/users/:name", getUserName)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/saveuser", saveUser)
	e.POST("/send", sendMessage)

	e.Logger.Fatal(e.Start(":1323"))
	
}

func sendMessage(c echo.Context) error {

	m := new(Message)
	if error := c.Bind(m); error != nil {
		return error
	}

	r := new(Response)
	r.Name = m.Name
	r.Email = m.Email
	r.Message = m.Message
	r.Stusts = "success"

	return c.JSON(http.StatusOK, r)
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
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