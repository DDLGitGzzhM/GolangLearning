package internal

import "github.com/gin-gonic/gin"

func RegisterRoute() *gin.Engine {
	sever := gin.Default()
	registerUserRoutes(sever)
	return sever
}

func registerUserRoutes(server *gin.Engine) {
	UserRoute.RegisterRoutes(server)
}
