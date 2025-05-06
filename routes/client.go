package routes

import (
	"go-rest-api/internal/controllers/client"
	"go-rest-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ClientRoutes(rg *gin.RouterGroup) {

	profile := rg.Group("profile").Use(middleware.AuthMiddleware())
	{
		profile.GET("/", client.GetUsers)
		profile.PUT("/update", client.GetUsers).Use()
	}
}
