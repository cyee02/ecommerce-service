package handlers

import (
	"net/http"

	"github.com/cyee02/ecommerce-service/db"
	"github.com/cyee02/ecommerce-service/helper"
	"github.com/cyee02/ecommerce-service/models"
	"github.com/cyee02/ecommerce-service/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := service.UserService.Login(&user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetUsers(c *gin.Context) {
	users, err := db.UserDB.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	getUserResp := helper.ConvertUserArrayToPublic(users)
	c.IndentedJSON(http.StatusOK, getUserResp)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := service.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, resp)
}
