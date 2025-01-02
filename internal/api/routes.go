package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/purisaurabh/car-rental/internal/app"
	"github.com/purisaurabh/car-rental/internal/pkg/middleware"
)

func Routes(ctx context.Context, deps app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	// Public routes (no auth required)
	publicRouter := router.NewRoute().Subrouter()
	publicRouter.HandleFunc("/registration", UserRegistration(ctx, deps.UserService)).Methods(http.MethodPost)
	publicRouter.HandleFunc("/login", UserLogin(ctx, deps.UserService)).Methods(http.MethodPost)

	// Protected routes (auth required)
	privateRouter := router.NewRoute().Subrouter()
	privateRouter.Use(middleware.AuthMiddleware)
	privateRouter.HandleFunc("/car", CreateCar(ctx, deps.CarService)).Methods(http.MethodPost)

	return router
}
