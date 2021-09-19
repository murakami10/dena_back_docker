package repository

import (
	"context"
	"dena-hackathon21/entity"
	"dena-hackathon21/sql_handler"
	"fmt"
)

type UserRepository struct {
	sqlHandler *sql_handler.SQLHandler
}

func NewUserRepository(sqlHandler *sql_handler.SQLHandler) *UserRepository {
	return &UserRepository{
		sqlHandler: sqlHandler,
	}
}

func (u UserRepository) GetUser(ctx context.Context, userID uint64) (*entity.User, error) {
	var user entity.User

	query := "select id, username, display_name, twitter_user_id, icon_url from users where id=?"
	rows, err := u.sqlHandler.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(&user.ID, &user.Username, &user.DisplayName, &user.TwitterUserID, &user.IconURL)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) GetUserByTwitterID(ctx context.Context, twitterID string) (*entity.User, error) {
	var user entity.User

	query := "select id, username, display_name, twitter_user_id, icon_url from users where twitter_user_id=?"
	rows, err := u.sqlHandler.QueryContext(ctx, query, twitterID)

	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(&user.ID, &user.Username, &user.DisplayName, &user.TwitterUserID, &user.IconURL)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	query := `insert into users (username, display_name, twitter_user_id, icon_url) values (?, ?, ?, ?)`

	_, err := u.sqlHandler.QueryContext(ctx, query, user.Username, user.DisplayName, user.TwitterUserID, user.IconURL)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	newUser, err := u.GetUserByTwitterID(ctx, user.TwitterUserID)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u UserRepository) GetFriends(ctx context.Context, userID uint64) ([]entity.User, error) {
	query := "select users.id, users.username, users.display_name, users.twitter_user_id, users.icon_url from friends JOIN users on users.id = friends.friend_user_id where friends.user_id=?"

	rows, err := u.sqlHandler.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, err
	}

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Username, &user.DisplayName, &user.TwitterUserID, &user.IconURL)
		users = append(users, user)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}
