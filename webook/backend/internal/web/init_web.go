package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	UserRoute.RegisterRoutes(server)
}
