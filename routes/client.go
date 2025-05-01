package routes

import (
	"go-rest-api/internal/controllers/client"

	"github.com/gin-gonic/gin"
)

func ClientRoutes(rg *gin.RouterGroup) {

	profile := rg.Group("profile")
	{
		profile.GET("/", client.GetUsers)
	}
}
