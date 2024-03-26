package usecase

import "bcc-freepass-2023/internal/repository"

type IStudentUsecase interface {
}

type StudentUsecase struct {
	studentQuerier *repository.Queries
}

func NewStudentUsecase(studentQuerier *repository.Queries) IStudentUsecase {
	return &StudentUsecase{
		studentQuerier: studentQuerier,
	}
}

func (s *StudentUsecase) CreateStudent() {

}
