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

// user registration columns
var UserRegistrationColumns = []string{"name", "email", "password", "mobile", "role", "created_at", "updated_at"}

// user login columns
var UserLoginColumns = []string{"id", "name", "password", "mobile", "role"}

const (
	UsersTable = "users"
)
