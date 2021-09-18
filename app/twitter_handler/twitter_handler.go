package twitter_handler

import (
	// "context"
	// "fmt"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"os"
	// "io/ioutil"
	// "net/http"
	"net/url"
)

type TwitterHandler struct {
	oauth1Config oauth1.Config
}

func NewTwitterHandler() (*TwitterHandler, error) {
	config := oauth1.Config{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECERT"),
		CallbackURL:    os.Getenv("CALLBACK_URL"),
		Endpoint:       twitter.AuthorizeEndpoint,
	}
	return &TwitterHandler{
		oauth1Config: config,
	}, nil
}

func (t TwitterHandler) GetRequestToken() (string, string, error) {
	requestToken, requestSecret, err := t.oauth1Config.RequestToken()
	return requestToken, requestSecret, err
}

func (t TwitterHandler) GetAuthorizationURL(requestToken string) (*url.URL, error) {
	authorizationURL, err := t.oauth1Config.AuthorizationURL(requestToken)
	return authorizationURL, err
}

func (t TwitterHandler) GetAccessToken(requestToken string, requestSecret string, verifier string) (*oauth1.Token, error) {
	accessToken, accessSecret, _ := t.oauth1Config.AccessToken(requestToken, requestSecret, verifier)
	// handle error
	token := oauth1.NewToken(accessToken, accessSecret)

	return token, nil
}
