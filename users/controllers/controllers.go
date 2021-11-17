package controllers

import (
	"encoding/json"
	"errors"
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

// Read all users.
// No parameters are required.
func ReadAll(w http.ResponseWriter, r *http.Request) {
	// Array of users where we are going to store all fetched values
	var arr []models.User

	// Set up query
	q := `SELECT * FROM users`

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
	rows, err := stmt.Query()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Error fetching users", err)
		tx.Rollback()
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User

		err = rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, "Problem scanning fetched users", err)
			tx.Rollback()
			return
		}

		arr = append(arr, u)
	}

	// Commit transaction
	tx.Commit()

	// Encode user
	json, _ := json.Marshal(arr)

	data := fmt.Sprintf(`{
		"message": "%v users where fetched",
		"user": %v
	}`, len(arr), string(json))

	sendResponse(w, http.StatusOK, []byte(data))
}

// Filter by NAME and/or EMAIL
// NAME and/or EMAIL must be sent through the body
func Filter(w http.ResponseWriter, r *http.Request) {
	// Decode NAME and EMAIL from r.body and assign it to u var
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		u.Name = ""
		u.Email = ""
	}

	// Array of users where we are going to store all fetched values
	var arr []models.User

	// Set up query
	q := `SELECT * FROM users 
	WHERE name LIKE '%' || $1 || '%' 
	AND email LIKE '%' || $2 || '%';`

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
	rows, err := stmt.Query(u.Name, u.Email)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Error fetching users", err)
		tx.Rollback()
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, "Problem scanning fetched users", err)
			tx.Rollback()
			return
		}

		arr = append(arr, user)
	}

	// Commit transaction
	tx.Commit()

	// Encode user
	json, _ := json.Marshal(arr)

	data := fmt.Sprintf(`{
		"message": "%v users where fetched",
		"user": %v
	}`, len(arr), string(json))

	sendResponse(w, http.StatusOK, []byte(data))
}

// Update NAME & EMAIL from user
// ID must be sent as query parameter
func Update(w http.ResponseWriter, r *http.Request) {
	// Fetch id from query params
	id := r.URL.Query().Get("id")
	if id == "" {
		errorResponse(w, http.StatusBadRequest, "Bad request", errors.New("'id' parameter must be sent through URL"))
		return
	}

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

	u.ID = id

	// Set up query
	q := `UPDATE users SET name = $2, email = $3 
	WHERE id = $1`

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
		"message": "User updated successfully",
		"user": %v
	}`, string(json))

	sendResponse(w, http.StatusOK, []byte(data))
}

// Delete user by id
// ID must be sent as query parameter
func Delete(w http.ResponseWriter, r *http.Request) {
	// Fetch the id from queryparams
	id := r.URL.Query().Get("id")
	if id == "" {
		errorResponse(w, http.StatusBadRequest, "Bad request", errors.New("'id' parameter must be sent through URL"))
		return
	}

	// Create a user object where store the data from deleted user
	var u models.User

	// Set up query
	q := `DELETE FROM users WHERE id = $1 RETURNING *`

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
	err = stmt.QueryRow(id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Couldn't fetch any user with provided 'id'", err)
		tx.Rollback()
		return
	}

	// Commit transaction
	tx.Commit()

	// Encode user
	json, _ := json.Marshal(u)

	data := fmt.Sprintf(`{
		"message": "User deleted successfully",
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
