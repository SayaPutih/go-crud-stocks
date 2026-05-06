package main

import (
	"log"
	"net/http"

	"go-crud/config"
	"go-crud/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRoutes()

	// http.HandleFunc("/test-live", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Server Running2222s 🚀"))
	// })

	// http.HandleFunc("/test-db", func(w http.ResponseWriter, r *http.Request) {
	// 	err := config.DB.Ping()
	// 	if err != nil {
	// 		http.Error(w, "DB NOT CONNECTED ❌", 500)
	// 		return
	// 	}
	// 	w.Write([]byte("CONNECTED TO THE DATABASE!!! ✅"))
	//})

	log.Println("Server running on :8080")
	log.Println("ROUTES LOADED")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
