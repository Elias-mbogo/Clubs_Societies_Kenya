package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username     string
	Firstname    string
	Lastname     string
	UserEmail    string
	UserPassword string
	GroupsID     uint
	Chat         []Chats
}
