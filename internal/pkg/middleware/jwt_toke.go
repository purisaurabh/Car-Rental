package middleware

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/purisaurabh/car-rental/internal/config"
	"github.com/purisaurabh/car-rental/internal/pkg/errors"
	"github.com/purisaurabh/car-rental/internal/pkg/specs"
	"go.uber.org/zap"
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

func VerifyJWTToken(tokenString string) (jwt.MapClaims, error) {
	if tokenString == "" {
		return nil, errors.ErrTokenEmpty
	}
	secretKey := config.GetSecretKey()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrSigningMethod
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		zap.S().Error("Error in parsing token: %v", err)
		return nil, errors.ErrInvalidToken
	}

	if !token.Valid {
		zap.S().Error("Token is not valid")
		return nil, errors.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims == nil {
		zap.S().Error("Error in parsing claims")
		return nil, errors.ErrInvalidToken
	}

	return claims, nil
}
