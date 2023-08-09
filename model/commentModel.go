package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:varchar(45);primary_key"`
	Content   string    `gorm:"not null"`
	UserId    uuid.UUID `gorm:"type:varchar(45);not null"`
	PostId    uuid.UUID `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
