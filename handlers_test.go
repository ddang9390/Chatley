package main

import (
	"bytes"
	"chatley/internal/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	testDBURL := os.Getenv("TEST_DB_URL")

	var err error
	mainDB, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer mainDB.Close()

	// Create test database
	_, err = mainDB.Exec("CREATE DATABASE test_db_name")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to test database
	testDB, err = sql.Open("postgres", testDBURL)
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations on testDB here

	// Run tests
	code := m.Run()

	// Drop test database
	_, err = mainDB.Exec("DROP DATABASE test_db_name")
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code) // Exit with test code
}

func TestCreateUser(t *testing.T) {
	godotenv.Load()
	testDBURL := os.Getenv("TEST_DB_URL")

	db, err := sql.Open("postgres", testDBURL)
	if err != nil {
		fmt.Println(err)
	}
	dbQueries := database.New(db)

	cfg := &apiConfig{DB: dbQueries}

	reqBody := map[string]string{
		"email":    "test@email.com",
		"password": "123456",
	}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	fmt.Println(err)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(createUser(cfg))
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
