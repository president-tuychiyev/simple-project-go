package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

func RunMigrationUp(db *sql.DB) error {
	upPath, _ := filepath.Abs("database/migrations/000001_initial_migrate_up.sql")
	sqlBytes, err := os.ReadFile(upPath)
	if err != nil {
		return fmt.Errorf("❌ Faylni o‘qishda xatolik: %v", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		fmt.Println("❌ Xatolik yuz berdi, rollback (down) bajarilmoqda...")
		errDown := RunMigrationDown(db)
		if errDown != nil {
			return fmt.Errorf("down migratsiyasini bajarishda xatolik: %v", errDown)
		}

		return fmt.Errorf("❌ SQLni bajarishda xatolik: %v, va rollback amalga oshirildi", err)
	}

	log.Println("✅ Migratsiya muvaffaqiyatli bajarildi!")
	return nil
}

func RunMigrationDown(db *sql.DB) error {
	downPath, _ := filepath.Abs("database/migrations/000001_initial_migrate_down.sql")
	sqlBytes, err := os.ReadFile(downPath)
	if err != nil {
		return fmt.Errorf("faylni o‘qishda xatolik: %v", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return fmt.Errorf("SQLni bajarishda xatolik: %v", err)
	}

	log.Println("✅ Migratsiya bekor qilindi!")
	return nil
}
