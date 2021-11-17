package routes

import (
	users "github.com/RamiroCuenca/users-crud-test/users/controllers"
	"github.com/go-chi/chi/v5"
)

// Fetch app multiplexor
func GetRouter() *chi.Mux {
	mux := chi.NewMux()

	// Path prefix
	pp := "/api/users"

	mux.Post(pp+"/create", users.Create)
	mux.Get(pp+"/readall", users.ReadAll)
	mux.Get(pp+"/filter", users.Filter)
	mux.Put(pp+"/update", nil)
	mux.Delete(pp+"/delete", nil)

	return mux
}
