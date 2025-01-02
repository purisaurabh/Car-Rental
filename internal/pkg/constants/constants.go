package constants

import (
	"net/http"

	"github.com/rs/cors"
)

const (
	EmailRegex    = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
	MobileRegex   = "^([+]\\d{2})?\\d{10}$"
	PasswordRegex = `^[A-Za-z\d@$!%*#?&]{8,}$`
)

var CorsOptions = cors.Options{
	AllowedOrigins:   []string{"*"},
	AllowCredentials: true,
	AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodPatch},
	AllowedHeaders:   []string{"*"},
}

type ContextKey string

const (
	UserIDKey    ContextKey = "user_id"
	UserRoleKey  ContextKey = "role"
)