package main

import (
	"net/http"

	"dena-hackathon21/repository"
	"dena-hackathon21/sql_handler"
	"dena-hackathon21/twitter_handler"
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
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

	e.GET("/request_url", func(c echo.Context) error {

		twitterHandler, _ := twitter_handler.NewTwitterHandler()

		token, secret, _ := twitterHandler.GetRequestToken()
		url, _ := twitterHandler.GetAuthorizationURL(token)
		fmt.Println(secret)
		return c.String(http.StatusOK, fmt.Sprintf("url:%s, secret:%s", url.String(), secret))
	})

	e.GET("/twitter_login", func(c echo.Context) error {
		oauthToken := c.QueryParam("oauth_token")
		oauthVerifier := c.QueryParam("oauth_verifier")
		url := fmt.Sprintf("localhost:8080/token?oauth_token=%s&oauth_verifier=%s&oauth_secret=?", oauthToken, oauthVerifier)

		return c.String(http.StatusOK, fmt.Sprintf("token: %s, verifier: %s, url: %s", oauthToken, oauthVerifier, url))
	})

	e.GET("/token", func(c echo.Context) error {
		oauthToken := c.QueryParam("oauth_token")
		oauthVerifier := c.QueryParam("oauth_verifier")
		oauthSecret := c.QueryParam("oauth_secret")

		twitterHandler, _ := twitter_handler.NewTwitterHandler()
		token, _ := twitterHandler.GetAccessToken(oauthToken, oauthSecret, oauthVerifier)
		// userInfo, _ := twitterHandler.GetTwitterUserInfo(c.Request().Context(), token)

		return c.String(http.StatusOK, token.Token)
		// return c.String(http.StatusOK, fmt.Sprintf("token: %s, secret:%s", token.Token, token.TokenSecret))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
