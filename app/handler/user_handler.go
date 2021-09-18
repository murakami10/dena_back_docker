package handler

import (
	"dena-hackathon21/auth"
	"dena-hackathon21/entity"
	"dena-hackathon21/repository"
	"dena-hackathon21/twitter_handler"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

type UserHandler struct {
	userRepository *repository.UserRepository
	twitterHandler *twitter_handler.TwitterHandler
	jwtHandler     *auth.JWTHandler
}

func NewUserHandler(userRepository *repository.UserRepository, twitterHandler *twitter_handler.TwitterHandler, jwtHandler *auth.JWTHandler) (*UserHandler, error) {
	return &UserHandler{
		userRepository: userRepository,
		twitterHandler: twitterHandler,
		jwtHandler:     jwtHandler,
	}, nil
}

type TwitterAuthToken struct {
	OAuthToken    string `json:"oauth_token"`
	OAuthSecret   string `json:"oauth_secret"`
	OAuthVerifier string `json:"oauth_verifier"`
}

func (u UserHandler) SignIn(c echo.Context) error {
	tat := TwitterAuthToken{}
	if err := c.Bind(tat); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, _ := u.twitterHandler.GetAccessToken(tat.OAuthToken, tat.OAuthSecret, tat.OAuthVerifier)
	twitterUser, err := u.twitterHandler.GetUserByToken(token)

	if err != nil {
		return c.String(401, "not auth")
	}
	user, _ := u.userRepository.GetUserByTwitterID(c.Request().Context(), twitterUser.ID)

	if user == nil {
		newUser := entity.User{
			Username:      twitterUser.Username,
			DisplayName:   twitterUser.Name,
			TwitterUserID: twitterUser.ID,
			IconURL:       twitterUser.ProfileImageURL,
		}
		user, _ = u.userRepository.CreateUser(c.Request().Context(), newUser)
	}

	jwtToken, _ := u.jwtHandler.GenerateJWTToken(user.ID)

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
}

func (u UserHandler) GetTwitterSignUpURL(c echo.Context) error {

	token, secret, _ := u.twitterHandler.GetRequestToken()
	url, _ := u.twitterHandler.GetAuthorizationURL(token, os.Getenv("SIGNUP_CALLBACK_URL"))
	jsonMap := map[string]string{
		"url":          url.String(),
		"oauth_token":  token,
		"oauth_secret": secret,
	}
	return c.JSON(http.StatusOK, jsonMap)
}

func (u UserHandler) GetTwitterSignInURL(c echo.Context) error {

	token, secret, _ := u.twitterHandler.GetRequestToken()
	url, _ := u.twitterHandler.GetAuthorizationURL(token, os.Getenv("SIGNIN_CALLBACK_URL"))
	jsonMap := map[string]string{
		"url":          url.String(),
		"oauth_token":  token,
		"oauth_secret": secret,
	}

	return c.JSON(http.StatusOK, jsonMap)
}
