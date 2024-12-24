package middleware

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/purisaurabh/car-rental/internal/config"
	"github.com/purisaurabh/car-rental/internal/pkg/specs"
)

func createClaims(payload specs.TokenPayload, expiration time.Duration) jwt.MapClaims {
	return jwt.MapClaims{
		"authorized": true,
		"user_id":    payload.UserID,
		"email":      payload.Email,
		"role":       payload.Role,
		"exp":        time.Now().Add(expiration).Unix(),
	}
}

func CreateToken(payload specs.TokenPayload) (string, error) {
	secretKey := config.GetSecretKey()
	if secretKey == "" {
		return "", fmt.Errorf("secret key not found")
	}

	expirationTimeInHours := config.GetExpiryTime()
	if expirationTimeInHours == "" {
		return "", fmt.Errorf("expiration time not found")
	}

	expirationHours, err := strconv.Atoi(expirationTimeInHours)
	if err != nil {
		log.Fatalf("Error parsing TOKEN_EXPIRATION_HOURS: %v", err)
		return "", err
	}

	claims := createClaims(payload, time.Duration(expirationHours)*time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("error while signing the token: %v", err)
	}

	return tokenString, nil

}
