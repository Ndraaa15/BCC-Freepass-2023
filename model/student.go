package model

type StudentRegister struct {
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	NIM      string `json:"nim" validate:"required,max=15,min=15"`
	Major    string `json:"major" validate:"required"`
	Faculty  string `json:"faculty" validate:"required"`
	Semester uint   `json:"semester" validate:"required,number"`
	Contact  string `json:"contact" validate:"required,e164,min=10,max=13"`
}
