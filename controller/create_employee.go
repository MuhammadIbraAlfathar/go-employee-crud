package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewCreateEmployeeController(db *sql.DB) func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Method == "POST" {

			_ = request.ParseForm()

			name := request.Form["name"][0]
			address := request.Form["address"][0]
			npwp := request.Form["npwp"][0]

			_, err := db.Exec("INSERT INTO employee (name, npwp, address) VALUES (?, ?, ?)", name, npwp, address)
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

		} else if request.Method == "GET" {
			fp := filepath.Join("views", "create.html")
			files, err := template.ParseFiles(fp)
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = files.Execute(writer, nil)
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

	}

}
