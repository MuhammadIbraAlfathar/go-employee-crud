package main

import (
	"github.com/MuhammadIbraAlfathar/go-employee-crud.git/database"
	"github.com/MuhammadIbraAlfathar/go-employee-crud.git/routes"
	"net/http"
)

func main() {
	database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server)

	err := http.ListenAndServe("localhost:8080", server)

	if err != nil {
		panic(err)
	}

}
