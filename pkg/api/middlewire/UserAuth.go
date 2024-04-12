package middlewire

import (
	"Laptop_Lounge/pkg/config"
	"Laptop_Lounge/pkg/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type TokenRequirement struct {
	tokenSecurityKey config.Token
}

var token TokenRequirement



func NewJwtTokenMiddleWire(keys config.Token) {
	
	token = TokenRequirement{tokenSecurityKey: keys}
}

func UserAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Fetch the secret key securely from environment variables or config file
		return []byte("comebuyLaptops"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to extract token claims"})
		c.Abort()
		return
	}

	role, ok := claims["role"].(string)
	if !ok || role != "client" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}

	id, ok := claims["id"].(float64)
	if !ok || id == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Error in retrieving ID"})
		c.Abort()
		return
	}

	c.Set("userRole", role)
	c.Set("userID", int(id))

	fmt.Println("User authenticated:", role, id)

	c.Next()
}

func AdminAuthorization(c *gin.Context) {
	adminToken := c.GetHeader("Authorization")
	fmt.Println("--", token)

	err := service.VerifyRefreshToken(adminToken, token.tokenSecurityKey.AdminSecurityKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		c.Abort()
	}
	c.Next()
}
