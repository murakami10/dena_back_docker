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
	err = rows.Scan(&user.Id, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
