package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:varchar(45);primary_key"`
	Title      string    `gorm:"not null"`
	Content    string    `gorm:"not null"`
	CategoryId uuid.UUID `gorm:"type:varchar(45);not null"`
	UserId     uuid.UUID `gorm:"type:varchar(45);not null"`
	Author     string    `gorm:"not null"`
	Like       uint
	Image      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
