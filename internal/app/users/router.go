package users

import "github.com/go-chi/chi/v5"

func SetupRoutes(s Service) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getHandler(s))
	router.Post("/", postHandler(s))
	return router
}
