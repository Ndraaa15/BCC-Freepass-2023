package entity

import "time"

type Student struct {
	ID        uint      `gorm:"primarykey" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FullName  string    `json:"fullName" binding:"required"`
	NIM       string    `json:"nim" binding:"required" gorm:"unique"`
	Major     string    `json:"major" binding:"required"`
	TotalSKS  uint      `json:"total_sks" binding:"required"`
	Email     string    `json:"email" binding:"required" gorm:"unique"`
	Username  string    `json:"username" binding:"required" gorm:"unique"`
	Password  string    `json:"password" binding:"required"`
	Contact   string    `json:"contact" binding:"required,e164"`
}
