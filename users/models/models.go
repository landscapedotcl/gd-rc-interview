package models

import (
	"errors"
	"strings"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func validate(u User) error {
	// Check that the Name is not empty.
	if u.Name == "" {
		return errors.New("Name can not be empty")
	}

	// Check that the Name is larger than 3 digits but shorter than 20 digits.
	if len(u.Name) < 4 || len(u.Name) > 20 {
		return errors.New("Name length must be between 4 and 20 digits")
	}

	// Check that the Email is not empty.
	if u.Email == "" {
		return errors.New("Email can not be empty")
	}

	// Check that the Email is valid
	if !strings.Contains(u.Email, "@") {
		return errors.New("Email can not be empty")
	}

	return nil
}
