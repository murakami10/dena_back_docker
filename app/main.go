package main

import (
	"net/http"

	"dena-hackathon21/SQLHandler"
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// TODO 環境変数から取りたい
	sqlHandler, err := SQLHandler.NewHandler("user:password@tcp(db:3306)/test_database")

	if err != nil {
		fmt.Printf("connect error: %s\n", err.Error())
		panic(1)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})

	e.GET("/users/1", func(c echo.Context) error {
		query := "select id, username from users where id=1"
		rows, err := sqlHandler.QueryContext(c.Request().Context(), query)

		if err != nil {
			return c.String(500, "db exec error")
		}

		var id string
		var username string
		rows.Next()
		err = rows.Scan(&id, &username)
		if err != nil {
			return c.String(500, fmt.Sprintf("db scan error: %s", err.Error()))
		}
		return c.String(http.StatusOK, fmt.Sprintf("id: %s, username: %s", id, username))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
