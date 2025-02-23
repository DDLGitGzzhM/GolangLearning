package main

import (
	"GolangLearning/webook/backend/internal"
	"GolangLearning/webook/backend/pkg"
)

func main() {
	server := internal.RegisterRoute()
	pkg.PanicIf(server.Run(":8080"))
}
