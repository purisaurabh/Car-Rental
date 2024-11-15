package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/purisaurabh/car-rental/internal/app"
	"github.com/purisaurabh/car-rental/internal/pkg/errors"
	specs "github.com/purisaurabh/car-rental/internal/pkg/responses"
	"go.uber.org/zap"
)

func UserRegistration(ctx context.Context, userService app.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeUserRegistrationRequest(r)
		if err != nil {
			specs.ErrorResponse(w, http.StatusBadRequest, err)
			zap.S().Errorf("Error decoding request: %v", err)
			return
		}

		err = req.Validate()
		if err != nil {
			specs.ErrorResponse(w, http.StatusBadRequest, err)
			zap.S().Errorf("Error validating request: %v", err)
			return
		}

		userID, err := userService.UserRegistration(ctx, &req)
		if err != nil {
			specs.ErrorResponse(w, http.StatusInternalServerError, errors.ErrInternalServer)
			zap.S().Errorf("Error registering user: %v", err)
			return
		}

		specs.SuccessResponse(w, http.StatusCreated, specs.MessageResponseWithUserID{
			UserId:  userID,
			Message: "User registered successfully",
		})

	}
}

func UserLogin(ctx context.Context, userService app.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeUserLoginRequest(r)
		if err != nil {
			specs.ErrorResponse(w, http.StatusBadRequest, err)
			zap.S().Errorf("Error decoding request: %v", err)
			return
		}

		err = req.Validate()
		if err != nil {
			specs.ErrorResponse(w, http.StatusBadRequest, err)
			zap.S().Errorf("Error validating request: %v", err)
			return
		}

		userLoginResponse, err := userService.UserLogin(ctx, &req)
		if err != nil {
			fmt.Println("error while loggin is : ", err)
			specs.ErrorResponse(w, http.StatusInternalServerError, errors.ErrInternalServer)
			zap.S().Errorf("Error logging in user: %v", err)
			return
		}

		specs.SuccessResponse(w, http.StatusOK, userLoginResponse)
	}
}
