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
	sqlHandler, _ := SQLHandler.NewHandler("user:password@tcp(db:3306)test_database")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})

	e.GET("/users/1", func(c echo.Context) error {
		query := "select id, name from users where id=1"
		rows, err := sqlHandler.QueryContext(c.Request().Context(), query)

		if err != nil {
			return c.String(500, "db connect error")
		}

		var id string
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return c.String(500, "db scan error")
		}
		return c.String(http.StatusOK, fmt.Sprintf("id: %s, name: %s", id, name))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
