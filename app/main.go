package main

import (
	"net/http"

<<<<<<< HEAD
	"dena-hackathon21/auth"
=======
>>>>>>> aab192966f06bae071bbc70cc8eecf61b5730a4d
	"dena-hackathon21/handler"
	"dena-hackathon21/repository"
	"dena-hackathon21/sql_handler"
	"dena-hackathon21/twitter_handler"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// TODO 環境変数から取りたい
	sqlHandler, err := sql_handler.NewHandler("user:password@tcp(db:3306)/test_database")
<<<<<<< HEAD
	userRepository := repository.NewUserRepository(sqlHandler)
	twitterHandler, _ := twitter_handler.NewTwitterHandler()
	jwtHandler, _ := auth.NewJWTHandler()

=======
>>>>>>> aab192966f06bae071bbc70cc8eecf61b5730a4d
	if err != nil {
		fmt.Printf("connect error: %s\n", err.Error())
		panic(1)
	}

	contactHandler := handler.NewContactHandler(
		repository.NewContactRepository(sqlHandler),
	)

	// TODO 内容がダミーなので後で消す
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})

	e.POST("/api/contact", contactHandler.Send)

	// TODO issue-7
	e.GET("/api/contact", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// TODO 疎通確認用なので後で消す
	e.GET("/users/1", func(c echo.Context) error {

		userRepository := repository.NewUserRepository(sqlHandler)
		user, err := userRepository.GetUser(c.Request().Context(), 1)

		if err != nil {
			return c.String(500, fmt.Sprintf("db scan error: %s", err.Error()))
		}
		return c.String(http.StatusOK, fmt.Sprintf("id: %d, username: %s", user.ID, user.Username))
	})

<<<<<<< HEAD
	userHandler, _ := handler.NewUserHandler(userRepository, twitterHandler, jwtHandler)
	e.GET("/api/users/twitter_signup_url", userHandler.GetTwitterSignUpURL)
	e.GET("/api/users/twitter_signin_url", userHandler.GetTwitterSignInURL)
	e.POST("/api/users/signin", userHandler.SignIn)
	e.POST("/api/users/signup", userHandler.SignIn)

=======
	// Start
>>>>>>> aab192966f06bae071bbc70cc8eecf61b5730a4d
	e.Logger.Fatal(e.Start(":8080"))
}