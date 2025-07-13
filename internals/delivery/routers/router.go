package routers

import (
	"e-procurement/internals/delivery/https"
	"e-procurement/internals/usecases"
	"e-procurement/pkg/auth"
	"e-procurement/pkg/chi_middlewar"

	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	User    usecases.UserUseCase
	Auth    usecases.AuthUseCase
	Vendor  usecases.VendorUseCase
	Product usecases.ProductUseCase
	Category usecases.CategoryUsecase
	JWT *auth.JWT
}

// routes

func registerAuthRoutes(r chi.Router, authHandler *https.AuthHttp) {
	r.Post("/auth/login", authHandler.Authentication)
	r.Post("/auth/register", authHandler.Create)
}

func registerUserRoutes(r chi.Router, userHandler *https.UserHttp) {
	r.Get("/user", userHandler.GetUserByID)
	r.Put("/user", userHandler.UpdateUser)
	r.Delete("/user", userHandler.DeleteUser)
	r.Put("/user/password", userHandler.ChangePassword)
}

func registerProductRoutes(r chi.Router, productHandler *https.ProductHttp) {
	r.Post("/vendor/product",productHandler.CreateProduct)
	r.Get("/vendor/products", productHandler.GetAllProducts)
	r.Get("/vendor/products/{id}", productHandler.GetProductByID)
	r.Put("/vendor/products/{id}", productHandler.UpdateProduct)
	r.Delete("/vendor/products/{id}", productHandler.DeleteProduct)
	r.Get("/vendor/products/category/{categoryID}", productHandler.GetProductsByCategory)
}

func registerCategoryRoutes(r chi.Router, categoryHandler *https.CategoryHttp) {
	r.Post("/vendor/product_category", categoryHandler.CreateCategory)
	r.Get("/vendor/product_categories", categoryHandler.GetCategory)
	r.Get("/vendor/product_categories/{id}", categoryHandler.GetCategoryByID)
	r.Put("/vendor/product_categories/{id}", categoryHandler.UpdateCategory)
	r.Delete("/vendor/product_categories/{id}", categoryHandler.DeleteCategory)
}

func registerVendorRouters(r chi.Router, vendorHandler *https.VendorHttp) {
	r.Post("/vendor", vendorHandler.CreateVendor)
	r.Get("/vendor", vendorHandler.GetAllVendors)
	r.Get("/vendor/{id}", vendorHandler.GetVendorByID)
	r.Put("/vendor/{id}", vendorHandler.UpdateVendor)
	r.Delete("/vendor/{id}", vendorHandler.DeleteVendor)
}

func NewRouter(r *Router) http.Handler {
	router := chi.NewRouter()
	// setting body is json by default
	router.Use(chi_middlewar.JSONContentTypeMiddleware)
	// Middleware can be added here if needed
	jwtMiddleware := auth.NewAuthMiddleware(r.JWT)

	authHandler := https.NewAuthHttp(r.Auth)
	userHandler := https.NewUserHttp(r.User)
	productHandler := https.NewProductHttp(r.Product)
	categoryHandler := https.NewCategoryHttp(r.Category)
	vendorHandler := https.NewVendortHttp(r.Vendor)
	router.Route("/api/v1/", func(r chi.Router) {
		// public routes
		r.Get("/hallo", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})
		registerAuthRoutes(r, authHandler)

		// protected routes
		r.Group(func(protected chi.Router) {
			protected.Use(jwtMiddleware.VerifyToken)
			registerUserRoutes(protected, userHandler)
			registerProductRoutes(protected,productHandler)
			registerCategoryRoutes(protected, categoryHandler)
			registerVendorRouters(protected, vendorHandler)
		})
	})
	return router
}