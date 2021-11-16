package routes

import "github.com/go-chi/chi/v5"

// Fetch app multiplexor
func GetRouter() *chi.Mux {
	mux := chi.NewMux()

	// Path prefix
	pp := "/api/users"

	mux.Post(pp+"/create", nil)
	mux.Get(pp+"/readall", nil)
	mux.Get(pp+"/readbyname", nil)
	mux.Get(pp+"/readbyemail", nil)
	mux.Put(pp+"/update", nil)
	mux.Delete(pp+"/delete", nil)

	return mux
}
