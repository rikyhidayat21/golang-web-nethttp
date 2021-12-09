package golang_web_nethttp

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`

	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "SIMPLE", "HELLO HTML TEMPLATE")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "HELLO TEMPLATE GLOB")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "HELLO TEMPLATE GLOB")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
