package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RamiroCuenca/users-crud-test/database/connection"
	"github.com/RamiroCuenca/users-crud-test/users/models"
	"github.com/google/uuid"
)

// Creates a new user.
// NAME and EMAIL must be sent through request body.
func Create(w http.ResponseWriter, r *http.Request) {
	// Decode NAME and EMAIL from r.body and assign it to u var
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Couldn't decode request body", err)
		return
	}

	// Check that sent values are valid
	err = models.Validate(u)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Sent parameters are invalid", err)
		return
	}

	// Generate UUID
	u.ID = uuid.NewString()

	// Set up query
	q := `INSERT INTO users (id, name, email) VALUES ($1, $2, $3)`

	// Fetch database
	db := connection.GetPostgreClient()

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Error starting db transaction", err)
		return
	}

	// Prepare transaction
	stmt, err := tx.Prepare(q)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Error preparing db transaction", err)
		tx.Rollback()
		return
	}
	defer stmt.Close()

	// Execute the query
	stmt.QueryRow(u.ID, u.Name, u.Email)

	// Commit transaction
	tx.Commit()

	// Encode user
	json, _ := json.Marshal(u)

	data := fmt.Sprintf(`{
		"message": "User created successfully",
		"user": %v
	}`, string(json))

	sendResponse(w, http.StatusOK, []byte(data))
}

// Send response to HTTP requests
func sendResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func errorResponse(w http.ResponseWriter, status int, message string, err error) {
	data := fmt.Sprintf(`{
		"message": "%v: %v"
	}`, message, err)
	sendResponse(w, status, []byte(data))
	return
}
