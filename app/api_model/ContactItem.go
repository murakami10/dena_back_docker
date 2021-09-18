package api_model

type ConctactItem struct {
	SenderID uint64 `json:"sender_id"`
	Message string `json:"message"`
	RoomID uint64 `json:"room_id"`
	SendAt time.Time `json:"send_at"`
}