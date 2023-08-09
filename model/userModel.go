package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:varchar(45);primary_key"`
	Username  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Status    string
	Role      string `gorm:"not null"`
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Birthday  time.Time `gorm:"not null"`
}

type APIUser struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:varchar(45);primary_key"`
	Username  string    `gorm:"not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Status    string
	Role      string `gorm:"not null"`
	Avatar    string
	Birthday  time.Time `gorm:"not null"`
}
