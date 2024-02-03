package controller

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func NewCreateEmployeeController() func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Method == "POST" {
			_, _ = writer.Write([]byte("Test"))
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
