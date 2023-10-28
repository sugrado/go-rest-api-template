package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sugrado/go-rest-api-template/internal/app"
	"github.com/sugrado/go-rest-api-template/internal/app/users"
)

func New(s *app.Service) *chi.Mux {
	router := chi.NewRouter()
	setMiddlewares(router)

	router.Mount("/api/users", users.SetupRoutes(s.User()))
	router.Mount("/api/debug", middleware.Profiler()) //pprof
	return router
}

func setMiddlewares(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // yol geçen hanı
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//r.Use(cache.New())
}
