package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique_index"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool
}

func (u *User) BeforeCreate(scope *gorm.DB) error {
	scope.UpdateColumn("CreatedAt", time.Now())
	return nil
}

func (u *User) BeforeUpdate(scope *gorm.DB) error {
	scope.UpdateColumn("UpdatedAt", time.Now())
	return nil
}
