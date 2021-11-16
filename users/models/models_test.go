package models

import (
	"testing"
)

func TestUserWithCorrectParams(t *testing.T) {
	u := User{
		ID:    "somerandomid",
		Name:  "somerandomname",
		Email: "somerandomtext@example.com",
	}

	if err := validate(u); err != nil {
		t.Errorf("❌ Could not create user with correct params: %v.", err)
	} else {
		t.Log("✅ User created with correct params successfully.")
	}
}

func TestUserWithEmptyName(t *testing.T) {
	u := User{
		ID:    "somerandomid",
		Name:  "",
		Email: "somerandomtext@example.com",
	}

	if err := validate(u); err == nil {
		t.Errorf("❌ User created with empty name: %v.", err)
	} else {
		t.Log("✅ Validator allow user creation as expected.")
	}
}

func TestUserWithLargeName(t *testing.T) {
	u := User{
		ID:    "somerandomid",
		Name:  "123456789123456789123456789123456789",
		Email: "somerandomtext@example.com",
	}

	if err := validate(u); err == nil {
		t.Errorf("❌ User created with large name: %v.", err)
	} else {
		t.Log("✅ Validator stopped the creation as expected.")
	}
}

func TestUserWithEmptyEmail(t *testing.T) {
	u := User{
		ID:    "somerandomid",
		Name:  "somerandomname",
		Email: "",
	}

	if err := validate(u); err == nil {
		t.Errorf("❌ User created with an empty email: %v.", err)
	} else {
		t.Log("✅ Validator stopped the creation as expected.")
	}
}

func TestUserWithInvalidEmail(t *testing.T) {
	u := User{
		ID:    "somerandomid",
		Name:  "somerandomname",
		Email: "testtest.com",
	}

	if err := validate(u); err == nil {
		t.Errorf("❌ User created with an invalid email: %v.", err)
	} else {
		t.Log("✅ Validator stopped the creation as expected.")
	}
}
