package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id" yaml:"id"`
	Name      string         `gorm:"size:100;not null" json:"name" yaml:"name"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email" yaml:"email"`
	Password  string         `gorm:"not null" json:"password" yaml:"password"`
	CreatedAt time.Time      `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" yaml:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" yaml:"-"`
}
