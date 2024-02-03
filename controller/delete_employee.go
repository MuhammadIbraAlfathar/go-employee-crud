package controller

import (
	"database/sql"
	"net/http"
)

func NewDeleteEmployeeController(db *sql.DB) func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		id := request.URL.Query().Get("id")

		_ = request.ParseForm()

		_, err := db.Exec("DELETE from employee WHERE id=?", id)
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/employee", http.StatusMovedPermanently)

	}

}
