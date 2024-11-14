package model

import "time"

type Link struct {
	Id          int64
	Link        string
	Message     string
	ChatId      int64
	ScheduledAt time.Time
	PostedAt    time.Time
	CreatedAt   time.Time
}
