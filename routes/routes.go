package routes

import (
	"database/sql"
	"github.com/MuhammadIbraAlfathar/go-employee-crud.git/controller"
	"net/http"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.HelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexController(db))
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeController(db))
	server.HandleFunc("/employee/update", controller.NewUpdateEmployeeController(db))
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployeeController(db))
}
