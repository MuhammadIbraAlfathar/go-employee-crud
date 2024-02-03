package controller

import "net/http"

func HelloWorldController() func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World"))
	}

}
