package seeders

import (
	"database/sql"
	"fmt"
	"log"
)

func SeedUsers(db *sql.DB) error {
	query := `
		INSERT INTO users (name, email, password) VALUES
		('president', 'admin@example.com', 'hashed_password'),
		('tuychiyev', 'user@example.com', 'hashed_password');
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("users jadvaliga seeder yozishda xatolik: %v", err)
	}

	log.Println("✅ Users jadvaliga seeder yozildi")
	return nil
}
