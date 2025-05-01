package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(path string) {
	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatalf("❌ .env faylni yuklab bo‘lmadi: %v", err)
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")

	dbDriver := os.Getenv("DB_CONNECTION")
	if dbDriver != "pgsql" {
		log.Fatalf("❌ Faqat 'pgsql' qo‘llab-quvvatlanadi, topildi: %s", dbDriver)
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, sslMode,
	)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ DB ulanishda xatolik: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("❌ DB ping ishlamadi: %v", err)
	}

	log.Println("✅ Postgres bazasiga ulanish muvaffaqiyatli!")
}
