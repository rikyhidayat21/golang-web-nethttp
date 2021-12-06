package golang_web_nethttp

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:4000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}