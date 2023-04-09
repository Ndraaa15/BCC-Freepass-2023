package repository

import (
	"database/sql"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type StudentInterface interface {
	CreateStudent()
	GetByEmail()
}

type Student struct {
	db       sql.DB
	supabase supabasestorageuploader.SupabaseClientService
}

func InitStudent(db sql.DB, supabase supabasestorageuploader.SupabaseClientService) StudentInterface {
	student := Student{}
	student.db = db
	student.supabase = supabase
	return &student
}

func (r *Student) CreateStudent() {

}

func (r *Student) GetByEmail() {

}
