package golang_web_nethttp

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprint(writer, "Hello world")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	// Untuk mengecek hasilnya
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}