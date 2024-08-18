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
		// Step 1: Parse the request body
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

func loginUser(cfg *apiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		fmt.Println(r.Body)
		// Parse the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		fmt.Println(user)

		ctx := r.Context()
		users, err := cfg.DB.GetAllUsers(ctx)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error getting users", http.StatusInternalServerError)
			return
		}
		foundEmail := false
		var foundUser User
		for _, u := range users {
			if u.Email == user.Email {
				foundEmail = true
				foundUser.Email = u.Email
				foundUser.ID = u.ID
				foundUser.Password = u.Password

				break
			}
		}
		if !foundEmail {
			http.Error(w, "Could not find user", http.StatusUnauthorized)
			return
		}

		// Decrypt found user's password and compare it
		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
		if err != nil {
			fmt.Printf("Input PW:%s, Actual PW:%s\n\n", user.Password, foundUser.Password)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		token := jwtCreation(user, cfg.jwtSecret)

		response := map[string]interface{}{
			"id":    foundUser.ID,
			"email": foundUser.Email,
			"token": token,
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)

	}
}

func deleteUser(cfg *apiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// JWT validation for user authentication
		_, err := jwtValidate(r, cfg.jwtSecret)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		var user User
		// Parse the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		u, err := cfg.DB.GetOneUser(ctx, user.ID)
		if err != nil {
			http.Error(w, "Error deleting user", http.StatusInternalServerError)
			return
		}
		var foundUser User
		foundUser.Email = u.Email
		foundUser.ID = u.ID
		foundUser.Password = u.Password

		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		err = cfg.DB.DeleteUser(ctx, database.DeleteUserParams{
			Email:    user.Email,
			Password: foundUser.Password,
		})
		if err != nil {
			http.Error(w, "Error deleting user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
	}
}

func updateUser(cfg *apiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// JWT validation for user authentication
		_, err := jwtValidate(r, cfg.jwtSecret)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		var user User
		// Parse the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Encode the password
		encPW, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Could not use password", http.StatusInternalServerError)
			return
		}

		ctx := r.Context()
		err = cfg.DB.UpdateUser(ctx, database.UpdateUserParams{
			ID:       user.ID,
			Email:    user.Email,
			Password: string(encPW),
		})
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
	}
}
