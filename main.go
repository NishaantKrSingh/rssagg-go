package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not found in the environment!") //Immediately stops code and also outputs the proper date and time
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	serv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Printf("Server starting on port %v", portString)
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal("An error occured", err)
	}
}
