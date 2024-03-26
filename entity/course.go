package entity

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	SKS       uint      `json:"sks"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
