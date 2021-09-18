package twitter_handler

import (
	"dena-hackathon21/entity"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	oauth1Twitter "github.com/dghubble/oauth1/twitter"
	"net/url"
	"os"
)

type TwitterHandler struct {
	oauth1Config oauth1.Config
}

func NewTwitterHandler() (*TwitterHandler, error) {
	config := oauth1.Config{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECERT"),
		CallbackURL:    os.Getenv("CALLBACK_URL"),
		Endpoint:       oauth1Twitter.AuthorizeEndpoint,
	}
	return &TwitterHandler{
		oauth1Config: config,
	}, nil
}

func (t TwitterHandler) GetRequestToken() (string, string, error) {
	requestToken, requestSecret, err := t.oauth1Config.RequestToken()
	return requestToken, requestSecret, err
}

func (t TwitterHandler) GetAuthorizationURL(requestToken string, callbackURL string) (*url.URL, error) {
	t.oauth1Config.CallbackURL = callbackURL
	authorizationURL, err := t.oauth1Config.AuthorizationURL(requestToken)
	return authorizationURL, err
}

func (t TwitterHandler) GetAccessToken(requestToken string, requestSecret string, verifier string) (*oauth1.Token, error) {
	accessToken, accessSecret, _ := t.oauth1Config.AccessToken(requestToken, requestSecret, verifier)
	// handle error
	token := oauth1.NewToken(accessToken, accessSecret)

	return token, nil
}

func (t TwitterHandler) GetUserByToken(token *oauth1.Token) (*entity.TwitterUser, error) {
	httpClient := t.oauth1Config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// TODO 一旦コピペしたので後でせいさする
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	return &entity.TwitterUser{
		ID:              user.IDStr,
		Name:            user.Name,
		Username:        user.ScreenName,
		Description:     user.Description,
		ProfileImageURL: user.ProfileImageURL,
	}, nil
}
