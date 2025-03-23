package main

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"GolangLearning/webook/backend/internal/web"
	"GolangLearning/webook/backend/pkg"
)

func main() {
	server := initRestServer()
	web.RegisterUserRoutes(server)
	pkg.PanicIf(server.Run(":8080"))
}

func initRestServer() *gin.Engine {
	server := gin.Default()
	// 确保 CORS 中间件在最前面
	server.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"content-type", "authorization"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return false
		},
		MaxAge: 100 * time.Millisecond,
	}))
	return server
}
