package golang_web_nethttp

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	//mux.Handle("/static/", fileServer)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer ))

	server := http.Server{Addr: "localhost:4000", Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// membuat go embed

//go:embed resources
var resources embed.FS

func TestFileServerGoEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{Addr: "localhost:4000", Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
