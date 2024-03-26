package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashedPassword"`
	FullName       string    `json:"fullName"`
	NIM            string    `json:"nim"`
	Major          string    `json:"major"`
	Faculty        string    `json:"faculty"`
	TotalSKS       uint      `json:"totalSks"`
	Semester       uint      `json:"semester"`
	Contact        string    `json:"contact"`
	Roles          uint      `json:"roles"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
