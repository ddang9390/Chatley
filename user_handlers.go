package main

import (
	"chatley/internal/database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

		// Encode the password
		encPW, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err1 != nil {
			http.Error(w, "Could not use password", http.StatusInternalServerError)
			return
		}

		// Step 2: Insert into the database
		ctx := r.Context()
		user.ID = uuid.New().String()
		_, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
			ID:       user.ID,
			Email:    user.Email,
			Password: string(encPW),
		})
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		// // Respond with the created user
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}

}
