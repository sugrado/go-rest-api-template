package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sugrado/tama-server/internal/app"
	"github.com/sugrado/tama-server/internal/app/users"
)

func New(s *app.Service) *chi.Mux {
	router := chi.NewRouter()
	setMiddlewares(router)

	router.Mount("/api/users", users.SetupRoutes(s.User()))
	return router
}
