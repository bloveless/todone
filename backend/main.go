package main

import (
	"todone-backend/controllers"
	"todone-backend/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Add the cors headers if there is an Origin header.
	r.Use(middleware.CORSMiddleware())

	homeController := new(controllers.HomeController)
	homeController.RegisterRoutes(r)

	pingController := new(controllers.PingController)
	pingController.RegisterRoutes(r)

	r.Run(":9090")
	// r.RunTLS(":4443", "keys/server.crt", "keys/server.key")
}
