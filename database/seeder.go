package database

import (
	"database/sql"
	"fmt"
	"go-rest-api/database/seeders"
)

func RunSeeders(db *sql.DB) error {
	if err := seeders.SeedUsers(db); err != nil {
		return fmt.Errorf("❌ Seederni bajarishda xatolik: %v", err)
	}
	return nil
}
