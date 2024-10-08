package main

import (
	"chatley/internal/database"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

// using config struct to store shared data that handlers will need access to
type apiConfig struct {
	DB        *database.Queries
	jwtSecret string
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
	}
	dbQueries := database.New(db)
	jwtSecret := os.Getenv("JWT_SECRET")

	cfg := &apiConfig{DB: dbQueries, jwtSecret: jwtSecret}
	router := http.NewServeMux()

	//Handler functions
	router.HandleFunc("/", readyHandler)
	router.HandleFunc("POST /users", createUser(cfg))
	router.HandleFunc("GET /users", loginUser(cfg))
	router.HandleFunc("DELETE /users", deleteUser(cfg))
	router.HandleFunc("PATCH /users", updateUser(cfg))

	//Handle image and css files
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.Handle("/", router)
	http.ListenAndServe(":"+port, router)
}
