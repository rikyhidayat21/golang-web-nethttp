package golang_web_nethttp

import (
	"fmt"
	"net/http"
	"testing"
)

var baseURL = "localhost:4000"

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello world")
	}

	server := http.Server{
		Addr: "localhost:4000",
		Handler: handler,
	}

	// handle errornya
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}


}

func TestServeMux(t *testing.T) {
	// inisialisasi NewServeMux
	mux := http.NewServeMux()

	// handle -> url, functionHandlernya || Sementara HandleFunc() -> Memasukan url, functionHandle(anonymous function) [Baca Functionnya saja]
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi memet")
	})

	mux.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Login page")
	})

	// URL PATTERN
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Image")
	})
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Thumbnails")
	})

	// inisialisasi server
	server := http.Server{Addr: "localhost:4000", Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{Addr: baseURL, Handler: handler}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}