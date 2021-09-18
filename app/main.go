package main

import (
	"net/http"

	"dena-hackathon21/auth"
	"dena-hackathon21/entity"
	"dena-hackathon21/repository"
	"dena-hackathon21/sql_handler"
	"dena-hackathon21/twitter_handler"
	"fmt"
	"github.com/labstack/echo"
	"os"
	"time"
)

type TwitterAuthToken struct {
	OAuthToken    string `json:"oauth_token"`
	OAuthSecret   string `json:"oauth_secret"`
	OAuthVerifier string `json:"oauth_verifier"`
}

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
		return c.String(http.StatusOK, fmt.Sprintf("id: %d, username: %s", user.ID, user.Username))
	})

	e.GET("/api/users/twitter_signin_url", func(c echo.Context) error {

		twitterHandler, _ := twitter_handler.NewTwitterHandler()

		token, secret, _ := twitterHandler.GetRequestToken()
		url, _ := twitterHandler.GetAuthorizationURL(token, os.Getenv("SIGNIN_CALLBACK_URL"))
		jsonMap := map[string]string{
			"url":          url.String(),
			"oauth_token":  token,
			"oauth_secret": secret,
		}

		return c.JSON(http.StatusOK, jsonMap)
	})

	e.GET("/api/users/twitter_signup_url", func(c echo.Context) error {

		twitterHandler, _ := twitter_handler.NewTwitterHandler()

		token, secret, _ := twitterHandler.GetRequestToken()
		url, _ := twitterHandler.GetAuthorizationURL(token, os.Getenv("SIGNUP_CALLBACK_URL"))
		jsonMap := map[string]string{
			"url":          url.String(),
			"oauth_token":  token,
			"oauth_secret": secret,
		}
		return c.JSON(http.StatusOK, jsonMap)
	})

	e.GET("/twitter_login", func(c echo.Context) error {
		oauthToken := c.QueryParam("oauth_token")
		oauthVerifier := c.QueryParam("oauth_verifier")
		url := fmt.Sprintf("localhost:8080/token?oauth_token=%s&oauth_verifier=%s&oauth_secret=?", oauthToken, oauthVerifier)

		return c.String(http.StatusOK, fmt.Sprintf("token: %s, verifier: %s, url: %s", oauthToken, oauthVerifier, url))
	})

	e.POST("/api/users/signin", func(c echo.Context) error {

		tat := TwitterAuthToken{}
		if err = c.Bind(tat); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		twitterHandler, _ := twitter_handler.NewTwitterHandler()
		token, _ := twitterHandler.GetAccessToken(tat.OAuthToken, tat.OAuthSecret, tat.OAuthVerifier)
		twitterUser, err := twitterHandler.GetUserByToken(token)

		if err != nil {
			return c.String(401, "not auth")
		}

		userRepository := repository.NewUserRepository(sqlHandler)
		user, _ := userRepository.GetUserByTwitterID(c.Request().Context(), twitterUser.ID)

		if user == nil {
			newUser := entity.User{
				Username:      twitterUser.Username,
				DisplayName:   twitterUser.Name,
				TwitterUserID: twitterUser.ID,
				IconURL:       twitterUser.ProfileImageURL,
			}
			user, _ = userRepository.GetUserByTwitterID(c.Request().Context(), newUser)
		}

		jwtHandler, _ := auth.NewJWTHandler()
		jwtToken, _ := jwtHandler.GenerateJWTToken(user.ID)

		// set cookie
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = jwtToken
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		jsonMap := map[string]entity.User{
			"user": *user,
		}
		return c.JSON(http.StatusOK, jsonMap)
	})

	e.POST("/api/users/signup", func(c echo.Context) error {

		tat := TwitterAuthToken{}
		if err = c.Bind(tat); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		twitterHandler, _ := twitter_handler.NewTwitterHandler()
		token, _ := twitterHandler.GetAccessToken(tat.OAuthToken, tat.OAuthSecret, tat.OAuthVerifier)
		twitterUser, err := twitterHandler.GetUserByToken(token)

		if err != nil {
			return c.String(401, "not auth")
		}

		userRepository := repository.NewUserRepository(sqlHandler)
		user, _ := userRepository.GetUserByTwitterID(c.Request().Context(), twitterUser.ID)

		if user == nil {
			newUser := entity.User{
				Username:      twitterUser.Username,
				DisplayName:   twitterUser.Name,
				TwitterUserID: twitterUser.ID,
				IconURL:       twitterUser.ProfileImageURL,
			}
			user, _ = userRepository.GetUserByTwitterID(c.Request().Context(), newUser)
		}

		jwtHandler, _ := auth.NewJWTHandler()
		jwtToken, _ := jwtHandler.GenerateJWTToken(user.ID)

		// set cookie
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = jwtToken
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		jsonMap := map[string]entity.User{
			"user": *user,
		}
		return c.JSON(http.StatusOK, jsonMap)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
