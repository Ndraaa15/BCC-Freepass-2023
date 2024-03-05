package usecase

import "bcc-freepass-2023/internal/repository"

//Struct Usecase berfungsi untuk menyimpan interface (berisi function) dari dari setiap entitiy yang ada
type Usecase struct {
	Student StudentInterface
}

/*
Method Init berfungsi untuk melakukan melakukan injection "Repository" ke "Usecase"
Hal ini dilakukan agar function yang ada di "Repository" bisa di akses di "Usecase"
*/
func Init(repo *repository.Repository) *Usecase {
	return &Usecase{
		Student: InitStudent(repo.Student),
	}
}
