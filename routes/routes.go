package routes

import (
	"go-crud/handlers"

	"go-crud/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {

	r := chi.NewRouter()
	r.Post("/auth/register", handlers.Register)
	r.Post("/auth/login", handlers.Login)

	r.Route("/", func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware)
		protected.Get("/users/me", handlers.GetMe)

		protected.Route("/", func(admin chi.Router) {
			admin.Use(middleware.AdminMiddleware)
			admin.Post("/stocks/create", handlers.CreateStock)
		})
	})

	return r
}

// r.Post("/users/post-user", handlers.RegisterAUser)

// // pakai chi param
// r.Put("/users/update-user/{id}", handlers.UpdateAUser)
// r.Delete("/users/delete-user/{id}", handlers.DeleteAUser)

// r.Get("/stocks/get-all", handlers.GetAllStock)
// r.Post("/stocks/post-stock", handlers.InsertAStock)
// r.Put("/stocks/update-stock/{id}", handlers.UpdateAStock)
// r.Delete("/stocks/delete-stock/{id}", handlers.DeleteAStock)

// http.HandleFunc("/categories", handlers.GetCategories)
// http.HandleFunc("/category/create", handlers.CreateCategory)

// // http.HandleFunc("/products", handlers.GetProducts)
// http.HandleFunc("/product/create", handlers.CreateProduct)
// // http.HandleFunc("/product/update", handlers.UpdateProduct)
// // http.HandleFunc("/product/delete", handlers.DeleteProduct)

// http.HandleFunc("/departements", handlers.GetDepartements)
// http.HandleFunc("/departements/create", handlers.CreateDepartement)

// // http.HandleFunc("/register", handlers.Register)
// // http.HandleFunc("/login", handlers.Login)

// http.HandleFunc("/users/get-all", handlers.GetAllUsers)
// http.HandleFunc("/users/post-user", handlers.RegisterAUser)
// http.HandleFunc("/users/update-user", handlers.UpdateAUser)
// http.HandleFunc("/users/delete-user", handlers.DeleteAUser)

// http.Handle(
// 	"/users",
// 	middleware.JWTMiddleware(
// 		http.HandlerFunc(handlers.GetAllUsers),
// 	),
// )
