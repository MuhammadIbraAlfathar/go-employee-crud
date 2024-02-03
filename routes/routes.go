package routes

import (
	"github.com/MuhammadIbraAlfathar/go-employee-crud.git/controller"
	"net/http"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.HelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexController())
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeController())
}
