package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:varchar(36);primary_key"`
	Title      string
	Content    string
	CategoryId string `gorm:"type:varchar(36)"`
	UserId     string `gorm:"type:varchar(36)"`
	Author     string
	Like       uint
	Image      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
