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

func (c ContactRepository) IsExistRoom(ctx context.Context, sender_id uint64, receiver_id uint64) (bool, error) {
	return false, nil
}

func (c ContactRepository) GetReceivedContact(ctx context.Context, user_id uint64) ([]api_model.ContactItem, error) {
	var contactItemList []api_model.ContactItem
	
	query := `
	select sender_id, message, created_at
	from requests
	where receiver_id = ?
	`
	rows, err := c.sqlHandler.QueryContext(ctx, query, user_id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var contactItem api_model.ContactItem
		err = rows.Scan(&contactItem.SenderID, &contactItem.Message, &contactItem.SendAt)
		if err != nil {
			return nil, err
		}

		contactItemList = append(contactItemList, contactItem)
	}
	
	return contactItemList, nil
}

func (c ContactRepository) GetPastContact(ctx context.Context, user_id uint64) ([]api_model.ContactItem, error) {
	var contactItemList []api_model.ContactItem
	
	query := `
	select c2.sender_id, c2.text, c2.room_id, c2.created_at 
	from chats as c2 
	where c2.created_at in (
	select max(t.created_at) 
	from (
		select c.sender_id, c.text, c.room_id, c.created_at 
		from chats as c 
		inner join room_members as m 
		on c.room_id = m.room_id 
		where m.user_id = ?
	) as t 
	group by room_id
	);
	`
	rows, err := c.sqlHandler.QueryContext(ctx, query, user_id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var contactItem api_model.ContactItem
		err = rows.Scan(&contactItem.SenderID, &contactItem.Message, &contactItem.RoomID, &contactItem.SendAt)
		if err != nil {
			return nil, err
		}

		contactItemList = append(contactItemList, contactItem)
	}
	
	return contactItemList, nil
}

func (c ContactRepository) GetRoomId(ctx context.Context, user_id uint64, sender_id uint64) (uint64, error) {
	var roomId uint64
	
	query := `
	select c. room_id 
	from chats as c 
	inner join room_members as m on 
	c.room_id = m.room_id 
	where m.user_id = ? and c.sender_id =?;
	`
	rows, err := c.sqlHandler.QueryContext(ctx, query, user_id, sender_id)
	if err != nil {
		return 0, err
	}

	rows.Next()
	err = rows.Scan(&roomId)
	if err != nil {
		return 0, err
	}
	
	return roomId, nil
}