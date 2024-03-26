package model

type StudentRegister struct {
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	NIM      string `json:"nim" binding:"required,max=15,min=15"`
	Major    string `json:"major" binding:"required"`
	Faculty  string `json:"faculty" binding:"required"`
	Semester uint   `json:"semester" binding:"required,number"`
	Contact  string `json:"contact" binding:"required,e164,min=10,max=13"`
}
