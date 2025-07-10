package routers

import (
	"e-procurement/internals/delivery/https"
	"e-procurement/internals/usecases"
	"e-procurement/pkg/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	User    string
	Auth    usecases.AuthUseCase
	Vendor  string
	Product usecases.ProductUseCase
	Category usecases.CategoryUsecase
	JWT *auth.JWT
}

// routes

func registerAuthRoutes(r chi.Router, authHandler *https.AuthHttp) {
	r.Post("/auth/login", authHandler.Authentication)
	r.Post("/auth/register", authHandler.Create)
}

func registerProductRoutes(r chi.Router, productHandler *https.ProductHttp) {
	r.Post("/vendor/product",productHandler.CreateProduct)
}

func registerCategoryRoutes(r chi.Router, categoryHandler *https.CategoryHttp) {
	r.Post("/vendor/product_category", categoryHandler.CreateCategory)
}

func NewRouter(r *Router) http.Handler {
	router := chi.NewRouter()
	// Middleware can be added here if needed
	jwtMiddleware := auth.NewAuthMiddleware(r.JWT)

	authHandler := https.NewAuthHttp(r.Auth)
	productHandler := https.NewProductHttp(r.Product)
	categoryHandler := https.NewCategoryHttp(r.Category)
	router.Route("/api/v1/", func(r chi.Router) {
		// public routes
		r.Get("/hallo", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})
		registerAuthRoutes(r, authHandler)

		// protected routes
		r.Group(func(protected chi.Router) {
			protected.Use(jwtMiddleware.VerifyToken)
			registerProductRoutes(r,productHandler)
			registerCategoryRoutes(r, categoryHandler)
		})
	})
	return router
}