package main

import (
	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/Personal-Website/backend/internal/routing"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	*r = routing.SetupRouting(*r)
	server.Server(
		&server.Data{
			Router: r,
			Port:   "44440",
		},
	)
}
