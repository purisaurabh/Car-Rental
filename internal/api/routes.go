package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/purisaurabh/car-rental/internal/app"
)

func Routes(ctx context.Context, svc app.Service) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/registration", UserRegistration(ctx, svc)).Methods(http.MethodPost)
	router.HandleFunc("/login", UserLogin(ctx, svc)).Methods(http.MethodPost)
	return router
}
