package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {

	t.Run("valid port number", func(t *testing.T) {

		_, err := NewApp(5050)

		assert.NoError(t, err)

	})

	t.Run("invalid port number", func(t *testing.T) {

		_, err := NewApp(0)

		assert.Equal(t, ErrInvalidPort, err)

	})

}

func TestHandleGetEnv(t *testing.T) {

	want := make(map[string]string)

	for _, environ := range os.Environ() {

		pair := strings.SplitN(environ, "=", 2)

		want[pair[0]] = pair[1]
	}

	t.Run("get environment variables with valid request", func(t *testing.T) {

		var got map[string]string
		req, _ := http.NewRequest(http.MethodGet, "/env", nil)
		res := httptest.NewRecorder()

		app, err := NewApp(5050)
		assert.NoError(t, err)

		app.envhandler(res, req)

		assert.Equal(t, http.StatusOK, res.Code, "got %d status code but want status code 200", res.Code)

		json.NewDecoder(res.Body).Decode(&got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected map %v but got %v", want, got)
		}
	})

	t.Run("invalid request (404 status code)", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodGet, "/enviro", nil)
		res := httptest.NewRecorder()

		app, err := NewApp(5050)
		assert.NoError(t, err)

		app.envhandler(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code, "got %d status code but want status code 404", res.Code)
	})

	t.Run("invalid method (500 status code)", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodPost, "/env", nil)
		res := httptest.NewRecorder()

		app, err := NewApp(5050)
		assert.NoError(t, err)

		app.envhandler(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code, "got %d status code but want status code 500", res.Code)
	})

	t.Run("valid JSON content format", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodPost, "/env", nil)
		res := httptest.NewRecorder()

		app, err := NewApp(5050)
		assert.NoError(t, err)

		app.handleGetEnv(res, req)

		assert.Equal(t, "application/json", res.Header().Get("Content-Type"), " got %s content type than json", res.Header().Get("Content-Type"))
	})

}

func TestHandleGetKey(t *testing.T) {

	t.Run("get existing key", func(t *testing.T) {

		got := ""

		key := "home"
		value := "environment"
		want := value

		os.Setenv(key, value)
		defer os.Unsetenv(key)

		req, _ := http.NewRequest("GET", "/env/"+key, nil)
		res := httptest.NewRecorder()

		app, err := NewApp(5050)
		assert.NoError(t, err)

		app.envhandler(res, req)

		assert.Equal(t, http.StatusOK, res.Code, "got %d status code but want status code 200", res.Code)

		json.NewDecoder(res.Body).Decode(&got)

		if got != want {
			t.Errorf("expected body %v, but got %v", want, got)
		}
	})

	t.Run(" get non existing key", func(t *testing.T) {

		want, got := "", ""

		key := "home"

		req, _ := http.NewRequest("GET", "/env/"+key, nil)
		res := httptest.NewRecorder()

		app, err := NewApp(5050)
		assert.NoError(t, err)

		app.envhandler(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code, "got %d status code but want status code 404", res.Code)

		json.NewDecoder(res.Body).Decode(&got)

		if got != want {
			t.Errorf("expected body %v, but got %v", want, got)
		}
	})

	t.Run("invalid method (500 status code)", func(t *testing.T) {

		got := ""

		key := "home"
		value := "environment"

		os.Setenv(key, value)
		defer os.Unsetenv(key)

		req, _ := http.NewRequest("POST", "/env/"+key, nil)
		res := httptest.NewRecorder()

		app, err := NewApp(5050)
		assert.NoError(t, err)

		app.envhandler(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code, "got %d status code but want status code 500", res.Code)

		json.NewDecoder(res.Body).Decode(&got)

		if got != "" {
			t.Errorf("expected body %v, but got %v", "", got)
		}
	})

}
