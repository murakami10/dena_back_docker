package entity

import (
	"time"
)

type Contact struct {
	Id			uint64
	sender_id	uint64
	receiver_id	uint64
	message		string
	created_at	time.Time
}