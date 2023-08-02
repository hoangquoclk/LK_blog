package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Like struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:varchar(36);primary_key"`
	PostId    string    `gorm:"type:varchar(36)"`
	UserId    string    `gorm:"type:varchar(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
