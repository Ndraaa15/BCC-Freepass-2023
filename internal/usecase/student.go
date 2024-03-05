package usecase

import "bcc-freepass-2023/internal/repository"

/*
Dalam bahasa Go, sebuah tipe struct dapat dianggap sebagai sebuah implementasi
dari sebuah interface jika struktur tersebut memiliki semua method yang diperlukan oleh interface tersebut.
*/

/*
Example :
			func (uc *Student) doSomething() {
			...
			}
Fungsi di atas merupakan salah satu contoh function yang memiliki implementasi dari sebuah struct
Key : func (nameVar *StructName) functionName (return) {...}
*/

type StudentInterface interface {
	Register()
	Login()
}

type Student struct {
	studentRepo repository.StudentInterface
}

/*
Meskipun dalam permintaan return (interface) berbeda dengan apa yang di return funtion (struct)
Hal ini tetap bisa dilakukan apabila struct mengimplemntasikan semua method yang ada di interface
*/
func InitStudent(studentRepo repository.StudentInterface) StudentInterface {
	return &Student{
		studentRepo: studentRepo,
	}
}

func (uc *Student) Register() {
	uc.studentRepo.CreateStudent()
}

func (uc *Student) Login() {

}
