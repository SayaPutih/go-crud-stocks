package routes

import (
	"go-crud/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
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
	r := chi.NewRouter()
	r.Get("/users/get-all", handlers.GetAllUsers)
	r.Post("/users/post-user", handlers.RegisterAUser)

	// pakai chi param
	r.Put("/users/update-user/{id}", handlers.UpdateAUser)
	r.Delete("/users/delete-user/{id}", handlers.DeleteAUser)

	r.Get("/stocks/get-all", handlers.GetAllStock)
	r.Post("/stocks/post-stock", handlers.InsertAStock)
	r.Put("/stocks/update-stock/{id}", handlers.UpdateAStock)
	r.Delete("/stocks/delete-stock/{id}", handlers.DeleteAStock)
	return r
}
