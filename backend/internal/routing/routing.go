package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouting adds all the routes for the server.
func SetupRouting(r mux.Router) mux.Router {
	r.HandleFunc("/code", codingProblems).Methods(http.MethodGet)
	r.HandleFunc("/blogs", blogs).Methods(http.MethodGet)
	r.HandleFunc("/books", books).Methods(http.MethodGet)
	r.HandleFunc("/tv", tv).Methods(http.MethodGet)
	r.HandleFunc("/trips", trips).Methods(http.MethodGet)
	r.HandleFunc("/movies", movies).Methods(http.MethodGet)
	return r
}
