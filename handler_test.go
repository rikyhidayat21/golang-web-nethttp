package golang_web_nethttp

import (
	"fmt"
	"net/http"
	"testing"
)

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