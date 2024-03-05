package repository

import (
	"database/sql"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type Repository struct {
	Student StudentInterface
}

func Init(db sql.DB, supabase supabasestorageuploader.SupabaseClientService) *Repository {
	return &Repository{
		Student: InitStudent(db, supabase),
	}
}
