package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func show(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, u)
}

func display(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})
	e.GET("/user", show)
	e.POST("/users", display)
	e.Logger.Fatal(e.Start(":8080"))
}
