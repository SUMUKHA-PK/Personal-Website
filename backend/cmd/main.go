package main

import (
	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/SUMUKHA-PK.github.io/backend/routing"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	*r = routing.SetupRouting(*r)
	server.Server(
		&server.Data{
			Router: r,
			Port:   "8008",
		},
	)
}
