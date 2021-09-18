package api_model

import (
	"time"
)

type GetContactResponse struct {
	ReceivedContactList []ContactItem `json:"received_contact_list"`
	PastMessageList	 []ContactItem `json:"past_message_list"`
}

type ConctactItem struct {
	SenderID uint64 `json:"sender_id"`
	Message string `json:"message"`
	RoomID uint64 `json:"room_id"`
	SendAt time.Time `json:"send_at"`
}