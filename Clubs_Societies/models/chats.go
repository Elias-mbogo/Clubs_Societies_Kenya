package models

import (
	"gorm.io/gorm"
)

type Chats struct {
	gorm.Model
	Body     string
	UsersID  uint
	GroupsID uint
}
