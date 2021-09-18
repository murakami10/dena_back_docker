package repository

import (
	"context"
	"dena-hackathon21/entity"
	"dena-hackathon21/sql_handler"
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

	query := "select id, username from users where id=?"
	rows, err := u.sqlHandler.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) GetUserByTwitterID(ctx context.Context, twitterID string) (*entity.User, error) {
	var user entity.User

	query := "select id, username from users where twitter_user_id=?"
	rows, err := u.sqlHandler.QueryContext(ctx, query, twitterID)

	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepository) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	query := `insert into Users (username, display_name, twitter_user_id, icon_url)　values (?, ?, ?, ?)`

	_, err := u.sqlHandler.QueryContext(ctx, query, user.Username, user.DisplayName, user.TwitterUserID, user.IconURL)

	if err != nil {
		return nil, err
	}

	query = "SELECT LAST_INSERT_ID()"
	rows, err := u.sqlHandler.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	rows.Next()
	err = rows.Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
