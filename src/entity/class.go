package entity

import "time"

type Class struct {
	ID           uint      `gorm:"primarykey" binding:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name" binding:"required"`
	CodeClass    string    `json:"code_class" binding:"required" gorm:"unique"`
	TotalStudent uint      `json:"total_student" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Lecturer     string    `json:"lecturer" binding:"required"`
	StatusID     int       `json:"status_id" binding:"required"`
	CourseID     int       `json:"course_id" binding:"required"`
	Course       []Course  `json:"course" binding:"required" foreignKey:"courseID"`
	StudentID    int       `json:"student_id" binding:"required"`
	Student      []Student `json:"student" binding:"required" foreignKey:"studentID"`
}
