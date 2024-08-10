package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	router := http.NewServeMux()

	//Handler functions
	router.HandleFunc("/", readyHandler)

	//Handle image and css files
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.Handle("/", router)
	http.ListenAndServe(":"+port, router)
}
