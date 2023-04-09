package supabase

import (
	"os"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

func InitSupabase() *supabasestorageuploader.SupabaseClientService {
	supClient := supabasestorageuploader.NewSupabaseClient(
		os.Getenv("PROJECT_URL"),
		os.Getenv("PROJECT_API_KEYS"),
		os.Getenv("STORAGE_NAME"),
		os.Getenv("STORAGE_FOLDER"),
	)
	return &supClient
}
