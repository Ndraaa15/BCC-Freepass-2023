package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func Init() (*DB, error) {
	//Pemanggilan method initDatabase untuk mendapatkan koneksi database dan dimasukkan ke dalam struct DB
	db, err := InitDatabase()
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func InitDatabase() (*gorm.DB, error) {
	//Membuat sumber data agar bisa terhubung dengan database
	sourceDatabase := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	//Proses menghubungkan ke database
	db, err := gorm.Open(mysql.Open(sourceDatabase), &gorm.Config{
		//logger.silent artinya tidak menampilkan proses yang sedang berlangsung ke database
		Logger: logger.Default.LogMode(logger.Error),
	})

	//Error handling apabila terjadi kesalahan waktu mencoba terhubung ke database
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	//Mendapatkan koneksi database sehingga dapat memanggil method di bawah
	database, err := db.DB()
	if err != nil {
		return nil, err
	}

	//Jumlah koneksi maksimum yang boleh tidak aktif pada suatu waktu.
	database.SetMaxIdleConns(2)

	//Jumlah koneksi maksimum yang boleh dibuka secara bersamaan
	database.SetMaxOpenConns(5)

	//Waktu maksimum koneksi dapat tetap idle sebelum ditutup secara otomatis
	database.SetConnMaxIdleTime(time.Minute * 10)

	//Waktu maksimum koneksi dapat digunakan sebelum ditutup secara otomatis
	database.SetConnMaxLifetime(time.Hour * 1)

	return db, nil
}
