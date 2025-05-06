package bootstrap

import (
	"flag"
	"go-rest-api/config"
	"go-rest-api/database"
	"go-rest-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	client := router.Group("/api/client")
	admin := router.Group("/api/admin")
	auth := router.Group("/api/auth")
	routes.ClientRoutes(client)
	routes.AdminRoutes(admin)
	routes.AuthRoutes(auth)

	return router
}

func SetupDB() {
	config.InitDB(".")
}

func SetupCommands() {
	migrateFlag := flag.Bool("migrate", false, "🏃 Run migrations")
	freshFlag := flag.Bool("fresh", false, "🏃 Run fresh DB")
	seederFlag := flag.Bool("seed", false, "🏃 Run seeders")

	flag.Parse()

	// start -> run fresh DB
	if *freshFlag {
		errDown := database.RunMigrationDown(config.DB)
		if errDown != nil {
			log.Fatal("Migratsiya bajarishda xatolik:", errDown)
		}
	}
	// end -> run fresh DB

	// start -> run migrations
	if *migrateFlag {
		err := database.RunMigrationUp(config.DB)
		if err != nil {
			log.Fatal("Migratsiya bajarishda xatolik:", err)
		}
	}
	// end -> run migrations

	// start -> run seeders
	if *seederFlag {
		err := database.RunSeeders(config.DB)
		if err != nil {
			log.Fatal("Seeder bajarishda xatolik:", err)
		}
	}
	// end -> run seeders
}
