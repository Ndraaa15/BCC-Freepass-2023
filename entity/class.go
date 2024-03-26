package entity

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	CodeClass    string    `json:"codeClass"`
	TotalStudent uint      `json:"totalStudent"`
	Lecturer     string    `json:"lecturer"`
	Course       Course    `json:"course"`
	User         []User    `json:"student"`
	TimeStart    string    `json:"timeStart"`
	TimeEnd      string    `json:"timeEnd"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
