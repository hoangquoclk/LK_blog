package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Like struct {
	gorm.Model
	PostId    uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	UserId    uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
