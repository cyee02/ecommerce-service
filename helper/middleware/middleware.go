package middleware

import (
	"net/http"
	"time"

	"github.com/cyee02/ecommerce-service/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SignedDetails struct {
	User *models.User
	jwt.StandardClaims
}

const SECRET_KEY = "my_secret_key"

func AuthToken(c *gin.Context) {
	// get JWT from auth header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
		c.Abort()
		return
	}

	// Parse the JWT
	token, err := jwt.ParseWithClaims(tokenString, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
		c.Abort()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is Expired"})
		c.Abort()
		return
	}
	c.Set("currentUser", claims.User)
}

func GenToken(user *models.User) (*string, error) {
	claims := &SignedDetails{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, err
	}
	return &token, nil
}
