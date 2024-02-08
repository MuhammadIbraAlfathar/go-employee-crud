package controller

import (
	"database/sql"
	"github.com/MuhammadIbraAlfathar/go-employee-crud.git/model"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewUpdateEmployeeController(db *sql.DB) func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Method == "POST" {

			id := request.URL.Query().Get("id")

			//test

			_ = request.ParseForm()

			name := request.Form["name"][0]
			address := request.Form["address"][0]
			npwp := request.Form["npwp"][0]

			_, err := db.Exec("UPDATE employee set name=?, npwp=?, address=? WHERE id=?", name, npwp, address, id)
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(writer, request, "/employee", http.StatusMovedPermanently)

		} else if request.Method == "GET" {

			id := request.URL.Query().Get("id")
			row := db.QueryRow("SELECT name, npwp, address from employee WHERE id = ?", id)

			var employe model.Employee

			err := row.Scan(
				&employe.Name,
				&employe.NPWP,
				&employe.Address,
			)

			employe.Id = id

			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			fp := filepath.Join("views", "update.html")
			files, err := template.ParseFiles(fp)
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["employee"] = employe

			err = files.Execute(writer, data)
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

	}

}
