package entities

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Login     string    `gorm:"column:login"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()"`
	UpdateAt  time.Time `gorm:"default:current_timestamp()"`
}

func (u *User) TableName() string {
	return "user"
}
