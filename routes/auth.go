package routes

import (
	"go-rest-api/internal/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", auth.Login)
}
