package main

import (
	"chatley/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func createUser(cfg *apiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		// // Step 1: Parse the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		fmt.Println(user)

		// Step 2: Insert into the database
		ctx := r.Context()
		_, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
			Email:    user.Email,
			Password: user.Password,
		})
		if err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		// // Respond with the created user
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}

}
