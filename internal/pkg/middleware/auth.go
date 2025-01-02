package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/purisaurabh/car-rental/internal/pkg/constants"
	specs "github.com/purisaurabh/car-rental/internal/pkg/responses"
	"go.uber.org/zap"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			zap.S().Error("No token found")
			specs.HandleError(w, http.StatusUnauthorized, "No token found", nil)
			return
		}

		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 || splitToken[0] == "Bearer" {
			zap.S().Error("Invalid token format")
			specs.HandleError(w, http.StatusUnauthorized, "Invalid token format", nil)
			return
		}

		tokenString := splitToken[1]
		claims, err := VerifyJWTToken(tokenString)
		if err != nil {
			zap.S().Error("Error verifying token", err)
			specs.HandleError(w, http.StatusUnauthorized, "Error verifying token", nil)
			return
		}

		userID, ok := claims["user_id"]
		if !ok {
			zap.S().Error("Error getting user_id from token")
			specs.HandleError(w, http.StatusUnauthorized, "Error getting user_id from token", nil)
			return
		}

		userRole, ok := claims["role"]
		if !ok {
			zap.S().Error("Error getting role from token")
			specs.HandleError(w, http.StatusUnauthorized, "Error getting role from token", nil)
			return
		}

		ctx := context.WithValue(r.Context(), constants.UserIDKey, userID)
		ctx = context.WithValue(ctx, constants.UserRoleKey, userRole)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
