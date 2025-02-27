package main

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"

	"GolangLearning/webook/backend/internal"
	"GolangLearning/webook/backend/pkg"
)

func main() {
	server := internal.RegisterRoute()

	// 确保 CORS 中间件在最前面
	server.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"content-type", "authorization"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return false
		},
		MaxAge: 100 * time.Millisecond,
	}))

	pkg.PanicIf(server.Run(":8080"))
}
