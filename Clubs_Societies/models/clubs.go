package models

import (
	"gorm.io/gorm"
)

var Club *Clubs

type Clubs struct {
	gorm.Model
	Name  string
	Group []Groups
	Post  []Posts
}
