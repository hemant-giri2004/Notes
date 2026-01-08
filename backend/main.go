package main

import (
	"backend/db"
	"backend/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load .env file
	//not need to load in production.
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	db.Connect()

	r := mux.NewRouter()
	// 🔥 CORS MUST be the FIRST middleware
	// r.Use(middlewares.CORSMiddleware)

	// Routes
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	})
	// Notes routes
	r.HandleFunc("/notes", handlers.GetNotes).Methods("GET")
	r.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	r.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")

	// 🔥 GLOBAL OPTIONS HANDLER (THIS FIXES EVERYTHING)
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")
	// ✅ CORS configuration (GUARANTEED)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
