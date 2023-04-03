package entity

import "time"

type Course struct {
	ID          uint      `gorm:"primarykey" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name" binding:"required"`
	Code        string    `json:"code" binding:"required" gorm:"unique"`
	SKS         uint      `json:"sks" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Lecturer    string    `json:"lecturer" binding:"required"`
}
