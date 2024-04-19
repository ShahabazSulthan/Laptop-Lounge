package middlewire

import (
	"Laptop_Lounge/pkg/config"
	"Laptop_Lounge/pkg/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequirement struct {
	tokenSecurityKey config.Token
}

var token TokenRequirement

func NewJwtTokenMiddleWire(keys config.Token) {
	token = TokenRequirement{tokenSecurityKey: keys}
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

// func SellerAuthorization(c *gin.Context) {
// 	accessToken := c.Request.Header.Get("Authorization")
// 	refreshToken := c.Request.Header.Get("Refreshtoken")

// 	if accessToken == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing access token"})
// 		c.Abort()
// 		return
// 	}

// 	id, err := service.VerifyAccessToken(accessToken, token.tokenSecurityKey.SellerSecurityKey)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
// 		c.Abort()
// 		return
// 	}

// 	if id == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to extract ID from token"})
// 		c.Abort()
// 		return
// 	}

// 	fmt.Println("Seller ID:", id)

// 	// Check if access token is valid, generate a new one if not
// 	if err := service.VerifyRefreshToken(refreshToken, token.tokenSecurityKey.SellerSecurityKey); err != nil {
// 		newAccessToken, err := service.GenerateAcessToken(token.tokenSecurityKey.SellerSecurityKey, id)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to generate new access token"})
// 			c.Abort()
// 			return
// 		}
// 		fmt.Println("New Access Token:", newAccessToken)
// 		// Update the Authorization header with the new access token
// 		c.Writer.Header().Set("Authorization", newAccessToken)
// 	}

// 	// Set SellerID in context
// 	c.Set("SellerID", id)
// 	c.Next()
// }

func SellerAuthorization(c *gin.Context) {
	accessToken := c.Request.Header.Get("Authorization")
	refreshToken := c.Request.Header.Get("Refreshtoken")

	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing access token"})
		c.Abort()
		return
	}

	id, err := service.VerifyAccessToken(accessToken, token.tokenSecurityKey.SellerSecurityKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
		c.Abort()
		return
	}

	if id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to extract ID from token"})
		c.Abort()
		return
	}

	fmt.Println("Seller ID:", id)

	// Check if access token is valid, generate a new one if not
	if err := service.VerifyRefreshToken(refreshToken, token.tokenSecurityKey.SellerSecurityKey); err != nil {
		newAccessToken, err := service.GenerateAcessToken(token.tokenSecurityKey.SellerSecurityKey, id)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to generate new access token"})
			c.Abort()
			return
		}
		fmt.Println("New Access Token:", newAccessToken)
		// Update the Authorization header with the new access token
		c.Writer.Header().Set("Authorization", newAccessToken)
	}

	// Set SellerID in context
	c.Set("SellerID", id)
	c.Next()
}

func UserAuthorization(c *gin.Context) {
	accessToken := c.Request.Header.Get("authorization")
	refreshToken := c.Request.Header.Get("refreshtoken")

	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "There is no access token"})
		c.Abort()
		return
	}

	id, err := service.VerifyAccessToken(accessToken, token.tokenSecurityKey.UsersSecurityKey)
	if err != nil {
		err := service.VerifyRefreshToken(refreshToken, token.tokenSecurityKey.UsersSecurityKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
			c.Abort()
		} else {
			newAccessToken, err := service.GenerateAcessToken(token.tokenSecurityKey.UsersSecurityKey, id)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
				c.Abort()
			} else {
				fmt.Println("New Access Token:", newAccessToken)
				c.Set("UserID", id)
				c.Next()
			}
		}
	} else {
		// c.JSON(http.StatusOK, "All perfect, your access token is up-to-date")
		c.Set("UserID", id)
		c.Next()
	}
}
