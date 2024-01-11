package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string
}

type Users = []User
