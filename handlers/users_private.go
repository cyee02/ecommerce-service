package handlers

import (
	"errors"
	"net/http"

	"github.com/cyee02/ecommerce-service/helper"
	"github.com/cyee02/ecommerce-service/models"
	"github.com/cyee02/ecommerce-service/service"
	"github.com/gin-gonic/gin"
)

func GetCurrUser(c *gin.Context) {
	// Get current user information based on SignedToken value
	val, ok := c.Get("currentUser")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, "Error getting current user information")
	}
	contextUser, err := helper.CastAnyToUser(val)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Error casting value to User struct")
		return
	}

	// Call db to find user_id
	currentUser, err := service.UserService.GetUserById(contextUser.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, currentUser)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user to be updated belongs to current user
	if err := validateUserAuth(c, user.UserId); err != nil {
		return
	}

	// Update current user
	resp, err := service.UserService.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user to be updated belongs to current user
	if err := validateUserAuth(c, user.UserId); err != nil {
		return
	}

	// Delete current user
	resp, err := service.UserService.DeleteUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, resp)
}

func validateUserAuth(c *gin.Context, userId string) error {
	// Get current user information based on SignedToken
	val, ok := c.Get("currentUser")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, "Error getting current user information")
		return errors.New("Error getting current user information")
	}
	currentUser, err := helper.CastAnyToUser(val)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Error casting value to User struct")
		return errors.New("Error casting value to User struct")
	}

	// Check if current user is same as request
	if currentUser.UserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "User to be modified does not match with current user"})
		return errors.New("User to be updated does not match with current user")
	}
	return nil
}
