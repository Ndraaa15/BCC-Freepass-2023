package usecase

import "bcc-freepass-2023/src/repository"

type StudentInterface interface {
	Register()
	Login()
}

type Student struct {
	studentRepo repository.StudentInterface
}

func InitStudent(studentRepo repository.StudentInterface) StudentInterface {
	student := Student{}
	student.studentRepo = studentRepo
	return &student
}

func (uc *Student) Register() {
	uc.studentRepo.CreateStudent()
}

func (uc *Student) Login() {

}
