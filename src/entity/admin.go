package entity

import "time"

type admin struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username" binding:"required" gorm:"unique"`
	Password  string    `json:"password" binding:"required"`
}
