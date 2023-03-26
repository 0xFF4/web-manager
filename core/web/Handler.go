package web

import (
	"manager/core/config"
	"net/http"

	"github.com/gorilla/mux"
)

// New will start a new web server instance
func New() {

	// Initialize a brand new router
	handler := mux.NewRouter()

	server := &http.Server{
		Addr:    config.Options.String("server", "host"),
		Handler: handler,
	}

	// Serve the HTTP server using TLS if it is enabled inside the configuration
	server.ListenAndServe()

}
