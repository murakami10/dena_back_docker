package main

import (
	"net/http"

	"dena-hackathon21/repository"
	"dena-hackathon21/sql_handler"
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// TODO 環境変数から取りたい
	sqlHandler, err := sql_handler.NewHandler("user:password@tcp(db:3306)/test_database")

	if err != nil {
		fmt.Printf("connect error: %s\n", err.Error())
		panic(1)
	}

	// TODO 内容がダミーなのでなんとかする
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})

	// TODO issue-7
	e.GET("/contact", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// TODO 疎通確認用なので後で消す
	e.GET("/users/1", func(c echo.Context) error {

		userRepository := repository.NewUserRepository(sqlHandler)
		user, err := userRepository.GetUser(c.Request().Context(), 1)

		if err != nil {
			return c.String(500, fmt.Sprintf("db scan error: %s", err.Error()))
		}
		return c.String(http.StatusOK, fmt.Sprintf("id: %d, username: %s", user.Id, user.Username))
	})

	// Start
	e.Logger.Fatal(e.Start(":8080"))
}
