package main

import (
	"github.com/cyee02/ecommerce-service/db"
	"github.com/cyee02/ecommerce-service/helper/config"
	"github.com/cyee02/ecommerce-service/helper/middleware"
	"github.com/cyee02/ecommerce-service/routes"
	"github.com/cyee02/ecommerce-service/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Init Config
	config.InitConfig("./config")

	// Init clients
	db.InitMySQL()
	service.InitUserService()

	routes.GenRoutePathWithoutAuth(router) // Init routes without requirement to authenticate
	router.Use(middleware.AuthToken)       // Init middleware for auth
	routes.GenRoutePathWithAuth(router)    // Init routes with requirement to auth authenticate
	router.Run("localhost:8080")
}
