package main

import (
	"go-rest-api/internal/bootstrap"
)

func main() {
	bootstrap.SetupDB()
	bootstrap.SetupCommands()
	app := bootstrap.SetupRouter()
	app.Run(":8080")
}
