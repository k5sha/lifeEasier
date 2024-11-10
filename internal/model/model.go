package model

import "time"

type User struct {
	Id        int64
	Username  string
	ChatId    int64
	CreatedAt time.Time
}

type Link struct {
	Id        int64
	Link      string
	Message   string
	Used      bool
	UserId    int64
	CreatedAt time.Time
}
