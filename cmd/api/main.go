package main


func main(){

}package main

import (
	"log"
	"net/http"
	
	"my-api/internal/handlers"
	"my-api/pkg/config"
	"my-api/pkg/database"
	"my-api/internal/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()
	
	// Initialize database
	db, err := database.Connect(cfg.DB)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	
	// Set up routes
	router := http.NewServeMux()
	
	// User routes
	userHandler := handlers.NewUserHandler(db)
	router.HandleFunc("GET /users", userHandler.GetUsers)
	router.HandleFunc("POST /users", userHandler.CreateUser)
	
	// Product routes
	productHandler := handlers.NewProductHandler(db)
	router.HandleFunc("GET /products", productHandler.GetProducts)
	
	// Add middleware
	handler := middleware.Logging(router)
	handler = middleware.Auth(handler)
	
	// Start server
	log.Printf("Server starting on %s", cfg.Server.Address)
	log.Fatal(http.ListenAndServe(cfg.Server.Address, handler))
}