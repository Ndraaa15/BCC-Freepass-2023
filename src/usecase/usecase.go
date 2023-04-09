package usecase

import "bcc-freepass-2023/src/repository"

type Usecase struct {
	Student StudentInterface
}

func Init(repo *repository.Repository) *Usecase {
	return &Usecase{
		Student: InitStudent(repo.Student),
	}
}
