package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Comment  string
	GroupsID uint
	ClubsID  uint
}
