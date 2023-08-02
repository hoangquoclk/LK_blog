package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:varchar(36);primary_key"`
	Username  string
	Password  string
	FirstName string
	LastName  string
	Email     string
	Status    string
	Role      string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Birthday  time.Time
}
