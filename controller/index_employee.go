package controller

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func NewIndexController() func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		fp := filepath.Join("views", "index.html")
		files, err := template.ParseFiles(fp)
		if err != nil {
			writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = files.Execute(writer, nil)
		if err != nil {
			writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}
