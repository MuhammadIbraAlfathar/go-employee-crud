package controller

import (
	"database/sql"
	"github.com/MuhammadIbraAlfathar/go-employee-crud.git/model"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewIndexController(db *sql.DB) func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		var employees []model.Employee

		rows, err := db.Query("SELECT id, name, npwp, address from employee")

		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		for rows.Next() {
			var employee model.Employee

			err := rows.Scan(
				&employee.Id,
				&employee.Name,
				&employee.NPWP,
				&employee.Address,
			)

			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			employees = append(employees, employee)
		}

		fp := filepath.Join("views", "index.html")
		files, err := template.ParseFiles(fp)
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["employees"] = employees

		err = files.Execute(writer, data)
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}
