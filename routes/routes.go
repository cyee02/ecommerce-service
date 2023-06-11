package routes

import (
	"github.com/cyee02/ecommerce-service/handlers"
	"github.com/gin-gonic/gin"
)

func GenRoutePathWithAuth(router *gin.Engine) {
	// User management
	router.POST("users/myinfo", handlers.GetCurrUser)
	router.POST("users/update", handlers.UpdateUser)
	router.POST("users/delete", handlers.DeleteUser)
}

func GenRoutePathWithoutAuth(router *gin.Engine) {
	// User management
	router.POST("users/search", handlers.GetUsers)
	router.POST("users/signup", handlers.CreateUser)
	router.POST("users/login", handlers.Login)
}
