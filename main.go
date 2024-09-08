package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func loadPortFromEnv() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	return port
}

func main() {
	// load env file
	godotenv.Load(".env")
	// create new router
	router := chi.NewRouter()

	// get port from env
	port := loadPortFromEnv()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	
	server := &http.Server{
		Handler: router,
		Addr: ":"+port,
	}

	fmt.Printf("Server is listening on port %s \n",port)
	serverError := server.ListenAndServe()

	if serverError != nil {
		fmt.Printf("Error: %s", serverError)
	}

}