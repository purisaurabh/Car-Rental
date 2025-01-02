package api

import (
	"context"
	"net/http"

	"github.com/purisaurabh/car-rental/internal/api/decoder"
	"github.com/purisaurabh/car-rental/internal/app/user"
	specs "github.com/purisaurabh/car-rental/internal/pkg/responses"
)

func UserRegistration(ctx context.Context, userService user.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode request
		req, err := decoder.DecodeUserRegistrationRequest(r)
		if err != nil {
			specs.HandleError(w, http.StatusBadRequest, "Failed to decode registration request", err)
			return
		}

		// Validate request
		if err := req.Validate(); err != nil {
			specs.HandleError(w, http.StatusBadRequest, "Invalid registration request", err)
			return
		}

		// Register user
		userID, err := userService.UserRegistration(ctx, &req)
		if err != nil {
			specs.HandleError(w, http.StatusInternalServerError, "Failed to register user", err)
			return
		}

		// Respond with success
		specs.SuccessResponse(w, http.StatusCreated, specs.MessageResponseWithID{
			ID:      userID,
			Message: "User registered successfully",
		})
	}
}

func UserLogin(ctx context.Context, userService user.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode request
		req, err := decoder.DecodeUserLoginRequest(r)
		if err != nil {
			specs.HandleError(w, http.StatusBadRequest, "Failed to decode login request", err)
			return
		}

		// Validate request
		if err := req.Validate(); err != nil {
			specs.HandleError(w, http.StatusBadRequest, "Invalid login request", err)
			return
		}

		// Login user
		userLoginResponse, err := userService.UserLogin(ctx, &req)
		if err != nil {
			specs.HandleError(w, http.StatusInternalServerError, "Failed to login user", err)
			return
		}

		// Respond with success
		specs.SuccessResponse(w, http.StatusOK, userLoginResponse)
	}
}
