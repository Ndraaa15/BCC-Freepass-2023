package entity

import "time"

type Admin struct {
	ID        uint      `gorm:"primarykey" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username" binding:"required" gorm:"unique"`
	Password  string    `json:"password" binding:"required"`
}
