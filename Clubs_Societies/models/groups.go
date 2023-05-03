package models

import (
	"gorm.io/gorm"
)

type Groups struct {
	gorm.Model
	Name    string
	ClubsID uint
	User    []Users
	Post    []Posts
	Chat    []Chats
}
