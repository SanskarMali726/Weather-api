package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)



func main() {
	router := chi.NewRouter()
	router.Get("/weather",wetherhandler)

	err := godotenv.Load()
	if err != nil {
    fmt.Println("Error loading .env file")
	}
	

	fmt.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
		return
	}

}
