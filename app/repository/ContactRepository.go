package repository

import (
	"context"
	"dena-hackathon21/api_model"
	"dena-hackathon21/sql_handler"
)

type ContactRepository struct {
	sqlHandler *sql_handler.SQLHandler
}

func NewContactRepository(sqlHandler *sql_handler.SQLHandler) *ContactRepository {
	return &ContactRepository{
		sqlHandler: sqlHandler,
	}
}

func (c ContactRepository) SendContact(ctx context.Context, sender_id uint64, receiver_id uint64, message string) error {
	query := `
	insert into requests (sender_id, receiver_id, message)
	values (?, ?, ?)
	`
	_, err := c.sqlHandler.QueryContext(ctx, query, sender_id, receiver_id, message)
	if err != nil {
		return err
	}
	return nil
}

func (c ContactRepository) GetReceivedContact(ctx context.Context, sender_id uint64) ([]*api_model.ContactItem, error) {
	return nil, nil
}
