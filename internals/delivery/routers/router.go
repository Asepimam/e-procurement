package routers

import (
	"e-procurement/internals/delivery/https"
	"e-procurement/internals/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	User    string
	Auth    usecases.AuthUseCase
	Vendor  string
	Product string
}

// routes

func registerAuthRoutes(r chi.Router, authHandler *https.AuthHttp) {
	r.Post("/auth/login", authHandler.Authentication)
	r.Post("/auth/register", authHandler.Create)
}

func NewRouter(r *Router) http.Handler {
	router := chi.NewRouter()

	authHandler := https.NewAuthHttp(r.Auth)

	router.Route("/api/v1/", func(r chi.Router) {
		r.Get("/hallo", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})

		registerAuthRoutes(r, authHandler)
	})
	return router
}