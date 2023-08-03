package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Like struct {
	gorm.Model
	PostId    uuid.UUID `gorm:"type:varchar(36);primaryKey;not null"`
	UserId    uuid.UUID `gorm:"type:varchar(36);primaryKey;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
