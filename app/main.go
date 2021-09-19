package main

import (
	"dena-hackathon21/auth"
	"dena-hackathon21/handler"
	"dena-hackathon21/repository"
	"dena-hackathon21/sql_handler"
	"dena-hackathon21/twitter_handler"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	sqlAuthentication := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	sqlHandler, err := sql_handler.NewHandler(sqlAuthentication)

	if err != nil {
		fmt.Printf("connect error: %s\n", err.Error())
		panic(1)
	}

	userRepository := repository.NewUserRepository(sqlHandler)
	twitterHandler, _ := twitter_handler.NewTwitterHandler()
	jwtHandler, _ := auth.NewJWTHandler()

	contactRepository := repository.NewContactRepository(sqlHandler)

	// TODO 内容がダミーなので後で消す
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})

	contactHandler := handler.NewContactHandler(contactRepository, jwtHandler)
	e.POST("/api/contacts/:id", contactHandler.Send)
	e.GET("/api/contacts/:id", contactHandler.Get)

	// TODO 疎通確認用なので後で消す
	e.GET("/users/1", func(c echo.Context) error {

		userRepository := repository.NewUserRepository(sqlHandler)
		user, err := userRepository.GetUser(c.Request().Context(), 1)

		if err != nil {
			return c.String(500, fmt.Sprintf("db scan error: %s", err.Error()))
		}
		return c.String(http.StatusOK, fmt.Sprintf("id: %d, username: %s", user.ID, user.Username))
	})

	userHandler, _ := handler.NewUserHandler(userRepository, twitterHandler, jwtHandler)
	e.GET("/api/users/twitter_signup_url", userHandler.GetTwitterSignUpURL)
	e.GET("/api/users/twitter_signin_url", userHandler.GetTwitterSignInURL)
	e.POST("/api/users/signin", userHandler.SignIn)
	e.POST("/api/users/signup", userHandler.SignIn)

	e.GET("/api/users/:id/frends", userHandler.GetFrends)

	e.GET("/api/users/session", userHandler.GetSession)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
