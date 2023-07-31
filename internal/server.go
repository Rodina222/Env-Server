package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// ErrInvalidPort is returned when the listening port is outside the range (from 1 to 65535)
var ErrInvalidPort = errors.New("invalid port number insert a port from 1 to 65535")

// App is a struct that involves the listening port number of the server
type App struct {
	listenPort int
}

// NewApp  validates the user input port, and if it is valid, it assigns the app's port number to it.
func NewApp(port int) (*App, error) {

	// check input port from user
	if port < 1 || port > 65535 {
		return nil, ErrInvalidPort
	}

	return &App{listenPort: port}, nil
}

// Run function calls registerHandler() method which is an internal method that calls other internal methods for running the app
func (app *App) Run() error {

	return app.registerHandlers()

}

func (app *App) registerHandlers() error {

	router := http.NewServeMux()

	router.HandleFunc("/env", app.envhandler)
	router.HandleFunc("/env/", app.envhandler)

	port := fmt.Sprintf(":%d", app.listenPort)
	fmt.Println("Server started listening on port", port)
	err := http.ListenAndServe(port, router)

	if err == http.ErrServerClosed {
		return fmt.Errorf("server closed suddenly")
	}

	if err != nil {
		return fmt.Errorf("failed to start the server %w", err)
	}

	return nil

}

func (app *App) envhandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		err := errors.New("http method should be get")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	key := strings.TrimPrefix(r.URL.Path, "/env")

	if key == "" {

		app.handleGetEnv(w, r)

	} else {

		app.handleGetKey(w, r)

	}

}

func (app *App) handleGetEnv(w http.ResponseWriter, r *http.Request) {

	envMap := make(map[string]string)

	for _, envVarString := range os.Environ() {

		pair := strings.SplitN(envVarString, "=", 2)

		envMap[pair[0]] = pair[1]
	}

	w.Header().Set("Content-Type", "application/json")

	// encoding the map into json format
	err := json.NewEncoder(w).Encode(envMap)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (app *App) handleGetKey(w http.ResponseWriter, r *http.Request) {

	key := strings.TrimPrefix(r.URL.Path, "/env/")

	value := os.Getenv(key)

	if value == "" {

		w.WriteHeader(http.StatusNotFound)
		return

	}

	// encoding the value into json format
	err := json.NewEncoder(w).Encode(value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
