package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:varchar(36);primary_key"`
	Content   string
	UserId    string `gorm:"type:varchar(36)"`
	PostId    string `gorm:"type:varchar(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
