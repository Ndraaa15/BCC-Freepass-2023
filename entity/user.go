package entity

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email" bindig:"email"`
	Password  string    `json:"password"`
	FullName  string    `json:"fullName"`
	NIM       string    `json:"nim"`
	Faculty   string    `json:"faculty"`
	Major     string    `json:"major"`
	TotalSKS  uint      `json:"totalSKS"`
	Username  string    `json:"username"`
	Contact   string    `json:"contact" binding:"e164"`
	Roles     uint      `json:"roles"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
