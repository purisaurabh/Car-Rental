package api

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/purisaurabh/car-rental/internal/app"
)

func Routes(ctx context.Context, svc app.Service) *mux.Router {
	router := mux.NewRouter()
	return router
}
