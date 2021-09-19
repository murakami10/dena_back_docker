package repository

import (
	"fmt"
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
	query := `
	select room_id from room_members where user_id in (?, ?) group by room_id having count(*) >= 2;
	`
	rows, err := c.sqlHandler.QueryContext(ctx, query, sender_id, receiver_id)
	if err != nil {
		return false, err
	}

	// TODO 準備中エラーと区別
	if ok := rows.Next(); !ok {
		// Roomが無かった場合
		return false, nil
	} else {
		// Roomがあった場合
		return true, nil
	}
}

func (c ContactRepository) CreateRoom(ctx context.Context, sender_id uint64, receiver_id uint64, message string) error {
	// ルームを作成
	query := `
	insert into rooms values ();
	`
	_, err := c.sqlHandler.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Create Room Error:", err)
		return err
	}

	// ルームIDを取得
	// TODO lastinsertidを使う
	var roomId uint64
	query = `
	select max(id) from rooms;
	`
	rows, err := c.sqlHandler.QueryContext(ctx, query)
	rows.Next()
	err = rows.Scan(&roomId)
	if err != nil {
		fmt.Println("Get RoomID Error:", err)
		return err
	}

	// RoomIDとuser_idを紐づける
	query = `
	insert into room_members (user_id, room_id) values (?, ?);
	`
	_, err = c.sqlHandler.QueryContext(ctx, query, sender_id, roomId)
	if err != nil {
		fmt.Println("Insert room_members Error:", err)
		return err
	}
	_, err = c.sqlHandler.QueryContext(ctx, query, receiver_id, roomId)
	if err != nil {
		fmt.Println("Insert room_members Error:", err)
		return err
	}

	// Chatsテーブルにレコードを追加
	query = `
	insert into chats (room_id, sender_id, text) values (?, ?, ?);
	`
	_, err = c.sqlHandler.QueryContext(ctx, query, roomId, sender_id, message)
	if err != nil {
		fmt.Println("Insert chats Error:", err)
		return err
	}
	return nil
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