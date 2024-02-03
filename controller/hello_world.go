package controller

import "net/http"

func HelloWorldController() func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	}

}

func TestController() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("TESTTTT"))
	}
}
