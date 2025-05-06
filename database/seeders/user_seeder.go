package seeders

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers(db *sql.DB) error {
	adminPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("admin parolini hashlashda xatolik: %v", err)
	}

	userPassword, err := bcrypt.GenerateFromPassword([]byte("user123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("user parolini hashlashda xatolik: %v", err)
	}

	query := `
		INSERT INTO users (name, email, password) VALUES
		('president', 'admin@example.com', $1),
		('tuychiyev', 'user@example.com', $2);
	`

	_, err = db.Exec(query, string(adminPassword), string(userPassword))
	if err != nil {
		return fmt.Errorf("users jadvaliga seeder yozishda xatolik: %v", err)
	}

	log.Println("✅ Users jadvaliga seeder yozildi")
	return nil
}
