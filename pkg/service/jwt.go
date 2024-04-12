package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// TemperveryTokenForOtpVerification generates a JWT token for OTP verification.
// It takes a security key and arbitrary data (Laptop) as input.
// Returns the JWT token string and an error if token creation fails.
func TemperveryTokenForOtpVerification(securityKey string, Laptop string) (string, error) {
	key := []byte(securityKey)
	claims := jwt.MapClaims{
		"exp":    time.Now().Unix() + 3000, // Adjust expiration time as needed
		"Laptop": Laptop,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err // Return the error instead of printing it
	}
	return tokenString, nil // Return tokenString and nil error on success
}

//  Generates an access token for authentication purposes with a specific expiration time.
//  This function creates an access token with an expiration (exp) claim set to 300 seconds from the current time and includes the user's ID (id). It signs the token using the HS256 signing method.

func GenerateAcessToken(securityKey string, id string) (string, error) {
	key := []byte(securityKey)
	claims := jwt.MapClaims{
		"exp": time.Now().Unix() + 300, // Adjust expiration time as needed
		"id":  id,                      // Include the user ID in the token claims
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Create a new token with specified claims and signing method
	tokenString, err := token.SignedString(key)                // Sign the token using the security key
	if err != nil {
		return "", err // Return the error instead of printing it
	}
	return tokenString, nil // Return tokenString and nil error on success
}

//    Generates a refresh token for obtaining new access tokens.
//   This function creates a refresh token with a long expiration time (3600000 seconds) and no additional claims. It signs the token using the HS256 signing method.

func GenerateRefreshToken(securityKey string) (string, error) {
	key := []byte(securityKey)
	claims := jwt.MapClaims{
		"exp": time.Now().Unix() + 3600000, // Set expiration time to 1 hour from now
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", errors.New("failed to generate refresh token: signing error")
	}
	return signedToken, nil
}

//   Verifies and extracts information from an access token.
//   his function parses and verifies the access token using the provided secret key. It extracts the user ID from the token's claims (id) if the token is valid and not tampered with.

func VerifyAccessToken(token string, secretKey string) (string, error) {
	key := []byte(secretKey)
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return "", err // Return the parsing error directly
	}

	if len(parsedToken.Header) == 0 {
		return "", errors.New("token tampered: header is missing or invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to parse token claims")
	}

	id, idOk := claims["id"].(string)
	if !idOk {
		return "", errors.New("id claim not found or not a string")
	}

	return id, nil // Return the extracted user ID and nil error on success
}

//    Verifies the validity of a refresh token.
//    This function parses and verifies the refresh token using the provided security key. It checks for token validity and expiration.

func VerifyRefreshToken(token string, securityKey string) error {
	key := []byte(securityKey)

	// Parse the JWT token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		// Error occurred during token parsing
		return errors.New("token tampered or expired")
	}

	// Check if the token is valid
	if !parsedToken.Valid {
		return errors.New("token is not valid")
	}

	// Token verification successful
	return nil
}

//    Retrieves data (Laptops) from a JWT token.
//    : This function parses and verifies the JWT token using the provided secret key. It extracts the Laptop data from the token's claims if the token is valid and not expired.

func FetchPhoneFromToken(tokenString string, secretKey string) (string, error) {
	secret := []byte(secretKey)

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !parsedToken.Valid {
		return "", errors.New("token validation failed: token is expired or invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("could not parse token claims")
	}

	LaptopClaim, ok := claims["Laptop"].(string)
	if !ok {
		return "", errors.New("phone claim not found or is not a string")
	}

	return LaptopClaim, nil
}
