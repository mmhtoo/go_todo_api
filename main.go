package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/mmhtoo/go-todo-api/helpers"
)

func main() {
	// load env file
	godotenv.Load(".env")
	// create new router
	router := chi.NewRouter()
	// config cors
	configCors(router)

	// get port from env
	port := loadPortFromEnv()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.NewDataResponse(w, 200, "Success", struct{}{})
	})

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Server is listening on port %s \n", port)
	serverError := server.ListenAndServe()

	if serverError != nil {
		fmt.Printf("Error: %s", serverError)
	}

}

func loadPortFromEnv() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	return port
}

func configCors(router *chi.Mux) {
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept"},
		ExposedHeaders:   []string{"Content-Length", "Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}
