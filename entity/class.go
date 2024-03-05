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
	Description  string    `json:"description"`
	Lecturer     string    `json:"lecturer"`
	Course       []Course  `json:"course" foreignKey:"courseID"`
	StudentID    int       `json:"student_id"`
	Student      []Student `json:"student" foreignKey:"studentID"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
