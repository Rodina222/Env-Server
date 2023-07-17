package internal

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	ErrInvalidPort = errors.New("invalid port number insert a port from 1 to 65535")
	ErrKeyNotFound = errors.New("entered key is not found")
)

type App struct {
	listenPort int
}

func CreateApp(port int) (*App, error) {

	if port < 1 || port > 65535 {
		return nil, ErrInvalidPort
	}

	return &App{listenPort: port}, nil
}

func (app *App) handleEnv(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/env/") {
		key := strings.TrimPrefix(r.URL.Path, "/env/")

		app.handleGetKey(w, r, key)
	} else {
		app.handleGetEnv(w, r)
	}
}

func (app *App) handleGetEnv(w http.ResponseWriter, r *http.Request) {

	for _, environ := range os.Environ() {
		fmt.Fprintln(w, environ)
	}
}

func (app *App) handleGetKey(w http.ResponseWriter, r *http.Request, key string) error {
	value, ok := os.LookupEnv(key)
	if !ok {
		http.NotFound(w, r)
		return ErrKeyNotFound
	}
	fmt.Fprintf(w, "%s=%s\n", key, value)
	return nil
}
